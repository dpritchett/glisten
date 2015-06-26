package main

import (
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"strings"
)

// Dump an irc message to the log
func PrintEvent(e *irc.Event) {
	channel := e.Arguments[0]
	nick := e.Nick
	msg := e.Message()
	log.Println(channel + " <" + nick + "> " + msg)
}

func main() {
	if len(os.Args) != 5 {
		log.Fatal("Usage example: glisten \"irc.freenode.net:6667\" glisten_botnick glisten_botname \"#memtechbot\"")
	}

	network, nick, username, channel := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// connect
	cnxn := irc.IRC(nick, username)

	debugEnabled := strings.ToLower(strings.Trim(os.Getenv("DEBUG"), " ")) == "true"

	// pretty noisy, probably need to squelch it at some point
	if debugEnabled {
		cnxn.VerboseCallbackHandler = true
		cnxn.Debug = true
	} else {
		cnxn.Debug = false
		cnxn.VerboseCallbackHandler = false
	}

	// join specified channel on successful network connect
	cnxn.AddCallback("001", func(e *irc.Event) { cnxn.Join(channel) })

	err := cnxn.Connect(network)

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Can't connect to " + network + ".")
	}

	// Despite the name, PRIVMSG seems to pass on non-private messages too
	cnxn.AddCallback("PRIVMSG", PrintEvent)

	// run until killed
	cnxn.Loop()
}
