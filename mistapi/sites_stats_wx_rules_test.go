// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesStatsWxRulesTestGetSiteWxRulesUsage tests the behavior of the SitesStatsWxRules
func TestSitesStatsWxRulesTestGetSiteWxRulesUsage(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesStatsWxRules.GetSiteWxRulesUsage(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","client_mac":["3bbbf819bb6f","bd96cbc4910f"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_wxtags":["d4134466-eec0-3756-6765-381c728a6037","1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"name":"Guest","order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"],"usage":{"1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3":{"num_flows":60},"d4134466-eec0-3756-6765-381c728a6037":{"num_flows":60}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsWxRulesTestGetSiteWxRulesUsage1 tests the behavior of the SitesStatsWxRules
func TestSitesStatsWxRulesTestGetSiteWxRulesUsage1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesStatsWxRules.GetSiteWxRulesUsage(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"action":"allow","client_mac":["3bbbf819bb6f","bd96cbc4910f"],"dst_allow_wxtags":["fff34466-eec0-3756-6765-381c728a6037","eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_deny_wxtags":["aaa34466-eec0-3756-6765-381c728a6037","bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"dst_wxtags":["d4134466-eec0-3756-6765-381c728a6037","1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3"],"name":"Guest","order":1,"src_wxtags":["8bfc2490-d726-3587-038d-cb2e71bd2330","3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"],"usage":{"1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3":{"num_flows":60},"d4134466-eec0-3756-6765-381c728a6037":{"num_flows":60}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
