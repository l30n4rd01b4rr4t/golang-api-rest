package clients

import (
	"context"
	"fmt"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/domain"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/rest_client"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/utils"
)

func GetItem(ctx context.Context, itemId string, version string) (*domain.Item, error) {
	item := new(domain.Item)
	client := rest_client.NewClient()
	url := fmt.Sprintf("/items/%s?%s", itemId, utils.GetScope(version))
	client.Get(url)
	response := client.Response(item)
	return item, response
}
