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

// TestSitesLocationTestGetSiteBeamCoverageOverview tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteBeamCoverageOverview(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"

	duration := "1d"
	resolution := models.ResolutionEnum("default")

	apiResponse, err := sitesLocation.GetSiteBeamCoverageOverview(ctx, siteId, &mapId, nil, nil, &duration, &resolution, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"beams_means":[[1,3,3.2],[6,10,6.5]],"end":1428954000,"gridsize":1,"result_def":["x","y","beams_mean","beacons_mean","max_rssi","avg_rssi"],"results":[[1,3,3.2,18.5,-68,-70],[6,10,6.5,30,1,-72.5,-75]],"start":1428939600}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestGetSiteBeamCoverageOverview1 tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteBeamCoverageOverview1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"

	duration := "1d"
	resolution := models.ResolutionEnum("default")

	apiResponse, err := sitesLocation.GetSiteBeamCoverageOverview(ctx, siteId, &mapId, nil, nil, &duration, &resolution, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"beams_means":[[1,3,3.2],[6,10,6.5]],"end":1428954000,"gridsize":1,"result_def":["x","y","beams_mean","beacons_mean","max_rssi","avg_rssi"],"results":[[1,3,3.2,18.5,-68,-70],[6,10,6.5,30,1,-72.5,-75]],"start":1428939600}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestGetSiteMachineLearningCurrentStat tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteMachineLearningCurrentStat(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"
	apiResponse, err := sitesLocation.GetSiteMachineLearningCurrentStat(ctx, siteId, &mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestGetSiteMachineLearningCurrentStat1 tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteMachineLearningCurrentStat1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"
	apiResponse, err := sitesLocation.GetSiteMachineLearningCurrentStat(ctx, siteId, &mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestGetSiteDefaultPlfForModels tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteDefaultPlfForModels(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesLocation.GetSiteDefaultPlfForModels(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestGetSiteDefaultPlfForModels1 tests the behavior of the SitesLocation
func TestSitesLocationTestGetSiteDefaultPlfForModels1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesLocation.GetSiteDefaultPlfForModels(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestClearSiteMlOverwriteForDevice tests the behavior of the SitesLocation
func TestSitesLocationTestClearSiteMlOverwriteForDevice(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesLocation.ClearSiteMlOverwriteForDevice(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesLocationTestOverwriteSiteMlForDevice tests the behavior of the SitesLocation
func TestSitesLocationTestOverwriteSiteMlForDevice(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body map[string]models.MlOverwriteAdditionalProperties
	errBody := json.Unmarshal([]byte(`{"iOS":{"int":6,"ple":-3},"iPod":{"int":-10,"ple":-5}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesLocation.OverwriteSiteMlForDevice(ctx, siteId, deviceId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestOverwriteSiteMlForDevice1 tests the behavior of the SitesLocation
func TestSitesLocationTestOverwriteSiteMlForDevice1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body map[string]models.MlOverwriteAdditionalProperties
	errBody := json.Unmarshal([]byte(`{"iOS":{"int":6,"ple":-3},"iPod":{"int":-10,"ple":-5}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesLocation.OverwriteSiteMlForDevice(ctx, siteId, deviceId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestClearSiteMlOverwriteForMap tests the behavior of the SitesLocation
func TestSitesLocationTestClearSiteMlOverwriteForMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesLocation.ClearSiteMlOverwriteForMap(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesLocationTestOverwriteSiteMlForMap tests the behavior of the SitesLocation
func TestSitesLocationTestOverwriteSiteMlForMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body map[string]models.MlOverwriteAdditionalProperties
	errBody := json.Unmarshal([]byte(`{"iOS":{"int":6,"ple":-3},"iPod":{"int":-10,"ple":-5}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesLocation.OverwriteSiteMlForMap(ctx, siteId, mapId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestOverwriteSiteMlForMap1 tests the behavior of the SitesLocation
func TestSitesLocationTestOverwriteSiteMlForMap1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body map[string]models.MlOverwriteAdditionalProperties
	errBody := json.Unmarshal([]byte(`{"iOS":{"int":6,"ple":-3},"iPod":{"int":-10,"ple":-5}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesLocation.OverwriteSiteMlForMap(ctx, siteId, mapId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"current":{"Android":{"completed":36,"int":-6,"level":3,"ple":-3,"quality":"4","src":"device","timestamp":1442854794},"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"2","src":"default","timestamp":1442854704},"iPod":{"int":-10,"overwrite":true,"ple":-5,"src":"overwrite"}},"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"},{"beacon_id":"7913f032-aab4-c3ae-e83e-5a2756ef4d40","current":{"iOS":{"completed":16,"int":-6,"level":6,"ple":-3,"quality":"last","src":"device","timestamp":1442854704}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLocationTestResetSiteMlStatsByMap tests the behavior of the SitesLocation
func TestSitesLocationTestResetSiteMlStatsByMap(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesLocation.ResetSiteMlStatsByMap(ctx, siteId, mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
