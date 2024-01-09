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

const MAX_STEP = 5

type BlockCrawler struct {
	Database   mongo.Database
	Collection string
}

type Block struct {
	ID               string `bson:"_id" json:"id"`
	Type             string `bson:"type" json:"code"`
	Number           int32  `bson:"number" json:"number"`
	Hash             string `bson:"hash"  json:"hash"`
	ParentHash       string `bson:"parent_hash" json:"parent_hash"`
	Nonce            string `bson:"nonce" json:"nonce"`
	Sha3Uncles       string `bson:"sha3_uncles" json:"sha3_uncles"`
	LogsBloom        string `bson:"logs_bloom" json:"logs_bloom"`
	TransactionsRoot string `bson:"transactions_root" json:"transactions_root"`
	StateRoot        string `bson:"state_root" json:"state_root"`
	ReceiptsRoot     string `bson:"receipts_root" json:"receipts_root"`
	Miner            string `bson:"miner" json:"miner"`
	Difficultly      string `bson:"difficultly" json:"difficultly"`
	TotalDifficulty  string `bson:"total_difficulty" json:"total_difficulty"`
	Size             string `bson:"size" json:"size"`
	ExtraData        string `bson:"extra_data" json:"extra_data"`
	GasLimit         string `bson:"gas_limit" json:"gas_limit"`
	GasUsed          string `bson:"gas_used" json:"gas_used"`
	Timestamp        int32  `bson:"timestamp" json:"timestamp"`
	TransactionCount int32  `bson:"transaction_count" json:"transaction_count"`
	ItemTimestamp    string `bson:"item_timestamp" json:"item_timestamp"`
}

func (bc *BlockCrawler) Crawl(env *bootstrap.Env) {
	collection := bc.Database.Collection(bc.Collection)

	step := env.NumberOfBlocks / MAX_STEP
	var allBlocks []Block

	for i := 0; i < MAX_STEP; i++ {
		fmt.Println("block_step_", i)
		var blocks []Block
		filter := bson.A{
			bson.D{{"$skip", i * int(step)}},
			bson.D{{"$limit", (i + 1) * int(step)}},
		}
		cursor, err := collection.Aggregate(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		cursor.All(context.Background(), &blocks)
		allBlocks = append(allBlocks, blocks...)
	}

	// Write file
	blocksJson, _ := json.Marshal(allBlocks)
	file := fmt.Sprintf("data/%s_blocks.json", env.DBName)
	err := ioutil.WriteFile(file, blocksJson, 0644)
	if err != nil {
		panic(err)
	}
}
