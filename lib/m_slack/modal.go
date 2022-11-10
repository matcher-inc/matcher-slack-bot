package mSlack

import (
	"github.com/slack-go/slack"
)

type Modal struct {
	Title           string
	CloseButton     string
	SubmitButton    string
	Blocks          Blocks
	CallbackID      string
	ExternalID      string
	PrivateMetadata string
}

func (m Modal) ToViewRequest(params RequestParams) slack.ModalViewRequest {
	modalRequest := slack.ModalViewRequest{
		Type:   slack.ViewType("modal"),
		Title:  Text{Body: m.Title}.toBlockObject(params),
		Close:  Text{Body: m.CloseButton}.toBlockObject(params),
		Submit: Text{Body: m.SubmitButton}.toBlockObject(params),
		Blocks: slack.Blocks{
			BlockSet: m.Blocks.toBlockArray(params),
		},
	}
	modalRequest.CallbackID = m.CallbackID
	modalRequest.ExternalID = m.ExternalID
	modalRequest.PrivateMetadata = m.PrivateMetadata
	return modalRequest
}
