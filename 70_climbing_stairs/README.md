### [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

Solution Summary:
- It's a DP problem and a variant of the fibonacci series problem.
- Basically, it can be recursively formulated as stairs(n) = stairs(n-1) + stairs(n-2). In other words, if you start from the top of the stairs, the number of ways to reach there is the sum of two terms: 1) number of ways of reaching stairs(n-1) and 2) number of ways of reaching stairs(n-2). Because after reaching either of those stairs, you can reach the top by taking one steps or two steps, both of which are allowed.
