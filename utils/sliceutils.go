package dramautils

// FindInt - find a int into []int
func FindInt(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}

	return -1
}

// FindInt3 - find 3 ints into []int
func FindInt3(arr []int, x, y, z int) int {
	if len(arr)%3 == 0 {
		for i := 0; i < len(arr)/3; i++ {
			if arr[i*3] == x && arr[i*3+1] == y && arr[i*3+2] == z {
				return i * 3
			}
		}
	}

	return -1
}

// FindIntArr - find a int array into [][]int
func FindIntArr(arr [][]int, vals []int) int {
	for i, arr1 := range arr {
		if len(arr1) == len(vals) {
			isok := true
			for j, av := range arr1 {
				if av != vals[j] {
					isok = false
					break
				}
			}

			if isok {
				return i
			}
		}
	}

	return -1
}

// Int3Arr2ToInt4Arr2 - []int{x,y,z} -> []int{x,y,z,v}
func Int3Arr2ToInt4Arr2(arr [][]int, val int) [][]int {
	narr := [][]int{}
	for _, arr1 := range arr {
		carr1 := append(arr1, val)

		narr = append(narr, carr1)
	}

	return narr
}

// IsSameIntArr2 -
func IsSameIntArr2(arr0 [][]int, arr1 [][]int) bool {
	if len(arr0) == len(arr1) {
		for i := 0; i < len(arr0); i++ {
			if len(arr0[i]) != len(arr1[i]) {
				return false
			}

			for j := 0; j < len(arr0[i]); j++ {
				if arr0[i][j] != arr1[i][j] {
					return false
				}
			}
		}

		return true
	}

	return false
}

// IsSameIntArr2Ex - 只比较前x个
func IsSameIntArr2Ex2(arr0 [][]int, arr1 [][]int, x int) bool {
	if len(arr0) == len(arr1) {
		for i := 0; i < len(arr0); i++ {
			if len(arr0[i]) != len(arr1[i]) {
				return false
			}

			for j := 0; j < len(arr0[i]) && j < x; j++ {
				if arr0[i][j] != arr1[i][j] {
					return false
				}
			}
		}

		return true
	}

	return false
}

// CloneArr3 - clone a [][][]int
func CloneArr3(src [][][]int) [][][]int {
	arr := [][][]int{}

	for _, src2 := range src {
		arr2 := [][]int{}

		for _, src1 := range src2 {
			arr1 := append([]int{}, src1[0:]...)
			arr2 = append(arr2, arr1)
		}

		arr = append(arr, arr2)
	}

	return arr
}

// IsSameIntArr2Ex -
func IsSameIntArr2Ex(arr0 [][]int, arr1 [][]int32) bool {
	if len(arr0) == len(arr1) {
		for i := 0; i < len(arr0); i++ {
			if len(arr0[i]) != len(arr1[i]) {
				return false
			}

			for j := 0; j < len(arr0[i]); j++ {
				if arr0[i][j] != int(arr1[i][j]) {
					return false
				}
			}
		}

		return true
	}

	return false
}
