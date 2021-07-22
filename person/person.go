package person

import "fmt"

func PerFun() {
	// 继承
	// zs := &Man{}
	// zs.Name = "张三"
	// zs.ShowInfo()
	// zs.Work()

	// ls := &Women{}
	// fmt.Println(ls)

	//嵌套结构体继承
	tv := TV{
		Goods{Name: "dd", Price: 5988},
		Brand{Name: "haier", Address: "qingdao"},
	}
	fmt.Println(tv)
}

//继承
type Person struct {
	Name string
	Age  int
	Add  string
}

func (person Person) ShowInfo() {
	fmt.Println(person.Name, person.Age, person.Add)
}

//如若出现同名属性和方法，就近寻找
type Man struct {
	Person
}

func (man Man) Work() {
	fmt.Println("努力搬砖")
}

type Women struct {
	Person // 匿名结构体
	// per    Person //有名结构体，访问其方法和属性时必须带上名称，a.per.Name
}

// 嵌套结构体
type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}
