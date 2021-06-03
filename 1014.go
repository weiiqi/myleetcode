package main


func maxScoreSightseeingPair(A []int) int {
	total := 0
	max := 0
	for i:=0;i<len(A);i++{
		if max + A[i] -i > total{
			total = max +A[i]-i
		}
		if A[i]+i > max{
			max = A[i]+i
		}
	}
	return total
}