// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsClientsMarvisTestCountOrgMarvisClientEvents tests the behavior of the OrgsClientsMarvis
func TestOrgsClientsMarvisTestCountOrgMarvisClientEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "type"

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsClientsMarvis.CountOrgMarvisClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsMarvisTestCountOrgMarvisClientEvents1 tests the behavior of the OrgsClientsMarvis
func TestOrgsClientsMarvisTestCountOrgMarvisClientEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "type"

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsClientsMarvis.CountOrgMarvisClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsMarvisTestSearchOrgMarvisClientEvents tests the behavior of the OrgsClientsMarvis
func TestOrgsClientsMarvisTestSearchOrgMarvisClientEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsClientsMarvis.SearchOrgMarvisClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":1000,"results":[{"bssid":"5c5b35000002","channel":11,"device_id":"da088609-8e8d-6d8b-0e40-fe1dc94b9218","hostname":"jdoe123-dell","location":{"map_id":"7735ef91-83b4-6116-1c1a-57819af48867","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1678377926,"x":12.5,"y":45.0},"neighbor_ap_report":[{"band":"5","bssid":"5c5b35000003","channel":44,"rssi":-55},{"band":"5","bssid":"5c5b35000004","channel":36,"rssi":-70}],"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","pre_bssid":"5c5b35000001","pre_channel":7,"pre_rssi":-76,"rssi":-53,"ssid":"Corp","timestamp":1678377926,"type":"MARVISCLIENT_ROAMED","wifi_ip":"10.10.20.1","wifi_mac":"f01c2df166e0"},{"device_id":"da088609-8e8d-6d8b-0e40-fe1dc94b9218","hostname":"jdoe123-dell","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","percent":9,"timestamp":1678378040,"type":"MARVISCLIENT_LOW_BATTERY","wifi_ip":"10.10.20.1","wifi_mac":"f01c2df166e0"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsMarvisTestSearchOrgMarvisClientEvents1 tests the behavior of the OrgsClientsMarvis
func TestOrgsClientsMarvisTestSearchOrgMarvisClientEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsClientsMarvis.SearchOrgMarvisClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":1000,"results":[{"bssid":"5c5b35000002","channel":11,"device_id":"da088609-8e8d-6d8b-0e40-fe1dc94b9218","hostname":"jdoe123-dell","location":{"map_id":"7735ef91-83b4-6116-1c1a-57819af48867","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1678377926,"x":12.5,"y":45.0},"neighbor_ap_report":[{"band":"5","bssid":"5c5b35000003","channel":44,"rssi":-55},{"band":"5","bssid":"5c5b35000004","channel":36,"rssi":-70}],"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","pre_bssid":"5c5b35000001","pre_channel":7,"pre_rssi":-76,"rssi":-53,"ssid":"Corp","timestamp":1678377926,"type":"MARVISCLIENT_ROAMED","wifi_ip":"10.10.20.1","wifi_mac":"f01c2df166e0"},{"device_id":"da088609-8e8d-6d8b-0e40-fe1dc94b9218","hostname":"jdoe123-dell","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","percent":9,"timestamp":1678378040,"type":"MARVISCLIENT_LOW_BATTERY","wifi_ip":"10.10.20.1","wifi_mac":"f01c2df166e0"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsMarvisTestDeleteOrgMarvisClient tests the behavior of the OrgsClientsMarvis
func TestOrgsClientsMarvisTestDeleteOrgMarvisClient(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsClientsMarvis.DeleteOrgMarvisClient(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
