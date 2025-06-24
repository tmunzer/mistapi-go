package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsClientsWanTestCountOrgWanClientEvents tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestCountOrgWanClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgWanClientsEventsCountDistinctEnum("type")
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.CountOrgWanClientEvents(ctx, orgId, &distinct, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsWanTestCountOrgWanClientEvents1 tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestCountOrgWanClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgWanClientsEventsCountDistinctEnum("type")
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.CountOrgWanClientEvents(ctx, orgId, &distinct, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsWanTestCountOrgWanClients tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestCountOrgWanClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgWanClientsCountDistinctEnum("mac")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.CountOrgWanClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsClientsWanTestCountOrgWanClients1 tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestCountOrgWanClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgWanClientsCountDistinctEnum("mac")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.CountOrgWanClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsClientsWanTestSearchOrgWanClientEvents tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestSearchOrgWanClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.SearchOrgWanClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":{"When":"2022-12-31 23:59:59.293000+00:00","ev_type":"CLIENT_IP_ASSIGNED","metadata":{},"org_id":"b0b9f142-aaba-11e6-aafc-0242ac110002","random_mac":true,"site_id":"fc656275-b157-43fd-b922-5f4f341c19bf","text":"DHCP Ack IP 192.168.88.216","wcid":"62bbfb75-10d8-49d1-dec7-d2df91624287"},"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWanTestSearchOrgWanClientEvents1 tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestSearchOrgWanClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWan.SearchOrgWanClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":{"When":"2022-12-31 23:59:59.293000+00:00","ev_type":"CLIENT_IP_ASSIGNED","metadata":{},"org_id":"b0b9f142-aaba-11e6-aafc-0242ac110002","random_mac":true,"site_id":"fc656275-b157-43fd-b922-5f4f341c19bf","text":"DHCP Ack IP 192.168.88.216","wcid":"62bbfb75-10d8-49d1-dec7-d2df91624287"},"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWanTestSearchOrgWanClients tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestSearchOrgWanClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsClientsWan.SearchOrgWanClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"When":"2022-12-31T23:59:43.497+0000","hostname":["sonoszp"],"ip":["192.168.1.139"],"last_hostname":"sonoszp","last_ip":"192.168.1.139","mfg":"Sonos","org_id":"b4e16c72-d50e-4c03-a952-a3217e231e2c","site_id":"f688779c-e335-4f88-8d7c-9c5e9964528b","wcid":"8bbe7389-212b-c65d-2208-00fab2017936"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWanTestSearchOrgWanClients1 tests the behavior of the OrgsClientsWan
func TestOrgsClientsWanTestSearchOrgWanClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsClientsWan.SearchOrgWanClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"When":"2022-12-31T23:59:43.497+0000","hostname":["sonoszp"],"ip":["192.168.1.139"],"last_hostname":"sonoszp","last_ip":"192.168.1.139","mfg":"Sonos","org_id":"b4e16c72-d50e-4c03-a952-a3217e231e2c","site_id":"f688779c-e335-4f88-8d7c-9c5e9964528b","wcid":"8bbe7389-212b-c65d-2208-00fab2017936"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
