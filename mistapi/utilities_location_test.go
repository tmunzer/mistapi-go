// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestUtilitiesLocationTestSendSiteDevicesArbitraryBleBeacon tests the behavior of the UtilitiesLocation
func TestUtilitiesLocationTestSendSiteDevicesArbitraryBleBeacon(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UtilsSendBleBeacon
    errBody := json.Unmarshal([]byte(`{"beacon_frame":"68b329da9893e34099c7d8ad5cb9c940","beacon_freq":100,"duration":10,"macs":["5c5b35584a6f","5c5b350ea3b3"],"map_ids":["845a23bf-bed9-e43c-4c86-6fa474be7ae5"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesLocation.SendSiteDevicesArbitraryBleBeacon(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
