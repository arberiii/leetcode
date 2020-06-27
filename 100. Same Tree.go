package isSameTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        if q == nil {
            return true
        } else {
            return false
        }
    } else {
       if q == nil {
            return false
        } 
    }
    
    if p.Val != q.Val {
        return false
    }
    var left, right bool
    if p.Left == nil {
        if q.Left != nil {
            return false
        }
        left = true
    } else if q.Left == nil {
        return false
    } else {
        left = isSameTree(p.Left, q.Left)
    }
    
    if p.Right == nil {
        if q.Right != nil {
            return false
        }
        right = true
    } else if q.Right == nil {
        return false
    } else {
        right = isSameTree(p.Right, q.Right)
    }
    
    return left && right
}
