package flatten

import (
	"fmt"
	"regexp"
	"strconv"
)

var reg = regexp.MustCompile(`[^-+\d]+`)

func Flatten(nested interface{}) []interface{} {
	sanitized := reg.Split(fmt.Sprintf("%v", nested), -1)
	flat := []interface{}{}

	for _, e := range sanitized {
		if n, err := strconv.Atoi(e); err == nil {
			flat = append(flat, n)
		}
	}

	return flat
}
