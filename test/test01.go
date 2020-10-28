package main

import (
	"fmt"
)

//自定义接口
type Imoocer interface{
	Hello()
}
//自定义结构体
type Student struct {
	Id   int64
	Name string
}

type Teacher struct {
	Addr  string
	Group string
}

type ImoocPtr string


//这样就可以说Student实现了此方法
func (a *Student)Hello() {
	fmt.Printf("Student(%d, %s) Hello imooc！ \n", a.Id, a.Name)
}
// 实现Hello（）函数 并且绑Teacher 结构体中，作为Teacher的方法；
func (a *Teacher) Hello() {
	fmt.Printf("Teacher(%s, %s) Hello cap1573！\n", a.Addr, a.Group)
}
// 实现Hello（）函数 ImoocPtr 结构体中，作为ImoocPtr的方法；
func (a *ImoocPtr) Hello() {
	fmt.Printf("ImoocPtr(%s) \n", *a)
}


// 定义一个普通函数，函数的参数为接口类型
func SayHello(i Imoocer) {
	i.Hello()
}

// 创建main（）函数
func main() {

	// 只要实现了此接口方法的类型，那么这个类型的变量就可以给i赋值
	s := &Student{1, "小明"}
	//使用Student已经实现的方法
	SayHello(s)

	//创建并且初始化Teacher实例
	t := &Teacher{"南京", "IT"}
	//使用Teacher已经实现的方法
	SayHello(t)

	var str ImoocPtr = "同学们好"
	//使用ImoocPtr已经实现的方法
	SayHello(&str)

	fmt.Println()
}