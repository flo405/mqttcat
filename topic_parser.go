package mqttcat

import (
  MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
  "regexp"
  "strconv"
  "strings"
)

// Get topics
func GetTopics(topics string) ([]MQTT.TopicFilter, error){
  topicsList := strings.Split(topics, ",")
  var result []MQTT.TopicFilter
  re := regexp.MustCompile(`(.*):([0|1|2])`)
  for _,topicAndQos := range topicsList {
    groups := re.FindAllStringSubmatch(topicAndQos, -1)
    if len(groups) ==1 {
      topic := groups[0][1]
      qos, err := strconv.Atoi(groups[0][2])
      if err != nil || qos > 2{
        qos = 0
      }
      filter, err := MQTT.NewTopicFilter(topic, byte(qos))
      if(err != nil){
        return make([]MQTT.TopicFilter, 0), err
      }
      result = append(result, *filter)
    }
  }
  return result, nil
}

