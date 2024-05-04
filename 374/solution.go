package _374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func GuessNumber(n int) int {

	lo := 1
	hi := n

	for lo <= hi {
		mid := lo + (hi-lo)/2

		result := guess(mid)

		if result == 0 {
			return mid
		} else if result == -1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1

}
