package main

import "fmt"

// リストから指定のインデッックスの要素を削除
func delEle(list []int, i int) []int {
	result := []int{}
	result = append(result, list[:i]...)
	result = append(result, list[i+1:]...)
	return result
}

func permutation(list []int) [][]int {
	var result [][]int            // 結果の格納
	var inner func(in, out []int) // 内部の再帰関数
	// in を捜査して、 out へ足していく
	inner = func(in, out []int) {
		if len(in) == 1 { // len(in) が 1 の時は結果を格納して終了
			result = append(result, append(out, in[0]))
		} else { // in から 1 つ取り出して out に付け足して再帰
			for i, n := range in {
				inner(delEle(in, i), append(out, n)) // 勘違いしやすいが参照渡しではないので引数は関数内部へコピー（clone_）される
			}
		}
	}
	inner(list, []int{}) // 初回呼び出しの out は空
	return result
}

func main() {
	fmt.Println(permutation([]int{1, 2, 3}))
}
