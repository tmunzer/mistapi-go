package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsTicketsTestListMspTickets tests the behavior of the MSPsTickets
func TestMSPsTicketsTestListMspTickets(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := msPsTickets.ListMspTickets(ctx, mspId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsTicketsTestListMspTickets1 tests the behavior of the MSPsTickets
func TestMSPsTicketsTestListMspTickets1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := msPsTickets.ListMspTickets(ctx, mspId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsTicketsTestCountMspTickets tests the behavior of the MSPsTickets
func TestMSPsTicketsTestCountMspTickets(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.MspTicketsCountDistinctEnum("status")
    limit := int(100)
    apiResponse, err := msPsTickets.CountMspTickets(ctx, mspId, &distinct, &limit)
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

// TestMSPsTicketsTestCountMspTickets1 tests the behavior of the MSPsTickets
func TestMSPsTicketsTestCountMspTickets1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.MspTicketsCountDistinctEnum("status")
    limit := int(100)
    apiResponse, err := msPsTickets.CountMspTickets(ctx, mspId, &distinct, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
