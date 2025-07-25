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

// TestOrgsLicensesTestClaimOrgLicense tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestClaimOrgLicense(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.ClaimActivation
    errBody := json.Unmarshal([]byte(`{"code":"ZHT3K-H36DT-MG85D-M61AC","type":"all"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsLicenses.ClaimOrgLicense(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"inventory_pending":[{"mac":"5c5b35000012"}],"license_added":[{"end":1520380800,"quantity":180,"start":1504828800,"type":"SUB-MAN"},{"end":1520380800,"quantity":120,"start":1504828800,"type":"SUB-LOC"}],"license_duplicated":[{"end":1520380800,"quantity":180,"start":1504828800,"type":"SUB-MAN"}],"license_error":[{"order":"00000464","reason":""}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestClaimOrgLicense1 tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestClaimOrgLicense1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.ClaimActivation
    errBody := json.Unmarshal([]byte(`{"code":"ZHT3K-H36DT-MG85D-M61AC","type":"all"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsLicenses.ClaimOrgLicense(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"inventory_pending":[{"mac":"5c5b35000012"}],"license_added":[{"end":1520380800,"quantity":180,"start":1504828800,"type":"SUB-MAN"},{"end":1520380800,"quantity":120,"start":1504828800,"type":"SUB-LOC"}],"license_duplicated":[{"end":1520380800,"quantity":180,"start":1504828800,"type":"SUB-MAN"}],"license_error":[{"order":"00000464","reason":""}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestGetOrgLicenseAsyncClaimStatus tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicenseAsyncClaimStatus(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    detail := bool(true)
    apiResponse, err := orgsLicenses.GetOrgLicenseAsyncClaimStatus(ctx, orgId, &detail)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"completed":["000000000022","000000000011"],"details":[{"mac":"000000000022","status":"added","timestamp":1709598053}],"failed":0,"incompleted":[],"processed":2,"scheduled_at":1709598052,"status":"done","succeed":2,"timestamp":1709598053,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestGetOrgLicenseAsyncClaimStatus1 tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicenseAsyncClaimStatus1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    detail := bool(true)
    apiResponse, err := orgsLicenses.GetOrgLicenseAsyncClaimStatus(ctx, orgId, &detail)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"completed":["000000000022","000000000011"],"details":[{"mac":"000000000022","status":"added","timestamp":1709598053}],"failed":0,"incompleted":[],"processed":2,"scheduled_at":1709598052,"status":"done","succeed":2,"timestamp":1709598053,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestGetOrgLicensesSummary tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicensesSummary(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsLicenses.GetOrgLicensesSummary(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"amendments":[{"created_time":1644684842,"end_time":1744156799,"id":"ff0a14f6-1234-5678-90ab-c8e64d4bc6c6","modified_time":1644684842,"quantity":-1,"start_time":1632873600,"subscription_id":"VNA-000000af","type":"SUB-VNA"},{"created_time":1644684842,"end_time":1744156799,"id":"c1c28812-1234-5678-90ab-dc95680da61e","modified_time":1644684842,"quantity":-1,"start_time":1632873600,"subscription_id":"MAN-000008be","type":"SUB-MAN"},{"created_time":1644684842,"end_time":1744243199,"id":"96c0a41f-1234-5678-90ab-afe74817e9fd","modified_time":1644684842,"quantity":-1,"start_time":1586476800,"subscription_id":"EX24-000000bc","type":"SUB-EX24"}],"entitled":{"SUB-ENG":26,"SUB-EX24":9,"SUB-MAN":26,"SUB-VNA":26},"licenses":[{"created_time":1555353534,"end_time":1586822399,"id":"693a41a6-1234-5678-90ab-f53dbd3a31c0","modified_time":1555353534,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":0,"start_time":1555286400,"subscription_id":"VNA-000000aa","type":"SUB-VNA"},{"created_time":1576132516,"end_time":1586822399,"id":"656607cf-1234-5678-90ab-fc9035614ea5","modified_time":1576132516,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":8,"remaining_quantity":0,"start_time":1576022400,"subscription_id":"VNA-000000ab","type":"SUB-VNA"},{"created_time":1579204568,"end_time":1730764800,"id":"db50d0bc-1234-5678-90ab-e439958cb06b","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"MAN-000000ac","type":"SUB-MAN"},{"created_time":1579204568,"end_time":1730764800,"id":"2ff9e84a-1234-5678-90ab-fb9ec0726e01","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"ENG-000000ad","type":"SUB-ENG"},{"created_time":1579204568,"end_time":1730764800,"id":"16df7ea6-1234-5678-90ab-78018cd4024d","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"VNA-000000ae","type":"SUB-VNA"},{"created_time":1586237081,"end_time":1744243199,"id":"1b6f68d5-1234-5678-90ab-70d3e6d18c73","modified_time":1586237081,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"VNA-000000af","type":"SUB-VNA"},{"created_time":1586237097,"end_time":1744243199,"id":"1375c9bf-1234-5678-90ab-9c636708c89e","modified_time":1586237097,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"MAN-000000ba","type":"SUB-MAN"},{"created_time":1586237137,"end_time":1744243199,"id":"5974e979-1234-5678-90ab-438f833ec1c9","modified_time":1586237137,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"ENG-000000bb","type":"SUB-ENG"},{"created_time":1629947267,"end_time":1744243199,"id":"340a9cb3-1234-5678-90ab-b009344dbf3c","modified_time":1629947267,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":10,"remaining_quantity":9,"start_time":1586476800,"subscription_id":"EX24-000000bc","type":"SUB-EX24"},{"created_time":1632941870,"end_time":1744156799,"id":"9b599b0f-1234-5678-90ab-406081b58e7f","modified_time":1632941870,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":10,"remaining_quantity":10,"start_time":1632873600,"subscription_id":"ENG-000000bd","type":"SUB-ENG"},{"created_time":1632941882,"end_time":1744156799,"id":"d6d8ead3-1234-5678-90ab-98badeac7287","modified_time":1632941882,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":11,"remaining_quantity":9,"start_time":1632873600,"subscription_id":"MAN-000008be","type":"SUB-MAN"}],"summary":{"SUB-ENG":18,"SUB-EX24":3,"SUB-MAN":22,"SUB-VNA":20}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestGetOrgLicensesSummary1 tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicensesSummary1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsLicenses.GetOrgLicensesSummary(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"amendments":[{"created_time":1644684842,"end_time":1744156799,"id":"ff0a14f6-1234-5678-90ab-c8e64d4bc6c6","modified_time":1644684842,"quantity":-1,"start_time":1632873600,"subscription_id":"VNA-000000af","type":"SUB-VNA"},{"created_time":1644684842,"end_time":1744156799,"id":"c1c28812-1234-5678-90ab-dc95680da61e","modified_time":1644684842,"quantity":-1,"start_time":1632873600,"subscription_id":"MAN-000008be","type":"SUB-MAN"},{"created_time":1644684842,"end_time":1744243199,"id":"96c0a41f-1234-5678-90ab-afe74817e9fd","modified_time":1644684842,"quantity":-1,"start_time":1586476800,"subscription_id":"EX24-000000bc","type":"SUB-EX24"}],"entitled":{"SUB-ENG":26,"SUB-EX24":9,"SUB-MAN":26,"SUB-VNA":26},"licenses":[{"created_time":1555353534,"end_time":1586822399,"id":"693a41a6-1234-5678-90ab-f53dbd3a31c0","modified_time":1555353534,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":0,"start_time":1555286400,"subscription_id":"VNA-000000aa","type":"SUB-VNA"},{"created_time":1576132516,"end_time":1586822399,"id":"656607cf-1234-5678-90ab-fc9035614ea5","modified_time":1576132516,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":8,"remaining_quantity":0,"start_time":1576022400,"subscription_id":"VNA-000000ab","type":"SUB-VNA"},{"created_time":1579204568,"end_time":1730764800,"id":"db50d0bc-1234-5678-90ab-e439958cb06b","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"MAN-000000ac","type":"SUB-MAN"},{"created_time":1579204568,"end_time":1730764800,"id":"2ff9e84a-1234-5678-90ab-fb9ec0726e01","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"ENG-000000ad","type":"SUB-ENG"},{"created_time":1579204568,"end_time":1730764800,"id":"16df7ea6-1234-5678-90ab-78018cd4024d","modified_time":1579204568,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":2,"remaining_quantity":2,"start_time":1572998400,"subscription_id":"VNA-000000ae","type":"SUB-VNA"},{"created_time":1586237081,"end_time":1744243199,"id":"1b6f68d5-1234-5678-90ab-70d3e6d18c73","modified_time":1586237081,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"VNA-000000af","type":"SUB-VNA"},{"created_time":1586237097,"end_time":1744243199,"id":"1375c9bf-1234-5678-90ab-9c636708c89e","modified_time":1586237097,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"MAN-000000ba","type":"SUB-MAN"},{"created_time":1586237137,"end_time":1744243199,"id":"5974e979-1234-5678-90ab-438f833ec1c9","modified_time":1586237137,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":14,"remaining_quantity":14,"start_time":1586563200,"subscription_id":"ENG-000000bb","type":"SUB-ENG"},{"created_time":1629947267,"end_time":1744243199,"id":"340a9cb3-1234-5678-90ab-b009344dbf3c","modified_time":1629947267,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":10,"remaining_quantity":9,"start_time":1586476800,"subscription_id":"EX24-000000bc","type":"SUB-EX24"},{"created_time":1632941870,"end_time":1744156799,"id":"9b599b0f-1234-5678-90ab-406081b58e7f","modified_time":1632941870,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":10,"remaining_quantity":10,"start_time":1632873600,"subscription_id":"ENG-000000bd","type":"SUB-ENG"},{"created_time":1632941882,"end_time":1744156799,"id":"d6d8ead3-1234-5678-90ab-98badeac7287","modified_time":1632941882,"order_id":"00000000","org_id":"9777c1a0-1234-5678-90ab-02e208b2d34f","quantity":11,"remaining_quantity":9,"start_time":1632873600,"subscription_id":"MAN-000008be","type":"SUB-MAN"}],"summary":{"SUB-ENG":18,"SUB-EX24":3,"SUB-MAN":22,"SUB-VNA":20}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestMoveOrDeleteOrgLicenseToAnotherOrg tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestMoveOrDeleteOrgLicenseToAnotherOrg(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgLicenseAction
    errBody := json.Unmarshal([]byte(`{"notes":"customer notes","op":"annotate","subscription_id":"SUB-000144"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsLicenses.MoveOrDeleteOrgLicenseToAnotherOrg(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsLicensesTestGetOrgLicensesBySite tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicensesBySite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsLicenses.GetOrgLicensesBySite(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"fully_loaded":{"SUB-LOC":30,"SUB-MAN":80},"num_devices":80,"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","usages":{"SUB-LOC":30,"SUB-MAN":60}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsLicensesTestGetOrgLicensesBySite1 tests the behavior of the OrgsLicenses
func TestOrgsLicensesTestGetOrgLicensesBySite1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsLicenses.GetOrgLicensesBySite(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"fully_loaded":{"SUB-LOC":30,"SUB-MAN":80},"num_devices":80,"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","usages":{"SUB-LOC":30,"SUB-MAN":60}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
