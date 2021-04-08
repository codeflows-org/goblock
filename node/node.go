package node

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"goblock/packages/api"
)

type Node struct {
	// WIP
	httpAddr string // The endpoint to listen for HTTP incoming traffic.
}

func NewNode(httpAddr string) *Node {
	return &Node{
		httpAddr: httpAddr,
	}
}

func (n *Node) Run() error {
	mux := makeMuxRouter()
	log.Println("Listening on", n.httpAddr)

	s := &http.Server{
		Addr:           ":" + n.httpAddr,
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

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/txns", handleTxnSubmit).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	// TODO: Use a blockchain store/repo.
	// This is just for compilation only.
	bc := api.NewBlockchain()

	bytes, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
}

func handleTxnSubmit(w http.ResponseWriter, r *http.Request) {
	// TODO
}
