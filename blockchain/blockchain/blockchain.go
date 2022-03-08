package blockchain

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/raibru/HomeWork/blockchain/config"

	"github.com/dgraph-io/badger/v3"
)

const (
	dbPath        = "./tmp/blocks"
	dbFile        = "./tmp/blocks/MANIFEST"
	rootBlockData = "Root Block"
)

// BlockChain holds the collection of blocks
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// HashIterator data type
type HashIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// Init initialize the blockcahin with initial values
func Init() error {
	return nil
}

// DBexists check that the database file exists
func DBexists() bool {
	dbf := path.Join((*config.GetConfig()).DbPath, (*config.GetConfig()).DbFile)
	if _, err := os.Stat(dbf); os.IsNotExist(err) {
		return false
	}

	return true
}

// ContinueBlockChain access the last block of an existing blockchain
func ContinueBlockChain(address string) *BlockChain {
	if !DBexists() {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions((config.GetConfig()).DbPath)
	//opts.Dir = (*config.GetConfig()).DbPath
	//opts.ValueDir = (*config.GetConfig()).DbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		Handle(err)
		return err
	})
	Handle(err)

	chain := BlockChain{lastHash, db}

	return &chain
}

// CreateBlockChain create a new blockchain with one root block
func CreateBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	//opts.Dir = dbPath
	//opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		var lherr error
		_, err := txn.Get([]byte("lh"))
		if err == badger.ErrKeyNotFound {
			lherr = initialLastHash(&lastHash, txn)
		} else {
			lherr = getLastHash(&lastHash, txn)
		}
		return lherr
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

// AddBlock create a new block with proof of work and append it in the blockchain
func (chain *BlockChain) AddBlock(data []byte) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)
}

// Iterator returns a HashIterator
func (chain *BlockChain) Iterator() *HashIterator {
	iter := &HashIterator{chain.LastHash, chain.Database}

	return iter
}

// Next answer the next block from HashIterator
func (iter *HashIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)

		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})
		Handle(err)
		return err
	})
	Handle(err)

	iter.CurrentHash = block.HashPrev
	return block
}

func initialLastHash(lastHash *[]byte, txn *badger.Txn) error {
	fmt.Println("No existing blockchain found")
	gb := CreateGenesisBlock([]byte((*config.GetConfig()).Root))
	fmt.Println("Genesis block proved")

	var err error
	err = txn.Set(gb.Hash, gb.Serialize())
	Handle(err)
	err = txn.Set([]byte("lh"), gb.Hash)

	*lastHash = gb.Hash[:]

	return err
}

func getLastHash(lastHash *[]byte, txn *badger.Txn) error {
	item, err := txn.Get([]byte("lh"))
	Handle(err)
	err = item.Value(func(val []byte) error {
		*lastHash = val
		return nil
	})
	return err
}
