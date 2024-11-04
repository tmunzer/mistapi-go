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
