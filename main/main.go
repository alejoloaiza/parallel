package main

import (
	"fmt"
	"os"
	"parallel/data"
	"parallel/db"
	"parallel/irc"
)

func main() {

	arg := "../config/config.json"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	//collyclient.Initcollyclient_Agency1()
	db.DBConnectPostgres(arg)
	db.DBConnectRedis()
	data.FillRawAssetsArray()
	go irc.StartIRCprocess(arg)
	data.PrintAssetsArray()
	data.AssetTypeClassifier()
	data.AssetBusinessClassifier()
	fmt.Println("=======================================")
	data.PrintAssetsArray()
}
