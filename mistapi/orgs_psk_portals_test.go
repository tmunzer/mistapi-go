// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsPskPortalsTestListOrgPskPortals tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortals(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortals(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsPskPortalsTestListOrgPskPortals1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortals1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortals(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsPskPortalsTestListOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestListOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestCountOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCountOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgPskPortalLogsCountDistinctEnum("pskportal_id")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsPskPortals.CountOrgPskPortalLogs(ctx, orgId, &distinct, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestCountOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCountOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgPskPortalLogsCountDistinctEnum("pskportal_id")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsPskPortals.CountOrgPskPortalLogs(ctx, orgId, &distinct, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestSearchOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestSearchOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsPskPortals.SearchOrgPskPortalLogs(ctx, orgId, &limit, &page, nil, nil, &duration, &sort, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestSearchOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestSearchOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsPskPortals.SearchOrgPskPortalLogs(ctx, orgId, &limit, &page, nil, nil, &duration, &sort, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestDeleteOrgPskPortal tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestDeleteOrgPskPortal(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsPskPortals.DeleteOrgPskPortal(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestDeleteOrgPskPortalImage tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestDeleteOrgPskPortalImage(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsPskPortals.DeleteOrgPskPortalImage(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestUploadOrgPskPortalImage tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUploadOrgPskPortalImage(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsPskPortals.UploadOrgPskPortalImage(ctx, orgId, pskportalId, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestUpdateOrgPskPortalTemplate tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUpdateOrgPskPortalTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsPskPortals.UpdateOrgPskPortalTemplate(ctx, orgId, pskportalId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
