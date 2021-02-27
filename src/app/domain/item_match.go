package domain

type ItemMatch struct {
	Score      float32      `json:"score"`
	Attributes []Attributes `json:"attributes"`
}
