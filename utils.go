package gonigsberg

func concat(slices ...[]int) []int{
        totalLen := 0
	for _,v := range slices{
		totalLen += len(v)
	}
        concatSlice := make([]int, totalLen)
        copy(concatSlice, slices[0])
	start := 0
        for i := 1; i < len(slices); i++{
		start += len(slices[i-1])
		copy(concatSlice[start:], slices[i])
	}
        return concatSlice
}


func filter(data, exclude []int) []int {
	retSlice  := make([]int, 0)
	for _, v := range data{
		add := true
		for _, val := range exclude{
			if val == v{
				add = false
				break
			}
		}
		if add {
			retSlice = append(retSlice, v)
		}
	}
	return retSlice
}
	
func sumLength(slices ...[]int) int{
	totalLen := 0
	for _, v := range slices{
		totalLen += len(v)
	}
	return totalLen
}