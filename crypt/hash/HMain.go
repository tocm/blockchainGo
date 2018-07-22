package hash

import (
	"fmt"
	"strconv"
)

func HashMain()  {

	fmt.Println("---- hash function test BEGIN -----")
	//Test Simple hash
	fmt.Println(strconv.Itoa(SimpleHash(18888888)))
	fmt.Println(strconv.Itoa(SimpleHash(12)))
	fmt.Println(strconv.Itoa(SimpleHash(555555)))

	//典型Hash散列算法
	fmt.Println(HashCode("139999"))
	fmt.Println(HashCode("1558888"))


	//创建头节点
	headNode := CreateHeadNode(KValue{"a", "node"})
	//添加新节点
	newNode := AddNewNode(KValue{"b", "15"},headNode)
	newNode = AddNewNode(KValue{"c", "25"},newNode)
	newNode = AddNewNode(KValue{"d", "35"},newNode)

	//遍历所有节点
	ShowAllNode(headNode);

	//查找节点
	foundNode := FindNode(KValue{"b","15"}, headNode)
	if foundNode != nil {
		fmt.Println("the found node is ",foundNode.Data )

	}

	tailNode := TailNode(headNode)
	fmt.Println("the tail node is ", tailNode.Data)


	fmt.Println("--------- Test the Hash node Begin ------------")
	InitBuckets();
	AddKeyValue("andy","16888")
	AddKeyValue("tocm","55555")
	AddKeyValue("yy","165555")

	value := GetKeyValue("tocm")
	fmt.Println("get key value ",value)
	fmt.Println("--------- Test the Hash node end ------------")
}