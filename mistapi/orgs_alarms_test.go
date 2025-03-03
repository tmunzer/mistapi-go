package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsAlarmsTestAckOrgMultipleAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestAckOrgMultipleAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Alarms
    errBody := json.Unmarshal([]byte(`{"alarm_ids":["ccb8c94d-ca56-4075-932f-1f2ab444ff2c","98ff4a3d-ec9b-4138-a42e-54fc3335179d"],"note":"maintenance window"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAlarms.AckOrgMultipleAlarms(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestAckOrgAllAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestAckOrgAllAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NoteString
    errBody := json.Unmarshal([]byte(`{"note":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAlarms.AckOrgAllAlarms(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestCountOrgAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestCountOrgAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsAlarms.CountOrgAlarms(ctx, orgId, nil, nil, nil, &duration, &limit, &page)
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

// TestOrgsAlarmsTestSearchOrgAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestSearchOrgAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsAlarms.SearchOrgAlarms(ctx, orgId, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmsTestUnackOrgMultipleAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestUnackOrgMultipleAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Alarms
    errBody := json.Unmarshal([]byte(`{"alarm_ids":["ccb8c94d-ca56-4075-932f-1f2ab444ff2c","98ff4a3d-ec9b-4138-a42e-54fc3335179d"],"note":"maintenance window"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAlarms.UnackOrgMultipleAlarms(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestUnackOrgAllAlarms tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestUnackOrgAllAlarms(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NoteString
    errBody := json.Unmarshal([]byte(`{"note":"maintenance window"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAlarms.UnackOrgAllAlarms(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestAckOrgAlarm tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestAckOrgAlarm(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    alarmId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NoteString
    errBody := json.Unmarshal([]byte(`{"note":"maintenance window"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAlarms.AckOrgAlarm(ctx, orgId, alarmId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestUnsubscribeOrgAlarmsReports tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestUnsubscribeOrgAlarmsReports(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAlarms.UnsubscribeOrgAlarmsReports(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmsTestSubscribeOrgAlarmsReports tests the behavior of the OrgsAlarms
func TestOrgsAlarmsTestSubscribeOrgAlarmsReports(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAlarms.SubscribeOrgAlarmsReports(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
