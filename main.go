package main

import (
	"strings"
)

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
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

// wrong it not prefix it check all sub string TT
func longestCommonPrefix(strs []string) string {

	i := 0
	s := strs[i]
	for in, v := range strs {
		if len(s) > len(v) {
			s = v
			i = in
		}
	}

	ns := append(strs[:i], strs[i+1:]...)
	var r string
	found := false

	for i := 0; i < len(s); i++ {
		t := s[:i+1]
		for _, v := range ns {
			if !strings.Contains(v, t) {
				found = true
				break
			}
		}
		if found {
			break
		}
		r = t
	}

	// OUT break [to] keyword

	// out:
	// 	for i := 0; i < len(s); i++ {
	// 		t := s[:i+1]
	// 		for _, v := range ns {
	// 			if !strings.Contains(v, t) {
	// 				break out
	// 			}
	// 		}
	// 		r = t
	// 	}

	return r
}

func longestCommonPrefixChatGPT(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Find the shortest string in the slice
	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	// Iterate over characters of the shortest string
	for i := range shortest {
		// Check if the character at index i is common in all strings
		for _, str := range strs {
			if str[i] != shortest[i] {
				return shortest[:i] // Return the prefix found so far
			}
		}
	}

	return shortest // All strings are identical up to the length of the shortest one
}
