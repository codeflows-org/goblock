package api

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int    // Index of the block.
	Timestamp    string // The moment the block has been added to the chain.
	Data         string
	Hash         string
	PreviousHash string
}

type Blockchain struct {
	blocks []Block
}

func NewBlockchain() Blockchain {
	genesis := Block{
		Index:        0,
		Timestamp:    "",
		Data:         "",
		Hash:         "",
		PreviousHash: "0",
	}
	return Blockchain{
		blocks: []Block{genesis},
	}
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func createBlock(previousBLock Block, data string) Block {
	var newBlock Block

	newBlock.Index = previousBLock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PreviousHash = previousBLock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func validateBlock(newBlock Block, oldBlock Block) bool {
	return oldBlock.Index+1 != newBlock.Index ||
		oldBlock.Hash != newBlock.PreviousHash ||
		calculateHash(newBlock) != newBlock.Hash
}
