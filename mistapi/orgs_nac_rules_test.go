package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsNACRulesTestListOrgNacRules tests the behavior of the OrgsNACRules
func TestOrgsNACRulesTestListOrgNacRules(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsNacRules.ListOrgNacRules(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","apply_tags":["string"],"created_time":0,"id":"455f6eca-6276-4993-bfeb-53cbbbba6208","matching":{"auth_type":"eap-tls","nactags":["string"],"port_types":["wireless"],"site_ids":["454f6eca-6276-4993-bfeb-53cbbbba6308"],"sitegroup_ids":["453f6eca-6276-4993-bfeb-53cbbbba6408"]},"modified_time":0,"name":"string","not_matching":{"auth_type":"eap-tls","nactags":["string"],"port_types":["wireless"],"site_ids":["452f6eca-6276-4993-bfeb-53cbbbba6508"],"sitegroup_ids":["451f6eca-6276-4993-bfeb-53cbbbba6608"]},"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACRulesTestCreateOrgNacRule tests the behavior of the OrgsNACRules
func TestOrgsNACRulesTestCreateOrgNacRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NacRule
    errBody := json.Unmarshal([]byte(`{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5\""],"matching":{"auth_type":"eap-tls","nactags":["041d5d36-716c-4cfb-4988-3857c6aa14a2","a809a97f-d599-f812-eb8c-c3f84aabf6ba"],"port_types":["wired"],"site_ids":["bb19fc3e-4124-4b57-80d9-c3f6edce47c4","bb19fc3e-6564-4b57-80d9-c3f6edce47c1"],"sitegroup_ids":["bb19fc3e-4124-4b57-80d9-c3f6edce47c4","bb19fc3e-6564-4b57-80d9-c3f6edce47c1"]},"name":"name1","not_matching":{},"order":1}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsNacRules.CreateOrgNacRule(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsNACRulesTestDeleteOrgNacRule tests the behavior of the OrgsNACRules
func TestOrgsNACRulesTestDeleteOrgNacRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsNacRules.DeleteOrgNacRule(ctx, orgId, nacruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACRulesTestGetOrgNacRule tests the behavior of the OrgsNACRules
func TestOrgsNACRulesTestGetOrgNacRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsNacRules.GetOrgNacRule(ctx, orgId, nacruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsNACRulesTestUpdateOrgNacRule tests the behavior of the OrgsNACRules
func TestOrgsNACRulesTestUpdateOrgNacRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsNacRules.UpdateOrgNacRule(ctx, orgId, nacruleId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
