// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsStatsSitesTestListOrgSiteStats tests the behavior of the OrgsStatsSites
func TestOrgsStatsSitesTestListOrgSiteStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsSites.ListOrgSiteStats(ctx, orgId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"address":"1601 S De Anza Blvd, Cupertino, CA 95014, USA","alarmtemplate_id":null,"analyticEnabled":true,"aptemplate_id":null,"country_code":"US","created_time":1472591606,"engagementEnabled":true,"gatewaytemplate_id":"e571f2a2-d748-4ad4-bd6c-895467957c21","id":"83bc290a-b76d-47fa-a294-d34e47f30f7f","lat":37.295553,"latlng":{"lat":37.295553,"lng":-122.033007},"lng":-122.033007,"modified_time":1728057857,"msp_id":"a9af4951-a1de-4520-b398-c95a58947349","name":"Live-Demo","networktemplate_id":"964cb213-deb2-469d-8c1e-a5f8661c6886","notes":"This site is used for demonstration purposes.","num_ap":17,"num_ap_connected":14,"num_clients":14,"num_devices":26,"num_devices_connected":22,"num_gateway":1,"num_gateway_connected":1,"num_switch":8,"num_switch_connected":7,"org_id":"b9814b40-ac4b-4424-86a8-b787eb68b86a","rftemplate_id":"2c134c07-3c57-46b3-a53b-8aea92ed7234","secpolicy_id":null,"sitegroup_ids":["5644a432-eea9-4a2f-a30a-ddaf4dbc79cf","5fc0f305-f626-49db-8869-10b87f201bba","882796ef-190b-405e-98ef-cb487140cf64"],"sitetemplate_id":null,"timezone":"America/Los_Angeles","tzoffset":960}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsSitesTestListOrgSiteStats1 tests the behavior of the OrgsStatsSites
func TestOrgsStatsSitesTestListOrgSiteStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsSites.ListOrgSiteStats(ctx, orgId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"address":"1601 S De Anza Blvd, Cupertino, CA 95014, USA","alarmtemplate_id":null,"analyticEnabled":true,"aptemplate_id":null,"country_code":"US","created_time":1472591606,"engagementEnabled":true,"gatewaytemplate_id":"e571f2a2-d748-4ad4-bd6c-895467957c21","id":"83bc290a-b76d-47fa-a294-d34e47f30f7f","lat":37.295553,"latlng":{"lat":37.295553,"lng":-122.033007},"lng":-122.033007,"modified_time":1728057857,"msp_id":"a9af4951-a1de-4520-b398-c95a58947349","name":"Live-Demo","networktemplate_id":"964cb213-deb2-469d-8c1e-a5f8661c6886","notes":"This site is used for demonstration purposes.","num_ap":17,"num_ap_connected":14,"num_clients":14,"num_devices":26,"num_devices_connected":22,"num_gateway":1,"num_gateway_connected":1,"num_switch":8,"num_switch_connected":7,"org_id":"b9814b40-ac4b-4424-86a8-b787eb68b86a","rftemplate_id":"2c134c07-3c57-46b3-a53b-8aea92ed7234","secpolicy_id":null,"sitegroup_ids":["5644a432-eea9-4a2f-a30a-ddaf4dbc79cf","5fc0f305-f626-49db-8869-10b87f201bba","882796ef-190b-405e-98ef-cb487140cf64"],"sitetemplate_id":null,"timezone":"America/Los_Angeles","tzoffset":960}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
