// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsAdminsTestListOrgAdmins tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestListOrgAdmins(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsAdmins.ListOrgAdmins(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"admin_id":"456b7016-a916-a4b1-78dd-72b947c152b7","email":"jsmith@mycorp.org","first_name":"Joe","last_name":"Smith","privileges":[{"role":"admin","scope":"msp"},{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"admin","scope":"org"},{"orggroup_ids":["507f1bab-13ba-73e2-f291-2bcb8d1362b0"],"role":"read","scope":"orggroup"}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAdminsTestListOrgAdmins1 tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestListOrgAdmins1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsAdmins.ListOrgAdmins(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"admin_id":"456b7016-a916-a4b1-78dd-72b947c152b7","email":"jsmith@mycorp.org","first_name":"Joe","last_name":"Smith","privileges":[{"role":"admin","scope":"msp"},{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"admin","scope":"org"},{"orggroup_ids":["507f1bab-13ba-73e2-f291-2bcb8d1362b0"],"role":"read","scope":"orggroup"}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAdminsTestRevokeOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestRevokeOrgAdmin(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsAdmins.RevokeOrgAdmin(ctx, orgId, adminId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUpdateOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUpdateOrgAdmin(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsAdmins.UpdateOrgAdmin(ctx, orgId, adminId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAdminsTestUpdateOrgAdmin1 tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUpdateOrgAdmin1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsAdmins.UpdateOrgAdmin(ctx, orgId, adminId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAdminsTestInviteOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestInviteOrgAdmin(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsAdmins.InviteOrgAdmin(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUninviteOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUninviteOrgAdmin(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	inviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsAdmins.UninviteOrgAdmin(ctx, orgId, inviteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUpdateOrgAdminInvite tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUpdateOrgAdminInvite(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	inviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsAdmins.UpdateOrgAdminInvite(ctx, orgId, inviteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
