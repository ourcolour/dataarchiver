package services

type IDumpSvs interface {
	Backup(host string, port int, user string, pass string, dbName string, tableName string, outputDirPath string, compress bool, args ...string) (string, error)
	DeleteOldArchiveByOverDayCount(dirPath string, overDays int) ([]string, error)
}
