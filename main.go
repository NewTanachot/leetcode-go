package main

import (
	"fmt"
	"strings"
)

type Covid struct {
	ConfirmDate    string
	No             int
	Age            *int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

type tets struct {
	Data []Covid
}

func main() {

	// c := make(chan string)

	// go func() {
	// 	c <- "a"
	// 	time.Sleep(time.Second * 1)
	// 	c <- "b"
	// 	time.Sleep(time.Second * 1)
	// 	c <- "c"
	// 	close(c)
	// }()

	// for v := range c {
	// 	println(v)
	// }

	// if a is int, a must be 0 value
	// if a is string, a must be empty value
	// a := <-c
	// fmt.Println(a)

	// go func() {
	// 	for v := range c {
	// 		println(v)
	// 	}
	// }()

	// c <- 1
	// time.Sleep(time.Second * 2)
	// c <- 2
	// time.Sleep(time.Second * 2)
	// c <- 3
	// time.Sleep(time.Second * 2)
	// c <- 4
	// time.Sleep(time.Second * 2)
	// c <- 5

	// a := make([]string, 0, 5)
	// b := []string{}

	// a = append(a, "1", "2", "3", "4", "5", "6")
	// b = append(b, "1", "2")

	// fmt.Println(a, cap(a))
	// fmt.Println(b, cap(b))

	// fmt.Println("start")
	// defer fmt.Println("nndsa;dkakljbkaebkd")
	// log.Fatal("this is fatal")
	// fmt.Println("never happen")

	// response, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")

	// if err != nil {
	// 	// return nil, err
	// 	fmt.Println(err)
	// }

	// body, err := io.ReadAll(response.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// result := new(tets)

	// if err = json.Unmarshal(body, result); err != nil {
	// 	fmt.Println(err)
	// }

	// a := 0
	// for _, i := range result.Data {
	// 	// fmt.Println(i.Age)
	// 	if i.Age != nil && *i.Age == 0 {
	// 		a++
	// 		fmt.Println(i)
	// 	}
	// }

	// println(a)

	// var waitGroup sync.WaitGroup
	// var mutex sync.Mutex
	// var str []string

	// for i := 0; i < 5; i++ {
	// 	waitGroup.Add(1)
	// 	go func() {
	// 		defer waitGroup.Done()
	// 		mutex.Lock()
	// 		str = append(str, "*")
	// 		mutex.Unlock()
	// 	}()
	// }

	// waitGroup.Wait()
	// fmt.Println(str)

	// fmt.Println(time.Now().Unix())

	// ages := map[string]int{
	// 	"john": 18,
	// }

	// if _, hasValue := ages["john"]; hasValue {
	// 	fmt.Println("john found")
	// } else {
	// 	fmt.Println("john not found")
	// }

	// item := "Chocolate"
	// price := 15.33333

	// line := fmt.Sprintf("%s (%.2f)", item, price)

	// fmt.Println(line)

	// type SimpleStruct struct {
	// 	GreetingMessage string `json:"greeting_message"`
	// }

	// var employees []string
	// employees := []string{"new", "not", "main"}
	// fmt.Println(employees)

	// for _, employee := range employees {
	// 	fmt.Println(employee)
	// }

	// if result, err := json.Marshal(&SimpleStruct{GreetingMessage: "some message"}); err == nil {
	// 	fmt.Println(string(result))
	// }

	// fmt.Println("")

	// a := searchInsert([]int{1, 3, 5, 6}, 5)
	// fmt.Println(a)
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left >= right {
		mid := (right + left) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right += mid - 1
		} else {
			left += mid + 1
		}
	}

	return -1
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := new(ListNode)
	head := newList

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = list1
			list1 = list1.Next
		} else {
			newList.Next = list2
			list2 = list2.Next
		}

		newList = newList.Next
	}

	if list1 == nil {
		newList.Next = list2
	} else {
		newList.Next = list1
	}

	return head.Next
}

// to do
func searchInsert_LogN(nums []int, target int) int {
	return 0
}

func searchInsert_ON(nums []int, target int) int {
	index := 0
	for i, v := range nums {
		if v >= target {
			return i
		}

		index++
	}

	return index
}

func isValidMap(s string) bool {
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	} // can use number of char in ascii table instead of hashmap

	r := []rune{}
	for _, i := range s {
		println(string(m[i]))
		if v := m[i]; v != 0 {
			l := len(r)
			if l != 0 {
				if v == r[l-1] {
					r = r[:l-1]
					continue
				}
			}

			return false
		}

		r = append(r, i)
		println(string(r))
	}

	if len(r) != 0 {
		return false
	}

	return true
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	r := []rune{}
	for _, v := range s {
		if l := len(r); l == 0 {
			r = append(r, v)
		} else {

			if j := l - 1; v-1 == r[j] || v-2 == r[j] {
				r = r[:j]
			} else {
				r = append(r, v)
			}
		}
	}

	return len(r) == 0
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

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		fmt.Println(mid)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// binar searh for get value
// func search(nums []int, target int) int {
// 	tempSlice := nums
// 	for {
// 		mIndex := len(tempSlice) / 2
// 		fmt.Println(tempSlice)
// 		fmt.Println(mIndex)

