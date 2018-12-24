package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	list := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rnd := rand.Intn(10)
		list = append(list, rnd)
	}
	fmt.Println("before bubbleSort ", list)
	num, err := bubbleSort(list)
	if err == nil {
		fmt.Println("after bubbleSort ", list, num, err)
	} else {
		fmt.Println(err)
	}
}

func TestQuickSort(t *testing.T) {
	list := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rnd := rand.Intn(10)
		list = append(list, rnd)
	}
	fmt.Println("before quickSort ", list)
	num, err := quickSort(list)
	if err == nil {
		fmt.Println("after quickSort  ", list, num, err)
	} else {
		fmt.Println(err)
	}
}

func TestMergeSort(t *testing.T) {
	list := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rnd := rand.Intn(10)
		list = append(list, rnd)
	}
	fmt.Println("before mergeSort ", list)
	num, err := mergeSort(list)
	if err == nil {
		fmt.Println("after mergeSort  ", list, num, err)
	} else {
		fmt.Println(err)
	}
}
