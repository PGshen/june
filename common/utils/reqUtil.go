package utils

func GetFilter(filter map[string]interface{}) map[string]interface{} {
	filterV := make(map[string]interface{})
	for k, v := range filter {
		if v != nil && v != "" {
			filterV[k] = v
		}
	}
	return filterV
}
