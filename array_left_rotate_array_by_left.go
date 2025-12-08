package dsa_brushup

//array revrsal approach 

func reverse(arr []int,start int,end int)[]int{
	left:=start
	right:=end

	for left<right{
		arr[left],arr[right]= arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func Rotate_array_by_k_left(arr []int,k int) []int{

	if len(arr)==0{
		return arr
	}
	k=k% len(arr)
	
	reverse(arr,0,k-1)
	
	reverse(arr,k,len(arr)-1)

	reverse(arr,0,len(arr)-1)
	return arr
}