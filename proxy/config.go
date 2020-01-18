package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ResponseConfig struct {
	RemoteHost    string  `json:"remote"`
	LocalAddress  string  `json:"local"`
	Per404        float32 `json:"404"`
	Per500        float32 `json:"500"`
	PerCustom     float32 `json:"custom"`
	CustomMessage string  `json:"customMessage"`
}

func (rc *ResponseConfig) CheckConfig() bool {
	return (rc.Per404 + rc.Per500 + rc.PerCustom) <= 1.0
}

func LoadConfig(fileName string) ResponseConfig {
	jsonFile, err := os.Open(fileName)
	errorPrintl(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config ResponseConfig
	json.Unmarshal(byteValue, &config)
	return config
}

func errorPrintl(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
