/*
Mist API

Testing UtilitiesMxEdgeAPIService

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

func Test_mistapigo_UtilitiesMxEdgeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UtilitiesMxEdgeAPIService PreemptSitesMxTunnel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxtunnelId string

		resp, httpRes, err := apiClient.UtilitiesMxEdgeAPI.PreemptSitesMxTunnel(context.Background(), siteId, mxtunnelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
