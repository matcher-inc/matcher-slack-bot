package feature

import (
	"strings"

	"github.com/slack-go/slack"
)

type Action struct {
	Key      string
	Callback func(slack.InteractionCallback) error
}

func (f Feature) RunAction(payload slack.InteractionCallback) error {
	for _, action := range f.Actions {
		if actionIsMatchingToRoute(payload, action) {
			return action.Callback(payload)
		}
	}
	return nil
}

func actionIsMatchingToRoute(payload slack.InteractionCallback, action Action) bool {
	switch payload.Type {
	case slack.InteractionTypeBlockActions:
		if len(payload.ActionCallback.BlockActions) == 0 {
			return false
		}
		blockAction := payload.ActionCallback.BlockActions[0]
		path := strings.Split(blockAction.BlockID, ":")[1]
		return path == action.Key
	}
	return false
}
