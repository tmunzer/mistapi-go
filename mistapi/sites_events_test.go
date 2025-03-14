package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesEventsTestListSiteRoamingEvents tests the behavior of the SitesEvents
func TestSitesEventsTestListSiteRoamingEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesEvents.ListSiteRoamingEvents(ctx, siteId, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1501023379,"limit":2,"next":"/api/v1/sites/dca0a44b-324c-11e6-a776-0243ad110007/events/fast_roam?type=success&start=1428939600&end=1428949600&limit=200&token=AAAAEgAIAAVVJh4hF8AAAARzc2lkAH%2F%2F%2F%2F0%3D","results":[{"ap_mac":"5c5b350e040b","client_mac":"dc2b2a3fb13d","fromap":"5c5b350e0569","latency":0.1874195,"ssid":"marvis_test","subtype":"CLIENT_AUTHENTICATED_11R","timestamp":1501000002283782}],"start":1500940800}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesEventsTestCountSiteSystemEvents tests the behavior of the SitesEvents
func TestSitesEventsTestCountSiteSystemEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteSystemEventsCountDistinctEnum("type")
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesEvents.CountSiteSystemEvents(ctx, siteId, &distinct, nil, &limit, nil, nil, &duration)
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

// TestSitesEventsTestSearchSiteSystemEvents tests the behavior of the SitesEvents
func TestSitesEventsTestSearchSiteSystemEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesEvents.SearchSiteSystemEvents(ctx, siteId, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"next":"string","results":[{"ap":"5c5b351e13b5","apfw":"5c5b351e13b5","model":"BT11-WW","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Succeeding DNS query from 172.29.101.134 to 172.29.101.7 for \"portal.mistsys.com\" on vlan 1, id 60224","timestamp":1547235620.89,"type":"CLIENT_DNS_OK"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
