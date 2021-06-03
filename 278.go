package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 思路： 二分法搜索
func firstBadVersion(n int) int {
	start := 0
	end := n
	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else if isBadVersion(mid) == false {
			start = mid
		}
	}
	if isBadVersion(start) {
		return start
	}
	return end
}
