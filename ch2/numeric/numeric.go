// ! 看似简单，实则深奥的数字类型
package main

func main() {
	// ? Integer types
	var int8Num int8 = 127
	var int16Num int16 = 32767
	var int32Num int32 = 2147483647
	var int64Num int64 = 9223372036854775807

	// ? Unsigned integer types
	var uint8Num uint8 = 255
	var uint16Num uint16 = 65535
	var uint32Num uint32 = 4294967295
	var uint64Num uint64 = 18446744073709551615

	println(int8Num, int16Num, int32Num, int64Num)
	println(uint8Num, uint16Num, uint32Num, uint64Num)
	println(123)
}
