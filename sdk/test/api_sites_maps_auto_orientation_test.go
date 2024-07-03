/*
Mist API

Testing SitesMapsAutoOrientationAPIService

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

func Test_mistapigo_SitesMapsAutoOrientationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesMapsAutoOrientationAPIService ClearSiteApAutoOrient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesMapsAutoOrientationAPI.ClearSiteApAutoOrient(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoOrientationAPIService DeleteSiteApAutoOrientation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mapId string
		var siteId string

		httpRes, err := apiClient.SitesMapsAutoOrientationAPI.DeleteSiteApAutoOrientation(context.Background(), mapId, siteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoOrientationAPIService StartSiteApAutoOrientation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mapId string
		var siteId string

		resp, httpRes, err := apiClient.SitesMapsAutoOrientationAPI.StartSiteApAutoOrientation(context.Background(), mapId, siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
