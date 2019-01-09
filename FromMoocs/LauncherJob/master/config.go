package master

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ApiSvrPrt          int `json:"apiSvrPrt"`
	ApiSvrReadTimeOut  int `json:"apiSvrReadTimeOut"`
	ApiSvrWriteTimeOut int `json:"apiSvrWriteTimeOut"`
	EtcdEndPoints []string`json:"etcdEndPoints"`
	EtcdDDialTimeOut int `json:"etcdDDialTimeOut"`
}

var (
	G_Config *Config
)

func InitConfig(fileName string) (err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(fileName); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	G_Config = &conf

	log.Println(G_Config)
	return
}
