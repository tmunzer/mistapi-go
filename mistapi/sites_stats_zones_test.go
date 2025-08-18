// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsZonesTestListSiteRssiZonesStats tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestListSiteRssiZonesStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.ListSiteRssiZonesStats(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"assets_wait":{"avg":0,"max":0,"min":0,"p95":0},"clients_wait":{"avg":39259.333333333336,"max":58361,"min":12376,"p95":58361},"created_time":1733864928,"devices":[{"device_id":"00000000-0000-0000-1000-c8786708bb5d","rssi":-70}],"discovered_assets_wait":{"avg":0,"max":0,"min":0,"p95":0},"display_is_group":false,"id":"17ef7169-e000-4dcd-abc7-f721f0a8ffda","modified_time":1733864928,"name":"proximity openspace","num_assets":0,"num_clients":3,"num_discovered_assets":0,"num_sdkclients":0,"num_unconnected_clients":7,"org_id":"c5fbc9e4-12bf-436e-98c4-1c842c66ab6c","sdkclients_wait":{"avg":0,"max":0,"min":0,"p95":0},"site_id":"079fafd3-ef5c-4b23-90f0-9fcebec0023a","unconnected_clients_wait":{"avg":37552.857142857145,"max":68342,"min":6649,"p95":68342}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestListSiteRssiZonesStats1 tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestListSiteRssiZonesStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.ListSiteRssiZonesStats(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"assets_wait":{"avg":0,"max":0,"min":0,"p95":0},"clients_wait":{"avg":39259.333333333336,"max":58361,"min":12376,"p95":58361},"created_time":1733864928,"devices":[{"device_id":"00000000-0000-0000-1000-c8786708bb5d","rssi":-70}],"discovered_assets_wait":{"avg":0,"max":0,"min":0,"p95":0},"display_is_group":false,"id":"17ef7169-e000-4dcd-abc7-f721f0a8ffda","modified_time":1733864928,"name":"proximity openspace","num_assets":0,"num_clients":3,"num_discovered_assets":0,"num_sdkclients":0,"num_unconnected_clients":7,"org_id":"c5fbc9e4-12bf-436e-98c4-1c842c66ab6c","sdkclients_wait":{"avg":0,"max":0,"min":0,"p95":0},"site_id":"079fafd3-ef5c-4b23-90f0-9fcebec0023a","unconnected_clients_wait":{"avg":37552.857142857145,"max":68342,"min":6649,"p95":68342}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestGetSiteRssiZoneStats tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestGetSiteRssiZoneStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.GetSiteRssiZoneStats(ctx, siteId, zoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"assets":["df8dff06ae90"],"client_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"clients":["5684dae9ac8b"],"id":"8ac84899-32db-6327-334c-9b6d58544cfe","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","name":"Board Room","num_clients":80,"num_sdkclients":0,"sdkclients":["7e2b463d-c91c-ff7d-f3c0-6eccc6949ff8"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestGetSiteRssiZoneStats1 tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestGetSiteRssiZoneStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.GetSiteRssiZoneStats(ctx, siteId, zoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"assets":["df8dff06ae90"],"client_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"clients":["5684dae9ac8b"],"id":"8ac84899-32db-6327-334c-9b6d58544cfe","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","name":"Board Room","num_clients":80,"num_sdkclients":0,"sdkclients":["7e2b463d-c91c-ff7d-f3c0-6eccc6949ff8"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestListSiteZonesStats tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestListSiteZonesStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"
	minDuration := int(120)
	apiResponse, err := sitesStatsZones.ListSiteZonesStats(ctx, siteId, &mapId, &minDuration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"assets_waits":{"avg":0,"max":0,"min":0,"p95":0},"clients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"created_time":1616625211,"id":"123470c7-5d9d-424a-8475-8b344c621234","map_id":"123449d4-d12f-4feb-b40f-5be0e2ae1234","modified_time":1616625211,"name":"Zone A","num_assets":0,"num_clients":80,"num_sdkclients":10,"occupancy_limit":4,"org_id":"1234c1a0-6ef6-11e6-8bbf-02e208b21234","sdkclients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"site_id":"123448e6-6ef6-11e6-8bbf-02e208b21234","vertices":[{"x":732,"y":1821},{"x":732.5,"y":1731},{"x":837.5,"y":1731.5},{"x":839,"y":1821}],"vertices_m":[{"x":24.1983341951072,"y":60.198314985369144},{"x":24.21486311190714,"y":57.22310996138056},{"x":27.685935639893827,"y":57.23963887818049},{"x":27.73552239029364,"y":60.198314985369144}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestListSiteZonesStats1 tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestListSiteZonesStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"
	minDuration := int(120)
	apiResponse, err := sitesStatsZones.ListSiteZonesStats(ctx, siteId, &mapId, &minDuration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"assets_waits":{"avg":0,"max":0,"min":0,"p95":0},"clients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"created_time":1616625211,"id":"123470c7-5d9d-424a-8475-8b344c621234","map_id":"123449d4-d12f-4feb-b40f-5be0e2ae1234","modified_time":1616625211,"name":"Zone A","num_assets":0,"num_clients":80,"num_sdkclients":10,"occupancy_limit":4,"org_id":"1234c1a0-6ef6-11e6-8bbf-02e208b21234","sdkclients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"site_id":"123448e6-6ef6-11e6-8bbf-02e208b21234","vertices":[{"x":732,"y":1821},{"x":732.5,"y":1731},{"x":837.5,"y":1731.5},{"x":839,"y":1821}],"vertices_m":[{"x":24.1983341951072,"y":60.198314985369144},{"x":24.21486311190714,"y":57.22310996138056},{"x":27.685935639893827,"y":57.23963887818049},{"x":27.73552239029364,"y":60.198314985369144}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestGetSiteZoneStats tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestGetSiteZoneStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.GetSiteZoneStats(ctx, siteId, zoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"assets":["df8dff06ae90"],"client_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"clients":["5684dae9ac8b"],"id":"8ac84899-32db-6327-334c-9b6d58544cfe","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","name":"Board Room","num_clients":80,"num_sdkclients":0,"sdkclients":["7e2b463d-c91c-ff7d-f3c0-6eccc6949ff8"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsZonesTestGetSiteZoneStats1 tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestGetSiteZoneStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsZones.GetSiteZoneStats(ctx, siteId, zoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"assets":["df8dff06ae90"],"client_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"clients":["5684dae9ac8b"],"id":"8ac84899-32db-6327-334c-9b6d58544cfe","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","name":"Board Room","num_clients":80,"num_sdkclients":0,"sdkclients":["7e2b463d-c91c-ff7d-f3c0-6eccc6949ff8"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
