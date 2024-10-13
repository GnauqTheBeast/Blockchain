package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nounce       int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nounce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nounce:       nounce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\nnonce: %d\nprevious hash: %x\n", b.timestamp, b.nounce, b.previousHash)
	fmt.Printf(strings.Repeat("-", 50))
	fmt.Printf("\nTx:\n")
	for _, tx := range b.transactions {
		tx.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nounce       int            `json:"nounce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nounce:       b.nounce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}
