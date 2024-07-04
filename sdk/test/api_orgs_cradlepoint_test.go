/*
Mist API

Testing OrgsCradlepointAPIService

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

func Test_mistapigo_OrgsCradlepointAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsCradlepointAPIService DeleteOrgCradlepointConnection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsCradlepointAPI.DeleteOrgCradlepointConnection(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsCradlepointAPIService SetupOrgCradlepointConnectionToMist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsCradlepointAPI.SetupOrgCradlepointConnectionToMist(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsCradlepointAPIService SyncOrgCradlepointRouters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsCradlepointAPI.SyncOrgCradlepointRouters(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsCradlepointAPIService UpdateOrgCradlepointConnectionToMist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsCradlepointAPI.UpdateOrgCradlepointConnectionToMist(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
