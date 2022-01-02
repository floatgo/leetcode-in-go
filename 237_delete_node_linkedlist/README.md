### [237. Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/)

Solution Summary:
- The task is to delete a node in the linked list where only the pointer to the *node to be deleted* is given.
- Since we can't access the previous node, we make the current node to mimic the next node i.e., we copy the next node's *Val* and next node's *Next* pointer to the current node. 
