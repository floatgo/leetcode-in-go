func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    
    if n == 1 || n == 2 {
        return 1
    }
    
    tribs := make([]int, n+1)
    
    tribs[0] = 0
    tribs[1] = 1
    tribs[2] = 1
    
    for i:=3; i<len(tribs); i++ {
        tribs[i] = tribs[i-1] + tribs[i-2] + tribs[i-3]
    }
    
    return tribs[n]
    
}
