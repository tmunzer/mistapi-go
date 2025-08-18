// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsClientsSDKTestGetSiteSdkStatsByMap tests the behavior of the SitesStatsClientsSDK
func TestSitesStatsClientsSDKTestGetSiteSdkStatsByMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsClientsSdk.GetSiteSdkStatsByMap(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsClientsSDKTestGetSiteSdkStats tests the behavior of the SitesStatsClientsSDK
func TestSitesStatsClientsSDKTestGetSiteSdkStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	sdkclientId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsClientsSdk.GetSiteSdkStats(ctx, siteId, sdkclientId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"d56bd5fa-0a0a-4861-a9df-5ac83d3a2eeb","last_seen":1428939600,"name":"John's iPhone","network_connection":{"mac":"c3-b6-e5-af-41-15","rssi":-75,"signal_level":3,"type":"WiFi"},"uuid":"ada72f8f-1643-e5c6-94db-f2a5636f1a64","vbeacons":[{"id":"d379d29d-24b4-96c5-5dd4-6f2a2dc5aaeb","since":1428939300}],"x":60,"y":80,"zones":[{"id":"8ac84899-32db-6327-334c-9b6d58544cfe","since":1428939600}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
