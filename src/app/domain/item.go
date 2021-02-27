package domain

type Item struct {
	Id         string       `json:"id"`
	SellerId   int          `json:"seller_id"`
	BrandId    int          `json:"brand_id"`
	Price      float32      `json:"price"`
	Attributes []Attributes `json:"attributes"`
}
