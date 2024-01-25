package main

import "fmt"

// 单一职责原则: 类的职责单一，对外只提供一种功能，而引起类的变化的原因都应该只有一个

type ClothesShop struct {
}

type ClothesWork struct {
}

func (cs *ClothesShop) OnShop() {
	fmt.Println("在逛街时的装扮...")
}

func (cw *ClothesWork) OnWork() {
	fmt.Println("在工作时的装扮...")
}

func main() {
	cs := new(ClothesShop)
	fmt.Println("在逛街时....")
	cs.OnShop()

	cw := new(ClothesWork)
	fmt.Println("在工作时....")
	cw.OnWork()

}

/*
// 错误的例子
type Clothes struct {
}

func (c *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}

// OnShop() 方法的修改可能会影响到类里面其余方法的使用或者修改
func (c *Clothes) OnShop() {
	fmt.Println("工作的装扮")
}

func main() {
	c := Clothes{}

	fmt.Println("工作时...")
	c.OnWork()

	fmt.Println("逛街时...")
	// c.OnWork()
	c.OnShop()
}
*/
