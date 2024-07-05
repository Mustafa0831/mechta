package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

type DataProcessor interface {
	Process(data []Data) int
}

func ReadJSONFile(filePath string) ([]Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var dataSlice []Data
	if err := json.Unmarshal(bytes, &dataSlice); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return dataSlice, nil
}
