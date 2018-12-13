package listfileservice

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

func ListFile(writer http.ResponseWriter, request *http.Request) error{
	if strings.Index(request.URL.Path,prefix)!=0{
		return userError("prefix must be "+prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err!= nil{
		return err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err!=nil{
		return err
	}
	writer.Write(contents)
	return nil
}