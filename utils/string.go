package utils

func ArrayContainString(s []interface{}, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
