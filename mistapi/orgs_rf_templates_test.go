// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsRFTemplatesTestListOrgRfTemplates tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestListOrgRfTemplates(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsRfTemplates.ListOrgRfTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestListOrgRfTemplates1 tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestListOrgRfTemplates1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsRfTemplates.ListOrgRfTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestCreateOrgRfTemplate tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestCreateOrgRfTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsRfTemplates.CreateOrgRfTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestCreateOrgRfTemplate1 tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestCreateOrgRfTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsRfTemplates.CreateOrgRfTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestDeleteOrgRfTemplate tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestDeleteOrgRfTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rftemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsRfTemplates.DeleteOrgRfTemplate(ctx, orgId, rftemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsRFTemplatesTestGetOrgRfTemplate tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestGetOrgRfTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rftemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsRfTemplates.GetOrgRfTemplate(ctx, orgId, rftemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestGetOrgRfTemplate1 tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestGetOrgRfTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rftemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsRfTemplates.GetOrgRfTemplate(ctx, orgId, rftemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestUpdateOrgRfTemplate tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestUpdateOrgRfTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rftemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsRfTemplates.UpdateOrgRfTemplate(ctx, orgId, rftemplateId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsRFTemplatesTestUpdateOrgRfTemplate1 tests the behavior of the OrgsRFTemplates
func TestOrgsRFTemplatesTestUpdateOrgRfTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rftemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsRfTemplates.UpdateOrgRfTemplate(ctx, orgId, rftemplateId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"ant_gain":0,"bandwidth":20,"channels":[1,6,11],"disabled":false,"power_max":11,"power_min":3,"preamble":"short"},"band_24_usage":"auto","band_5":{"allow_rrm_disable":false,"ant_gain":0,"bandwidth":80,"channels":[36,40,44,48,52,56,60,64,100,104,149,153,157,161],"disabled":false,"power_max":16,"power_min":9,"preamble":"short"},"country_code":"FR","created_time":1594743723,"id":"b3f20330-f76a-49f1-bc65-0d8727140b1d","model_specific":{},"modified_time":1613582192,"name":"Lab","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
