package main

import (
	"goblock/node"
	"log"
	"os"
)

func main() {
	if _, is := os.LookupEnv("HTTP_ENDPOINT"); !is {
		log.Panic("'HTTP_ENDPOINT' env var is not provided.")
	}
	httpAddr := os.Getenv("HTTP_ENDPOINT")
	node := node.NewNode(httpAddr)
	log.Fatalln(node.Run())
}
