package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsClientsNACTestCountOrgNacClients tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestCountOrgNacClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgNacClientsCountDistinctEnum("type")
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.CountOrgNacClients(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsNACTestCountOrgNacClients1 tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestCountOrgNacClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgNacClientsCountDistinctEnum("type")
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.CountOrgNacClients(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsNACTestCountOrgNacClientEvents tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestCountOrgNacClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.CountOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsNACTestCountOrgNacClientEvents1 tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestCountOrgNacClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.CountOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, &duration, &limit)
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

// TestOrgsClientsNACTestSearchOrgNacClientEvents tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestSearchOrgNacClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    respAttrs := []string{ "Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous" }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.SearchOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, respAttrs, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513176951,"limit":10,"results":[{"ap":"5c5b35513227","auth_type":"eap-ttls","bssid":"5c5b355fafcc","dryrun_nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264012","dryrun_nacrule_matched":true,"idp_id":"912ef72e-2239-4996-b81e-469e87a27cd6","idp_role":["itsuperusers","vip"],"mac":"ac3eb179e535","nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62","nacrule_matched":true,"nas_vendor":"juniper-mist","org_id":"27547ac2-d114-4e04-beb1-f3f1e6e81ec6","random_mac":false,"resp_attrs":["Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous"],"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","ssid":"mist_nac","timestamp":1691512031.358188,"type":"NAC_CLIENT_PERMIT","username":"user@deaflyz.net","vlan":"750"}],"start":1512572151,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsNACTestSearchOrgNacClientEvents1 tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestSearchOrgNacClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    respAttrs := []string{ "Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous" }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsNac.SearchOrgNacClientEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, respAttrs, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513176951,"limit":10,"results":[{"ap":"5c5b35513227","auth_type":"eap-ttls","bssid":"5c5b355fafcc","dryrun_nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264012","dryrun_nacrule_matched":true,"idp_id":"912ef72e-2239-4996-b81e-469e87a27cd6","idp_role":["itsuperusers","vip"],"mac":"ac3eb179e535","nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62","nacrule_matched":true,"nas_vendor":"juniper-mist","org_id":"27547ac2-d114-4e04-beb1-f3f1e6e81ec6","random_mac":false,"resp_attrs":["Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous"],"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","ssid":"mist_nac","timestamp":1691512031.358188,"type":"NAC_CLIENT_PERMIT","username":"user@deaflyz.net","vlan":"750"}],"start":1512572151,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsNACTestSearchOrgNacClients tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestSearchOrgNacClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    status := models.NacClientLastStatusEnum("permitted")
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsClientsNac.SearchOrgNacClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &status, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513362753,"limit":3,"results":[{"ap":["5c5b35bf16bb","d4dc090041b4"],"auth_type":"eap-tls","cert_cn":["john@mycorp.net"],"cert_issuer":["/C=US/ST=CA/CN=MyCorp"],"client_ip":["10.7.51.74"],"idp_id":"string","idp_role":["string"],"last_ap":"string","last_cert_cn":"john@mycorp.net","last_cert_expiry":1746711240,"last_cert_issuer":"/C=US/ST=CA/CN=MyCorp","last_cert_serial":"2c63510123456789","last_cert_subject":"/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net","last_client_ip":"10.7.51.74","last_nacrule_id":"603b62db-d839-4152-9f7f-f2578443de8d","last_nacrule_name":"Wireless Cert Auth","last_nas_vendor":"juniper-mist","last_ssid":"string","last_status":"permitted","mac":"string","nacrule_id":["603b62db-d839-4152-9f7f-f2578443de8d"],"nacrule_matched":true,"nacrule_name":["Wireless Cert Auth"],"nas_vendor":["juniper-mist"],"org_id":"31f27122-68a9-47a4-b526-8fb8a62a8acb","random_mac":true,"site_id":"832b1d74-9531-409b-ae37-4d7f3edbde92","ssid":["string"],"timestamp":1694689718.612,"type":"wireless","usermac_label":["non-compliant","building26","floor52"]}],"start":1513276353,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsNACTestSearchOrgNacClients1 tests the behavior of the OrgsClientsNAC
func TestOrgsClientsNACTestSearchOrgNacClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    status := models.NacClientLastStatusEnum("permitted")
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsClientsNac.SearchOrgNacClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &status, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513362753,"limit":3,"results":[{"ap":["5c5b35bf16bb","d4dc090041b4"],"auth_type":"eap-tls","cert_cn":["john@mycorp.net"],"cert_issuer":["/C=US/ST=CA/CN=MyCorp"],"client_ip":["10.7.51.74"],"idp_id":"string","idp_role":["string"],"last_ap":"string","last_cert_cn":"john@mycorp.net","last_cert_expiry":1746711240,"last_cert_issuer":"/C=US/ST=CA/CN=MyCorp","last_cert_serial":"2c63510123456789","last_cert_subject":"/C=US/O=MyCorp/CN=john@mycorp.net/emailAddress=john@mycorp.net","last_client_ip":"10.7.51.74","last_nacrule_id":"603b62db-d839-4152-9f7f-f2578443de8d","last_nacrule_name":"Wireless Cert Auth","last_nas_vendor":"juniper-mist","last_ssid":"string","last_status":"permitted","mac":"string","nacrule_id":["603b62db-d839-4152-9f7f-f2578443de8d"],"nacrule_matched":true,"nacrule_name":["Wireless Cert Auth"],"nas_vendor":["juniper-mist"],"org_id":"31f27122-68a9-47a4-b526-8fb8a62a8acb","random_mac":true,"site_id":"832b1d74-9531-409b-ae37-4d7f3edbde92","ssid":["string"],"timestamp":1694689718.612,"type":"wireless","usermac_label":["non-compliant","building26","floor52"]}],"start":1513276353,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
