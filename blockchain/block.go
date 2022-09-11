package blockchain

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

func PrintBlock(block Block) {
	fmt.Printf("Block:\nData-%s\nHash-%x\nPrevHash-%x\n", block.Data, block.Hash, block.PrevHash)
}

// func DeriveHash updates the Hash field
// of the block with the hash calculated using
// the current blocks data and previpus blocks hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// lets use SHA256 as the placeholder for now
	// and we can move on to a better suitable Hash later
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// func CreateBlock take the new block data and previous blocks hash
// and returns a new block created with the supplied data
func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()

	return &block
}
