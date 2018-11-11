package main

import (
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/ClosurePrac/fibonacci/web/filelistservice"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(selfhandler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Panic: ",r)
				http.Error(writer, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			}
		}()
		err := selfhandler(writer, request)
		code := http.StatusOK
		if err != nil {

			if userRrr, ok := err.(userError); ok{
				http.Error(writer,userRrr.Message(),http.StatusBadRequest)
				return
			}

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelistservice.ListFileService))
	http.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})

	fmt.Println("Server running...")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}

}
