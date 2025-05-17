package pipeline

import (
	"context"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/debug"
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/config"
	"github.com/dcentraldev/scratch/calibr8/data-pipeline/pkg/transform"
)

func Run(cfg *config.Config) {
	beam.Init()
	pipeline, scope := beam.NewPipelineWithRoot()
	ConstructPipeline(pipeline, scope, cfg)
}

func ConstructPipeline(pipeline *beam.Pipeline, s beam.Scope, cfg *config.Config) {
	s = s.Scope("SolDataPipeline")
	ctx := context.Background()

	impulse := beam.Impulse(s)

	solClientFn, err := transform.NewChainClient(cfg)
	if err != nil {
		log.Fatalf(ctx, "Failed to create Solana client: %v", err)
	}
	logs := beam.ParDo(s.Scope("SubscribeSol"), solClientFn, impulse)
	debug.Printf(s.Scope("EventInfo"), "Recevied event: %v", logs)

	beamx.Run(context.Background(), pipeline)
}
