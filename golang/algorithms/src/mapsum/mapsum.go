package mapsum

import "strings"

var mapSum = make(map[string]int)

func Insert(key string, value int) {
	mapSum[key] = value
}

func Sum(prefix string) int {
	var result int = 0

	for key, value := range mapSum {
		if strings.Index(key, prefix) == 0 {
			result += value
		}
	}

	return result
}
