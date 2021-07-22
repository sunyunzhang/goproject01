package lib

import (
	"encoding/json"
	"flag"
	"fmt"
	. "fmt"
	"time"
)

func Toc() {
	// fmt.Println("lib1")
	Ld()
}

var ar = [...]int{1, 2, 3, 4}
var Slice4_1 = []string{"ad", "bd", "cd"}

//函数调用测试
func LibFunc() {
	Toc()
	// arr(&ar)
	// makeSlice()
	// makeString()

	// sort()
	// makeMap()
	// getParams()
	// getStruct()
	// testPresonMethods()
	// vis()
}

// 闭包
func Add() func(int) int {
	var n int = 1
	return func(x int) int {
		n = x + n
		return n
	}
}

//defer关键字
func DeferTest(n1 int, n2 int) int {
	defer fmt.Println("defer标记语句", n1)
	defer fmt.Println("defer标记语句", n2)
	res := n1 + n2
	fmt.Println("defer标记语句", n1+n2)
	return res
}

//字符串长度
func strlen() {
	var str = "hello"
	fmt.Println(len(str))
}

// 当前时间
func getTime() {
	var time = time.Now()

	fmt.Printf("时间是%v", time)
}

// 休眠
func sleep() {
	for i := 0; i < 20; i++ {
		fmt.Printf("时间是%v\n", time.Now())
		time.Sleep(time.Second)
	}
}

//数组
func arr(arr *[4]int) {
	// 数组声明
	// var ar = [...]int{1, 2, 3, 4}
	// ar := [...]int{1, 2}

	// 数组赋值
	// ar[0] = 9.0
	// ar[1] = 3

	// 循环数组
	// for i, v := range ar {
	// 	fmt.Printf("i=%v,v=%v", i, v)
	// }
	// for i := 0; i < len(ar); i++ {
	// 	fmt.Printf("i=%v,v=%v", i, ar[i])
	// }

	// 函数传参数组时 arr [4]int 规定数量的数组
	// 传数组指针 arr *[4]int   使用*arr

	(*arr)[0] = 888
	Println(arr, ar)

	//二维数组
	// var  dbarr  [4][4]int
	// dbarr[0][0] = 1
	// Println(dbarr)
}

//切片
func makeSlice() {
	//切片是一个引用类型
	// var ar = [...]int{1, 2, 3, 4}

	// 声明切片三种方式
	// var slice1 = ar[1:3]
	// slice2 := make([]int, 2)
	// var slice3 = []string{"ad", "cd"}

	// 切片追加
	// var slice3 = []string{"ad", "bd"} //[ad bd]
	// Println(slice3)
	// slice3 = append(slice3, "cd") //[ad bd cd]

	//切片的拷贝
	var slice4_1 = []string{"ad", "bd", "cd"}
	var slice4_2 = []string{"af", "bf"}

	copy(slice4_2, slice4_1)
	Println(slice4_2)
}

//字符串
func makeString() {
	//字符串不可修改，需转成byte数组，包含中文转换成rune数组
	var str = "ghgc的"
	Println(str)
	strbyte := []byte(str)
	strbyte2 := []rune(str)
	strbyte[0] = 'd'
	strbyte2[0] = '的'
	str2 := str[3:]
	Println(str2)
}

//冒泡排序
func sort() {
	var arr [5]int = [5]int{2, 4, 3, 1, 8}

	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				var temp int = 0
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}

	Println(arr)
}

//map
func makeMap() {
	// 三种声明方式
	// var m1 map[string]string
	// m1 = make(map[string]string, 10)

	// var m2_1 = make(map[string]string, 10)
	// var m2_2 = make(map[string]string)
	// var m3 = map[string]string{"key": "value"}

	// map赋值
	// m3["key2"] = "vvv" //存在就更改，不存在创建
	// Printf("%v", m3)

	//删除
	// delete(m3, "key2") //存在就删除，不存在就跳过

	//判断key是否存在
	// _, flag := m3["key"]
	// Printf("\nflag  %t", flag)

	//map遍历
	// var m4 = map[string]string{
	// 	"key1": "a",
	// 	"key2": "b",
	// 	"key3": "c",
	// 	"key4": "d",
	// }
	// for key, val := range m4 {
	// 	Printf("\nkey %v\nvalue %v", key, val)
	// }

	// map长度统计
	// var m5 = map[string]string{
	// 	"key1": "a",
	// 	"key2": "b",
	// 	"key3": "c",
	// 	"key4": "d",
	// }

	// Printf("\n长度 %v", len(m5))

	// map切片类似于js [{},{}]
	var mapslice = make([]map[string]string, 0)
	// mapslice = append(mapslice, map[string]string{  //追加map切片值
	// 	"key1": "a",
	// 	"key2": "b",
	// 	"key3": "c",
	// 	"key4": "d",
	// })
	// Printf("mapslice %v", mapslice)

	// 终端输入追加切片
	// 终端输入方式

	var key, value string
	for i := 0; i < 4; i++ {
		// Printf("请输入第%dkey", i)
		// Scanln(&key)
		// Printf("请输入第%d个", i)
		// Scanf("key:%s value:%s", &key, &value)
		Scanf("key:%s,value:%s", &key, &value)
		mapslice = append(mapslice, map[string]string{ //追加map切片值
			key: value,
		})
	}
	for _, scliceItem := range mapslice {
		// for _, mapItem := range scliceItem {
		// 	mapItem["88"] = "88"
		// }
		scliceItem["88"] = "88"
	}
	Println("mapslice", mapslice)
}

