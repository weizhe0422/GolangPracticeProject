package main

import (
	"fmt"
	"log"
)

/*
 * Complete the 'getTimes' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY time
 *  2. INTEGER_ARRAY direction
 */

func getTimes(time []int32, direction []int32) []int32 {
// Write your code here
	timeMap := make(map[int32][]int)
	result := make([]int32,len(time))
	timer := 0
	preDir := 1

	for idx, value := range time{
		timeMap[value] = append(timeMap[value], idx)
	}


	log.Println("timeMap",timeMap)

	for len(timeMap) > 0{
		for time , person := range timeMap{
			log.Println("time 00",time)
			if time<=int32(timer){
				log.Println("time 0",time)
				log.Println("person 0",person)
				if len(person) > 1{
					for _, perNo := range person{
						if direction[perNo] == int32(preDir){
							result[perNo] = int32(timer)
							log.Println("perNo 0",perNo)
							//timer++
							timeMap[time] = append(timeMap[time][:perNo],timeMap[time][perNo+1:]...)
							log.Println(result)
							goto END
						}
					}
				}else{
					log.Println("time 1",time)
					if direction[person[0]] == int32(preDir){
						result[person[0]] = int32(timer)
						delete(timeMap,int32(time))
						log.Println("person 1:",person)
						log.Println(result)
						goto END
					}else{
						result[person[0]] = int32(timer)
						preDir = int(direction[person[0]])
						delete(timeMap,int32(time))
						log.Println("person 2:",person)
						log.Println(result)
						goto END
					}
				}
			END:
				log.Println("timer F1",timer)
				log.Println("timeMap",timeMap)
			}
			timer++
			log.Println("timer F2",timer)
			log.Println("timeMap",timeMap)
		}
	}



	return result

}

func main() {
	time := []int32{0,0,1,5}
	direction := []int32{0,1,1,0}
	fmt.Println(getTimes(time,direction))
}
