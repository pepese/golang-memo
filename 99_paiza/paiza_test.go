package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

type Test struct {
	in   io.Reader
	want string
}

func readFile(f string) []Test {
	file, _ := os.Open(f)
	defer file.Close()
	sc := bufio.NewScanner(file)
	tests := make([]Test, 0)
	var ins []string
	var mode int
	for sc.Scan() {
		in := sc.Text()
		if in == `[input]` {
			mode = 0
		} else if in == `[output]` {
			mode = 1
		} else if mode == 0 {
			ins = append(ins, in)
		} else if mode == 1 {
			tests = append(tests, Test{in: strings.NewReader(strings.Join(ins, "\n")), want: in})
			ins = make([]string, 0)
		}
	}
	return tests
}

func TestExec(t *testing.T) {
	tests := readFile(`testdata/exec.txt`)
	for _, tt := range tests {
		if got := exec(tt.in); got != tt.want {
			t.Fatalf("want = %s, got = %s", tt.want, got)
		}
	}
}
