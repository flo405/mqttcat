package main

import (
 "testing"
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)


func TestGetTopics(t *testing.T){
  topics, err := GetTopics("a:1,b:2,c:0")
  if err != nil{
    t.Log(err.Error())
    t.Log("There should be no error")
    t.FailNow()
  }
  length := len(topics)
  if length != 3 {
    t.Logf("Length should be 3 and it's '%s'", length)
    t.FailNow()

  }
}

func TestGetTopicsiWithInvalidList(t *testing.T){
  _, err := GetTopics("invalid data")
  if err == nil{
    t.Log("There should be an error")
    t.FailNow()
  }
}

func TestGetTopic(t *testing.T){
  topic,qos,err := GetTopic("a:1")
  if err != nil{
    t.Log("There should be no error")
    t.Fail()
  }
  if topic != "a" {
    t.Log("Topic should equal a")
    t.Fail()
  }
  if qos != MQTT.QOS_ONE{
    t.Log("Qos should be QOS_ONE")
    t.Fail()
  }

}

func TestGetTopicWithInvalidQOS(t *testing.T){
  _,_,err := GetTopic("a:3")
  if err == nil{
    t.Log("There should be an error")
    t.Fail()
  }
  message := "Error : invalid topic"
  if err.Error() != message {
    t.Logf("Error should be 'message' and got '%s'", message, err.Error())
    t.FailNow()
  }
}

func TestGetQOS(t *testing.T){
  qos,_ := GetQOS(1)
  if qos != MQTT.QOS_ONE{
    t.Logf("QOS should be QOS_ONE")
    t.FailNow()
  }

  qos,_ = GetQOS(2)
  if qos != MQTT.QOS_TWO{
    t.Logf("QOS should be QOS_TWO")
    t.FailNow()
  }

  qos,_ = GetQOS(0)
  if qos != MQTT.QOS_ZERO{
    t.Logf("QOS should be QOS_ZERO")
    t.FailNow()
  }
}

func TestGetQOSWithInvalidQOS(t *testing.T){
  _, err := GetQOS(3)
  if err == nil{
    t.Log("There should be an error")
    t.FailNow()
  }
  message := "Error : invalid qos"
  if err.Error() != message{
    t.Logf("Error should eq '%s' and got '%s'", message, err.Error())
    t.FailNow()
  }

}

