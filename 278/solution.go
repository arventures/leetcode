package _278

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func FirstBadVersion(n int) int {

	lo := 1
	hi := n

	for lo <= hi {
		mid := lo + (hi-lo)/2

		response := isBadVersion(mid)

		if response == true {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo

}
