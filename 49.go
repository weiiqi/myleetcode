package main

import "sort"

/*
思路一：使用map存储 异位数，然后遍历map输出
时间复杂度: O(n*k * log k) => n = len(strs) k = max(anagrams)
空间复杂度: O(n * k)
*/
func groupAnagrams(strs []string) [][]string {
	tmpHash := make(map[string][]string)
	for k := range strs {
		v1 := []byte(strs[k])
		sort.Slice(v1, func(i, j int) bool { return v1[i] < v1[j] })
		if _, ok := tmpHash[string(v1)]; ok {
			tmpHash[string(v1)] = append(tmpHash[string(v1)], strs[k])
		} else {
			tmpHash[string(v1)] = []string{strs[k]}
		}
	}
	ret := make([][]string, 0)
	for _, v := range tmpHash {
		ret = append(ret, v)
	}
	return ret
}

/*
思路二：计数法
通过计算每个单词出现的字母个数, 作为hash的key
*/

func groupAnagrams_2(strs []string) [][]string {
	hashT := make(map[[26]int][]string)
	for k := range strs {
		cnt := [26]int{}
		for _, v := range strs[k] {
			cnt[v-'a']++
		}
		hashT[cnt] = append(hashT[cnt], strs[k])
	}
	ret := make([][]string, 0)
	for _, v := range hashT {
		ret = append(ret, v)
	}
	return ret
}
