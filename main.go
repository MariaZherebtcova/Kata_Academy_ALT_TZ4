package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Panic3 = "Ошибка. Можно умножать на числа от 1 до 10"
	Panic4 = "Ошибка. Можно делить на числа от 1 до 10"
	Panic5 = "Введенные данные не подходят под условия калькулятор"
	Panic6 = "Ошибка.Операнд parts[0] должен быть в кавычках"
	Panic7 = "Ошибка.Операнд parts[1] должен быть в кавычках"
	Panic8 = "Ошибка.Операнд не может быть длиннее 10 символов"
)

func operation(text string) string {
	var operator string
	for i := range text {
		switch rune(text[i]) {
		case '+':
			operator = "+"
		case '-':
			operator = "-"
		case '*':
			operator = "*"
		case '/':
			operator = "/"
		}
	}
	return operator
}

func main() {
	var input string
	fmt.Println("Введите пример:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	operator := operation(input)

	switch operator {
	//Сложение
	case "+":
		parts := strings.Split(input, "+")
		for i := range parts {
			parts[i] = strings.Trim(parts[i], " ")
		}
		if len(parts[0]) > 12 || len(parts[1]) > 12 {
			fmt.Println(Panic8)
		} else {
			if !strings.HasPrefix(parts[0], `"`) || !strings.HasSuffix(parts[0], `" `) && !strings.HasPrefix(parts[1], `"`) || !strings.HasSuffix(parts[1], `"`) {

				fmt.Println(Panic6)
			} else {

				result := parts[0] + parts[1]
				result = strings.ReplaceAll(result, `"`, "")

				fmt.Printf("%q", result)
			}
		}

	//Вычитание
	case "-":
		parts := strings.Split(input, " - ")

		for i := range parts {
			parts[i] = strings.Trim(parts[i], "-")

		}
		if len(parts[0]) > 12 || len(parts[1]) > 12 {
			fmt.Println(Panic8)
		} else {
			if !strings.HasPrefix(parts[0], `"`) && !strings.HasSuffix(parts[0], `" `) {

				fmt.Println(Panic6)
			} else {
				if !strings.HasPrefix(parts[1], `"`) && !strings.HasSuffix(parts[1], `"`) {
					fmt.Println(Panic7)

				} else {
					parts[0] = strings.Trim(parts[0], `"`)
					parts[1] = strings.Trim(parts[1], `"`)
					parts[0] = strings.ReplaceAll(parts[0], "\"", "")
					parts[1] = strings.ReplaceAll(parts[1], "\"", "")

					result := strings.Replace(parts[0], parts[1], "", -1)

					fmt.Printf("%q", result)

				}
			}

		}
	//Умножение
	case "*":
		parts := strings.Split(input, "*")
		for i := range parts {
			parts[i] = strings.Trim(parts[i], " ")
		}
		if len(parts[0]) > 12 {
			fmt.Println(Panic8)
		} else {
			if !strings.HasPrefix(parts[0], `"`) && !strings.HasSuffix(parts[0], `" `) {

				fmt.Println(Panic6)
			} else {
				num, err := strconv.Atoi(parts[1])
				if err != nil || num < 1 || num > 10 {
					fmt.Println(Panic3)
					break
				}
				result := strings.ReplaceAll(parts[0], "\"", "")
				result = strings.Repeat(result, num)
				if len(result) > 42 {

					fmt.Printf("%q", result[0:42]+"...")
				} else {
					fmt.Printf("%q", result)
				}
			}
		}
	//Деление
	case "/":
		parts := strings.Split(input, "/")
		for i := range parts {
			parts[i] = strings.Trim(parts[i], " ")

		}
		if len(parts[0]) > 12 {
			fmt.Println(Panic8)
		} else {
			if !strings.HasPrefix(parts[0], `"`) && !strings.HasSuffix(parts[0], `" `) {

				fmt.Println(Panic6)
			} else {
				num, err := strconv.Atoi(parts[1])
				if err != nil || num < 1 || num > 10 {
					fmt.Println(Panic4)

				}

				var result string
				if num == len(parts[0]) {
					result = ""
				} else {
					parts[0] = strings.ReplaceAll(parts[0], "\"", "")
					result = parts[0][:len(parts[0])/num]
				}
				fmt.Printf("%q", result)
			}
		}
	default:
		fmt.Println(Panic5)

	}
}
