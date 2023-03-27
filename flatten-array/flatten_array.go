package flatten

func Flatten(nested interface{}) []interface{} {
	flat := []interface{}{}
	for _, e := range nested.([]interface{}) {
		switch t := e.(type) {
		case int:
			flat = append(flat, t)
		case []interface{}:
			flat = append(flat, Flatten(t)...)
		}
	}
	return flat
}
