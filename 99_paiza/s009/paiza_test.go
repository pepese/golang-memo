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
	var inputs, outputs []string
	var mode int
	for i := 0; sc.Scan(); i++ {
		in := sc.Text()
		if in == `[input]` {
			mode = 0
			if i > 0 {
				tests = append(tests, Test{in: strings.NewReader(strings.Join(inputs, "\n")), want: strings.Join(outputs, "\n")})
			}
			inputs = make([]string, 0)
		} else if in == `[output]` {
			mode = 1
			outputs = make([]string, 0)
		} else if mode == 0 {
			inputs = append(inputs, in)
		} else if mode == 1 {
			outputs = append(outputs, in)
		}
	}
	tests = append(tests, Test{in: strings.NewReader(strings.Join(inputs, "\n")), want: strings.Join(outputs, "\n")})
	return tests
}

func TestExec(t *testing.T) {
	tests := readFile(`testdata/testdata.txt`)
	for _, tt := range tests {
		if got := exec(tt.in); got != tt.want {
			t.Fatalf("want = %s, got = %s", tt.want, got)
		}
	}
}
