package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const dbPath = "./tmp/blocks"

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("no existing blockchain found")
			genesis := Genesis()
			fmt.Println("genesis created")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				panic("an error occured while creating genesis block")
			}
			lastHash = genesis.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			if err != nil {
				panic("error occured while getting last hash")
			}
			item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})
	if err != nil {
		panic("error occured while initializing blockchain")
	}
	return &BlockChain{
		Database: db,
		LastHash: lastHash,
	}

}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}
