package dsa_brushup

func Array_serach(arr []int,key int)bool{
	for i:=0;i<len(arr);i++{
		if arr[i]==key {
			return true
		}
	}
	return false
}