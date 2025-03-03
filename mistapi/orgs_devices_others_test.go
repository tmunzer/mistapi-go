package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsDevicesOthersTestListOrgOtherDevices tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestListOrgOtherDevices(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsDevicesOthers.ListOrgOtherDevices(ctx, orgId, nil, nil, nil, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1676983730,"device_mac":"001122334455","id":"ae9dee49-69e7-4710-a114-5b827a777738","mac":"5c5b35000018","model":"AP41","modified_time":1676983730,"name":"hallway","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","vendor":"cradlepoint"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesOthersTestUpdateOrgOtherDevices tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestUpdateOrgOtherDevices(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OtherDeviceUpdateMulti
    errBody := json.Unmarshal([]byte(`{"device_mac":"0adfea67e65b","macs":["5c5b350e0001","5c5b350e0003"],"op":"assign","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsDevicesOthers.UpdateOrgOtherDevices(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsDevicesOthersTestCountOrgOtherDeviceEvents tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestCountOrgOtherDeviceEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgOtherdevicesEventsCountDistinctEnum("mac")
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsDevicesOthers.CountOrgOtherDeviceEvents(ctx, orgId, &distinct, nil, nil, nil, &duration, &limit)
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

// TestOrgsDevicesOthersTestSearchOrgOtherDeviceEvents tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestSearchOrgOtherDeviceEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsDevicesOthers.SearchOrgOtherDeviceEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":{"device_mac":"string","mac":"5c5b351e13b5","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Plugged: The Internal 5GB (SIM1) has been inserted into Internal 1.","timestamp":547235620.89,"type":"CELLULAR_EDGE_MODEM_WAN_PLUGGED","vendor":"cradlepoint"},"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesOthersTestDeleteOrgOtherDevice tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestDeleteOrgOtherDevice(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    resp, err := orgsDevicesOthers.DeleteOrgOtherDevice(ctx, orgId, deviceMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsDevicesOthersTestGetOrgOtherDevice tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestGetOrgOtherDevice(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    apiResponse, err := orgsDevicesOthers.GetOrgOtherDevice(ctx, orgId, deviceMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"device_mac":"00112233abcd","id":"ae9dee49-69e7-4710-a114-5b827a777738","mac":"5c5b35000018","model":"W2005","name":"W2005-268","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","serial":"WB23015E025468","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","state":"online","vendor":"cradlepoint","vendor_api_id":"4658714"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesOthersTestUpdateOrgOtherDevice tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestUpdateOrgOtherDevice(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    var body models.OtherDeviceUpdate
    errBody := json.Unmarshal([]byte(`{"device_mac":"0adfea67e65b","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsDevicesOthers.UpdateOrgOtherDevice(ctx, orgId, deviceMac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsDevicesOthersTestRebootOrgOtherDevice tests the behavior of the OrgsDevicesOthers
func TestOrgsDevicesOthersTestRebootOrgOtherDevice(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    resp, err := orgsDevicesOthers.RebootOrgOtherDevice(ctx, orgId, deviceMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
