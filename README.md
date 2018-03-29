# IRC WEBSCRAPING BOT

## Steps

1. Clone ```parallel``` repository
2. Go to your $GOPATH 
3. Install dependencies with command ``` go get -v -u github.com/gocolly/colly/... ```
4. Be patient, it takes some time.
5. Install another dependency to connect to PostgreSQL ``` go get -v -u github.com/lib/pq ```
6. Create a config file located in ``` parallel/config/config.json ``` use this sample
```
{ 
	"IRCNick":["YourNick"],
	"IRCChannels":["#Yourchannel"],
	"IRCUser":["user user user user user user"],
	"API":["Your api get address"],
	"IRCServerPort":["irc.ircserver.org:6667"],
	"DBHost":["Hostname of your postgresSQL DB"],
	"DBPort":["Port of your postgresSQL DB"],
	"DBUser":["User of your postgresSQL DB"],
	"DBPass":["Password of your postgresSQL DB"],
	"DBName":["Database Name your postgresSQL DB"]
}
```
7. Run ``` go build ``` located in your main folder
8. Run the generated file ``` ./main ``` 
