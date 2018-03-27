package config

import "encoding/json"
import "os"
import "fmt"

type Configuration struct {
	Nick     []string
	Channels []string
	User     []string
	API      []string
	ServerPort      []string
}

func GetConfig(configpath string) *Configuration {
	file, _ := os.Open(configpath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}
