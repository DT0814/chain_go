package main

import (
	"chain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainsetHandler)
	http.ListenAndServe("localhost:8888", nil)
}
func blockchainGetHandler(write http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(write, error.Error(), http.StatusInternalServerError)
		return
	} else {
		io.WriteString(write, string(bytes))
	}
}
func blockchainsetHandler(write http.ResponseWriter, r *http.Request) {
	blockchain.Print()
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(write, r)
}
func main() {
	blockchain = core.NewBlockChain()
	run()
}
