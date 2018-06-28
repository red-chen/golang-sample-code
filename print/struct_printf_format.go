package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	/*
		一般占位符
		符号	说明
		%v	相应值的默认格式
		%+v	在打印结构体时，默认格式，会添加字段名
		%#v	相应值的 Go 语法表示
		%T	相应值的类型的 Go 语法表示
		%%	字面上的百分号，并非值的占位符
	*/

	p := point{1, 2}

	// {1 2}
	fmt.Printf("%v\n", p)

	// {x:1 y:2}
	fmt.Printf("%+v\n", p)

	// main.point{x:1, y:2}
	fmt.Printf("%#v\n", p)

	// main.point
	fmt.Printf("%T\n", p)

	fmt.Printf("%p\n", &p)
}
