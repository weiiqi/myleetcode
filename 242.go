package main

import (
	"fmt"
	"sort"
)

/*
思路一： 将连个字符排序，然后判等即可
时间复杂度: O(nlogn)
空间复杂度: O(n)
*/

func isAnagram(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

/*
思路二：
通过hash表来实现
时间复杂度: O(n)
空间复杂度: O(1) 字符集大小
*/

func isAnagram_2(s string, t string) bool {
	hasht := make(map[byte]int)
	for k := range s {
		if _, ok := hasht[s[k]]; ok {
			hasht[s[k]]++
		} else {
			hasht[s[k]] = 1
		}
	}

	for k := range t {
		if v, ok := hasht[t[k]]; ok {
			if v > 1 {
				hasht[t[k]]--
			} else {
				delete(hasht, t[k])
			}
		} else {
			return false
		}
	}
	return len(hasht) == 0
}
func main() {
	fmt.Println(isAnagram_2("anagram", "nagaram"))
}
