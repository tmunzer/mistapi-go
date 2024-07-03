/*
Mist API

Testing SitesWxTagsAPIService

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

func Test_mistapigo_SitesWxTagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesWxTagsAPIService CreateSiteWxTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.CreateSiteWxTag(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService DeleteSiteWxTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxtagId string

		httpRes, err := apiClient.SitesWxTagsAPI.DeleteSiteWxTag(context.Background(), siteId, wxtagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService GetSiteApplicationList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.GetSiteApplicationList(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService GetSiteCurrentMatchingClientsOfAWxTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxtagId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.GetSiteCurrentMatchingClientsOfAWxTag(context.Background(), siteId, wxtagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService GetSiteWxTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxtagId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.GetSiteWxTag(context.Background(), siteId, wxtagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService ListSiteWxTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.ListSiteWxTags(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxTagsAPIService UpdateSiteWxTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxtagId string

		resp, httpRes, err := apiClient.SitesWxTagsAPI.UpdateSiteWxTag(context.Background(), siteId, wxtagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
