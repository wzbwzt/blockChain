package main

import (
	"blockChain/block"
	"blockChain/chain"
	"fmt"
)

func main(){
	firstBlock := block.CreateFirstBlock("创世区块")
	newBlock := block.CreatNewBlock(firstBlock, "区块2")

	firstNode := chain.CreateFirstNode(&firstBlock)
	chain.Add(&newBlock, firstNode)
	list := chain.ShowNodeList(firstNode)
	fmt.Println(list)
}
