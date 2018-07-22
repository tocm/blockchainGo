package oop

import "fmt"

//测试接口: 实现结构类必须要全部实现接口内方法
type Phone interface{
	Call()
	GetUUID() int64
}

type Iphone struct{

}
type AndrodPhone struct{

}

func(iphone Iphone) Call() {
	fmt.Println("this is call function");
}
func(iphone Iphone) GetUUID() int64 {
	return 555
}

func(androidPhone AndrodPhone) Call() {
	fmt.Println("this is android phone");
}

func(androidPhone AndrodPhone) GetUUID() int64 {
	return 8888
}
