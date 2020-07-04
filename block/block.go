package block

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const Diff =3

//区块
type  Block struct {
	PreHash string
	HashCode string
	TimeStamp int
	Diff int
	Data string
	Index int
	Nonce int
}
//生成哈希值
func  GenerateHashCode(block Block)string{
	hashData:=strconv.Itoa(block.Nonce)+strconv.Itoa(block.Diff)+strconv.Itoa(block.Index)+block.PreHash
	h:=sha256.New()
	h.Write([]byte(hashData))
	restmp := h.Sum(nil)
	res := hex.EncodeToString(restmp)
	return res
}
//生成一个创世区块
func CreateFirstBlock(data string)(firstBlock Block){
	firstBlock.PreHash="0"
	firstBlock.TimeStamp=int(time.Now().Unix())
	firstBlock.Diff=Diff
	firstBlock.Data=data
	firstBlock.Index=1
	firstBlock.Nonce=123
	firstBlock.HashCode=GenerateHashCode(firstBlock)
	return
}

//共识算法pow
func Pow(block *Block)string{
	for  {
		hashCode := GenerateHashCode(*block)
		if strings.HasPrefix(hashCode,strings.Repeat("0",block.Diff)) {
			return hashCode
		}
		rand.Seed(time.Now().Unix())
		block.Nonce+=rand.Intn(100)
	}
}

//生成新区块
func CreatNewBlock(preBlock Block, data string)(newBlock Block){
	newBlock.PreHash=preBlock.HashCode
	newBlock.Data=data
	newBlock.Index=preBlock.Index+1
	newBlock.Diff=Diff
	newBlock.Nonce=123
	newBlock.TimeStamp=int(time.Now().Unix())
	hash := Pow(&newBlock)
	newBlock.HashCode=hash
	return
}


