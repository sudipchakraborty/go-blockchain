package storage

import (
	"log"
	"../blockchain"
	"github.com/dgraph-io/badger/v3"
)

type BlockchainDB struct{
DB *badger.DB
}
//____________________________________________________________________________________________________
func OpenDB(path string) *BlockchainDB{
	opts:=badger.DefaultOptions(path)
	//
	db,err :=badger.Open(opts)
	if err :=nil{
		log.Panic(err)
	}
	return &BlockchainDB{DB:db}
}
//____________________________________________________________________________________________________
// saveBlock: store a block in the badgerDB
func (bdb *BlockchainDB) saveBlock(block *blockchain.Block) error{
	return bdb.DB.Update(func(txn *badger.Txn) error {
		err!=nil {
			return err
		}
		return nil
	})
}
//____________________________________________________________________________________________________
// GetBlock: get a block from the badgerDB
func (bdb *BlockchainDB) GetBlock(hash [] byte) (*blockchain.Block, error){
	var block *blockchain.Block
	err :=bdb.DB.View(func(txn *badger.Txn) error {
		item, err:=txn.Get(hash)
		if err !=nil{
			return err
		}
		err=item.Value(func(val []byte) error {
			block=blockchain.DeserializeBlock(val)
			return nil
		})
		return err
	})
		if err !=nil{
			return nil, err
		}
		return block, nil
}
//____________________________________________________________________________________________________
func(bdb *BlockchainDB) CloseDB(){
	err :=bdb.DB.Close()
	if err != nil {
		log.Panic(err)
	}
}