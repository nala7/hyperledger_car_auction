/*
Copyright 2020 IBM All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	log.Println("============ application-golang starts ============")

	var channelID = "mychannel"
	// Identidad con la que se conecta al nodo-par
	var identity = "User1@org1.example.com"

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists(identity) {
		log.Fatalf("Failed to populate wallet contents: %v", err)
	}

	ccpPath := filepath.Join(
		"ccp.yaml",
	)

	fmt.Println(ccpPath)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, identity),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelID)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("car-contract")

	/****************************   Init Ledger   ***********************************/
	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	/****************************   Utils   ***********************************/
	log.Println("--> Evaluate Transaction: QueryCar, return the car with the id specified")
	result, err = contract.EvaluateTransaction("QueryCar", "8")
	if err != nil {
		log.Fatalf("Failed to Evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: QueryAllCars, function returns all the current cars that are being auctioned")
	result, err = contract.EvaluateTransaction("QueryAllCars")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	/****************************   Bids   ***********************************/
	log.Println("--> Evaluate Transaction: Querybid, returns the bid stored in the world state with given id")
	result, err = contract.EvaluateTransaction("Querybid", "2")
	if err != nil {
		log.Fatalf("Failed to Evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: QueryAllBidsForCarNumber, returns all bids made for the car with a given id found in world state")
	result, err = contract.EvaluateTransaction("QueryAllCars", "6")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))
	
	/****************************   Auctioner Methods   ***********************************/
	log.Println("--> Submit Transaction: CreateCar, function create an auction process for a car with details provided")
	result, err = contract.SubmitTransaction("CreateCar", "88", "77", "Lada", "Blue", "Rafael")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: StartAuction, function start the auction of the car with the given carNumber")
	result, err = contract.SubmitTransaction("StartAuction", "88", "Rafael")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: CloseAuction, function end the auction for the car with the given carNumber")
	result, err = contract.SubmitTransaction("CloseAuction", "88", "Rafael")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: VerifyAuction, function validate the auction process for the car with the given carNumber")
	result, err = contract.SubmitTransaction("VerifyAuction", "88", "Rafael")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	/****************************   Buyer Methods   ***********************************/
	log.Println("--> Submit Transaction: CreateBit, user bid for the car with the given details the amount specified")
	result, err = contract.SubmitTransaction("CreateBid", "100", "101", "10 000", "Nadia", "5")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println(string(result))
	log.Println("============ application-golang ends ============")


}
