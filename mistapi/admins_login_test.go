package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestAdminsLoginTestLogin tests the behavior of the AdminsLogin
func TestAdminsLoginTestLogin(t *testing.T) {
    ctx := context.Background()
    var body models.Login
    errBody := json.Unmarshal([]byte(`{"email":"test@mistsys.com","password":"foryoureyesonly","two_factor":"123456"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := adminsLogin.Login(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(false,"Set-Cookie",""),
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"email":"test@mistsys.com","two_factor_passed":false,"two_factor_required":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsLoginTestLogin1 tests the behavior of the AdminsLogin
func TestAdminsLoginTestLogin1(t *testing.T) {
    ctx := context.Background()
    var body models.Login
    errBody := json.Unmarshal([]byte(`{"email":"test@mistsys.com","password":"foryoureyesonly","two_factor":"123456"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := adminsLogin.Login(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(false,"Set-Cookie",""),
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"email":"test@mistsys.com","two_factor_passed":false,"two_factor_required":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsLoginTestTwoFactor tests the behavior of the AdminsLogin
func TestAdminsLoginTestTwoFactor(t *testing.T) {
    ctx := context.Background()
    var body models.TwoFactorString
    errBody := json.Unmarshal([]byte(`{"two_factor":"123456"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := adminsLogin.TwoFactor(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
