package storage

import (
	"log"
	"github.com/dgraph-io/badger"
)

type BlockchainDB struct{
DB *badger.DB
}

func OpenDB(path string) *BlockchainDB{
	opts:=badger.DefaultOptions(path)
	//
	db,err :=badger.Open(opts)
	if err :=nil{
		log.Panic(err)
	}
	return &BlockchainDB{DB:db}
}


// saveBlock: store a block in the badgerDB
func (bdb *BlockchainDB) saveBlock(block *Block) error{
	return bdb.DB.Update(func(txn *badger.Txn) error {
		err!=nil {
			return err
		}
		return nil
	})
}