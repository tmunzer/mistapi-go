/*
Mist API

Testing SitesDevicesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistapi-go"
)

func Test_mistapigo_SitesDevicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesAPIService AddSiteDeviceImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string
		var imageNumber int32

		httpRes, err := apiClient.SitesDevicesAPI.AddSiteDeviceImage(context.Background(), siteId, deviceId, imageNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService CountSiteDeviceConfigHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.CountSiteDeviceConfigHistory(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService CountSiteDeviceEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.CountSiteDeviceEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService CountSiteDeviceLastConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.CountSiteDeviceLastConfig(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService CountSiteDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.CountSiteDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService DeleteSiteDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesAPI.DeleteSiteDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService DeleteSiteDeviceImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string
		var imageNumber int32

		httpRes, err := apiClient.SitesDevicesAPI.DeleteSiteDeviceImage(context.Background(), siteId, deviceId, imageNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService ExportSiteDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.ExportSiteDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService GetSiteDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.GetSiteDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService ImportSiteDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.ImportSiteDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService ListSiteDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.ListSiteDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService RestartSiteMultipleDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		httpRes, err := apiClient.SitesDevicesAPI.RestartSiteMultipleDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService SearchSiteDeviceConfigHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.SearchSiteDeviceConfigHistory(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService SearchSiteDeviceEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.SearchSiteDeviceEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService SearchSiteDeviceLastConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.SearchSiteDeviceLastConfigs(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService SearchSiteDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.SearchSiteDevices(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesAPIService UpdateSiteDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesAPI.UpdateSiteDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
