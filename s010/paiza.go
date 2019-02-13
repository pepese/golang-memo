package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	T, _ := strconv.Atoi(ins[0])
	B, _ := strconv.Atoi(ins[1])
	U, _ := strconv.Atoi(ins[2])
	D, _ := strconv.Atoi(ins[3])
	L, _ := strconv.Atoi(ins[4])
	R, _ := strconv.Atoi(ins[5])
	sc.Scan()
	// どこにある判定
	a := func(in, t, b, u, d, l, r int) (int, int, int, int, int, int, int) {
		switch in {
		case t: // そのまま、回転しない
			//fmt.Println("t")
			return 0, t, b, u, d, l, r
		case u: // 右、左回転:t->u u->b, b->d, d->t
			//fmt.Println("u")
			return 1, u, d, b, t, l, r
		case b: // 下、右に二回転:t->b, u->d, b->t, d->u
			//fmt.Println("b")
			return 2, b, t, d, u, l, r
		case d: // 左、右回転:t->d, u->t,b->u,d->b
			//fmt.Println("d")
			return 1, d, u, t, b, l, r
		case l: // 前、奥回転:t->l, l->r, r->t
			//fmt.Println("l")
			return 1, l, b, u, d, r, t
		case r: // 奥、前回転:t->r, r->l ,l->t
			//fmt.Println("r")
			return 1, r, b, u, d, t, l
		}
		//fmt.Println("hoge")
		return 0, t, b, u, d, l, r
	}
	// 回転
	//N, _ := strconv.Atoi(sc.Text())
	sum := 1
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		var out int
		out, T, B, U, D, L, R = a(num, T, B, U, D, L, R)
		//fmt.Println("num: ", out)
		sum += out
	}
	return fmt.Sprintf("%d", sum)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
