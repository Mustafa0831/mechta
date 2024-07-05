package data

type SumProcessor struct{}

func (sp SumProcessor) Process(data []Data) int {
	totalSum := 0
	for _, item := range data {
		totalSum += item.A + item.B
	}
	return totalSum
}
