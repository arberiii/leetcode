package getDecimalValue
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    result := head.Val
    for head.Next != nil {
        head = head.Next
        result = result << 1 | head.Val
    }
    return result
}
