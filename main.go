package main

import "fmt"

/*
Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).
*/
func question2(x int) (int, bool) {
	isEven := x % 2 == 0
  return n / 2, isEven
}

/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers
*/
func question3(args ...int) int {
	largest := args[0]

	for _, value := range args {
		if value > largest {
			largest = value
		}
	}
	return largest
}

/*
Using makeEvenGenerator as an example, write a
makeOddGenerator function that generates odd
numbers.

*/
func makeOddGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		i++
		return
	}
}

/*
The Fibonacci sequence is defined as: fib(0) =
0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function which can find
fib(n).
*/
func question5(n int) int {
	if n < 2 {
		return n
	}
	return question5(n-1) + question5(n-2)
}

func main() {
	// x, y := question2(2)
	// fmt.Println(x, y)

	// z := []int{
	// 	48, 96, 86, 68,
	// 	57, 82, 63, 70,
	// 	37, 34, 83, 27,
	// 	19, 97, 9, 17,
	// }
	// fmt.Println(question3(z...))

	// nextOdd := makeOddGenerator()
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())

	fmt.Println(question5(10))
}
