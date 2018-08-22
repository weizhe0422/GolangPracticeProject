package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type inputData struct {
	accountNum int
	saveMoney  int
}

var in []chan inputData
var out []chan int

const maxUser = 100
const maxThread = 10

var account [maxUser]string
var sum [maxUser]int

// Random get random value
func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func saveMoney(accountNum, money int) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		channelNum := accountNum % maxThread

		in[channelNum] <- inputData{accountNum, money}
		for {
			select {
			//case result := <-out[channelNum]:
			case <-out[channelNum]:
				//fmt.Printf("User%d's money is %d\n", accountNum, result)
				//wg.Done()
				return
			}
		}
	}(&wg)

	wg.Wait()
}

func main() {
	in = make([]chan inputData, maxThread)
	out = make([]chan int, maxThread)

	for i := range in {
		in[i] = make(chan inputData)
		out[i] = make(chan int)
	}

	for i := 0; i < maxUser; i++ {
		account[i] = "User" + strconv.Itoa(i)
		sum[i] = 50
	}

	for i := range in {
		go func(in *chan inputData, i int) {
			for {
				select {
				case inputMoney := <-*in:
					sum[inputMoney.accountNum-1] = sum[inputMoney.accountNum-1] + inputMoney.saveMoney

					out[inputMoney.accountNum%maxThread] <- sum[inputMoney.accountNum-1]
				}
			}
		}(&in[i], i)
	}

	idx := 0
	saveCNT := 10000
	for idx < saveCNT {
		number := Random(1, maxUser)
		saveMoney(number, 100)
		idx++
	}

	for j := 0; j < maxUser; j++ {
		fmt.Printf("User_%d: %d \n", j, sum[j])
	}

}
