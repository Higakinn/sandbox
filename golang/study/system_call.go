// go build ignore
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main(){
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	fmt.Println("バイナリ表示")
	fmt.Println(binary)

	args := []string{"ls", "-a", "-l"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
