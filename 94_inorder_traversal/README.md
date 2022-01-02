### [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

Solution Summary:
- I have added two approaches in the solution. Both approaches use the same logic but the first one is slightly verbose.
- Approach 1: Define a slice which will be populated by a nested function. The nested function will use standard inorder traversal logic to populate the slice. We finally return the slice.
- Approach 2: The inorder logic is the same, however we create the slice on the fly instead of declaring it separately. We first call inorder on the left and then on the right. We first add the left elements, then the root element and finally the right elements and then return the slice.
