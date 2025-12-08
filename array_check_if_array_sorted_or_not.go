package dsa_brushup

func Check_array_sorted_or_not(arr []int)bool{
	for i:=0;i<len(arr)-1;i++{
     if arr[i]>arr[i+1]{
		return false
	 }
	}
	return true

}
