package domain

type Attributes struct {
	ID        string `json:"id"`
	ValueID   string `json:"value_id,omitempty"`
	ValueName string `json:"value_name"`
}
