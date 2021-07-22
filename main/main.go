package main

import (
	"fmt"
	"project01/lib"
)

//全局匿名函数
var (
	As = func() {
		fmt.Print("as")
	}
	Bs = func() {
		fmt.Print("bs")
	}
)

// 工场模式使用lib内部结构体
// var car1 = *lib.NewCar("fote", 12)

func main() {
	// fmt.Println("使用外部包测试", configor.Config{})
	// fmt.Println(car1)
	// lib.GetOs()
	// car1.GetName() //工厂模式调用内部结构体方法
	lib.LibFunc() //调用lib包中的函数
	// person.PerFun()
	// lib.GetOs()

}
