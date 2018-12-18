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
	var(
		cfg clientv3.Config
		client *clientv3.Client
		err error
		watcher clientv3.Watcher
		kv clientv3.KV
		getResp *clientv3.GetResponse
		startRevisionId int64
		watchRespChan <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
		)

	cfg = clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5*time.Second,
	}
	if client, err = clientv3.New(cfg); err!=nil{
		log.Println(err)
		return
	}

	//Simulate a kv keep modified
	kv = clientv3.NewKV(client)
	go func(){
		for{
			kv.Put(context.TODO(),"/cron/jobs/job7","i am job7")
			kv.Delete(context.TODO(),"/cron/jobs/job7")
			time.Sleep(1*time.Second)
		}
	}()

	//get current revision, and set the start revision number
	if getResp, err = kv.Get(context.TODO(),"/cron/jobs/job7"); err!=nil{
		log.Println(err)
		return
	}
	startRevisionId = getResp.Header.Revision + 1

	//New a watcher
	watcher = clientv3.NewWatcher(client)
	//start to watch "/cron/jobs/job7"
	watchRespChan = watcher.Watch(context.TODO(),"/cron/jobs/job7",clientv3.WithRev(startRevisionId))
	for watchResp = range watchRespChan{
		for _, event = range watchResp.Events{
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("Modified: ",string(event.Kv.Value),"Revision: ",event.Kv.CreateRevision,event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("Delete: ",event.Kv.ModRevision)

			}
		}
	}

}
