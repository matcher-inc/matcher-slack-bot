package mSlack

type RequestParams struct {
	Token string
	// RequestBodyを使うのは、eventのparseで取得して、verificationURLでチェックするときだけ
	// RequestBody []byte
	// Type        EventType
	FeaturePath string
	UserID      string
	ChannelID   string
}
