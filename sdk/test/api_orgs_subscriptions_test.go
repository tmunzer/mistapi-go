/*
Mist API

Testing OrgsSubscriptionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_OrgsSubscriptionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSubscriptionsAPIService SubscribeOrgAlarmsReports", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsSubscriptionsAPI.SubscribeOrgAlarmsReports(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSubscriptionsAPIService UnsubscribeOrgAlarmsReports", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsSubscriptionsAPI.UnsubscribeOrgAlarmsReports(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}