// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsWxRulesTestListOrgWxRules tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestListOrgWxRules(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWxRules.ListOrgWxRules(ctx, orgId, &limit, &page)
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

// TestOrgsWxRulesTestListOrgWxRules1 tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestListOrgWxRules1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWxRules.ListOrgWxRules(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-bfeb-53ebbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxRulesTestCreateOrgWxRule tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestCreateOrgWxRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanRule
    errBody := json.Unmarshal([]byte(`{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWxRules.CreateOrgWxRule(ctx, orgId, &body)
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

// TestOrgsWxRulesTestCreateOrgWxRule1 tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestCreateOrgWxRule1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WxlanRule
    errBody := json.Unmarshal([]byte(`{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWxRules.CreateOrgWxRule(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxRulesTestDeleteOrgWxRule tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestDeleteOrgWxRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWxRules.DeleteOrgWxRule(ctx, orgId, wxruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWxRulesTestGetOrgWxRule tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestGetOrgWxRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxRules.GetOrgWxRule(ctx, orgId, wxruleId)
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

// TestOrgsWxRulesTestGetOrgWxRule1 tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestGetOrgWxRule1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wxruleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWxRules.GetOrgWxRule(ctx, orgId, wxruleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxRulesTestUpdateOrgWxRule tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestUpdateOrgWxRule(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
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
    apiResponse, err := orgsWxRules.UpdateOrgWxRule(ctx, orgId, wxruleId, &body)
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

// TestOrgsWxRulesTestUpdateOrgWxRule1 tests the behavior of the OrgsWxRules
func TestOrgsWxRulesTestUpdateOrgWxRule1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
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
    apiResponse, err := orgsWxRules.UpdateOrgWxRule(ctx, orgId, wxruleId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"action":"allow","apply_tags":["c049dfcd-0c73-5014-1c64-062e9903f1e5"],"blocked_apps":["mist","all-videos"],"created_time":0,"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"enabled":true,"for_site":true,"id":"497f6eca-6276-4993-9feb-53cbbbba6f08","modified_time":0,"order":1,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
