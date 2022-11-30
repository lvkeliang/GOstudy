package model

type SecurityQuestion struct {
	ID       int64  `json:"ID"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
