package utils

func IsStrBlank(data string) bool {
	if data == "" || len(data) == 0 {
		return true
	}
	return false
}
