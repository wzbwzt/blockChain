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
	nextNode := chain.Add(&newBlock, firstNode)
	fmt.Println(nextNode)
	chain.ShowNodeList(firstNode)
}
