package main

import (
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
	irc.StartIRCprocess(arg)
	//data.PrintAssetsArray()
	/*
		data.AssetClassifier()
		fmt.Println("=======================================")
		data.PrintAssetsArray()
	*/
}
