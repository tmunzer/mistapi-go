/*
Mist API

Testing SitesDevicesWiredVirtualChassisAPIService

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

func Test_mistapigo_SitesDevicesWiredVirtualChassisAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesDevicesWiredVirtualChassisAPIService CreateSiteVirtualChassis", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesWiredVirtualChassisAPI.CreateSiteVirtualChassis(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredVirtualChassisAPIService DeleteSiteVirtualChassis", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWiredVirtualChassisAPI.DeleteSiteVirtualChassis(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredVirtualChassisAPIService GetSiteDeviceVirtualChassis", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesWiredVirtualChassisAPI.GetSiteDeviceVirtualChassis(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredVirtualChassisAPIService SetSiteVcPort", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesDevicesWiredVirtualChassisAPI.SetSiteVcPort(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesDevicesWiredVirtualChassisAPIService UpdateSiteVirtualChassisMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesDevicesWiredVirtualChassisAPI.UpdateSiteVirtualChassisMember(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
