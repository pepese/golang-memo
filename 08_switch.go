package main

import "fmt"

func main() {
	// 普通のswitch
	signal := "blue"
	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("caution")
	case "green", "blue":
		fmt.Println("Go")
	default:
		fmt.Println("wrong")
	}

	// if-else的にも書く事が可能
	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!")
	default:
		fmt.Println("Bad")
	}
}
