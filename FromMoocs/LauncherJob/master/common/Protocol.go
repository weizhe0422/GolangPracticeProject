package common

import (
	"encoding/json"
	"log"
)

type Job struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	CronExpr string `json:"cronExpr"`
}

type HttpResponse struct {
	ErrNo int `json:"errNo"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}


func BuildResponse(errNo int, msg string,data interface{})(respContent []byte, err error){
	var(
		httpResp HttpResponse
	)

	httpResp.ErrNo=errNo
	httpResp.Msg=msg
	httpResp.Data=data

	log.Println("data",data)
	log.Println("httpResp",httpResp)
	respContent, err = json.Marshal(httpResp)

	return
}