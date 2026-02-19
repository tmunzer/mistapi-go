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

// TestOrgsStatsPortsTestCountOrgSwOrGwPorts tests the behavior of the OrgsStatsPorts
func TestOrgsStatsPortsTestCountOrgSwOrGwPorts(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SitePortsCountDistinctEnum("mac")
	fullDuplex := bool(true)
	mac := "5c5b350e0410"
	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txMcastPkts := int(1000000)
	txBcastPkts := int(1000000)
	rxMcastPkts := int(1000000)
	rxBcastPkts := int(1000000)
	speed := int(1000000000)

	up := bool(true)
	siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsStatsPorts.CountOrgSwOrGwPorts(ctx, orgId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, &siteId, nil, nil, &duration, &limit)
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

// TestOrgsStatsPortsTestCountOrgSwOrGwPorts1 tests the behavior of the OrgsStatsPorts
func TestOrgsStatsPortsTestCountOrgSwOrGwPorts1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SitePortsCountDistinctEnum("mac")
	fullDuplex := bool(true)
	mac := "5c5b350e0410"
	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txMcastPkts := int(1000000)
	txBcastPkts := int(1000000)
	rxMcastPkts := int(1000000)
	rxBcastPkts := int(1000000)
	speed := int(1000000000)

	up := bool(true)
	siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsStatsPorts.CountOrgSwOrGwPorts(ctx, orgId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, &siteId, nil, nil, &duration, &limit)
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

// TestOrgsStatsPortsTestSearchOrgSwOrGwPorts tests the behavior of the OrgsStatsPorts
func TestOrgsStatsPortsTestSearchOrgSwOrGwPorts(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceType := models.SearchOrgSwOrGwPortsTypeEnum("all")

	lteImsi := "310260000000001"
	lteIccid := "89014103211118510720"
	lteImei := "123456789012345"

	xcvrPartNumber := "SFP-10G-SR"
	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsStatsPorts.SearchOrgSwOrGwPorts(ctx, orgId, &deviceType, nil, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &xcvrPartNumber, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsStatsPortsTestSearchOrgSwOrGwPorts1 tests the behavior of the OrgsStatsPorts
func TestOrgsStatsPortsTestSearchOrgSwOrGwPorts1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceType := models.SearchOrgSwOrGwPortsTypeEnum("all")

	lteImsi := "310260000000001"
	lteIccid := "89014103211118510720"
	lteImei := "123456789012345"

	xcvrPartNumber := "SFP-10G-SR"
	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsStatsPorts.SearchOrgSwOrGwPorts(ctx, orgId, &deviceType, nil, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &xcvrPartNumber, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
