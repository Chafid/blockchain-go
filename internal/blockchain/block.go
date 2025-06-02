package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

// CalculateHash returns the SHA 256 hash of the block contents.
func (b *Block) CalculateHash() []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(int64(b.Index)),
			IntToHex(b.Timestamp),
			[]byte(b.Data),
			b.PrevHash,
			IntToHex(int64(b.Nonce)),
		}, []byte{},
	)
	hash := sha256.Sum256(data)

	return hash[:]
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
