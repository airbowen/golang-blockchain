package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
func main() {
	chain := InitBlockChain()
	chain.AddBlock("First BLOCK after genesis")
	chain.AddBlock("Second Block after genesis")
	chain.AddBlock("Third Block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("data in block %s\n", block.Data)
		fmt.Print("hash: %x\n", block.Hash)

	}
}
