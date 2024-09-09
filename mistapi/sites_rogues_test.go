package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesRoguesTestListSiteRogueAPs tests the behavior of the SitesRogues
func TestSitesRoguesTestListSiteRogueAPs(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    limit := int(100)
    
    
    duration := "1d"
    interval := "10m"
    apiResponse, err := sitesRogues.ListSiteRogueAPs(ctx, siteId, nil, &limit, nil, nil, &duration, &interval)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1428954000,"limit":100,"next":"/api/v1/sites/a3eda150-ab3f-11e4-aa18-13e21dd250cc/rogues?start=1498482000&end=1498485600&limit=10&interval=1h&type=others","results":[{"ap_mac":"5c5b350e021c","avg_rssi":-72,"bssid":"d8-97-ba-76-b5-aa","channel":"11","num_aps":4,"ssid":"xfinitywifi","times_heard":8}],"start":1428939600}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRoguesTestListSiteRogueClients tests the behavior of the SitesRogues
func TestSitesRoguesTestListSiteRogueClients(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    
    
    duration := "1d"
    interval := "10m"
    apiResponse, err := sitesRogues.ListSiteRogueClients(ctx, siteId, &limit, nil, nil, &duration, &interval)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1428954000,"limit":100,"next":"/api/v1/sites/a3eda150-ab3f-11e4-aa18-13e21dd250cc/rogues/clients?start=1498482000&end=1498485600&limit=10&interval=1h","results":[{"annotation":"whitelist","ap_mac":"5c-5b-35-0e-02-1c","avg_rssi":-63.9,"band":"5","bssid":"d8-97-ba-76-b5-aa","client_mac":"34-f8-32-13-57-c2","num_aps":2}],"start":1428939600}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRoguesTestCountSiteRogueEvents tests the behavior of the SitesRogues
func TestSitesRoguesTestCountSiteRogueEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteRogueEventsCountDistinctEnum("bssid")
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesRogues.CountSiteRogueEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestSitesRoguesTestSearchSiteRogueEvents tests the behavior of the SitesRogues
func TestSitesRoguesTestSearchSiteRogueEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesRogues.SearchSiteRogueEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1538074800,"limit":10,"results":[{"ap":"5c5b350e10030","bssid":"38ff363c8c4c","channel":136,"rssi":-54,"ssid":"MyHomeNetwork","timestamp":1538074612},{"ap":"5c5b350e10030","bssid":"60d02c2394cc","channel":11,"rssi":-59,"ssid":"Home-Office","timestamp":1538074612}],"start":1538071200,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRoguesTestGetSiteRogueAP tests the behavior of the SitesRogues
func TestSitesRoguesTestGetSiteRogueAP(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    rogueBssid := "0000000000ab"
    apiResponse, err := sitesRogues.GetSiteRogueAP(ctx, siteId, rogueBssid)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"manufacture":"Intel Corporate","seen_as_client":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
