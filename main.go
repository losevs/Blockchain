package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/losevs/Blockchain/chain"
	"github.com/losevs/Blockchain/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	go func() {
		genesisBlock := chain.Block{Index: 0, Timestamp: time.Now().String(), Age: 0, Hash: "", PrevHash: ""}
		spew.Dump(genesisBlock)
		chain.Blockchain = append(chain.Blockchain, genesisBlock)
	}()
	log.Fatal(server.Run())
}
