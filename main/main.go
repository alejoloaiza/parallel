package main

import (
	"parallel/irc"
	"os"
	"parallel/db"
)

func main() {
	arg := "../config/config.json"
	if len(os.Args) > 1 { arg = os.Args[1] }

	//collyclient.Initcollyclient_Agency1()
	//db.DBConnectPostgres(arg)
	db.DBConnectRedis()
	irc.StartIRCprocess(arg)
	

}
