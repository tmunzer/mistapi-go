package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesStatsAssetsTestListSiteAssetsStats tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestListSiteAssetsStats(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsAssets.ListSiteAssetsStats(ctx, siteId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"478f6eca-6276-4993-bfeb-5bcbbbbacf08","since":0}],"x":0,"y":0,"zones":[{"id":"477f6eca-6276-4993-bfeb-5ccbbbbadf08","since":0}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestListSiteAssetsStats1 tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestListSiteAssetsStats1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsAssets.ListSiteAssetsStats(ctx, siteId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"478f6eca-6276-4993-bfeb-5bcbbbbacf08","since":0}],"x":0,"y":0,"zones":[{"id":"477f6eca-6276-4993-bfeb-5ccbbbbadf08","since":0}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestCountSiteAssets tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestCountSiteAssets(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteAssetsCountDistinctEnum("map_id")
    limit := int(100)
    apiResponse, err := sitesStatsAssets.CountSiteAssets(ctx, siteId, &distinct, &limit)
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

// TestSitesStatsAssetsTestCountSiteAssets1 tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestCountSiteAssets1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteAssetsCountDistinctEnum("map_id")
    limit := int(100)
    apiResponse, err := sitesStatsAssets.CountSiteAssets(ctx, siteId, &distinct, &limit)
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

// TestSitesStatsAssetsTestSearchSiteAssets tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestSearchSiteAssets(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mac := "001122334455"
    mapId := "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ibeaconUuid := "3ce2ef69-4414-469d-9d55-3ec7fcc38520"
    ibeaconMajor := int(1)
    ibeaconMinor := int(1)
    eddystoneUidNamespace := "1234567890abcdef1234567890abcdef"
    eddystoneUidInstance := "1234567890abcdef1234567890abcdef"
    eddystoneUrl := "https://example.com"
    deviceName := "Device Name"
    by := "mac"
    name := "Asset Name"
    apMac := "001122334455"
    beam := "0"
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsAssets.SearchSiteAssets(ctx, siteId, &mac, &mapId, &ibeaconUuid, &ibeaconMajor, &ibeaconMinor, &eddystoneUidNamespace, &eddystoneUidInstance, &eddystoneUrl, &deviceName, &by, &name, &apMac, &beam, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"next":"string","results":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"476f6eca-6276-4993-bfeb-5dcbbbbaef08","since":0}],"x":0,"y":0,"zones":[{"id":"475f6eca-6276-4993-bfeb-5ecbbbbf6f08","since":0}]}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestSearchSiteAssets1 tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestSearchSiteAssets1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mac := "001122334455"
    mapId := "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ibeaconUuid := "3ce2ef69-4414-469d-9d55-3ec7fcc38520"
    ibeaconMajor := int(1)
    ibeaconMinor := int(1)
    eddystoneUidNamespace := "1234567890abcdef1234567890abcdef"
    eddystoneUidInstance := "1234567890abcdef1234567890abcdef"
    eddystoneUrl := "https://example.com"
    deviceName := "Device Name"
    by := "mac"
    name := "Asset Name"
    apMac := "001122334455"
    beam := "0"
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsAssets.SearchSiteAssets(ctx, siteId, &mac, &mapId, &ibeaconUuid, &ibeaconMajor, &ibeaconMinor, &eddystoneUidNamespace, &eddystoneUidInstance, &eddystoneUrl, &deviceName, &by, &name, &apMac, &beam, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"next":"string","results":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"476f6eca-6276-4993-bfeb-5dcbbbbaef08","since":0}],"x":0,"y":0,"zones":[{"id":"475f6eca-6276-4993-bfeb-5ecbbbbf6f08","since":0}]}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestGetSiteAssetStats tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestGetSiteAssetStats(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := sitesStatsAssets.GetSiteAssetStats(ctx, siteId, assetId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"480f6eca-6276-4993-bfeb-59cbbbbaaf08","since":0}],"x":0,"y":0,"zones":[{"id":"479f6eca-6276-4993-bfeb-5acbbbbabf08","since":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestGetSiteAssetStats1 tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestGetSiteAssetStats1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := sitesStatsAssets.GetSiteAssetStats(ctx, siteId, assetId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"480f6eca-6276-4993-bfeb-59cbbbbaaf08","since":0}],"x":0,"y":0,"zones":[{"id":"479f6eca-6276-4993-bfeb-5acbbbbabf08","since":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestListSiteDiscoveredAssets tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestListSiteDiscoveredAssets(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsAssets.ListSiteDiscoveredAssets(ctx, siteId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesStatsAssetsTestGetSiteAssetsOfInterest tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestGetSiteAssetsOfInterest(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    duration := "1d"
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsAssets.GetSiteAssetsOfInterest(ctx, siteId, &duration, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"ap_mac":"string","beam":0,"by":"string","curr_site":"string","device_name":"string","id":"string","last_seen":0,"mac":"string","manufacture":"string","map_id":"string","name":"string","rssi":0}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestGetSiteAssetsOfInterest1 tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestGetSiteAssetsOfInterest1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    duration := "1d"
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsAssets.GetSiteAssetsOfInterest(ctx, siteId, &duration, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"ap_mac":"string","beam":0,"by":"string","curr_site":"string","device_name":"string","id":"string","last_seen":0,"mac":"string","manufacture":"string","map_id":"string","name":"string","rssi":0}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsAssetsTestGetSiteDiscoveredAssetByMap tests the behavior of the SitesStatsAssets
func TestSitesStatsAssetsTestGetSiteDiscoveredAssetByMap(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesStatsAssets.GetSiteDiscoveredAssetByMap(ctx, siteId, mapId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
