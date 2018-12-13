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

func quickSort(list []int) (int, error) {
	roop := 0
	var roopNum *int
	roopNum = &roop
	if list == nil {
		return *roopNum, errors.New("")
	} else if len(list) < 2 {
		return *roopNum, nil
	}
	var calcQuick func(list []int, left int, right int, roop *int)
	calcQuick = func(list []int, left int, right int, roop *int) {
		if left < right {
			i, j := left, right
			pivot := func(num1 int, num2 int, num3 int) int {
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
			}(list[left], list[int(i+(j-i)/2)], list[right])
			for true {
				for list[i] < pivot {
					i++
				}
				for pivot < list[j] {
					j--
				}
				*roop++
				if i >= j {
					break
				}
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
				i++
				j--
			}
			calcQuick(list, left, i-1, roop)
			calcQuick(list, j+1, right, roop)
		}
	}
	calcQuick(list, 0, len(list)-1, roopNum)
	return *roopNum, nil
}

func mergeSort(list []int) (int, error) {
	roopNum := 0
	if list == nil {
		return roopNum, errors.New("")
	}
	var sort func(list []int, left int, right int)
	sort = func(list []int, left int, right int) {
		i, k := left, right
		j := int((i + k) / 2)
		if i == k {
			return
		}
		sort(list, i, j)
		sort(list, j+1, k)
		func(_list []int, _i int, _j int, _k int) {
			l, m := _i, _j
			var tmp []int
			for l <= _j && m <= _k {
				if _list[l] < _list[m] {
					tmp = append(tmp, _list[l])
					l++
				} else {
					tmp = append(tmp, _list[m])
					m++
				}
			}
			copy(_list[_i:_k+1], tmp[:])
		}(list, i, j, k)
	}
	sort(list, 0, len(list)-1)
	return roopNum, nil
}

func main() {
	list1 := []int{}
	list2 := []int{}
	list3 := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rnd := rand.Intn(10)
		list1 = append(list1, rnd)
		list2 = append(list2, rnd)
		list3 = append(list2, rnd)
	}
	fmt.Println("before sort ", list1)
	num, err := bubbleSort(list1)
	if err == nil {
		fmt.Println("after bubbleSort  ", list1, num, err)
	} else {
		fmt.Println(err)
	}
	num, err = quickSort(list2)
	if err == nil {
		fmt.Println("after quickSort  ", list2, num, err)
	} else {
		fmt.Println(err)
	}
	num, err = mergeSort(list3)
	if err == nil {
		fmt.Println("after mergeSort  ", list3, num, err)
	} else {
		fmt.Println(err)
	}
}
