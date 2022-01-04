func generate(numRows int) [][]int {
    result := make([][]int, 0)
    result = append(result, []int{1})
    result = append(result, []int{1,1})
    
    if numRows <= 2 {
        return result[:numRows]
    }
    
    for i := 2; i < numRows; i ++ {
        new_arr := make([]int, 0)
        
        new_arr = append(new_arr, 1)
        
        for j := 0; j < len(result[i-1])-1; j++ {
            to_append := result[i-1][j] + result[i-1][j+1]
            new_arr = append(new_arr, to_append)
        }
        
        new_arr = append(new_arr, 1)
        result = append(result, new_arr)
    }
    
    return result
}
