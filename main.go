package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	// ошибка когда оба числа больше 10 и они разные

	var number1 string
	var number2 string
	var operator string
	var operator2 string
	romanNumbers := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	err1 := errors.New("division by zero is impossible")
	err2 := errors.New("non-integer or 0 used")
	err3 := errors.New("a number greater than 10 is used")
	err4 := errors.New("more than two operands and one operator")
	err5 := errors.New("unknown operator")
	err6 := errors.New("there are no negative numbers in the roman system")
	err7 := errors.New("different number systems are used at the same time")
	//err8 := errors.New("")
	/*
		второй способ чтения входящих данных с консоли

			reader := bufio.NewReader(os.Stdin)

			for {
				fmt.Print("Введите выражение ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
		        if len(text) > 2 {
		        fmt.Println(err4)
		        } else {
		           toNumber1, _ := strconv.Atoi(string(text[0]))
		           operator := string(text[1])
				   toNumber2, _ := strconv.Atoi(string(text[2]))
		        }
			}
	*/

	fmt.Print("Введите выражение ")
	fmt.Scanln(&number1, &operator, &number2, &operator2)

	if len(operator2) == 1 {
		fmt.Println(err4)
		return
	}

	if _, exists := romanNumbers[number2]; !exists {
		/*
			//ошибка когда оба числа больше 10 и они разные
			toNumber1, _ := strconv.Atoi(number1)
			if _, exists := romanNumbers[number1]; exists && toNumber1 > 10 {
				if _, exists := romanNumbers[number2]; exists && toNumber1 > 10 {
					fmt.Println(err3, "or", err7)
					return
				}
				fmt.Println(err3, "or", err7)
				return
			}
		*/
		if _, exists := romanNumbers[number1]; !exists {

			resultNumber1 := romanToNumbers1(number1)

			if resultNumber1 > 10 {
				fmt.Println(err3)
				return
			}

			resultNumber2 := romanToNumbers2(number2)

			if resultNumber2 > 10 {
				fmt.Println(err3)
				return
			}

			toNumber1, _ := strconv.Atoi(number1)
			toNumber2, _ := strconv.Atoi(number2)

			if toNumber1 > 10 || toNumber2 > 10 {
				fmt.Print(err3)
				return
			} else if toNumber1 == 0 || toNumber2 == 0 {
				fmt.Print(err2)
			} else {
				switch operator {
				case "+":
					//fmt.Printf("%d %s %d = %d", toNumber1, operator, toNumber2, toNumber1+toNumber2)
					fmt.Println(toNumber1 + toNumber2)
				case "-":
					//fmt.Printf("%d %s %d = %d", toNumber1, operator, toNumber2, toNumber1-toNumber2)
					fmt.Println(toNumber1 - toNumber2)
				case "*":
					//fmt.Printf("%d %s %d = %d", toNumber1, operator, toNumber2, toNumber1*toNumber2)
					fmt.Println(toNumber1 * toNumber2)
				case "/":
					if toNumber2 == 0 {
						fmt.Println(err1)
					}
					//fmt.Printf("%d %s %d = %d", toNumber1, operator, toNumber2, toNumber1/toNumber2)
					fmt.Println(toNumber1 / toNumber2)
				default:
					fmt.Println(err5)
				}
			}
		}
	}

	if _, exists := romanNumbers[number2]; exists {

		resultNumber2 := romanToNumbers2(number2)

		if _, exists := romanNumbers[number1]; exists {

			resultNumber1 := romanToNumbers1(number1)

			if resultNumber1 > 10 {
				fmt.Println(err3)
				return
			} else {
				switch operator {
				case "+":
					result1 := resultNumber1 + resultNumber2

					//fmt.Printf("%s %s %s = %s", number1, operator, number2, result)
					fmt.Printf(intToRoman(result1))
				case "-":
					if resultNumber2 > resultNumber1 {
						fmt.Println(err6)
						return
					}
					result1 := resultNumber1 - resultNumber2
					if result1 == 0 {
						fmt.Println(result1)
					}

					//fmt.Printf("%s %s %s = %s", number1, operator, number2, result)
					fmt.Printf(intToRoman(result1))
				case "*":
					result1 := resultNumber1 * resultNumber2

					//fmt.Printf("%s %s %s = %s", number1, operator, number2, result)
					fmt.Printf(intToRoman(result1))

				case "/":
					if resultNumber2 == 0 {
						fmt.Println(err1)
					}
					result1 := resultNumber1 / resultNumber2

					//fmt.Printf("%s %s %s = %s", number1, operator, number2, result)
					fmt.Printf(intToRoman(result1))
				default:

					fmt.Println(err5)
				}
			}

		} else {

			resultNumber1 := romanToNumbers1(number1)

			if resultNumber1 > 10 {
				fmt.Println(err3)
				return
			} else {
				fmt.Println(err7)
				return
			}
		}
	} else {

		resultNumber2 := romanToNumbers2(number2)

		if resultNumber2 > 10 {
			fmt.Println(err3)
			return
		} else if _, exists := romanNumbers[number1]; exists {
			fmt.Println(err7)
			return
		}
	}
}

func romanToNumbers1(number1 string) int {
	romanNumbersMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var resultNumber1 = romanNumbersMap[number1[len(number1)-1]]
	for i := len(number1) - 2; i >= 0; i-- {
		if romanNumbersMap[number1[i]] < romanNumbersMap[number1[i+1]] {
			resultNumber1 -= romanNumbersMap[number1[i]]
		} else {
			resultNumber1 += romanNumbersMap[number1[i]]
		}
	}
	return resultNumber1
}

func romanToNumbers2(number2 string) int {
	romanNumbersMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var resultNumber2 = romanNumbersMap[number2[len(number2)-1]]
	for i := len(number2) - 2; i >= 0; i-- {
		if romanNumbersMap[number2[i]] < romanNumbersMap[number2[i+1]] {
			resultNumber2 -= romanNumbersMap[number2[i]]
		} else {
			resultNumber2 += romanNumbersMap[number2[i]]
		}
	}
	return resultNumber2
}

func intToRoman(result1 int) string {
	romanNumeral := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicNumeral := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string
	for result1 > 0 {
		for i := range arabicNumeral {
			if result1 >= arabicNumeral[i] {
				result += romanNumeral[i]
				result1 -= arabicNumeral[i]
				break
			}
		}
	}
	return result
}
