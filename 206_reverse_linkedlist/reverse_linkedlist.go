/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    // initialize prev to a nil pointer
    tmp := ListNode{10,nil}
    prev := tmp.Next
    
    curr := head
    
    for curr != nil {
        nxt := curr.Next
        
        curr.Next = prev
        
        prev = curr
        curr = nxt
    }
    
    return prev
}
