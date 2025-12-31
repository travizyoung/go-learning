package main

import "fmt"

func main() {
	// ? 使用下划线_分隔数字
	salary := 15_000
	fmt.Println(salary)

	// ? 使用e来表示指数
	floatNum := 6.03e23
	fmt.Println(floatNum)

	// ? 使用p来表示十六进制浮点数
	floatNumHex := 0x12.34p5
	fmt.Println(floatNumHex)

	// ? rune 卢恩类型，使用单引号表示
	fmt.Println('a', '\n', '\t', '\'', '\\')

	// ? 解释字符串
	fmt.Println("Greetings and\n\"Salutations\"")
	// ! 反斜杠不能出现 fmt.Println("\")

	// ? 原始字符串，使用反引号表示
	fmt.Println(`Greetings and
"Salutations"`)

	// ? bitwise 操作符
	a := 60             // 60 = 0011 1100
	b := 13             // 13 = 0000 1101
	fmt.Println(a & b)  // AND  = 0000 1100 = 12
	fmt.Println(a | b)  // OR   = 0011 1101 = 61
	fmt.Println(a ^ b)  // XOR  = 0011 0001 = 49
	fmt.Println(a &^ b) // AND NOT = 0011 0000 = 48
}
