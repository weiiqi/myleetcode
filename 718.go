package main

import "fmt"

func findLength(A []int, B []int) int {
	m,n := len(A), len(B)
	dp := make([][]int, m+1)
	for k := range dp{
		dp[k] = make([]int, n+1)
	}
	maxL := 0
	for i:= m-1; i>=0; i--{
		for j := n-1;j>=0;j--{
			if A[i] == B[j]{
				dp[i][j] = dp[i+1][j+1]+1
			}else{
				dp[i][j] = 0
			}
			if maxL < dp[i][j]{
				maxL = dp[i][j]
			}
		}
	}
	return maxL
}

func main()  {
	A:= []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	fmt.Print(findLength(A,B))
}