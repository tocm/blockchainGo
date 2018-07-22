package hash

import "fmt"

//========	链表实现  =========//


 /**
 	存放数据元素结构
  */
type KValue struct {
	Key string
	Value string
}

/**
	节点结构
 */
type LNode struct {
	Data KValue
	NextNode *LNode;
}

/*
	创建头结点
 */
func CreateHeadNode(data KValue) *LNode {
	//fmt.Println("add new node ",data)
	var pHeadNode = new(LNode)
	pHeadNode.Data = data
	pHeadNode.NextNode = nil
	return pHeadNode
}

/**
	添加新节点
 */
func AddNewNode(data KValue, lnode *LNode) *LNode {
	var newNode = CreateHeadNode(data)
	lnode.NextNode = newNode
	return newNode
}

/**
	尾节点
 */
func TailNode(node *LNode) *LNode  {
	tailNode := node
	for tailNode != nil && tailNode.NextNode != nil {
		tailNode = tailNode.NextNode
	}
	return tailNode;
}

/*
	查找节点
 */
func FindNode(key KValue, node *LNode) *LNode  {
	if node != nil {
		for node.NextNode != nil {
			if node.Data == key || node.Data.Key == key.Key {
				return node;
			}
			node = node.NextNode
		}
	}
	return node
}
/**
	查询所有节点
 */
func ShowAllNode(headNode *LNode)  {
	note := headNode
	for note != nil {
		fmt.Println("show current link all node", note.Data)
		note = note.NextNode

	}
	
}