package actions

import "go-bot-test/lib/feature"

var ReceiveFirstNumAction = feature.Action{
	ActionPath: "receiveFirstNum",
	Callback:   receiveFirstNumCallback,
}

var ReceiveOperatorAction = feature.Action{
	ActionPath: "receiveOperator",
	Callback:   receiveOperatorCallback,
}

var CalcResultAction = feature.Action{
	ActionPath: "calcResult",
	Callback:   calcResultCallback,
}
