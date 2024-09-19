package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	arabToRome := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение (например, II + III или 3 * 4):")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		}
		text = strings.TrimSpace(text)
		parts := strings.Fields(text)

		if len(parts) != 3 {
			fmt.Println("Неправильный формат ввода. Ожидается: число операция число.")
			continue
		}

		firstOperand := parts[0]
		operator := parts[1]
		secondOperand := parts[2]

		var a, b int
		var isRomanA, isRomanB bool

		if val, exists := romeToArab[firstOperand]; exists {
			a = val
			isRomanA = true
		} else if val, err := strconv.Atoi(firstOperand); err == nil {
			a = val
		} else {
			fmt.Println("Неправильное первое число.")
			continue
		}

		if val, exists := romeToArab[secondOperand]; exists {
			b = val
			isRomanB = true
		} else if val, err := strconv.Atoi(secondOperand); err == nil {
			b = val
		} else {
			fmt.Println("Неправильное второе число.")
			continue
		}

		if isRomanA != isRomanB {
			fmt.Println("Нельзя выполнять операции с римскими и арабскими числами.")
			continue
		}

		var result int
		var validOperation bool

		switch operator {
		case "+":
			result = a + b
			validOperation = true
		case "-":
			result = a - b
			validOperation = true
		case "*":
			result = a * b
			validOperation = true
		case "/":
			if b != 0 {
				result = a / b
				validOperation = true
			} else {
				fmt.Println("Деление на ноль.")
				continue
			}
		default:
			fmt.Println("Неправильная операция. Используйте +, -, * или /.")
			continue
		}

		if validOperation {
			if isRomanA {
				if romanResult, exists := arabToRome[result]; exists {
					fmt.Printf("Результат: %s\n", romanResult)
				} else {
					fmt.Println("Результат выходит за пределы допустимого диапазона для римских чисел.")
				}
			} else {
				fmt.Printf("Результат: %d\n", result)
			}
		}
	}
}
