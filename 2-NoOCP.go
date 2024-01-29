package main

import "fmt"

// 开闭原则: 类的改动是通过增加代码进行的, 而不是修改源代码

type Banker struct {
}

// 存款业务
func (b *Banker) Save() {
	fmt.Println("...进行了存款业务...")
}

// 转账业务
func (b *Banker) Transfer() {
	fmt.Println("...进行了转账业务...")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("...进行了支付业务...")
}

// 股票功能
// 增加该功能导致原有类结构发生改变，该写法不好
func (b *Banker) Share() {
	fmt.Println("...进行了股票业务...")
}

func main() {
	b := &Banker{}
	b.Pay()

	b.Transfer()

	b.Save()
	b.Share()
}
