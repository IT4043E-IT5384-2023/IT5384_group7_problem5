package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/quandat10/bigdata-crawl/bootstrap"
	"github.com/quandat10/bigdata-crawl/crawler"
	"github.com/quandat10/bigdata-crawl/mongo"
	"sync"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	// Goroutine for crawl wallet block and wallet
	var wg sync.WaitGroup
	wg.Add(2)
	go walletCrawl(&wg, db, s, env)
	go blockCrawl(&wg, db, s, env)
	wg.Wait()

	// Goroutine for crawl transactions
	var wt sync.WaitGroup
	wt.Add(1)
	go transactionCrawl(&wt, db, s, env)
	wt.Wait()

	fmt.Println("+++++COMPLETE+++++")
}

func transactionCrawl(wg *sync.WaitGroup, db mongo.Database, s *spinner.Spinner, env *bootstrap.Env) {
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

func blockCrawl(wg *sync.WaitGroup, db mongo.Database, s *spinner.Spinner, env *bootstrap.Env) {
	fmt.Println("====BLOCKS====")
	bc := crawler.BlockCrawler{
		Database:   db,
		Collection: "blocks",
	}
	s.Start() // Start the spinner
	bc.Crawl(env)
	s.Stop()
	defer wg.Done()
}

func walletCrawl(wg *sync.WaitGroup, db mongo.Database, s *spinner.Spinner, env *bootstrap.Env) {
	fmt.Println("====WALLETS====")
	wc := crawler.WalletCrawler{
		Database:   db,
		Collection: "wallets",
	}
	s.Start() // Start the spinner
	wc.Crawl(env)
	s.Stop()
	defer wg.Done()
}
