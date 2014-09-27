package mqttcat


import (

  "regexp"
  "encoding/hex"
  "strconv"
  "errors"
)
//Create a message from a string by parsing
func GetMessage(mess string) (*Message, error) {
  re := regexp.MustCompile(`(.*):([0|1|2]):(.*)`)
  groups := re.FindAllStringSubmatch(mess, -1)
  if len(groups) == 1 {
    p := new(Message)
    p.topic = groups[0][1]
    qos, err := strconv.Atoi(groups[0][2])
    if err == nil || qos > 2 {
      p.qos = qos
    } else{
      return nil, errors.New("Error: invalid qos")
    }
    data, err := hex.DecodeString(groups[0][3])
    if err == nil{
      p.data = data
    } else{
      return nil, err
    }
    return p, nil
  }
  return nil, errors.New("Error : cannot parse message : "+ mess)
}

//message definition
type Message struct{
  topic string
  qos int
  data []byte
}
