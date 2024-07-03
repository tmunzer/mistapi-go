/*
Mist API

Testing OrgsNACPortalsAPIService

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

func Test_mistapigo_OrgsNACPortalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsNACPortalsAPIService CreateOrgNacPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.CreateOrgNacPortal(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService DeleteOrgNacPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		httpRes, err := apiClient.OrgsNACPortalsAPI.DeleteOrgNacPortal(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService DownloadOrgNacPortalSsoSamlMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.DownloadOrgNacPortalSsoSamlMetadata(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService GetOrgNacPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.GetOrgNacPortal(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService GetOrgNacPortalSsoSamlMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.GetOrgNacPortalSsoSamlMetadata(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService ListOrgNacPortalSsoLatestFailures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.ListOrgNacPortalSsoLatestFailures(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService ListOrgNacPortals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.ListOrgNacPortals(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNACPortalsAPIService UpdateOrgNacPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var nacportalId string

		resp, httpRes, err := apiClient.OrgsNACPortalsAPI.UpdateOrgNacPortal(context.Background(), orgId, nacportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
