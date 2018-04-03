package main

import (
	"parallel/irc"
	"os"
	"parallel/db"
	"parallel/data"
)

func main() {
	
	arg := "../config/config.json"
	if len(os.Args) > 1 { arg = os.Args[1] }

	//collyclient.Initcollyclient_Agency1()
	//db.DBConnectPostgres(arg)
	db.DBConnectRedis()
	go irc.StartIRCprocess(arg)
	data.GetAllKeys()

}
