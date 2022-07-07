package actions

import "go-bot-test/lib/feature"

var SelectVersionAction = feature.Action{
	Key:      "selectVersion",
	Callback: selectVersionActionCallback,
}

var ConfirmDeploymentAction = feature.Action{
	Key:      "confirmDeployment",
	Callback: confirmDeploymentActionCallback,
}
