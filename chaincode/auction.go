package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for an open ascending auction
type SmartContract struct {
	contractapi.Contract
}

// Bid records one accepted bid.
type Bid struct {
	Bidder string `json:"Bidder"`
	Amount int    `json:"Amount"`
}

// Auction represents an open auction and its bids.
type Auction struct {
	AuctionID string `json:"AuctionID"`
	Item      string `json:"Item"`
	Status    string `json:"Status"`
	Bids      []Bid  `json:"Bids"`
	Winner    string `json:"Winner"`
	HighBid   int    `json:"HighBid"`
}

const (
	statusOpen      = "OPEN"
	statusClosed    = "CLOSED"
	statusCancelled = "CANCELLED"
)

// CreateAuction opens a new auction with status "OPEN" and no bids.
func (s *SmartContract) CreateAuction(ctx contractapi.TransactionContextInterface, auctionID string, item string) error {
	existing, err := ctx.GetStub().GetState(auctionID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existing != nil {
		return fmt.Errorf("auction %s already exists", auctionID)
	}

	auction := Auction{
		AuctionID: auctionID,
		Item:      item,
		Status:    statusOpen,
		Bids:      []Bid{},
	}
	return putAuction(ctx, &auction)
}

// GetAuction returns the auction identified by auctionID.
func (s *SmartContract) GetAuction(ctx contractapi.TransactionContextInterface, auctionID string) (*Auction, error) {
	data, err := ctx.GetStub().GetState(auctionID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("auction %s does not exist", auctionID)
	}

	var auction Auction
	if err := json.Unmarshal(data, &auction); err != nil {
		return nil, err
	}
	return &auction, nil
}

// PlaceBid records a bid that is strictly higher than the current HighBid.
func (s *SmartContract) PlaceBid(ctx contractapi.TransactionContextInterface, auctionID string, bidder string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("bid amount must be positive")
	}

	auction, err := s.GetAuction(ctx, auctionID)
	if err != nil {
		return err
	}
	if auction.Status != statusOpen {
		return fmt.Errorf("auction %s is not OPEN (current status: %s)", auctionID, auction.Status)
	}
	if amount <= auction.HighBid {
		return fmt.Errorf("bid %d does not beat the current high bid of %d", amount, auction.HighBid)
	}

	auction.Bids = append(auction.Bids, Bid{Bidder: bidder, Amount: amount})
	auction.HighBid = amount
	return putAuction(ctx, auction)
}

// GetHighestBid returns the current highest bid amount.
func (s *SmartContract) GetHighestBid(ctx contractapi.TransactionContextInterface, auctionID string) (int, error) {
	auction, err := s.GetAuction(ctx, auctionID)
	if err != nil {
		return 0, err
	}
	return auction.HighBid, nil
}

// CloseAuction closes the auction and records the highest bidder as Winner.
func (s *SmartContract) CloseAuction(ctx contractapi.TransactionContextInterface, auctionID string) error {
	auction, err := s.GetAuction(ctx, auctionID)
	if err != nil {
		return err
	}
	if auction.Status != statusOpen {
		return fmt.Errorf("auction %s is not OPEN (current status: %s)", auctionID, auction.Status)
	}
	if len(auction.Bids) == 0 {
		return fmt.Errorf("auction %s has no bids and cannot be closed with a winner", auctionID)
	}

	// Every accepted bid raised HighBid, so the last bid in the list is the
	// highest one.
	auction.Winner = auction.Bids[len(auction.Bids)-1].Bidder
	auction.Status = statusClosed
	return putAuction(ctx, auction)
}

// CancelAuction sets the auction's status to "CANCELLED".
func (s *SmartContract) CancelAuction(ctx contractapi.TransactionContextInterface, auctionID string) error {
	auction, err := s.GetAuction(ctx, auctionID)
	if err != nil {
		return err
	}
	if auction.Status != statusOpen {
		return fmt.Errorf("auction %s is not OPEN (current status: %s)", auctionID, auction.Status)
	}

	auction.Status = statusCancelled
	return putAuction(ctx, auction)
}

// --- helpers ---

func putAuction(ctx contractapi.TransactionContextInterface, auction *Auction) error {
	bytes, err := json.Marshal(auction)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(auction.AuctionID, bytes)
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
