package actions

import "go-bot-test/lib/feature"

var SelectVersionAction = feature.Action{
	ActionPath: "selectVersion",
	Callback:   selectVersionActionCallback,
}

var ConfirmDeploymentAction = feature.Action{
	ActionPath: "confirmDeployment",
	Callback:   confirmDeploymentActionCallback,
}
