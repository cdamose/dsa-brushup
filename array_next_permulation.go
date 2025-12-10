package dsa_brushup

func find_pivot(arr []int) int {
	n:=len(arr)
	for i:=n-2;i>=0;i--{
		if arr[i]<arr[i+1] {
          return i
		}
	}
	return -1
}


func find_successor(arr []int,pivot int ) int {
	n:=len(arr)
	for i:=n-1;i>pivot;i--{
		if arr[i]>arr[pivot] {
          return i
		}
	}
	return -1
}
func reverse_range(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
// Algorithm: Next Permutation
// 1. Find pivot: Scan from right to left, find rightmost position i where arr[i] < arr[i+1]
//    - This is the position we need to change to get next permutation
//    - If no such position exists (array is descending), reverse entire array
// 2. Find successor: Scan from right to left, find smallest element > arr[pivot]
//    - This will be the next value at pivot position
// 3. Swap: Swap arr[pivot] with arr[successor]
// 4. Reverse: Reverse the subarray after pivot position to get smallest arrangement
//
// Example: [1, 3, 5, 4, 2]
//   Step 1: pivot = 1 (arr[1]=3 < arr[2]=5)
//   Step 2: successor = 3 (arr[3]=4 is smallest element > 3)
//   Step 3: swap → [1, 4, 5, 3, 2]
//   Step 4: reverse after pivot → [1, 4, 2, 3, 5]
//
// Time: O(n), Space: O(1)
func Next_permutation(arr []int) []int {
	pivot := find_pivot(arr)
	
	if pivot == -1 {
		reverse_range(arr, 0, len(arr)-1)
		return arr
	}
	
	successor := find_successor(arr, pivot)
	
	arr[pivot], arr[successor] = arr[successor], arr[pivot]
	
	reverse_range(arr, pivot+1, len(arr)-1)
	
	return arr
}