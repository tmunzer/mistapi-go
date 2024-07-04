/*
Mist API

Testing OrgsServicesAPIService

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

func Test_mistapigo_OrgsServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsServicesAPIService CreateOrgService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsServicesAPI.CreateOrgService(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicesAPIService DeleteOrgService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var serviceId string

		httpRes, err := apiClient.OrgsServicesAPI.DeleteOrgService(context.Background(), orgId, serviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicesAPIService GetOrgService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var serviceId string

		resp, httpRes, err := apiClient.OrgsServicesAPI.GetOrgService(context.Background(), orgId, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicesAPIService ListOrgServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsServicesAPI.ListOrgServices(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicesAPIService UpdateOrgService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var serviceId string

		resp, httpRes, err := apiClient.OrgsServicesAPI.UpdateOrgService(context.Background(), orgId, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
