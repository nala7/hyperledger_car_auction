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

	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: GetCars, function returns all the current cars that are being auctioned")
	result, err = contract.EvaluateTransaction("GetCars")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: GetCar, return the car with the id specified")
	result, err = contract.EvaluateTransaction("GetCar", "524")
	if err != nil {
		log.Fatalf("Failed to Evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Create, function create an auction process for a car with details provided")
	result, err = contract.SubmitTransaction("Create", "Lada", "Blue")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Start, function start the auction of the car with the given id")
	result, err = contract.SubmitTransaction("Start", "car18552")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Close, function end the auction for the car with the given id")
	result, err = contract.SubmitTransaction("Close", "car18552")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Validate, function validate the auction process for the car with the given id")
	result, err = contract.SubmitTransaction("Validate", "car18552")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Bid, user bid for the car with the given id the amount specified")
	result, err = contract.SubmitTransaction("Bid", "car18552", "15 000")
	if err != nil {
		log.Fatalf("Failed to submit transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println(string(result))
	log.Println("============ application-golang ends ============")


}