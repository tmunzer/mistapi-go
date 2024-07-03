/*
Mist API

Testing SitesServicesAPIService

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

func Test_mistapigo_SitesServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesServicesAPIService CountSiteServicePathEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesServicesAPI.CountSiteServicePathEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesServicesAPIService ListSiteServicesDerived", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesServicesAPI.ListSiteServicesDerived(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesServicesAPIService SearchSiteServicePathEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesServicesAPI.SearchSiteServicePathEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
