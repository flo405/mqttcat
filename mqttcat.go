package mqttcat

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
  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
}

var(
  url string = "tcp://localhost:1883"
)

// mqttcat --id=test --topics:/topic/a:0,/topic/b:1 --protocol=tcp localhost 1883
var (
  id = kingpin.Flag("id", "The id of the connected device").Default("mqtt-cat").String()
  protocol = kingpin.Flag("protocol", "The protocol to use (tcp, ssl, ws)").Default("tcp").String()
  host = kingpin.Arg("host", "IP address of the host").Required().String()
  port = kingpin.Arg("port", "Port of the server").Required().Int()
  topics = kingpin.Flag("topics", "List of topics").Default("").String()

)

func main() {

  //parse args
  kingpin.Parse()
  topicsList, err := GetTopics(*topics)
  url := fmt.Sprintf("%s://%s:%d",*protocol,*host,*port)

  //create client option
  opts := MQTT.NewClientOptions().AddBroker(url)
  opts.SetClientId(*id)
  opts.SetDefaultPublishHandler(f)

  // create the client
  c := MQTT.NewClient(opts)
  _, startError := c.Start()
  if startError != nil {
    panic(startError)
  }

  for _, topic := range topicsList {
    _, err = c.StartSubscription(nil, &topic)
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
      c.Publish(MQTT.QOS_ONE, message.topic, message.data)
    }
  }

}

