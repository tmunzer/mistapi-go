package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsMarvisTestCountMspsMarvisActions tests the behavior of the MSPsMarvis
func TestMSPsMarvisTestCountMspsMarvisActions(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.MspMarvisSuggestionsCountDistinctEnum("org_id")
    limit := int(100)
    apiResponse, err := msPsMarvis.CountMspsMarvisActions(ctx, mspId, &distinct, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"status","limit":1000,"results":[{"count":24,"status":"002e176a-0000-000-1111-002e208b20e1"},{"count":12,"status":"2d3f176a-0000-000-2222-002e208f176a"},{"count":15,"status":"08b2176a-0000-000-3333-002e208b2d3f"}],"total":3}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsMarvisTestCountMspsMarvisActions1 tests the behavior of the MSPsMarvis
func TestMSPsMarvisTestCountMspsMarvisActions1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.MspMarvisSuggestionsCountDistinctEnum("org_id")
    limit := int(100)
    apiResponse, err := msPsMarvis.CountMspsMarvisActions(ctx, mspId, &distinct, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"status","limit":1000,"results":[{"count":24,"status":"002e176a-0000-000-1111-002e208b20e1"},{"count":12,"status":"2d3f176a-0000-000-2222-002e208f176a"},{"count":15,"status":"08b2176a-0000-000-3333-002e208b2d3f"}],"total":3}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
