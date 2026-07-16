package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for an open ascending auction
type SmartContract struct {
	contractapi.Contract
}

// Bid records one accepted bid.
type Bid struct {
	Bidder string `json:"Bidder"` // bidder name
	Amount int    `json:"Amount"` // bid amount in paise
}

// Auction represents an open auction and its bids.
type Auction struct {
	AuctionID string `json:"AuctionID"` // unique auction id, e.g. "auc1"
	Item      string `json:"Item"`      // item being auctioned
	Status    string `json:"Status"`    // OPEN | CLOSED | CANCELLED
	Bids      []Bid  `json:"Bids"`      // accepted bids in order
	Winner    string `json:"Winner"`    // highest bidder, set when closed
	HighBid   int    `json:"HighBid"`   // current highest amount
}

// CreateAuction opens a new auction with status "OPEN" and no bids.
// It must fail if the auction already exists.
func (s *SmartContract) CreateAuction(ctx contractapi.TransactionContextInterface, auctionID string, item string) error {

	return nil
}

// GetAuction returns the auction identified by auctionID.
// It must fail if the auction does not exist.
func (s *SmartContract) GetAuction(ctx contractapi.TransactionContextInterface, auctionID string) (*Auction, error) {

	return nil, nil
}

// PlaceBid records a bid. A bid is only accepted if it is strictly higher than
// the current HighBid; accepted bids update HighBid.
// It must fail if the auction does not exist, is not OPEN, amount is not
// positive, or amount is not higher than the current HighBid.
func (s *SmartContract) PlaceBid(ctx contractapi.TransactionContextInterface, auctionID string, bidder string, amount int) error {

	return nil
}

// GetHighestBid returns the current highest bid amount.
// It must fail if the auction does not exist.
func (s *SmartContract) GetHighestBid(ctx contractapi.TransactionContextInterface, auctionID string) (int, error) {

	return 0, nil
}

// CloseAuction transitions the auction from "OPEN" to "CLOSED" and records the
// highest bidder as the Winner.
// It must fail if the auction does not exist, is not OPEN, or has no bids.
func (s *SmartContract) CloseAuction(ctx contractapi.TransactionContextInterface, auctionID string) error {

	return nil
}

// CancelAuction sets the auction's status to "CANCELLED". A cancelled auction
// has no winner.
// It must fail if the auction does not exist or is not OPEN.
func (s *SmartContract) CancelAuction(ctx contractapi.TransactionContextInterface, auctionID string) error {

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic("Error creating auction chaincode: " + err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic("Error starting auction chaincode: " + err.Error())
	}
}
