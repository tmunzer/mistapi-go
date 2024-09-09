package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesWxRulesTestListSiteWxRules tests the behavior of the SitesWxRules
func TestSitesWxRulesTestListSiteWxRules(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesWxRules.ListSiteWxRules(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-bfeb-53ebbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxRulesTestCreateSiteWxRule tests the behavior of the SitesWxRules
func TestSitesWxRulesTestCreateSiteWxRule(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanRule
    errBody := json.Unmarshal([]byte(`{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesWxRules.CreateSiteWxRule(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxRulesTestGetSiteWxRulesDerived tests the behavior of the SitesWxRules
func TestSitesWxRulesTestGetSiteWxRulesDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesWxRules.GetSiteWxRulesDerived(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-bfeb-53ebbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxRulesTestDeleteSiteWxRule tests the behavior of the SitesWxRules
func TestSitesWxRulesTestDeleteSiteWxRule(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesWxRules.DeleteSiteWxRule(ctx, siteId, wxruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesWxRulesTestGetSiteWxRule tests the behavior of the SitesWxRules
func TestSitesWxRulesTestGetSiteWxRule(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesWxRules.GetSiteWxRule(ctx, siteId, wxruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxRulesTestUpdateSiteWxRule tests the behavior of the SitesWxRules
func TestSitesWxRulesTestUpdateSiteWxRule(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanRule
    errBody := json.Unmarshal([]byte(`{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesWxRules.UpdateSiteWxRule(ctx, siteId, wxruleId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
