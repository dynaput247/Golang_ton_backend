package utils

func SumStrings(strs ...string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}