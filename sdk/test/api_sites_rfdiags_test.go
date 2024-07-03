/*
Mist API

Testing SitesRfdiagsAPIService

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

func Test_mistapigo_SitesRfdiagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesRfdiagsAPIService DeleteSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var rfdiagId string

		httpRes, err := apiClient.SitesRfdiagsAPI.DeleteSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService DownloadSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var rfdiagId string

		resp, httpRes, err := apiClient.SitesRfdiagsAPI.DownloadSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService GetSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var rfdiagId string

		resp, httpRes, err := apiClient.SitesRfdiagsAPI.GetSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService GetSiteSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesRfdiagsAPI.GetSiteSiteRfdiagRecording(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService StartSiteRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesRfdiagsAPI.StartSiteRecording(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService StopSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var rfdiagId string

		httpRes, err := apiClient.SitesRfdiagsAPI.StopSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesRfdiagsAPIService UpdateSiteRfdiagRecording", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var rfdiagId string

		resp, httpRes, err := apiClient.SitesRfdiagsAPI.UpdateSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
