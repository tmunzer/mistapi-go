package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesStatsCallsTestCountSiteCalls tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestCountSiteCalls(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distrinct := models.CountSiteCallsDistrinctEnum("mac")
    
    
    
    
    apiResponse, err := sitesStatsCalls.CountSiteCalls(ctx, siteId, &distrinct, nil, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    
    app := "zoom"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsCalls.SearchSiteCalls(ctx, siteId, nil, &app, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsCallsTestListSiteTroubleshootCalls tests the behavior of the SitesStatsCalls
func TestSitesStatsCallsTestListSiteTroubleshootCalls(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    app := "zoom"
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsCalls.ListSiteTroubleshootCalls(ctx, siteId, nil, nil, nil, &app, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"mac":"983a78ea4a44","meeting_id":"b784d744-9a7c-4fad-9af0-f78858a319b1","results":[{"audio_in":{"ap_num_clients":45.483067936367462,"ap_rtt":6.352042701509264,"client_cpu":9.3234525786505813,"client_radio_band":1.5325644097982959E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463192,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.15336827271514469,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.28558229323890472,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.78755196597841048,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.81311179531945121},"audio_out":{"ap_num_clients":45.483067936367462,"ap_rtt":6.352042701509264,"client_cpu":9.3234525786505813,"client_radio_band":1.5325644097982959E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463192,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.15336827271514469,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.28558229323890472,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.78755196597841048,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.81311179531945121},"timestamp":1695425115,"video_in":{"ap_num_clients":45.483067936367462,"ap_rtt":6.352042701509264,"client_cpu":9.3234525786505813,"client_radio_band":1.5325644097982959E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463192,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.15336827271514469,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.28558229323890472,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.78755196597841048,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.81311179531945121},"video_out":{"ap_num_clients":45.483067936367462,"ap_rtt":6.352042701509264,"client_cpu":9.3234525786505813,"client_radio_band":1.5325644097982959E-06,"client_rssi":17.251008563571506,"client_tx_bytes":1.8379177401463192,"client_tx_rates":10.668423069847954,"client_tx_retries":43.323209603627525,"client_vpn_distance":112.4420166015625,"expected":29.74261474609375,"radio_bandwidth":-0.15336827271514469,"radio_channel":0.662909648484654,"radio_util":27.891777674357098,"radio_util_interference":4.38913492154744,"site_num_clients":-0.28558229323890472,"site_wan_avg_upload_mpbs":-0.988989942603641,"site_wan_jitter":0.78755196597841048,"site_wan_rtt":15.094849904378256,"site_wan_upload_mpbs":-0.81311179531945121}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}