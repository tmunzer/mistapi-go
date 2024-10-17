package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsUserMACsTestCreateOrgUserMacs tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestCreateOrgUserMacs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsUserMaCs.CreateOrgUserMacs(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"added":["921b638445cd"],"errors":["921b638445ce - mac invalid","921b638445cf - mac already provided"],"updated":["721b638445ef","721b638445ee"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestSearchOrgUserMacs tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestSearchOrgUserMacs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	page := int(1)
	apiResponse, err := orgsUserMaCs.SearchOrgUserMacs(ctx, orgId, nil, nil, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":100,"page":1,"results":[{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestDeleteOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestDeleteOrgUserMac(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsUserMaCs.DeleteOrgUserMac(ctx, orgId, usermacId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsUserMACsTestGetOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestGetOrgUserMac(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsUserMaCs.GetOrgUserMac(ctx, orgId, usermacId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestUpdateOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestUpdateOrgUserMac(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsUserMaCs.UpdateOrgUserMac(ctx, orgId, usermacId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
