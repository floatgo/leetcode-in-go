### [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

Solution Summary:
- Initialize a prev pointer to Nil. Set curr = head.
- Loop through the end of the list. In each iteration, remember the next of curr. Assign curr's new next to previous pointer (thus reversing that portion of the list). Update prev to curr and curr to next.
- Return prev, because at the end of the loop, curr will be Nil.
