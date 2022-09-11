package main

import (
	"fmt"

	"github.com/KavetiRohith/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("I am Awesome")
	chain.AddBlock("So i am trying to build a toy Blockchain :)")
	blockchain.PrintBlock(*chain.Blocks[0])
	blockchain.PrintBlock(*chain.Blocks[1])
	blockchain.PrintBlock(*chain.Blocks[2])

	fmt.Println("-------------------------------------------------------------")

	chainTwo := blockchain.InitBlockChain()
	chainTwo.AddBlock("I am awesome") // notice the small a in awesome here
	chainTwo.AddBlock("So i am trying to build a toy Blockchain :)")
	blockchain.PrintBlock(*chainTwo.Blocks[0])
	blockchain.PrintBlock(*chainTwo.Blocks[1])
	blockchain.PrintBlock(*chainTwo.Blocks[2])
}
