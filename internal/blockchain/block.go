package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

// CalculateHash returns the SHA 256 hash of the block contents.
func CalculateHash(block Block) string {
	txData := ""
	for _, tx := range block.Transactions {
		txData += fmt.Sprintf("%s->%s:%.2f|", tx.Sender, tx.Receiver, tx.Amount)
	}
	record := block.Timestamp.String() + txData + block.PrevHash + strconv.Itoa(block.Nonce)

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}

// Serialize the block for storage/transmission
func (b *Block) Serialize() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(b)

	return buf.Bytes()
}

// DeserializeBlock converts bytes back into a Block
func DeserializeBlock(data []byte) *Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	_ = dec.Decode(&block)

	return &block
}
