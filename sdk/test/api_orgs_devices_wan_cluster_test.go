/*
Mist API

Testing OrgsDevicesWANClusterAPIService

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

func Test_mistapigo_OrgsDevicesWANClusterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsDevicesWANClusterAPIService CreateOrgGatewayHaCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsDevicesWANClusterAPI.CreateOrgGatewayHaCluster(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDevicesWANClusterAPIService DeleteOrgGatewayHaCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsDevicesWANClusterAPI.DeleteOrgGatewayHaCluster(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}