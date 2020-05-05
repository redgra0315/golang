package main

import "fmt"

func vatIntValue() {
	a := 3
	b := "中国"
	fmt.Print(a, b)

}
func main() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
	vatIntValue()
}
