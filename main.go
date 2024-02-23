package main

import (
	"strings"
)

func main() {
	romanToInt("MCMXCIV")
}

func twoSum(nums []int, target int) []int {

	len := len(nums)

	for i := 0; i < len; i++ {
		for j := 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func romanToInt(s string) int {
	mapNumber := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	chars := strings.Split(s, "")
	lastIndex := len(chars) - 1

	var sum int
	tempValue := 0

	for index, value := range chars {
		currenValue := mapNumber[value]
		nextValue := 0

		if index != lastIndex {
			nextValue = mapNumber[chars[index+1]]
		}

		if nextValue != 0 && currenValue < nextValue {
			tempValue = currenValue
			continue
		}

		sum += mapNumber[value] - tempValue
		tempValue = 0
	}

	return sum
}
