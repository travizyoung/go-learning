package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== 十进制和二进制转换练习 ===")
	fmt.Println("请选择练习模式:")
	fmt.Println("1. 十进制转二进制")
	fmt.Println("2. 二进制转十进制")
	fmt.Println("3. 随机混合练习")
	fmt.Println("4. 位运算练习")
	fmt.Println("5. 退出")

	for {
		fmt.Print("\n请输入选项 (1-5): ")
		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			practiceDecimalToBinary(scanner)
		case "2":
			practiceBinaryToDecimal(scanner)
		case "3":
			practiceMixed(scanner)
		case "4":
			practiceBitwise(scanner)
		case "5":
			fmt.Println("再见！")
			return
		default:
			fmt.Println("无效选项，请重新输入")
		}
	}
}

func practiceDecimalToBinary(scanner *bufio.Scanner) {
	fmt.Println("\n=== 十进制转二进制练习 ===")
	fmt.Println("输入 'q' 退出练习模式")

	correct := 0
	total := 0

	for {
		// 生成随机十进制数 (0-255)
		decimal := rand.Intn(256)
		fmt.Printf("\n问题 %d: 将 %d 转换为二进制: ", total+1, decimal)

		if !scanner.Scan() {
			break
		}

		answer := strings.TrimSpace(scanner.Text())
		if answer == "q" {
			break
		}

		// 计算正确答案
		correctAnswer := strconv.FormatInt(int64(decimal), 2)

		// 移除可能的 "0b" 前缀
		userAnswer := answer
		if strings.HasPrefix(answer, "0b") {
			userAnswer = answer[2:]
		}

		if userAnswer == correctAnswer {
			fmt.Println("✅ 正确！")
			correct++
		} else {
			fmt.Printf("❌ 错误。正确答案是: %s (0b%s)\n", correctAnswer, correctAnswer)
			fmt.Printf("   计算方法: %d = ", decimal)
			printDivisionSteps(decimal)
		}

		total++
	}

	printScore(correct, total)
}

func practiceBinaryToDecimal(scanner *bufio.Scanner) {
	fmt.Println("\n=== 二进制转十进制练习 ===")
	fmt.Println("输入 'q' 退出练习模式")

	correct := 0
	total := 0

	for {
		// 生成随机二进制数 (最多8位)
		decimal := rand.Intn(256)
		binary := strconv.FormatInt(int64(decimal), 2)

		// 确保至少4位
		for len(binary) < 4 {
			binary = "0" + binary
		}

		fmt.Printf("\n问题 %d: 将 %s 转换为十进制: ", total+1, binary)

		if !scanner.Scan() {
			break
		}

		answer := strings.TrimSpace(scanner.Text())
		if answer == "q" {
			break
		}

		userAnswer, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Println("❌ 请输入有效的数字")
			continue
		}

		if userAnswer == decimal {
			fmt.Println("✅ 正确！")
			correct++
		} else {
			fmt.Printf("❌ 错误。正确答案是: %d\n", decimal)
			fmt.Printf("   计算方法: %s = ", binary)
			printExpansionSteps(binary)
		}

		total++
	}

	printScore(correct, total)
}

func practiceMixed(scanner *bufio.Scanner) {
	fmt.Println("\n=== 混合练习 ===")
	fmt.Println("输入 'q' 退出练习模式")

	correct := 0
	total := 0

	for {
		// 随机选择练习类型
		exerciseType := rand.Intn(2)

		if exerciseType == 0 {
			// 十进制转二进制
			decimal := rand.Intn(256)
			fmt.Printf("\n问题 %d (十进制转二进制): 将 %d 转换为二进制: ", total+1, decimal)

			if !scanner.Scan() {
				break
			}

			answer := strings.TrimSpace(scanner.Text())
			if answer == "q" {
				break
			}

			correctAnswer := strconv.FormatInt(int64(decimal), 2)
			userAnswer := answer
			if strings.HasPrefix(answer, "0b") {
				userAnswer = answer[2:]
			}

			if userAnswer == correctAnswer {
				fmt.Println("✅ 正确！")
				correct++
			} else {
				fmt.Printf("❌ 错误。正确答案是: %s\n", correctAnswer)
			}
		} else {
			// 二进制转十进制
			decimal := rand.Intn(256)
			binary := strconv.FormatInt(int64(decimal), 2)
			for len(binary) < 4 {
				binary = "0" + binary
			}

			fmt.Printf("\n问题 %d (二进制转十进制): 将 %s 转换为十进制: ", total+1, binary)

			if !scanner.Scan() {
				break
			}

			answer := strings.TrimSpace(scanner.Text())
			if answer == "q" {
				break
			}

			userAnswer, err := strconv.Atoi(answer)
			if err != nil {
				fmt.Println("❌ 请输入有效的数字")
				continue
			}

			if userAnswer == decimal {
				fmt.Println("✅ 正确！")
				correct++
			} else {
				fmt.Printf("❌ 错误。正确答案是: %d\n", decimal)
			}
		}

		total++
	}

	printScore(correct, total)
}

