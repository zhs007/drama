package dramautils

import (
	"testing"
)

func Test_Int3Arr2ToInt4Arr2(t *testing.T) {

	in := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	out := [][]int{
		{1, 2, 3, 7},
		{4, 5, 6, 7},
	}

	cout := Int3Arr2ToInt4Arr2(in, 7)
	if !IsSameIntArr2(out, cout) {
		t.Logf("Test_Int3Arr2ToInt4Arr2 Fail")

		return
	}

	t.Logf("Test_Int3Arr2ToInt4Arr2 OK")
}
