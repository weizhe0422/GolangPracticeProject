package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"log"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
		kvPair *mvccpb.KeyValue
	)
	config = clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout: 5 *time.Second,
	}

	if client, err = clientv3.New(config); err!=nil{
		log.Println("failed to connect to etcd:", err)
		return
	}

	//Put a new key-value into etcd
	kv = clientv3.NewKV(client)
	if putResp,err = kv.Put(context.TODO(),"/cron/jobs/task1","hello,task 1",clientv3.WithPrevKV());err!=nil{
		log.Println("failed to put: ",err)
	}else{
		fmt.Println("Revision: ",putResp.Header.Revision)
		if putResp.PrevKv!=nil{
			fmt.Println("Prev. Value:",string(putResp.PrevKv.Value))
		}
	}
	if putResp,err = kv.Put(context.TODO(),"/cron/jobs/task2","hello,task 2",clientv3.WithPrevKV());err!=nil{
		log.Println("failed to put: ",err)
		return
	}else{
		fmt.Println("Revision: ",putResp.Header.Revision)
		if putResp.PrevKv!=nil{
			fmt.Println("Prev. Value:",string(putResp.PrevKv.Value))
		}
	}

	//Get a value with a specific key
	if getResp,err = kv.Get(context.TODO(),"/cron/jobs",clientv3.WithPrefix());err!=nil{
		log.Println("failed to get value:",err)
		return
	}else{
		for _,value := range getResp.Kvs{
			fmt.Println(" Content:",value)
		}

	}

	//Delete a Key-value pair
	if delResp, err = kv.Delete(context.TODO(),"/cron/jobs/task2",clientv3.WithPrevKV(),clientv3.WithLimit(0));err!=nil{
		log.Println(err)
		return
	}
	//get previous key-value before deleted
	if len(delResp.PrevKvs) != 0 {
		for _,kvPair = range delResp.PrevKvs {
			fmt.Println("Delete :",string(kvPair.Key), string(kvPair.Value))
		}
	}


}
