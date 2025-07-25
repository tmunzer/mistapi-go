// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsStatsBGPPeersTestCountOrgBgpStats tests the behavior of the OrgsStatsBGPPeers
func TestOrgsStatsBGPPeersTestCountOrgBgpStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    apiResponse, err := orgsStatsBgpPeers.CountOrgBgpStats(ctx, orgId, &limit)
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

// TestOrgsStatsBGPPeersTestCountOrgBgpStats1 tests the behavior of the OrgsStatsBGPPeers
func TestOrgsStatsBGPPeersTestCountOrgBgpStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    apiResponse, err := orgsStatsBgpPeers.CountOrgBgpStats(ctx, orgId, &limit)
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

// TestOrgsStatsBGPPeersTestSearchOrgBgpStats tests the behavior of the OrgsStatsBGPPeers
func TestOrgsStatsBGPPeersTestSearchOrgBgpStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsBgpPeers.SearchOrgBgpStats(ctx, orgId, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"evpn_overlay":true,"for_overlay":true,"local_as":65000,"mac":"020001c04668","neighbor":"15.8.3.5","neighbor_as":65000,"neighbor_mac":"020001c04600","node":"node0","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","rx_pkts":63366,"rx_routes":60,"site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","state":"established","timestamp":1666251056.07,"tx_pkts":1735,"tx_routes":60,"up":true,"uptime":31355,"vrf_name":"default"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsBGPPeersTestSearchOrgBgpStats1 tests the behavior of the OrgsStatsBGPPeers
func TestOrgsStatsBGPPeersTestSearchOrgBgpStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsBgpPeers.SearchOrgBgpStats(ctx, orgId, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"evpn_overlay":true,"for_overlay":true,"local_as":65000,"mac":"020001c04668","neighbor":"15.8.3.5","neighbor_as":65000,"neighbor_mac":"020001c04600","node":"node0","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","rx_pkts":63366,"rx_routes":60,"site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","state":"established","timestamp":1666251056.07,"tx_pkts":1735,"tx_routes":60,"up":true,"uptime":31355,"vrf_name":"default"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
