package dramautils

import (
	"bytes"
	"strconv"
	"strings"
)

// AppendString - append string
func AppendString(strs ...string) string {
	var buffer bytes.Buffer

	for _, str := range strs {
		if len(str) > 0 {
			buffer.WriteString(str)
		}
	}

	return buffer.String()
}

func String2Int64(str string) (int64, error) {
	if strings.Contains(str, ".") {
		nf, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, err
		}

		return int64(nf), nil
	}

	iv, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return iv, nil
}

func String2Float64(str string) (float64, error) {
	// if strings.Contains(str, ".") {
	nf, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return nf, nil
	// }

	// iv, err := strconv.ParseInt(str, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }

	// return , nil
}
