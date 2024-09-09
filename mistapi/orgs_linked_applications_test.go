package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsLinkedApplicationsTestLinkOrgToJuniperJuniperAccount tests the behavior of the OrgsLinkedApplications
func TestOrgsLinkedApplicationsTestLinkOrgToJuniperJuniperAccount(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountJuniperConfig
    errBody := json.Unmarshal([]byte(`{"password":"password","username":"john@nmo.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsLinkedApplications.LinkOrgToJuniperJuniperAccount(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"account":{"linked_by":"John Smith (john@abccorp.com)","name":"ABC Corp"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLinkedApplicationsTestUnlinkOrgFromJuniperCustomerId tests the behavior of the OrgsLinkedApplications
func TestOrgsLinkedApplicationsTestUnlinkOrgFromJuniperCustomerId(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsLinkedApplications.UnlinkOrgFromJuniperCustomerId(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
