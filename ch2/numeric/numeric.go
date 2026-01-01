// ! 看似简单，实则深奥的数字类型
package main

func main() {
	// ? Integer types
	var int8Num int8 = 127
	var int16Num int16 = 32767
	var int32Num int32 = 2147483647
	var int64Num int64 = 9223372036854775807

	var b2 = 0b0000000
	println(b2)

	// 使用下划线提高数字字面量的可读性
	var decimalWithUnderscores int = 1_234_567_890        // 十进制，按千位分组
	var binaryWithUnderscores int = 0b1010_1100_0011_1101 // 二进制，每4位分组
	var octalWithUnderscores int = 0o755_421              // 八进制，按3位分组（对应八进制数字）
	var hexWithUnderscores int = 0xDEAD_BEEF_CAFE         // 十六进制，每4位分组（对应2字节）

	// 下划线也可以用于其他类型的数字
	var floatWithUnderscores float64 = 3.141_592_653_589_793
	var longIntWithUnderscores int64 = 9_223_372_036_854_775_807

	// 打印这些变量以验证
	println(decimalWithUnderscores)
	println(binaryWithUnderscores)
	println(octalWithUnderscores)
	println(hexWithUnderscores)
	println(floatWithUnderscores)
	println(longIntWithUnderscores)

	// ? Unsigned integer types
	var uint8Num uint8 = 255
	var uint16Num uint16 = 65535
	var uint32Num uint32 = 4294967295
	var uint64Num uint64 = 18446744073709551615

	println(int8Num, int16Num, int32Num, int64Num)
	println(uint8Num, uint16Num, uint32Num, uint64Num)
	println(123)
}
