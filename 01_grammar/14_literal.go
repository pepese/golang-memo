package main

import (
	"fmt"
	"os"
)

func main() {
	var in string
	fmt.Fscanf(os.Stdin, "%s\n", &in) // 標準入力待ち

	//words := []byte(in)
	//words[len(words)-1] = byte('0')
	//words[len(words)-2] = byte('0')

	//words := []int32(in)
	//words[len(words)-1] = 48
	//words[len(words)-2] = 48

	words := []rune(in)
	words[len(words)-1] = 48
	words[len(words)-2] = 48

	fmt.Println(string(words))
}
