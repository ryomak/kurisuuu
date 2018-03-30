package core

import (
	"net/http"
	"time"
	"encoding/json"
)

var client = &http.Client{Timeout: 3 * time.Second}

func GetJsonByUrl(url string,data interface{})error{
	res,err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return  json.NewDecoder(res.Body).Decode(data)
}

