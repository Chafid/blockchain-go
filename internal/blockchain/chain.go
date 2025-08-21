package blockchain

import (
	"strings"
	"time"
)

type Blockchain struct {
	Blocks []Block
}

const difficulty = 2

// CreateGenesisBlock to initializes the blockchain
func CreateGenesisBlock() Block {
	genesis := Block{
		Timestamp:    time.Now(),
		Transactions: []Transaction{{Sender: "genesis", Receiver: "network", Amount: 0}},
		PrevHash:     "",
		Nonce:        0,
	}
	genesis.Hash = CalculateHash(genesis)
	return genesis
}

func NewBlockchain() *Blockchain {
	genesis := CreateGenesisBlock()
	return &Blockchain{Blocks: []Block{genesis}}
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := mineBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func mineBlock(transactions []Transaction, prevHash string) Block {
	var block Block
	block.Timestamp = time.Now()
	block.Transactions = transactions
	block.PrevHash = prevHash

	for {
		hash := CalculateHash(block)
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			block.Hash = hash
			break
		}
		block.Nonce++
	}

	return block
}

func (bc *Blockchain) GetBalance(address string) float64 {
	var balance float64
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.Sender == address {
				balance -= tx.Amount
			}
			if tx.Receiver == address {
				balance += tx.Amount
			}
		}
	}

	return balance
}
