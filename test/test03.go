package main

import (
	"sync"
	"time"
)

/*
//例子1，多线程同时操作一个变量，导致值没有达到预期的4万
func main() {
	var x = 0
	for num := 0; num < 4; num++ {
		go func() {
			for i := 0; i < 10000; i++ {
				increase(&x)
				//x = x + 1
			}
		}()
	}
	// 等待子线程退出
	time.Sleep(time.Millisecond * 2000)
	fmt.Println(x)
}

// 这个操作是Java做不了的.只有C,C++和Golang可以
func increase(x *int) {
	*x = *x + 1
}*/






var ILock=sync.Mutex{}

//例子2，使用sync.Mutex实现锁执行有序
func main() {
	var x = 0
	for num := 0; num < 4; num++ {
		go func() {

			ILock.Lock()
			defer ILock.Unlock()
			for i := 0; i < 10000; i++ {
				increase(&x)
				//x = x + 1
			}
		}()
	}
	// 等待子线程退出
	time.Sleep(time.Millisecond * 2000)
	fmt.Println(x)
}

// 这个操作是Java做不了的.只有C,C++和Golang可以
func increase(x *int) {
	*x = *x + 1
}


