package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	TimeStamp     int64  //区块时间戳 区块创建的时间
	PrevBlockHash string //前一区块哈希
	Hash          string //当前区块哈希
	Data          string //当前区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.TimeStamp) + string(b.PrevBlockHash)
	hashByte := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashByte[:])
}

func generateNewBlock(preBlockHash Block, data string) Block {
	newBlock := Block{}
	newBlock.PrevBlockHash = preBlockHash.Hash
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.Index = preBlockHash.Index + 1
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Data = data
	return newBlock
}
func generategenersisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return generateNewBlock(preBlock, "0")
}
