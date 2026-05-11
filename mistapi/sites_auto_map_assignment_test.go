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

// TestSitesAutoMapAssignmentTestApplySiteAutoMapAssignment tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestApplySiteAutoMapAssignment(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AutoMapAssignmentRequest
	errBody := json.Unmarshal([]byte(`{}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesAutoMapAssignment.ApplySiteAutoMapAssignment(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"accepted_maps":["d3c42998-9012-4859-9743-6b9bee475309","f7a21456-7891-4abc-def0-123456789abc"],"message":"Accepted map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestApplySiteAutoMapAssignment1 tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestApplySiteAutoMapAssignment1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AutoMapAssignmentRequest
	errBody := json.Unmarshal([]byte(`{}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesAutoMapAssignment.ApplySiteAutoMapAssignment(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"accepted_maps":["d3c42998-9012-4859-9743-6b9bee475309","f7a21456-7891-4abc-def0-123456789abc"],"message":"Accepted map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestCancelSiteAutoMapAssignment tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestCancelSiteAutoMapAssignment(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesAutoMapAssignment.CancelSiteAutoMapAssignment(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesAutoMapAssignmentTestGetSiteAutoMapAssignmentStatus tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestGetSiteAutoMapAssignmentStatus(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesAutoMapAssignment.GetSiteAutoMapAssignmentStatus(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"est_time_left":280.5,"start_time":1678900062,"status":"in_progress","time_updated":1678900100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestGetSiteAutoMapAssignmentStatus1 tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestGetSiteAutoMapAssignmentStatus1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesAutoMapAssignment.GetSiteAutoMapAssignmentStatus(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"est_time_left":280.5,"start_time":1678900062,"status":"in_progress","time_updated":1678900100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestStartSiteAutoMapAssignment tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestStartSiteAutoMapAssignment(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAutoMapAssignment.StartSiteAutoMapAssignment(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"devices":{"5c5b35000001":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true},"5c5b35000002":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true},"5c5b35000003":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true}},"estimated_runtime":300,"reason":"Started auto map assignment","started":true,"valid":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestStartSiteAutoMapAssignment1 tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestStartSiteAutoMapAssignment1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAutoMapAssignment.StartSiteAutoMapAssignment(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"devices":{"5c5b35000001":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true},"5c5b35000002":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true},"5c5b35000003":{"reason":"Device meets the minimum requirements for auto map assignment","valid":true}},"estimated_runtime":300,"reason":"Started auto map assignment","started":true,"valid":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestClearSiteAutoMapAssignment tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestClearSiteAutoMapAssignment(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AutoMapAssignmentRequest
	errBody := json.Unmarshal([]byte(`{}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesAutoMapAssignment.ClearSiteAutoMapAssignment(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"message":"Rejected map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']","rejected_maps":["d3c42998-9012-4859-9743-6b9bee475309","f7a21456-7891-4abc-def0-123456789abc"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAutoMapAssignmentTestClearSiteAutoMapAssignment1 tests the behavior of the SitesAutoMapAssignment
func TestSitesAutoMapAssignmentTestClearSiteAutoMapAssignment1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AutoMapAssignmentRequest
	errBody := json.Unmarshal([]byte(`{}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesAutoMapAssignment.ClearSiteAutoMapAssignment(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"message":"Rejected map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']","rejected_maps":["d3c42998-9012-4859-9743-6b9bee475309","f7a21456-7891-4abc-def0-123456789abc"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
