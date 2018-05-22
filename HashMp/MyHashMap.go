package HashMp

import (
	"zhqGo/MyhashMap/linknodes"
	"fmt"
	"strconv"
)

var buletArr  [16]*linknodes.Node

//给数组初始化
func CreateBulet()  {
	var arr  [16]*linknodes.Node
	for i := 0; i < 16; i++ {
		arr[i] = linknodes.CreateHeadNode("头key"+strconv.Itoa(i),"头结点")
	}
	buletArr = arr
}
//向数组中添加键值对
func AddKeyValue(k,v string )  {
	//计算键所对应的木桶下标
	var pos = HashCode(k)
	//获得木桶数组
	//var arr = buletArr
	var head *linknodes.Node = buletArr[pos]
	//向指定下标的头结点，添加节点
	linknodes.AddNode(k,v,head)
}
//获取数据
func GetValueByKey(k string)string  {
	var pos = HashCode(k)
	var head * linknodes.Node = buletArr[pos]
	fmt.Println(buletArr)
	//通过头结点遍历链表
	//linknodes.ShowNodes(head)
	//查找对应下标下的链表,只有判断在key与节点中的key一致时打印
	var n = head
	for  {
		//如果 节点为nil说明已经到链表最后，也就是链表中不存在这个值
		if n == nil {
			fmt.Println("没有对应的值")
			break
		}
		if n.Data.K == k {
			fmt.Println(n.Data.V)
			break
		}else {
			n = n.NextNode
		}
	}
	return ""
}
//将key转换成数组下标的散列算法
func HashCode(key string) int {
	var index int = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		//散列因子
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1

	return index
}
