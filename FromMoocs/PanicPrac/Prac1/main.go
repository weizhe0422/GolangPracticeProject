package main

import "fmt"

func tryPanic() {
	defer func() {
		r:=recover()
		if err,ok := r.(error); ok {
			fmt.Println("Error occured:",err)
		}else{
			panic(fmt.Sprint("Don't know error: ",r))
		}
	}()
	//b:=0
	//a:=5/b
	//fmt.Println(a)
	panic(123)
}

func main() {
	tryPanic()
}
