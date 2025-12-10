package dsa_brushup

func ArraySortColorDutchFlagAlgorithm(arr []int) []int{
	low:=0
	mid:=0
	temp:=0
	high:=len(arr)-1

	for mid<=high {
		if arr[mid] == 0 {
			temp=arr[low]
			arr[low ]=arr[mid]
			arr[mid ]=temp
			low++
			mid++

		}else if arr[mid]== 1 {
			mid ++

		}else if arr[mid] == 2 {
			temp=arr[mid]
			arr[mid]=arr[high]
			arr[high ]=temp
			high--

		}
	}
	return arr

}

func ArraySortColorCountLogic(arr []int) []int{
	count_0,count_1,count_2:=0,0,0
	for i:=0 ;i< len(arr) ;i ++ {
		if arr[i]==0 {
			count_0++
		}else if arr[i]==1 {
			count_1++
		}else if arr[i]==2 {
			count_2++
		}
	}
	i:=0
	for count_0>0{
      arr[i]=0
	  i++
	  count_0--  
	}
	for count_1>0{
      arr[i]=1
	  i++
	  count_1--  
	}

	for count_2>0{
      arr[i]=2
	  i++
	  count_2--  
	}
	return arr
}