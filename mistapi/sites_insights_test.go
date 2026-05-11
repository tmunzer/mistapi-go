// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesInsightsTestGetSiteInsightMetrics tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetrics(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "num_clients,num_aps"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetrics(ctx, siteId, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"interval":0,"results":[{}],"start":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetrics1 tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetrics1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "num_clients,num_aps"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetrics(ctx, siteId, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"interval":0,"results":[{}],"start":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForAP tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForAP(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "num_clients,num_stressed_clients"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForAP(ctx, siteId, deviceId, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1604347200,"interval":3600,"limit":168,"page":1,"results":[10,11,12,12,10,9,9,9,10,10,11,11,11,11,11,11,11,10,11,11,10,11,11,10],"rt":["2020-11-01 20:00:00+00:00","2020-11-01 21:00:00+00:00","2020-11-01 22:00:00+00:00","2020-11-01 23:00:00+00:00","2020-11-02 00:00:00+00:00","2020-11-02 01:00:00+00:00","2020-11-02 02:00:00+00:00","2020-11-02 03:00:00+00:00","2020-11-02 04:00:00+00:00","2020-11-02 05:00:00+00:00","2020-11-02 06:00:00+00:00","2020-11-02 07:00:00+00:00","2020-11-02 08:00:00+00:00","2020-11-02 09:00:00+00:00","2020-11-02 10:00:00+00:00","2020-11-02 11:00:00+00:00","2020-11-02 12:00:00+00:00","2020-11-02 13:00:00+00:00","2020-11-02 14:00:00+00:00","2020-11-02 15:00:00+00:00","2020-11-02 16:00:00+00:00","2020-11-02 17:00:00+00:00","2020-11-02 18:00:00+00:00","2020-11-02 19:00:00+00:00"],"start":1604260800}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForAP1 tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForAP1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "num_clients,num_stressed_clients"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForAP(ctx, siteId, deviceId, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1604347200,"interval":3600,"limit":168,"page":1,"results":[10,11,12,12,10,9,9,9,10,10,11,11,11,11,11,11,11,10,11,11,10,11,11,10],"rt":["2020-11-01 20:00:00+00:00","2020-11-01 21:00:00+00:00","2020-11-01 22:00:00+00:00","2020-11-01 23:00:00+00:00","2020-11-02 00:00:00+00:00","2020-11-02 01:00:00+00:00","2020-11-02 02:00:00+00:00","2020-11-02 03:00:00+00:00","2020-11-02 04:00:00+00:00","2020-11-02 05:00:00+00:00","2020-11-02 06:00:00+00:00","2020-11-02 07:00:00+00:00","2020-11-02 08:00:00+00:00","2020-11-02 09:00:00+00:00","2020-11-02 10:00:00+00:00","2020-11-02 11:00:00+00:00","2020-11-02 12:00:00+00:00","2020-11-02 13:00:00+00:00","2020-11-02 14:00:00+00:00","2020-11-02 15:00:00+00:00","2020-11-02 16:00:00+00:00","2020-11-02 17:00:00+00:00","2020-11-02 18:00:00+00:00","2020-11-02 19:00:00+00:00"],"start":1604260800}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForClient tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForClient(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	metrics := "top-app-by-num_client,top-app-by-bytes"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"interval":0,"results":[{}],"start":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForClient1 tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForClient1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	metrics := "top-app-by-num_client,top-app-by-bytes"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metrics, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"interval":0,"results":[{}],"start":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForGateway tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForGateway(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "tx_bps,rx_bps"
	portId := "ge-0/0/1"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForGateway(ctx, siteId, deviceId, metrics, &portId, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1604347200,"interval":3600,"limit":168,"page":1,"results":[10,11,12,12,10,9,9,9,10,10,11,11,11,11,11,11,11,10,11,11,10,11,11,10],"rt":["2020-11-01 20:00:00+00:00","2020-11-01 21:00:00+00:00","2020-11-01 22:00:00+00:00","2020-11-01 23:00:00+00:00","2020-11-02 00:00:00+00:00","2020-11-02 01:00:00+00:00","2020-11-02 02:00:00+00:00","2020-11-02 03:00:00+00:00","2020-11-02 04:00:00+00:00","2020-11-02 05:00:00+00:00","2020-11-02 06:00:00+00:00","2020-11-02 07:00:00+00:00","2020-11-02 08:00:00+00:00","2020-11-02 09:00:00+00:00","2020-11-02 10:00:00+00:00","2020-11-02 11:00:00+00:00","2020-11-02 12:00:00+00:00","2020-11-02 13:00:00+00:00","2020-11-02 14:00:00+00:00","2020-11-02 15:00:00+00:00","2020-11-02 16:00:00+00:00","2020-11-02 17:00:00+00:00","2020-11-02 18:00:00+00:00","2020-11-02 19:00:00+00:00"],"start":1604260800}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesInsightsTestGetSiteInsightMetricsForGateway1 tests the behavior of the SitesInsights
func TestSitesInsightsTestGetSiteInsightMetricsForGateway1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	metrics := "tx_bps,rx_bps"
	portId := "ge-0/0/1"

	duration := "1d"
	interval := "10m"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesInsights.GetSiteInsightMetricsForGateway(ctx, siteId, deviceId, metrics, &portId, nil, nil, &duration, &interval, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1604347200,"interval":3600,"limit":168,"page":1,"results":[10,11,12,12,10,9,9,9,10,10,11,11,11,11,11,11,11,10,11,11,10,11,11,10],"rt":["2020-11-01 20:00:00+00:00","2020-11-01 21:00:00+00:00","2020-11-01 22:00:00+00:00","2020-11-01 23:00:00+00:00","2020-11-02 00:00:00+00:00","2020-11-02 01:00:00+00:00","2020-11-02 02:00:00+00:00","2020-11-02 03:00:00+00:00","2020-11-02 04:00:00+00:00","2020-11-02 05:00:00+00:00","2020-11-02 06:00:00+00:00","2020-11-02 07:00:00+00:00","2020-11-02 08:00:00+00:00","2020-11-02 09:00:00+00:00","2020-11-02 10:00:00+00:00","2020-11-02 11:00:00+00:00","2020-11-02 12:00:00+00:00","2020-11-02 13:00:00+00:00","2020-11-02 14:00:00+00:00","2020-11-02 15:00:00+00:00","2020-11-02 16:00:00+00:00","2020-11-02 17:00:00+00:00","2020-11-02 18:00:00+00:00","2020-11-02 19:00:00+00:00"],"start":1604260800}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
