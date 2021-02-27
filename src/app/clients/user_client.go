package clients

import (
	"context"
	"fmt"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/domain"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/rest_client"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/utils"
)

func GetUser(ctx context.Context, userId int, version string) (*domain.User, error) {
	user := new(domain.User)
	client := rest_client.NewClient()
	url := fmt.Sprintf("/users/%v?%s", userId, utils.GetScope(version))
	client.Get(url)
	response := client.Response(user)
	return user, response
}
