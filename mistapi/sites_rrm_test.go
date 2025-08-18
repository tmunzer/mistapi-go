// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesRRMTestGetSiteCurrentChannelPlanning tests the behavior of the SitesRRM
func TestSitesRRMTestGetSiteCurrentChannelPlanning(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRrm.GetSiteCurrentChannelPlanning(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band_24":{},"band_24_metric":{"avg_aps_per_channel":0,"channel_distribution_uniformity":0,"cochannel_neighbors":0,"density":0,"naps_by_channel":{},"naps_by_power":{},"neighbors":0,"noise":0},"band_5":{},"band_5_metric":{"avg_aps_per_channel":0,"channel_distribution_uniformity":0,"cochannel_neighbors":0,"density":0,"naps_by_channel":{},"naps_by_power":{},"neighbors":0,"noise":0},"rftemplate":{"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":5,"power_max":3,"power_min":18,"preamble":"auto","usage":"24"},"band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":100,"disabled":true,"power":9,"power_max":6,"power_min":17,"preamble":"auto"},"country_code":"string","created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","model_specific":{"property1":{"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":9,"power_max":6,"power_min":17,"preamble":"auto","usage":"rrm"},"band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":112,"disabled":true,"power":10,"power_max":6,"power_min":15,"preamble":"auto"}}},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},"rftemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rftemplate_name":"string","status":"updating","timestamp":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRRMTestGetSiteCurrentChannelPlanning1 tests the behavior of the SitesRRM
func TestSitesRRMTestGetSiteCurrentChannelPlanning1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRrm.GetSiteCurrentChannelPlanning(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band_24":{},"band_24_metric":{"avg_aps_per_channel":0,"channel_distribution_uniformity":0,"cochannel_neighbors":0,"density":0,"naps_by_channel":{},"naps_by_power":{},"neighbors":0,"noise":0},"band_5":{},"band_5_metric":{"avg_aps_per_channel":0,"channel_distribution_uniformity":0,"cochannel_neighbors":0,"density":0,"naps_by_channel":{},"naps_by_power":{},"neighbors":0,"noise":0},"rftemplate":{"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":5,"power_max":3,"power_min":18,"preamble":"auto","usage":"24"},"band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":100,"disabled":true,"power":9,"power_max":6,"power_min":17,"preamble":"auto"},"country_code":"string","created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","model_specific":{"property1":{"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":9,"power_max":6,"power_min":17,"preamble":"auto","usage":"rrm"},"band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":112,"disabled":true,"power":10,"power_max":6,"power_min":15,"preamble":"auto"}}},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},"rftemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rftemplate_name":"string","status":"updating","timestamp":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRRMTestListSiteRrmEvents tests the behavior of the SitesRRM
func TestSitesRRMTestListSiteRrmEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesRrm.ListSiteRrmEvents(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"next":"/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/rrm?start=1428939600&end=1428949600&limit=200&token=001a0010000000120010000005005880ec18000004776c616e007fffffeb067ab8e29c1d659b6a7c8cf698bf81490003","results":[{"ap_id":"00000000-0000-0000-1000-5c5b359e4fe0","band":"24","bandwidth":20,"channel":6,"event":"scheduled-site_rrm","power":5,"pre_bandwidth":20,"pre_channel":1,"pre_power":11,"pre_usage":"24","timestamp":1428939600,"usage":"24"}],"start":1428939600}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRRMTestListSiteRrmEvents1 tests the behavior of the SitesRRM
func TestSitesRRMTestListSiteRrmEvents1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesRrm.ListSiteRrmEvents(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"next":"/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/rrm?start=1428939600&end=1428949600&limit=200&token=001a0010000000120010000005005880ec18000004776c616e007fffffeb067ab8e29c1d659b6a7c8cf698bf81490003","results":[{"ap_id":"00000000-0000-0000-1000-5c5b359e4fe0","band":"24","bandwidth":20,"channel":6,"event":"scheduled-site_rrm","power":5,"pre_bandwidth":20,"pre_channel":1,"pre_power":11,"pre_usage":"24","timestamp":1428939600,"usage":"24"}],"start":1428939600}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
