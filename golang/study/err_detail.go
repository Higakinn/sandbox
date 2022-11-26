//go build ignore
package main

import (
	"os"
	"fmt"
	//"log"
)

func main() {
	_, err := os.Open("not_found.txt")
	if err != nil {
		// err.Error, err どっちの記述でもエラーの出力内容は変わらない
		fmt.Println("==========err.Error()===========")
		fmt.Println(err.Error())
		fmt.Println("==========err =============")
		fmt.Println(err)
	}
}
