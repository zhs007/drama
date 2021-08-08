package dramautils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetJsonString(t *testing.T) {
	s, err := GetJsonString([]byte(`{"abc":"123"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123")

	s, err = GetJsonString([]byte(`{"abc":123}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123")

	s, err = GetJsonString([]byte(`{"abc":123.456}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123.456")

	s, err = GetJsonString([]byte(`{"abc":null}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "")

	t.Logf("Test_GetJsonString OK")
}

func Test_GetJsonInt(t *testing.T) {
	i64, err := GetJsonInt([]byte(`{"abc":"123"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":123}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":123.456}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":"123.456"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":""}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(0))

	i64, err = GetJsonInt([]byte(`{"abc":null}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(0))

	t.Logf("Test_GetJsonInt OK")
}

func Test_GetJsonFloat(t *testing.T) {
	f64, err := GetJsonFloat([]byte(`{"abc":"123"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(123))

	f64, err = GetJsonFloat([]byte(`{"abc":123}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(123))

	f64, err = GetJsonFloat([]byte(`{"abc":123.456}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(123.456))

	f64, err = GetJsonFloat([]byte(`{"abc":"123.456"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(123.456))

	f64, err = GetJsonFloat([]byte(`{"abc":""}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(0))

	f64, err = GetJsonFloat([]byte(`{"abc":null}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, f64, float64(0))

	t.Logf("Test_GetJsonFloat OK")
}
