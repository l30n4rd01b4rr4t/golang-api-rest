package domain

type FormInit struct {
	Vehicle       Vehicle
	Customer      Customer
	Dealership    Dealership
	TransactionId string
}

type Vehicle struct {
	BrandID       string
	ModelID       string
	StateID       string
	FullYearID    string
	PurchaseValue float32
	VehicleTypeID int64
	Attributes    []Attributes
}
type Customer struct {
	Nickname string
}
type Dealership struct {
	UserType string
}
