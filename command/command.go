package command

import (
	"fmt"
	"parallel/collyclient"
	"parallel/data"
	"strings"
)

func ProcessCommand(command []string) string {
	var bodyString string
	fmt.Println("Command request inside process: " + command[0])
	if strings.TrimSpace(command[0]) == "process" {
		api := "on"
		data.FlushData()
		data.FillRawAssetsArray()
		data.PrintAssetsArray()

		if strings.TrimSpace(command[1]) == "apioff" {
			api = "off"
		}
		data.AssetClassifier(api)
		fmt.Println("=======================================")
		data.PrintAssetsArray()
	}
	if strings.TrimSpace(command[0]) == "webscraping" {
		if strings.TrimSpace(command[1]) == "agency1" {
			go collyclient.Initcollyclient_Agency1()
		} else if strings.TrimSpace(command[1]) == "agency2" {
			go collyclient.Initcollyclient_Agency2()
		}
		bodyString = "Executed successfully in background as gorutine"
	}
	return bodyString
}
