package main

import "fmt"

func main() {

	/*
		布尔占位符
		符号	说明
		%t	单词 true 或 false
	*/
	fmt.Printf("%t\n", true) // true

	/*
		整数占位符
		符号	说明
		%b	二进制表示
		%c	相应 Unicode 码点所表示的字符
		%d	十进制表示
		%o	八进制表示
		%f	十进制浮点数
		%q	单引号围绕的字符字面值，由 Go 语法安全地转义
		%x	十六进制表示，字母形式为小写 a-f
		%X	十六进制表示，字母形式为大写 A-F
		%U	Unicode 格式：U+1234，等同于 "U+%04X"
	*/
	fmt.Printf("%b\n", 14)          // 1110
	fmt.Printf("%c\n", 33)          // !
	fmt.Printf("%q\n", 33)          // '!'
	fmt.Printf("%d\n", 123)         // 123
	fmt.Printf("%d\n", int64(123))  // 123
	fmt.Printf("%o\n", 123)         // 173
	fmt.Printf("%x\n", 123)         // 7b
	fmt.Printf("%X\n", 123)         // 7B
	fmt.Printf("%f\n", 78.9)        // 78.900000
	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) // 1.234000E+08

	// 固定格式化输出
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// out
	/*
		|    12|   345|
		|  1.20|  3.45|
		|1.20  |3.45  |
		|   foo|     b|
		|foo   |b     |
	*/

}
