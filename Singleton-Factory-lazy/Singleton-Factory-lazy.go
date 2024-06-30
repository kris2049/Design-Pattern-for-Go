package main

// type singleton struct{}

// var s *singleton

// func newSingleton() *singleton{
// 	return &singleton{}
// }

// func (s *singleton)Work(){}

// type Instance interface{
// 	Work()
// }

// func GetInstance() Instance{
// 	if s == nil{
// 		s=newSingleton()
// 	}
// 	return s
// }

/**
	当有多个goroutine，同时执行并访问全局变量s时，会出现资源竞争的情况。
	例如同时调用两次GetInstance，就会对s初始化两次，为了避免这种情况，要加锁
**/

import (
	"sync"
)

// type singleton struct{}

// var (
// 	s   *singleton
// 	mux sync.Mutex
// )

// func newSingleton() *singleton {
// 	return &singleton{}
// }

// func (s *singleton) Work() {}

// type Instance interface {
// 	Work()
// }

// 这样写存在一些浪费，因为如果s没有竞争情况，加锁就会无意义的损耗性能
// func GetInstance() Instance {
// 	mux.Lock()
// 	defer mux.Unlock()
// 	if s == nil {
// 		s = newSingleton()
// 	}

// 	return s
// }

// 这样写还是有问题。如果两个线程并发同时走到抢锁这一步，虽然由于锁的存在使代码串行，但是s仍然
// 被初始化了两次
// func GetInstance() Instance{
// 	if s!=nil{
// 		return s
// 	}
// 	mux.Lock()
// 	defer mux.Unlock()

// 	s = newSingleton()

// 	return s
// }

// 进行二次检查，由于锁的存在，即使两个线程同时抢到锁，也只有一个线程对s进行初始化，当初始化线程完成后释放锁
// 另一个线程要先对变量s进行二次判断，如果初始化了就返回。避免了二次初始化问题
// 其实sync库中现有的sync.Once来实现
// func GetInstance() Instance{
// 	if s!=nil{
// 		return s
// 	}
// 	mux.Lock()
// 	defer mux.Unlock()

// 	if s!=nil{
// 		return s
// 	}
// 	s = newSingleton()
// 	return s
// }

// sync.Once 可以保证全局变量或者高级别的结构体只被处理一次，它的作用周期就是变量或者结构体的生命周期
type singleton struct{}

var (
	once sync.Once
	s    *singleton
)

func newSingleton() *singleton {
	return &singleton{}
}

func (s *singleton) Work() {}

type Instance interface {
	Work()
}

func GetInstance() Instance {
	once.Do(func() {
		s = newSingleton()
	})
	return s
}
