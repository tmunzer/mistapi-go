// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesMapsAutoZoneTestDeleteSiteMapAutoZone tests the behavior of the SitesMapsAutoZone
func TestSitesMapsAutoZoneTestDeleteSiteMapAutoZone(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMapsAutoZone.DeleteSiteMapAutoZone(ctx, mapId, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoZoneTestGetSiteMapAutoZoneStatus tests the behavior of the SitesMapsAutoZone
func TestSitesMapsAutoZoneTestGetSiteMapAutoZoneStatus(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesMapsAutoZone.GetSiteMapAutoZoneStatus(ctx, mapId, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"status":"awaiting_review","zones":[{"name":"zone1","vertices":[{"x":0,"y":0},{"x":0,"y":10},{"x":10,"y":10},{"x":10,"y":0}]},{"name":"zone2","vertices":[{"x":0,"y":0},{"x":0,"y":20},{"x":20,"y":20},{"x":20,"y":0}]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsAutoZoneTestGetSiteMapAutoZoneStatus1 tests the behavior of the SitesMapsAutoZone
func TestSitesMapsAutoZoneTestGetSiteMapAutoZoneStatus1(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesMapsAutoZone.GetSiteMapAutoZoneStatus(ctx, mapId, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"status":"awaiting_review","zones":[{"name":"zone1","vertices":[{"x":0,"y":0},{"x":0,"y":10},{"x":10,"y":10},{"x":10,"y":0}]},{"name":"zone2","vertices":[{"x":0,"y":0},{"x":0,"y":20},{"x":20,"y":20},{"x":20,"y":0}]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsAutoZoneTestStartSiteMapAutoZone tests the behavior of the SitesMapsAutoZone
func TestSitesMapsAutoZoneTestStartSiteMapAutoZone(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMapsAutoZone.StartSiteMapAutoZone(ctx, mapId, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
