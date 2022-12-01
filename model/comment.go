package model

type AllComment struct {
	Comment          Message     `json:"comment"`
	NextLayerComment interface{} `json:"nextLayerComment"`
}
