package pipeline

import (
	"context"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
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

	//beam.ParDo0(s.Scope("LogInfo"), Temp, impulse)
	solClientFn, err := transform.NewChainClient(
		cfg.HeliusGRPC,
		cfg.HeliusAPIKey)
	if err != nil {
		log.Fatalf(ctx, "Failed to create Solana client: %v", err)
	}
	beam.ParDo(s.Scope("SubscribeSol"), solClientFn, impulse)

	beamx.Run(context.Background(), pipeline)
}

func Temp(c []uint8) {
	log.Info(context.Background(), "Hello world in beam")

}

func init() {
	register.Function1x0(Temp)
}
