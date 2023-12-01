package dayone

import (
	"log"
	"testing"
	adventofcode2023 "theouteredge/adventOfCode2023"

	"github.com/stretchr/testify/assert"
)

func Test_DayOneChallengeOne_ExampleInput(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
33sdkhsjd`

	product := FindCalibrationValues(input)
	log.Printf("The Final Product is: [%v]", product)

	assert.Greater(t, product, -1, "Product should be greater than -1")
	assert.Equal(t, 142+33, product, "Product should be 175`")
}

func Test_DayOneChallengeOne(t *testing.T) {
	if input, err := adventofcode2023.GetInputFromFile("1.1.txt"); err != nil {
		assert.Fail(t, "Error getting input: %v", err)

	} else {
		product := FindCalibrationValues(input)
		log.Printf("The Final Product is: [%v]", product)

		assert.Greater(t, product, -1, "Product should be greater than -1")
		assert.Equal(t, 54632, product, "Product should be 54632`")
	}
}

func Test_DayOneChallengeTwo_MixedNumber_Example(t *testing.T) {
	input := `8ninefivegzk7ftqbceightwogfv
8ninefivegzk7ftqbceightwogfv`
	sum := FindCalibrationValuesExtended(input)

	assert.Equal(t, 164, sum, "sum should be 164")
}

func Test_DayOneChallengeTwo_ExampleInput(t *testing.T) {
	// eightwothree should this return [8, 2, 3] or [8, 3]?
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	sum := FindCalibrationValuesExtended(input)

	assert.Greater(t, sum, -1, "sum should be greater than -1")
	assert.Equal(t, 281, sum, "sum should be 281`")
}

func Test_DayOneChallengeTwo(t *testing.T) {
	// eightwothree should this return [8, 2, 3] or [8, 3]?
	if input, err := adventofcode2023.GetInputFromFile("1.1.txt"); err != nil {
		assert.Fail(t, "Error getting input: %v", err)

	} else {
		sum := FindCalibrationValuesExtended(input)

		assert.Greater(t, sum, -1, "sum should be greater than -1")
		assert.Equal(t, 54019, sum, "sum should be 54019`")
	}
}

//54019
