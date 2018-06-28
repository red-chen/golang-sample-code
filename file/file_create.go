package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fp, err := os.Create("f1.data")
	check(err)
	defer fp.Close()

	n, err := fp.WriteString("hello world")
	check(err)

	fmt.Printf("Wrote %d bytes\n", n)

	// 获取当前的指针位置, SeekStart, SeekCurrent, SeekEnd
	poz, err := fp.Seek(0, io.SeekCurrent)
	fmt.Printf("current position %d\n", poz)
}
