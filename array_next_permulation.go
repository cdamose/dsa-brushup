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