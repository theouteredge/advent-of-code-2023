package dayone

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FindCalibrationValues(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		ints := regexp.MustCompile(`\d`).FindAllString(line, -1)
		if len(ints) > 0 {
			digits := getFirstLast(ints)
			if num, err := strconv.Atoi(digits); err != nil {
				log.Printf("Error converting [%s] to int: [%v]", ints, err)
				return -1
			} else {
				log.Printf("LINE [%s] EXTRACTED %v CONVERTED [%v]", line, ints, num)
				sum += num
			}
		}
	}

	return sum
}

func getFirstLast(s []string) string {
	if len(s) > 1 {
		return s[0] + s[len(s)-1]

	} else if len(s) == 1 {
		return fmt.Sprintf("%s%s", s[0], s[0])

	} else {
		return ""
	}
}

func FindCalibrationValuesExtended(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		ints := extractNumbers(line)
		digits := getFirstLast(ints)
		if num, err := strconv.Atoi(digits); err != nil {
			log.Printf("Error converting [%s] to int: [%v]", ints, err)
			return -1

		} else {
			log.Printf("LINE [%s] EXTRACTED %v CONVERTED [%v]", line, ints, num)
			sum += num
		}
	}

	return sum
}

func extractNumbers(line string) []string {
	digits := []string{}

	for len(line) > 0 {
		if digit, i := findDigit(line); digit != "" {
			digits = append(digits, digit)

			if i == 0 {
				line = line[i+1:]
			} else {
				line = line[i:]
			}
		} else {
			line = line[1:]
		}
	}

	return digits
}

func findDigit(line string) (string, int) {
	word := ""
	for i, char := range line {
		word += string(char)
		if digit, ok := matchNumericDigit(word); ok {
			return digit, i

		} else if digit, ok := matchStringDigit(word); ok {
			return digit, i
		}
	}

	return "", -1
}

func matchNumericDigit(word string) (string, bool) {
	if len(word) > 1 {
		return "", false

	} else if _, err := strconv.Atoi(word); err != nil {
		return "", false

	} else {
		return word, true
	}
}

func matchStringDigit(word string) (string, bool) {
	if word == "one" {
		return "1", true
	} else if word == "two" {
		return "2", true
	} else if word == "three" {
		return "3", true
	} else if word == "four" {
		return "4", true
	} else if word == "five" {
		return "5", true
	} else if word == "six" {
		return "6", true
	} else if word == "seven" {
		return "7", true
	} else if word == "eight" {
		return "8", true
	} else if word == "nine" {
		return "9", true
	} else {
		return "", false
	}
}
