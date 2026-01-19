package main

import "fmt"

func func1(array []int) []int {
	var result []int
	for _, el := range array {
		new := true
		for _, el2 := range result {
			if el == el2 {
				new = false
				break
			}
		}
		if new {
			result = append(result, el)
		}
	}

	return result
}

func func2(a []int, b []int) []int {
	var result []int
	for _, el := range a {
		new := true
		for _, el2 := range result {
			if el2 == el {
				new = false
				break
			}
		}
		if new {
			result = append(result, el)
		}
	}

	for _, el := range result {
		new := false
		var index = 1
		for _, el2 := range b {
			if el2 == el {
				new = true
				break
			}
		}
		if !new {
			result = append(result[:index], result[:index+1]...)
		}
		index = index + 1
	}
	return result
}

func func3(a []int, b []int) []int {
	var result []int
	for _, el := range a {
		new := true
		for _, el2 := range result {
			if el2 == el {
				new = false
				break
			}
		}
		if new {
			result = append(result, el)
		}
	}

	for _, el := range b {
		new := true
		for _, el2 := range result {
			if el2 == el {
				new = false
				break
			}
		}
		if new {
			result = append(result, el)
		}
	}
	return result
}

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	arraya := func1(a)
	arrayb := func1(b)
	result := func2(a, b)
	result2 := func3(a, b)
	fmt.Println(arraya, arrayb)
	fmt.Println(result)
	fmt.Println(result2)
}
