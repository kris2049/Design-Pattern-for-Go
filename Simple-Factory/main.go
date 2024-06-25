package main

import (
	"fmt"
)

/**
	对要构造的类根据其共性定义一个公共的interface
	比如无论是什么水果，都要被种植和采摘
**/
//水果接口
type Fruit interface {
	// 种植方法
	grant()
	// 采摘方法
	pick()
}

/**
	使用实例去实现interface，即创建苹果和桔子实例实现Fruit接口
**/
// 苹果结构体实现水果接口  只要实现了水果接口中的方法就可以说这个结构体实现了这个接口
type Apple struct {
	name string
}

// grant implements Fruit.
func (*Apple) grant() {
	fmt.Println("Plant an Apple!")
}

// pick implements Fruit.
func (*Apple) pick() {
	fmt.Println("Pick an Apple!")
}

// 桔子结构体实现水果接口
type Orange struct {
	name string
}

func (*Orange) grant() {
	fmt.Println("Plant an Orange!")
}

func (*Orange) pick() {
	fmt.Println("Pick an Orange!")
}

// 为什么Fruit不是指针而要使用取地址符：
// 1. 共享性和可修改性， 返回值类型是一个副本，对原始实例没有改变
// 2. 指针比返回结构体更有效率
// 3. 指针可以是nil，但是值类型必须要有一个非零的默认值
func NewApple(name string) Fruit {
	return &Apple{
		name: name,
	}
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

/**
	创建水果工厂类，用来创建水果
**/
// 首先定义一个结构体表示水果工厂类
// type FruitFactory struct{}

// 定义一个函数来创建水果工厂类
// func NewFruitFactory() *FruitFactory {
// 	return &FruitFactory{}
// }

// 定义水果工厂类的创建水果函数
// func (*FruitFactory) CreateFruit(name string) Fruit {
// 	switch name {
// 	case "apple":
// 		return NewApple(name)
// 	case "orange":
// 		return NewOrange(name)
// 	}
// 	return nil
// }
/*************************************************************************************************************************************************/
/**
	使用switch语句有个问题，如果水果种类很多很复杂的话，关系会变得很复杂。
	所以可以将水果构造函数定义成一个方法
**/
// fruitCrerator 的类型是一个函数，返回值是Fruit
type fruitCreator func(name string) Fruit

// 在水果工厂类中增加一个map，将水果类型映射到构造函数
type FruitFactory struct {
	creators map[string]fruitCreator
}

// 在定义一个函数来创建水果工厂类的时候要初始化类中的map
func NewFruitFactory() *FruitFactory {
	return &FruitFactory{
		creators: map[string]fruitCreator{
			"apple":  NewApple,
			"orange": NewOrange,
		},
	}
}

// 定义一个水果工厂类的创建水果函数
func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	fruitCreator, ok := f.creators[typ]
	if !ok {
		return nil, fmt.Errorf("fruit type: %s is not support yet", typ)
	}

	return fruitCreator(typ), nil
}

func main() {

	//首先创建水果工厂类
	fruitFactory := NewFruitFactory()

	//使用水果工厂类的函数来创建水果，这些水果已经实现了抽象接口中的方法，
	//所以可以使用接口中的函数
	apple, _ := fruitFactory.CreateFruit("apple")
	apple.grant()
	apple.pick()

	orange, _ := fruitFactory.CreateFruit("orange")
	orange.grant()
	orange.pick()

}
