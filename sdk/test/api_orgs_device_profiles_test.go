/*
Mist API

Testing OrgsDeviceProfilesAPIService

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

func Test_mistapigo_OrgsDeviceProfilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsDeviceProfilesAPIService AssignOrgDeviceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceprofileId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.AssignOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService CreateOrgDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.CreateOrgDeviceProfiles(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService DeleteOrgDeviceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceprofileId string

		httpRes, err := apiClient.OrgsDeviceProfilesAPI.DeleteOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService GetOrgDeviceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceprofileId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.GetOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService ListOrgDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.ListOrgDeviceProfiles(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService UnassignOrgDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceprofileId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.UnassignOrgDeviceProfiles(context.Background(), orgId, deviceprofileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsDeviceProfilesAPIService UpdateOrgDeviceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var deviceprofileId string

		resp, httpRes, err := apiClient.OrgsDeviceProfilesAPI.UpdateOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
