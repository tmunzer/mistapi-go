package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesZonesTestListSiteZones tests the behavior of the SitesZones
func TestSitesZonesTestListSiteZones(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesZones.ListSiteZones(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vertices":[{"x":0,"y":0}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesZonesTestCreateSiteZone tests the behavior of the SitesZones
func TestSitesZonesTestCreateSiteZone(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Zone
    errBody := json.Unmarshal([]byte(`{"name":"string","vertices":[{"x":0,"y":0}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesZones.CreateSiteZone(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vertices":[{"x":0,"y":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesZonesTestDeleteSiteZone tests the behavior of the SitesZones
func TestSitesZonesTestDeleteSiteZone(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesZones.DeleteSiteZone(ctx, siteId, zoneId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesZonesTestGetSiteZone tests the behavior of the SitesZones
func TestSitesZonesTestGetSiteZone(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesZones.GetSiteZone(ctx, siteId, zoneId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vertices":[{"x":0,"y":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesZonesTestUpdateSiteZone tests the behavior of the SitesZones
func TestSitesZonesTestUpdateSiteZone(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    zoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Zone
    errBody := json.Unmarshal([]byte(`{"name":"string","vertices":[{"x":0,"y":0}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesZones.UpdateSiteZone(ctx, siteId, zoneId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vertices":[{"x":0,"y":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
