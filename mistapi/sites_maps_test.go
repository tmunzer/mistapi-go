package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesMapsTestListSiteMaps tests the behavior of the SitesMaps
func TestSitesMapsTestListSiteMaps(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesMaps.ListSiteMaps(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"flags":{},"height":0,"height_m":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","latlng_br":{"lat":"string","lng":"string"},"latlng_tl":{"lat":"string","lng":"string"},"locked":true,"modified_time":0,"name":"string","occupancy_limit":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","orientation":0,"origin_x":0,"origin_y":0,"ppm":0,"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitesurvey_path":[{"coordinate":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]}],"thumbnail_url":"string","type":"image","url":"string","view":"roadmap","wall_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"wayfinding":{"micello":{"account_key":"string","default_level_id":0,"map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},"snap_to_path":true},"wayfinding_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"width":0,"width_m":0}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsTestCreateSiteMap tests the behavior of the SitesMaps
func TestSitesMapsTestCreateSiteMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Map
	errBody := json.Unmarshal([]byte(`{"flags":{},"height":0,"height_m":0,"latlng_br":{"lat":"string","lng":"string"},"latlng_tl":{"lat":"string","lng":"string"},"locked":true,"name":"string","occupancy_limit":0,"orientation":0,"origin_x":0,"origin_y":0,"ppm":0,"sitesurvey_path":[{"coordinate":"string","name":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]}],"type":"image","view":"roadmap","wall_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"wayfinding":{"micello":{"account_key":"string","default_level_id":0},"snap_to_path":true},"wayfinding_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"width":0,"width_m":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesMaps.CreateSiteMap(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesMapsTestImportSiteMaps tests the behavior of the SitesMaps
func TestSitesMapsTestImportSiteMaps(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	autoDeviceprofileAssignment := bool(true)

	apiResponse, err := sitesMaps.ImportSiteMaps(ctx, siteId, &autoDeviceprofileAssignment, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aps":[{"action":"assigned-placed","floorplan_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","height":3,"mac":"5c5b35000001","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","orientation":45}],"floorplans":[{"action":"imported","id":"cbdb7f0b-3be0-4872-88f9-58790b509c23","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","name":"map1"}],"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","summary":{"num_ap_assigned":1,"num_inv_assigned":1,"num_map_assigned":1}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsTestDeleteSiteMap tests the behavior of the SitesMaps
func TestSitesMapsTestDeleteSiteMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesMaps.DeleteSiteMap(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsTestGetSiteMap tests the behavior of the SitesMaps
func TestSitesMapsTestGetSiteMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesMaps.GetSiteMap(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesMapsTestUpdateSiteMap tests the behavior of the SitesMaps
func TestSitesMapsTestUpdateSiteMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Map
	errBody := json.Unmarshal([]byte(`{"flags":{},"height":0,"height_m":0,"latlng_br":{"lat":"string","lng":"string"},"latlng_tl":{"lat":"string","lng":"string"},"locked":true,"name":"string","occupancy_limit":0,"orientation":0,"origin_x":0,"origin_y":0,"ppm":0,"sitesurvey_path":[{"coordinate":"string","name":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]}],"type":"image","view":"roadmap","wall_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"wayfinding":{"micello":{"account_key":"string","default_level_id":0},"snap_to_path":true},"wayfinding_path":{"coordinate":"string","nodes":[{"edges":{"N2":"string"},"name":"string","position":{"x":0,"y":0}}]},"width":0,"width_m":0}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesMaps.UpdateSiteMap(ctx, siteId, mapId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesMapsTestDeleteSiteMapImage tests the behavior of the SitesMaps
func TestSitesMapsTestDeleteSiteMapImage(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesMaps.DeleteSiteMapImage(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsTestBulkAssignSiteApsToMap tests the behavior of the SitesMaps
func TestSitesMapsTestBulkAssignSiteApsToMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["5c5b35000001","5c5b35584a6f"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesMaps.BulkAssignSiteApsToMap(ctx, siteId, mapId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"locked":["5c5b35584a6f"],"moved":["5c5b35000001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsTestImportSiteWayfindings tests the behavior of the SitesMaps
func TestSitesMapsTestImportSiteWayfindings(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := sitesMaps.ImportSiteWayfindings(ctx, siteId, mapId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
