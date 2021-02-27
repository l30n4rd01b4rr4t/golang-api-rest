package services

import (
	"context"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/builders"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/clients"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/domain"
)

func GetFormInit(ctx context.Context, siteId string, itemId string, version string) (*domain.FormInit, error) {

	itemMatchData, err_match := clients.ItemMatch(ctx, siteId, itemId, version)
	if err_match != nil {
		return nil, err_match
	}

	item, err_item := clients.GetItem(ctx, itemId, version)
	if err_item != nil {
		return nil, err_item
	}

	user, err_user := clients.GetUser(ctx, item.SellerId, version)
	if err_user != nil {
		return nil, err_user
	}

	return builders.BuildFormInitData(item, itemMatchData, user), nil
}
