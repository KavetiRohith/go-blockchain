package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// struct Block represnts a single block
// Data holds the data of the block
// PrevHash is the Hash of previous block
// the current blocks hash is calculated with the
// current blocks data and and previous blocks hash
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func printBlock(block Block) {
	fmt.Printf("Block:\nData-%s\nHash-%x\nPrevHash-%x\n", block.Data, block.Hash, block.PrevHash)
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// lets use SHA256 as the placeholder for now
	// and we can move on to a better suitable Hash later
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()

	return &block
}

type BlockChain struct {
	blocks []*Block
}

// func Genesis Creates the first block of our blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	prevHash := prevBlock.Hash
	chain.blocks = append(chain.blocks, CreateBlock(data, prevHash))
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("I am Awesome")
	chain.AddBlock("So i am trying to build a toy Blockchain :)")
	printBlock(*chain.blocks[0])
	printBlock(*chain.blocks[1])
	printBlock(*chain.blocks[2])
	fmt.Println("-------------------------------------------------------------")
	chainTwo := InitBlockChain()
	chainTwo.AddBlock("I am awesome") // notice the small a in awesome here
	chainTwo.AddBlock("So i am trying to build a toy Blockchain :)")
	printBlock(*chainTwo.blocks[0])
	printBlock(*chainTwo.blocks[1])
	printBlock(*chainTwo.blocks[2])
}
