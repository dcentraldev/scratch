package main

import (
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/pipeline"
	"github.com/finiteloopme/goutils/pkg/log"
)

type Config struct {
	Dataset      string  `required:"true"`
	TxTable      string  `default:"FILTERED_TRANSACTIONS"`
	PriceTable   string  `default:"TOKEN_PRICES"`
	HeliusGRPC   string  `required:"true"`
	HeliusAPIKey string  `required:"true"`
	Threshold    float64 `default:"100.0"`
}

var targetTokens = map[string]string{
	"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v": "USDC",
	"So11111111111111111111111111111111111111112":  "SOL", // Wrapped SOL, native SOL handled by lamport transfers
	"JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN":  "JUP",
	"DRIFT6t7MAcxgfaX6B2sFfKCW5x8tUgtr2a3F8T8sZyh": "DRIFT",
}

func processConfig() *Config {
	cfg := Config{}
	// env.ProcessEnvconfig("", &cfg)
	return &cfg
}

func main() {
	log.Info("Hello World")
	processConfig()
	pipeline.Run()
}
