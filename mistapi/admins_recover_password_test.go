package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestAdminsRecoverPasswordTestRecoverPassword tests the behavior of the AdminsRecoverPassword
func TestAdminsRecoverPasswordTestRecoverPassword(t *testing.T) {
    ctx := context.Background()
    var body models.Recover
    errBody := json.Unmarshal([]byte(`{"email":"test@mistsys.com","recaptcha":"string","recaptcha_flavor":"hcaptcha"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := adminsRecoverPassword.RecoverPassword(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
