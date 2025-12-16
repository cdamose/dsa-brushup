package dsa_brushup

func Segregate_0_1(arr []int) []int {
	count:=0

	for i:=0;i<len(arr);i++ {
      if arr[i]==0{
		count++
	  }
	}

	for i:=0;i<count;i++{
		arr[i]=0
	}

	for i:=count;i<len(arr);i++{
		arr[i]=1
	}
	return arr

}

func Segregate_0_1_haroe_partion_algo(arr []int) []int{
	low:=0
	hi:=len(arr)-1
	for low<hi{
     if arr[low] == 1{
		arr[low] ,arr[hi] =arr[hi],arr[low]
		hi--
		
	 }else {
         low++
	 }
	}
	return arr

}