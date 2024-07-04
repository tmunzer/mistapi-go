/*
Mist API

Testing OrgsSecPoliciesAPIService

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

func Test_mistapigo_OrgsSecPoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSecPoliciesAPIService CreateOrgSecPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSecPoliciesAPI.CreateOrgSecPolicies(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSecPoliciesAPIService DeleteOrgSecPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var secpolicyId string

		httpRes, err := apiClient.OrgsSecPoliciesAPI.DeleteOrgSecPolicy(context.Background(), orgId, secpolicyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSecPoliciesAPIService GetOrgSecPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var secpolicyId string

		resp, httpRes, err := apiClient.OrgsSecPoliciesAPI.GetOrgSecPolicy(context.Background(), orgId, secpolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSecPoliciesAPIService ListOrgSecPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSecPoliciesAPI.ListOrgSecPolicies(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSecPoliciesAPIService UpdateOrgSecPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var secpolicyId string

		resp, httpRes, err := apiClient.OrgsSecPoliciesAPI.UpdateOrgSecPolicies(context.Background(), orgId, secpolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}