// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesStatsCallsTestTroubleshootSiteCall tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestTroubleshootSiteCall(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	meetingId := "1234567890abcdef"
	mac := "001122334455"
	app := "zoom"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsCalls.TroubleshootSiteCall(ctx, siteId, clientMac, meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"983a78ea4a44","meeting_id":"b784d744-9a7c-4fad-9af0-f78858a319b1","results":[{"ap_num_clients":-0.08802365511655807,"ap_rtt":0.09924473613500595,"audio_in":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"audio_out":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"client_cpu":0.00834270566701889,"client_n_streams":0.00734270566701889,"client_radio_band":0.5841414928436279,"client_rssi":0.7594696879386902,"client_rx_bytes":2.365511655807E-05,"client_rx_rates":0.02441493794322014,"client_rx_retries":-0.14325742423534393,"client_tx_bytes":0.00102365511655807,"client_tx_rates":0.22236637771129608,"client_tx_retries":0.3308201730251312,"client_vpn_distance":-0.0001660545531194657,"client_wifi_version":7.0566701889E-07,"expected":-2.8630001056670187,"radio_ap_change":0.01850946433842182,"radio_bandwidth":-0.021175479516386986,"radio_channel":0.11686426401138306,"radio_rx_failed":1.1782013177871704,"radio_tx_power":0.121039018034935,"radio_util":0.2452986091375351,"radio_util_interference":3.4367904663085938,"site_num_clients":0.055026158690452576,"site_wan_avg_download_mbps":3.0566701889E-07,"site_wan_avg_upload_mbps":5.566701889E-08,"site_wan_download_mbps":8.0566701889E-07,"site_wan_jitter":1.00566701889E-06,"site_wan_rtt":4.0566701889E-07,"site_wan_upload_mbps":2.0566701889E-07,"timestamp":1695425115,"video_in":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"video_out":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899}}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsCallsTestTroubleshootSiteCall1 tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestTroubleshootSiteCall1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"
	meetingId := "1234567890abcdef"
	mac := "001122334455"
	app := "zoom"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsCalls.TroubleshootSiteCall(ctx, siteId, clientMac, meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"983a78ea4a44","meeting_id":"b784d744-9a7c-4fad-9af0-f78858a319b1","results":[{"ap_num_clients":-0.08802365511655807,"ap_rtt":0.09924473613500595,"audio_in":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"audio_out":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"client_cpu":0.00834270566701889,"client_n_streams":0.00734270566701889,"client_radio_band":0.5841414928436279,"client_rssi":0.7594696879386902,"client_rx_bytes":2.365511655807E-05,"client_rx_rates":0.02441493794322014,"client_rx_retries":-0.14325742423534393,"client_tx_bytes":0.00102365511655807,"client_tx_rates":0.22236637771129608,"client_tx_retries":0.3308201730251312,"client_vpn_distance":-0.0001660545531194657,"client_wifi_version":7.0566701889E-07,"expected":-2.8630001056670187,"radio_ap_change":0.01850946433842182,"radio_bandwidth":-0.021175479516386986,"radio_channel":0.11686426401138306,"radio_rx_failed":1.1782013177871704,"radio_tx_power":0.121039018034935,"radio_util":0.2452986091375351,"radio_util_interference":3.4367904663085938,"site_num_clients":0.055026158690452576,"site_wan_avg_download_mbps":3.0566701889E-07,"site_wan_avg_upload_mbps":5.566701889E-08,"site_wan_download_mbps":8.0566701889E-07,"site_wan_jitter":1.00566701889E-06,"site_wan_rtt":4.0566701889E-07,"site_wan_upload_mbps":2.0566701889E-07,"timestamp":1695425115,"video_in":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899},"video_out":{"ap_num_clients":-0.6565111,"ap_rtt":0.16559607,"client_cpu":3.7028809,"client_n_streams":0.15803306,"client_radio_band":0.5576923,"client_rssi":-1.0839354,"client_rx_bytes":2.2622051,"client_rx_rates":0.62357205,"client_rx_retries":0.26726437,"client_tx_bytes":0.15803306,"client_tx_rates":0.62357205,"client_tx_retries":0.77553505,"client_vpn_distance":1.6474955,"client_wifi_version":0.18267937,"expected":30.941595,"radio_bandwidth":-0.06538621,"radio_channel":-0.73391086,"radio_tx_power":0.10027129,"radio_util":12.770318,"radio_util_interference":-3.079999,"site_num_clients":0.017364305,"wan_avg_download_mbps":1.4803165,"wan_avg_upload_mbps":-0.038184267,"wan_jitter":5.9680853,"wan_max_download_mbps":1.4803165,"wan_max_upload_mbps":-0.038184267,"wan_rtt":46.77899}}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsCallsTestCountSiteCalls tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestCountSiteCalls(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.CountSiteCallsDistinctEnum("mac")
	rating := int(5)
	app := "zoom"

	limit := int(100)
	apiResponse, err := sitesStatsCalls.CountSiteCalls(ctx, siteId, &distinct, &rating, &app, nil, nil, &limit)
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

// TestSitesStatsCallsTestCountSiteCalls1 tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestCountSiteCalls1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.CountSiteCallsDistinctEnum("mac")
	rating := int(5)
	app := "zoom"

	limit := int(100)
	apiResponse, err := sitesStatsCalls.CountSiteCalls(ctx, siteId, &distinct, &rating, &app, nil, nil, &limit)
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

// TestSitesStatsCallsTestSearchSiteCalls tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestSearchSiteCalls(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	app := "zoom"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesStatsCalls.SearchSiteCalls(ctx, siteId, &mac, &app, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsCallsTestSearchSiteCalls1 tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestSearchSiteCalls1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	app := "zoom"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesStatsCalls.SearchSiteCalls(ctx, siteId, &mac, &app, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsCallsTestGetSiteCallsSummary tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestGetSiteCallsSummary(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apMac := "001122334455"
	app := "zoom"

	apiResponse, err := sitesStatsCalls.GetSiteCallsSummary(ctx, siteId, &apMac, &app, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"bad_minutes":5566,"bad_minutes_client":526,"bad_minutes_site_wan":3612,"bad_minutes_wireless":1428,"num_aps":1,"num_users":3,"total_minutes":575217}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsCallsTestGetSiteCallsSummary1 tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestGetSiteCallsSummary1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apMac := "001122334455"
	app := "zoom"

	apiResponse, err := sitesStatsCalls.GetSiteCallsSummary(ctx, siteId, &apMac, &app, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"bad_minutes":5566,"bad_minutes_client":526,"bad_minutes_site_wan":3612,"bad_minutes_wireless":1428,"num_aps":1,"num_users":3,"total_minutes":575217}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsCallsTestListSiteTroubleshootCalls tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestListSiteTroubleshootCalls(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ap := "001122334455"
	meetingId := "1234567890abcdef"
	mac := "001122334455"
	app := "zoom"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsCalls.ListSiteTroubleshootCalls(ctx, siteId, &ap, &meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"983a78ea4a44","meeting_id":"b784d744-9a7c-4fad-9af0-f78858a319b1","results":[{"ap_num_clients":-0.08802365511655807,"ap_rtt":0.09924473613500595,"audio_in":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"audio_out":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"client_cpu":0.00834270566701889,"client_n_streams":0.00734270566701889,"client_radio_band":0.5841414928436279,"client_rssi":0.7594696879386902,"client_rx_bytes":2.365511655807E-05,"client_rx_rates":0.02441493794322014,"client_rx_retries":-0.14325742423534393,"client_tx_bytes":0.00102365511655807,"client_tx_rates":0.22236637771129608,"client_tx_retries":0.3308201730251312,"client_vpn_distance":-0.0001660545531194657,"client_wifi_version":7.0566701889E-07,"expected":-2.8630001056670187,"radio_ap_change":0.01850946433842182,"radio_bandwidth":-0.021175479516386986,"radio_channel":0.11686426401138306,"radio_rx_failed":1.1782013177871704,"radio_tx_power":0.121039018034935,"radio_util":0.2452986091375351,"radio_util_interference":3.4367904663085938,"site_num_clients":0.055026158690452576,"site_wan_avg_download_mbps":3.0566701889E-07,"site_wan_avg_upload_mbps":5.566701889E-08,"site_wan_download_mbps":8.0566701889E-07,"site_wan_jitter":1.00566701889E-06,"site_wan_rtt":4.0566701889E-07,"site_wan_upload_mbps":2.0566701889E-07,"timestamp":1695425115,"video_in":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"video_out":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512}}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsCallsTestListSiteTroubleshootCalls1 tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestListSiteTroubleshootCalls1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ap := "001122334455"
	meetingId := "1234567890abcdef"
	mac := "001122334455"
	app := "zoom"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesStatsCalls.ListSiteTroubleshootCalls(ctx, siteId, &ap, &meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"983a78ea4a44","meeting_id":"b784d744-9a7c-4fad-9af0-f78858a319b1","results":[{"ap_num_clients":-0.08802365511655807,"ap_rtt":0.09924473613500595,"audio_in":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"audio_out":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"client_cpu":0.00834270566701889,"client_n_streams":0.00734270566701889,"client_radio_band":0.5841414928436279,"client_rssi":0.7594696879386902,"client_rx_bytes":2.365511655807E-05,"client_rx_rates":0.02441493794322014,"client_rx_retries":-0.14325742423534393,"client_tx_bytes":0.00102365511655807,"client_tx_rates":0.22236637771129608,"client_tx_retries":0.3308201730251312,"client_vpn_distance":-0.0001660545531194657,"client_wifi_version":7.0566701889E-07,"expected":-2.8630001056670187,"radio_ap_change":0.01850946433842182,"radio_bandwidth":-0.021175479516386986,"radio_channel":0.11686426401138306,"radio_rx_failed":1.1782013177871704,"radio_tx_power":0.121039018034935,"radio_util":0.2452986091375351,"radio_util_interference":3.4367904663085938,"site_num_clients":0.055026158690452576,"site_wan_avg_download_mbps":3.0566701889E-07,"site_wan_avg_upload_mbps":5.566701889E-08,"site_wan_download_mbps":8.0566701889E-07,"site_wan_jitter":1.00566701889E-06,"site_wan_rtt":4.0566701889E-07,"site_wan_upload_mbps":2.0566701889E-07,"timestamp":1695425115,"video_in":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512},"video_out":{"ap_num_clients":45.48306793636746,"ap_rtt":6.352042701509264,"client_cpu":9.323452578650581,"client_radio_band":1.5325644097982958E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463191,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.1533682727151447,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.2855822932389047,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.7875519659784105,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.8131117953194512}}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
