package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsLicensesTestClaimMspLicense tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestClaimMspLicense(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CodeString
    errBody := json.Unmarshal([]byte(`{"code":"ZHT3K-H36DT-MG85D-M61AC"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsLicenses.ClaimMspLicense(ctx, mspId, &body)
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

// TestMSPsLicensesTestClaimMspLicense1 tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestClaimMspLicense1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.CodeString
    errBody := json.Unmarshal([]byte(`{"code":"ZHT3K-H36DT-MG85D-M61AC"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsLicenses.ClaimMspLicense(ctx, mspId, &body)
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

// TestMSPsLicensesTestListMspLicenses tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestListMspLicenses(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsLicenses.ListMspLicenses(ctx, mspId)
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

// TestMSPsLicensesTestListMspLicenses1 tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestListMspLicenses1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsLicenses.ListMspLicenses(ctx, mspId)
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

// TestMSPsLicensesTestMoveOrDeleteMspLicenseToAnotherOrg tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestMoveOrDeleteMspLicenseToAnotherOrg(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MspLicenseAction
    errBody := json.Unmarshal([]byte(`{"op":"delete","subscription_id":"SUB-0000144"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := msPsLicenses.MoveOrDeleteMspLicenseToAnotherOrg(ctx, mspId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsLicensesTestListMspOrgLicenses tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestListMspOrgLicenses(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsLicenses.ListMspOrgLicenses(ctx, mspId)
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

// TestMSPsLicensesTestListMspOrgLicenses1 tests the behavior of the MSPsLicenses
func TestMSPsLicensesTestListMspOrgLicenses1(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsLicenses.ListMspOrgLicenses(ctx, mspId)
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
