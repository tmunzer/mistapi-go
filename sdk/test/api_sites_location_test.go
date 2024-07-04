/*
Mist API

Testing SitesLocationAPIService

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

func Test_mistapigo_SitesLocationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesLocationAPIService ClearSiteMlOverwriteForDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		httpRes, err := apiClient.SitesLocationAPI.ClearSiteMlOverwriteForDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService ClearSiteMlOverwriteForMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesLocationAPI.ClearSiteMlOverwriteForMap(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService GetSiteBeamCoverageOverview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesLocationAPI.GetSiteBeamCoverageOverview(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService GetSiteDefaultPlfForModels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesLocationAPI.GetSiteDefaultPlfForModels(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService GetSiteMachineLearningCurrentStat", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesLocationAPI.GetSiteMachineLearningCurrentStat(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService OverwriteSiteMlForDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var deviceId string

		resp, httpRes, err := apiClient.SitesLocationAPI.OverwriteSiteMlForDevice(context.Background(), siteId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService OverwriteSiteMlForMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var mapId string

		resp, httpRes, err := apiClient.SitesLocationAPI.OverwriteSiteMlForMap(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesLocationAPIService ResetSiteMlStatsByMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesLocationAPI.ResetSiteMlStatsByMap(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}