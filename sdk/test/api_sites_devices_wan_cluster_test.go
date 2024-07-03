/*
Mist API

Testing SitesDevicesWANClusterAPIService

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

func Test_mistapigo_SitesDevicesWANClusterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesWANClusterAPIService CreateSiteDeviceHaCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWANClusterAPI.CreateSiteDeviceHaCluster(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWANClusterAPIService DeleteSiteDeviceHaCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWANClusterAPI.DeleteSiteDeviceHaCluster(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWANClusterAPIService SwapSiteDeviceHaClusterNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWANClusterAPI.SwapSiteDeviceHaClusterNode(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
