package main

import (
	"flag"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/LauncherJob/master"
	"log"
	"runtime"
	"time"
)

var(
	confFilePath string
)

func initEnv(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func initArgs(){
	flag.StringVar(&confFilePath,"config","./master.json","Config file path")
	flag.Parse()
}

func main() {
	var(
		err error
	)

	initEnv()

	initArgs()

	if err = master.InitConfig(confFilePath); err!=nil{
		goto ERR
	}

	if err = master.InitJobMgr(); err!=nil{
		goto ERR
	}

	if err = master.InitApiServer(); err !=nil{
		log.Println("InitApiServer",err)
		goto ERR
	}


	for{
		time.Sleep(1*time.Second)
	}

	ERR:
		log.Println(err)
	
}
