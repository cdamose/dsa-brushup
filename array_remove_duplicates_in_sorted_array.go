package dsa_brushup

func Remove_duplicates_sorted_array(arr []int) []int{
	new_arr:=[]int{}
	new_arr =append(new_arr,arr[0])
	for i:=0;i<len(arr)-1;i++ {
      if arr[i]!=arr[i+1]{
        new_arr =append(new_arr,arr[i+1])
	  }
	}
	return new_arr
}