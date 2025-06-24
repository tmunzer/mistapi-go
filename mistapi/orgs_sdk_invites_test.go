package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsSDKInvitesTestListSdkInvites tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestListSdkInvites(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.ListSdkInvites(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestListSdkInvites1 tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestListSdkInvites1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.ListSdkInvites(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestCreateSdkInvite tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestCreateSdkInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Sdkinvite
    errBody := json.Unmarshal([]byte(`{"enabled":true,"name":"string","quota":0,"quota_limited":true}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSdkInvites.CreateSdkInvite(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestCreateSdkInvite1 tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestCreateSdkInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Sdkinvite
    errBody := json.Unmarshal([]byte(`{"enabled":true,"name":"string","quota":0,"quota_limited":true}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSdkInvites.CreateSdkInvite(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestRevokeSdkInvite tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestRevokeSdkInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSdkInvites.RevokeSdkInvite(ctx, orgId, sdkinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSDKInvitesTestGetSdkInvite tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestGetSdkInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.GetSdkInvite(ctx, orgId, sdkinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestGetSdkInvite1 tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestGetSdkInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.GetSdkInvite(ctx, orgId, sdkinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestUpdateSdkInvite tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestUpdateSdkInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkInvites.UpdateSdkInvite(ctx, orgId, sdkinviteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestUpdateSdkInvite1 tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestUpdateSdkInvite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkInvites.UpdateSdkInvite(ctx, orgId, sdkinviteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1428954000,"enabled":true,"expire_time":1428954000,"id":"5034b980-b49e-501c-66e0-9de4c38f18a2","name":"Macy's","quota":-1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKInvitesTestSendSdkInviteEmail tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestSendSdkInviteEmail(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.EmailString
    errBody := json.Unmarshal([]byte(`{"email":"test@abc.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsSdkInvites.SendSdkInviteEmail(ctx, orgId, sdkinviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSDKInvitesTestGetSdkInviteQrCode tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestGetSdkInviteQrCode(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.GetSdkInviteQrCode(ctx, orgId, sdkinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsSDKInvitesTestGetSdkInviteQrCode1 tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestGetSdkInviteQrCode1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkInvites.GetSdkInviteQrCode(ctx, orgId, sdkinviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsSDKInvitesTestSendSdkInviteSms tests the behavior of the OrgsSDKInvites
func TestOrgsSDKInvitesTestSendSdkInviteSms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdkinviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SdkInviteSms
    errBody := json.Unmarshal([]byte(`{"number":"14081234567"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsSdkInvites.SendSdkInviteSms(ctx, orgId, sdkinviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
