Gossip
======

Simple IRC client written in Go.

How to use it
--------------
Build with
~~~sh
go build
~~~

then run
~~~sh
./gossip
~~~

you might want to edit pr1.json to whatever you want:
~~~json
{
"Server": "irc.network.net",
"Port": "6667",
"Nick": "yournick",
"User": "youruser",
"Channel": "#somechannel"
}
~~~

It doesn't support identification yet

TODO
-------
* Switch from JSON to YAML for config
* Have multiple profiles you can load (no hardcoded)
* Use channels and concurrency features
* Clean code in general (decodeJSON for example)
* Delete ugly hack when asking for input
* Add support for identification
* Curated output (Colorized names, etc.)
* Curses interface (Separate output/input)