package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte,error){
	resp, err := http.Get(url)
	if err!=nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		return nil, fmt.Errorf("fail to link to %s",url)
	}

	encoding := determinEncoding(resp.Body)
	reader := transform.NewReader(resp.Body, encoding.NewDecoder())

	return ioutil.ReadAll(reader)

}


func determinEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err!=nil{
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}