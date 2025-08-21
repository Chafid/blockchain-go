package main

import (
	"blockchain-go/internal/blockchain"
	"fmt"
)

func main() {
	// Initialize
	chain := blockchain.NewBlockchain()

	//Add blocks with some transaction
	chain.AddBlock([]blockchain.Transaction{
		{Sender: "alice", Receiver: "bob", Amount: 10},
		{Sender: "bob", Receiver: "carol", Amount: 5},
	})

	chain.AddBlock([]blockchain.Transaction{
		{Sender: "carol", Receiver: "dave", Amount: 2.5},
		{Sender: "alice", Receiver: "dave", Amount: 1},
	})

	// Print the full chain
	for i, block := range chain.Blocks {
		fmt.Printf("========= Block %d ==========\n", i)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		for _, tx := range block.Transactions {
			fmt.Printf(" %s -> %s: %.2f\n", tx.Sender, tx.Receiver, tx.Amount)
		}
		fmt.Println()
	}

	addresses := []string{"alice", "bob", "carol", "dave"}
	for _, addr := range addresses {
		balance := chain.GetBalance(addr)
		fmt.Printf("Balance of %s: %.2f\n", addr, balance)
	}
}
