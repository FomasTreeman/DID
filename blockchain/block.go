package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  string
	Hash      string
	// Commenting out PoW related fields
	// Nonce      int
	// Difficulty int
}

func NewBlock(data []byte, prevHash string) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		// Removing Difficulty initialization
	}
	// Instead of mining, just calculate hash once
	block.Hash = block.CalculateHash()
	return block
}

func (b *Block) CalculateHash() string {
	// Simplified hash calculation without nonce
	record := string(b.Data) + b.PrevHash + string(b.Timestamp)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// Commenting out the Mine function as we're not using PoW for now
/*
func (b *Block) Mine() {
	target := string(make([]byte, b.Difficulty))
	for i := 0; i < b.Difficulty; i++ {
		target = "0" + target
	}

	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:b.Difficulty] == target {
			break
		}
		b.Nonce++
	}
}
*/
