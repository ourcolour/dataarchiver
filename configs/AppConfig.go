package configs

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"xjh.com/dataarchiver/constants/errs"
)

const (
	DUMP_FILE_EXT = ".dmp"
	ZIP_FILE_EXT  = ".zip"
)

var MYSQL_HOME = ""         //os.Getenv("MYSQL_HOME")
var MYSQL_BIN_DIR_PATH = "" //filepath.Join(MYSQL_HOME, "bin")
var MYSQLDUMP_PATH = ""     //filepath.Join(MYSQL_BIN_DIR_PATH, "mysqldump")

func CheckMysqlDump() error {
	MYSQL_HOME = os.Getenv("MYSQL_HOME")
	MYSQL_BIN_DIR_PATH = filepath.Join(MYSQL_HOME, "bin")
	MYSQLDUMP_PATH = filepath.Join(MYSQL_BIN_DIR_PATH, "mysqldump")
	if "windows" == strings.ToLower(runtime.GOOS) {
		MYSQLDUMP_PATH += ".exe"
	}

	log.Printf("OS PLAT: %s", runtime.GOOS)
	log.Printf("OS ARCH: %s", runtime.GOARCH)
	log.Printf("MYSQL_HOME: %s", MYSQL_HOME)
	log.Printf("MYSQL_BIN : %s", MYSQL_BIN_DIR_PATH)
	log.Printf("MYSQL_DUMP: %s", MYSQLDUMP_PATH)

	// 检查 mysqldump 是否存在
	fileInfo, err := os.Stat(MYSQLDUMP_PATH)
	if nil != err {
		return err
	}
	if fileInfo.IsDir() {
		return errs.ERR_INVALID_PARAMETERS
	}

	return nil
}
