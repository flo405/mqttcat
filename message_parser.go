package main

import (
  "regexp"
  "encoding/hex"
  "errors"
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)


//Create a message from a string by parsing
func GetMessage(mess string) (*Message, error) {
  re := regexp.MustCompile(`(.*:[0|1|2]):(.*)`)
  groups := re.FindAllStringSubmatch(mess, -1)
  if len(groups) == 1 {
    p := new(Message)
    topic, qos, err := GetTopic(groups[0][1])
    if err != nil {
      return nil, errors.New("Error : invalid topic")
    }
    p.topic = topic
    p.qos = qos
    data, err := hex.DecodeString(groups[0][2])
    if err == nil{
      p.data = data
    } else{
      return nil, err
    }
    return p, nil
  }
  return nil, errors.New("Error : invalid message")
}

//message definition
type Message struct{
  topic string
  qos MQTT.QoS
  data []byte
}

