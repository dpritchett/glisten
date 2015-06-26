package main

import (
	"github.com/thoj/go-ircevent"
	"log"
)

func main() {
	// connect
	irccon1 := irc.IRC("glisten-tester", "glisten-tester")
	irccon1.VerboseCallbackHandler = true
	irccon1.Debug = true

	err := irccon1.Connect("irc.freenode.net:6667")

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Can't connect to freenode.")
	}

	// 001: Welcome to the network [https://tools.ietf.org/html/rfc2812]
	irccon1.AddCallback("001", func(e *irc.Event) { irccon1.Join("#memtechbot") })

	irccon1.AddCallback("PRIVMSG", func(e *irc.Event) { log.Println("<" + e.Nick + "> " + e.Message()) })

	irccon1.Loop()
}
