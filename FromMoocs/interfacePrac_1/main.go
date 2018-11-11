package main

import (
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/interfacePrac_1/mock"
	real2 "github.com/weizhe0422/GolangPractice/FromUdemy/interfacePrac_1/real"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string{
	return r.Get("http://www.google.com")
}
func main() {
	var r Retriver
	r = mock.Retriver{"this is fake retirver"}
	r = real2.Retriver{"http://www.yahoo.com.tw",100}
	fmt.Println(download(r))

}
