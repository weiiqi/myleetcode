package main

import "fmt"

/*
思路一: 回溯法
时间复杂度: O(3^n * 4^m)  总的来说这个算法是枚举算法，指数级别
空间复杂度: O(m + n) 总的来说为常数的空间复杂度
*/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	backTrack(digits, 0, "", &ret)
	return ret
}

func backTrack(digits string, index int, combination string, ret *[]string) {
	// terminate
	if index == len(digits) {
		*ret = append((*ret), combination)
		return
	}
	// process current logic
	digit := string(digits[index])
	letters := phoneMap[digit]
	for i := 0; i < len(letters); i++ {
		backTrack(digits, index+1, combination+string(letters[i]), ret)
	}
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
