package main

import (
	"io/ioutil"
	"os"

	"./fabslang"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: fabslang code_file.fabs")
		return
	}

	f := os.Args[1]
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		println("File can't read")
		return
	}

	b := fabslang.New(string(bytes), os.Stdin, os.Stdout)
	b.Run()
}
