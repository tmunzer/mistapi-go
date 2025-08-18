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

// TestSitesUISettingsTestListSiteUiSettings tests the behavior of the SitesUISettings
func TestSitesUISettingsTestListSiteUiSettings(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.ListSiteUiSettings(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestListSiteUiSettings1 tests the behavior of the SitesUISettings
func TestSitesUISettingsTestListSiteUiSettings1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.ListSiteUiSettings(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestCreateSiteUiSettings tests the behavior of the SitesUISettings
func TestSitesUISettingsTestCreateSiteUiSettings(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UiSettings
	errBody := json.Unmarshal([]byte(`{"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"name":"New Databoard","purpose":"databoard","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesUiSettings.CreateSiteUiSettings(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestCreateSiteUiSettings1 tests the behavior of the SitesUISettings
func TestSitesUISettingsTestCreateSiteUiSettings1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UiSettings
	errBody := json.Unmarshal([]byte(`{"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"name":"New Databoard","purpose":"databoard","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesUiSettings.CreateSiteUiSettings(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestListSiteUiSettingDerived tests the behavior of the SitesUISettings
func TestSitesUISettingsTestListSiteUiSettingDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.ListSiteUiSettingDerived(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestListSiteUiSettingDerived1 tests the behavior of the SitesUISettings
func TestSitesUISettingsTestListSiteUiSettingDerived1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.ListSiteUiSettingDerived(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestDeleteSiteUiSetting tests the behavior of the SitesUISettings
func TestSitesUISettingsTestDeleteSiteUiSetting(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesUiSettings.DeleteSiteUiSetting(ctx, siteId, uisettingId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesUISettingsTestGetSiteUiSetting tests the behavior of the SitesUISettings
func TestSitesUISettingsTestGetSiteUiSetting(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.GetSiteUiSetting(ctx, siteId, uisettingId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestGetSiteUiSetting1 tests the behavior of the SitesUISettings
func TestSitesUISettingsTestGetSiteUiSetting1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesUiSettings.GetSiteUiSetting(ctx, siteId, uisettingId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestUpdateSiteUiSetting tests the behavior of the SitesUISettings
func TestSitesUISettingsTestUpdateSiteUiSetting(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UiSettings
	errBody := json.Unmarshal([]byte(`{"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"name":"New Databoard","purpose":"databoard","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesUiSettings.UpdateSiteUiSetting(ctx, siteId, uisettingId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesUISettingsTestUpdateSiteUiSetting1 tests the behavior of the SitesUISettings
func TestSitesUISettingsTestUpdateSiteUiSetting1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UiSettings
	errBody := json.Unmarshal([]byte(`{"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"name":"New Databoard","purpose":"databoard","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesUiSettings.UpdateSiteUiSetting(ctx, siteId, uisettingId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1508823803,"defaultScopeId":"67970e46-4e12-11e6-9188-0242ad112847","defaultScopeType":"site","defaultTimeRange":{"end":1508828400,"endDate":"10/23/2017","interval":"1d","name":"This Week","shortName":"thisWeek","start":1508655600,"usePreset":true},"description":"Description of the databoard","for_site":true,"id":"3bdcc7e8-c04d-4512-b4fc-093da9057eb0","isCustomDataboard":true,"isScopeLinked":true,"isTimeRangeLinked":true,"modified_time":0,"name":"New Databoard","org_id":"cc079380-5029-4d4a-9125-858de85731ff","purpose":"databoard","site_id":"67970e46-4e12-11e6-9188-0242ad112847","tiles":[{"chartBand":"2.4 ghz","chartColor":"#00B4AD","chartDirection":"tx + rx","chartRankBy":"string","chartType":"timeSeries","colspan":5,"column":1,"hideEmptyRows":true,"id":"7a9ab38c-cfc3-483d-b51a-0aec571fadc0","metric":{"apiName":"client_dhcp_latency"},"name":"New Analysis","row":1,"rowspan":2,"scopeId":"e0c767834b4c","scopeType":"client","sortedColumns":null,"timeRange":{"end":1508823743,"endDate":"10/23/2017","interval":"1d","name":"Past 7 Days","shortName":"7d","start":1508223600,"usePreset":true},"trendType":"line","vizType":"averageTimeSeriesChart"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
