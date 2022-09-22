package actions

import "go-bot-test/lib/feature"

var ShowDialogAction = feature.Action{
	Key:      "showDialog",
	Callback: showDialogCallback,
}

var ReceiveFormAction = feature.Action{
	Key:      "receiveForm",
	Callback: receiveFormCallback,
}

var ConfirmAction = feature.Action{
	Key:      "confirm",
	Callback: ConfirmCallback,
}
