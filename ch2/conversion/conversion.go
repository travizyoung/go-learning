package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== 十进制和二进制相互转换技巧 ===")
	fmt.Println()

	// 1. 十进制转二进制的基本方法
	fmt.Println("1. 十进制转二进制的基本方法:")
	decimalToBinaryBasic(42)
	decimalToBinaryBasic(255)
	decimalToBinaryBasic(1024)
	fmt.Println()

	// 2. 二进制转十进制的基本方法
	fmt.Println("2. 二进制转十进制的基本方法:")
	binaryToDecimalBasic("101010")
	binaryToDecimalBasic("11111111")
	binaryToDecimalBasic("10000000000")
	fmt.Println()

	// 3. 使用Go语言内置函数进行转换
	fmt.Println("3. 使用Go语言内置函数进行转换:")
	useBuiltInFunctions()
	fmt.Println()

	// 4. 位运算技巧
	fmt.Println("4. 位运算技巧:")
	bitwiseTechniques()
	fmt.Println()

	// 5. 快速转换技巧
	fmt.Println("5. 快速转换技巧:")
	quickConversionTips()
	fmt.Println()

	// 6. 负数的二进制表示（补码）
	fmt.Println("6. 负数的二进制表示（补码）:")
	twosComplement()
	fmt.Println()

	// 7. 实际应用示例
	fmt.Println("7. 实际应用示例:")
	practicalExamples()
}

// 十进制转二进制的基本方法 - 除2取余法
func decimalToBinaryBasic(n int) {
	original := n
	var binary []int

	if n == 0 {
		fmt.Printf("  %d (十进制) = 0 (二进制)\n", n)
		return
	}

	for n > 0 {
		remainder := n % 2
		binary = append([]int{remainder}, binary...)
		n = n / 2
	}

	fmt.Printf("  %d (十进制) = ", original)
	for _, bit := range binary {
		fmt.Printf("%d", bit)
	}
	fmt.Printf(" (二进制)\n")
}

// 二进制转十进制的基本方法 - 按权展开法
func binaryToDecimalBasic(binaryStr string) {
	decimal := 0
	power := 1 // 2^0 = 1

	// 从右向左处理二进制字符串
	for i := len(binaryStr) - 1; i >= 0; i-- {
		if binaryStr[i] == '1' {
			decimal += power
		}
		power *= 2
	}

	fmt.Printf("  %s (二进制) = %d (十进制)\n", binaryStr, decimal)
}

// 使用Go语言内置函数进行转换
func useBuiltInFunctions() {
	// 十进制转二进制
	decimal := 42
	binaryStr := strconv.FormatInt(int64(decimal), 2)
	fmt.Printf("  strconv.FormatInt(%d, 2) = %s\n", decimal, binaryStr)

	// 二进制转十进制
	binary := "101010"
	if decimal, err := strconv.ParseInt(binary, 2, 64); err == nil {
		fmt.Printf("  strconv.ParseInt(%s, 2, 64) = %d\n", binary, decimal)
	}

	// 使用fmt包格式化输出
	fmt.Printf("  fmt.Sprintf(\"%%b\", %d) = %s\n", decimal, fmt.Sprintf("%b", decimal))
	fmt.Printf("  fmt.Sprintf(\"%%#b\", %d) = %s\n", decimal, fmt.Sprintf("%#b", decimal))
}

// 位运算技巧
func bitwiseTechniques() {
	fmt.Println("  a) 检查奇偶性:")
	num := 42
	if num&1 == 0 {
		fmt.Printf("    %d & 1 = 0, 所以 %d 是偶数\n", num, num)
	} else {
		fmt.Printf("    %d & 1 = 1, 所以 %d 是奇数\n", num, num)
	}

	fmt.Println("  b) 乘以2的幂:")
	num = 5
	fmt.Printf("    %d << 1 = %d (相当于 %d * 2)\n", num, num<<1, num)
	fmt.Printf("    %d << 2 = %d (相当于 %d * 4)\n", num, num<<2, num)
	fmt.Printf("    %d << 3 = %d (相当于 %d * 8)\n", num, num<<3, num)

	fmt.Println("  c) 除以2的幂:")
	num = 40
	fmt.Printf("    %d >> 1 = %d (相当于 %d / 2)\n", num, num>>1, num)
	fmt.Printf("    %d >> 2 = %d (相当于 %d / 4)\n", num, num>>2, num)
	fmt.Printf("    %d >> 3 = %d (相当于 %d / 8)\n", num, num>>3, num)

	fmt.Println("  d) 设置特定位:")
	num = 0b1010 // 十进制10
	// 设置第2位（从0开始计数）
	num = num | (1 << 2) // 0b1010 | 0b0100 = 0b1110 (14)
	fmt.Printf("    设置第2位: 0b1010 | (1 << 2) = 0b%04b (%d)\n", num, num)

	fmt.Println("  e) 清除特定位:")
	num = 0b1110 // 十进制14
	// 清除第2位
	num = num & ^(1 << 2) // 0b1110 & 0b1011 = 0b1010 (10)
	fmt.Printf("    清除第2位: 0b1110 & ^(1 << 2) = 0b%04b (%d)\n", num, num)

	fmt.Println("  f) 切换特定位:")
	num = 0b1010 // 十进制10
	// 切换第1位
	num = num ^ (1 << 1) // 0b1010 ^ 0b0010 = 0b1000 (8)
	fmt.Printf("    切换第1位: 0b1010 ^ (1 << 1) = 0b%04b (%d)\n", num, num)
}

