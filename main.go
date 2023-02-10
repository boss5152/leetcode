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
		fmt.Println(problems.LengthOfLongestSubstring(s))
	case 26:
		slice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		fmt.Println(problems.RemoveDuplicates(slice))
	case 39:
		candidates := []int{2, 3, 6, 7}
		target := 7
		fmt.Println(problems.CombinationSum(candidates, target))
	case 108:
		slice := []int{-10, -3, 0, 5, 9}
		fmt.Println(problems.SortedArrayToBST(slice))
	case 136:
		s := []int{2, 2, 1}
		fmt.Println(problems.SingleNumber(s))
	case 137:
		s := []int{0, 1, 0, 1, 0, 1, 99}
		fmt.Println(problems.SingleNumber2(s))
	case 169:
		nums := []int{3, 2, 3}
		fmt.Println(problems.MajorityElement(nums))
	case 242:
		s := "anagram"
		t := "nagaram"
		fmt.Println(problems.IsAnagram(s, t))
	case 338:
		n := 5
		fmt.Println(problems.CountBits(n))
	case 344:
		s := []byte("hello")
		fmt.Println(problems.ReverseString(s))
	case 349:
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2}
		fmt.Println(problems.Intersection349(nums1, nums2))
	case 389:
		s := "abcd"
		t := "abcde"
		fmt.Println(problems.FindTheDifference(s, t))
	case 412:
		n := 15
		fmt.Println(problems.FizzBuzz(n))
	case 496:
		nums1 := []int{1, 3, 5, 2, 4}
		nums2 := []int{5, 4, 3, 2, 1}
		fmt.Println(problems.NextGreaterElement(nums1, nums2))
	case 500:
		slice := []string{"Hello", "Alaska", "Dad", "Peace"}
		fmt.Println(problems.FindWords(slice))
	case 557:
		s := "Let's take LeetCode contest"
		fmt.Println(problems.ReverseWords(s))
	case 561:
		slice := []int{6, 2, 6, 5, 1, 2}
		fmt.Println(problems.ArrayPairSum(slice))
	case 657:
		s := "LL"
		fmt.Println(problems.JudgeCircle(s))
	case 771:
		jewels := "aA"
		stones := "aAAbbbb"
		fmt.Println(problems.NumJewelsInStones(jewels, stones))
	case 905:
		nums := []int{3, 1, 2, 4}
		fmt.Println(problems.SortArrayByParity(nums))
	case 1217:
		position := []int{1, 2, 3}
		fmt.Println(problems.MinCostToMoveChips(position))
	case 1221:
		s := "RLRRLLRLRL"
		fmt.Println(problems.BalancedStringSplit(s))
	case 1323:
		num := 9669
		fmt.Println(problems.Maximum69Number(num))
	case 2160:
		num := 2932
		fmt.Println(problems.MinimumSum(num))
	case 2248:
		slice := [][]int{
			{3, 1, 2, 4, 5},
			{1, 2, 3, 4},
			{3, 4, 5, 6},
		}
		fmt.Println(problems.Intersection2248(slice))
	default:
		problems.Learn()
	}
}
