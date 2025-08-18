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

// TestOrgsWLANTemplatesTestListOrgTemplates tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestListOrgTemplates(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsWlanTemplates.ListOrgTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestListOrgTemplates1 tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestListOrgTemplates1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsWlanTemplates.ListOrgTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestCreateOrgTemplate tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestCreateOrgTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWlanTemplates.CreateOrgTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestCreateOrgTemplate1 tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestCreateOrgTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWlanTemplates.CreateOrgTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestDeleteOrgTemplate tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestDeleteOrgTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsWlanTemplates.DeleteOrgTemplate(ctx, orgId, templateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWLANTemplatesTestGetOrgTemplate tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestGetOrgTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsWlanTemplates.GetOrgTemplate(ctx, orgId, templateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestGetOrgTemplate1 tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestGetOrgTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsWlanTemplates.GetOrgTemplate(ctx, orgId, templateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestUpdateOrgTemplate tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestUpdateOrgTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWlanTemplates.UpdateOrgTemplate(ctx, orgId, templateId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestUpdateOrgTemplate1 tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestUpdateOrgTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWlanTemplates.UpdateOrgTemplate(ctx, orgId, templateId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestCloneOrgTemplate tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestCloneOrgTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.NameString
	errBody := json.Unmarshal([]byte(`{"name":"Cloned"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsWlanTemplates.CloneOrgTemplate(ctx, orgId, templateId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWLANTemplatesTestCloneOrgTemplate1 tests the behavior of the OrgsWLANTemplates
func TestOrgsWLANTemplatesTestCloneOrgTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	templateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.NameString
	errBody := json.Unmarshal([]byte(`{"name":"Cloned"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsWlanTemplates.CloneOrgTemplate(ctx, orgId, templateId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"applies":{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"created_time":0,"deviceprofile_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"exceptions":{"site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},"filter_by_deviceprofile":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
