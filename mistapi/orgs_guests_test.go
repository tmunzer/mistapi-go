package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsGuestsTestListOrgGuestAuthorizations tests the behavior of the OrgsGuests
func TestOrgsGuestsTestListOrgGuestAuthorizations(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsGuests.ListOrgGuestAuthorizations(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsGuestsTestCountOrgGuestAuthorizations tests the behavior of the OrgsGuests
func TestOrgsGuestsTestCountOrgGuestAuthorizations(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgGuestsCountDistinctEnum("auth_method")
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsGuests.CountOrgGuestAuthorizations(ctx, orgId, &distinct, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsGuestsTestSearchOrgGuestAuthorization tests the behavior of the OrgsGuests
func TestOrgsGuestsTestSearchOrgGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId := "00000000-0000-0000-0000-000000000000"
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsGuests.SearchOrgGuestAuthorization(ctx, orgId, &wlanId, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1531862583,"limit":2,"next":"/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/guests/search?wlan_id=88ffe630-95b8-11e8-b294-346895ed1b7d&end=1531855849.000&limit=2&start=1531776183.0","results":[{"ap":"5c5b350e0001","auth_method":"passphrase","authorized_expiring_time":1531810258.1862731,"authorized_time":1531782218,"company":"mistsystems","email":"user@mistsys.com","name":"john","ssid":"openNet","timestamp":1531782218},{"ap":"5c5b350e0001","auth_method":"facebook","authorized_expiring_time":1531810821.145,"authorized_time":1531782632,"company":"xyz inc.","email":"cool_user@yahoo.com","name":"John White","ssid":"openNet","timestamp":1531782632}],"start":1531776183,"total":14}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsGuestsTestDeleteOrgGuestAuthorization tests the behavior of the OrgsGuests
func TestOrgsGuestsTestDeleteOrgGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    resp, err := orgsGuests.DeleteOrgGuestAuthorization(ctx, orgId, guestMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsGuestsTestGetOrgGuestAuthorization tests the behavior of the OrgsGuests
func TestOrgsGuestsTestGetOrgGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    apiResponse, err := orgsGuests.GetOrgGuestAuthorization(ctx, orgId, guestMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsGuestsTestUpdateOrgGuestAuthorization tests the behavior of the OrgsGuests
func TestOrgsGuestsTestUpdateOrgGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    var body models.Guest
    errBody := json.Unmarshal([]byte(`{"authorized":true,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsGuests.UpdateOrgGuestAuthorization(ctx, orgId, guestMac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
