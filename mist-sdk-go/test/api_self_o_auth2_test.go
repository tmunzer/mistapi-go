/*
Mist API

Testing SelfOAuth2APIService

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

func Test_mistapigo_SelfOAuth2APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SelfOAuth2APIService GetOauth2UrlForLinking", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var provider string

		resp, httpRes, err := apiClient.SelfOAuth2API.GetOauth2UrlForLinking(context.Background(), provider).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SelfOAuth2APIService LinkOauth2MistAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var provider string

		resp, httpRes, err := apiClient.SelfOAuth2API.LinkOauth2MistAccount(context.Background(), provider).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
