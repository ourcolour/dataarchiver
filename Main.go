package main

import (
	"fmt"
	"github.com/ourcolour/dataarchiver/cmd"
	"github.com/ourcolour/dataarchiver/configs"
	"github.com/ourcolour/dataarchiver/services/impl/bll"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//unittest1()
	//return
	//unittest()
	//log.Printf("%s\n", filepath.Base("/root/1/t.txt"))
	//return

	if err := configs.CheckMysqlDump(); nil != err {
		log.Printf("请设置环境变量 {MYSQL_HOME}\n")
		log.Printf("%s\n", err)
	}

	// backup -H "127.0.0.1" -P 3306 -u root -p 123456 -d chinaloyalty -o "./dmp"
	// clearold -o "./dmp" -d 31
	cmd.Execute()
}

func unittest1() {
	//srcPath := "/Users/cc/Desktop/u2.py"
	//dstPath := "/Volumes/Data/u2.zip"
	//comment := "我是注释内容"
	//
	//videoSvs := &impl.VideoSvs{}
	//str, err := videoSvs.Archive(srcPath, dstPath, comment, true)
	//if nil != err {
	//	log.Panicf("%s\n", err.Error())
	//} else {
	//	log.Printf("%s\n", str)
	//}
}

func unittest() {
	//filePath := "/Volumes/Data/阿丽塔：战斗天使.Alita.Battle.Angel.2019.HD1080P.X264.AAC.CHS.mp4"
	filePath := "/Volumes/Data/dd/129-1.mkv"
	file, err := os.Open(filePath)
	if nil != err {
		log.Panicln(err)
	}
	defer file.Close()

	_, err = file.Stat()
	if nil != err {
		log.Panicln(err)
	}

	len, err := bll.WriteGZip(
		io.Reader(file),
		filePath,
		filepath.Join("/Volumes/Data/", fmt.Sprintf("%s.%s", filepath.Base(filePath), "gz")),
	)
	if nil != err {
		log.Panicln(err)
	}
	log.Printf("FFF: %d", len)

}
