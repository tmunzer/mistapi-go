/*
Mist API

Testing SitesDevicesWiredAPIService

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

func Test_mistapigo_SitesDevicesWiredAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesWiredAPIService DeleteSiteLocalSwitchPortConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWiredAPI.DeleteSiteLocalSwitchPortConfig(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredAPIService GetSiteSwitchesMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesWiredAPI.GetSiteSwitchesMetrics(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredAPIService UpdateSiteLocalSwitchPortConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWiredAPI.UpdateSiteLocalSwitchPortConfig(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
