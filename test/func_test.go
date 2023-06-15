package test

import (
	"fmt"
	"testing"
)

/**
函数作为一等公民的一些特殊用法
*/

/**
func 类型强转:
Afunc 为 func(x, y int) int 类型, 同时实现了 AInterface 接口的 Add 方法, 所以 Afunc 可以强转为 AInterface 类型.
someAction 本身的定义和 AInterface 接口的 Add 方法一致, 所以 someAction 也可以强转为 Afunc 类型, 从而实现了 AInterface 接口.
*/

type AInterface interface {
	Add(x, y int) int
}

type Afunc func(x, y int) int

func (a Afunc) Add(x, y int) int {
	return a(x, y)
}

func doSomething(da AInterface) {
	fmt.Println(da.Add(4, 5))
}

func someAction(x, y int) int {
	return x + y
}

func TestInterface(t *testing.T) {
	// 强转为 Afunc 类型
	a := Afunc(someAction)
	doSomething(a)
}

// 函数式编程之函子（functor）
// 函子本身是一个容器类型，该容器类型需要实现一个方法，该方法接受一个函数类型参数，并在容器的每个元素上应用那个函数，得到一个新函子，原函子容器内部的元素值不受影响

type IntSliceFunctor interface {
	CalcuateMap(f func(x int) int) IntSliceFunctor
}

func NewIntSliceFunctor(s []int) IntSliceFunctor {
	return MySlice{s: s}
}

type MySlice struct {
	s []int
}

func (ms MySlice) CalcuateMap(f func(x int) int) IntSliceFunctor {
	var result []int
	for _, v := range ms.s {
		result = append(result, f(v))
	}
	return MySlice{s: result}
}

func TestMySlice(t *testing.T) {
	originSlice := []int{1, 2, 3}
	fmt.Println(originSlice)

	m := NewIntSliceFunctor(originSlice)

	mf := func(i int) int {
		return i * 10
	}
	ms := m.CalcuateMap(mf)

	fmt.Println(ms.CalcuateMap(mf))
	fmt.Println(originSlice)
}

// 函数式编程之CPS (Continuation Passing Style) 传递函数
// 传递函数的一种方式，将函数作为参数传递给另一个函数，这种方式称为 CPS，CPS 是一种函数编程风格，它的核心思想是将计算的控制流进行显示的传递，这样做的好处是可以控制计算的执行顺序，从而实现一些特殊的功能，比如异常处理、回溯等。
// CPS 风格中，函数的最后一个参数是一个函数，这个函数就是 continuation，它是计算的控制流，它的作用是接收计算的结果，然后进行处理，比如打印、异常处理等。且函数不再返回结果，而是将结果传递给 continuation 函数。
// 但 CPS 风格的代码可读性很差，所以在实际开发中，我们很少使用这种风格，但是它的思想是很重要的，它可以帮助我们理解一些特殊的函数，比如 map、filter、reduce 等。

// 以下是两个常规函数，一个是比较大小，一个是阶乘
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 将上面连个函数改写成 CPS 风格
func maxCPS(x, y int, cont func(int)) {
	if x > y {
		cont(x)
	} else {
		cont(y)
	}
}

func factorialCPS(n int, cont func(int)) {
	if n == 0 {
		cont(1)
	} else {
		factorialCPS(n-1, func(i int) {
			cont(n * i)
		})
	}
}

// CPS 风格的函数调用
func TestCPS(t *testing.T) {
	maxCPS(10, 20, func(i int) {
		fmt.Println(i)
	})

	factorialCPS(5, func(i int) {
		fmt.Println(i)
	})
}
