package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/quandat10/bigdata-crawl/bootstrap"
	"github.com/quandat10/bigdata-crawl/crawler"
	"github.com/quandat10/bigdata-crawl/mongo"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	// Goroutine for crawl wallet block and wallet
	var wg sync.WaitGroup
	wg.Add(1)
	go walletCrawl(&wg, db, env)
	// go blockCrawl(&wg, db, env)
	wg.Wait()

	// Goroutine for crawl transactions
	var wt sync.WaitGroup
	wt.Add(1)
	go transactionCrawl(&wt, db, env)
	wt.Wait()

	fmt.Println("+++++COMPLETE+++++")
}

func transactionCrawl(wg *sync.WaitGroup, db mongo.Database, env *bootstrap.Env) {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	fmt.Println("====TRANSACTIONS====")
	tc := crawler.TransactionCrawler{
		Database:   db,
		Collection: "transactions",
	}
	s.Start() // Start the spinner
	tc.Crawl(env)
	s.Stop()
	defer wg.Done()
}

func blockCrawl(wg *sync.WaitGroup, db mongo.Database, env *bootstrap.Env) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	bc := crawler.BlockCrawler{
		Database:   db,
		Collection: "blocks",
	}
	s.Start() // Start the spinner
	bc.Crawl(env)
	s.Stop()
	defer wg.Done()
}

func walletCrawl(wg *sync.WaitGroup, db mongo.Database, env *bootstrap.Env) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	wc := crawler.WalletCrawler{
		Database:   db,
		Collection: "wallets",
	}
	s.Start() // Start the spinner
	wc.Crawl(env)
	s.Stop()
	defer wg.Done()
}
