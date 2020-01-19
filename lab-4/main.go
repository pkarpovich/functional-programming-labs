package main

import "fmt"

func main() {
	fmt.Println(modificate(func(i int) int {
		return i * i
	}, 2))

	fmt.Println(modificate(func(i int) int {
		return i * i * i
	}, 2))

	fmt.Println(modificate(func(i int) int {
		return -i
	}, 2))

	fmt.Println(filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(i int) bool {
		return i == 2
	}))

	fmt.Println(compose(sum, sum, sum)(0))

}

func modificate(f func(int) int, num int) int {
	return f(num)
}

func filter(arr []int, predicate func(int) bool) []int {
	for index, item := range arr {
		if predicate(item) {
			arr[index] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		}
	}

	return arr
}

func sum(i int) int {
	return i + 1
}

func compose(funcs ...func(int) int) func(int) int {
	return func(i int) int {
		for _, v := range funcs {
			i = v(i)
		}
		return i
	}
}
