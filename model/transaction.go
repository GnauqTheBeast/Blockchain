package model

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (tx *Transaction) Print() {
	fmt.Printf("Sender: %s\n", tx.senderBlockchainAddress)
	fmt.Printf("Recipient: %s\n", tx.recipientBlockchainAddress)
	fmt.Printf("Value: %.2f\n", tx.value)
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    tx.recipientBlockchainAddress,
		RecipientBlockchainAddress: tx.senderBlockchainAddress,
		Value:                      tx.value,
	})
}
