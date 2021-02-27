package clients

import (
	"context"
	"errors"
	"fmt"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/domain"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/rest_client"
	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/utils"
)

func ItemMatch(ctx context.Context, siteId string, itemId string, version string) (*domain.ItemMatch, error) {
	santanderMatch, itauMatch := ItemMatchByBank(ctx, siteId, "BANCO1", itemId, version), ItemMatchByBank(ctx, siteId, "BANCO2", itemId, version)
	dataSantanderMatch, dataItauMatch := <-santanderMatch, <-itauMatch

	if dataSantanderMatch.Attributes == nil && dataItauMatch.Attributes == nil {

		return nil, errors.New("Item no match")
	}

	return getNotEmptyMatchData(dataSantanderMatch, dataItauMatch), nil
}

func ItemMatchByBank(ctx context.Context, siteId string, bankId string, itemId string, version string) <-chan *domain.ItemMatch {
	r := make(chan *domain.ItemMatch)

	go func() {
		defer close(r)

		itemMatch := new(domain.ItemMatch)
		client := rest_client.NewClient()

		url := fmt.Sprintf("/match/%s/banks/%s/items/%s?%s", siteId, bankId, itemId, utils.GetScope(version))
		client.Get(url)

		if client.Response(itemMatch) != nil {
			r <- &domain.ItemMatch{}
		}

		r <- itemMatch
	}()

	return r
}

func getNotEmptyMatchData(santanderData *domain.ItemMatch, itauData *domain.ItemMatch) *domain.ItemMatch {
	if santanderData.Attributes != nil && itauData.Attributes != nil {
		return santanderData
	}

	if santanderData.Attributes == nil && itauData.Attributes != nil {
		return itauData
	}

	return santanderData
}
