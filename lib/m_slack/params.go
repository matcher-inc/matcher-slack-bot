package mSlack

type EventParams struct {
	Token       string
	RequestBody []byte
	Type        EventType
	RequestKey  string
	UserID      string
	ChannelID   string
}
