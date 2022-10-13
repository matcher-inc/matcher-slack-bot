package mSlack

type RequestParams struct {
	Token string
	// RequestBodyを使うのは、eventのparseで取得して、verificationURLでチェックするときだけ
	RequestBody []byte
	Type        EventType
	RequestKey  string
	UserID      string
	ChannelID   string
}
