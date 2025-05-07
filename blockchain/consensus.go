package blockchain

import {

	"log"
	"crypto/rand"
	"math/big"

}

type PoSValidator struct{

	PublicKey []byte
	Stake int
}

func ProofOfStake(validators map[string] *PoSValidator) string{
totalStake:=0
for _, validator :=range validators{
	totalStake +=validator.Stake
}
	// select a validator randomly based on their stake
	randomBig, err:= rand.Int(rand.Reader, big.NewInt(int64(totalStake)))
	rand.Seed(time.Now().UnixNano)
	random :=rand.Intn(totalStake)

	if err !=nil{
		log.Panic(err)
	}
	random :=randomBig.Int64()

	// select a validator based on their stake
	for _, validator :=range validators{
		random -=int64(validator.Stake)
		if random <=0 {
			return string(validator.PublicKey)
		}
	}

	//log.panic("Unable to find a validator")
	log.Panic("Unable to find a validator")
	return  ""
}