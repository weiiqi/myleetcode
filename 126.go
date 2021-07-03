package main

import "fmt"

/*
思路一： 自信来于 127 题，先构建图，然后使用广度优先遍历 + 回溯算法求最短路径
由于需要准确路径，所以这里不能使用虚拟节点

学习了刷友们的题解，这方面的题目基本都要参考题解，有点难, 思路不太对
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	if !dict[endWord] {
		return ret
	}
	dict[beginWord] = true
	graph := make(map[string][]string)
	distance := bfs_127(endWord, dict, graph)
	dfs_127(beginWord, endWord, &ret, []string{}, distance, graph)
	return ret
}

func dfs_127(beginWord, endWord string, res *[][]string, path []string, distance map[string]int, graph map[string][]string) {
	// terminate
	if beginWord == endWord {
		path = append(path, beginWord)
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// process current level logic
	for _, v := range graph[beginWord] {
		if distance[beginWord] == distance[v]+1 {
			path = append(path, beginWord)
			// drill down
			dfs_127(v, endWord, res, path, distance, graph)
			// reverse state
			path = path[:len(path)-1]
		}
	}
}

func expand_127(word string, dict map[string]bool) []string {
	ret := make([]string, 0)
	bs := []rune(word)
	for k, v := range bs {
		for i := 'a'; i <= 'z'; i++ {
			if v == i {
				continue
			}
			bs[k] = i
			newStr := string(bs)
			if dict[newStr] {
				ret = append(ret, newStr)
			}
		}
		bs[k] = v
	}
	return ret
}

func bfs_127(endWord string, dict map[string]bool, graph map[string][]string) map[string]int {
	distance := make(map[string]int, 0)
	queue := []string{endWord}
	distance[endWord] = 0
	for len(queue) > 0 {
		curSize := len(queue)
		for i := 0; i < curSize; i++ {
			word := queue[0]
			queue = queue[1:]
			expandNode := expand_127(word, dict)
			for _, v := range expandNode {
				graph[v] = append(graph[v], word)
				if _, ok := distance[v]; !ok {
					distance[v] = distance[word] + 1
					queue = append(queue, v)
				}
			}
		}
	}
	return distance
}

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
