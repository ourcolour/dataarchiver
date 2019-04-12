package bll

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/ourcolour/dataarchiver/configs"
	"github.com/ourcolour/dataarchiver/constants/errs"
	"github.com/ourcolour/dataarchiver/utils"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func createDirectoryIfNotExists(targetDirPath string) error {
	fileInfo, err := os.Stat(targetDirPath)
	if nil != err {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(targetDirPath, os.ModePerm); nil != err {
				log.Printf("Failed to create folder `%s`.\n", targetDirPath)
				return err
			}
		} else {
			log.Printf("%s\n", err)
			return err
		}
	} else {
		if !fileInfo.IsDir() {
			log.Printf("Target path existed but not a valid dir `%s`.\n", targetDirPath)
			return errs.ERR_DUPLICATED
		}
	}

	return nil
}

/**
 *
 * 备份MySql数据库
 * @param 	host: 			数据库地址: localhost
 * @param 	port:			端口: 3306
 * @param 	user:			用户名: root
 * @param 	password:		密码: root
 * @param 	databaseName:	需要被分的数据库名: test
 * @param 	tableName:		需要备份的表名: user
 * @param 	sqlPath:		备份SQL存储路径: D:/backup/test/
 * @return 	backupPath
 *
 */
func Dump(host string, port int, user, password, dbName, tableName, targetDirPath string, compress bool) (string, error) {
	// 检查 mysqldump 是否存在
	fileInfo, err := os.Stat(configs.MYSQLDUMP_PATH)
	if nil != err {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	var cmd *exec.Cmd
	if strings.TrimSpace(dbName) == "" {
		cmd = exec.Command(configs.MYSQLDUMP_PATH, "--opt", "-h"+host, "-P"+strconv.Itoa(port), "-u"+user, "-p"+password, "--all-databases")
	} else if strings.TrimSpace(tableName) == "" {
		cmd = exec.Command(configs.MYSQLDUMP_PATH, "--opt", "-h"+host, "-P"+strconv.Itoa(port), "-u"+user, "-p"+password, dbName)
	} else {
		cmd = exec.Command(configs.MYSQLDUMP_PATH, "--opt", "-h"+host, "-P"+strconv.Itoa(port), "-u"+user, "-p"+password, dbName, tableName)
	}

	// 生成带时间戳的文件名
	var backupFileName string
	if !compress {
		backupFileName = GenerateDumpFileBaseName(dbName, tableName)
	} else {
		backupFileName = GenerateGzipFileBaseName(dbName, tableName)
	}
	var backupPath string = filepath.Join(targetDirPath, backupFileName)

	// 创建目录
	if err = createDirectoryIfNotExists(targetDirPath); nil != err {
		log.Panic(err)
		return "", err
	}

	// 标准输出流校验
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return "", err
	}

	// 输出不同文件类型
	if compress {
		_, err = writeGZip(stdout, backupFileName, backupPath)
	} else {
		_, err = writeDump(stdout, backupFileName, backupPath)
	}

	return backupPath, err
}

func WriteGZip(reader io.Reader, srcPath string, dstPath string) (int, error) {
	return writeGZip(reader, srcPath, dstPath)
}

const BUFFER_SIZE = 1024 * 4

func writeDump(reader io.ReadCloser, srcPath string, dstPath string) (int, error) {
	srcFileName := filepath.Base(srcPath)

	// 文件输出流
	saveFileObj, err := os.Create(dstPath)
	if nil != err {
		log.Fatal(err)
		return 0, err
	}
	defer saveFileObj.Close()

	// 创建 writer
	writer := bufio.NewWriter(saveFileObj)

	// 从输入流读取数据
	total := 0
	for {
		buf := make([]byte, BUFFER_SIZE)
		length, err := reader.Read(buf)

		if nil == err { // 没有错误
			length, err = writer.Write(buf) // 写到压缩文件
			if nil != err {
				log.Fatalln(err)
				return 0, err
			}
			writer.Flush()

			// 计算已输出长度
			total += length
		} else if io.EOF == err { // EOF
			log.Printf("Total %d bytes were written into target file.\n", total, srcFileName)
			break
		} else {
			log.Fatal(err)
			return 0, err
			break
		}
	}

	return total, err
}

func writeGZip(reader io.Reader, srcPath string, dstPath string) (int, error) {
	srcFileName := filepath.Base(srcPath)

	// 文件输出流
	saveFileObj, err := os.Create(dstPath)
	if nil != err {
		log.Fatal(err)
		return 0, err
	}
	defer saveFileObj.Close()
	// 源文件大小
	//srcFileStat, _ := os.Stat(srcPath)
	//srcFileLength := float64(srcFileStat.Size())

	// 创建 gzip
	writer := gzip.NewWriter(saveFileObj)
	defer writer.Close()
	writer.Header.Name = srcFileName
	writer.ModTime = time.Now()

	// 从输入流读取数据
	total := 0
	for {
		buf := make([]byte, BUFFER_SIZE)
		length, err := reader.Read(buf)

		if nil == err { // 没有错误
			length, err = writer.Write(buf) // 写到压缩文件
			if nil != err {
				log.Fatalln(err)
				return 0, err
			}
			writer.Flush()

			// 计算已输出长度
			total += length
			// 计算完成百分比
			//cur, _ := strconv.ParseInt(strconv.Itoa(total), 10, 64)
			//percent := int64(math.Floor(float64(cur) / srcFileLength * 100))
			//log.Printf("[%3d%%] Wrote %d of %d byte(s) to file \"%s\"\n", percent, length, total, srcFileName)
		} else if io.EOF == err { // EOF
			log.Printf("Total %d bytes were written to file \"%s\"\n", total, srcFileName)
			break
		} else {
			log.Fatal(err)
			return 0, err
			break
		}
	}

	return total, err
}

func GenerateGzipFileBaseName(dbname string, tableName string) string {
	return strings.Replace(GenerateDumpFileBaseName(dbname, tableName), configs.DUMP_FILE_EXT, "", -1) + configs.GZIP_FILE_EXT
}

func GenerateDumpFileBaseName(dbname string, tableName string) string {
	return GenerateDumpFileBaseNameByTime(dbname, tableName, time.Now())
}

func GenerateDumpFileBaseNameByTime(dbName string, tableName string, t time.Time) string {
	timeString := GenerateTimeString(t)

	var result string
	if "" == strings.TrimSpace(dbName) {
		result = strings.Join([]string{"all", timeString + configs.DUMP_FILE_EXT}, "-")
	} else if "" == strings.TrimSpace(tableName) {
		result = strings.Join([]string{dbName, timeString + configs.DUMP_FILE_EXT}, "-")
	} else {
		result = strings.Join([]string{dbName, tableName, timeString + configs.DUMP_FILE_EXT}, "-")
	}

	return result
}

func GenerateTimeString(t time.Time) string {
	timeFormat := "20060102_150405_123"
	timeString := utils.FormatDatetime(t, timeFormat)
	fileName := fmt.Sprintf("%s", timeString)
	return fileName
}

func GenerateZipNameByDumpFilePath(dumpFilePath string) (string, error) {
	if "" == strings.TrimSpace(dumpFilePath) {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	fileName := filepath.Base(dumpFilePath)
	extName := filepath.Ext(dumpFilePath)
	zipName := strings.Replace(fileName, extName, configs.ZIP_FILE_EXT, -1)

	return zipName, nil
}
