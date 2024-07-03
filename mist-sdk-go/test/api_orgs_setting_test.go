/*
Mist API

Testing OrgsSettingAPIService

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

func Test_mistapigo_OrgsSettingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSettingAPIService CreateOrgWirelessClientsBlocklist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSettingAPI.CreateOrgWirelessClientsBlocklist(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSettingAPIService DeleteOrgWirelessClientsBlocklist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsSettingAPI.DeleteOrgWirelessClientsBlocklist(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSettingAPIService GetOrgSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSettingAPI.GetOrgSettings(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSettingAPIService SetOrgCustomBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSettingAPI.SetOrgCustomBucket(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSettingAPIService UpdateOrgSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSettingAPI.UpdateOrgSettings(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSettingAPIService VerifyOrgCustomBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsSettingAPI.VerifyOrgCustomBucket(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}