package main

import "fmt"

/*
	OCP: 开闭原则, 平时写代码一定要具备的思想, 相当重要
*/

type AbstractBanker interface {
	DoBus()
}

type SaveBanker struct {
	// AbstractBanker
}

func (sb *SaveBanker) DoBus() {
	fmt.Println("...进行了存款业务...")
}

type TransferBanker struct {
	// AbstractBanker
}

func (tb *TransferBanker) DoBus() {
	fmt.Println("...进行了转账业务...")
}

type PayBanker struct {
	// AbstractBanker
}

func (pb *PayBanker) DoBus() {
	fmt.Println("...进行了支付业务...")
}

// 如果要新加一个功能, 只需要在原有接口上实现一个类即可，不影响其余功能的实现
type ShareBanker struct {
	// AbstractBanker
}

func (*ShareBanker) DoBus() {
	fmt.Println("...进行了股票业务...")
}

// 实现一个架构层, 基于抽象层进行业务封装-针对interface接口进行封装
func BankBusiness(b AbstractBanker) {
	// 通过接口向下来调用(多态的实现)
	b.DoBus()
}

func main() {
	/*
		sb := &SaveBanker{}
		sb.DoBus()

		pb := &PayBanker{}
		pb.DoBus()

		tb := &TransferBanker{}
		tb.DoBus()
		sb1 := &ShareBanker{}
		sb1.DoBus()
	*/

	// 另一种实现方式, 更为简洁
	BankBusiness(&SaveBanker{})
	BankBusiness(&PayBanker{})
	BankBusiness(&ShareBanker{})
	BankBusiness(&TransferBanker{})
}
