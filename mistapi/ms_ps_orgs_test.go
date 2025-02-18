package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsOrgsTestListMspOrgs tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestListMspOrgs(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsOrgs.ListMspOrgs(ctx, mspId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":1440}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgsTestCreateMspOrg tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestCreateMspOrg(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Org
    errBody := json.Unmarshal([]byte(`{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":10}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsOrgs.CreateMspOrg(ctx, mspId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":1440}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgsTestManageMspOrgs tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestManageMspOrgs(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MspOrgChange
    errBody := json.Unmarshal([]byte(`{"op":"assign","org_ids":["2b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := msPsOrgs.ManageMspOrgs(ctx, mspId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsOrgsTestSearchMspOrgs tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestSearchMspOrgs(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    limit := int(100)
    apiResponse, err := msPsOrgs.SearchMspOrgs(ctx, mspId, nil, nil, nil, nil, nil, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1614383378.4365287,"limit":10,"results":[{"msp_id":"d287e62f-0000-0000-0000-f2b9ba0a531f","name":"Test Org","num_aps":9,"num_sites":5,"num_switches":1,"num_unassigned_aps":1,"org_id":"bb1a8bf6-0000-0000-0000-8053a663cf65","sub_ana_required":9,"sub_ast_entitled":5,"sub_ast_required":3,"sub_eng_required":3,"sub_ex12_required":1,"sub_insufficient":true,"sub_man_required":9,"sub_vna_entitled":1,"timestamp":1614322563.513937,"trial_enabled":false,"usage_types":["sub_eng"]},{"msp_id":"d287e62f-0000-0000-0000-f2b9ba0a531f","name":"Rogue Test1","num_aps":1,"num_sites":1,"org_id":"0fb81690-0000-0000-0000-9596d1d1534f","sub_ana_entitled":1,"sub_ana_required":1,"sub_insufficient":false,"sub_man_entitled":1,"sub_man_required":1,"timestamp":1614309876.500955}],"start":1613778578.4365668,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgsTestDeleteMspOrg tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestDeleteMspOrg(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := msPsOrgs.DeleteMspOrg(ctx, mspId, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsOrgsTestGetMspOrg tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestGetMspOrg(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsOrgs.GetMspOrg(ctx, mspId, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":1440}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgsTestUpdateMspOrg tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestUpdateMspOrg(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := msPsOrgs.UpdateMspOrg(ctx, mspId, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":1440}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgsTestListMspOrgStats tests the behavior of the MSPsOrgs
func TestMSPsOrgsTestListMspOrgStats(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := msPsOrgs.ListMspOrgStats(ctx, mspId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
