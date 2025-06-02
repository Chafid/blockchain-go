package main

import (
	"fmt"

	"blockchain-go/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	// Add a few blocks
	bc.AddBlock("First block after Genesis")
	bc.AddBlock("Another block")

	// Print the chain
	for _, block := range bc.Blocks {
		fmt.Println("========= BLOCK ==========")
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
	}
}
