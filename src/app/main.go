package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"ghostchaser/kolxi-blockchain/src/models/genesisblock"
	"ghostchaser/kolxi-blockchain/src/responses/healthresponse"
	"ghostchaser/kolxi-blockchain/src/service"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var hResponse healthresponse.HealthResponse
	hResponse.Ok = true
	hResponse.Message = "Server is alive."

	json.NewEncoder(response).Encode(hResponse)

}

func main() {

	fmt.Println("Init server...")

	service.LoadGenesisFromDisk()

	absPath, _ := filepath.Abs("./src/database/genesis/genesis.json")
	file, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal("Failed to read genesis block file")
	}

	var genBlock genesisblock.GenesisBlock

	// Read json file and map tp GenesisBlock struct
	_ = json.Unmarshal([]byte(file), &genBlock)

	fmt.Println(genBlock.CreatedAt)

	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheckHandler)

	http.ListenAndServe(":"+"5200", router)
}
