package Queue

import (
	"fmt"
)

type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}

func (q *Queue) Print(){
	for _, value := range *q{
		fmt.Printf("%d ",value)
	}
}