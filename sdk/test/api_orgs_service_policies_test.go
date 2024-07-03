/*
Mist API

Testing OrgsServicePoliciesAPIService

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

func Test_mistapigo_OrgsServicePoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsServicePoliciesAPIService CreateOrgServicePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsServicePoliciesAPI.CreateOrgServicePolicy(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicePoliciesAPIService DeleteOrgServicePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var servicepolicyId string

		httpRes, err := apiClient.OrgsServicePoliciesAPI.DeleteOrgServicePolicy(context.Background(), orgId, servicepolicyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicePoliciesAPIService GetOrgServicePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var servicepolicyId string

		resp, httpRes, err := apiClient.OrgsServicePoliciesAPI.GetOrgServicePolicy(context.Background(), orgId, servicepolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicePoliciesAPIService ListOrgServicePolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsServicePoliciesAPI.ListOrgServicePolicies(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsServicePoliciesAPIService UpdateOrgServicePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var servicepolicyId string

		resp, httpRes, err := apiClient.OrgsServicePoliciesAPI.UpdateOrgServicePolicy(context.Background(), orgId, servicepolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
