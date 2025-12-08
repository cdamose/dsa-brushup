package dsa_brushup

func Second_largest_elemnt_in_array(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest_element := arr[0]
	second_largest := -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest_element {
			second_largest = largest_element
			largest_element = arr[i]
		
		 } else if arr[i] > second_largest || second_largest == -1 {
		 	second_largest = arr[i]
		 }
	}
	return second_largest
}