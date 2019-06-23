package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

// {1 2} &{1 2} {1 0} {0 0}

func main() {
	// defer 外层函数返回之后执行
	s := []int{2, 3, 5, 7, 11, 13}
	var m = map[string]Vertex{
		"Bell Labs": {
			40, -74,
		},
		"Google": {
			37, -1220840,
		},
	}
	// printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	// printSlice(s)

	fmt.Println(Vertex{1, 2}, "struct")
	ArrayDemo()
	Pointer()
	const PI = 3.14159265457
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println("print Sqrt return", Sqrt(4))
	sum := 1

	// for index := 0; index < 10; index++ {
	// 	fmt.Println(index)
	// }

	for sum < 1000 {
		sum += sum
	}

	// 没有条件的 switch 同 switch true 一样
	// switch sum {
	// case 1024:
	// 	fmt.Println(sum, "-----")
	// default:
	// 	fmt.Println("sum not equal 1024")
	// }

	// if num := 999; sum < 1000 {
	// 	fmt.Println(num + sum)
	// } else if sum > 1001 {
	// 	fmt.Println(sum + 999)
	// }

}

// Sqrt custom
func Sqrt(x float64) float64 {
	defer fmt.Println("the fun Sqrt is end")
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Pointer() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

func ArrayDemo() {
	var arr = [6]int{2, 3, 5, 7, 11, 13}
	// 数组改变其切片也会改变
	var slice = arr[1:4]
	fmt.Println(slice)
	arr[2] = 999
	fmt.Println(slice)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s, "struct array")
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
