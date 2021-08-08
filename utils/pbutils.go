package dramautils

import "go.uber.org/zap"

// Int32ArrToIntArr2 - []int32 to [][]int
func Int32ArrToIntArr2(arr []int32, x, y int) ([][]int, error) {
	arr2 := [][]int{}

	if len(arr) != x*y {
		Error("Int32ArrToIntArr2",
			zap.Int("len", len(arr)),
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Error(ErrInvalidArrayLength))

		return nil, ErrInvalidArrayLength
	}

	for i := 0; i < len(arr)/x; i++ {
		carr := []int{}

		for j := 0; j < x; j++ {
			carr = append(carr, int(arr[i*x+j]))
		}

		arr2 = append(arr2, carr)
	}

	return arr2, nil
}

// IntArr2ToInt32Arr - [][]int to []int32
func IntArr2ToInt32Arr(arr [][]int) ([]int32, int, int) {
	arr2 := []int32{}

	for _, arr1 := range arr {
		for _, v := range arr1 {
			arr2 = append(arr2, int32(v))
		}
	}

	return arr2, len(arr[0]), len(arr)
}

// Int32ArrToIntArr3 - []int32 to [][][]int
func Int32ArrToIntArr3(arr []int32, x, y, z int) ([][][]int, error) {
	arr3 := [][][]int{}

	if len(arr) != x*y*z {
		Error("Int32ArrToIntArr3",
			zap.Int("len", len(arr)),
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Int("z", z),
			zap.Error(ErrInvalidArrayLength))

		return nil, ErrInvalidArrayLength
	}

	for cz := 0; cz < z; cz++ {
		carr2 := [][]int{}

		for cy := 0; cy < y; cy++ {
			carr := []int{}

			for cx := 0; cx < x; cx++ {
				carr = append(carr, int(arr[cz*x*y+cy*x+cx]))
			}

			carr2 = append(carr2, carr)
		}

		arr3 = append(arr3, carr2)
	}

	return arr3, nil
}

// IntArr3ToInt32Arr - [][][]int to []int32
func IntArr3ToInt32Arr(arr [][][]int) ([]int32, int, int, int) {
	arr3 := []int32{}

	for _, arr2 := range arr {
		for _, arr1 := range arr2 {
			for _, v := range arr1 {
				arr3 = append(arr3, int32(v))
			}
		}
	}

	return arr3, len(arr[0][0]), len(arr[0]), len(arr)
}
