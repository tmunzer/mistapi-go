/*
Mist API

Testing SitesDevicesDiscoveredSwitchesAPIService

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

func Test_mistapigo_SitesDevicesDiscoveredSwitchesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesDiscoveredSwitchesAPIService CountSiteDiscoveredSwitches", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.CountSiteDiscoveredSwitches(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesDiscoveredSwitchesAPIService GetSiteDiscoveredSwitchesMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.GetSiteDiscoveredSwitchesMetrics(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesDiscoveredSwitchesAPIService SearchSiteDiscoveredSwitches", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitches(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesDiscoveredSwitchesAPIService SearchSiteDiscoveredSwitchesMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitchesMetrics(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}