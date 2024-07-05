package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"mechta/pkg/data"
)

type App struct {
	numGoroutines int
	profile       bool
	dataFile      string
	processor     data.DataProcessor
}

func NewApp() *App {
	var (
		numGoroutines int
		profile       bool
	)

	flag.IntVar(&numGoroutines, "goroutines", 1, "Number of goroutines to use")
	flag.BoolVar(&profile, "profile", false, "Run in profiling mode")
	flag.Parse()

	return &App{
		numGoroutines: numGoroutines,
		profile:       profile,
		dataFile:      filepath.Join(".", "data", "data.json"),
		processor:     data.SumProcessor{},
	}
}

func (app *App) Run() {
	if app.numGoroutines <= 0 {
		log.Println("Number of goroutines must be greater than 0")
		flag.Usage()
		return
	}

	dataSlice, err := data.ReadJSONFile(app.dataFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	totalSum, duration := app.calculateSum(ctx, dataSlice)

	log.Printf("Total sum: %d", totalSum)
	log.Printf("Time taken: %v", duration)

	if app.profile {
		app.startProfiling()
	}
}

func (app *App) calculateSum(ctx context.Context, dataSlice []data.Data) (int, time.Duration) {
	startTime := time.Now()

	dataLen := len(dataSlice)
	chunkSize := (dataLen + app.numGoroutines - 1) / app.numGoroutines

	var wg sync.WaitGroup
	sumCh := make(chan int, app.numGoroutines)

	for i := 0; i < app.numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > dataLen {
			end = dataLen
		}

		wg.Add(1)
		go app.worker(ctx, dataSlice[start:end], &wg, sumCh)
	}

	wg.Wait()
	close(sumCh)

	totalSum := 0
	for sum := range sumCh {
		totalSum += sum
	}

	duration := time.Since(startTime)
	return totalSum, duration
}

func (app *App) worker(ctx context.Context, data []data.Data, wg *sync.WaitGroup, sumCh chan int) {
	defer wg.Done()
	sum := app.processor.Process(data)
	select {
	case <-ctx.Done():
		log.Println("Worker canceled")
		return
	case sumCh <- sum:
	}
}

func (app *App) startProfiling() {
	log.Println("Profiling mode enabled. Keeping the application running.")
	fmt.Println("Access pprof at: http://localhost:6060/debug/pprof/")
	fmt.Println("Press Ctrl+C to exit")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("Received interrupt signal, shutting down...")
}

func main() {
	app := NewApp()
	app.Run()
}

func init() {
	go func() {
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			log.Fatalf("pprof server failed: %v", err)
		}
	}()
}
