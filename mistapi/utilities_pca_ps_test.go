// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestUtilitiesPCAPsTestListOrgPacketCaptures tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestListOrgPacketCaptures(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := utilitiesPcaPs.ListOrgPacketCaptures(ctx, orgId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1461089816,"limit":100,"next":"/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816","results":[{"ap_macs":["5c5b35000010"],"timestamp":1461869041,"type":"new_assoc","url":"https://..."}],"start":1461099816}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestListOrgPacketCaptures1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestListOrgPacketCaptures1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := utilitiesPcaPs.ListOrgPacketCaptures(ctx, orgId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1461089816,"limit":100,"next":"/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816","results":[{"ap_macs":["5c5b35000010"],"timestamp":1461869041,"type":"new_assoc","url":"https://..."}],"start":1461099816}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStopOrgPacketCapture tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStopOrgPacketCapture(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesPcaPs.StopOrgPacketCapture(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesPCAPsTestGetOrgCapturingStatus tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestGetOrgCapturingStatus(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesPcaPs.GetOrgCapturingStatus(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":["5c5b350e001c","5c5b350e001b"],"client_mac":"60a10a773412","duration":300,"failed":[],"id":"a9a84e13-a714-b1eb-152f-a434416217d5","includes_mcast":false,"max_pkt_len":128,"num_packets":1000,"ok":["5c5b350e001c","5c5b350e001b"],"started_time":1435080709,"type":"client"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestGetOrgCapturingStatus1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestGetOrgCapturingStatus1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesPcaPs.GetOrgCapturingStatus(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":["5c5b350e001c","5c5b350e001b"],"client_mac":"60a10a773412","duration":300,"failed":[],"id":"a9a84e13-a714-b1eb-152f-a434416217d5","includes_mcast":false,"max_pkt_len":128,"num_packets":1000,"ok":["5c5b350e001c","5c5b350e001b"],"started_time":1435080709,"type":"client"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStartOrgPacketCapture tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStartOrgPacketCapture(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CaptureOrg
    errBody := json.Unmarshal([]byte(`{"duration":600,"format":"stream","max_pkt_len":1500,"mxedges":{"00000000-0000-0000-1000-001122334455":{"interfaces":{"port1":{"tcpdump_expression":"udp port 67 or udp port 68"}}}},"num_packets":100,"tcpdump_expression":"vlan 999","type":"mxedge"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesPcaPs.StartOrgPacketCapture(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"ap_count":3,"aps":[],"duration":600,"enabled":true,"expiry":1614886726.5411825,"format":"stream","id":"a9a84e13-a714-b1eb-152f-a434416217d5","include_mcast":false,"max_pkt_len":68,"num_packets":100,"org_id":"a9346fba-f920-e99a-cc51-2e8dcc57fa3c","raw":true,"site_id":"67970e46-4e12-11e6-9188-0242ac110007","ssid":"","timestamp":1614886126.5411825,"type":"radiotap"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStartOrgPacketCapture1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStartOrgPacketCapture1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CaptureOrg
    errBody := json.Unmarshal([]byte(`{"duration":600,"format":"stream","max_pkt_len":1500,"mxedges":{"00000000-0000-0000-1000-001122334455":{"interfaces":{"port1":{"tcpdump_expression":"udp port 67 or udp port 68"}}}},"num_packets":100,"tcpdump_expression":"vlan 999","type":"mxedge"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesPcaPs.StartOrgPacketCapture(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"ap_count":3,"aps":[],"duration":600,"enabled":true,"expiry":1614886726.5411825,"format":"stream","id":"a9a84e13-a714-b1eb-152f-a434416217d5","include_mcast":false,"max_pkt_len":68,"num_packets":100,"org_id":"a9346fba-f920-e99a-cc51-2e8dcc57fa3c","raw":true,"site_id":"67970e46-4e12-11e6-9188-0242ac110007","ssid":"","timestamp":1614886126.5411825,"type":"radiotap"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestListSitePacketCaptures tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestListSitePacketCaptures(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := utilitiesPcaPs.ListSitePacketCaptures(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1461089816,"limit":100,"next":"/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816","results":[{"ap_macs":["5c5b35000010"],"timestamp":1461869041,"type":"new_assoc","url":"https://..."}],"start":1461099816}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestListSitePacketCaptures1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestListSitePacketCaptures1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := utilitiesPcaPs.ListSitePacketCaptures(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1461089816,"limit":100,"next":"/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816","results":[{"ap_macs":["5c5b35000010"],"timestamp":1461869041,"type":"new_assoc","url":"https://..."}],"start":1461099816}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStopSitePacketCapture tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStopSitePacketCapture(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesPcaPs.StopSitePacketCapture(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesPCAPsTestGetSiteCapturingStatus tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestGetSiteCapturingStatus(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesPcaPs.GetSiteCapturingStatus(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":["5c5b350e001c","5c5b350e001b"],"client_mac":"60a10a773412","duration":300,"failed":[],"id":"a9a84e13-a714-b1eb-152f-a434416217d5","includes_mcast":false,"max_pkt_len":128,"num_packets":1000,"ok":["5c5b350e001c","5c5b350e001b"],"started_time":1435080709,"type":"client"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestGetSiteCapturingStatus1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestGetSiteCapturingStatus1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesPcaPs.GetSiteCapturingStatus(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":["5c5b350e001c","5c5b350e001b"],"client_mac":"60a10a773412","duration":300,"failed":[],"id":"a9a84e13-a714-b1eb-152f-a434416217d5","includes_mcast":false,"max_pkt_len":128,"num_packets":1000,"ok":["5c5b350e001c","5c5b350e001b"],"started_time":1435080709,"type":"client"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStartSitePacketCapture tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStartSitePacketCapture(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CaptureSite
    errBody := json.Unmarshal([]byte(`{"client_mac":"60a10a773412","duration":600,"includes_mcast":false,"max_pkt_len":128,"num_packets":100,"type":"new_assoc"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesPcaPs.StartSitePacketCapture(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"ap_count":3,"aps":[],"duration":600,"enabled":true,"expiry":1614886726.5411825,"format":"stream","id":"a9a84e13-a714-b1eb-152f-a434416217d5","include_mcast":false,"max_pkt_len":68,"num_packets":100,"org_id":"a9346fba-f920-e99a-cc51-2e8dcc57fa3c","raw":true,"site_id":"67970e46-4e12-11e6-9188-0242ac110007","ssid":"","timestamp":1614886126.5411825,"type":"radiotap"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesPCAPsTestStartSitePacketCapture1 tests the behavior of the UtilitiesPCAPs
func TestUtilitiesPCAPsTestStartSitePacketCapture1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CaptureSite
    errBody := json.Unmarshal([]byte(`{"client_mac":"60a10a773412","duration":600,"includes_mcast":false,"max_pkt_len":128,"num_packets":100,"type":"new_assoc"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesPcaPs.StartSitePacketCapture(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"ap_count":3,"aps":[],"duration":600,"enabled":true,"expiry":1614886726.5411825,"format":"stream","id":"a9a84e13-a714-b1eb-152f-a434416217d5","include_mcast":false,"max_pkt_len":68,"num_packets":100,"org_id":"a9346fba-f920-e99a-cc51-2e8dcc57fa3c","raw":true,"site_id":"67970e46-4e12-11e6-9188-0242ac110007","ssid":"","timestamp":1614886126.5411825,"type":"radiotap"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
