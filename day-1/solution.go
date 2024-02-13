package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func Solution1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	res := 0

	for scanner.Scan() {
		digits := make([]rune, 0)
		line := scanner.Text()

		for _, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
			}
		}

		firstDigit := digits[0]
		secondDigit := digits[len(digits)-1]
		newInt, _ := strconv.Atoi(string(firstDigit) + string(secondDigit))

		res += newInt

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func Solution2(r io.Reader) int {
	words := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	scanner := bufio.NewScanner(r)
	res := 0

	for scanner.Scan() {
		digits := make([]rune, 0)
		line := scanner.Text()

		for i, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
			} else {
				for word := range words {
					if strings.HasPrefix(line[i:], word) {
						digits = append(digits, words[word])
					}
				}
			}
		}

		firstDigit := digits[0]
		secondDigit := digits[len(digits)-1]
		newInt, _ := strconv.Atoi(string(firstDigit) + string(secondDigit))

		res += newInt

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
