// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsStatsTunnelsTestCountOrgTunnelsStats tests the behavior of the OrgsStatsTunnels
func TestOrgsStatsTunnelsTestCountOrgTunnelsStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgTunnelCountDistinctEnum("wxtunnel_id")
	mType := models.OrgTunnelTypeCountEnum("wxtunnel")
	limit := int(100)
	apiResponse, err := orgsStatsTunnels.CountOrgTunnelsStats(ctx, orgId, &distinct, &mType, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsTunnelsTestCountOrgTunnelsStats1 tests the behavior of the OrgsStatsTunnels
func TestOrgsStatsTunnelsTestCountOrgTunnelsStats1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgTunnelCountDistinctEnum("wxtunnel_id")
	mType := models.OrgTunnelTypeCountEnum("wxtunnel")
	limit := int(100)
	apiResponse, err := orgsStatsTunnels.CountOrgTunnelsStats(ctx, orgId, &distinct, &mType, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsTunnelsTestSearchOrgTunnelsStats tests the behavior of the OrgsStatsTunnels
func TestOrgsStatsTunnelsTestSearchOrgTunnelsStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	mType := models.TunnelTypeEnum("wxtunnel")
	limit := int(100)

	duration := "5m"
	sort := "timestamp"
	apiResponse, err := orgsStatsTunnels.SearchOrgTunnelsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &mType, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583,"limit":10,"results":[{"auth_algo":"hmac-md5-96","encrypt_algo":"aes-256-cbc","ike_version":"2","ip":"192.168.233.0","last_event":"down reason","mac":"020001ae9dd5","node":"node0","org_id":"78c11da8-f984-4425-bedb-a7ddd7d0f6da","peer_host":"sunnyvale1-vpn.zscalerbeta.net","peer_ip":"10.224.8.16","protocol":"ipsec","rx_bytes":150,"rx_pkts":75,"site_id":"e83e7928-eda1-4e93-82db-df3dd42ab726","tunnel_name":"Device-ipsec-1","tx_bytes":100,"tx_pkts":50,"up":true,"uptime":10,"wan_name":"wan"}],"start":1531776183,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsTunnelsTestSearchOrgTunnelsStats1 tests the behavior of the OrgsStatsTunnels
func TestOrgsStatsTunnelsTestSearchOrgTunnelsStats1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	mType := models.TunnelTypeEnum("wxtunnel")
	limit := int(100)

	duration := "5m"
	sort := "timestamp"
	apiResponse, err := orgsStatsTunnels.SearchOrgTunnelsStats(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &mType, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583,"limit":10,"results":[{"auth_algo":"hmac-md5-96","encrypt_algo":"aes-256-cbc","ike_version":"2","ip":"192.168.233.0","last_event":"down reason","mac":"020001ae9dd5","node":"node0","org_id":"78c11da8-f984-4425-bedb-a7ddd7d0f6da","peer_host":"sunnyvale1-vpn.zscalerbeta.net","peer_ip":"10.224.8.16","protocol":"ipsec","rx_bytes":150,"rx_pkts":75,"site_id":"e83e7928-eda1-4e93-82db-df3dd42ab726","tunnel_name":"Device-ipsec-1","tx_bytes":100,"tx_pkts":50,"up":true,"uptime":10,"wan_name":"wan"}],"start":1531776183,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
