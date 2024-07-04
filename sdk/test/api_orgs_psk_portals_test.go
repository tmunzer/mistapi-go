/*
Mist API

Testing OrgsPskPortalsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_OrgsPskPortalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsPskPortalsAPIService CountOrgPskPortalLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.CountOrgPskPortalLogs(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService CreateOrgPskPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.CreateOrgPskPortal(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService DeleteOrgPskPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		httpRes, err := apiClient.OrgsPskPortalsAPI.DeleteOrgPskPortal(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService DeleteOrgPskPortalImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		httpRes, err := apiClient.OrgsPskPortalsAPI.DeleteOrgPskPortalImage(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService GetOrgPskPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.GetOrgPskPortal(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService ListOrgPskPortalLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.ListOrgPskPortalLogs(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService ListOrgPskPortals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.ListOrgPskPortals(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService SearchOrgPskPortalLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.SearchOrgPskPortalLogs(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService UpdateOrgPskPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		resp, httpRes, err := apiClient.OrgsPskPortalsAPI.UpdateOrgPskPortal(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService UpdateOrgPskPortalTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		httpRes, err := apiClient.OrgsPskPortalsAPI.UpdateOrgPskPortalTemplate(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsPskPortalsAPIService UploadOrgPskPortalImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var pskportalId string

		httpRes, err := apiClient.OrgsPskPortalsAPI.UploadOrgPskPortalImage(context.Background(), orgId, pskportalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
