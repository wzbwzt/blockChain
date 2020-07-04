package chain

import (
	"blockChain/block"
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

//判断节点中的区块是否合法
func CheckBlock(data *block.Block,preNode *Node)bool{
	if data.PreHash!=preNode.Data.HashCode {
		return false
	}
	if data.Index!=preNode.Data.Index+1 {
		return false
	}
	return true
}


//添加节点，当挖矿成功的话添加新节点
func Add(data *block.Block,preNode *Node) *Node{
	if CheckBlock(data,preNode){
		nextNode:=new(Node)
		nextNode.NextNode=nil
		nextNode.Data=data
		preNode.NextNode=nextNode
		return nextNode
	}
	return nil
}

//查看链表中的数据
func ShowNodeList(node *Node)[]block.Block{
	tmp:=node
	list:=make([]block.Block,0)
	for {
		if tmp.NextNode==nil {
			list=append(list,*tmp.Data)
			break
		}
		list=append(list,*tmp.Data)
		tmp=tmp.NextNode
	}
	return list
}



