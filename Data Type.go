// 数据类型的出现是为了把数据拆分成所需内存大小不同的数据

// 布尔型， 数字类型， 字符串类型
// 派生类型：指针类型、数组类型、结构化类型、Channel类型、函数类型、切片类型、接口类型、Map类型

package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main1() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
