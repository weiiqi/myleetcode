package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	ret := make([]int, 0)
	for k, v := range nums {
		if v1, ok := tmp[v]; ok {
			ret = append(ret, v1, k)
			break
		}
		tmp[target-v] = k
	}
	return ret
}

// 2021-05-27
// 毫无思路，先使用暴力解法试试看

func twoSum20210527(nums []int, target int) []int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{i, j}
}

// 看到题解方法感觉很明朗
// 通过hash表存储target - x = y 的 x 的下表，后面在遇到 target -y = x 直接从hash表中O(1)的效率找出来

func twoSum20210527_by_hash(nums []int, target int) []int {
	hasht := make(map[int]int)
	for k, v := range nums {
		if x, ok := hasht[target-v]; ok {
			return []int{k, x}
		} else {
			hasht[v] = k
		}
	}
	return nil
}
