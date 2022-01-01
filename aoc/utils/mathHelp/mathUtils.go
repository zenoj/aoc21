package mathHelp

import "math"

func IntIntersection(list1, list2 []int) []int {
	resArray := make([]int, len(list1))
	for _, e := range list1 {
		if IntContains(e, list2) {
			resArray = append(resArray, e)
		}
	}
	return resArray
}

func IntContains(e int, list []int) bool {
	for _, i := range list {
		if e == i {
			return true
		}
	}
	return false
}

func MaxInt(list ...int) int {
	max := math.MinInt
	for _, i2 := range list {
		if i2 > max {
			max = i2
		}
	}
	return max
}

func MinInt(list ...int) int {
	min := math.MaxInt
	for _, i2 := range list {
		if i2 < min {
			min = i2
		}
	}
	return min
}

func Sum(list1 []int) int {
	sum := 0
	for _, i2 := range list1 {
		sum += i2
	}
	return sum
}

func Product(list1 []int) int {
	p := 1
	for _, e := range list1 {
		p *= e
	}
	return p
}
