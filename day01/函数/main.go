package main

import (
	"fmt"
	"math"
	"os"
)

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("not support operate: %s", op)
	}
}
func swap(a, b int) (int, int) {
	//x, y = a ,b
	return b, a
}
func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}
func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type greeting func(name string) string

func say(g greeting, word string) {
	fmt.Println(g(word))
}
func english(name string) string {
	return "Hello," + name
}
func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[1]
	}
	return s
}
func main() {
	a, b := 10, 5
	result, err := operate(a, b, "m")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	a1, b1 := swap(a, b)
	fmt.Println(swap(a1, b1))
	file, err := os.Open("F:\\code\\go\\src\\github.com\\golang\\day01\\函数\\main.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("打开正常", file)
	}
	fmt.Println(compute(powInt, 3, 4))
	say(english, "word")
	sums := sum(1, 2, 34, 54, 354, 6534, 43, 6, 64, 3)
	fmt.Println(sums)
}
