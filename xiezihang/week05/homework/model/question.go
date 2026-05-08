package model

type QuestionRequest struct {
	Model    string `json:"model"`
	Language string `json:"language"`
	Type     int    `json:"type"`
	Keyword  string `json:"keyword"`
}

type QuestionResponse struct {
	Title   string   `json:"title"`
	Answers []string `json:"answers"`
	Right   []int    `json:"right"`
}

type APIResult struct {
	AIStartTime string           `json:"aiStartTime"`
	AIEndTime   string           `json:"aiEndTime"`
	AICostTime  string           `json:"aiCostTime"`
	AIReq       QuestionRequest  `json:"aiReq"`
	AIRes       QuestionResponse `json:"aiRes"`
	HTTPCode    int              `json:"httpCode"`
	HTTPMsg     string           `json:"httpMsg"`
}
