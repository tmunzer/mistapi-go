package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesClientsWiredTestCountSiteWiredClients tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestCountSiteWiredClients(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteWiredClientsCountDistinctEnum("mac")
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesClientsWired.CountSiteWiredClients(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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

// TestSitesClientsWiredTestSearchSiteWiredClients tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestSearchSiteWiredClients(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesClientsWired.SearchSiteWiredClients(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1648529800.8221116,"limit":1000,"results":[{"device_mac":["001122334455"],"device_mac_port":[{"device_mac":"001122334455","ip":"","port_id":"et-0/0/1","port_parent":"","start":"2020-12-10T00:07:36.262+0000","vlan":1,"when":"2022-03-29T04:56:05.172+0000"}],"ip":["11.216.202.61"],"mac":"112233445566","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","port_id":["et-0/0/1"],"site_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","timestamp":1571174567.807,"vlan":[0,1001]}],"start":1648443400.8221116,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
