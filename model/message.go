package model

type Message struct {
	MID       int64  `json:"mid"`
	SenderUID int64  `json:"senderUID"`
	RecUID    int64  `json:"recUID"`
	Detail    string `json:"detail"`
}
