// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsStatsAssetsTestListOrgAssetsStats tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestListOrgAssetsStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsAssets.ListOrgAssetsStats(ctx, orgId, nil, nil, &duration, &limit, &page)
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

// TestOrgsStatsAssetsTestListOrgAssetsStats1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestListOrgAssetsStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsAssets.ListOrgAssetsStats(ctx, orgId, nil, nil, &duration, &limit, &page)
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

// TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    limit := int(100)
    apiResponse, err := orgsStatsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil, &limit)
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

// TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    limit := int(100)
    apiResponse, err := orgsStatsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil, &limit)
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

// TestOrgsStatsAssetsTestSearchOrgAssets tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestSearchOrgAssets(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsStatsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestOrgsStatsAssetsTestSearchOrgAssets1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestSearchOrgAssets1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsStatsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
