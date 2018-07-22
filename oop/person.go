package oop
import (
	"strconv"
	"fmt"
)

type Person struct{
	name string
	age int
	address string
}

func (p Person) GetInfo() string {
	return "name = "+p.name +", age = "+strconv.Itoa(p.age) +",address = "+p.address
}

/**
* 必须要是传指针，否则不会保存当前set 的值，在GetInfo中无法取出
*/
func (p *Person) SetInfo(name string, age int, address string) {
	p.name = name
	p.age = age
	p.address = address
	fmt.Println(p.name + "," + strconv.Itoa(p.age) +", " + p.address)

}

