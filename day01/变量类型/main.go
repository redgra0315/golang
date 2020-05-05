package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main() {
	var a int = 1
	var b int32 = 2
	//显示类型转换
	b = int32(a)
	fmt.Printf("%T %T\n", a, b)
	cpuArch := runtime.GOARCH
	intSize := strconv.IntSize
	fmt.Println(cpuArch, intSize)

	var f1 float32
	var f2 float64
	fmt.Printf("%f %f\n", f1, f2)

	var bt byte = 2
	var ru rune = '中'
	fmt.Printf("%T %T\n", bt, ru)

	//强制类型转换
	var a1, a2 = 3, 4
	c := int(math.Sqrt(float64(a1*a2 + a2*a2)))
	fmt.Println(c)
	//指针
	p := 1
	ptr := &p
	fmt.Println(ptr)
	fmt.Printf("%T %T", p, ptr)
}
