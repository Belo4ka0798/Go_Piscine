package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var args []string
	args = append(args, os.Args[2:]...)
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		args = append(args, scan.Text())
	}
	cmd := exec.Command(os.Args[1], args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Command is not valid!")
		os.Exit(1)
	}
	fmt.Println(out.String())
}
