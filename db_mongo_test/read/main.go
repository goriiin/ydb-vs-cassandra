package main

import (
	"context"
	"db-bench/lib"
	"db-bench/lib/conf"
	"log"
	"os"
	"sync"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	cfg, err := conf.LoadConfig("mongo", configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), cfg.TestDuration)
	defer cancel()
	tester, err := lib.GetTester("mongo", cfg)
	if err != nil {
		log.Fatalf("Failed to initialize mongo tester: %v", err)
	}
	defer tester.Close()
	var wg sync.WaitGroup
	tester.RunTest(ctx, &wg)
	wg.Wait()
	log.Println("Test for mongo completed.")
}
