/*
Mist API

Testing SitesVPNsAPIService

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

func Test_mistapigo_SitesVPNsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesVPNsAPIService ListSiteVpnsDerived", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesVPNsAPI.ListSiteVpnsDerived(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
