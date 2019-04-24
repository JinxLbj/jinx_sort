package awesomeProject

func Insert(arrays []int) {

	var i = 1
	var j = i
	for i = 1; i < len(arrays); i++ {
		for j = i; j > 0; j-- {
			if arrays[j] < arrays[j-1] {
				var temp = arrays[j]
				arrays[j] = arrays[j-1]
				arrays[j-1] = temp
			} else {
				break
			}
		}
	}

}
