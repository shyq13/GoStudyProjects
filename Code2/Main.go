package main

import "fmt"

func main() {
	// 变量

	// 1. 声明
	var age int

	// 2. 初始化
	age = 13
	// 也可以这样写：
	var name = "shyq13" // 使用这种方法定义变量时，变量类型可以省略
	sex := "男"          // 使用这种方式定义变量时，不需要写var和类型

	var school string

	// 变量使用
	fmt.Println("name:", name)
	fmt.Println("sex:", sex)
	fmt.Println("age:", age)
	fmt.Println("school:", school) // string型变量默认为空

	fmt.Println("-----------------------------------------------------------------------------")

	// 3. 一次性定义多个变量
	var a, b, c = 1, "age", "name"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	d, e, f := "sex", "id", "address"
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
}
