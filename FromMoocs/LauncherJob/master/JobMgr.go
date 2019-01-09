package master

import (
	"context"
	"encoding/json"
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


func (j *JobMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error){
	var(
		jobKey string
		jobValue []byte
		putResp *clientv3.PutResponse
		oldJobObj common.Job
	)

	jobKey = common.Job_Save_Dir + job.Name
	if jobValue, err = json.Marshal(job); err!=nil{
		return
	}

	if putResp, err = j.jobKV.Put(context.TODO(),jobKey,string(jobValue),clientv3.WithPrevKV()); err!=nil{
		return
	}

	if putResp.PrevKv != nil{
		if err = json.Unmarshal(putResp.PrevKv.Value,&oldJobObj); err!=nil{
			err=nil
			return
		}
		oldJob = &oldJobObj
	}

	return
}



func (j *JobMgr) DeleteJob(name string) (oldJob *common.Job, err error){
	var(
		jobKey string
		delResp *clientv3.DeleteResponse
		oldJobObj common.Job
	)

	jobKey = common.Job_Save_Dir + name

	if delResp, err = j.jobKV.Delete(context.TODO(),jobKey,clientv3.WithPrevKV()); err!=nil{
		return
	}

	if len(delResp.PrevKvs) != 0 {
		if err = json.Unmarshal(delResp.PrevKvs[0].Value,&oldJobObj); err!=nil{
			err=nil
			return
		}
	}

	oldJob = &oldJobObj

	return
}



func (j *JobMgr) ListJobs() (jobList []*common.Job, err error){
	var(
		jobPath string
		getResp *clientv3.GetResponse
		keyPair *mvccpb.KeyValue
		job *common.Job
	)

	jobPath = common.Job_Save_Dir

	if getResp, err = j.jobKV.Get(context.TODO(),jobPath,clientv3.WithPrefix()); err!=nil{
		return
	}

	jobList = make([]*common.Job,0)


	for _, keyPair = range getResp.Kvs{
		job = &common.Job{}
		if err = json.Unmarshal(keyPair.Value,job); err!=nil{
			err = nil
			continue
		}
		jobList = append(jobList,job)
	}

	return
}


func (j *JobMgr) KillJob(name string) (err error){
	var(
		killKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
	)

	killKey = common.Job_Kill_dir + name

	if leaseGrantResp, err = j.jobLease.Grant(context.TODO(),1); err!=nil{
		return
	}

	leaseId = leaseGrantResp.ID

	if _, err = j.jobKV.Put(context.TODO(),killKey,"", clientv3.WithLease(leaseId)); err!=nil{
		return
	}

	return
}