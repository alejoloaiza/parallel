package main

import "net"
import "net/http"
import "fmt"
import "bufio"
import "io"
import "io/ioutil"
import "strings"
import "time"
//import "os"
//import "strconv"



func main() {

	/*var attr os.ProcAttr 
	attr.Sys.HideWindow = true
	p, err := os.StartProcess("whatever", nil, &attr)
	*/
	for {
  		conn, err := net.Dial("tcp", "chat.freenode.net:6667")

		if err != nil {
			fmt.Println(err)
			time.Sleep(2000 * time.Millisecond)
			continue
		}
		fmt.Fprintln(conn, "NICK atomixxxbot")
		fmt.Fprintln(conn, "USER golang  8 * :golang ircbot")
		fmt.Fprintln(conn, "JOIN #ALEJOCHAN")
		
		MyReader := bufio.NewReader(conn)
	  	for { 
	    	message, err := MyReader.ReadString('\n')
			// atomixxx: To handle if connection is closed, and terminate the program.
			if err != nil {
				fmt.Println(err)
				if io.EOF == err {
					conn.Close()
					fmt.Println("server closed connection")
					break
				}
			}
			
	    	fmt.Print(">>"+message)
	
			// atomixxx: Split the message into words to better compare between different commands
			text := strings.Split(message," ")
			//fmt.Println("Number of objects in text: "+ strconv.Itoa(len(text)))
			
			// atomixxx: Logic to detect messages, BOT logic should go inside this
			if len(text) >= 4 && text[1] == "PRIVMSG" {
				var response string
				var repeat bool = true
				var respondTo string
				// logic to differ if message is channel or private from user
				if text[2][0:1] == "#" {
					fmt.Println("Message detected from channel")
					// logic to respond the same thing to a channel / repeater BOT
					respondTo =  text[2]
				} else {
					fmt.Println("Message detected from user")
					userto := strings.Split(text[0],"!")
					respondTo = userto[0][1:]
					// logic to respond the same thing to a user / repeater BOT
				}
				if text[3] == ":!command:" {
					repeat = false
					commandresponse := processCommand(text[4:])
					response = "PRIVMSG " + respondTo + " :" + commandresponse 
					fmt.Fprintln(conn,response)
					fmt.Println("<<"+response)
				} 
				if repeat == true {
					response = "PRIVMSG " + respondTo + " " + strings.Join(text[3:]," ") 
					fmt.Fprintln(conn,response)
					fmt.Println("<<"+response)
				}
			}
			// atomixxx: Ping/Pong handler to avoid timeout disconnect from the irc server
			if len(text) == 2 && text[0] == "PING"  {
				pong := "PONG "+text[1]
				fmt.Fprintln(conn, pong)
				fmt.Println("<<"+pong)
			}
			
	  	}
		// atomixxx: If connection is closed, will try to reconnect after 2 seconds
		time.Sleep(2000 * time.Millisecond)
	}

}
func processCommand(command[] string) string {
		var bodyString string
		fmt.Println("Command request inside process: " +command[0] +"|" )
		if command[0] == "api" {
			fmt.Println("api!")
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
