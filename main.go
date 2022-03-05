package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/migueleliasweb/prom-metrics-trimmer/pkg/prom"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func main1() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, "up", time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func main() {
	writer := bytes.NewBufferString("your string")

	fmt.Println(prom.GenerateBackfillData(writer))

	fmt.Println(writer.String())
}
