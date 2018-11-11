package main

import (
	"fmt"
	"github.com/weizhe0422/GolangPractice/FromUdemy/QueuePrac/Queue"
)

func main() {
	newQueue := Queue.Queue{}

	newQueue.Push(3)
	newQueue.Push(6)
	newQueue.Push(19)
	newQueue.Print()
	fmt.Println()

	fmt.Println("POP: ",newQueue.Pop())
	newQueue.Print()
}
