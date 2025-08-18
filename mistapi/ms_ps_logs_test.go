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

// TestMSPsLogsTestListMspAuditLogs tests the behavior of the MSPsLogs
func TestMSPsLogsTestListMspAuditLogs(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	siteId := "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
	adminName := "jsnow"
	message := "Added new site"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := msPsLogs.ListMspAuditLogs(ctx, mspId, &siteId, &adminName, &message, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"admin_id":"72bfa2bd-e58a-4670-9d20-a1468f7a6f58","admin_name":"test@mistsys.com","id":"c6f9347b-b0a4-4a23-b927-fa9249f2ffb2","message":"TEST AUDIT","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1431382121}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsLogsTestListMspAuditLogs1 tests the behavior of the MSPsLogs
func TestMSPsLogsTestListMspAuditLogs1(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	siteId := "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
	adminName := "jsnow"
	message := "Added new site"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := msPsLogs.ListMspAuditLogs(ctx, mspId, &siteId, &adminName, &message, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"admin_id":"72bfa2bd-e58a-4670-9d20-a1468f7a6f58","admin_name":"test@mistsys.com","id":"c6f9347b-b0a4-4a23-b927-fa9249f2ffb2","message":"TEST AUDIT","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1431382121}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsLogsTestCountMspAuditLogs tests the behavior of the MSPsLogs
func TestMSPsLogsTestCountMspAuditLogs(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.MspLogsCountDistinctEnum("admin_name")
	limit := int(100)
	apiResponse, err := msPsLogs.CountMspAuditLogs(ctx, mspId, &distinct, &limit)
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

// TestMSPsLogsTestCountMspAuditLogs1 tests the behavior of the MSPsLogs
func TestMSPsLogsTestCountMspAuditLogs1(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.MspLogsCountDistinctEnum("admin_name")
	limit := int(100)
	apiResponse, err := msPsLogs.CountMspAuditLogs(ctx, mspId, &distinct, &limit)
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
