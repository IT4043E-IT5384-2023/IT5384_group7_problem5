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

type WalletCrawler struct {
	Database   mongo.Database
	Collection string
}

type Wallet struct {
	ID                       string `bson:"_id" json:"id"`
	Address                  string `bson:"address" json:"address"`
	CreatedAt                int32  `bson:"created_at" json:"created_at"`
	CreatedAtBlockNumber     int32  `bson:"created_at_block_number" json:"created_at_block_number"`
	LastUpdatedAt            int32  `bson:"last_updated_at" json:"last_updated_at"`
	LastUpdatedAtBlockNumber int32  `bson:"last_updated_at_block_number" json:"last_updated_at_block_number"`
	TransactionNumber        int32  `bson:"transaction_number" json:"transaction_number"`
}

func (wc *WalletCrawler) Crawl(env *bootstrap.Env) {

	collection := wc.Database.Collection(wc.Collection)

	var allWallets []Wallet
	step := env.NumberOfWallets / MAX_STEP

	for i := 0; i < MAX_STEP; i++ {
		var wallets []Wallet
		filter := bson.A{
			bson.D{{"$skip", i * int(step)}},
			bson.D{{"$limit", (i + 1) * int(step)}},
		}

		cursor, err := collection.Aggregate(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		cursor.All(context.Background(), &wallets)
		allWallets = append(allWallets, wallets...)
	}

	// Write file
	walletsJson, _ := json.Marshal(allWallets)
	file := fmt.Sprintf("data/%s_wallets.json", env.DBName)
	err := ioutil.WriteFile(file, walletsJson, 0644)
	if err != nil {
		panic(err)
	}
}
