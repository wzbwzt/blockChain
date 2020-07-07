package main

import (
	"blockChain/http/service"
)

func main(){
	//firstBlock := block.CreateFirstBlock("创世区块")
	//newBlock := block.CreatNewBlock(firstBlock, "区块2")
	//
	//firstNode := chain.CreateFirstNode(&firstBlock)
	//chain.Add(&newBlock, firstNode)
	//list := chain.ShowNodeList(firstNode)
	//indent, err := json.MarshalIndent(list, "", "\t")
	//if err != nil {
	//	log.Println("json.MarshalIndent",err)
	//	return
	//}
	service.Run()
}
