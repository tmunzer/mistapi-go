package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsWxTagsTestListOrgWxTags tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestListOrgWxTags(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWxTags.ListOrgWxTags(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestListOrgWxTags1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestListOrgWxTags1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWxTags.ListOrgWxTags(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestCreateOrgWxTag tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestCreateOrgWxTag(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanTag
    errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWxTags.CreateOrgWxTag(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestCreateOrgWxTag1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestCreateOrgWxTag1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanTag
    errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWxTags.CreateOrgWxTag(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestGetOrgApplicationList tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgApplicationList(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgApplicationList(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"group":"Emails","key":"gmail","name":"Gmail - web/app"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestGetOrgApplicationList1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgApplicationList1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgApplicationList(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"group":"Emails","key":"gmail","name":"Gmail - web/app"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestDeleteOrgWxTag tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestDeleteOrgWxTag(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWxTags.DeleteOrgWxTag(ctx, orgId, wxtagId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWxTagsTestGetOrgWxTag tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgWxTag(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgWxTag(ctx, orgId, wxtagId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestGetOrgWxTag1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgWxTag1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgWxTag(ctx, orgId, wxtagId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestUpdateOrgWxTag tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestUpdateOrgWxTag(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWxTags.UpdateOrgWxTag(ctx, orgId, wxtagId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestUpdateOrgWxTag1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestUpdateOrgWxTag1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWxTags.UpdateOrgWxTag(ctx, orgId, wxtagId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTagsTestGetOrgCurrentMatchingClientsOfAWxTag tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgCurrentMatchingClientsOfAWxTag(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgCurrentMatchingClientsOfAWxTag(ctx, orgId, wxtagId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsWxTagsTestGetOrgCurrentMatchingClientsOfAWxTag1 tests the behavior of the OrgsWxTags
func TestOrgsWxTagsTestGetOrgCurrentMatchingClientsOfAWxTag1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxTags.GetOrgCurrentMatchingClientsOfAWxTag(ctx, orgId, wxtagId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
