// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsVarsTestSearchOrgVars tests the behavior of the OrgsVars
func TestOrgsVarsTestSearchOrgVars(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsVars.SearchOrgVars(ctx, orgId, nil, nil, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1693952979,"limit":10,"results":[{"created_time":1618457655.384858,"modified_time":1693610886.477805,"org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","site_id":"1519f016-4e41-47c0-a396-cce4d04bac0b","src":"site","var":"mvp"}],"start":1693949379,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsVarsTestSearchOrgVars1 tests the behavior of the OrgsVars
func TestOrgsVarsTestSearchOrgVars1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsVars.SearchOrgVars(ctx, orgId, nil, nil, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1693952979,"limit":10,"results":[{"created_time":1618457655.384858,"modified_time":1693610886.477805,"org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","site_id":"1519f016-4e41-47c0-a396-cce4d04bac0b","src":"site","var":"mvp"}],"start":1693949379,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
