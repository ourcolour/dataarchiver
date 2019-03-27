package main

import (
	"github.com/ourcolour/dataarchiver/cmd"
	"github.com/ourcolour/dataarchiver/configs"
	"log"
)

func main() {
	////filePath := "/Volumes/Data/阿丽塔：战斗天使.Alita.Battle.Angel.2019.HD1080P.X264.AAC.CHS.mp4"
	//filePath := "/Volumes/Data/dd/129-1.mkv"
	//file, err := os.Open(filePath)
	//if nil != err {
	//	log.Panicln(err)
	//}
	//defer file.Close()
	//
	//_, err = file.Stat()
	//if nil != err {
	//	log.Panicln(err)
	//}
	//
	//len, err := bll.WriteGZip(
	//	io.ReadCloser(file),
	//	filePath,
	//	filepath.Join("/Volumes/Data/", fmt.Sprintf("%s.%s", filepath.Base(filePath), "gzip")),
	//)
	//if nil != err {
	//	log.Panicln(err)
	//}
	//log.Printf("FFF: %d", len)
	//
	//return

	if err := configs.CheckMysqlDump(); nil != err {
		log.Printf("请设置环境变量 {MYSQL_HOME}", err)
		log.Printf("%s\n", err)
	}

	// backup -H "127.0.0.1" -P 3306 -u root -p 123456 -d chinaloyalty -o "./dmp"
	// clearold -o "./dmp" -d 31
	cmd.Execute()
}
