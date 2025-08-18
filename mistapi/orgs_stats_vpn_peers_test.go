// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsStatsVPNPeersTestCountOrgPeerPathStats tests the behavior of the OrgsStatsVPNPeers
func TestOrgsStatsVPNPeersTestCountOrgPeerPathStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsVpnPeers.CountOrgPeerPathStats(ctx, orgId, nil, nil, nil, &duration, &limit)
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

// TestOrgsStatsVPNPeersTestCountOrgPeerPathStats1 tests the behavior of the OrgsStatsVPNPeers
func TestOrgsStatsVPNPeersTestCountOrgPeerPathStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsVpnPeers.CountOrgPeerPathStats(ctx, orgId, nil, nil, nil, &duration, &limit)
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

// TestOrgsStatsVPNPeersTestSearchOrgPeerPathStats tests the behavior of the OrgsStatsVPNPeers
func TestOrgsStatsVPNPeersTestSearchOrgPeerPathStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsVpnPeers.SearchOrgPeerPathStats(ctx, orgId, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1619518989.4989712,"limit":10,"results":[{"auth_algo":"hmac-sha1-96","enc_algo":"aes-cbc-128","ike_version":"1","is_active":true,"last_seen":1619518709.222,"mac":"020001c04668","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","peer_mac":"020001367edd","peer_port_id":"DC_Internet","peer_site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","port_id":"Lte","site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","type":"svr","up":true,"uptime":1527128046},{"is_active":true,"last_seen":1619518709.222,"latency":91,"mac":"020001c04668","mos":4,"mtu":1500,"org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","peer_mac":"020001367edd","peer_port_id":"DC_Internet","peer_router_name":"RIDCBBP1","peer_site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","port_id":"Lte","router_name":"RIST01544AP1","site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","type":"svr","up":true,"uptime":1527128046}],"start":1619518689.4989705,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsVPNPeersTestSearchOrgPeerPathStats1 tests the behavior of the OrgsStatsVPNPeers
func TestOrgsStatsVPNPeersTestSearchOrgPeerPathStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsStatsVpnPeers.SearchOrgPeerPathStats(ctx, orgId, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1619518989.4989712,"limit":10,"results":[{"auth_algo":"hmac-sha1-96","enc_algo":"aes-cbc-128","ike_version":"1","is_active":true,"last_seen":1619518709.222,"mac":"020001c04668","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","peer_mac":"020001367edd","peer_port_id":"DC_Internet","peer_site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","port_id":"Lte","site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","type":"svr","up":true,"uptime":1527128046},{"is_active":true,"last_seen":1619518709.222,"latency":91,"mac":"020001c04668","mos":4,"mtu":1500,"org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","peer_mac":"020001367edd","peer_port_id":"DC_Internet","peer_router_name":"RIDCBBP1","peer_site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","port_id":"Lte","router_name":"RIST01544AP1","site_id":"725a8d34-a126-4f2c-b990-d1219421cb75","type":"svr","up":true,"uptime":1527128046}],"start":1619518689.4989705,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
