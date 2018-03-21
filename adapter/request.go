package adapter

import "strings"

type Request struct{
	Device string
	Query map[string]string
}

var spType = []string{"iPhone","iPad","Android","iPod"}

func SetDevice(agent string)string{
	for _,v := range spType{
		if strings.Index(agent,v) != -1{
			return "sp"
		}
	}
	return "pc"
}