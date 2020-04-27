package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Circlemono/simpelBlock/Block"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

//获取区块内容
func handleGetBlockchain(w http.ResponseWriter,
	r *http.Request) {
	bytes, err := json.MarshalIndent(Block.Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))

}

type Message struct {
	BPM int
}

//写入区块信息
func handleWriteBlock(w http.ResponseWriter,
	r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}

	defer r.Body.Close()

	newBlock, err := Block.GenerateBlock(&Block.Blockchain[len(Block.Blockchain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if Block.IsBlockValid(newBlock, Block.Blockchain[len(Block.Blockchain)-1]) {
		newBlockchain := append(Block.Blockchain, newBlock)
		Block.ReplaceChain(newBlockchain)
		spew.Dump(Block.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

//回发消息转Json
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func run() error {
	mux := makeMuxRouter()
	s := &http.Server{
		Addr:           ":8088",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {

	fmt.Println("main run")
	t := time.Now()
	genesisBlock := Block.BlockData{
		0,
		t.String(),
		0,
		"",
		"",
	}
	spew.Dump(genesisBlock)
	Block.Blockchain = append(Block.Blockchain, genesisBlock)
	log.Fatal(run())
}
