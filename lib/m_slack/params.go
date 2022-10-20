package mSlack

type RequestParams struct {
	Token string
	// RequestBodyを使うのは、eventのparseで取得して、verificationURLでチェックするときだけ
	// RequestBody []byte
	// Type        EventType
	FeaturePath  string
	ActionPath   string
	UserID       string
	ChannelID    string
	TriggerID    string
	ActionParams ActionParams
}

type ActionParams struct {
	Value  string
	Values map[string]string
}
