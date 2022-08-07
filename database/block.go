package database

import (
	"crypto/sha256"
	"encoding/json"
)

type Block struct {
	Header BlockHeader
	TXs    []Tx
}

type BlockHeader struct {
	Parent Hash
	Time   uint64
}

func (b *Block) Hash() (Hash, error) {
	blockJson, _ := json.Marshal(b)
	return sha256.Sum256(blockJson), nil
}
