package main

import "fmt"

/*
继承有个问题，如果父类改变，很可能会导致子类改变
所以最好用合成，即两个模块合并一起
*/

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("eatting")
}

//给小猫增加一个睡觉的方法(使用继承来实现)
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("sleeping")
}

//使用组合的方式来实现睡觉
// type CatC struct {
// 	C *Cat
// }

// func (cc *CatC) Eat() {
// 	cc.C.Eat()
// }

// func (cc *CatC) Sleep() {
// 	fmt.Println("sleeping")
// }

// 在创建CatC类的时候，依赖于Cat类，这样耦合度还是高。下面解决这个问题。

type CatC struct{}

func (cc *CatC) Eat(c *Cat) {
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("sleeping")
}

func main() {
	c := &Cat{}
	c.Eat()

	fmt.Println("----------------")

	cb := &CatB{}
	cb.Eat()
	cb.Sleep()

	fmt.Println("----------------")

	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}
