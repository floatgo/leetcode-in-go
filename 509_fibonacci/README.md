### [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

Solution Summary:
- The Dynamic Programming way to solve this is to maintain an array of n numbers.
- Initialize 0th and 1st index with 0 and 1 respectively, and calculate arr[i] as arr[i-1] + arr[i-2] for the rest of the indices.
- Finally return arr[n]
