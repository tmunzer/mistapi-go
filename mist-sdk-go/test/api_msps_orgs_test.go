/*
Mist API

Testing MSPsOrgsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func Test_mistapigo_MSPsOrgsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MSPsOrgsAPIService CreateMspOrg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.CreateMspOrg(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService DeleteMspOrg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string
		var orgId string

		httpRes, err := apiClient.MSPsOrgsAPI.DeleteMspOrg(context.Background(), mspId, orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService GetMspOrg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string
		var orgId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.GetMspOrg(context.Background(), mspId, orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService ListMspOrgStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.ListMspOrgStats(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService ListMspOrgs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.ListMspOrgs(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService ManageMspOrgs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		httpRes, err := apiClient.MSPsOrgsAPI.ManageMspOrgs(context.Background(), mspId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService SearchMspOrgs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.SearchMspOrgs(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsOrgsAPIService UpdateMspOrg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string
		var orgId string

		resp, httpRes, err := apiClient.MSPsOrgsAPI.UpdateMspOrg(context.Background(), mspId, orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
