package blockchain

import (
	"bytes"
	"fmt"
)

const targetPrefix = "00" // Difficulty: hash must start with this

type ProofOfWork struct {
	block *Block
	//target []byte
}

func NewProofOfWork(b *Block) *ProofOfWork {
	return &ProofOfWork{block: b}
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hash []byte
	var nonce int
	for {
		pow.block.Nonce = nonce
		hash = pow.block.CalculateHash()
		if bytes.HasPrefix(hash, []byte(targetPrefix)) {
			break
		}
		nonce++
		if nonce%100000 == 0 {
			fmt.Println("Still mining... nonce:", nonce)
		}
	}
	fmt.Printf("Mined block: %x\n", hash)

	return nonce, hash
}
