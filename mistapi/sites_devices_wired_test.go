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
	var body map[string]models.JunosLocalPortConfig
	errBody := json.Unmarshal([]byte(`{"ge-0/0/0-1":{"poe_disabled":true,"usage":"iot"}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := sitesDevicesWired.UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
