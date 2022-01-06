func climbStairs(n int) int {
    if n == 0 {
        return 1
    }
    
    if n == 1 {
        return 1
    }
    
    ways := make([]int, n+1)
    
    ways[0] = 1
    ways[1] = 1
    
    for i := 2; i < len(ways); i++ {
        ways[i] = ways[i-1] + ways[i-2]
    }
    
    return ways[n]
    
}
