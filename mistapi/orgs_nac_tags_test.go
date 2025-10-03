// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsNACTagsTestListOrgNacTags tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestListOrgNacTags(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	page := int(1)
	apiResponse, err := orgsNacTags.ListOrgNacTags(ctx, orgId, nil, nil, nil, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsNACTagsTestListOrgNacTags1 tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestListOrgNacTags1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	page := int(1)
	apiResponse, err := orgsNacTags.ListOrgNacTags(ctx, orgId, nil, nil, nil, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsNACTagsTestCreateOrgNacTag tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestCreateOrgNacTag(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsNacTags.CreateOrgNacTag(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACTagsTestCreateOrgNacTag1 tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestCreateOrgNacTag1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsNacTags.CreateOrgNacTag(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACTagsTestDeleteOrgNacTag tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestDeleteOrgNacTag(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	nactagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsNacTags.DeleteOrgNacTag(ctx, orgId, nactagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACTagsTestGetOrgNacTag tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestGetOrgNacTag(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	nactagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsNacTags.GetOrgNacTag(ctx, orgId, nactagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACTagsTestGetOrgNacTag1 tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestGetOrgNacTag1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	nactagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsNacTags.GetOrgNacTag(ctx, orgId, nactagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACTagsTestUpdateOrgNacTag tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestUpdateOrgNacTag(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	nactagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsNacTags.UpdateOrgNacTag(ctx, orgId, nactagId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACTagsTestUpdateOrgNacTag1 tests the behavior of the OrgsNACTags
func TestOrgsNACTagsTestUpdateOrgNacTag1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	nactagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsNacTags.UpdateOrgNacTag(ctx, orgId, nactagId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"egress_vlan_names":["1vlan-30","1vlan-20","2vlan10"],"name":"trunk_ap","type":"egress_vlan_names"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
