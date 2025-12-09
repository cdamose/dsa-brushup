package dsa_brushup

func Maximum_consecutive_one_in_array(arr []int) int{
	count:=0
	max_count:=0
	for i:=0;i<len(arr);i++{
      if arr[i]==1{
		count=count+1
		if count>max_count{
			max_count=count
		}
	  }else{
			count=0
	}
	}
	return max_count
}