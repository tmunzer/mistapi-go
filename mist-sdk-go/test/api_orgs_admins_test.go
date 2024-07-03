/*
Mist API

Testing OrgsAdminsAPIService

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

func Test_mistapigo_OrgsAdminsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsAdminsAPIService InviteOrgAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsAdminsAPI.InviteOrgAdmin(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAdminsAPIService ListOrgAdmins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsAdminsAPI.ListOrgAdmins(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAdminsAPIService RevokeOrgAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var adminId string

		httpRes, err := apiClient.OrgsAdminsAPI.RevokeOrgAdmin(context.Background(), orgId, adminId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAdminsAPIService UninviteOrgAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var inviteId string

		httpRes, err := apiClient.OrgsAdminsAPI.UninviteOrgAdmin(context.Background(), orgId, inviteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAdminsAPIService UpdateOrgAdmin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var adminId string

		resp, httpRes, err := apiClient.OrgsAdminsAPI.UpdateOrgAdmin(context.Background(), orgId, adminId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAdminsAPIService UpdateOrgAdminInvite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var inviteId string

		httpRes, err := apiClient.OrgsAdminsAPI.UpdateOrgAdminInvite(context.Background(), orgId, inviteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}