package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesDevicesWiredTestDeleteSiteLocalSwitchPortConfig tests the behavior of the SitesDevicesWired
func TestSitesDevicesWiredTestDeleteSiteLocalSwitchPortConfig(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesDevicesWired.DeleteSiteLocalSwitchPortConfig(ctx, siteId, deviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesDevicesWiredTestUpdateSiteLocalSwitchPortConfig tests the behavior of the SitesDevicesWired
func TestSitesDevicesWiredTestUpdateSiteLocalSwitchPortConfig(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.JunosPortConfig
    errBody := json.Unmarshal([]byte(`{"ae_disable_lacp":true,"ae_idx":0,"aggregated":false,"description":"string","disable_autoneg":true,"duplex":"auto","dynamic_usage":"string","esilag":true,"poe_disabled":true,"speed":"auto","usage":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := sitesDevicesWired.UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
