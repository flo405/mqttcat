package main

import (
  "fmt"
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
  "os"
  "bufio"
  "strings"
  "gopkg.in/alecthomas/kingpin.v1"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
  fmt.Printf("%s:%s\n", msg.Topic(),msg.Payload())
}



// mqttcat --id=test --topics:/topic/a:0,/topic/b:1 --protocol=tcp localhost 1883
var (
  app = kingpin.New("mqttcat", "A netcat for mqtt.")
  id = app.Flag("id", "The id of the connected device").Default("mqtt-cat").String()
  protocol = app.Flag("protocol", "The protocol to use (tcp, ssl, ws)").Default("tcp").String()
  topics = app.Flag("topics", "List of topics : /topic/1:1, /topic/2:0").Default("").String()
  username = app.Flag("username", "MQTT server username").Default("").String()
  password = app.Flag("password", "MQTT server password").Default("").String()
  cleanSession = app.Flag("clean", "Set clean session flag").Default("false").Bool()
  keepAlive = app.Flag("keepalive", "Set the keep alive timer (in sec)").Default("10").Int()
  host = app.Arg("host", "IP address of the host").Required().String()
  port = app.Arg("port", "Port of the server").Required().Int()
)

func main() {

  //parse args
  kingpin.Version("0.1.0")
  app.Parse(os.Args[1:])
  topicsList, err := GetTopics(*topics)
  url := fmt.Sprintf("%s://%s:%d",*protocol,*host,*port)

  //create client option
  opts := MQTT.NewClientOptions().AddBroker(url)
  opts.SetClientId(*id)
  opts.SetDefaultPublishHandler(f)
  opts.SetCleanSession(*cleanSession)
  opts.SetKeepAlive(uint(*keepAlive))

  if *username != ""{
    opts.SetUsername(*username)
  }
  if *password != ""{
    opts.SetPassword(*password)
  }

  // create the client
  c := MQTT.NewClient(opts)
  _, startError := c.Start()
  if startError != nil {
    panic(startError)
  }

  //subscribe to topic list
  for _, topic := range topicsList {
    _, err = c.StartSubscription(f, &topic)
    if err != nil{
      panic(err)
    }
  }

  //read message on stdin, parse and send them
  reader := bufio.NewReader(os.Stdin)
  for{
    text, err := reader.ReadString('\n')
    if err != nil {
       // You may check here if err == io.EOF
       break
     }
    message, err := GetMessage(strings.TrimSuffix(text, "\n"))
    if err == nil {
      fmt.Println(message.topic)
      receipt := c.Publish(message.qos, message.topic, message.data)
      <-  receipt
    }
  }

}

