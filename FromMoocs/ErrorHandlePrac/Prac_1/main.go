package listfileservice

import (
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/ErrorHandlePrac/Prac_1/listfileservice"
	"log"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}
type errHandler func(http.ResponseWriter,*http.Request) error
func errWrapper(handler errHandler) func(http.ResponseWriter,*http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			r := recover()
			if r!=nil{
				log.Print(r)
				http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)

		if userErr, ok := err.(userError); ok {
			http.Error(w,userErr.Message(),http.StatusBadRequest)
		}

		code := http.StatusOK

		switch {
		case os.IsNotExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}

		http.Error(w,http.StatusText(code),code)

	}
}

func main(){

	http.HandleFunc("/list/",errWrapper(listfileservice.ListFile))

	http.ListenAndServe(":7890",nil)
}