func practiceBitwise(scanner *bufio.Scanner) {
	fmt.Println("\n=== 位运算练习 ===")
	fmt.Println("输入 'q' 退出练习模式")

	correct := 0
	total := 0
	operations := []string{"&", "|", "^", "<<", ">>"}

	for {
		// 生成两个随机数
		a := rand.Intn(16)
		b := rand.Intn(5) // 对于移位操作，限制移位位数

		// 随机选择操作
		op := operations[rand.Intn(len(operations))]

		var question string
		var correctAnswer int

		switch op {
		case "&":
			question = fmt.Sprintf("%d %s %d", a, op, 1<<b)
			correctAnswer = a & (1 << b)
		case "|":
			question = fmt.Sprintf("%d %s %d", a, op, 1<<b)
			correctAnswer = a | (1 << b)
		case "^":
			question = fmt.Sprintf("%d %s %d", a, op, 1<<b)
			correctAnswer = a ^ (1 << b)
		case "<<":
			question = fmt.Sprintf("%d %s %d", a, op, b)
			correctAnswer = a << b
		case ">>":
			question = fmt.Sprintf("%d %s %d", a, op, b)
			correctAnswer = a >> b
		}

		fmt.Printf("\n问题 %d: 计算 %s = ", total+1, question)

		if !scanner.Scan() {
			break
		}

		answer := strings.TrimSpace(scanner.Text())
		if answer == "q" {
			break
		}

		userAnswer, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Println("❌ 请输入有效的数字")
			continue
		}

		if userAnswer == correctAnswer {
			fmt.Println("✅ 正确！")
			correct++
		} else {
			fmt.Printf("❌ 错误。正确答案是: %d\n", correctAnswer)
			fmt.Printf("   二进制表示:\n")
			fmt.Printf("   %d = %04b\n", a, a)
			if op == "<<" || op == ">>" {
				fmt.Printf("   %s %d 位\n", op, b)
			} else {
				fmt.Printf("   %d = %04b\n", 1<<b, 1<<b)
				fmt.Printf("   运算: %04b %s %04b = %04b\n", a, op, 1<<b, correctAnswer)
			}
		}

		total++
	}

	printScore(correct, total)
}

func printDivisionSteps(n int) {
	if n == 0 {
		fmt.Println("0")
		return
	}

	var steps []string
	original := n

	for n > 0 {
		remainder := n % 2
		steps = append([]string{fmt.Sprintf("%d ÷ 2 = %d ... %d", n, n/2, remainder)}, steps...)
		n = n / 2
	}

	fmt.Println()
	for i, step := range steps {
		fmt.Printf("   %s", step)
		if i == len(steps)-1 {
			fmt.Print(" ↑")
		}
		fmt.Println()
	}
	fmt.Printf("   从下往上读余数: ")

	// 重新计算二进制
	binary := strconv.FormatInt(int64(original), 2)
	for _, ch := range binary {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}

func printExpansionSteps(binary string) {
	fmt.Println()
	total := 0
	power := 1

	// 从右向左处理
	for i := len(binary) - 1; i >= 0; i-- {
		bit := binary[i]
		position := len(binary) - 1 - i

		if bit == '1' {
			value := power
			fmt.Printf("   %c×2¹⁽%d⁾ = %d\n", bit, position, value)
			total += value
		} else {
			fmt.Printf("   %c×2¹⁽%d⁾ = 0\n", bit, position)
		}
		power *= 2
	}

	fmt.Printf("   总和 = %d\n", total)
}

func printScore(correct, total int) {
	if total == 0 {
		fmt.Println("\n没有完成任何练习")
		return
	}

	score := float64(correct) / float64(total) * 100
	fmt.Printf("\n=== 练习结果 ===\n")
	fmt.Printf("正确: %d/%d\n", correct, total)
	fmt.Printf("得分: %.1f%%\n", score)

	if score == 100 {
		fmt.Println("🎉 完美！太棒了！")
	} else if score >= 80 {
		fmt.Println("👍 做得很好！")
	} else if score >= 60 {
		fmt.Println("👌 不错，继续努力！")
	} else {
		fmt.Println("💪 多加练习会更好！")
	}
}
