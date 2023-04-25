package main

import (
	"fmt"
	"leetcode/problems"
	"leetcode/problems/easy"
	"leetcode/problems/medium"
	"os"

	Sort "leetcode/sort"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "3":
			s := "pwwkew"
			fmt.Println(medium.LengthOfLongestSubstring(s))
		case "11":
			nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
			fmt.Println(medium.MaxArea(nums))
		case "12":
			num := 3
			fmt.Println(medium.IntToRoman(num))
		case "22":
			n := 3
			fmt.Println(medium.GenerateParenthesis(n))
		case "26":
			slice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
			fmt.Println(easy.RemoveDuplicates(slice))
		case "39":
			candidates := []int{2, 3, 6, 7}
			target := 7
			fmt.Println(easy.CombinationSum(candidates, target))
		case "46":
			nums := []int{1, 2, 3}
			fmt.Println(medium.Permute(nums))
		case "48":
			matrix := [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}
			medium.Rotate(matrix)
		case "49":
			strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
			fmt.Println(medium.GroupAnagrams(strs))
		case "59":
			n := 3
			fmt.Println(medium.GenerateMatrix(n))
		case "62":
			m := 3
			n := 7
			fmt.Println(medium.UniquePaths(m, n))
		case "64":
			grid := [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}
			fmt.Println(medium.MinPathSum(grid))
		case "122":
			prices := []int{2, 1, 2, 1, 0, 1, 2}
			fmt.Println(medium.MaxProfit(prices))
		case "108":
			slice := []int{-10, -3, 0, 5, 9}
			fmt.Println(easy.SortedArrayToBST(slice))
		case "131":
			s := "ababbbabbaba"
			fmt.Println(medium.Partition(s))
		case "136":
			s := []int{2, 2, 1}
			fmt.Println(easy.SingleNumber(s))
		case "137":
			s := []int{0, 1, 0, 1, 0, 1, 99}
			fmt.Println(medium.SingleNumber2(s))
		case "169":
			nums := []int{3, 2, 3}
			fmt.Println(easy.MajorityElement(nums))
		case "242":
			s := "anagram"
			t := "nagaram"
			fmt.Println(easy.IsAnagram(s, t))
		case "318":
			word := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
			fmt.Println(medium.MaxProduct(word))
		case "338":
			n := 5
			fmt.Println(easy.CountBits(n))
		case "344":
			s := []byte("hello")
			fmt.Println(easy.ReverseString(s))
		case "349":
			nums1 := []int{1, 2, 2, 1}
			nums2 := []int{2, 2}
			fmt.Println(easy.Intersection349(nums1, nums2))
		case "389":
			s := "abcd"
			t := "abcde"
			fmt.Println(easy.FindTheDifference(s, t))
		case "412":
			n := 15
			fmt.Println(easy.FizzBuzz(n))
		case "496":
			nums1 := []int{1, 3, 5, 2, 4}
			nums2 := []int{5, 4, 3, 2, 1}
			fmt.Println(easy.NextGreaterElement(nums1, nums2))
		case "500":
			slice := []string{"Hello", "Alaska", "Dad", "Peace"}
			fmt.Println(easy.FindWords(slice))
		case "557":
			s := "Let's take LeetCode contest"
			fmt.Println(easy.ReverseWords(s))
		case "561":
			slice := []int{6, 2, 6, 5, 1, 2}
			fmt.Println(easy.ArrayPairSum(slice))
		case "657":
			s := "LL"
			fmt.Println(easy.JudgeCircle(s))
		case "771":
			jewels := "aA"
			stones := "aAAbbbb"
			fmt.Println(easy.NumJewelsInStones(jewels, stones))
		case "905":
			nums := []int{3, 1, 2, 4}
			fmt.Println(easy.SortArrayByParity(nums))
		case "961":
			nums := []int{1, 2, 3, 3}
			fmt.Println(easy.RepeatedNTimes(nums))
		case "1217":
			position := []int{1, 2, 3}
			fmt.Println(easy.MinCostToMoveChips(position))
		case "1221":
			s := "RLRRLLRLRL"
			fmt.Println(easy.BalancedStringSplit(s))
		case "1323":
			num := 9669
			fmt.Println(easy.Maximum69Number(num))
		case "1827":
			nums := []int{1, 5, 2, 4, 1}
			fmt.Println(easy.MinOperations(nums))
		case "2160":
			num := 2932
			fmt.Println(easy.MinimumSum(num))
		case "2248":
			slice := [][]int{
				{3, 1, 2, 4, 5},
				{1, 2, 3, 4},
				{3, 4, 5, 6},
			}
			fmt.Println(easy.Intersection2248(slice))
		case "2482":
			slice := [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{0, 0, 1},
			}
			fmt.Println(medium.OnesMinusZeros(slice))
		case "bubble":
			nums := []int{40, 20, 30, 10, 60, 50}
			Sort.BubbleSort(nums)
		case "insert":
			nums := []int{40, 20, 30, 10, 60, 50}
			Sort.InsertSort(nums)
		case "select":
			nums := []int{40, 20, 30, 10, 60, 50}
			Sort.SelectSort(nums)
		case "learn":
			problems.Learn()
		case "try":
			problems.Try()
		}
	}
}
