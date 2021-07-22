package lib

import (
	"fmt"
)

func Ld() {
	// 声明一个接口将一个符合已定义的接口的结构体实例赋值给这个声明的接口才能调用方法
	// var vn = Adc{}
	// var inface Interface01 = vn
	// inface.Test01()

	//当结构体实现接口后，只需将结构体传入函数，即可调用接口定义的方法
	// var hero01 = Hero{}
	// var vn = Adc{
	// 	Name: "vn",
	// }
	// hero01.InterfaceTest(&vn)
	// showAbility(&vn)

	//自定义数据类型使用接口
	// var num Number = 1
	// hero01.InterfaceTest(num)

	// 类型断言使用
	var computer = Computer{}
	var deviceArr = [2]USB{Phone{
		Name: "苹果",
	}, Camera{
		Name: "佳能",
	}}
	for _, dev := range deviceArr {
		computer.Working(dev)
	}

}

//接口
//接口中定义的函数，只要结构体的方法能够全都实现，则这个接口实现
type Interface01 interface {
	Test01()
}
type Interface02 interface {
	Test02()
}

//接口可以组合
type Interface03 interface {
	Interface01
	Interface02
}

//接口也可以声明
//空接口可以进行任意赋值
var Interface04 interface{}

//struct
type Adc struct {
	Name string
}

func (adc *Adc) Test01() {
	fmt.Println("射手")
}
func (adc *Adc) Test02() {
	fmt.Println("枪")
}

type Sd struct {
	Name string
}

func (sd Sd) Test01() {
	fmt.Println("上单")
}
func (sd Sd) Test02() {
	fmt.Println("斧子")
}

type Hero struct {
	Name string
}
type Number int

func (num Number) Test01() { //自定义数据类型定义方法，也可以实现接口
	fmt.Println(num + 8)
}

// 给函数或者方法传参为接口时，可以传递任意一个实现了该接口的结构体,能传进去的结构体必实现了这个接口
//多态参数

func (hero Hero) InterfaceTest(i Interface01) {
	fmt.Println(i)
	i.Test01()
}

func showAbility(hero Interface02) {
	hero.Test02()
}

// 多态数组
//通过接口可以实现一个数组中存放不同的结构体
var usbArr = [2]Interface01{&Adc{}, Sd{}}

//类型断言
//用来判断接口赋值前的数据类型

type USB interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (phone Phone) Call() {
	fmt.Println("手机打电话")
}

type Camera struct {
	Name string
}

func (camera Camera) Start() {
	fmt.Println("相机开始工作")
}
func (camera Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct{}

func (computer Computer) Working(usb USB) {
	usb.Start()

	//类型断言，判断usb接口赋值前是否是Phone类型
	if item, ok := usb.(Phone); ok {
		item.Call()
	}
	usb.Stop()
}
