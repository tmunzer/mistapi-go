/*
Mist API

Testing SitesAnomalyAPIService

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

func Test_mistapigo_SitesAnomalyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesAnomalyAPIService GetSiteAnomalyEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var metric string

		resp, httpRes, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEvents(context.Background(), siteId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAnomalyAPIService GetSiteAnomalyEventsForClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var clientMac string
		var metric string

		resp, httpRes, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEventsForClient(context.Background(), siteId, clientMac, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAnomalyAPIService GetSiteAnomalyEventsforDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var metric string
		var deviceMac string

		resp, httpRes, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEventsforDevice(context.Background(), siteId, metric, deviceMac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
