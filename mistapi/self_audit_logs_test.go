package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestSelfAuditLogsTestListSelfAuditLogs tests the behavior of the SelfAuditLogs
func TestSelfAuditLogsTestListSelfAuditLogs(t *testing.T) {
    ctx := context.Background()
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := selfAuditLogs.ListSelfAuditLogs(ctx, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1428954000,"limit":100,"page":1,"results":[{"admin_id":"72bfa2bd-e58a-4670-9d20-a1468f7a6f58","admin_name":"test@mistsys.com","after":{"auth":{"type":"open"}},"before":{"auth":{"type":"psk"}},"id":"c6f9347b-b0a4-4a23-b927-fa9249f2ffb2","message":"Update WLAN \"Corporate\"","org_id":"423f6eca-6276-4994-bfeb-53cbbbba6f04","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1431382121}],"start":1428939600,"total":135}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
