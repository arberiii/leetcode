/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return addDigits(l1, l2, 0)
}

func addDigits(l1 *ListNode, l2 *ListNode, curry int) *ListNode {
    if l1 == nil && l2 == nil {
        if curry != 0 {
            return &ListNode{Val: curry, Next: nil}
        }
        return nil
    }
    val1, val2 := 0, 0
    var next1, next2 *ListNode
    if l1 != nil {
        val1 = l1.Val
        next1 = l1.Next
    }
    if l2 != nil {
        val2 = l2.Val
        next2 = l2.Next
    }
    
    sum := val1 + val2 + curry
    next := addDigits(next1, next2, sum / 10)
    return &ListNode{Val: sum % 10, Next: next}
}
