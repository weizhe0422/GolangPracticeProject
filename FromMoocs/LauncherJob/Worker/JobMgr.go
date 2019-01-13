package worker

import (
	"context"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/LauncherJob/master/common"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

type JobMgr struct {
	jobClient *clientv3.Client
	jobLease  clientv3.Lease
	jobKV     clientv3.KV
}

var (
	G_JobMgr *JobMgr
)

func(j *JobMgr) watchJobs() (err error){
	var(
		getResp *clientv3.GetResponse
		kvPair *mvccpb.KeyValue
		job *common.Job
	)
	if getResp, err = j.jobKV.Get(context.TODO(),common.Job_Save_Dir,clientv3.WithPrefix()); err!=nil{
		return
	}

	for _,kvPair := range getResp.Kvs{
		if job, err = common.UnpackJob(kvPair.Value); err == nil{
			//TODO: send job to scheduler
		}
	}




}

func InitJobMgr() (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)

	config = clientv3.Config{
		Endpoints:   G_Config.EtcdEndPoints,
		DialTimeout: time.Duration(G_Config.EtcdDDialTimeOut) * time.Millisecond,
	}

	if client, err = clientv3.New(config); err != nil {
		return
	}

	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	G_JobMgr = &JobMgr{
		jobClient: client,
		jobLease:  lease,
		jobKV:     kv,
	}

	return

}