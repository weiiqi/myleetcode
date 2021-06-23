package main

/*
思路一：BFS， 求最优路径
时间复杂度: O(n^m) n -> 字符长度，m -> 层数
空间复杂度: O(3n)  n -> bank 长度 , used长度， queue 长度
*/
func minMutation(start, end string, bank []string) int {
	bankMap := bankToMap(bank)
	if _, ok := bankMap[end]; !ok {
		return -1
	}
	isUsed := make(map[string]bool)

	queue := []string{start}
	count := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}
			for j := 0; j < len(curr); j++ {
				for _, c := range mutationMap[curr[j]] {
					gen := curr[:j] + c + curr[j+1:]
					if _, ok := bankMap[gen]; ok && !isUsed[gen] {
						queue = append(queue, gen)
						isUsed[gen] = true
					}
				}
			}
		}
		count++
		queue = queue[1:]
	}
	return -1
}

var mutationMap = map[uint8][3]string{
	'A': {"T", "G", "C"},
	'C': {"T", "G", "A"},
	'T': {"C", "G", "A"},
	'G': {"T", "A", "C"},
}

func bankToMap(bank []string) map[string]struct{} {
	tmp := make(map[string]struct{})
	for _, v := range bank {
		tmp[v] = struct{}{}
	}
	return tmp
}
func idxOfBank(str string, bank []string) int {
	for k, v := range bank {
		if v == str {
			return k
		}
	}
	return -1
}
