package dsa_brushup

func Merge_sorted_array(arr1 []int,arr2 []int)[]int{
	result:=[]int{

	}
	i:=0
	j:=0

	for i<len(arr1) && j<len(arr2){
		if arr1[i]<arr2[j]{
			result=append(result,arr1[i])
			i++
		}else if arr2[j]<arr1[i]{
			result=append(result,arr2[j])
			j++
		}else if arr1[i]==arr2[j]{
			result=append(result,arr1[i])
			result=append(result,arr2[j])
			i++
			j++ 
		}
	}

	//iterate arr1 upto end
	for i<len(arr1){
		result=append(result,arr1[i])
		i++
	}
	//iterate arr2 upto end
	for j<len(arr2){
		result=append(result,arr2[j])

		j++
	}

	return result
  
}