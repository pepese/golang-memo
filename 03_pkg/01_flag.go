package main

import (
	//"flag"
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	intOpt := flag.Int("i", 1234, "help message for i option")
	boolOpt := flag.Bool("b", false, "help message for b option")
	strOpt := flag.String("s", "default", "help message for s option")

	flag.Parse()
	fmt.Println(*flag.CommandLine)
	fmt.Println(flag.Arg(1))
	fmt.Println(flag.Args())
	fmt.Println(*intOpt)
	fmt.Println(*boolOpt)
	fmt.Println(*strOpt)
}
