package main

// 算術演算について
// http://golang.jp/go_spec#Arithmetic_operators
/*
+    和                     整数、浮動小数点、複素数、文字列
-    差                     整数、浮動小数点、複素数
*    積                     整数、浮動小数点、複素数
/    商                     整数、浮動小数点、複素数
%    剰余                   整数

&    ビット演算 and          整数
|    ビット演算 or           整数
^    ビット演算 xor          整数
&^   ビットクリア(and not)   整数

<<   左シフト                整数 << 符号なし整数
>>   右シフト                整数 >> 符号なし整数
*/
// 上記以外は "math" パッケージを使う
// http://golang.jp/pkg/math

import (
	"fmt"
	"math"
	"os"
)

// 正四面体の表面積を求める：入力は一辺
func main() {
	var in int
	fmt.Fscanf(os.Stdin, "%d", &in)
	fmt.Println(6 * math.Pow(float64(in), 2.0))
}
