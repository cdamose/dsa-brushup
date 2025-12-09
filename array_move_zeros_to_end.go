package dsa_brushup

func Move_zero_to_end(arr []int) []int {
	pointer:=0
	for i:=0;i<len(arr);i++{
		if arr[i]!=0{
			arr[pointer]=arr[i]
			pointer++
		}
	}

	for pointer<len(arr){
		arr[pointer]=0
		pointer++
	}
	return arr
}