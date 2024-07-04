/*
Mist API

Testing SitesSLEsAPIService

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

func Test_mistapigo_SitesSLEsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesSLEsAPIService GetSiteSleClassifierDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SleSummaryScope
		var scopeId string
		var metric string
		var classifier string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleClassifierDetails(context.Background(), siteId, scope, scopeId, metric, classifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleHistogram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleHistogramScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleHistogram(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactSummaryScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactSummary(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleScope
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedApplications(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedAps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedApsScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedAps(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedChassis", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedChassisScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedChassis(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedGateways", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedGatewaysScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedGateways(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedInterfaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedInterfacesScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedInterfaces(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedSwitches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedSwitchesScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedSwitches(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedWiredClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedClientsScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedWiredClients(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleImpactedWirelessClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleImpactedUsersScopeParameter
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedWirelessClients(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleMetricClassifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleMetricClassifiersScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleMetricClassifiers(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleMetricSummaryScopeParameters
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleSummary(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSleThreshold", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleThresholdScopeParameter
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService GetSiteSlesMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleMetricsScopeParameters
		var scopeId string

		resp, httpRes, err := apiClient.SitesSLEsAPI.GetSiteSlesMetrics(context.Background(), siteId, scope, scopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService ReplaceSiteSleThreshold", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleThresholdScopeParameter
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.ReplaceSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesSLEsAPIService UpdateSiteSleThreshold", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var scope openapiclient.SiteSleThresholdScopeParameter
		var scopeId string
		var metric string

		resp, httpRes, err := apiClient.SitesSLEsAPI.UpdateSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
