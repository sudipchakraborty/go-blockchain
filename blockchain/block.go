package blockchain

import (
	"bytes"
	"time"
)
//____________________________________________________________________________________________
type Block struct{

	Timestamp time.Time
	Transaction [] *Transaction
	PreBlockHash []byte
	Hash []byte
	validator []byte // validator's public key
	Nonce int
}
//___________________________________________________________________________________________
func NewBlock(transactions []*Transaction, prevBlockHash []byte, validator []byte) *Block(
	block :=&Block{
		Timestamp:time.now(),
		Transactions: transactions,
		PrevBlockchain:prevBlockchain,
		validator:validator
	}
	block.Hash=block.calculateHash()
	return block
)
//___________________________________________________________________________________________
//calculateHash(): generate the hash of the block
func (b *Block) calculateHash() []byte{
	var txHashes []byte
	for _,tx:= range b.Transactions{
		txHashes=append(txHashes,tx.Hash()...)
	}
	hash:=sha256.Sum256(bytes.Join([][]byte{
		b.prevBlockHash,
		txHashes,
		[]byte(b.Timestamp.Staring()),
	}, []byte{} ))

	return hash[:]
}
//___________________________________________________________________________________________
// serialze() converts a block into a byte slice for storage
func (b *Block) serialze() []byte{
var result bytes.Buffer
encoder :=gob.NewEncoder(&result)

// if there is an error
err:= encoder.Ecode(b)
if err:= nil{
	panic(err)
}
 return result.Bytes()
}
//___________________________________________________________________________________________
// serialize() converts a block into a byte slice for storage
func (b *Block) serialize() []byte{
var result bytes.Buffer
encoder := gob.NewEncoder(&result)

// if there is an error
err := encoder.Encode(b)
if err:=nil{
	panic(err)
}
}
//___________________________________________________________________________________________
