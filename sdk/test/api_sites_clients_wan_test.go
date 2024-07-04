/*
Mist API

Testing SitesClientsWanAPIService

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

func Test_mistapigo_SitesClientsWanAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesClientsWanAPIService CountSiteWanClientEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesClientsWanAPI.CountSiteWanClientEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesClientsWanAPIService CountSiteWanClients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesClientsWanAPI.CountSiteWanClients(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesClientsWanAPIService SearchSiteWanClientEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesClientsWanAPI.SearchSiteWanClientEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesClientsWanAPIService SearchSiteWanClients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesClientsWanAPI.SearchSiteWanClients(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}