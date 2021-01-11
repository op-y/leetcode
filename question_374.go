package leetcode

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low := 1
    high := n
    for {
        mid := (low + high) / 2
        ok := guess(mid)
        if ok < 0 {
             high = mid -1
        } else if ok > 0 {
             low = mid + 1
        } else {
            return mid
        }
    }
}
