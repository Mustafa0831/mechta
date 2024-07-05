package data

import (
	"fmt"
	"testing"
)

func TestSumProcessor(t *testing.T) {
	data := []Data{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: -5, B: -6},
	}

	processor := SumProcessor{}
	sum := processor.Process(data)

	expectedSum := -1
	if sum != expectedSum {
		t.Errorf("expected %d, got %d", expectedSum, sum)
	} else {
		fmt.Printf("TestSumProcessor passed: expected %d, got %d\n", expectedSum, sum)
	}
}

func BenchmarkSumProcessor(b *testing.B) {
	// Создаем тестовые данные
	data := make([]Data, 1000000)
	for i := range data {
		data[i] = Data{A: i % 10, B: -(i % 10)}
	}

	processor := SumProcessor{}

	b.ResetTimer() // Сбрасываем таймер перед началом бенчмарка

	for i := 0; i < b.N; i++ {
		processor.Process(data)
	}
}
