/*
Mist API

Testing SitesDevicesStatsAPIService

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

func Test_mistapigo_SitesDevicesStatsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesStatsAPIService CountSiteBgpStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.CountSiteBgpStats(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService CountSiteSwOrGwPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.CountSiteSwOrGwPorts(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService CountSiteSwitchPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.CountSiteSwitchPorts(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService GetSiteAllClientsStatsByDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.GetSiteAllClientsStatsByDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService GetSiteDeviceStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.GetSiteDeviceStats(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService GetSiteGatewayMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.GetSiteGatewayMetrics(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService GetSiteMxEdgeStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxedgeId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.GetSiteMxEdgeStats(context.Background(), siteId, mxedgeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService ListSiteDevicesStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.ListSiteDevicesStats(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService ListSiteMxEdgesStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.ListSiteMxEdgesStats(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService SearchSiteBgpStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.SearchSiteBgpStats(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService SearchSiteSwOrGwPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.SearchSiteSwOrGwPorts(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesStatsAPIService SearchSiteSwitchPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesStatsAPI.SearchSiteSwitchPorts(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}