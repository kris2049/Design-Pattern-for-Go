package main

import "fmt"

// 对类的改变应该通过增加代码实现，不应该修改原来的类。解决方法是使用抽象接口

// 创建一个抽象的banker，这个抽象的banker会做业务，具体什么业务根据实现这个接口的类决定
type AbstractBanker interface {
	Dobusiness()
}

// 存款业务员
type SaveBanker struct{}

func (sb *SaveBanker) Dobusiness() {
	fmt.Println("进行存款业务。。。。")
}

// 转账业务员
type TransferBanker struct{}

func (tb *TransferBanker) Dobusiness() {
	fmt.Println("进行转账业务。。。。")
}

// 股票业务员
type SharesBanker struct{}

func (sb *SharesBanker) Dobusiness() {
	fmt.Println("进行股票业务。。。。")
}
