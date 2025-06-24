package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesDevicesOthersTestListSiteOtherDevices tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestListSiteOtherDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevicesOthers.ListSiteOtherDevices(ctx, siteId, nil, nil, nil, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1676983730,"device_mac":"001122334455","id":"ae9dee49-69e7-4710-a114-5b827a777738","mac":"5c5b35000018","model":"AP41","modified_time":1676983730,"name":"hallway","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","vendor":"cradlepoint"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesOthersTestListSiteOtherDevices1 tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestListSiteOtherDevices1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevicesOthers.ListSiteOtherDevices(ctx, siteId, nil, nil, nil, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1676983730,"device_mac":"001122334455","id":"ae9dee49-69e7-4710-a114-5b827a777738","mac":"5c5b35000018","model":"AP41","modified_time":1676983730,"name":"hallway","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","vendor":"cradlepoint"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesOthersTestCountSiteOtherDeviceEvents tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestCountSiteOtherDeviceEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteOtherDeviceEventsCountDistinctEnum("mac")
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesDevicesOthers.CountSiteOtherDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, &duration, &limit)
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

// TestSitesDevicesOthersTestCountSiteOtherDeviceEvents1 tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestCountSiteOtherDeviceEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteOtherDeviceEventsCountDistinctEnum("mac")
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesDevicesOthers.CountSiteOtherDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, &duration, &limit)
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

// TestSitesDevicesOthersTestSearchSiteOtherDeviceEvents tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestSearchSiteOtherDeviceEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesDevicesOthers.SearchSiteOtherDeviceEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":{"device_mac":"string","mac":"5c5b351e13b5","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Plugged: The Internal 5GB (SIM1) has been inserted into Internal 1.","timestamp":547235620.89,"type":"CELLULAR_EDGE_MODEM_WAN_PLUGGED","vendor":"cradlepoint"},"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesOthersTestSearchSiteOtherDeviceEvents1 tests the behavior of the SitesDevicesOthers
func TestSitesDevicesOthersTestSearchSiteOtherDeviceEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesDevicesOthers.SearchSiteOtherDeviceEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":{"device_mac":"string","mac":"5c5b351e13b5","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Plugged: The Internal 5GB (SIM1) has been inserted into Internal 1.","timestamp":547235620.89,"type":"CELLULAR_EDGE_MODEM_WAN_PLUGGED","vendor":"cradlepoint"},"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
