/*
Mist API

Testing SitesEventsAPIService

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

func Test_mistapigo_SitesEventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesEventsAPIService CountSiteSystemEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesEventsAPI.CountSiteSystemEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesEventsAPIService GetSiteRoamingEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesEventsAPI.GetSiteRoamingEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesEventsAPIService SearchSiteSystemEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesEventsAPI.SearchSiteSystemEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
