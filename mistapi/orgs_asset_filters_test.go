package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsAssetFiltersTestListOrgAssetFilters tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestListOrgAssetFilters(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsAssetFilters.ListOrgAssetFilters(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAssetFiltersTestListOrgAssetFilters1 tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestListOrgAssetFilters1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsAssetFilters.ListOrgAssetFilters(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAssetFiltersTestCreateOrgAssetFilter tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestCreateOrgAssetFilter(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AssetFilter
    errBody := json.Unmarshal([]byte(`{"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","mfg_company_id":0,"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAssetFilters.CreateOrgAssetFilter(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAssetFiltersTestCreateOrgAssetFilter1 tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestCreateOrgAssetFilter1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AssetFilter
    errBody := json.Unmarshal([]byte(`{"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","mfg_company_id":0,"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAssetFilters.CreateOrgAssetFilter(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAssetFiltersTestDeleteOrgAssetFilter tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestDeleteOrgAssetFilter(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAssetFilters.DeleteOrgAssetFilter(ctx, orgId, assetfilterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAssetFiltersTestGetOrgAssetFilter tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestGetOrgAssetFilter(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsAssetFilters.GetOrgAssetFilter(ctx, orgId, assetfilterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAssetFiltersTestGetOrgAssetFilter1 tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestGetOrgAssetFilter1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsAssetFilters.GetOrgAssetFilter(ctx, orgId, assetfilterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAssetFiltersTestUpdateOrgAssetFilter tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestUpdateOrgAssetFilter(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AssetFilter
    errBody := json.Unmarshal([]byte(`{"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab2","mfg_company_id":0,"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAssetFilters.UpdateOrgAssetFilter(ctx, orgId, assetfilterId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAssetFiltersTestUpdateOrgAssetFilter1 tests the behavior of the OrgsAssetFilters
func TestOrgsAssetFiltersTestUpdateOrgAssetFilter1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AssetFilter
    errBody := json.Unmarshal([]byte(`{"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab2","mfg_company_id":0,"name":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAssetFilters.UpdateOrgAssetFilter(ctx, orgId, assetfilterId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":0,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
