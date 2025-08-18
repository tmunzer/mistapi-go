// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsClientsWirelessTestListSiteWirelessClientsStats tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestListSiteWirelessClientsStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wired := bool(false)
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesStatsClientsWireless.ListSiteWirelessClientsStats(ctx, siteId, &wired, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestListSiteWirelessClientsStats1 tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestListSiteWirelessClientsStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wired := bool(false)
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesStatsClientsWireless.ListSiteWirelessClientsStats(ctx, siteId, &wired, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestGetSiteWirelessClientStats tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestGetSiteWirelessClientStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	wired := bool(false)
	apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientStats(ctx, siteId, clientMac, &wired)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestGetSiteWirelessClientStats1 tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestGetSiteWirelessClientStats1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	wired := bool(false)
	apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientStats(ctx, siteId, clientMac, &wired)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestGetSiteWirelessClientsStatsByMap tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestGetSiteWirelessClientsStatsByMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestGetSiteWirelessClientsStatsByMap1 tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestGetSiteWirelessClientsStatsByMap1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"annotation":"unknown","ap_id":"00000000-0000-0000-1000-5c5b35963d70","ap_mac":"5c5b358e6fea","assoc_time":1741152905,"band":"5","bssid":"5c5b358298f2","channel":157,"dual_band":true,"family":"","group":"","hostname":"android-9b228dc33690","idle_time":5,"ip":"10.100.0.47","is_guest":false,"key_mgmt":"WPA3-SAE-FT/CCMP","last_seen":1741257505,"mac":"dadbfc123456","manufacture":"Unknown","map_id":"ed7a0a4e-8835-4c94-ba78-6c1169c9f135","model":"","num_locating_aps":2,"os":"Android 10","proto":"ac","rssi":-39,"rx_bps":0,"rx_bytes":14451780,"rx_pkts":44175,"rx_rate":6,"rx_retries":2010,"site_id":"96c348a9-d6d7-4732-b4f5-23350a2843cd","snr":47,"ssid":"Live_demo_only","tx_bps":0,"tx_bytes":56364072,"tx_pkts":43685,"tx_rate":173.3,"tx_retries":5413,"uptime":104600,"vlan_id":"1","wlan_id":"497fc18a-79b5-405a-bf5a-192eed31ea60","x":695.3357339330526,"x_m":24.086588,"y":760.6746524247893,"y_m":26.349943}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsClientsWirelessTestListSiteUnconnectedClientStats tests the behavior of the SitesStatsClientsWireless
func TestSitesStatsClientsWirelessTestListSiteUnconnectedClientStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStatsClientsWireless.ListSiteUnconnectedClientStats(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ap_mac":"5c5b350e0410","last_seen":1428939600,"mac":"5684dae9ac8b","manufacture":"Apple","rssi":-75,"x":60,"y":80}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
