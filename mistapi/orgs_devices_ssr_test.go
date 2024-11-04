package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsDevicesSSRTestGetOrg128TRegistrationCommands tests the behavior of the OrgsDevicesSSR
func TestOrgsDevicesSSRTestGetOrg128TRegistrationCommands(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsDevicesSsr.GetOrg128TRegistrationCommands(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"conductor_cmd":"register mist eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ","registration_code":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ","router_shell_cmd":"128agent register --registration-code eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
