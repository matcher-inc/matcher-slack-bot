package actions

import "go-bot-test/lib/feature"

var ShowDialogAction = feature.Action{
	ActionPath: "showDialog",
	Callback:   showDialogCallback,
}

var ReceiveFormAction = feature.Action{
	ActionPath: "receiveForm",
	Callback:   receiveFormCallback,
}

var ConfirmAction = feature.Action{
	ActionPath: "confirm",
	Callback:   ConfirmCallback,
}
