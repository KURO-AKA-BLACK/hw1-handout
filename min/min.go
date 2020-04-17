package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	// TODO: implement this function.
	if arr == nil {
		return 0
	} else {
		var minimum int = 0
		for i:=0; i < len(arr); i++ {
			if i == 0 {
				minimum = arr[0]
			} else if minimum > arr[i] {
				minimum = arr[i]
			}
		}
		return minimum
	}
}
