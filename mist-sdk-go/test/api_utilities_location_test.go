/*
Mist API

Testing UtilitiesLocationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func Test_mistapigo_UtilitiesLocationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UtilitiesLocationAPIService SendSiteDevicesArbitratryBleBeacon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		httpRes, err := apiClient.UtilitiesLocationAPI.SendSiteDevicesArbitratryBleBeacon(context.Background(), siteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}