package blockchain

import (
	"bytes"
	"encoding/binary"
)

// IntToHex converts an int64 to a byte slice
func IntToHex(n int64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, n)

	return buf.Bytes()
}
