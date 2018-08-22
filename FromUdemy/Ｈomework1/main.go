package main

import (
	"fmt"
)

func main() {
	numList := make(map[int]string)

	for index := 0; index <= 11; index++ {
		if index%2 == 0 {
			numList[index] = "Odd"
		} else {
			numList[index] = "EVEN"
		}
	}

	for i, j := range numList {
		fmt.Println(i, ": ", j)
	}
}
