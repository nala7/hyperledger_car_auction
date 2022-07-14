/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	ID   string `json:"id"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	State string `json:"state"`  // sold/forSale/reservation
	Owner  string `json:"owner"` 
}

// Bid describes basic details of what makes up a bid
type Bid struct {
	ID   string `json:"id"`
	Price int `json:"price"` 
	Owner  string `json:"owner"` 
	Car string `json:"car"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// QueryResultBid structure used for handling result of query with bid
type QueryResultBid struct {
	Key    string `json:"Key"`
	Record *Bid
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{ID: "1", Model: "Toyota", Colour: "blue", State: "sold", Owner: "Ariel"},
		Car{ID: "2", Model: "Ford", Colour: "red", State: "forSale", Owner: "Luis"},
		Car{ID: "3", Model: "Hyundai", Colour: "green", State: "sold", Owner: "Nadia"},
		Car{ID: "4", Model: "Volkswagen", Colour: "yellow", State: "reservation", Owner: "Amalia"},
		Car{ID: "5", Model: "Tesla", Colour: "black", State: "reservation", Owner: "Luis"},
		Car{ID: "6", Model: "Peugeot", Colour: "purple", State: "sold", Owner: "Gabriela"},
		Car{ID: "7", Model: "Chery", Colour: "white", State: "forSale", Owner: "Amalia"},
		Car{ID: "8", Model: "Fiat", Colour: "violet", State: "forSale", Owner: "Labourdette"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	bids := []Bid{
		Bid{ID: "1", Price: 1500, Owner: "Lena", Car: "1"},
		Bid{ID: "2", Price: 1600, Owner: "karla", Car: "2"},
		Bid{ID: "3", Price: 2000, Owner: "Hugo", Car: "1"},
		Bid{ID: "4", Price: 2100, Owner: "Nadia", Car: "3"},
		Bid{ID: "5", Price: 3000, Owner: "Amalia", Car: "1"},
		Bid{ID: "6", Price: 1000, Owner: "Hector", Car: "3"},
		Bid{ID: "7", Price: 4500, Owner: "Ariel", Car: "1"},
		Bid{ID: "8", Price: 600, Owner: "Bonifacio", Car: "3"},
		Bid{ID: "9", Price: 5000, Owner: "SeoYong", Car: "8"},
		Bid{ID: "10", Price: 3200, Owner: "Luffy", Car: "1"},
		Bid{ID: "11", Price: 20, Owner: "Labourdette", Car: "3"},
		Bid{ID: "12", Price: 1020, Owner: "Luis", Car: "7"},
		Bid{ID: "13", Price: 1040, Owner: "Nadia", Car: "2"},
		Bid{ID: "14", Price: 7000, Owner: "Gabriela", Car: "8"},
		Bid{ID: "15", Price: 5200, Owner: "Ariel", Car: "2"},
		Bid{ID: "16", Price: 4300, Owner: "Amalia", Car: "8"},
		Bid{ID: "17", Price: 6000, Owner: "Mickey", Car: "7"},
		Bid{ID: "18", Price: 2000, Owner: "Minnie", Car: "8"},
		Bid{ID: "19", Price: 100, Owner: "Labourdette", Car: "7"},
		Bid{ID: "20", Price: 2200, Owner: "Roly", Car: "6"},
		Bid{ID: "21", Price: 2700, Owner: "Gabriela", Car: "6"},
	}

	for i, bid := range bids {
		bidAsBytes, _ := json.Marshal(bid)
		err := ctx.GetStub().PutState("BID"+strconv.Itoa(i), bidAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**************************   Utils  ***********************************************/

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

	return results, nil
}

// QueryBid returns the bid stored in the world state with given id
func (s *SmartContract) QueryBid(ctx contractapi.TransactionContextInterface, bidNumber string) (*Bid, error) {
	bidAsBytes, err := ctx.GetStub().GetState(bidNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if bidAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", bidNumber)
	}

	bid := new(Bid)
	_ = json.Unmarshal(bidAsBytes, bid)

	return bid, nil
}

// QueryAllBidsForCarNumber returns all bids made for the car with a given id found in world state
func (s *SmartContract) QueryAllBidsForCarNumber(ctx contractapi.TransactionContextInterface, carNumber string) ([]QueryResultBid, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResultBid{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		bid := new(Bid)
		_ = json.Unmarshal(queryResponse.Value, bid)

		queryResult := QueryResultBid{Key: queryResponse.Key, Record: bid}

		if bid.Car == carNumber
		{
			results = append(results, queryResult)
		}		
	}

	return results, nil
}

/**************************   Auctioneer's methods  ***********************************************/

// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, id string, model string, colour string, owner string) error {
	car := Car{
		ID:   id,
		Model:  model,
		Colour: colour,
		State: "reservation",
		Owner:  owner,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// StartAuction updates the status field of car with given id in world state to forSale
func (s *SmartContract) StartAuction(ctx contractapi.TransactionContextInterface, carNumber string, user string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.Owner != user {
		return fmt.Errorf("%s you don't have access to this auction", user)
	}

	if car.State != "reservation" {
		return fmt.Errorf("%s, this auction cannot be started because its status is not a reservation", user)
	}

	car.State = "forSale"

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// CloseAuction updates the owner field of car with given id in world state to sold
func (s *SmartContract) CloseAuction(ctx contractapi.TransactionContextInterface, carNumber string, user string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.Owner != user {
		return fmt.Errorf("%s you don't have access to this auction", user)
	}

	if car.State == "sold" {
		return fmt.Errorf("%s, this auction is already closed", user)
	}

	if car.State == "reservation" {
		return fmt.Errorf("%s, this auction has not started", user)
	}

	car.State = "sold"

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// VerifyAuction updates the owner field of car with given id in world state
func (s *SmartContract) VerifyAuction(ctx contractapi.TransactionContextInterface, carNumber string, user string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.Owner != user {
		return fmt.Errorf("%s you don't have access to this auction", user)
	}

	if car.State == "forSale" {
		return fmt.Errorf("%s, this auction has not closed", user)
	}

	if car.State == "reservation" {
		return fmt.Errorf("%s, this auction has not started", user)
	}

	buyers, err := s.QueryAllBidsForCarNumber(ctx, carNumber)

	max := buyers[0]

	for i := 1 ; i < len(buyers); i++ {
		if buyers[i].Price > max.Price{
			max = buyers[i]
		}
	} 

	car.Owner = max.Owner

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

/**************************   Buyer's methods  ***********************************************/

// CreateBid adds a new bid to the world state with given details
func (s *SmartContract) CreateBid(ctx contractapi.TransactionContextInterface, bidNumber string, id string, price int, owner string, carNumber string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	if car.Owner == user {
		return fmt.Errorf("%s you can't bid on your own auction", user)
	}

	if car.State == "sold" {
		return fmt.Errorf("%s, this auction is already closed", user)
	}

	if car.State == "reservation" {
		return fmt.Errorf("%s, this auction has not started", user)
	}

	bid := Bid{
		ID:   id,
		Price:  price,
		Owner: owner,
		Car: carNumber,
	}

	bidAsBytes, _ := json.Marshal(bid)

	return ctx.GetStub().PutState(bidNumber, bidAsBytes)
}

/**************************  Main function  ***********************************************/

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create auction chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting auction chaincode: %s", err.Error())
	}
}