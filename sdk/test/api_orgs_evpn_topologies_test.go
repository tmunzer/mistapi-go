/*
Mist API

Testing OrgsEVPNTopologiesAPIService

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

func Test_mistapigo_OrgsEVPNTopologiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsEVPNTopologiesAPIService CreateOrgEvpnTopology", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsEVPNTopologiesAPI.CreateOrgEvpnTopology(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsEVPNTopologiesAPIService DeleteOrgEvpnTopology", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var evpnTopologyId string

		httpRes, err := apiClient.OrgsEVPNTopologiesAPI.DeleteOrgEvpnTopology(context.Background(), orgId, evpnTopologyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsEVPNTopologiesAPIService GetOrgEvpnTolopogy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var evpnTopologyId string

		resp, httpRes, err := apiClient.OrgsEVPNTopologiesAPI.GetOrgEvpnTolopogy(context.Background(), orgId, evpnTopologyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsEVPNTopologiesAPIService ListOrgEvpnTopologies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsEVPNTopologiesAPI.ListOrgEvpnTopologies(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsEVPNTopologiesAPIService UpdateOrgEvpnTopology", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var evpnTopologyId string

		resp, httpRes, err := apiClient.OrgsEVPNTopologiesAPI.UpdateOrgEvpnTopology(context.Background(), orgId, evpnTopologyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
