/*
Mist API

Testing SitesServicePoliciesAPIService

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

func Test_mistapigo_SitesServicePoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesServicePoliciesAPIService ListSiteServicePoliciesDerived", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesServicePoliciesAPI.ListSiteServicePoliciesDerived(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
