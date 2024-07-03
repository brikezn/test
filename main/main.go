package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isValid(variables []string) bool {
	if len(variables) != 3 {
		println("Not 3 args!")
		return false
	}

	firstItem := variables[0]
	secondItem := variables[2]

	if !(isDigit(firstItem) == isDigit(secondItem)) {
		println("You use incorect types")
		return false
	}

	if getDigitFromRoman(firstItem) == -1 || getDigitFromRoman(secondItem) == -1 {
		println("Not valid arguments!")
		return false
	}

	return true
}

func isDigit(input string) bool {
	if _, err := strconv.Atoi(input); err == nil {
		return true
	}

	return false
}

func getDigitFromRoman(target string) int {
	switch target {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	}

	return -1
}

func toRoman(number int) string {
	table := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IV",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	arr := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	n := number

	for i := 0; i < 9; i++ {
		if n >= arr[i] {
			n -= arr[i]
			result += table[arr[i]]
			if n > arr[i] {
				i--
			}
		}
	}

	return result
}

func run(variables []string) {
	if !isValid(variables) {
		println("Invalid arguments!")
		return
	}

	isRoman := !isDigit(variables[0])
	var variableOne int
	var variableTwo int

	if isRoman {
		variableOne = getDigitFromRoman(variables[0])
		variableTwo = getDigitFromRoman(variables[2])
	} else {
		variableOne, _ = strconv.Atoi(variables[0])
		variableTwo, _ = strconv.Atoi(variables[2])
	}

	var result int

	switch variables[1] {
	case "+":
		result = variableOne + variableTwo
	case "-":
		result = variableOne - variableTwo
	case "*":
		result = variableOne * variableTwo
	case "/":
		result = variableOne / variableTwo
	default:
		println("Not valid operation!")
	}

	if isRoman {
		if result > 0 {
			println(toRoman(result))
		} else {
			println("Roman number less than 1!")
		}
	} else {
		println(result)
	}
}

func main() {
	for {
		myScanner := bufio.NewScanner(os.Stdin)
		myScanner.Scan()
		line := myScanner.Text()
		variables := strings.Split(line, " ")

		if isValid(variables) {
			run(variables)
		}
	}
}
