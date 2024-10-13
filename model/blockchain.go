package model

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, new(Block).Hash())
	bc.transactionPool = make([]*Transaction, 0)
	return bc
}

func (bc *Blockchain) CreateBlock(nounce int, previousHash [32]byte) *Block {
	block := NewBlock(nounce, previousHash, bc.transactionPool)
	bc.transactionPool = make([]*Transaction, 0)
	bc.chain = append(bc.chain, block)
	return block
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	bc.transactionPool = append(bc.transactionPool, NewTransaction(sender, recipient, value))
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
		fmt.Println()
	}
	fmt.Println(strings.Repeat("*", 25))
}
