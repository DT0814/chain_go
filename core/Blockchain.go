package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockChain() *Blockchain {
	genesisBlock := generategenersisBlock()
	blockChain := Blockchain{}
	blockChain.ApendBlock(&genesisBlock)
	return &blockChain
}
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := generateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}
func (bc *Blockchain) ApendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Print("invalid block")
	}

}
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index:%d \n", block.Index)
		fmt.Printf("Prevhash:%s \n", block.PrevBlockHash)
		fmt.Printf("CurrHash:%s \n", block.Hash)
		fmt.Printf("data:%s \n", block.Data)
		fmt.Printf("temp:%d \n", block.TimeStamp)
	}
}
