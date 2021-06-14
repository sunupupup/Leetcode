package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(guess int, target int) int {
	if target == guess {
		return 0
	} else if guess > target {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int, target int) int {

	left := 1
	right := n
	for left < right {
		mid := (right-left)/2 + left
		if guess(mid, target) == 0 {
			return mid
		} else if guess(mid, target) == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}
