package builders

import (
	"strconv"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/domain"
)

func BuildFormInitData(item *domain.Item, itemMatchData *domain.ItemMatch, user *domain.User) *domain.FormInit {
	var formInitData domain.FormInit

	model_attr := GetItemAttribute(item, "MODEL")
	brand_attr := GetItemAttribute(item, "BRAND")
	year_attr := GetItemAttribute(item, "VEHICLE_YEAR")
	vehicle_type_attr := GetItemAttribute(item, "VEHICLE_TYPE")

	formInitData.Vehicle.ModelID = model_attr.ValueID
	formInitData.Vehicle.BrandID = brand_attr.ValueID
	formInitData.Vehicle.FullYearID = year_attr.ValueID
	formInitData.Vehicle.VehicleTypeID, _ = strconv.ParseInt(vehicle_type_attr.ValueID, 10, 64)
	formInitData.Vehicle.PurchaseValue = item.Price
	formInitData.Vehicle.Attributes = itemMatchData.Attributes
	formInitData.Customer.Nickname = user.Nickname
	formInitData.Dealership.UserType = user.UserType

	return &formInitData
}

func GetItemAttribute(item *domain.Item, id string) domain.Attributes {
	attribute, _ := underscore.FindBy(item.Attributes, map[string]interface{}{
		"id": id,
	}).(domain.Attributes)
	return attribute
}
