// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsMapsTestImportOrgMaps tests the behavior of the OrgsMaps
func TestOrgsMapsTestImportOrgMaps(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    autoDeviceprofileAssignment := bool(true)
    
    
    
    apiResponse, err := orgsMaps.ImportOrgMaps(ctx, orgId, &autoDeviceprofileAssignment, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":[{"action":"assigned-placed","floorplan_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","height":3,"mac":"5c5b35000001","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","orientation":45}],"floorplans":[{"action":"imported","id":"cbdb7f0b-3be0-4872-88f9-58790b509c23","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","name":"map1"}],"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","summary":{"num_ap_assigned":1,"num_inv_assigned":1,"num_map_assigned":1}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMapsTestImportOrgMaps1 tests the behavior of the OrgsMaps
func TestOrgsMapsTestImportOrgMaps1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    autoDeviceprofileAssignment := bool(true)
    
    
    
    apiResponse, err := orgsMaps.ImportOrgMaps(ctx, orgId, &autoDeviceprofileAssignment, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aps":[{"action":"assigned-placed","floorplan_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","height":3,"mac":"5c5b35000001","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","orientation":45}],"floorplans":[{"action":"imported","id":"cbdb7f0b-3be0-4872-88f9-58790b509c23","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","name":"map1"}],"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","summary":{"num_ap_assigned":1,"num_inv_assigned":1,"num_map_assigned":1}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
