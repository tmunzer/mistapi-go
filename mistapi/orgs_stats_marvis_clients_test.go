// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsStatsMarvisClientsTestCountOrgMarvisClientsStats tests the behavior of the OrgsStatsMarvisClients
func TestOrgsStatsMarvisClientsTestCountOrgMarvisClientsStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "os_type"

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsStatsMarvisClients.CountOrgMarvisClientsStats(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestOrgsStatsMarvisClientsTestCountOrgMarvisClientsStats1 tests the behavior of the OrgsStatsMarvisClients
func TestOrgsStatsMarvisClientsTestCountOrgMarvisClientsStats1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "os_type"

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsStatsMarvisClients.CountOrgMarvisClientsStats(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestOrgsStatsMarvisClientsTestSearchOrgMarvisClientsStats tests the behavior of the OrgsStatsMarvisClients
func TestOrgsStatsMarvisClientsTestSearchOrgMarvisClientsStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsStatsMarvisClients.SearchOrgMarvisClientsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":100,"results":[{"battery_charging":false,"battery_level":82,"cpu_background":2.1,"cpu_idle":78.3,"cpu_system":5.4,"cpu_user":14.2,"device_id":"0c2a2c6c-5a95-4956-a02d-e1b39c2e5c6e","hostname":"my-android-phone","location":{"map_id":"b4695157-0d1d-4da0-8f9e-5c38149d8b81","site_id":"d14f16d8-d14f-11e5-8e81-1258369c38a9","timestamp":1717027100,"x":423.5,"y":201.0},"memory_total":8589934592,"memory_usage":3758096384,"mfg":"Samsung","model":"Galaxy S23","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","os_type":"Android","os_version":"14","serial":"R5CX123456A","storage_total":128849018880,"storage_usage":52428800000,"timestamp":1717027200,"wifi_band":"5GHz","wifi_bssid":"00:11:22:33:44:55","wifi_channel":36,"wifi_ip":"192.168.1.55","wifi_mac":"00:aa:bb:cc:dd:ee","wifi_rssi":-58,"wifi_ssid":"Corp-WiFi"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsMarvisClientsTestSearchOrgMarvisClientsStats1 tests the behavior of the OrgsStatsMarvisClients
func TestOrgsStatsMarvisClientsTestSearchOrgMarvisClientsStats1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsStatsMarvisClients.SearchOrgMarvisClientsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":100,"results":[{"battery_charging":false,"battery_level":82,"cpu_background":2.1,"cpu_idle":78.3,"cpu_system":5.4,"cpu_user":14.2,"device_id":"0c2a2c6c-5a95-4956-a02d-e1b39c2e5c6e","hostname":"my-android-phone","location":{"map_id":"b4695157-0d1d-4da0-8f9e-5c38149d8b81","site_id":"d14f16d8-d14f-11e5-8e81-1258369c38a9","timestamp":1717027100,"x":423.5,"y":201.0},"memory_total":8589934592,"memory_usage":3758096384,"mfg":"Samsung","model":"Galaxy S23","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","os_type":"Android","os_version":"14","serial":"R5CX123456A","storage_total":128849018880,"storage_usage":52428800000,"timestamp":1717027200,"wifi_band":"5GHz","wifi_bssid":"00:11:22:33:44:55","wifi_channel":36,"wifi_ip":"192.168.1.55","wifi_mac":"00:aa:bb:cc:dd:ee","wifi_rssi":-58,"wifi_ssid":"Corp-WiFi"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
