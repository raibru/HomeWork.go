package ledger

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Block holds the data for a single block
type Block struct {
	Hash     []byte
	Data     []byte
	HashPrev []byte
	Nonce    int
	Elapsed  string
}

// CreateBlock create a new Block with deduced hash value
func CreateBlock(data []byte, prevHash []byte) *Block {
	b := &Block{[]byte{}, data, prevHash, 0, "0"}
	pow := CreateProof(b)

	s := time.Now()
	nonce, hash := pow.Run()
	b.Elapsed = time.Since(s).String()
	fmt.Printf("Elepsed time: %s\n", b.Elapsed)

	b.Hash = hash[:]
	b.Nonce = nonce
	return b
}

// CreateGenesisBlock create new block without a hash value
func CreateGenesisBlock(data []byte) *Block {
	return CreateBlock(data, []byte{})
}

// Serialize the block type receiver
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

// Deserialize the block type receiver
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

// Handle an error and logout
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ToString answer a formated string represent current block and the proof of work check as boolean
func (b *Block) ToString(pck bool) string {
	var out bytes.Buffer
	fmt.Fprintln(&out, "----------------------------------------------------------------------------")
	fmt.Fprintf(&out, "Prev. hash: %x\n", b.HashPrev)
	fmt.Fprintf(&out, "Data      : %s\n", b.Data)
	fmt.Fprintf(&out, "Hash      : %x\n", b.Hash)
	fmt.Fprintf(&out, "PoW       : %s\n", strconv.FormatBool(pck))
	fmt.Fprintf(&out, "Elapsed   : %s\n", b.Elapsed)
	fmt.Fprintln(&out, "----------------------------------------------------------------------------")
	return out.String()
}
