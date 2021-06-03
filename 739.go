package main

import "fmt"

func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Print(dailyTemperaturesByStack(nums), "\n")
	fmt.Print(dailyTemperatures(nums), "\n")
}


func dailyTemperatures(T []int) []int {
	ret := make([]int,len(T))
	for i:=0;i<= len(T)-1; i++{
		for j:=i; j<= len(T)-1; j++{
			if T[i] < T[j]{
				ret[i] = j-i
				break
			}
		}
	}
	return ret
}

// 单调栈的实现
func dailyTemperaturesByStack(T []int) []int  {
	ret := make([]int, len(T))
	stack := make([]int, 0)
	for i:= 0; i<= len(T)-1; i++{
		if len(stack) <= 0{
			stack = append(stack, i)
			continue
		}
		for j := len(stack)-1; j>=0;j--{
			if T[i] > T[stack[j]]{
				ret[stack[j]] = i - stack[j]
				stack = stack[:j]
				if len(stack) <= 0{
					stack = append(stack, i)
				}
			}else{
				stack = append(stack, i)
				break
			}
		}

	}
	return ret
}