package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"parallel/config"
	"strings"
	"parallel/collyclient"
	"parallel/data"

)

func ProcessCommand(command []string, allconfig *config.Configuration) string {
	var bodyString string
	fmt.Println("Command request inside process: " + command[0])
	if (strings.TrimSpace(command[0]) == "process") {
		data.GetAllKeys()
	}
	if (strings.TrimSpace(command[0]) == "webscraping") {
		if (strings.TrimSpace(command[1]) == "agency1") {
			go collyclient.Initcollyclient_Agency1()
		}else if (strings.TrimSpace(command[1]) == "agency2") {
			go collyclient.Initcollyclient_Agency2()
		}
		bodyString = "Executed successfully in background as gorutine"
	}
	if strings.TrimSpace(command[0]) == "api" {
		req, err := http.NewRequest("GET", strings.Join(allconfig.API, ""), nil)
		if err != nil {
			fmt.Println("Error in newRequest: ", err)

		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error in Response: ", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString = string(bodyBytes)
			fmt.Println(bodyString)
		}
	}

	return bodyString
}
