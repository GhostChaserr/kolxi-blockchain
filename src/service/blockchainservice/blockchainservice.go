package blockchainservice

// Resources
// https://jeiwan.net/posts/building-blockchain-in-go-part-1/

import (
	"fmt"
	"ghostchaser/kolxi-blockchain/src/models/transaction"
)

type Block struct {
	Timestamp    string
	Index        int64
	Proof        int
	PreviousHash string
	Transactions []transaction.Transaction
}

type Blockchain struct {
	Chain        []Block
	Transactions []transaction.Transaction
}

// Creates new block
func (b *Blockchain) CreateBlock(proof int64, previousHash string) int64 {

	// New Block
	var block Block
	block.Index = 1
	block.Timestamp = "2020"
	block.Proof = 98715
	block.PreviousHash = "new hash"
	block.Transactions = b.Transactions

	// Adds new block to blockchain
	b.Chain = append(b.Chain, block)

	return block.Index

}

// Returns Latest block
func (b *Blockchain) GetLatestBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

// TODO.
// 1. String slicing
// 2. Generate jasj operation

func (b *Blockchain) ProofOfWork(previousProof int64) int64 {

	var newProof int64 = 1
	var proofFound bool = false

	for proofFound == false {

		fmt.Println("Running..")
	}

	return newProof

}

// type State struct {
// 	Balances  map[string]uint
// 	txMempool []transaction.Transaction
// 	dbFile    *os.File
// }

// // Adds new transaction
// func (s *State) AddTransaction(tx transaction.Transaction) bool {
// 	fmt.Println(s)
// 	return true
// }

// Recalculates new state from disk.
// func NewStateFromDisk() (*State, error) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 	return nil, err }
// }
