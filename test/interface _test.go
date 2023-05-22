package test

import (
	"fmt"
	"testing"
)

/**
func 类型强转例子:
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

// ———————————————————— 组合实现继承 ————————————————————

type HumanInterface interface {
	Say()
	Run()
}

type BaseHuman struct{}

func (h *BaseHuman) Say() {
	fmt.Println("BaseHuman Say")
}

type EarthHuman struct {
	BaseHuman
}

func (h *EarthHuman) Run() {
	fmt.Println("EarthHuman Run")
}

// func (h *EarthHuman) Say() {
// 	fmt.Println("EarthHuman Say")
// }

func execHuman(h HumanInterface) {
	h.Say()
	h.Run()
}

func TestInherit(t *testing.T) {
	h := &EarthHuman{}
	execHuman(h)
}
