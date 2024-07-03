/*
Mist API

Testing OrgsDevicesOthersAPIService

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

func Test_mistapigo_OrgsDevicesOthersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsDevicesOthersAPIService CountOrgOtherDeviceEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsDevicesOthersAPI.CountOrgOtherDeviceEvents(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService DeleteOrgOtherDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceMac string

		httpRes, err := apiClient.OrgsDevicesOthersAPI.DeleteOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService GetOrgOtherDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceMac string

		resp, httpRes, err := apiClient.OrgsDevicesOthersAPI.GetOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService GetOrgOtherDeviceStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceMac string

		resp, httpRes, err := apiClient.OrgsDevicesOthersAPI.GetOrgOtherDeviceStats(context.Background(), orgId, deviceMac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService ListOrgOtherDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsDevicesOthersAPI.ListOrgOtherDevices(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService RebootOrgOtherDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceMac string

		httpRes, err := apiClient.OrgsDevicesOthersAPI.RebootOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService SearchOrgOtherDeviceEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsDevicesOthersAPI.SearchOrgOtherDeviceEvents(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService UpdateOrgOtherDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceMac string

		httpRes, err := apiClient.OrgsDevicesOthersAPI.UpdateOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesOthersAPIService UpdateOrgOtherDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsDevicesOthersAPI.UpdateOrgOtherDevices(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
