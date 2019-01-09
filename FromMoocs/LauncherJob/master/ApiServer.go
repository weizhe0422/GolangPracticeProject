package master

import (
	"encoding/json"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/LauncherJob/master/common"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	G_ApiServer *ApiServer
)

func handleSaveJob(resp http.ResponseWriter, req *http.Request) {
	var (
		err           error
		postJob       string
		job           common.Job
		oldJob        *common.Job
		BuildResponse []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	postJob = req.PostForm.Get("job")

	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	if oldJob, err = G_JobMgr.SaveJob(&job); err != nil {
		goto ERR
	}

	if BuildResponse, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(BuildResponse)
	}

	//記得加return，否則會持續進到ERR區間
	return

ERR:
	log.Println("3", string(BuildResponse))
	if BuildResponse, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(BuildResponse)
	}
}

func handleDelJob(resp http.ResponseWriter, req *http.Request) {
	var (
		err       error
		jobName   string
		oldJob    *common.Job
		respBytes []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	jobName = req.PostForm.Get("name")

	if oldJob, err = G_JobMgr.DeleteJob(jobName); err != nil {
		goto ERR
	}

	log.Println(oldJob)

	if respBytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		log.Println("respBytes", string(respBytes))
		resp.Write(respBytes)
	}

	return

ERR:
	if respBytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(respBytes)
	}
}

func handleListJob(resp http.ResponseWriter, req *http.Request) {
	var (
		err       error
		jobList   []*common.Job
		respBytes []byte
	)

	if jobList, err = G_JobMgr.ListJobs(); err != nil {
		goto ERR

	}

	if respBytes, err = common.BuildResponse(0, "success", jobList); err == nil {
		resp.Write(respBytes)
	}

	return

ERR:
	if respBytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(respBytes)
	}
}

func handleKillJob(resp http.ResponseWriter, req *http.Request) {
	var (
		err      error
		jobName  string
		respByes []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	jobName = req.PostForm.Get("name")

	if err = G_JobMgr.KillJob(jobName); err != nil {
		goto ERR
	}

	if respByes, err = common.BuildResponse(0, "success", nil); err == nil {
		resp.Write(respByes)
	}

	return

ERR:
	if respByes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(respByes)
	}
}

func InitApiServer() (err error) {
	var (
		mux       *http.ServeMux
		httpSvr   *http.Server
		listener  net.Listener
		staticDir http.Dir
		staticHandler http.Handler
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/jobs/save", handleSaveJob)
	mux.HandleFunc("/jobs/delete", handleDelJob)
	mux.HandleFunc("/jobs/list", handleListJob)
	mux.HandleFunc("/jobs/kill", handleKillJob)

	staticDir = http.Dir(G_Config.WebRoot)
	staticHandler = http.FileServer(staticDir)
	mux.Handle("/",http.StripPrefix("/",staticHandler))

	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_Config.ApiSvrPrt)); err != nil {
		return
	}

	httpSvr = &http.Server{
		ReadTimeout:  time.Duration(G_Config.ApiSvrReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(G_Config.ApiSvrWriteTimeOut) * time.Millisecond,
		Handler:      mux,
	}

	G_ApiServer = &ApiServer{
		httpServer: httpSvr,
	}

	go httpSvr.Serve(listener)

	return

}
