package blockchain

// struct BlockChain is a slice of pointers to various
// Blocks comprising the Blockchain
type BlockChain struct {
	Blocks []*Block
}

// func Genesis Creates the first block of our blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// func InitBlockChain initialises a new chain with a genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// func AddBlock takes data iin form of string creates a new block
// based on given data and updates the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	prevHash := prevBlock.Hash
	chain.Blocks = append(chain.Blocks, CreateBlock(data, prevHash))
}
