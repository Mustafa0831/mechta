package data

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	// Создание временного файла для тестирования
	file, err := os.CreateTemp("", "testdata.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Запись тестовых данных в файл
	testData := []Data{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: -5, B: -6},
	}
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}
	if _, err := file.Write(jsonData); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Close()

	// Чтение данных из файла
	dataSlice, err := ReadJSONFile(file.Name())
	if err != nil {
		t.Fatalf("ReadJSONFile returned error: %v", err)
	}

	// Проверка результатов
	if len(dataSlice) != len(testData) {
		t.Errorf("expected %d items, got %d", len(testData), len(dataSlice))
	}
	for i, data := range dataSlice {
		if data != testData[i] {
			t.Errorf("expected %+v, got %+v", testData[i], data)
		}
	}
	fmt.Printf("TestReadJSONFile passed: expected %d items, got %d items\n", len(testData), len(dataSlice))
}
