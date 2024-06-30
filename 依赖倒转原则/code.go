package main

import "fmt"

/*
	实现层向上依赖抽象层，业务逻辑层向下依赖抽象层
*/

//---------抽象层----------//

// 汽车抽象类
type Car interface {
	Run()
}

// 司机抽象类
type Driver interface {
	Drive(car Car)
}

//---------实现层----------//
// 奔驰车类，对汽车抽象类进行实现
type Benz struct {
}

func (bz *Benz) Run() {
	fmt.Println("奔驰running")
}

// 宝马车类，对汽车抽象类进行实现
type Bmw struct {
}

func (bmw *Bmw) Run() {
	fmt.Println("宝马running")
}

// 张三和李四类，对司机抽象类进行实现
type zhang3 struct {
}

func (z3 *zhang3) Drive(car Car) {
	fmt.Println("张三drive car")
	car.Run()
}

type li4 struct {
}

func (l4 *li4) Drive(car Car) {
	fmt.Println("李四drive car")
	car.Run()
}

//---------业务逻辑层----------//
func main() {
	//张三开奔驰
	var benz Car
	var zhangsan Driver

	benz = new(Benz)
	zhangsan = new(zhang3)
	zhangsan.Drive(benz)

	//李四开宝马
	var bmw Car
	var lisi Driver

	bmw = new(Bmw)
	lisi = new(li4)
	lisi.Drive(bmw)
}
