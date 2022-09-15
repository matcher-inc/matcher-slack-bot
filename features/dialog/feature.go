package dialog

import (
	"go-bot-test/features/dialog/actions"
	"go-bot-test/lib/feature"
)

var Feature = feature.Feature{
	Event: event,
	Actions: []feature.Action{
		actions.ShowDialogAction,
		actions.ReceiveFormAction,
		actions.ConfirmAction,
	},
}
