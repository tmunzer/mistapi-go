/*
Mist API

Testing UtilitiesWANAPIService

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

func Test_mistapigo_UtilitiesWANAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UtilitiesWANAPIService ClearSiteSsrArpCache", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.ClearSiteSsrArpCache(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService ClearSiteSsrBgpRoutes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.ClearSiteSsrBgpRoutes(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService GetSiteSsrAndSrxRoutes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.GetSiteSsrAndSrxRoutes(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService GetSiteSsrAndSrxSessions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.GetSiteSsrAndSrxSessions(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService GetSiteSsrServicePath", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.GetSiteSsrServicePath(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService ReleaseSiteSsrDhcpLease", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.UtilitiesWANAPI.ReleaseSiteSsrDhcpLease(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService ServicePingFromSsr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.ServicePingFromSsr(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UtilitiesWANAPIService TestSiteSsrDnsResolution", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.UtilitiesWANAPI.TestSiteSsrDnsResolution(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
