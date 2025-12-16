package dsa_brushup

// time o(n) space o(n)
func Array_missing_and_repeated_number(arr []int) (int, int) {
	freq := make(map[int]int)
	missing := 0
	repeated := 0
	
	for i := 0; i < len(arr); i++ {
		freq[arr[i]] = freq[arr[i]] + 1  
	}
	
	n := len(arr)
	for num := 1; num <= n; num++ {
		if freq[num] == 0 {
			missing = num  
		} else if freq[num] == 2 {
			repeated = num  
		}
	}
	return missing, repeated
}

// ┌─────────────────────────────────────────┐
// │  Array as a Checklist                   │
// ├─────────────────────────────────────────┤
// │                                         │
// │  Index 0 → Tracks number 1              │
// │  Index 1 → Tracks number 2              │
// │  Index 2 → Tracks number 3              │
// │  Index 3 → Tracks number 4              │
// │                                         │
// │  Negative = "I've seen this number"     │
// │  Positive = "Haven't seen this number"  │
// │                                         │
// └─────────────────────────────────────────┘
func Array_repeated_missing_number_using_marking_logic(arr []int) (int,int){
	missing:=0
	repeated:=0

	// Mark elements by making the value at index (arr[i]-1) negative
	for i:=0;i<len(arr);i++{
		val := arr[i]
		if val < 0 {
			val = -val
		}
		
		index := val - 1
		
		// If already negative, this number appeared before (repeated)
		if arr[index] < 0 {
			repeated = val
		} else {
			// Mark as visited by making it negative
			arr[index] = -arr[index]
		}
	}

	// Find the missing number (index with positive value)
	for i:=0;i<len(arr);i++{
		if arr[i]>0{
			missing=i+1
			break
		}
	}

	return missing,repeated
}
