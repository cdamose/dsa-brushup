package dsa_brushup

func Largert_element_in_array(arr []int) int {
	largest:=arr[0];
	for i:=1;i<len(arr);i++{
      if arr[i]>largest{
		largest=arr[i]
	  }
	}
	return largest
}