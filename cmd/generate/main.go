package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"mechta/pkg/data"
)

type DataGenerator struct {
	numObjects int
	filePath   string
}

const numObjects = 1000000

func NewDataGenerator(numObjects int, filePath string) *DataGenerator {
	return &DataGenerator{
		numObjects: numObjects,
		filePath:   filePath,
	}
}

func (dg *DataGenerator) Generate() {
	dataSlice := make([]data.Data, dg.numObjects)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dg.numObjects; i++ {
		dataSlice[i] = data.Data{
			A: rand.Intn(21) - 10,
			B: rand.Intn(21) - 10,
		}
	}

	err := os.MkdirAll(filepath.Dir(dg.filePath), os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(dg.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(dataSlice); err != nil {
		panic(err)
	}

	println("data.json file has been generated successfully")
}

func main() {
	generator := NewDataGenerator(numObjects, "data/data.json")
	generator.Generate()
}
