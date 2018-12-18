package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main(){
	var(
		cfg clientv3.Config
		client *clientv3.Client
		err error
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseid clientv3.LeaseID
		kv clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		leaseKeepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
		leaseKeepAlive *clientv3.LeaseKeepAliveResponse

	)
	//Create a client with setting
	cfg = clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout:5*time.Second,
	}
	if client, err = clientv3.New(cfg); err!=nil{
		log.Println(err)
		return
	}
	log.Println("Success to create a client.")

	//New a lease with 10 seconds, and get the lease id
	lease = clientv3.NewLease(client)
	if leaseGrantResp, err = lease.Grant(context.TODO(),10); err!=nil{
		log.Println(err)
		return
	}
	leaseid = leaseGrantResp.ID
	log.Println("Success to create a lease: ",leaseid)


	//Auto extend lease
	/* Auto extend lease function disabled after 5 seconds, so the lease will be alive in 15 seconds.
	ctx,_ := context.WithTimeout(context.TODO(),5*time.Second)
	if leaseKeepAliveCh, err = lease.KeepAlive(ctx,leaseid); err!=nil{
		log.Println(err)
		return
	}
	*/
	if leaseKeepAliveCh, err = lease.KeepAlive(context.TODO(),leaseid); err!=nil{
		log.Println(err)
		return
	}
	go func(){
		for{
			select {
			case leaseKeepAlive = <-leaseKeepAliveCh:
				if leaseKeepAliveCh == nil{
					log.Println("fail to auto extend lease")
					goto END
				}else{
					log.Println("success to auto extend lease:",leaseKeepAlive.ID)
				}
			}
		}
		END:
	}()



	//Put a new KV with lease
	kv = clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(),"/cronjob/task/job1","Hello,WeiZhe",clientv3.WithLease(leaseid));err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("Success to put a kv: ",putResp.Header.Revision)

	//Create a for loop to check whether the lease is expired or not
	for{
		if getResp, err = kv.Get(context.TODO(),"/cronjob/task/job1"); err!=nil{
			log.Println("fail to get kv:", err)
		}

		if getResp.Count==0{
			fmt.Println(leaseid, "is expired.")
			break
		}
		fmt.Println(leaseid, " is not expired yet: ", getResp.Kvs)
		time.Sleep(2*time.Second)
	}


}