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
	expected := `{"limit":100,"results":[{"bssid":"00:11:22:33:44:55","channel":6,"device_id":"0c2a2c6c-5a95-4956-a02d-e1b39c2e5c6e","hostname":"my-android-phone","neighbor_ap_report":[{"band":"2.4GHz","bssid":"aa:bb:cc:dd:ee:ff","channel":1,"rssi":-72}],"org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","pre_bssid":"aa:bb:cc:dd:ee:ff","pre_channel":1,"pre_rssi":-70,"rssi":-58,"ssid":"Corp-WiFi","timestamp":1717027200,"type":"ROAM","wifi_ip":"192.168.1.55","wifi_mac":"00:aa:bb:cc:dd:ee"}],"total":1}`
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
	expected := `{"limit":100,"results":[{"bssid":"00:11:22:33:44:55","channel":6,"device_id":"0c2a2c6c-5a95-4956-a02d-e1b39c2e5c6e","hostname":"my-android-phone","neighbor_ap_report":[{"band":"2.4GHz","bssid":"aa:bb:cc:dd:ee:ff","channel":1,"rssi":-72}],"org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","pre_bssid":"aa:bb:cc:dd:ee:ff","pre_channel":1,"pre_rssi":-70,"rssi":-58,"ssid":"Corp-WiFi","timestamp":1717027200,"type":"ROAM","wifi_ip":"192.168.1.55","wifi_mac":"00:aa:bb:cc:dd:ee"}],"total":1}`
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
