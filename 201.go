package main

func rangeBitwiseAnd(m int, n int) int {
	count := 0
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
