package dsa_brushup

func Array_duplicates(arr []int) []int {

	dict:=make(map[int]int)
	result:=[]int{}

	for i:=0;i<len(arr);i++{
		dict[arr[i]]++
	}

	for num,count:=range dict{
		if count >1 {
          result=append(result,num)  
		}
	}
	return result
}