### [118. Pascal Triangle](https://leetcode.com/problems/pascals-triangle)

Solution Summary:
- Start with two rows: 1 and 1,1. 
- For each row from 3rd row to Nth row, add 1 at the beginning and the end of the row. In between the first 1 and the last 1, add the sum of the elements prev_row[i] and prev_row[i+1].
