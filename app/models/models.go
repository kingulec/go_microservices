package models

type Test struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
}

type Payload struct {
	BuildID   string `json:"build_id"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Tests     []Test `json:"tests"`
}

type Response struct {
	Received bool    `json:"received"`
	BuildId  string  `json:"build_id"`
	PassRate float64 `json:"pass_rate"`
}