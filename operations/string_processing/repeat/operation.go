package repeat

import (
	"strings"

	"github.com/project-flogo/catalystml-flogo/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	// params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Info("Starting Operation Repeat")
	a.logger.Debugf("Input for operation repeat.%s %d times", input.S, input.I)

	out := strings.Repeat(input.S, input.I)

	a.logger.Info("Operation Repeat completed.")
	a.logger.Debug("Output for operation repeat.", out)

	return out, nil
}
