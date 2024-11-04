package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesSyntheticTestsTestStartSiteSwitchRadiusSyntheticTest tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestStartSiteSwitchRadiusSyntheticTest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SynthetictestRadiusServer
    errBody := json.Unmarshal([]byte(`{"password":"string","profile":"dot1x","user":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesSyntheticTests.StartSiteSwitchRadiusSyntheticTest(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesSyntheticTestsTestGetSiteDeviceSyntheticTest tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestGetSiteDeviceSyntheticTest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesSyntheticTests.GetSiteDeviceSyntheticTest(ctx, siteId, deviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"device_type":"gateway","mac":"5c5b35584a6f","port_id":"ge-0/0/1.100","start_time":1675718807,"status":"inprogress","type":"speedtest"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSyntheticTestsTestTriggerSiteDeviceSyntheticTest tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestTriggerSiteDeviceSyntheticTest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := sitesSyntheticTests.TriggerSiteDeviceSyntheticTest(ctx, siteId, deviceId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesSyntheticTestsTestGetSiteSyntheticTestStatus tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestGetSiteSyntheticTestStatus(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesSyntheticTests.GetSiteSyntheticTestStatus(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"device_type":"gateway","mac":"5c5b35584a6f","port_id":"ge-0/0/1.100","start_time":1675718807,"status":"inprogress","type":"speedtest"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSyntheticTestsTestTriggerSiteSyntheticTest tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestTriggerSiteSyntheticTest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Synthetictest
    errBody := json.Unmarshal([]byte(`{"email":"test@mist.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesSyntheticTests.TriggerSiteSyntheticTest(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"a42775f6-edc8-69b5-f979-542fa1b43ff9","message":"Successfully queued synthetic test for the site.","status":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSyntheticTestsTestSearchSiteSyntheticTest tests the behavior of the SitesSyntheticTests
func TestSitesSyntheticTestsTestSearchSiteSyntheticTest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    apiResponse, err := sitesSyntheticTests.SearchSiteSyntheticTest(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"next":"string","results":[{"by":"user","device_type":"gateway","failed":false,"latency":40,"mac":"aff827549235","port_id":"ge-0/0/2","rx_mbps":322,"timestamp":1706824045.059036,"tx_mbps":199,"type":"speedtest","vlan_id":20},{"by":"marvis","device_type":"gateway","failed":true,"latency":0,"mac":"8396cd006c8c","port_id":"ge-0/0/2","reason":"interface not ready to perform test","rx_mbps":0,"timestamp":1706824045.059036,"tx_mbps":0,"type":"speedtest","vlan_id":100}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
