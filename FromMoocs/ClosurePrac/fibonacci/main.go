package main

import (
	"bufio"
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/ClosurePrac/fibonacci/fib"
	"log"
	"os"
)

func writeFile(fileName string) bool{
	//file, err := os.Create(fileName)
	file, err:= os.OpenFile(fileName,os.O_EXCL|os.O_CREATE,0666)
	if err!= nil{
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		}else{
			fmt.Printf("%s --- %s --- %s",  pathError.Err, pathError.Op, pathError.Path)
		}
		log.Fatalln(err)
		return false
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i:= 0; i < 20; i++{
		fmt.Fprintln(writer,f())
	}
	return true
}

func main() {
	writeFile("fib.txt")
}
