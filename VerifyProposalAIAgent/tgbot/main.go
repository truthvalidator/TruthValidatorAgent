package main

import (
	"log"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/tgbot/cmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {

	cmd.Execute()
}
