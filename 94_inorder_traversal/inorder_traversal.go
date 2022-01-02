/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// APPROACH 1
func inorderTraversal(root *TreeNode) []int {
    trav := make([]int, 0)
    
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode) {
        if root != nil {
            inorder(root.Left)
            trav = append(trav, root.Val)
            inorder(root.Right)
        }
    }
    
    inorder(root)
    return trav
    
}



// APPROACH 2
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)
    
    // first add left, then add root
    left = append(left, root.Val)   
    // finally add right, to complete the inorder traversal
    return append(left, right...)

}
