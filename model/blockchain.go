package model

import (
	"fmt"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
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

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, len(bc.transactionPool))
	copy(transactions, bc.transactionPool)
	return transactions
}

func (bc *Blockchain) ValidateProof(nounce int, previousHash [32]byte, transactions []*Transaction) bool {
	block := &Block{0, nounce, previousHash, transactions}
	return strings.HasPrefix(fmt.Sprintf("%x", block.Hash()), strings.Repeat("0", MINING_DIFFICULTY))
}

func (bc *Blockchain) ProofOfWork() int {
	nounce := 0
	previousHash := bc.LastBlock().Hash()
	transactions := bc.CopyTransactionPool()

	for !bc.ValidateProof(nounce, previousHash, transactions) {
		nounce++
	}

	return nounce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	fmt.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) CalcuateTotalAmount(blockchainAddress string) float32 {
	totalAmount := float32(0.0)
	for _, block := range bc.chain {
		for _, transaction := range block.transactions {
			if transaction.sender == blockchainAddress {
				totalAmount -= transaction.value
			}
			if transaction.recipient == blockchainAddress {
				totalAmount += transaction.value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
		fmt.Println()
	}
	fmt.Println(strings.Repeat("*", 25))
}
