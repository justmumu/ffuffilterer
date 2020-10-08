package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ffuf/ffuf/pkg/output"
)

type Record struct {
	Results []output.Result `json:"results"`
}

func LoadResults(path string) (*Record, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var record Record
	json.Unmarshal(byteValue, &record)

	return &record, nil
}
