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

// TestSitesDevicesWirelessTestListSiteDeviceRadioChannels tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestListSiteDeviceRadioChannels(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	countryCode := "US"
	apiResponse, err := sitesDevicesWireless.ListSiteDeviceRadioChannels(ctx, siteId, &countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band24_40mhz_allowed":false,"band24_channels":{"20":[1,2,3,4,5,6,7,8,9,10,11],"40":[1,2,3,4,5,6,7,8,9,10,11]},"band24_enabled":true,"band5_channels":{"20":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165],"40":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"80":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"dfs":[52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144],"outdoor":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165]},"band5_enabled":true,"certified":true,"code":840,"dfs_ok":true,"key":"US","name":"United States","uses":"US_FCC"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestListSiteDeviceRadioChannels1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestListSiteDeviceRadioChannels1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	countryCode := "US"
	apiResponse, err := sitesDevicesWireless.ListSiteDeviceRadioChannels(ctx, siteId, &countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"A1":1,"DO":0}`
	testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}

// TestSitesDevicesWirelessTestGetSiteDeviceIotPort1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestGetSiteDeviceIotPort1(t *testing.T) {
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"A1":1,"DO":0}`
	testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}

// TestSitesDevicesWirelessTestSetSiteDeviceIotPort1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestSetSiteDeviceIotPort1(t *testing.T) {
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"A1":1,"DO":0}`
	testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}

// TestSitesDevicesWirelessTestStartSiteDeviceZigbeeEventTrail tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestStartSiteDeviceZigbeeEventTrail(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesDevicesWireless.StartSiteDeviceZigbeeEventTrail(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session":"7a5f7796-83ee-11e5-95c6-1258369c38a9"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestStartSiteDeviceZigbeeEventTrail1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestStartSiteDeviceZigbeeEventTrail1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesDevicesWireless.StartSiteDeviceZigbeeEventTrail(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session":"7a5f7796-83ee-11e5-95c6-1258369c38a9"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestStopSiteDeviceZigbeeJoin tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestStopSiteDeviceZigbeeJoin(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesDevicesWireless.StopSiteDeviceZigbeeJoin(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesDevicesWirelessTestEnableSiteDeviceZigbeeJoin tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestEnableSiteDeviceZigbeeJoin(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UtilsZigbeeJoin
	errBody := json.Unmarshal([]byte(`{"duration":600}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesDevicesWireless.EnableSiteDeviceZigbeeJoin(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session_id":"19e73828-937f-05e6-f709-e29efdb0a82b"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestEnableSiteDeviceZigbeeJoin1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestEnableSiteDeviceZigbeeJoin1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UtilsZigbeeJoin
	errBody := json.Unmarshal([]byte(`{"duration":600}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesDevicesWireless.EnableSiteDeviceZigbeeJoin(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session_id":"19e73828-937f-05e6-f709-e29efdb0a82b"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestKickSiteDeviceZigbeeClients tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestKickSiteDeviceZigbeeClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UtilsZigbeeKick
	errBody := json.Unmarshal([]byte(`{"macs":["00177a01060cae9f","00177a01060caea1"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := sitesDevicesWireless.KickSiteDeviceZigbeeClients(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesDevicesWirelessTestStartSiteDeviceZigbeePacketTrail tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestStartSiteDeviceZigbeePacketTrail(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesDevicesWireless.StartSiteDeviceZigbeePacketTrail(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session":"7a5f7796-83ee-11e5-95c6-1258369c38a9"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWirelessTestStartSiteDeviceZigbeePacketTrail1 tests the behavior of the SitesDevicesWireless
func TestSitesDevicesWirelessTestStartSiteDeviceZigbeePacketTrail1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesDevicesWireless.StartSiteDeviceZigbeePacketTrail(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"session":"7a5f7796-83ee-11e5-95c6-1258369c38a9"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
