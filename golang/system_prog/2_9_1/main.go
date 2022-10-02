/*
* 2.4.7「フォーマットしてデータを io.Writer に書き出す」で説明したように、
* fmt. Fprintf(writer, フォーマット文字列, 値...)でio.Writerに数値や文字列を 出力できます。
* fmt.Fprintf() では、他の言語と同じように、"%d" が数値の表示 に、"%s" が文字列の表示に、"%f" が浮動小数点数の表示に使えます。
* これらを使って、数字や小数、文字列をファイルに書き出してみましょう。
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "数値: %d, 文字列: %s, 浮動小数点: %f", 1, "string", 1.4)
}
