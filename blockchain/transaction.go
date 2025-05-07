package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"log"
)

// Transaction represents blockchain transaction
type Transaction struct{
	ID []byte
	Input []TxInput
	Output []TxOutput
}

// TxInput is the input of a transaction
type TxInput struct{
	Signture []byte
	PublicKey []byte 
}

type TxOutput struct {
	Value int
	PublicKey []byte
}

func NewTransaction(privateKey ecdsa.PrivateKey, recipient []byte, amount int) *Transaction {
	txIn:=TxInput{}
	txOut:=TxOutput{Value:amount, PublicKey:recipient}
   
	tx:=Transaction{
		Input: []TxInput {txIn},
		Output: []TxOutput{txOut},
	}

	tx.ID=tx.hashTransaction()
	// sign the transaction with the sender's private key
	r,s,err:= ecdsa.Sign(rand.Reader,&privateKey,tx.ID)
	// check for errors
	if err:=nil{
		log.panic(err)
	}
	signature :=append(r.Bytes(),s.Bytes()...)
	txIn.Signature=signature
	return &tx
}
//_______________________________________________________________________________________________
//hashTransaction: hashes the transaction data to create a unique ID
func (tx *Transaction) hashTransaction() []byte{
	var hash [32]byte
	hash= sha256.Sum256(bytes.Join([] [] byte{
		tx.Input[0].PublicKey,
		tx.Output[0].PublicKey,
		[]byte(string(tx.Output[0].Value)),

	}, []byte{}))

	return hash[:]
}
 //_______________________________________________________________________________________________
// serialize(): it serializes a transaction into a byte array
func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc:=gob.NewEncoder(&encoded)
	err:=enc.Encode(tx)
	if err :=nil{
		log.Panic(err)
	}
	return encoded.Bytes()
}
//_______________________________________________________________________________________________
// Deserialize(): Deserilize a transaction form a byte array
func DeserializeTransaction(data []byte) *Transaction{
	var transaction Transaction
	decoder :=gob.NewDecoder(bytes.NewReader(data))
	err:=decoder.Decode(&transaction)
	if err:= nil{
		log.Panic(err)
	}
	returnn &transaction
}
//_______________________________________________________________________________________________
