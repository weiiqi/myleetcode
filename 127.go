package main

import (
	"fmt"
	"math"
)

/*
思路： 广度优先遍历BFS
拆分单词，使用虚拟节点，马冬梅 -> 马*梅， *冬梅， 马冬*
通过单词列表和虚拟节点，构建图，然后进行广度遍历，得出最短路径

时间复杂度: O(N * C^2) -> N 单词表长度，C单词长度
空间复杂度: O(N * C^2) -> N 单词表长度，C单词长度
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, ok := wordId[endWord]
	if !ok {
		return 0
	}
	const inf int = math.MaxInt64
	dlist := make([]int, len(wordId))
	for i := range dlist {
		dlist[i] = inf
	}
	dlist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dlist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dlist[w] == inf {
				dlist[w] = dlist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

func castlistTomap(wordl []string) map[string]bool {
	tmp := make(map[string]bool)
	for _, v := range wordl {
		tmp[v] = true
	}
	return tmp
}

func ladderLength_v2(beginWord string, endWord string, wordList []string) int {
	beginSet, endSet := make(map[string]bool), make(map[string]bool)
	wordL := castlistTomap(wordList)
	length := 1
	visited := make(map[string]bool)
	beginSet[beginWord] = true
	endSet[endWord] = true
	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		tmp := make(map[string]bool)
		for k := range beginSet {
			tmpk := []byte(k)
			for i := 0; i < len(k); i++ {
				for c := 'a'; c < 'z'; c++ {
					old := tmpk[i]
					tmpk[i] = byte(c)
					target := string(tmpk)
					if _, ok := endSet[target]; ok {
						return length + 1
					}
					if !visited[target] && wordL[target] {
						tmp[target] = true
						visited[target] = true
					}
					tmpk[i] = old
				}
			}
		}
		beginSet = tmp
		length++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength_v2("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
