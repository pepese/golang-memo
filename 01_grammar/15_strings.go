package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var in int
	fmt.Fscanf(os.Stdin, "%d\n", &in)
	var s []string
	for i := 0; i < in; i++ {
		s = append(s, "*")
	}
	fmt.Println(strings.Join(s, ""))
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	in := sc.Text()
	in = strings.Replace(in, "A", "4", -1)
	in = strings.Replace(in, "E", "3", -1)
	in = strings.Replace(in, "G", "6", -1)
	in = strings.Replace(in, "I", "1", -1)
	in = strings.Replace(in, "O", "0", -1)
	in = strings.Replace(in, "S", "5", -1)
	in = strings.Replace(in, "Z", "2", -1)
	fmt.Println(in)
}
*/
