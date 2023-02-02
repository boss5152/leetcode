package main

import (
	"flag"
	"fmt"
	"leetcode/problems"
)

var flagInt int

func init() {
	flag.IntVar(&flagInt, "Id", 0, "Please Enter Id")
	flag.Parse()
}

func main() {
	switch flagInt {
	case 3:
		s := "pwwkew"
		ans := problems.LengthOfLongestSubstring(s)
		fmt.Print(ans)
	case 26:
		slice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		ans := problems.RemoveDuplicates(slice)
		fmt.Print(ans)
	case 108:
		slice := []int{-10, -3, 0, 5, 9}
		ans := problems.SortedArrayToBST(slice)
		fmt.Print(ans)
	case 136:
		s := []int{2, 2, 1}
		ans := problems.SingleNumber(s)
		fmt.Print(ans)
	case 137:
		s := []int{0, 1, 0, 1, 0, 1, 99}
		ans := problems.SingleNumber2(s)
		fmt.Print(ans)
	case 242:
		s := "anagram"
		t := "nagaram"
		ans := problems.IsAnagram(s, t)
		fmt.Print(ans)
	case 338:
		n := 5
		ans := problems.CountBits(n)
		fmt.Print(ans)
	case 344:
		s := []byte("hello")
		ans := problems.ReverseString(s)
		fmt.Print(ans)
	case 349:
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2}
		ans := problems.Intersection(nums1, nums2)
		fmt.Print(ans)
	case 389:
		s := "abcd"
		t := "abcde"
		ans := problems.FindTheDifference(s, t)
		fmt.Print(ans)
	case 412:
		n := 15
		ans := problems.FizzBuzz(n)
		fmt.Print(ans)
	case 557:
		s := "Let's take LeetCode contest"
		ans := problems.ReverseWords(s)
		fmt.Print(ans)
	case 657:
		s := "LL"
		ans := problems.JudgeCircle(s)
		fmt.Print(ans)
	default:
		problems.Learn()
	}
}
