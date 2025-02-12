package main

import (
	"log"

	"ssaisearch/api/handler"
)

func main() {
	initAll()
	g.POST("/verify-proposal", handler.VerifyProposal)
	startGin()
}

func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}
