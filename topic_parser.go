package mqttcat

import (
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
  "regexp"
  "strconv"
  "strings"
  "errors"
)

//Extract a list of topic from a string
//input example : a:1,b:0,c:2
func GetTopics(topics string) ([]MQTT.TopicFilter, error){
  topicsList := strings.Split(topics, ",")
  var result []MQTT.TopicFilter
  for _,topicAndQos := range topicsList {
    topic, qos, err := GetTopic(topicAndQos)
    if(err == nil){
      filter, _ := MQTT.NewTopicFilter(topic, byte(qos))
      result = append(result, *filter)
    }
  }
  return result, nil
}

//Extract a topic from a string
//input example : a:1
func GetTopic(data string) (topic string, qos MQTT.QoS, err error){
  re := regexp.MustCompile(`(.*):([0|1|2])`)
  groups := re.FindAllStringSubmatch(data, -1)
  if len(groups) ==1 {
    topic := groups[0][1]
    qosInt, err := strconv.Atoi(groups[0][2])
    if err != nil || qos > 2{
      qos = 0
    }
    qos, err = getQOS(qosInt)
    return topic, qos, nil
  }
  return "",MQTT.QOS_ZERO, errors.New("Invalid topic : "+data)

}
func getQOS(qos int) (MQTT.QoS, error){
  switch qos {
    case 0: return  MQTT.QOS_ZERO, nil
    case 1: return MQTT.QOS_ONE, nil
    case 2: return MQTT.QOS_TWO, nil
    default: return MQTT.QOS_ZERO, errors.New("Invalid qos")
  }
}
