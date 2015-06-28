# ✨ glisten ✨

## what

IRC listener that streams a single channel to STDOUT

## why

So you can log public IRC channels to STDOUT using a daemonized oneliner

## how

### usage
![usage demo screenshot](http://i.imgur.com/rcJhBWp.png)

`$ glisten "irc.freenode.net:6667" bot_nick_here bot_name_here "#channel_name"`

#### Running as a CoreOS / `systemd` daemon

- Cut and paste the file from `examples/glistenMRUG.service` over to your CoreOS cluster
- Update the IRC connection info to match your desired network, channel, and username
- `fleetctl submit glistenMRUG.service`
- `fleetctl start  glistenMRUG.service`
- `fleetctl journal -f glistenMRUG.service`

### install

`go get -u github.com/dpritchett/glisten`

### prebuilt downloads for Linux, OS X, and Windows

* [https://github.com/dpritchett/glisten/releases](https://github.com/dpritchett/glisten/releases)

## contributors

* Daniel Pritchett <dpritchett@gmail.com>
