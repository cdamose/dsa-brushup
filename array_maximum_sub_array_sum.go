package dsa_brushup

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func Maximum_sub_array_sun_brute_force(arr []int) int{

if len(arr) == 0 {
	return 0
}	
result:=arr[0]
for i:=0;i<len(arr);i++ {
	sum:=0
	for j:=i; j< len(arr);j++ {
      sum=sum+arr[j]
	  result=max(sum,result)
	}
}
return result
}

func Maximum_sub_array_kadane_algorithm(arr []int)int{
	if len(arr)==0 {
      return 0
	}
	res:=arr[0]
	max_ending:=arr[0]
	for i:=1;i<len(arr);i++ {
     max_ending=max(max_ending+arr[i],arr[i])
	 res=max(max_ending,res)
	}
	return res
}