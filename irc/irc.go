package irc

import "net"
import "fmt"
import "bufio"
import "io"
import "strings"
import "time"
import "parallel/command"
import "parallel/config"

func StartIRCprocess() {

	//allconfig := config.GetConfig(configpath)

	for {
		conn, err := net.Dial("tcp", strings.Join(config.Localconfig.IRCServerPort, ""))

		if err != nil {
			fmt.Println(err)
			time.Sleep(2000 * time.Millisecond)
			continue
		}

		fmt.Fprintln(conn, "NICK "+strings.Join(config.Localconfig.IRCNick, ""))
		fmt.Fprintln(conn, "USER "+strings.Join(config.Localconfig.IRCUser, ""))
		fmt.Fprintln(conn, "JOIN "+strings.Join(config.Localconfig.IRCChannels, ""))

		MyReader := bufio.NewReader(conn)
		for {
			message, err := MyReader.ReadString('\n')
			// atomixxx: To handle if connection is closed, and jump to next execution.
			if err != nil {
				fmt.Println(time.Now().Format(time.Stamp) + ">>>" + err.Error())
				if io.EOF == err {
					conn.Close()
					fmt.Println("server closed connection")
				}
				time.Sleep(2000 * time.Millisecond)
				break
			}

			fmt.Print(time.Now().Format(time.Stamp) + ">>" + message)

			// atomixxx: Split the message into words to better compare between different commands
			text := strings.Split(message, " ")
			//fmt.Println("Number of objects in text: "+ strconv.Itoa(len(text)))
			var respond bool = false
			var response string
			// atomixxx: Logic to detect messages, BOT logic should go inside this
			if len(text) >= 4 && text[1] == "PRIVMSG" {
				respond = true
				var repeat bool = true
				var respondTo string
				//atomixxx logic to differ if message is channel or private from user
				if text[2][0:1] == "#" {
					fmt.Println("Message detected from channel")
					// logic to respond the same thing to a channel / repeater BOT
					respondTo = text[2]
				} else {
					fmt.Println("Message detected from user")
					userto := strings.Split(text[0], "!")
					respondTo = userto[0][1:]
					// logic to respond the same thing to a user / repeater BOT
				}
				// If its a command BOT will execute the command given
				if text[3] == ":!command:" {
					repeat = false
					commandresponse := command.ProcessCommand(text[4:])
					response = "PRIVMSG " + respondTo + " :" + commandresponse

				}
				// If is not a command BOT will repeat the same thing
				if repeat == true {
					response = "PRIVMSG " + respondTo + " " + strings.Join(text[3:], " ")

				}
			}
			// atomixxx: Ping/Pong handler to avoid timeout disconnect from the irc server
			if len(text) == 2 && text[0] == "PING" {
				response = "PONG " + text[1]
				respond = true
			}
			// This checks if the received text requires response or not, and respond according to the above logic
			if respond == true {
				fmt.Fprintln(conn, response)
				fmt.Println(time.Now().Format(time.Stamp) + "<<" + response)
			}
		}
		// atomixxx: If connection is closed, will try to reconnect after 2 seconds
		time.Sleep(2000 * time.Millisecond)
	}

}
