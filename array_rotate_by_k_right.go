package dsa_brushup


func Rotate_array_by_k_right(arr []int, k int) []int {
    if len(arr) == 0 {
        return arr
    }
    
    k = k % len(arr)  
    
    reverse(arr, 0, len(arr)-k-1)
    reverse(arr, len(arr)-k, len(arr)-1)
    reverse(arr, 0, len(arr)-1)  
    return arr
}
