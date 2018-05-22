package linknodes

import (
	"fmt"
)

//声明全局变量，保存头结点
var head *Node
var curr *Node

//声明节点类型
type Node struct {
	//数据域
	Data DM
	//地址域
	NextNode *Node
	//头结点的当前节点
	curr *Node

}
//用结构体做数据域的类型
type DM struct {
	K string
	V string
}

//创建头结点
func CreateHeadNode(k,v string) (*Node) {
	var node *Node = new(Node)
	//设置数据域结构体中的键值对
	node.Data.K = k
	node.Data.V = v
	//设置地址域
	node.NextNode = nil
	//保存头结点
	head = node
	head.curr = node
	return node
}

//在指定节点后添加新节点
func AddNode(k,v string,currHead *Node) *Node {
	fmt.Println("addNode",currHead)
	var newNode *Node = new(Node)
	newNode.Data.K = k
	newNode.Data.V = v
	//设置地址域
	newNode.NextNode = nil
	//挂接节点
	currHead.curr.NextNode = newNode
	fmt.Printf("newNode:%p\n",newNode)
	//
	currHead.curr = newNode
	return currHead
}

//在指定的节点后面遍历链表
func ShowNodes(n *Node) {
	var i *Node = n
	for ; i != nil; i = i.NextNode {
		fmt.Println(i.Data)
	}

	//var tempnode *Node = head
	//for  {
	//	var node = tempnode
	//	fmt.Println(node.Data)
	//	tempnode = node.NextNode
	//	if tempnode == nil {
	//		break
	//	}
	//}
}
func NodeCnt() int {
	fmt.Println(head)
	var i = head
	count := 0
	for ; i != nil; i = i.NextNode {
		count++
	}
	fmt.Println(count)
	return count
}
/*
func InsertNodeByIndex(index int, data string) *Node {
	if index == 0 {
		//添加的为新的头结点
		var node *Node = new(Node)
		node.Data = data
		node.NextNode = head
		head = node
	} else if index > NodeCnt()-1 {
		//添加节点
		AddNode(data)
	} else {
		//中间插入节点
		var n = head
		for i := 0;i<index-1;i++ {
			n = n.NextNode
		}
		var newNode *Node = new(Node)
		newNode.Data = data
		newNode.NextNode = n.NextNode
		n.NextNode = newNode
	}
	return nil
}
func DeleteNode(index int)  {

	//var tempNode = head


	if index == 0 {
		//添加的为新的头结点
		var node *Node = new(Node)
		node = head.NextNode
		head = node
	} else {
		//中间插入节点
		var n = head
		if index > NodeCnt()-1 {
			index = NodeCnt()-1
		}
		for i := 0;i<index-1;i++ {
			n = n.NextNode
		}
		n.NextNode = n.NextNode.NextNode
	}
}
func UpdateNodeByIndex(index int, data string)  {
	var node  = head
	if index == 0 {
		head.Data = data
	} else {
		if index > NodeCnt()-1 {
			index = NodeCnt()-1
		}
		for i := 0; i < index; i++ {
			node = node.NextNode
		}
		node.Data = data
	}

}
*/