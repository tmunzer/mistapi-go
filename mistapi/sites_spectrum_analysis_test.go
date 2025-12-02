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

// TestSitesSpectrumAnalysisTestGetSiteRunningSpectrumAnalysis tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestGetSiteRunningSpectrumAnalysis(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesSpectrumAnalysis.GetSiteRunningSpectrumAnalysis(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band":"5","device_id":"00000000-0000-0000-1000-5c5b35bd76bb","duration":600,"format":"stream","started_time":1435080709}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSpectrumAnalysisTestGetSiteRunningSpectrumAnalysis1 tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestGetSiteRunningSpectrumAnalysis1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesSpectrumAnalysis.GetSiteRunningSpectrumAnalysis(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band":"5","device_id":"00000000-0000-0000-1000-5c5b35bd76bb","duration":600,"format":"stream","started_time":1435080709}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSpectrumAnalysisTestInitiateSiteAnalyzeSpectrum tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestInitiateSiteAnalyzeSpectrum(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.SpectrumAnalysis
	errBody := json.Unmarshal([]byte(`{"band":"5","device_id":"00000000-0000-0000-1000-5c5b35bd76bb","duration":600,"format":"stream"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesSpectrumAnalysis.InitiateSiteAnalyzeSpectrum(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesSpectrumAnalysisTestInitiateSiteAnalyzeSpectrum1 tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestInitiateSiteAnalyzeSpectrum1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.SpectrumAnalysis
	errBody := json.Unmarshal([]byte(`{"band":"5","device_id":"00000000-0000-0000-1000-5c5b35bd76bb","duration":600,"format":"stream"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesSpectrumAnalysis.InitiateSiteAnalyzeSpectrum(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesSpectrumAnalysisTestListSiteSpectrumAnalysis tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestListSiteSpectrumAnalysis(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesSpectrumAnalysis.ListSiteSpectrumAnalysis(ctx, siteId, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1694708579,"limit":10,"results":[{"band":"5","channel_usage":[{"channel":36,"noise":-78,"non_wifi":0.08,"wifi":0.13}],"fft_samples":[{"frequency":2437,"rssi":-92,"signal7":-93}],"mac":"5c5b35bd76bb","org_id":"f2695c32-0e83-4936-b1b2-96fc88051213","timestamp":1694098696}],"start":1694622179,"total":4}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSpectrumAnalysisTestListSiteSpectrumAnalysis1 tests the behavior of the SitesSpectrumAnalysis
func TestSitesSpectrumAnalysisTestListSiteSpectrumAnalysis1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesSpectrumAnalysis.ListSiteSpectrumAnalysis(ctx, siteId, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1694708579,"limit":10,"results":[{"band":"5","channel_usage":[{"channel":36,"noise":-78,"non_wifi":0.08,"wifi":0.13}],"fft_samples":[{"frequency":2437,"rssi":-92,"signal7":-93}],"mac":"5c5b35bd76bb","org_id":"f2695c32-0e83-4936-b1b2-96fc88051213","timestamp":1694098696}],"start":1694622179,"total":4}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
