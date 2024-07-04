/*
Mist API

Testing OrgsSDKInvitesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_OrgsSDKInvitesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSDKInvitesAPIService ActivateSdkInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var secret string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.ActivateSdkInvite(context.Background(), secret).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService CreateSdkInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.CreateSdkInvite(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService GetSdkInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.GetSdkInvite(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService GetSdkInviteQrCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.GetSdkInviteQrCode(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService ListSdkInvites", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.ListSdkInvites(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService RevokeSdkInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		httpRes, err := apiClient.OrgsSDKInvitesAPI.RevokeSdkInvite(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService SendSdkInviteEmail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		httpRes, err := apiClient.OrgsSDKInvitesAPI.SendSdkInviteEmail(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService SendSdkInviteSms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		httpRes, err := apiClient.OrgsSDKInvitesAPI.SendSdkInviteSms(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSDKInvitesAPIService UpdateSdkInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var sdkinviteId string

		resp, httpRes, err := apiClient.OrgsSDKInvitesAPI.UpdateSdkInvite(context.Background(), orgId, sdkinviteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
