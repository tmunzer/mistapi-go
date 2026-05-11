// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsStatsAssetsTestListOrgAssetsStats tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestListOrgAssetsStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsStatsAssets.ListOrgAssetsStats(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"_ttl":86400,"battery_voltage":3370,"by":"asset","device_id":"00000000-0000-0000-1000-5c5b35000001","device_name":"BLE Device","eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_url":"https://www.abc.com","ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","last_seen":1480716946,"mac":"4a0222000e31","manufacture":"Asset Manufacturer Name","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","mfg_company_id":935,"mfg_data":"648520a1020000","name":"Asset Name","rssi":-45,"rssizones":[{"id":"478f6eca-6276-4993-bfeb-5bcbbbbacf08","since":0}],"service_packets":[{"data":"640","last_rx_time":1645855923,"rx_cnt":213065,"uuid":"00003e10-0000-1000-8000-00805f9b34fb"}],"temperature":23.5,"x":51.0,"y":29.0,"zones":[{"id":"477f6eca-6276-4993-bfeb-5ccbbbbadf08","since":0}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsAssetsTestListOrgAssetsStats1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestListOrgAssetsStats1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsStatsAssets.ListOrgAssetsStats(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"_ttl":86400,"battery_voltage":3370,"by":"asset","device_id":"00000000-0000-0000-1000-5c5b35000001","device_name":"BLE Device","eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_url":"https://www.abc.com","ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","last_seen":1480716946,"mac":"4a0222000e31","manufacture":"Asset Manufacturer Name","map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","mfg_company_id":935,"mfg_data":"648520a1020000","name":"Asset Name","rssi":-45,"rssizones":[{"id":"478f6eca-6276-4993-bfeb-5bcbbbbacf08","since":0}],"service_packets":[{"data":"640","last_rx_time":1645855923,"rx_cnt":213065,"uuid":"00003e10-0000-1000-8000-00805f9b34fb"}],"temperature":23.5,"x":51.0,"y":29.0,"zones":[{"id":"477f6eca-6276-4993-bfeb-5ccbbbbadf08","since":0}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	apiResponse, err := orgsStatsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil, &limit)
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

// TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestCountOrgAssetsByDistanceField1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	apiResponse, err := orgsStatsAssets.CountOrgAssetsByDistanceField(ctx, orgId, nil, &limit)
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

// TestOrgsStatsAssetsTestSearchOrgAssets tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestSearchOrgAssets(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsStatsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":1,"ibeacon_minor":1,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"476f6eca-6276-4993-bfeb-5dcbbbbaef08","since":0}],"x":0,"y":0,"zones":[{"id":"475f6eca-6276-4993-bfeb-5ecbbbbf6f08","since":0}]}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsAssetsTestSearchOrgAssets1 tests the behavior of the OrgsStatsAssets
func TestOrgsStatsAssetsTestSearchOrgAssets1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsStatsAssets.SearchOrgAssets(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":1,"ibeacon_minor":1,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","last_seen":0,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","rssizones":[{"id":"476f6eca-6276-4993-bfeb-5dcbbbbaef08","since":0}],"x":0,"y":0,"zones":[{"id":"475f6eca-6276-4993-bfeb-5ecbbbbf6f08","since":0}]}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
