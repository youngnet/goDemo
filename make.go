package main

import "fmt"

func main() {
	Pic(4, 4)
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	a = append(a, 1)
	printSlice("append a", a)
	// for i, _ := range pow
	// for _, value := range pow
	// 若你只需要索引，忽略第二个变量即可。
	// for i := range pow
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8
	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
	}
	return res
}
