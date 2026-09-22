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

// TestOrgsSecurityZonesTestListOrgSecurityZones tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestListOrgSecurityZones(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsSecurityZones.ListOrgSecurityZones(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestListOrgSecurityZones1 tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestListOrgSecurityZones1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsSecurityZones.ListOrgSecurityZones(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestCreateOrgSecurityZone tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestCreateOrgSecurityZone(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Securityzone
	errBody := json.Unmarshal([]byte(`{"name":"corp-zone"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSecurityZones.CreateOrgSecurityZone(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestCreateOrgSecurityZone1 tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestCreateOrgSecurityZone1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Securityzone
	errBody := json.Unmarshal([]byte(`{"name":"corp-zone"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSecurityZones.CreateOrgSecurityZone(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestDeleteOrgSecurityZone tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestDeleteOrgSecurityZone(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	securityzoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsSecurityZones.DeleteOrgSecurityZone(ctx, orgId, securityzoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSecurityZonesTestGetOrgSecurityZone tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestGetOrgSecurityZone(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	securityzoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsSecurityZones.GetOrgSecurityZone(ctx, orgId, securityzoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestGetOrgSecurityZone1 tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestGetOrgSecurityZone1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	securityzoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsSecurityZones.GetOrgSecurityZone(ctx, orgId, securityzoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestUpdateOrgSecurityZone tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestUpdateOrgSecurityZone(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	securityzoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Securityzone
	errBody := json.Unmarshal([]byte(`{"name":"corp-zone"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSecurityZones.UpdateOrgSecurityZone(ctx, orgId, securityzoneId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityZonesTestUpdateOrgSecurityZone1 tests the behavior of the OrgsSecurityZones
func TestOrgsSecurityZonesTestUpdateOrgSecurityZone1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	securityzoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Securityzone
	errBody := json.Unmarshal([]byte(`{"name":"corp-zone"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSecurityZones.UpdateOrgSecurityZone(ctx, orgId, securityzoneId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1656353525,"id":"a045ea40-ccd0-4660-8f35-0305fae26e9c","modified_time":1656353525,"name":"corp-zone","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
