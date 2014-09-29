# Install

The app is not available in any package manager.

~~~bash
go get
go build
~~~

# Feature

Mqttcat sends to mqtt broker formatted input from stdin
that should looks like topic:qos:hex_message.

~~~bash
usage: mqttcat [<flags>] <host> <port>

A netcat for mqtt.

Flags:
  --help               Show help.
  --id="mqtt-cat"      The id of the connected device
  --protocol="tcp"     The protocol to use (tcp, ssl, ws)
  --topics=TOPICS      List of topics
  --username=USERNAME  MQTT server username
  --password=PASSWORD  MQTT server password
  --clean              Set clean session flag
  --keepalive=10       Set the keep alive timer (in sec)

Args:
  <host>  IP address of the host
  <port>  Port of the server

~~~

#Usage

Cat a file

cat data.dat |  mqttcat localhost 1883

Read from stdin

~~~bash
./mqttcat localhost 1883 << EOF
a:1:48
a:1:46
b:1:45
EOF
~~~
