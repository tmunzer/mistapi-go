/*
Mist API

Testing MSPsLicensesAPIService

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

func Test_mistapigo_MSPsLicensesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MSPsLicensesAPIService ClaimMspLicence", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsLicensesAPI.ClaimMspLicence(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsLicensesAPIService ListMspLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsLicensesAPI.ListMspLicenses(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsLicensesAPIService ListMspOrgLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsLicensesAPI.ListMspOrgLicenses(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsLicensesAPIService MoveOrDeleteMspLicenseToAnotherOrg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		httpRes, err := apiClient.MSPsLicensesAPI.MoveOrDeleteMspLicenseToAnotherOrg(context.Background(), mspId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
