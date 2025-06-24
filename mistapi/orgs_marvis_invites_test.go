package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsMarvisInvitesTestListOrgMarvisClientInvites tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestListOrgMarvisClientInvites(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMarvisInvites.ListOrgMarvisClientInvites(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestListOrgMarvisClientInvites1 tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestListOrgMarvisClientInvites1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMarvisInvites.ListOrgMarvisClientInvites(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestCreateOrgMarvisClientInvite tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestCreateOrgMarvisClientInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MarvisClient
    errBody := json.Unmarshal([]byte(`{"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMarvisInvites.CreateOrgMarvisClientInvite(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestCreateOrgMarvisClientInvite1 tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestCreateOrgMarvisClientInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MarvisClient
    errBody := json.Unmarshal([]byte(`{"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMarvisInvites.CreateOrgMarvisClientInvite(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestDeleteOrgMarvisClientInvite tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestDeleteOrgMarvisClientInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    marvisinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsMarvisInvites.DeleteOrgMarvisClientInvite(ctx, orgId, marvisinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMarvisInvitesTestGetOrgMarvisClientInvite tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestGetOrgMarvisClientInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    marvisinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMarvisInvites.GetOrgMarvisClientInvite(ctx, orgId, marvisinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestGetOrgMarvisClientInvite1 tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestGetOrgMarvisClientInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    marvisinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMarvisInvites.GetOrgMarvisClientInvite(ctx, orgId, marvisinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestUpdateOrgMarvisClientInvite tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestUpdateOrgMarvisClientInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    marvisinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MarvisClient
    errBody := json.Unmarshal([]byte(`{"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMarvisInvites.UpdateOrgMarvisClientInvite(ctx, orgId, marvisinviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMarvisInvitesTestUpdateOrgMarvisClientInvite1 tests the behavior of the OrgsMarvisInvites
func TestOrgsMarvisInvitesTestUpdateOrgMarvisClientInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    marvisinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MarvisClient
    errBody := json.Unmarshal([]byte(`{"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMarvisInvites.UpdateOrgMarvisClientInvite(ctx, orgId, marvisinviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"3a14098f-b995-7552-b0a4-b8ee39b337a6","name":"Handhelds"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
