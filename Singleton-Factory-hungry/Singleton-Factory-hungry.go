package main

/**
	Hungry mode: Initialize at the beginning of the program
	The structure of singleton and construction method must be not first letter capital
**/

// Hungry mode
// 's' should not be capital
// 创建一个单例结构体
type singleton struct{}

// 创建一个单例结构体变量
var s *singleton

// 创建一个创建单例结构体变量的函数
func newSingleton() *singleton {
	return &singleton{}
}

// 单例结构体的函数
func (s *singleton) Work() {}

// 返回单例结构体的实例
// func GetInstance() *singleton{
// 	return s
// }

/**
	上面的代码有一个规范性问题，GetInstance函数返回的是不可导出的类型singleton
	为了避免上面的问题，再定义一层interface，将这个接口作为GetInstance的返回值类型
**/
type Instance interface {
	Work()
}

// 返回值是一个接口，而结构体变量s实现了这个接口(实现了接口中的Work()函数)，所以返回变量s
func GetInstance() Instance {
	return s
}

func main() {

}
