package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesClientsNACTestCountSiteNacClients tests the behavior of the SitesClientsNAC
func TestSitesClientsNACTestCountSiteNacClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteNacClientsCountDistinctEnum("type")

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesClientsNac.CountSiteNacClients(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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

// TestSitesClientsNACTestCountSiteNacClientEvents tests the behavior of the SitesClientsNAC
func TestSitesClientsNACTestCountSiteNacClientEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsNac.CountSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsNACTestSearchSiteNacClientEvents tests the behavior of the SitesClientsNAC
func TestSitesClientsNACTestSearchSiteNacClientEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsNac.SearchSiteNacClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513176951,"limit":10,"results":[{"ap":"5c5b35513227","auth_type":"eap-ttls","bssid":"5c5b355fafcc","dryrun_nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264012","dryrun_nacrule_matched":true,"idp_id":"912ef72e-2239-4996-b81e-469e87a27cd6","idp_role":["itsuperusers","vip"],"mac":"ac3eb179e535","nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62","nacrule_matched":true,"nas_vendor":"juniper-mist","org_id":"27547ac2-d114-4e04-beb1-f3f1e6e81ec6","random_mac":false,"resp_attrs":["Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous"],"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","ssid":"mist_nac","timestamp":1691512031.3581879,"type":"NAC_CLIENT_PERMIT","username":"user@deaflyz.net","vlan":"750"}],"start":1512572151,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsNACTestSearchSiteNacClients tests the behavior of the SitesClientsNAC
func TestSitesClientsNACTestSearchSiteNacClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesClientsNac.SearchSiteNacClients(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513362753,"limit":3,"results":[{"ap":["5c5b35bf16bb","d4dc090041b4"],"auth_type":"eap-tls","cert_cn":["string"],"cert_issuer":["string"],"idp_id":"string","idp_role":["string"],"last_ap":"string","last_cert_cn":"string","last_cert_expiry":0,"last_cert_issuer":"string","last_nacrule_id":"string","last_nacrule_name":"string","last_nas_vendor":"string","last_ssid":"string","last_status":"string","mac":"string","nacrule_id":["string"],"nacrule_matched":true,"nacrule_name":["string"],"nas_vendor":["string"],"org_id":"31f27122-68a9-47a4-b526-8fb8a62a8acb","random_mac":true,"site_id":"832b1d74-9531-409b-ae37-4d7f3edbde92","ssid":["string"],"timestamp":1694689718.612,"type":"wireless"}],"start":1513276353,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
