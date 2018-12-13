package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func countMax(input []int,start, end int ) (arrLow, arrHigh, arrSum int) {
	mid := (start + end) / 2

	log.Println("input", start,end,mid)

	corssLow, corssHigh, corssSum := findMaximumCorssSubArray(input, start, end, len(input))

	return corssLow, corssHigh, corssSum
}

func findMaximumCorssSubArray(A []int, low, mid, high int) (arrLow, arrHigh, arrSum int) {
	leftSum := math.MinInt64
	sum := 0
	for i := mid-1; i >= low; i-- {
		sum = sum + A[i]
		log.Println("sum",sum)
		log.Println("leftSum",leftSum)
		if sum > leftSum {
			leftSum = sum
			arrLow = i
		}
	}
	return arrLow, arrHigh, leftSum //+ rightSum
}
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var (
		inputSlice  []string
		inputString string
		inputValue  []int
		output int
		chkPos bool
	)

	fmt.Scan(&inputString)
	if strings.Contains(inputString, ",,") {
		fmt.Println(0)
		return
	}

	inputSlice = strings.Split(inputString, ",")

	chkPos = false
	inputValue = make([]int,len(inputSlice))
	for index, value := range inputSlice {
		i, _ := strconv.Atoi(value)
		inputValue[index] = i
		if i > 0 {
			chkPos = true
		}
	}

	if !chkPos {
		fmt.Println(0)
		return
	}
	_, _, output = countMax(inputValue,0, len(inputValue))
	fmt.Println(output)

}

//1,3,-2,1,2
