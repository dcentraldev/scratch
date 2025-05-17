package main

import (
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/config"
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/pipeline"
	"github.com/finiteloopme/goutils/pkg/log"
)

var targetTokens = map[string]string{
	"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v": "USDC",
	"So11111111111111111111111111111111111111112":  "SOL", // Wrapped SOL, native SOL handled by lamport transfers
	"JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN":  "JUP",
	"DRIFT6t7MAcxgfaX6B2sFfKCW5x8tUgtr2a3F8T8sZyh": "DRIFT",
}

func main() {
	log.Info("Hello World")
	cfg := config.GetConfig()
	pipeline.Run(cfg)
}
