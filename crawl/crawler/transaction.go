package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/quandat10/bigdata-crawl/bootstrap"
	"github.com/quandat10/bigdata-crawl/mongo"
	"go.mongodb.org/mongo-driver/bson"
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
	ID             string `bson:"_id" json:"id"`
	FromAddress    string `bson:"from_address" json:"from_address"`
	ToAddress      string `bson:"to_address" json:"to_address"`
	Value          string `bson:"value" json:"value"`
	BlockTimestamp int32  `bson:"block_timestamp" json:"block_timestamp"`
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
