package dramautils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppendString(t *testing.T) {

	in := [][]string{
		{"abc", "efg", "hijklmn"},
		{"abc", "", "hijklmn"},
	}

	out := []string{
		"abcefghijklmn",
		"abchijklmn",
	}

	for i, v := range in {
		ret := AppendString(v...)
		assert.Equal(t, ret, out[i])
		// if ret != out[i] {
		// 	t.Fatalf("Test_AppendString AppendString \"%s\" != \"%s\" [ %+v ]",
		// 		ret, out[i], in[i])
		// }
	}

	t.Logf("Test_AppendString OK")
}

func Test_String2Int64(t *testing.T) {
	i64, err := String2Int64("123")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = String2Int64("123.456")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	_, err = String2Int64("123.456,")
	assert.Error(t, err)

	t.Logf("Test_String2Int64 OK")
}

func Test_String2Float64(t *testing.T) {
	f64, err := String2Float64("123")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(123))

	f64, err = String2Float64("123.456")
	assert.NoError(t, err)
	assert.Equal(t, f64, 123.456)

	_, err = String2Float64("123.456,")
	assert.Error(t, err)

	t.Logf("Test_String2Float64 OK")
}
