package dsa_brushup

func Find_missing_number_in_array(arr []int) int{
	//using formula 
	sum:=0
	length:=len(arr)
	for i:=0;i<length;i++{
		sum=sum+arr[i]
	}
	length=length+1
	t_sum:= (length *(length+1))/2
	return t_sum-sum
}