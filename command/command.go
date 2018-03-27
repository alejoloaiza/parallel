package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"parallel/config"
	"strings"
)

func ProcessCommand(command []string, allconfig *config.Configuration) string {
	var bodyString string
	fmt.Println("Command request inside process: " + command[0])
	if command[0] == "api" {
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
