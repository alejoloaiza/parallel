package main

import (
	"os"
	"parallel/config"
	"parallel/db"
	"parallel/irc"
)

func main() {

	arg := "../config/config.json"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	_ = config.GetConfig(arg)
	//collyclient.Initcollyclient_Agency1()
	db.DBConnectPostgres(arg)
	db.DBConnectRedis()
	//data.FillRawAssetsArray()
	irc.StartIRCprocess()
	//data.PrintAssetsArray()
	/*
		data.AssetClassifier()
		fmt.Println("=======================================")
		data.PrintAssetsArray()
	*/

}
