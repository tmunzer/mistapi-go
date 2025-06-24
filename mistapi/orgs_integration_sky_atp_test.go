package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsIntegrationSkyATPTestDeleteOrgSkyAtpIntegration tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestDeleteOrgSkyAtpIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsIntegrationSkyAtp.DeleteOrgSkyAtpIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIntegrationSkyATPTestGetOrgSkyAtpIntegration tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestGetOrgSkyAtpIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationSkyAtp.GetOrgSkyAtpIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsIntegrationSkyATPTestGetOrgSkyAtpIntegration1 tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestGetOrgSkyAtpIntegration1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationSkyAtp.GetOrgSkyAtpIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsIntegrationSkyATPTestSetupOrgAtpIntegration tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestSetupOrgAtpIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountSkyatpConfig
    errBody := json.Unmarshal([]byte(`{"password":"foryoureyesonly","realm":"mist-team","username":"john@abc.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.SetupOrgAtpIntegration(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"Example":{"value":{"secintel":{"third_party_threat_feeds":["block_list"]},"secintel_allowlist_url":"https://papi.s3.amazonaws.com/secintel_allowlist/xxx...","secintel_blocklist_url":"https://papi.s3.amazonaws.com/secintel_blocklist/xxx..."}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationSkyATPTestSetupOrgAtpIntegration1 tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestSetupOrgAtpIntegration1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountSkyatpConfig
    errBody := json.Unmarshal([]byte(`{"password":"foryoureyesonly","realm":"mist-team","username":"john@abc.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.SetupOrgAtpIntegration(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpIntegration tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountSkyatpData
    errBody := json.Unmarshal([]byte(`{"secintel":{"third_party_threat_feeds":["block_list"]}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpIntegration(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpIntegration1 tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpIntegration1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountSkyatpData
    errBody := json.Unmarshal([]byte(`{"secintel":{"third_party_threat_feeds":["block_list"]}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpIntegration(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpAllowedList tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpAllowedList(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SkyatpList
    errBody := json.Unmarshal([]byte(`{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpAllowedList(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpAllowedList1 tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpAllowedList1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SkyatpList
    errBody := json.Unmarshal([]byte(`{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpAllowedList(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpBlockedList tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpBlockedList(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SkyatpList
    errBody := json.Unmarshal([]byte(`{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpBlockedList(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationSkyATPTestUdpateOrgAtpBlockedList1 tests the behavior of the OrgsIntegrationSkyATP
func TestOrgsIntegrationSkyATPTestUdpateOrgAtpBlockedList1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SkyatpList
    errBody := json.Unmarshal([]byte(`{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIntegrationSkyAtp.UdpateOrgAtpBlockedList(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"domains":[{"comment":"restricted","domain":"unsafe.xxx"}],"ips":[{"comment":"nas","ip":"10.1.3.5"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
