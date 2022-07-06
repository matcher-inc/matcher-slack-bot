package deploy

import (
	actions "go-bot-test/features/deploy/actions"
	"go-bot-test/lib/feature"
)

var Feature = feature.Feature{
	Event:   event,
	Actions: actionList,
}

var event = feature.MentionEvent{
	Callback: eventCallback,
}

var actionList = []feature.Action{
	{Key: "selectVersion", Callback: actions.SelectVersionActionCallback},
}
