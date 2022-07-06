package deploy

import (
	"go-bot-test/features/deploy/actions"
	"go-bot-test/lib/feature"
)

var Feature = feature.Feature{
	Event: event,
	Actions: []feature.Action{
		actions.SelectVersionAction,
		actions.ConfirmDeploymentAction,
	},
}
