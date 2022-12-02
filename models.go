package main

import "time"

type Transaction struct {
	ID                   int
	TransactionID        string
	TransactionUUID      string
	RelUUID              string
	BuyerID              string
	SellerID             string
	ProductID            string
	Price                int
	Volume               int
	Value                int
	TransactionDate      time.Time
	EntryDate            time.Time
	ConfirmDate          time.Time
	CompleteDateBuyer    time.Time
	CompleteDateSeller   time.Time
	BuySell              string
	IsAmend              bool
	IsCancel             bool
	ConfirmStatue        string
	CompleteStatusBuyer  string
	CompleteStatusSeller string
	Status               string
}

type Datamart1 struct {
	TransactionID        string
	BuyerID              string
	BuyerName            string
	SellerID             string
	SellerName           string
	ProductID            string
	ProductName          string
	Currency             string
	Price                int
	Volume               int
	Value                int
	TransactionDate      time.Time
	EntryDate            time.Time
	EntryMonth           int
	EntryYear            int
	BuySell              string
	ConfirmStatue        string
	CompleteStatusBuyer  string
	CompleteStatusSeller string
}

type Datamart2 struct {
	ProductId                              string
	ProductName                            string
	TotalTransaction                       int
	LatestPriceByTransactionDateTime       int
	LatestPriceByEntryDateTime             int
	LatestTotalVolumeByTransactionDateTime int
	LatestTotalVolumeByEntryDateTime       int
	LatestValueByTransactionDateTime       int
	LatestValueByEntryDateTime             int
}

type Datamart3 struct {
	CustomerId                             string
	CustomerName                           string
	TotalTransaction                       int
	LatestPriceByTransactionDateTime       int
	LatestPriceByEntryDateTime             int
	LatestTotalVolumeByTransactionDateTime int
	LatestTotalVolumeByEntryDateTime       int
	LatestValueByTransactionDateTime       int
	LatestValueByEntryDateTime             int
}
