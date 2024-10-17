package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesBeaconsTestListSiteBeacons tests the behavior of the SitesBeacons
func TestSitesBeaconsTestListSiteBeacons(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesBeacons.ListSiteBeacons(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":0,"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"eddystone-uid","x":0,"y":0}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesBeaconsTestCreateSiteBeacon tests the behavior of the SitesBeacons
func TestSitesBeaconsTestCreateSiteBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Beacon
	errBody := json.Unmarshal([]byte(`{"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","name":"string","power":0,"type":"eddystone-uid","x":0,"y":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesBeacons.CreateSiteBeacon(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":0,"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"eddystone-uid","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesBeaconsTestDeleteSiteBeacons tests the behavior of the SitesBeacons
func TestSitesBeaconsTestDeleteSiteBeacons(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	beaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesBeacons.DeleteSiteBeacons(ctx, siteId, beaconId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesBeaconsTestGetSiteBeacon tests the behavior of the SitesBeacons
func TestSitesBeaconsTestGetSiteBeacon(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	beaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesBeacons.GetSiteBeacon(ctx, siteId, beaconId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":0,"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"eddystone-uid","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesBeaconsTestUpdateSiteBeacons tests the behavior of the SitesBeacons
func TestSitesBeaconsTestUpdateSiteBeacons(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	beaconId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Beacon
	errBody := json.Unmarshal([]byte(`{"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","name":"string","power":0,"type":"eddystone-uid","x":0,"y":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesBeacons.UpdateSiteBeacons(ctx, siteId, beaconId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"eddystone_instance":"string","eddystone_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","power":0,"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"eddystone-uid","x":0,"y":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
