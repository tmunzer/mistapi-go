package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesPsksTestListSitePsks tests the behavior of the SitesPsks
func TestSitesPsksTestListSitePsks(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesPsks.ListSitePsks(ctx, siteId, nil, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestListSitePsks1 tests the behavior of the SitesPsks
func TestSitesPsksTestListSitePsks1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesPsks.ListSitePsks(ctx, siteId, nil, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestCreateSitePsk tests the behavior of the SitesPsks
func TestSitesPsksTestCreateSitePsk(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Psk
    errBody := json.Unmarshal([]byte(`{"mac":"string","name":"string","passphrase":"secretpsk","ssid":"string","usage":"multi","vlan_id":1}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesPsks.CreateSitePsk(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesPsksTestCreateSitePsk1 tests the behavior of the SitesPsks
func TestSitesPsksTestCreateSitePsk1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Psk
    errBody := json.Unmarshal([]byte(`{"mac":"string","name":"string","passphrase":"secretpsk","ssid":"string","usage":"multi","vlan_id":1}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesPsks.CreateSitePsk(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesPsksTestUpdateSiteMultiplePsks tests the behavior of the SitesPsks
func TestSitesPsksTestUpdateSiteMultiplePsks(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesPsks.UpdateSiteMultiplePsks(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestUpdateSiteMultiplePsks1 tests the behavior of the SitesPsks
func TestSitesPsksTestUpdateSiteMultiplePsks1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesPsks.UpdateSiteMultiplePsks(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestImportSitePsks tests the behavior of the SitesPsks
func TestSitesPsksTestImportSitePsks(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesPsks.ImportSitePsks(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestImportSitePsks1 tests the behavior of the SitesPsks
func TestSitesPsksTestImportSitePsks1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesPsks.ImportSitePsks(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","passphrase":"secretpsk","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ssid":"string","usage":"multi","vlan_id":1}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesPsksTestDeleteSitePsk tests the behavior of the SitesPsks
func TestSitesPsksTestDeleteSitePsk(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesPsks.DeleteSitePsk(ctx, siteId, pskId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesPsksTestGetSitePsk tests the behavior of the SitesPsks
func TestSitesPsksTestGetSitePsk(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesPsks.GetSitePsk(ctx, siteId, pskId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesPsksTestGetSitePsk1 tests the behavior of the SitesPsks
func TestSitesPsksTestGetSitePsk1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesPsks.GetSitePsk(ctx, siteId, pskId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesPsksTestUpdateSitePsk tests the behavior of the SitesPsks
func TestSitesPsksTestUpdateSitePsk(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Psk
    errBody := json.Unmarshal([]byte(`{"mac":"string","name":"string","passphrase":"secretpsk","ssid":"string","usage":"multi","vlan_id":1}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesPsks.UpdateSitePsk(ctx, siteId, pskId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesPsksTestUpdateSitePsk1 tests the behavior of the SitesPsks
func TestSitesPsksTestUpdateSitePsk1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Psk
    errBody := json.Unmarshal([]byte(`{"mac":"string","name":"string","passphrase":"secretpsk","ssid":"string","usage":"multi","vlan_id":1}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesPsks.UpdateSitePsk(ctx, siteId, pskId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
