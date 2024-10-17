package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesVBeaconsTestListSiteVBeacons tests the behavior of the SitesVBeacons
func TestSitesVBeaconsTestListSiteVBeacons(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesVBeacons.ListSiteVBeacons(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","major":0,"map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","message":"string","minor":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":4,"power_mode":"default","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","url":"string","uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","wayfinding_nodename":"string","x":0,"y":0}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesVBeaconsTestCreateSiteVBeacon tests the behavior of the SitesVBeacons
func TestSitesVBeaconsTestCreateSiteVBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Vbeacon
	errBody := json.Unmarshal([]byte(`{"major":0,"message":"string","minor":0,"name":"string","power":4,"power_mode":"default","url":"string","uuid":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","wayfinding_nodename":"string","x":0,"y":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesVBeacons.CreateSiteVBeacon(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","major":0,"map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","message":"string","minor":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":4,"power_mode":"default","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","url":"string","uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","wayfinding_nodename":"string","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesVBeaconsTestDeleteSiteVBeacon tests the behavior of the SitesVBeacons
func TestSitesVBeaconsTestDeleteSiteVBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	vbeaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesVBeacons.DeleteSiteVBeacon(ctx, siteId, vbeaconId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesVBeaconsTestGetSiteVBeacon tests the behavior of the SitesVBeacons
func TestSitesVBeaconsTestGetSiteVBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	vbeaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesVBeacons.GetSiteVBeacon(ctx, siteId, vbeaconId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","major":0,"map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","message":"string","minor":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":4,"power_mode":"default","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","url":"string","uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","wayfinding_nodename":"string","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesVBeaconsTestUpdateSiteVBeacon tests the behavior of the SitesVBeacons
func TestSitesVBeaconsTestUpdateSiteVBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	vbeaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Vbeacon
	errBody := json.Unmarshal([]byte(`{"major":0,"message":"string","minor":0,"name":"string","power":4,"power_mode":"default","url":"string","uuid":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","wayfinding_nodename":"string","x":0,"y":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesVBeacons.UpdateSiteVBeacon(ctx, siteId, vbeaconId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","major":0,"map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","message":"string","minor":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":4,"power_mode":"default","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","url":"string","uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","wayfinding_nodename":"string","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
