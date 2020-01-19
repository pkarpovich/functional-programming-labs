package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	pas := []int{2, 3, 1}
	tax := []int{10, 11, 97, 19, 18}

	taxiResult := findTaxi(pas, tax)
	fmt.Printf("Taxi result: %v\n", taxiResult)

	text := "public class Person {private final String name; private final int age;}"
	result := parseText(text)

	fmt.Printf("Parsed text: %v\n", result)
}

func findTaxi(pas []int, tax []int) (result [][]int) {
	sort.Ints(pas)
	sort.Ints(tax)

	for index, item := range pas {
		result = append(result, []int{item, tax[index]})
	}

	return
}

func parseText(input string) map[string]int {
	result := make(map[string]int)

	r, _ := regexp.Compile("[a-z]+")

	for _, item := range r.FindAllString(input, -1) {
		if val, ok := result[item]; ok {
			result[item] = val + 1
			continue
		}

		result[item] = 1
	}

	return result
}
