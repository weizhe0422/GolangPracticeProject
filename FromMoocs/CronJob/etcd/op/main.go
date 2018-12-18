package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	var (
		cfg clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putop clientv3.Op
		getop clientv3.Op
		opResp clientv3.OpResponse
	)

	cfg = clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5 *time.Second,
	}
	if client, err = clientv3.New(cfg); err!=nil{
		log.Println(err)
		return
	}

	//Create a kv operation
	kv = clientv3.NewKV(client)

	//Create a new OP with PUT request, and use kv operation to execute PUT
	putop = clientv3.OpPut("/cron/jobs/job4","i am job4")
	if opResp, err = kv.Do(context.TODO(),putop); err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("Put Revision:", opResp.Put().Header.Revision)

	//Create a new op with GET request, and use kv operation to execute GET
	getop = clientv3.OpGet("/cron/jobs/job4")
	if opResp, err = kv.Do(context.TODO(),getop); err!=nil{
		log.Println(err)
		return
	}
	//Judge the result of GET request, if the count of result bugger than zero, then print out the ModRevision and the current value
	if opResp.Get().Count > 0 {
		fmt.Println("Revision:", opResp.Get().Kvs[0].ModRevision,"Value: ",string(opResp.Get().Kvs[0].Value))
	}

	
}
