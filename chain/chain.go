package chain

import (
	"blockChain/block"
	"fmt"
)

//链表主要结构是指针域和数据域
type Node struct {
	NextNode *Node
	Data *block.Block
}

//创建第一个节点
func CreateFirstNode(block *block.Block)*Node{
	firstNode:=new(Node)
	firstNode.NextNode=nil
	firstNode.Data=block
	return firstNode
}

//添加节点，当挖矿成功的话添加新节点
func Add(data *block.Block,preNode *Node) *Node{
	nextNode:=new(Node)
	nextNode.NextNode=nil
	nextNode.Data=data
	preNode.NextNode=nextNode
	return nextNode
}

//查看链表中的数据
func ShowNodeList(node *Node){
	tmp:=node
	for {
		if tmp.NextNode==nil {
			fmt.Println(tmp.Data)
			break
		}
		fmt.Println(tmp.Data)
		tmp=tmp.NextNode
	}
}



