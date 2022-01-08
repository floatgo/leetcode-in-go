func min(a int, b int) int {
    if a <= b {
        return a
    }
    
    return b
}

func minCostClimbingStairs(cost []int) int {
    
    if len(cost) < 2 {
        return cost[len(cost)-1]
    }
    
    dp := make([]int, len(cost))
    
    dp[0] = cost[0]
    dp[1] = cost[1]
    
    for i := 2; i < len(dp); i++ {
        dp[i] = cost[i] + min(dp[i-1], dp[i-2])
    }
    
    return min(dp[len(cost)-1],dp[len(cost)-2])
}
