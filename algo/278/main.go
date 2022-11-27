package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start := 1
	end := n
	fb := n
	for i := n / 2; i >= start && i <= end; i = (start + end) / 2 {
		if isBadVersion(i) {
			fb = i
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return fb
}

func isBadVersion(n int) bool {
	if n >= 1 {
		return true
	}
	return false
}

func main() {
	firstBadVersion(4)
}
