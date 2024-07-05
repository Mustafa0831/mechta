package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"mechta/pkg/data"
)

func generateTestData(filePath string, numObjects int) error {
	dataSlice := make([]data.Data, numObjects)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numObjects; i++ {
		dataSlice[i] = data.Data{
			A: rand.Intn(21) - 10,
			B: rand.Intn(21) - 10,
		}
	}

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(dataSlice); err != nil {
		return err
	}

	return nil
}

func benchmarkCalculateSum(b *testing.B, numGoroutines int) {
	filePath := filepath.Join(".", "data", "data.json")
	if err := generateTestData(filePath, 1000000); err != nil {
		b.Fatalf("Error generating test data: %v", err)
	}

	app := &App{
		numGoroutines: numGoroutines,
		dataFile:      filePath,
		processor:     data.SumProcessor{},
	}
	dataSlice, err := data.ReadJSONFile(app.dataFile)
	if err != nil {
		b.Fatalf("Error reading JSON file: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.calculateSum(ctx, dataSlice)
	}
}

func BenchmarkCalculateSum1(b *testing.B)  { benchmarkCalculateSum(b, 1) }
func BenchmarkCalculateSum2(b *testing.B)  { benchmarkCalculateSum(b, 2) }
func BenchmarkCalculateSum4(b *testing.B)  { benchmarkCalculateSum(b, 4) }
func BenchmarkCalculateSum8(b *testing.B)  { benchmarkCalculateSum(b, 8) }
func BenchmarkCalculateSum16(b *testing.B) { benchmarkCalculateSum(b, 16) }
func BenchmarkCalculateSum32(b *testing.B) { benchmarkCalculateSum(b, 32) }
func BenchmarkCalculateSum64(b *testing.B) { benchmarkCalculateSum(b, 64) }
