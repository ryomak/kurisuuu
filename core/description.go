package core

import (
	"io/ioutil"
)

func GetDescription(filePath string) (string,error){
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "",err
	}
	return string(data), nil
}
