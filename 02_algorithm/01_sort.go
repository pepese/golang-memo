package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(list []int) (int, error) {
	roopNum := 0
	if list == nil {
		return roopNum, errors.New("")
	} else if len(list) < 2 {
		return roopNum, nil
	}
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				tmp := list[i]
				list[i] = list[i+1]
				list[i+1] = tmp
				flag = true
			}
			roopNum++
		}
	}
	return roopNum, nil
}

func main() {
	list := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		list = append(list, rand.Intn(10))
	}
	fmt.Println(list)
	num, err := bubbleSort(list)
	fmt.Println(list, num, err)
}
