package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesDevicesWirelessTestGetSiteDeviceRadioChannels tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestGetSiteDeviceRadioChannels(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    countryCode := "US"
    apiResponse, err := sitesDevicesWireless.GetSiteDeviceRadioChannels(ctx, siteId, &countryCode)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"band24_40mhz_allowed":false,"band24_channels":{"20":[1,2,3,4,5,6,7,8,9,10,11],"40":[1,2,3,4,5,6,7,8,9,10,11]},"band24_enabled":true,"band5_channels":{"20":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165],"40":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"80":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"dfs":[52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144],"outdoor":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165]},"band5_enabled":true,"certified":true,"code":840,"dfs_ok":true,"key":"US","name":"United States","uses":"US_FCC"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestGetSiteDeviceIotPort tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestGetSiteDeviceIotPort(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesDevicesWireless.GetSiteDeviceIotPort(ctx, siteId, deviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"A1":1,"DO":0}`
    testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}

// TestSitesDevicesWirelessTestSetSiteDeviceIotPort tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestSetSiteDeviceIotPort(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesDevicesWireless.SetSiteDeviceIotPort(ctx, siteId, deviceId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"A1":1,"DO":0}`
    testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}
