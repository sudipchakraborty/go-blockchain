package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


var bc *blockchain.blockchain
var db *blockchain.BlockchainDB

func StartServer(blockchain *blockchain.Blockchain, database *blockchain.BlockchainDB){
bc=blockchain
db=database

e :=echo.Now()

	//Define API routes
	e.POST("/transaction",handleTransaction)
	e.GET("/block/:hash",handeGetBlock)
	e.POST("/contract",handleDeployContract)
	e.POST("/contract/execute",handleExecuteContract)

	// Start the Server
	e.Logger.Fatal(e.Start(":1323"))
}

// HandleTransaction: process a new transaction
func handleTransaction(c echo.Context) error {

	from :=c.QueryParam("from")
	to :=c.QueryParam("to")
	amount :=c.Queryparam("amount")

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "transaction submitted",
		"from" :	from,
		"to" :	to,
		"amount" :amount,
	})
}

// handleGetBlock: retrieve a block by its hash
func handleGetBlock(c echo.Context) error{

	hash:= c.Param("hash")

	block,err := bc.GetBlock(hash)

	if err !=nil{
		return c.JSON(http.StatusNotFound,map[string]string){
			"message": "Block not found",
		})

	}
		return c.JSON(http.StatusOK,block)
}


func  handleDeployContract(c echo.Context) error  {
	id:= c.Queryparam("id")
	code :=c.QueryParam("code")

	contract :=contract.NewSmartContract(id,code)
	err:= contract.validate()

	if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"message":err.Error(),
		})
	}

	return c.JSON(http.StatusCretaed,map[string] string{
		"message": "Contract deployed",
		"id": is,
	})
	
}



func handleExecuteContract(c echo.Context) error  {
	id:= c.QueryParam("id")
	input :=map[string] interface{}{}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Contract executed",
		"id": id,
		"input":input,
	})
}