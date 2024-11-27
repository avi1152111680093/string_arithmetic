package main

import (
	"strconv"
)

func StrMult(s1, s2 string) string {
	result := make([]int, len(s1)+len(s2))

	for i, digit_i := range reverse(s2) {
		carry := 0
		for j, digit_j := range reverse(s1) {
			val1, _ := strconv.Atoi(string(digit_i))
			val2, _ := strconv.Atoi(string(digit_j))

			val := val1*val2 + result[i+j] + carry

			carry = val / 10
			unit_digit := val % 10

			result[i+j] = unit_digit
		}

		result[len(result)-1] = carry
	}

	finalResult := make([]byte, len(result))
	for i := range result {
		finalResult[len(result)-1-i] = byte(result[i]) + '0'
	}

	if finalResult[0] == '0' {
		return string(finalResult[1:])
	}

	return string(finalResult)
}

func reverse(s string) string {
	newStr := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newStr[i] = s[len(s)-1-i]
	}

	return string(newStr)
}
