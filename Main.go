package dataarchiver

import (
	"log"
	"xjh.com/dataarchiver/cmd"
	"xjh.com/dataarchiver/configs"
)

func main() {
	if err := configs.CheckMysqlDump(); nil != err {
		log.Printf("请设置环境变量 {MYSQL_HOME}", err)
		log.Printf("%s\n", err)
	}

	// backup -H "127.0.0.1" -P 3306 -u root -p 123456 -d northworld -o "./dmp"
	// clearold -o "./dmp" -d 31
	cmd.Execute()
}
