package calc

import (
	"go-bot-test/features/calc/actions"
	"go-bot-test/lib/feature"
)

var Feature = feature.Feature{
	Event: event,
	Actions: []feature.Action{
		actions.ReceiveFirstNumAction,
		actions.ReceiveOperatorAction,
		actions.CalcResultAction,
	},
}
