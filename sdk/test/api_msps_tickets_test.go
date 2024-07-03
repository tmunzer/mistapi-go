/*
Mist API

Testing MSPsTicketsAPIService

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

func Test_mistapigo_MSPsTicketsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MSPsTicketsAPIService CountMspTickets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsTicketsAPI.CountMspTickets(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsTicketsAPIService ListMspTickets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsTicketsAPI.ListMspTickets(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
