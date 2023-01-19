package main

import (
	"flag"
	"fmt"
	"leetcode/problems"
)

var flagInt int

func init() {
	flag.IntVar(&flagInt, "Id", 12, "Please Enter Id")
	flag.Parse()
}

func main() {
	switch flagInt {
	case 3:
		s := "pwwkew"
		ans := problems.LengthOfLongestSubstring(s)
		fmt.Print(ans)
	case 242:
		s := "anagram"
		t := "nagaram"
		ans := problems.IsAnagram(s, t)
		fmt.Print(ans)
	case 344:
		s := []byte("hello")
		ans := problems.ReverseString(s)
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
	default:
		problems.Learn()
	}
}