// 快速转换技巧
func quickConversionTips() {
	fmt.Println("  a) 2的幂次方快速识别:")
	powersOfTwo := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for _, n := range powersOfTwo {
		// 2的幂次方的二进制表示只有一个1
		if n&(n-1) == 0 {
			fmt.Printf("    %d = 2^%d, 二进制: %b\n", n, log2(n), n)
		}
	}

	fmt.Println("  b) 快速转换8位二进制:")
	fmt.Println("    记住这些常用值:")
	commonValues := map[string]int{
		"00000001": 1,
		"00000010": 2,
		"00000100": 4,
		"00001000": 8,
		"00010000": 16,
		"00100000": 32,
		"01000000": 64,
		"10000000": 128,
	}
	for binary, decimal := range commonValues {
		fmt.Printf("    %s = %d\n", binary, decimal)
	}

	fmt.Println("  c) 十六进制与二进制的快速转换:")
	fmt.Println("    每个十六进制数字对应4位二进制:")
	hexToBinary := map[byte]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	}
	for hex, binary := range hexToBinary {
		fmt.Printf("    %c = %s  ", hex, binary)
		if hex == '7' || hex == 'F' {
			fmt.Println()
		}
	}
}

// 负数的二进制表示（补码）
func twosComplement() {
	fmt.Println("  a) 正数的二进制表示:")
	positive := 42
	fmt.Printf("    +%d = %08b\n", positive, positive)

	fmt.Println("  b) 负数的二进制表示（补码）:")
	negative := -42
	// 在Go中，负数的二进制表示是补码形式
	fmt.Printf("    %d = %b\n", negative, negative)
	fmt.Printf("    %d 的32位表示 = %032b\n", negative, uint32(negative))

	fmt.Println("  c) 补码计算方法:")
	fmt.Println("    1. 取正数的二进制表示")
	fmt.Println("    2. 按位取反（得到反码）")
	fmt.Println("    3. 加1（得到补码）")

	// 演示-42的补码计算
	pos := 42
	posBinary := fmt.Sprintf("%08b", pos)
	fmt.Printf("    +42 = %s\n", posBinary)

	// 按位取反
	inverted := ^uint8(pos)
	fmt.Printf("    按位取反: ^%s = %08b\n", posBinary, inverted)

	// 加1得到补码
	twosComp := int8(inverted) + 1
	fmt.Printf("    加1: %08b + 1 = %08b (%d)\n", inverted, uint8(twosComp), twosComp)
}

// 实际应用示例
func practicalExamples() {
	fmt.Println("  a) IP地址转换:")
	ip := "192.168.1.1"
	fmt.Printf("    IP地址 %s 的二进制表示:\n", ip)
	// 简化示例，实际需要分割IP地址
	fmt.Println("    192 = 11000000")
	fmt.Println("    168 = 10101000")
	fmt.Println("    1   = 00000001")
	fmt.Println("    1   = 00000001")

	fmt.Println("  b) 权限控制（Linux文件权限）:")
	permissions := 0o755 // 八进制表示
	fmt.Printf("    权限 %o 的二进制表示:\n", permissions)
	fmt.Printf("    user:  %03b (读/写/执行)\n", (permissions>>6)&7)
	fmt.Printf("    group: %03b (读/执行)\n", (permissions>>3)&7)
	fmt.Printf("    other: %03b (读/执行)\n", permissions&7)

	fmt.Println("  c) 颜色值转换（RGB到二进制）:")
	color := 0xFFA500 // 橙色
	fmt.Printf("    颜色 #%06X 的二进制表示:\n", color)
	fmt.Printf("    Red:   %08b (%d)\n", (color>>16)&0xFF, (color>>16)&0xFF)
	fmt.Printf("    Green: %08b (%d)\n", (color>>8)&0xFF, (color>>8)&0xFF)
	fmt.Printf("    Blue:  %08b (%d)\n", color&0xFF, color&0xFF)

	fmt.Println("  d) 数据包校验:")
	data := []byte{0x01, 0x02, 0x03, 0x04}
	checksum := 0
	for _, b := range data {
		checksum ^= int(b) // 使用异或进行简单校验
	}
	fmt.Printf("    数据 %v 的校验和: %08b (%d)\n", data, checksum, checksum)
}

// 辅助函数：计算以2为底的对数（整数）
func log2(n int) int {
	count := 0
	for n > 1 {
		n >>= 1
		count++
	}
	return count
}
