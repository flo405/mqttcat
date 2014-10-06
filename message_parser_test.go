package main

import (
 "testing"
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func TestGetMessage(t *testing.T){
  message,err := GetMessage("a:1:EE")
  if err != nil{
    t.FailNow()
  }
  if message.topic != "a"{
    t.FailNow()
  }
  if message.qos != MQTT.QOS_ONE{
    t.FailNow()
  }
}

func TestGetMessageWithInvalidQOS(t *testing.T){
  _,err := GetMessage("a:3:DD")
  if err == nil{
    t.Log("There should be an error")
    t.FailNow()
  }
  message := "Error : invalid message"
  if err.Error() != message {
    t.Logf("Error should be '%s' and got '%s'", message, err.Error)
    t.FailNow()
  }
}
