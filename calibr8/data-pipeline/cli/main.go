package main

import (
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/config"
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/pipeline"
)

func main() {
	cfg := config.GetConfig()
	pipeline.Run(cfg)
}
