package exercises

import "strconv"

func ConvertNumber(value string) (int, string) {
	num, error := strconv.Atoi(value)
	if error != nil {
		return num, "Error: " + error.Error()
	}

	if num > 100 {
		return num, "It is greater than 100"
	} else {
		return num, "It is less than 100"
	}
}
