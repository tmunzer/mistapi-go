// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsIoTEndpointsTestSearchSiteIotEndpoints tests the behavior of the SitesStatsIoTEndpoints
func TestSitesStatsIoTEndpointsTestSearchSiteIotEndpoints(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apMac := "5c5b350e0001"
	mac := "63f9e299182b63f9"
	mType := "zigbee"
	mfg := "Assa Abloy"
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesStatsIoTEndpoints.SearchSiteIotEndpoints(ctx, siteId, &apMac, &mac, &mType, &mfg, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583.0,"results":[{"ap_mac":"5c5b350e0001","id":"63f9e299182b63f9","lqi":178,"mac":"63f9e299182b63f9","mfg":"Assa Abloy","model":"Assa Abloy","timestamp":1531782218,"type":"zigbee"}],"start":1531776183.0,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsIoTEndpointsTestSearchSiteIotEndpoints1 tests the behavior of the SitesStatsIoTEndpoints
func TestSitesStatsIoTEndpointsTestSearchSiteIotEndpoints1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apMac := "5c5b350e0001"
	mac := "63f9e299182b63f9"
	mType := "zigbee"
	mfg := "Assa Abloy"
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesStatsIoTEndpoints.SearchSiteIotEndpoints(ctx, siteId, &apMac, &mac, &mType, &mfg, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583.0,"results":[{"ap_mac":"5c5b350e0001","id":"63f9e299182b63f9","lqi":178,"mac":"63f9e299182b63f9","mfg":"Assa Abloy","model":"Assa Abloy","timestamp":1531782218,"type":"zigbee"}],"start":1531776183.0,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
