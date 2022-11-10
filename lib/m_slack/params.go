package mSlack

import "github.com/slack-go/slack"

type RequestParams struct {
	Token string
	// RequestBodyを使うのは、eventのparseで取得して、verificationURLでチェックするときだけ
	// RequestBody []byte
	// Type        EventType
	FeaturePath     string
	ActionPath      string
	UserID          string
	ChannelID       string
	TriggerID       string
	ViewID          string
	ExternalID      string
	ActionParams    ActionParams
	PrivateMetadata string
	responseURL     string
}

type ActionParams struct {
	Value  string
	Values map[string]map[string]slack.BlockAction
}
