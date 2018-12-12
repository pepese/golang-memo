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
	for i := len(list) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
			}
			roopNum++
		}
	}
	return roopNum, nil
}

func quickSort(list []int, left int, right int) (int, error) {
	roopNum := 0
	if list == nil {
		return roopNum, errors.New("")
	} else if len(list) < 2 {
		return roopNum, nil
	}
	pivot := func() int {
		num1, num2, num3 := list[left], list[int((left+right)/2)], list[right]
		if num1 > num2 {
			if num1 < num3 {
				return num1
			} else if num2 > num3 {
				return num2
			} else {
				return num3
			}
		} else {
			if num2 < num3 {
				return num2
			} else if num1 > num3 {
				return num1
			} else {
				return num3
			}
		}
	}()
	i, j := left, right
	if right > left {
		for true {
			for list[i] <= pivot {
				i++
			}
			for list[j] > pivot {
				j--
			}
			roopNum++
			if i >= j {
				break
			}
			tmp := list[i]
			list[i] = list[j]
			list[j] = tmp
		}
		roop1, _ := quickSort(list, left, i-1)
		roop2, _ := quickSort(list, j+1, right)
		roopNum += roop1 + roop2
	}
	return roopNum, nil
}

func main() {
	list := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		list = append(list, rand.Intn(10))
	}
	fmt.Println("before sort ", list)
	/*
		num, err := bubbleSort(list)
		if err == nil {
			fmt.Println("after bubbleSort  ", list, num, err)
		} else {
			fmt.Println(err)
		}*/
	num, _ := quickSort(list, 0, len(list)-1)
	fmt.Println("after sort  ", list, num)
}
