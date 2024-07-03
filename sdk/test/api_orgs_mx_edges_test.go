/*
Mist API

Testing OrgsMxEdgesAPIService

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

func Test_mistapigo_OrgsMxEdgesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsMxEdgesAPIService AddOrgMxEdgeImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string
		var imageNumber int32

		httpRes, err := apiClient.OrgsMxEdgesAPI.AddOrgMxEdgeImage(context.Background(), orgId, mxedgeId, imageNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService AssignOrgMxEdgeToSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.AssignOrgMxEdgeToSite(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService BounceOrgMxEdgeDataPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		httpRes, err := apiClient.OrgsMxEdgesAPI.BounceOrgMxEdgeDataPorts(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService ClaimOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.ClaimOrgMxEdge(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService ControlOrgMxEdgeServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string
		var name string
		var action string

		httpRes, err := apiClient.OrgsMxEdgesAPI.ControlOrgMxEdgeServices(context.Background(), orgId, mxedgeId, name, action).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService CountOrgMxEdges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.CountOrgMxEdges(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService CountOrgSiteMxEdgeEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.CountOrgSiteMxEdgeEvents(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService CreateOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.CreateOrgMxEdge(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService DeleteOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		httpRes, err := apiClient.OrgsMxEdgesAPI.DeleteOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService DeleteOrgMxEdgeImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string
		var imageNumber int32

		httpRes, err := apiClient.OrgsMxEdgesAPI.DeleteOrgMxEdgeImage(context.Background(), orgId, mxedgeId, imageNumber).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService GetOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService GetOrgMxEdgeStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdgeStats(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService GetOrgMxEdgeUpgradeInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdgeUpgradeInfo(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService ListOrgMxEdges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.ListOrgMxEdges(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService ListOrgMxEdgesStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.ListOrgMxEdgesStats(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService RestartOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		httpRes, err := apiClient.OrgsMxEdgesAPI.RestartOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService SearchOrgMistEdgeEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.SearchOrgMistEdgeEvents(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService SearchOrgMxEdges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.SearchOrgMxEdges(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService UnassignOrgMxEdgeFromSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.UnassignOrgMxEdgeFromSite(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService UnregisterOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		httpRes, err := apiClient.OrgsMxEdgesAPI.UnregisterOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService UpdateOrgMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		resp, httpRes, err := apiClient.OrgsMxEdgesAPI.UpdateOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxEdgesAPIService UploadOrgMxEdgeSupportFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxedgeId string

		httpRes, err := apiClient.OrgsMxEdgesAPI.UploadOrgMxEdgeSupportFiles(context.Background(), orgId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
