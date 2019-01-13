package main

import (
	"flag"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/LauncherJob/Worker"
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
	flag.StringVar(&confFilePath,"config","./worker.json","Config worker file path")
	flag.Parse()
}

func main() {
	var(
		err error
	)

	initEnv()

	initArgs()

	if err = worker.InitConfig(confFilePath); err!=nil{
		goto ERR
	}

	if err = worker.InitJobMgr(); err!=nil{
		goto ERR
	}

	for{
		time.Sleep(1*time.Second)
	}

ERR:
	log.Println(err)

}
