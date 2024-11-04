package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesStatsDiscoveredSwitchesTestSearchSiteDiscoveredSwitchesMetrics tests the behavior of the SitesStatsDiscoveredSwitches
func TestSitesStatsDiscoveredSwitchesTestSearchSiteDiscoveredSwitchesMetrics(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    scope := models.DiscoveredSwitchesMetricScopeEnum("site")
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitchesMetrics(ctx, siteId, &scope, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1675193686.0191767,"limit":1,"next":"/api/v1/sites/f5fcbee5-fbca-45b3-8bf1-1619ede87879/stats/discovered_switch_metrics/search?end=1675193686.0191767&limit=1&search_after=%5B1675193400000%5D&start=1675107286.0191767","results":[{"details":{},"org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4","scope":"site","score":100,"site_id":"f5fcbee5-fbca-45b3-8bf1-1619ede87879","timestamp":1675193400,"type":"inactive_wired_vlans"}],"start":1675107286.0191767,"total":3}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsDiscoveredSwitchesTestCountSiteDiscoveredSwitches tests the behavior of the SitesStatsDiscoveredSwitches
func TestSitesStatsDiscoveredSwitchesTestCountSiteDiscoveredSwitches(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteDiscoveredSwitchesCountDistinctEnum("system_name")
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsDiscoveredSwitches.CountSiteDiscoveredSwitches(ctx, siteId, &distinct, nil, nil, &duration, &limit, &page)
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

// TestSitesStatsDiscoveredSwitchesTestGetSiteDiscoveredSwitchesMetrics tests the behavior of the SitesStatsDiscoveredSwitches
func TestSitesStatsDiscoveredSwitchesTestGetSiteDiscoveredSwitchesMetrics(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    apiResponse, err := sitesStatsDiscoveredSwitches.GetSiteDiscoveredSwitchesMetrics(ctx, siteId, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"inactive_wired_vlans":{"details":{},"score":100},"poe_compliance":{"details":{"total_aps":63,"total_power":981500},"score":100},"switch_ap_affinity":{"details":{"system_name":["mist-lab-ex2300c","switch1"],"threshold":12},"score":33.3333},"version_compliance":{"details":{"major_versions":[{"major_count":2,"model":"EX2300-C-12P","system_names":["switch1","mist-lab-ex2300c"]},{"major_count":1,"model":"EX4300-48P","system_names":[]}],"total_switch_count":5},"score":75}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsDiscoveredSwitchesTestSearchSiteDiscoveredSwitches tests the behavior of the SitesStatsDiscoveredSwitches
func TestSitesStatsDiscoveredSwitchesTestSearchSiteDiscoveredSwitches(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitches(ctx, siteId, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1604496474.8978634,"limit":1000,"results":[{"aps":[{"hostname":"ap41nearlab","inactive_wired_vlans":[],"mac":"5c5b352e2001","poe_status":true,"when":"2019-06-13T19:53:16.870+0000"}],"mgmt_addr":"10.1.1.1","model":"EX2300-C-12P","org_id":"6748cfa6-4e12-11e6-9188-0242ac110007","site_id":"67970e46-4e12-11e6-9188-0242ac110007","system_desc":"Juniper Networks, Inc. ex2300-c-12p Ethernet Switch, kernel JUNOS 18.2R2.6, Build date: 2018-12-07 13:19:04 UTC Copyright (c) 1996-2018 Juniper Networks, Inc.","system_name":"mist-lab-ex2300c","timestamp":1560457177.037,"vendor":"Juniper Networks","version":"18.2R2.6"}],"start":1604410074.8978484,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
