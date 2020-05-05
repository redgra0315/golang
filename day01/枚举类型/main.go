package main

import "fmt"

// 枚举类型 iota

func enumDemo() {
	const (
		Monday = 1 + iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

}
func typeDefAndAlias() {
	type MyInt1 int   // 基于int定义的一个新的类型,这个用的比较多一点(别名)
	type MyInt2 = int // MyInt2 和int 就是完全一样的 type alias

	var i int = 1
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
func main() {
	enumDemo()
	typeDefAndAlias()
}
