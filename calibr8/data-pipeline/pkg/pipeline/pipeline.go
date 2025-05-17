package pipeline

import (
	"context"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

func Run() {
	beam.Init()
	pipeline, scope := beam.NewPipelineWithRoot()
	ConstructPipeline(pipeline, scope)
}

func ConstructPipeline(pipeline *beam.Pipeline, s beam.Scope) {
	s = s.Scope("SolDataPipeline")
	// ctx := context.Background()

	impulse := beam.Impulse(s)

	beam.ParDo0(s.Scope("LogInfo"), Temp, impulse)

	beamx.Run(context.Background(), pipeline)
}

func Temp(c []uint8) {
	log.Info(context.Background(), "Hello world in beam")

}

func init() {
	register.Function1x0(Temp)
}
