package command

import "net/http"
import "io/ioutil"
import "fmt"

func ProcessCommand(command[] string) string {
	var bodyString string
	fmt.Println("Command request inside process: " +command[0] )
	if command[0] == "api" {
		req, err := http.NewRequest("GET", "http://47.88.174.2:3000/api/transactions/71527525", nil)
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