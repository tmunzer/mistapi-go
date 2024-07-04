/*
Mist API

Testing SitesApplicationsAPIService

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

func Test_mistapigo_SitesApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesApplicationsAPIService CountSiteApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesApplicationsAPI.CountSiteApps(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesApplicationsAPIService ListSiteApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesApplicationsAPI.ListSiteApps(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}