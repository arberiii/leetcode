package firstBadVersion

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left := 1
    for {
        if left == n {
            return left
        }
        
        midpoint := (n-left)/2+left
        if isBadVersion(midpoint) {
            n = midpoint
        } else {
            left = midpoint + 1
        }
    }
    
    return n
}
