package main

import "parallel/irc"
import "os"

func main() {
	arg := os.Args[1]
	irc.StartIRCprocess(arg);


}