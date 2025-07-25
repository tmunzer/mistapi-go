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

// TestSitesServicesTestListSiteServicesDerived tests the behavior of the SitesServices
func TestSitesServicesTestListSiteServicesDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesServices.ListSiteServicesDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"addresses":["string"],"apps":["string"],"dscp":8,"hostnames":["string"],"max_jitter":0,"max_latency":0,"max_loss":0,"name":"string","specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"default","type":"custom"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesServicesTestListSiteServicesDerived1 tests the behavior of the SitesServices
func TestSitesServicesTestListSiteServicesDerived1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesServices.ListSiteServicesDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"addresses":["string"],"apps":["string"],"dscp":8,"hostnames":["string"],"max_jitter":0,"max_latency":0,"max_loss":0,"name":"string","specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"default","type":"custom"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesServicesTestCountSiteServicePathEvents tests the behavior of the SitesServices
func TestSitesServicesTestCountSiteServicePathEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteServiceEventsCountDistinctEnum("type")
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesServices.CountSiteServicePathEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesServicesTestCountSiteServicePathEvents1 tests the behavior of the SitesServices
func TestSitesServicesTestCountSiteServicePathEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteServiceEventsCountDistinctEnum("type")
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesServices.CountSiteServicePathEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesServicesTestSearchSiteServicePathEvents tests the behavior of the SitesServices
func TestSitesServicesTestSearchSiteServicePathEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesServices.SearchSiteServicePathEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1697096379,"limit":10,"results":[{"mac":"90ec7734b374","model":"SSR120","org_id":"a3c6718f-2823-4e48-bf5e-b841768a4c9b","policy":"INTERNET","port_id":"ge-1/0/6","site_id":"4279edbd-1d24-41ea-9505-2eb26c8590fa","text":"Peer Path Down","timestamp":1697037328.651775,"type":"GW_SERVICE_PATH_REMOVE","version":"6.1.5-14.lts","vpn_name":"Syracuse_HUB","vpn_path":"Syracuse_HUB-Wan0"}],"start":1697009979,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesServicesTestSearchSiteServicePathEvents1 tests the behavior of the SitesServices
func TestSitesServicesTestSearchSiteServicePathEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesServices.SearchSiteServicePathEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1697096379,"limit":10,"results":[{"mac":"90ec7734b374","model":"SSR120","org_id":"a3c6718f-2823-4e48-bf5e-b841768a4c9b","policy":"INTERNET","port_id":"ge-1/0/6","site_id":"4279edbd-1d24-41ea-9505-2eb26c8590fa","text":"Peer Path Down","timestamp":1697037328.651775,"type":"GW_SERVICE_PATH_REMOVE","version":"6.1.5-14.lts","vpn_name":"Syracuse_HUB","vpn_path":"Syracuse_HUB-Wan0"}],"start":1697009979,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
