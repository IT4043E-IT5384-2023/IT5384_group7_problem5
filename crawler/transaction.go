package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/quandat10/bigdata-crawl/bootstrap"
	"github.com/quandat10/bigdata-crawl/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
)

type TransactionCrawler struct {
	Database   mongo.Database
	Collection string
}

type DecodedInput struct {
	FromAddress  string  `bson:"from_address" json:"from_address"`
	ToAddress    string  `bson:"to_address" json:"to_address"`
	AssetAddress string  `bson:"asset_address" json:"asset_address"`
	Value        float32 `bson:"value" json:"value"`
}

type Transaction struct {
	ID               string `bson:"_id" json:"id"`
	Type             string `bson:"type" json:"type"`
	Hash             string `bson:"hash" json:"hash"`
	Nonce            int32  `bson:"nonce" json:"nonce"`
	TransactionIndex int32  `bson:"transaction_index" json:"transaction_index"`
	FromAddress      string `bson:"from_address" json:"from_address"`
	ToAddress        string `bson:"to_address" json:"to_address"`
	Value            string `bson:"value" json:"value"`
	Gas              string `bson:"gas" json:"gas"`
	GasPrice         string `bson:"gas_price" json:"gas_price"`
	Input            string `bson:"input" json:"input"`
	BlockTimestamp   int32  `bson:"block_timestamp" json:"block_timestamp"`
	BlockNumber      int32  `bson:"block_number" json:"block_number"`
	//BlockHash              *string        `bson:"block_hash" json:"block_hash"`
	//ReceiptGasUsed         *string       `bson:"receipt_gas_used" json:"receipt_gas_used"`
	//ReceiptContractAddress *string       `bson:"receipt_contract_address" json:"receipt_contract_address"`
	//ReceiptRoot            *string       `bson:"receipt_root" json:"receipt_root"`
	//ReceiptStatus          int32         `bson:"receipt_status" json:"receipt_status"`
	//ItemTimestamp          *string       `bson:"item_timestamp" json:"item_timestamp"`
	//DecodedInput           *DecodedInput `bson:"decoded_input" json:"decoded_input"`
	//TransactionType        *string       `bson:"transaction_type" json:"transaction_type"`
}

func (tc *TransactionCrawler) Crawl(env *bootstrap.Env) {
	collection := tc.Database.Collection(tc.Collection)
	var wallets []Wallet
	fileWallet := fmt.Sprintf("data/%s_wallets.json", env.DBName)
	data, _ := ioutil.ReadFile(fileWallet)

	err := json.Unmarshal(data, &wallets)

	var bsonA bson.A

	for _, wallet := range wallets {
		bsonA = append(bsonA, bson.D{{"from_address", wallet.Address}})
		bsonA = append(bsonA, bson.D{{"to_address", wallet.Address}})
	}

	filter := bson.D{
		{"$or",
			bsonA,
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	var transactions []Transaction

	cursor.All(context.Background(), &transactions)

	// Write file
	transactionsJson, _ := json.Marshal(transactions)
	file := fmt.Sprintf("data/%s_transactions.json", env.DBName)
	err = ioutil.WriteFile(file, transactionsJson, 0644)
}
