package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Account string

type Transaction struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

func (t Transaction) IsReward() bool {
	return t.Data == "reward"
}

type Genesis struct {
	Balances map[Account]uint `json:"balances"`
}

type State struct {
	Balances  map[Account]uint
	txMempool []Transaction
	dbFile    *os.File
}

func LoadGenesis(filePath string) (Genesis, error) {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)

	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}

func LoadGenesisFromDisk() {
	filePath, err := filepath.Abs("./src/database/genesis/genesis.json")
	if err != nil {
		fmt.Println(err)
	}

	gen, err := LoadGenesis(filePath)
	if err != nil {
		fmt.Println(err)
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	transactionsPath, _ := filepath.Abs("./src/database/db/txt.db")
	file, err := os.OpenFile(transactionsPath, os.O_APPEND|os.O_RDWR, 0600)

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	state := &State{balances, make([]Transaction, 0), file}

	for scanner.Scan() {

		var tx Transaction
		json.Unmarshal(scanner.Bytes(), &tx)

		state.apply(tx)

	}

	fmt.Println(state)

}

func (s *State) apply(tx Transaction) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if s.Balances[tx.From] < tx.Value {
		log.Fatal("no funds!")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}