// 		// fmt.Println(nums[mIndex])
// 		// fmt.Println(target)
// 		if tempSlice[mIndex] > target {
// 			tempSlice = tempSlice[:mIndex]
// 		} else if tempSlice[mIndex] < target {
// 			tempSlice = tempSlice[mIndex+1:]
// 		} else {
// 			return mIndex
// 		}
// 	}
// }

// func generate(numRows int) [][]int {
// 	tri := [][]int{}

// 	for i := 0; i < numRows; i++ { // 2
// 		if i == 0 {
// 			tri = append(tri, []int{1})
// 			continue
// 		}

// 		shard := []int{}
// 		shard = append(shard, 1)

// 		for j := 0; j < i-1; j++ { // 1
// 			shard = append(shard, tri[i-1][j]+tri[i-1][j+1])
// 		}

// 		shard = append(shard, 1)
// 		tri = append(tri, shard)
// 	}

// 	return tri
// }

// func singleNumber(nums []int) int {
// 	tempMap := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		tempMap[nums[i]]++
// 	}

// 	for key, value := range tempMap {
// 		if value == 1 {
// 			return key
// 		}
// 	}

// 	return 0
// }

// func searchInsert(nums []int, target int) int {
// 	for i, n := range nums {
// 		if target <= n {
// 			return i
// 		}
// 	}

// 	return 0
// }

// func longestCommonPrefix(strs []string) string {
// 	r := ""
// 	var shortestIndex = 0

// 	for i := 0; i < len(strs); i++ {
// 		if strs[i] == "" {
// 			return r
// 		}

// 		if i == 0 {
// 			shortestIndex = i
// 		} else {
// 			if len(strs[i]) < len(strs[shortestIndex]) {
// 				shortestIndex = i
// 			}
// 		}
// 	}

// 	mainStr := strs[shortestIndex]
// 	for i := 0; i < len(mainStr); i++ {
// 		for j := 0; j < len(strs); j++ {
// 			if j != shortestIndex && strs[j][i] != mainStr[i] {
// 				return r
// 			}
// 		}

// 		r += string(mainStr[i])
// 	}

// 	return r
// }

// func romanToInt(s string) int {
// 	mapNumber := map[rune]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}

// 	result := 0
// 	sLen := len(s)
// 	for i := 0; i < sLen; i++ {
// 		if i+1 != sLen && mapNumber[rune(s[i])] < mapNumber[rune(s[i+1])] {
// 			result += mapNumber[rune(s[i+1])] - mapNumber[rune(s[i])]
// 			i++
// 		} else {
// 			result += mapNumber[rune(s[i])]
// 		}
// 	}

// 	return result
// }

// func twoSum(nums []int, target int) []int {
// 	// for i := 0; i < len(nums); i++ {
// 	// 	for j := 1; j < len(nums); j++ {
// 	// 		if nums[i]+nums[j] == target {
// 	// 			return []int{i, j}
// 	// 		}
// 	// 	}
// 	// }

// 	for i := 0; i < len(nums); i++ {
// 		var ft int
// 		if target > nums[i] {
// 			ft = target - nums[i]
// 		} else {
// 			ft = nums[i] - target
// 		}

// 		for index, v := range nums {
// 			if v == ft && index != i {
// 				return []int{index, i}
// 			}
// 		}
// 	}

// 	return nil
// }

// func intersection(nums1 []int, nums2 []int) []int {
// 	index1 := 0
// 	r := []int{}

// 	for len(nums1) != 0 && len(nums2) != 0 {
// 		for index2, v := range nums2 {
// 			if nums1[index1] == v {
// 				r = append(r, v)
// 				nums1 = append(nums1[:index1], nums1[index1+1:]...)
// 				nums2 = append(nums2[:index2], nums2[index2+1:]...)

// 				fmt.Println(nums1)
// 				fmt.Println(nums2)
// 			} else {
// 				nums1 = append(nums1[:index1], nums1[index1+1:]...)
// 			}

// 			break
// 		}
// 	}

// 	return r
// }

// func majorityElement(nums []int) int {
// 	g := map[int]int{}

// 	rkey := 0
// 	rvalue := 0

// 	for _, v := range nums {
// 		g[v]++
// 	}

// 	for k, v := range g {
// 		if v > rvalue {
// 			rkey = k
// 			rvalue = v
// 		}
// 	}

// 	return rkey
// }

// func moveZeroes(nums []int) {
// 	r := []int{}
// 	zeroCount := 0

// 	for _, v := range nums {
// 		if v != 0 {
// 			r = append(r, v)
// 		} else {
// 			zeroCount++
// 		}
// 	}

// 	nums = append(r, make([]int, zeroCount)...)
// 	fmt.Println(nums)
// }

// func maxProfit(prices []int) int {

// 	minIndex := 0

// 	for i := 1; i < len(prices); i++ {
// 		if prices[minIndex] > prices[i] {
// 			minIndex = i
// 		}
// 	}

// 	fmt.Println(minIndex)

// 	found := false
// 	maxIndex := minIndex
// 	for i := 0; i < len(prices); i++ {
// 		if i == minIndex {
// 			found = true
// 			continue
// 		}

// 		if found && prices[maxIndex] < prices[i] {
// 			maxIndex = i
// 		}
// 	}

// 	return prices[maxIndex] - prices[minIndex]

// minValue := slices.Min(prices)
// found := false

// for i, v := range prices {
// 	if !found && v == minValue {
// 		found = true
// 		continue
// 	}

// 	if found {

// 	}
// }
// }
