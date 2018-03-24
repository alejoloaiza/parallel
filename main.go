package main

import "net"
import "fmt"
import "bufio"
import "io"
import "strings"
import "time"
//import "strconv"

func main() {

	for {
  		conn, err := net.Dial("tcp", "chat.freenode.net:6667")

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Fprintf(conn, "NICK atomixxxbot"+"\n")
		fmt.Fprintf(conn, "USER golang  8 * :golang ircbot"+"\n")
	    fmt.Fprintf(conn, "JOIN #ALEJOCHAN"+"\n")
	
		MyReader := bufio.NewReader(conn)
	  	for { 
	    	message, err := MyReader.ReadString('\n')
			// atomixxx: To handle if connection is closed, and terminate the program.
			if err != nil {
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
				// logic to differ if message is channel or private from user
				if text[2][0:1] == "#" {
					fmt.Println("Message detected from channel")
					// logic to respond the same thing to a channel / repeater BOT
					response = "PRIVMSG " + text[2] + " " + strings.Join(text[3:]," ") +"\n"
				} else {
					fmt.Println("Message detected for user")
					userto := strings.Split(text[0],"!")
					// logic to respond the same thing to a user / repeater BOT
					response = "PRIVMSG " + userto[0][1:] + " " + strings.Join(text[3:]," ") +"\n"
				}
				fmt.Fprintf(conn,response)
				fmt.Println("<<"+response)
				
			}
			// atomixxx: Ping/Pong handler to avoid timeout from the irc server
			if len(text) == 2 && text[0] == "PING"  {
				pong := "PONG "+text[1]+"\n"
				fmt.Fprintf(conn, pong)
				fmt.Println("<<"+pong)
			}
			
	  	}
		// atomixxx: If connection is closed, will try to reconnect after 2 seconds
		time.Sleep(2000 * time.Millisecond)
	}
}
