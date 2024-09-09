package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesStatsDevicesTestGetSiteAllClientsStatsByDevice tests the behavior of the SitesStatsDevices
func TestSitesStatsDevicesTestGetSiteAllClientsStatsByDevice(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesStatsDevices.GetSiteAllClientsStatsByDevice(ctx, siteId, deviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsDevicesTestGetSiteGatewayMetrics tests the behavior of the SitesStatsDevices
func TestSitesStatsDevicesTestGetSiteGatewayMetrics(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesStatsDevices.GetSiteGatewayMetrics(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_success":99.9,"version_compliance":{"major_version":{"SRX320":{"major_count":0,"major_version":"19.4R2-S1.2"}},"score":99.9,"type":"gateway"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsDevicesTestGetSiteSwitchesMetrics tests the behavior of the SitesStatsDevices
func TestSitesStatsDevicesTestGetSiteSwitchesMetrics(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    apiResponse, err := sitesStatsDevices.GetSiteSwitchesMetrics(ctx, siteId, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"active_ports_summary":{"details":{"active_port_count":4,"total_port_count":4},"score":100,"total_switch_count":2},"config_success":{"details":{"config_success_count":2},"score":100,"total_switch_count":2},"version_compliance":{"details":{"major_versions":[{"major_count":1,"major_version":"21.4R3.5","model":"EX2300-C-12P","system_names":[]},{"major_count":1,"major_version":"6.0.4-11","model":"SSR120","system_names":[]}]},"score":100,"total_switch_count":2}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
