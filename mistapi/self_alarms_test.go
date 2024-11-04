package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestSelfAlarmsTestListAlarmSubscriptions tests the behavior of the SelfAlarms
func TestSelfAlarmsTestListAlarmSubscriptions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := selfAlarms.ListAlarmSubscriptions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
