/*
Mist API

Testing SitesVBeaconsAPIService

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

func Test_mistapigo_SitesVBeaconsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesVBeaconsAPIService CreateSiteVBeacon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesVBeaconsAPI.CreateSiteVBeacon(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesVBeaconsAPIService DeleteSiteVBeacon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var vbeaconId string

		httpRes, err := apiClient.SitesVBeaconsAPI.DeleteSiteVBeacon(context.Background(), siteId, vbeaconId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesVBeaconsAPIService GetSiteVBeacon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var vbeaconId string

		resp, httpRes, err := apiClient.SitesVBeaconsAPI.GetSiteVBeacon(context.Background(), siteId, vbeaconId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesVBeaconsAPIService ListSiteVBeacons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesVBeaconsAPI.ListSiteVBeacons(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesVBeaconsAPIService UpdateSiteVBeacon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var vbeaconId string

		resp, httpRes, err := apiClient.SitesVBeaconsAPI.UpdateSiteVBeacon(context.Background(), siteId, vbeaconId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}