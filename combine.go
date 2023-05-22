package main

import "fmt"

/**
Go 中没有继承的概念，但是可以通过匿名字段来实现继承的效果。
把一个 struct 嵌入到另一个 struct 中，被嵌入的 struct 称为匿名字段。
匿名字段的数据类型必须是 struct 或者是指向 struct 的指针。
匿名字段默认采用类型名作为字段名，如果想要用其他的名字作为字段名可以采用下面的语法：
type Human struct {
	name string
	age int
}
type Student struct {
	Human // 匿名字段，那么默认Student就包含了Human的所有字段
	school string
}
如果一个 struct 嵌入了多个匿名字段，那么该 struct 可以直接访问被嵌入的所有字段。
如果多个匿名字段拥有相同的字段名，那么在访问该字段时就必须明确指定匿名字段的名字。

匿名字段的可见性规则遵循简单的基于名字的规则：
	1.首字母大写的字段名是可导出的；
	2.首字母小写的字段名是不可导出的。
	3.结构体中嵌入的匿名结构体，可以直接访问匿名结构体的方法，编译器会将方法与匿名结构体进行组合，如果结构体和匿名结构体有相同的方法，优先调用结构体中的方法。
	4.如果结构体中包含有多个匿名结构体，那么在访问匿名结构体中的方法时，就必须带上匿名结构体的名字。
	5.如果结构体中包含有多个字段名相同的匿名结构体，那么在访问这个字段时，就必须带上匿名结构体的名字。
*/

func doMain() {
	p := Person{"张三", 20}
	s := Student{p, "清华大学"}

	// 如果结构体和匿名结构体有相同的方法，优先调用结构体中的方法
	fmt.Println(s.String())
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type Student struct {
	Person
	School string
}

func (s *Student) String() string {
	return fmt.Sprintf("%v (%v years) %v", s.Name, s.Age, s.School)
}

// 常量定义风格推荐
var (
	a = 14
	b = int32(14)
)

// 不同类型的变量间的运算时不支持隐式的类型转换，如果要比较或者运算，必须有显示的转换类型.
// 两个类型即便拥有相同的底层类型（underlying type），也仍然是不同的数据类型
func compare() {
	type myInt int
	var a int = 5
	var b myInt = 6
	fmt.Println(a + int(b)) // 输出：11
}
