package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

func main() {
	// неизменяемую переменную
	const constant = 5

	pureFunction("BS", "TU")

	firstClassFunction()

	f := func(v int) int {
		return v + 3
	}
	twice(f, 7)

	prefixedGreet("Good Morning")("Pavel")

	fib := memoizedFib()

	fib(10)
	println(fib(10))

	// лямбда выражение
	anonymous := func() string {
		return "anonymous"
	}

	fmt.Println(anonymous())

	// операции Map, Filter и Reduce.
	r := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	fmt.Println(r)

	r = funk.Map([]int{1, 2, 3, 4}, func(x int) int {
		return x * 2
	})

	fmt.Println(r)

	r = funk.Reduce([]int{1, 2, 3, 4}, func(acc int, i int) int {
		return acc + i
	}, 0)

	fmt.Println(r)
}

// чистую функцию
func pureFunction(a string, b string) string {
	return a + b
}

// функцию первого порядка
func firstClassFunction() {
	fmt.Println("firstClassFunction")
}

// функцию высшего порядка
func twice(f func(int) int, v int) int {
	return f(f(v))
}

func greet(greeting, name string) string {
	return fmt.Sprintf("%v %v", greeting, name)
}

// каррирование функций
func prefixedGreet(p string) func(string) string {
	return func(n string) string {
		return greet(p, n)
	}
}

// меморизацию функции
func memoizedFib() func(int) int {
	cache := make(map[int]int)
	var fn func(int) int
	fn = func(n int) int {
		if n > 0 {
			result := n * fn(n-1)

			if _, ok := cache[n]; !ok {
				cache[n] = result

				return result
			}

			fmt.Println("From cache")

			return result
		}

		return 1
	}
	return fn
}
