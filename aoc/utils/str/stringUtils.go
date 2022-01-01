package str

import (
	"strconv"
	"strings"
)

func Diff(value string, subtract ...string) string {
	for _, s := range subtract {
		letters := strings.Split(s, "")
		for _, letter := range letters {
			value = strings.ReplaceAll(value, letter, "")
		}

	}
	return value
}

func Intersection(list1 string, list2 string) string {
	// Intersection is a method that gives back the characters that are contained in both strings
	resArray := make([]string, 0)
	for _, e := range list1 {
		if strings.Contains(list2, string(e)) {
			resArray = append(resArray, string(e))
		}
	}
	return strings.Join(resArray, "")
}

func StrArrayToIntArray(list1 []string) []int {
	intArray := make([]int, len(list1))
	for i, _ := range list1 {
		e, _ := strconv.Atoi(list1[i])
		intArray[i] = e
	}
	return intArray
}

func Contains(str string, a []string) bool {
	for _, s := range a {
		if str == s {
			return true
		}
	}
	return false
}
