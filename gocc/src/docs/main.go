package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	fmt.Println("Starting....")
	err := shim.Start(new(DocsChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
