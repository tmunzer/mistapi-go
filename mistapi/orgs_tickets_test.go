package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsTicketsTestListOrgTickets tests the behavior of the OrgsTickets
func TestOrgsTicketsTestListOrgTickets(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := orgsTickets.ListOrgTickets(ctx, orgId, nil, nil, &duration)
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

// TestOrgsTicketsTestCreateOrgTicket tests the behavior of the OrgsTickets
func TestOrgsTicketsTestCreateOrgTicket(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Ticket
    errBody := json.Unmarshal([]byte(`{"comment":"string","subject":"string","type":"question"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsTickets.CreateOrgTicket(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsTicketsTestCountOrgTickets tests the behavior of the OrgsTickets
func TestOrgsTicketsTestCountOrgTickets(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgTicketsCountDistinctEnum("status")
    apiResponse, err := orgsTickets.CountOrgTickets(ctx, orgId, &distinct)
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

// TestOrgsTicketsTestGetOrgTicket tests the behavior of the OrgsTickets
func TestOrgsTicketsTestGetOrgTicket(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ticketId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := orgsTickets.GetOrgTicket(ctx, orgId, ticketId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsTicketsTestUpdateOrgTicket tests the behavior of the OrgsTickets
func TestOrgsTicketsTestUpdateOrgTicket(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ticketId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Ticket
    errBody := json.Unmarshal([]byte(`{"comment":"string","subject":"string","type":"question"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsTickets.UpdateOrgTicket(ctx, orgId, ticketId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsTicketsTestUploadOrgTicketAttachment tests the behavior of the OrgsTickets
func TestOrgsTicketsTestUploadOrgTicketAttachment(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ticketId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsTickets.UploadOrgTicketAttachment(ctx, orgId, ticketId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsTicketsTestGetOrgTicketAttachment tests the behavior of the OrgsTickets
func TestOrgsTicketsTestGetOrgTicketAttachment(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ticketId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    attachmentId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := orgsTickets.GetOrgTicketAttachment(ctx, orgId, ticketId, attachmentId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"content_url":"https://api.mist.com/api/v1/forward/download?jwt=..."}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsTicketsTestAddOrgTicketComment tests the behavior of the OrgsTickets
func TestOrgsTicketsTestAddOrgTicketComment(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ticketId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    comment := "this is urgent"
    
    apiResponse, err := orgsTickets.AddOrgTicketComment(ctx, orgId, ticketId, &comment, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"comments":[{"attachments":[{"content_type":"string","content_url":"string","size":0}],"author":"string","comment":"string","created_at":0}],"created_at":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","requester":"string","status":"open","subject":"string","type":"string","updated_at":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
