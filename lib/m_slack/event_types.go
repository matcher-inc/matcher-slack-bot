package mSlack

type EventType string

const (
	AppMentionEvent EventType = "AppMentionEvent"
	SlashEvent      EventType = "SlashEvent"
	URLVerification EventType = "URLVerification"
)
