package bll

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"xjh.com/dataarchiver/configs"
	"xjh.com/dataarchiver/constants/errs"
	"xjh.com/dataarchiver/utils"
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
func Dump(host string, port int, user, password, dbName, tableName, targetDirPath string) (string, error) {
	// 检查 mysqldump 是否存在
	fileInfo, err := os.Stat(configs.MYSQLDUMP_PATH)
	if nil != err {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	var cmd *exec.Cmd
	if tableName == "" {
		cmd = exec.Command(configs.MYSQLDUMP_PATH, "--opt", "-h"+host, "-P"+strconv.Itoa(port), "-u"+user, "-p"+password, dbName)
	} else {
		cmd = exec.Command(configs.MYSQLDUMP_PATH, "--opt", "-h"+host, "-P"+strconv.Itoa(port), "-u"+user, "-p"+password, dbName, tableName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return "", err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// 生成带时间戳的文件名
	backupFileName := GenerateDumpFileBaseName(dbName, tableName)
	var backupPath string = filepath.Join(targetDirPath, backupFileName)

	// 创建目录
	if err = createDirectoryIfNotExists(targetDirPath); nil != err {
		log.Panic(err)
		return "", err
	}
	// 写文件
	if err = ioutil.WriteFile(backupPath, bytes, 0644); err != nil {
		log.Panic(err)
		return "", err
	}

	return backupPath, nil
}

func GenerateDumpFileBaseName(dbname string, tableName string) string {
	return GenerateDumpFileBaseNameByTime(dbname, tableName, time.Now())
}

func GenerateDumpFileBaseNameByTime(dbName string, tableName string, t time.Time) string {
	timeString := GenerateTimeString(t)

	var result string
	if "" == strings.TrimSpace(tableName) {
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
