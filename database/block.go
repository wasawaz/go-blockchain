package database

import (
	"crypto/sha256"
	"encoding/json"
)

type Hash [32]byte

type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []Tx        `json:"payload"`
}

type BlockFs struct {
	Key   Hash  `json:"hash"`
	Value Block `json:"block"`
}

type BlockHeader struct {
	Parent Hash   `json:"parent"`
	Time   uint64 `json:time`
}

func (b *Block) Hash() (Hash, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}
	return sha256.Sum256(blockJson), nil
}

func NewBlock(hash [32]byte, time uint64, txs []Tx) Block {
	return Block{BlockHeader{hash, time}, txs}
}
