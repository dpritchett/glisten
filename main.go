package main

import (
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"strings"
)

const (
	VERSION = "glisten v1.0.1"
)

type Message struct {
	Event   *irc.Event
	Channel string
	Nick    string
	Message string
}

func NewMessage(e *irc.Event) *Message {
	m := &Message{}
	m.Event = e

	m.Channel = e.Arguments[0]
	m.Nick = e.Nick
	m.Message = e.Message()

	return m
}

func (m *Message) ToString() string {
	return m.Channel + " <" + m.Nick + "> " + m.Message
}

// Dump an irc message to the log
func PrintEvent(e *irc.Event) {
	msg := NewMessage(e)
	log.Println(msg.ToString())
}

var progArgs = []string{}

func main() {
	progArgs = os.Args[1:]

	if len(progArgs) != 4 {
		log.Fatal("Usage example: glisten \"irc.freenode.net:6667\" glisten_botnick glisten_botname \"#memtechbot\"")
	}

	network, nick, username, channel := progArgs[0], progArgs[1], progArgs[2], progArgs[3]

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

	// Announce successful JOIN
	cnxn.AddCallback("366", func(e *irc.Event) {
		log.Println("** JOINed " + e.Arguments[1])
	})

	// Die when KICKed
	// Arguments:[]string{"#MemphisRuby", "glisten_botnick", "test kick bot"}
	cnxn.AddCallback("KICK", func(e *irc.Event) {
		channel, user := e.Arguments[0], e.Arguments[1]
		msg := "** User " + user + " was kicked from " + channel + "."

		if user == progArgs[1] {
			log.Fatal(msg)
		} else {
			log.Println(msg)
		}
	})

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
