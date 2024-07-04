/*
Mist API

Testing SitesAssetFiltersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_SitesAssetFiltersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesAssetFiltersAPIService CreateSiteAssetFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesAssetFiltersAPI.CreateSiteAssetFilters(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAssetFiltersAPIService DeleteSiteAssetFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var assetfilterId string

		httpRes, err := apiClient.SitesAssetFiltersAPI.DeleteSiteAssetFilter(context.Background(), siteId, assetfilterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAssetFiltersAPIService GetSiteAssetFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var assetfilterId string

		resp, httpRes, err := apiClient.SitesAssetFiltersAPI.GetSiteAssetFilter(context.Background(), siteId, assetfilterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAssetFiltersAPIService ListSiteAssetFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesAssetFiltersAPI.ListSiteAssetFilters(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAssetFiltersAPIService UpdateSiteAssetFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var assetfilterId string

		resp, httpRes, err := apiClient.SitesAssetFiltersAPI.UpdateSiteAssetFilter(context.Background(), siteId, assetfilterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
