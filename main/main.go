package main

import (
	"parallel/irc"
	"os"
)

func main() {
	arg := "../config/config.json"
	if len(os.Args) > 1 { arg = os.Args[1] }

	//collyclient.Initcollyclient_Agency1()
	irc.StartIRCprocess(arg)

}
