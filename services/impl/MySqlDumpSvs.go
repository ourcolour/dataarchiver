package impl

import (
	"github.com/ourcolour/dataarchiver/services"
	"github.com/ourcolour/dataarchiver/services/impl/bll"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MySqlDumpSvs struct {
}

func NewMySqlDumpSvs() services.IDumpSvs {
	return services.IDumpSvs(&MySqlDumpSvs{})
}

func (this *MySqlDumpSvs) Backup(host string, port int, user string, pass string, dbName string, tableName string, outputDirPath string, compress bool, args ...string) (string, error) {
	log.Printf("备份至目录：`%s`\n", outputDirPath)

	dumpFilePath, err := bll.Dump(host, port, user, pass, dbName, tableName, outputDirPath, compress, args...)
	//dumpFilePath, err := this.generateDumpFile(host, port, user, pass, dbName, tableName, outputDirPath, compress)

	log.Printf("备份完毕：`%s`\n", dumpFilePath)

	return dumpFilePath, err

	/*
		// 压缩包文件名
		var zipName string
		if zipName, err = bll.GenerateZipNameByDumpFilePath(dumpFilePath); nil != err {
			log.Panicf("%s\n", err)
			return err
		}
		zipPath := filepath.Join(outputDirPath, zipName)

		// 压缩文件
		if err = this.zipDumpFile([]string{dumpFilePath}, zipPath); nil != err {
			log.Panicf("%s\n", err)
			return err
		}

		// 删除 dump 文件
		if err = this.deleteDumpFile([]string{dumpFilePath}); nil != err {
			log.Panicf("%s\n", err)
			return err
		}
	*/
}

/*
func (this *MySqlDumpSvs) generateDumpFile(host string, port int, user string, pass string, dbName string, tableName string, outputDirPath string, compress bool) (string, error) {
	return bll.Dump(host, port, user, pass, dbName, tableName, outputDirPath, compress)
}

func (this *MySqlDumpSvs) zipDumpFile(srcPathArray []string, dstPath string) error {
	log.Printf("压缩文件：`%s`\n", dstPath)

	zipFile := archiver.NewZip()
	defer zipFile.Close()
	err := zipFile.Archive(srcPathArray, dstPath)

	return err
}

func (this *MySqlDumpSvs) deleteDumpFile(srcPathArray []string) error {
	for _, path := range srcPathArray {
		if err := os.Remove(path); nil != err {
			log.Panicf("文件不存在：`%s`\n", err)
			return err
		} else {
			log.Printf("删除文件：`%s`\n", path)
		}
	}

	return nil
}
*/

func (this *MySqlDumpSvs) DeleteOldArchiveByOverDayCount(dirPath string, overDays int) ([]string, error) {
	var (
		result []string = make([]string, 0)
		err    error
	)

	// 换算过期时间
	overHours := time.Hour * time.Duration(24) * time.Duration(overDays)

	// 遍历文件夹
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}

		// 只看文件
		if info.IsDir() {
			return nil
		}

		// 比较修改时间
		if info.ModTime().Add(overHours).Before(time.Now()) {
			// 过滤文件
			//extName := filepath.Ext(info.Name())

			if true { // configs.ZIP_FILE_EXT == strings.ToLower(strings.TrimSpace(extName)) || configs.GZIP_FILE_EXT == strings.ToLower(strings.TrimSpace(extName)) {
				if err := os.Remove(path); nil != err {
					log.Printf("删除文件失败：`%s`\n", err)
				} else {
					log.Printf("删除过期文件：%s\n", path)
					result = append(result, path)
				}
			} else {
				log.Printf("跳过无效文件：%s\n", path)
			}
		}

		return nil
	})

	return result, err
}