// 命令行参数
var name = flag.String("name", "ss", "input your name")
var age = flag.Int("Age", 0, "input your age")

func getParams() {
	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*name, *age)
}

//结构体
func getStruct() {

	//结构体是值类型，赋值会拷贝
	type cat struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Color string `json:"color"`
	}
	// 四种声明方式
	var cat1 cat //第一种
	cat1.Name = "xiaobai"
	cat1.Age = 10
	cat1.Color = "white"

	// var cat2 = cat{"xiaobai",10,"dfd"}//第二种 cat2:=cat{"xiaobai",10,"dfd"}
	// var cat2 = cat{Name:"xiaobai",Age:10}//第二种 cat2:=cat{Name:"xiaobai",Age:10} 可以加上属性名称，不必全部赋值

	var cat3 *cat = new(cat) //new会返回一个指针，但可以直接赋值也可以这样写var cat4 = new(cat)
	cat3.Name = "xiaohei"    //直接赋值的写法
	(*cat3).Name = "xiaohei" //标准赋值的写法二者等价

	// var cat4 *cat = &cat{} //第四种
	jsonstr, err := json.Marshal(cat1)
	if err != nil {
		Println("ddd", err)
	} else {
		Println(string(jsonstr))
	}

	// 结构体tag 用于给结构体指定标签转小写,转json时使用
	// type cat struct {
	// 	Name  string `json:"name"`
	// 	Age   int    `json:"age"`
	// 	Color string `json:"color"`
	// }
}

// 方法
// 方法是作用在指定数据类型上，和指定的数据类型绑定，自定义数据类型都可以有方法，不仅仅是struct
// struct相当于类的的属性，这里的方法可以当作类的方法
type Preson struct {
	Name string
}
type Student struct {
	Name string
	Age  int
}

func (preson Preson) presonMethods() string { //声明方法，Preson相当于形参，将新的Preson实例传进来
	// println("名字是", preson.Name)
	preson.Name = "lisi"
	return preson.Name
}

func (preson *Preson) presonMethods2() string { //声明方法
	// println("名字是", (*preson).Name) //标准写法，用指针取值的方式
	// println("名字是", preson.Name) // 语法糖写法，可以省去指针取值
	preson.Name = "lisi"
	return preson.Name
}

func (stu *Student) String() string {
	str := fmt.Sprintf("name=[%v] age=[%v]", stu.Name, stu.Age)
	return str
}

func testPresonMethods() {
	// var preson1 = new(Preson)
	preson1 := Preson{
		Name: "xiaoming",
		// Age:  10,
	}
	// preson1.Name = "zhangsan"
	// preson1.presonMethods()
	// Println("传值", preson1)
	preson1.presonMethods2()
	// Println("传指针", preson1)

	// student1 := Student{
	// 	Name: "xiaoming",
	// 	Age:  10,
	// }
	// var student2 = new(Student)
	// student2.Name = "zhangsan"
	// student2.Age = 1
	// fmt.Println(student2)
}

// 面向对象例子
type Visitor struct {
	Name string
	Age  int
}

func (visitor Visitor) DecidePrice() {
	if visitor.Age > 15 {
		Printf("游客的名字为 %v 游客的年龄为%v，门票价格为20元\n", visitor.Name, visitor.Age)
	} else {
		Printf("游客的名字为 %v 游客的年龄为%v，门票价格为10元\n", visitor.Name, visitor.Age)
	}
}
func vis() {
	var vis Visitor
	for {
		Println("请输入游客名字")
		Scanln(&vis.Name)
		if vis.Name == "n" {
			Println("退出程序。。。")
			break
		} else {
			Println("请输入游客年龄")
			Scanln(&vis.Age)
			vis.DecidePrice()
		}
	}
}

//工厂模式
//外部访问包内部的私有结构体,如以下结构体被main访问时，需要使用工厂模式返回内部变量
type car struct {
	Name string
	Age  int
}

func (c *car) GetName() string {
	Println(c.Name)
	return c.Name
}

func NewCar(name string, age int) *car {
	return &car{
		Name: name,
		Age:  age,
	}
}

//返回指针返回值前加*
func returnPoint() *int {
	var num = 1
	return &num
}
