# IRC WEBSCRAPING BOT

## Steps

1. Clone ```parallel``` repository
2. Go to your $GOPATH 
3. Install the following dependencies 
``` 
go get -v -u github.com/gocolly/colly/... 
go get -v -u "github.com/schollz/closestmatch"
go get -v -u "github.com/lib/pq"
go get -v -u "github.com/go-redis/redis"
```
4. Be patient, it takes some time.
5. Create a config file located in ``` parallel/config/config.json ``` use this sample
```
{ 
	"IRCNick":["YourNick"],
	"IRCChannels":["#Yourchannel"],
	"IRCUser":["user user user user user user"],
	"GoogleAPI":["Your api get address"],
	"IRCServerPort":["irc.ircserver.org:6667"],
	"DBHost":["Hostname of your postgresSQL DB"],
	"DBPort":["Port of your postgresSQL DB"],
	"DBUser":["User of your postgresSQL DB"],
	"DBPass":["Password of your postgresSQL DB"],
	"DBName":["Database Name your postgresSQL DB"]
}
```
6. Database setup: follow steps at  [Database Setup](DATABASE.md)
7. Run ``` go build ``` located in your main folder
8. Run the generated file ``` ./main ``` 
