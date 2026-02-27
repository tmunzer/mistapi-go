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

// TestSitesMapStacksTestListSiteMapStacks tests the behavior of the SitesMapStacks
func TestSitesMapStacksTestListSiteMapStacks(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	apiResponse, err := sitesMapStacks.ListSiteMapStacks(ctx, siteId, &limit, &page, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"Board Room","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapStacksTestListSiteMapStacks1 tests the behavior of the SitesMapStacks
func TestSitesMapStacksTestListSiteMapStacks1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	apiResponse, err := sitesMapStacks.ListSiteMapStacks(ctx, siteId, &limit, &page, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"Board Room","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapStacksTestCreateSiteMapStack tests the behavior of the SitesMapStacks
func TestSitesMapStacksTestCreateSiteMapStack(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Mapstack
	errBody := json.Unmarshal([]byte(`{"name":"Board Room"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesMapStacks.CreateSiteMapStack(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesMapStacksTestCreateSiteMapStack1 tests the behavior of the SitesMapStacks
func TestSitesMapStacksTestCreateSiteMapStack1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Mapstack
	errBody := json.Unmarshal([]byte(`{"name":"Board Room"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesMapStacks.CreateSiteMapStack(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
