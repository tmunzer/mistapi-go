package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestAdminsTestRegisterNewAdmin tests the behavior of the Admins
func TestAdminsTestRegisterNewAdmin(t *testing.T) {
    ctx := context.Background()
    var body models.AdminInvite
    errBody := json.Unmarshal([]byte(`{"account_only":false,"allow_mist":false,"city":"Cupertino","country":"United States","email":"test@mistsys.com","first_name":"John","invite_code":"MISTROCKS","last_name":"Smith","org_name":"Smith LLC","password":"foryoureyesonly","recaptcha":"string","recaptcha_flavor":"hcaptcha","referer_invite_token":"Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat","return_to":"https://mist.zendesk.com/hc/quickstart.pdf","state":"California","street_address":"1601 S De Anza Blvd Ste 248","street_address 2":"1601 S De Anza Blvd Ste 248","zipcode":"95014"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := admins.RegisterNewAdmin(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestAdminsTestGetAdminRegistrationInfo tests the behavior of the Admins
func TestAdminsTestGetAdminRegistrationInfo(t *testing.T) {
    ctx := context.Background()
    recaptchaFlavor := models.RecaptchaFlavorEnum("google")
    apiResponse, err := admins.GetAdminRegistrationInfo(ctx, &recaptchaFlavor)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"flavor":"google","required":true,"sitekey":"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsTestGetAdminRegistrationInfo1 tests the behavior of the Admins
func TestAdminsTestGetAdminRegistrationInfo1(t *testing.T) {
    ctx := context.Background()
    recaptchaFlavor := models.RecaptchaFlavorEnum("google")
    apiResponse, err := admins.GetAdminRegistrationInfo(ctx, &recaptchaFlavor)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"flavor":"google","required":true,"sitekey":"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
