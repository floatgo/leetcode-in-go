func fib(n int) int {
    fibs := make([]int, n+1)
    
    fibs[0] = 0
    
    if n == 0 {
        return fibs[0]
    }
    
    fibs[1] = 1
    
    for i:=2; i<len(fibs); i++ {
        fibs[i] = fibs[i-1] + fibs[i-2]
    }
    
    return fibs[n]
}
