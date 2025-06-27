package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestUtilitiesUpgradeTestListOrgDeviceUpgrades tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgDeviceUpgrades(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgDeviceUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"466f6eca-6276-4993-bfeb-53cbbbba6f88","site_upgrades":[{"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","upgrade_id":"174bda0-06a3-40ee-b918-d9cbde303690"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListOrgDeviceUpgrades1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgDeviceUpgrades1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgDeviceUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"466f6eca-6276-4993-bfeb-53cbbbba6f88","site_upgrades":[{"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","upgrade_id":"174bda0-06a3-40ee-b918-d9cbde303690"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeOrgDevices tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgDevices(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.UpgradeOrgDevices(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"enable_p2p":true,"force":true,"id":"31223085-405d-4b64-8aea-9c5b98098b4b","strategy":"big_bang","target_version":"0.14.29411","upgrades":[{"site_id":"1bbe6e79-2583-403c-be1a-9881b4691ab6","upgrade":{"id":"473f6eca-6276-4993-bfeb-53cbbbba6f18","start_time":1717658765,"status":"upgrading","targets":{"download_requested":["5c5b3550bd2e"],"downloaded":["003e7316ff9e"],"total":4,"upgraded":["5c5b3550bd2e","003e7316ff9e"]}}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeOrgDevices1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgDevices1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.UpgradeOrgDevices(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"enable_p2p":true,"force":true,"id":"31223085-405d-4b64-8aea-9c5b98098b4b","strategy":"big_bang","target_version":"0.14.29411","upgrades":[{"site_id":"1bbe6e79-2583-403c-be1a-9881b4691ab6","upgrade":{"id":"473f6eca-6276-4993-bfeb-53cbbbba6f18","start_time":1717658765,"status":"upgrading","targets":{"download_requested":["5c5b3550bd2e"],"downloaded":["003e7316ff9e"],"total":4,"upgraded":["5c5b3550bd2e","003e7316ff9e"]}}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetOrgDeviceUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgDeviceUpgrade(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgDeviceUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"enable_p2p":true,"force":true,"id":"31223085-405d-4b64-8aea-9c5b98098b4b","strategy":"big_bang","target_version":"0.14.29411","upgrades":[{"site_id":"1bbe6e79-2583-403c-be1a-9881b4691ab6","upgrade":{"id":"473f6eca-6276-4993-bfeb-53cbbbba6f18","start_time":1717658765,"status":"upgrading","targets":{"download_requested":["5c5b3550bd2e"],"downloaded":["003e7316ff9e"],"total":4,"upgraded":["5c5b3550bd2e","003e7316ff9e"]}}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetOrgDeviceUpgrade1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgDeviceUpgrade1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgDeviceUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"enable_p2p":true,"force":true,"id":"31223085-405d-4b64-8aea-9c5b98098b4b","strategy":"big_bang","target_version":"0.14.29411","upgrades":[{"site_id":"1bbe6e79-2583-403c-be1a-9881b4691ab6","upgrade":{"id":"473f6eca-6276-4993-bfeb-53cbbbba6f18","start_time":1717658765,"status":"upgrading","targets":{"download_requested":["5c5b3550bd2e"],"downloaded":["003e7316ff9e"],"total":4,"upgraded":["5c5b3550bd2e","003e7316ff9e"]}}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestCancelOrgDeviceUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestCancelOrgDeviceUpgrade(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesUpgrade.CancelOrgDeviceUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesUpgradeTestListOrgAvailableDeviceVersions tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgAvailableDeviceVersions(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    model := "AP43"
    apiResponse, err := utilitiesUpgrade.ListOrgAvailableDeviceVersions(ctx, orgId, &mType, &model)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"model":"AP41","tag":"stable","version":"v0.1.543"},{"model":"AP21","version":"v0.1.545"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListOrgAvailableDeviceVersions1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgAvailableDeviceVersions1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    model := "AP43"
    apiResponse, err := utilitiesUpgrade.ListOrgAvailableDeviceVersions(ctx, orgId, &mType, &model)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"model":"AP41","tag":"stable","version":"v0.1.543"},{"model":"AP21","version":"v0.1.545"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeOrgJsiDevice tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgJsiDevice(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    var body models.VersionString
    errBody := json.Unmarshal([]byte(`{"version":"3.1.5"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesUpgrade.UpgradeOrgJsiDevice(ctx, orgId, deviceMac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesUpgradeTestListOrgMxEdgeUpgrades tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgMxEdgeUpgrades(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgMxEdgeUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestUtilitiesUpgradeTestListOrgMxEdgeUpgrades1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgMxEdgeUpgrades1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgMxEdgeUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestUtilitiesUpgradeTestUpgradeOrgMxEdges tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgMxEdges(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MxedgeUpgradeMulti
    errBody := json.Unmarshal([]byte(`{"allow_downgrades":{"mxagent":false,"tunterm":true},"canary_phases":[5,25,50,100],"channel":"stable","max_failure_percentage":5,"max_failures":[1,1,5,5],"mxedge_ids":["387804a7-3474-85ce-15a2-f9a9684c9c90"],"start_time":1624399840,"strategy":"serial","versions":{"mxagent":"latest","radsecproxy":"<version>","tunterm":"default"}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesUpgrade.UpgradeOrgMxEdges(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesUpgradeTestGetOrgMxEdgeUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgMxEdgeUpgrade(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgMxEdgeUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestUtilitiesUpgradeTestGetOrgMxEdgeUpgrade1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgMxEdgeUpgrade1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgMxEdgeUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestUtilitiesUpgradeTestListOrgSsrUpgrades tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgSsrUpgrades(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgSsrUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListOrgSsrUpgrades1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgSsrUpgrades1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.ListOrgSsrUpgrades(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeOrgSsrs tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgSsrs(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsrUpgradeMulti
    errBody := json.Unmarshal([]byte(`{"channel":"stable","device_ids":["00000000-0000-0000-1000-5c5b3500001f","00000000-0000-0000-1000-5c5b35000020"],"strategy":"big_bang","version":"5.3.0-93"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeOrgSsrs(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeOrgSsrs1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeOrgSsrs1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsrUpgradeMulti
    errBody := json.Unmarshal([]byte(`{"channel":"stable","device_ids":["00000000-0000-0000-1000-5c5b3500001f","00000000-0000-0000-1000-5c5b35000020"],"strategy":"big_bang","version":"5.3.0-93"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeOrgSsrs(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetOrgSsrUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgSsrUpgrade(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgSsrUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","device_type":"gateway","id":"5cbcee0a-c620-4bb4-a25e-15000934e9d8","status":"upgrading","targets":{"failed":[],"queued":[],"success":[],"upgrading":["8e525f1d-4178-4ae1-a988-2b0176855e55"]},"versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetOrgSsrUpgrade1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetOrgSsrUpgrade1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetOrgSsrUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","device_type":"gateway","id":"5cbcee0a-c620-4bb4-a25e-15000934e9d8","status":"upgrading","targets":{"failed":[],"queued":[],"success":[],"upgrading":["8e525f1d-4178-4ae1-a988-2b0176855e55"]},"versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestCancelOrgSsrUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestCancelOrgSsrUpgrade(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesUpgrade.CancelOrgSsrUpgrade(ctx, orgId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesUpgradeTestListOrgAvailableSsrVersions tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgAvailableSsrVersions(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    channel := models.SsrVersionChannelEnum("stable")
    
    apiResponse, err := utilitiesUpgrade.ListOrgAvailableSsrVersions(ctx, orgId, &channel, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"default":true,"package":"SSR","version":"5.3.1-17"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListOrgAvailableSsrVersions1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListOrgAvailableSsrVersions1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    channel := models.SsrVersionChannelEnum("stable")
    
    apiResponse, err := utilitiesUpgrade.ListOrgAvailableSsrVersions(ctx, orgId, &channel, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"default":true,"package":"SSR","version":"5.3.1-17"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListSiteDeviceUpgrades tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListSiteDeviceUpgrades(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.ListSiteDeviceUpgrades(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"counts":{"download_requested":0,"downloaded":0,"failed":0,"reboot_in_progress":0,"rebooted":0,"skipped":0,"total":0},"enable_p2p":true,"force":true,"id":"472f6eca-6276-4993-bfeb-53cbbbba6f28","start_time":0,"status":"created","strategy":"big_bang","target_version":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListSiteDeviceUpgrades1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListSiteDeviceUpgrades1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.ListSiteDeviceUpgrades(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"counts":{"download_requested":0,"downloaded":0,"failed":0,"reboot_in_progress":0,"rebooted":0,"skipped":0,"total":0},"enable_p2p":true,"force":true,"id":"472f6eca-6276-4993-bfeb-53cbbbba6f28","start_time":0,"status":"created","strategy":"big_bang","target_version":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeSiteDevices tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeSiteDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.UpgradeSiteDevices(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"upgrade_id":"4316c116-0acb-4c43-8f06-6723154e741e"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeSiteDevices1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeSiteDevices1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := utilitiesUpgrade.UpgradeSiteDevices(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"upgrade_id":"4316c116-0acb-4c43-8f06-6723154e741e"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetSiteDeviceUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetSiteDeviceUpgrade(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetSiteDeviceUpgrade(ctx, siteId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"canary_phases":[1,10,50,100],"current_phase":0,"enable_p2p":true,"force":false,"id":"b910de88-0e0e-43a0-b20e-a43aa8788917","max_failure_percentage":5,"max_failures":[5,5,5,5],"p2p_cluster_size":10,"p2p_parallelism":2,"reboot_at":1738233315,"start_time":1738232698,"status":"created","strategy":"canary","target_version":"0.14.29633","targets":{"downloaded":["30783a619182"],"total":2,"upgraded":["a8378280da5e"]},"upgrade_plan":{"1":["30783a619182"],"2":["a8378280da5e"]}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetSiteDeviceUpgrade1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetSiteDeviceUpgrade1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetSiteDeviceUpgrade(ctx, siteId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"canary_phases":[1,10,50,100],"current_phase":0,"enable_p2p":true,"force":false,"id":"b910de88-0e0e-43a0-b20e-a43aa8788917","max_failure_percentage":5,"max_failures":[5,5,5,5],"p2p_cluster_size":10,"p2p_parallelism":2,"reboot_at":1738233315,"start_time":1738232698,"status":"created","strategy":"canary","target_version":"0.14.29633","targets":{"downloaded":["30783a619182"],"total":2,"upgraded":["a8378280da5e"]},"upgrade_plan":{"1":["30783a619182"],"2":["a8378280da5e"]}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestCancelSiteDeviceUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestCancelSiteDeviceUpgrade(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesUpgrade.CancelSiteDeviceUpgrade(ctx, siteId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesUpgradeTestListSiteAvailableDeviceVersions tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListSiteAvailableDeviceVersions(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    
    apiResponse, err := utilitiesUpgrade.ListSiteAvailableDeviceVersions(ctx, siteId, &mType, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"model":"AP41","tag":"stable","version":"v0.1.543"},{"model":"AP21","version":"v0.1.545"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestListSiteAvailableDeviceVersions1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestListSiteAvailableDeviceVersions1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    
    apiResponse, err := utilitiesUpgrade.ListSiteAvailableDeviceVersions(ctx, siteId, &mType, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"model":"AP41","tag":"stable","version":"v0.1.543"},{"model":"AP21","version":"v0.1.545"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeDevice tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeDevice(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.DeviceUpgrade
    errBody := json.Unmarshal([]byte(`{"version":"3.1.5"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeDevice(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"status":"inprogress","timestamp":1428949501}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeDevice1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeDevice1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.DeviceUpgrade
    errBody := json.Unmarshal([]byte(`{"version":"3.1.5"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeDevice(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"status":"inprogress","timestamp":1428949501}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetSiteSsrUpgrade tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetSiteSsrUpgrade(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetSiteSsrUpgrade(ctx, siteId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","device_type":"gateway","id":"5cbcee0a-c620-4bb4-a25e-15000934e9d8","status":"upgrading","targets":{"failed":[],"queued":[],"success":[],"upgrading":["8e525f1d-4178-4ae1-a988-2b0176855e55"]},"versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestGetSiteSsrUpgrade1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestGetSiteSsrUpgrade1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    upgradeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := utilitiesUpgrade.GetSiteSsrUpgrade(ctx, siteId, upgradeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","device_type":"gateway","id":"5cbcee0a-c620-4bb4-a25e-15000934e9d8","status":"upgrading","targets":{"failed":[],"queued":[],"success":[],"upgrading":["8e525f1d-4178-4ae1-a988-2b0176855e55"]},"versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeSsr tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeSsr(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsrUpgrade
    errBody := json.Unmarshal([]byte(`{"channel":"stable","version":"5.3.1-170-93"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeSsr(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestUtilitiesUpgradeTestUpgradeSsr1 tests the behavior of the UtilitiesUpgrade
func TestUtilitiesUpgradeTestUpgradeSsr1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsrUpgrade
    errBody := json.Unmarshal([]byte(`{"channel":"stable","version":"5.3.1-170-93"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := utilitiesUpgrade.UpgradeSsr(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"channel":"stable","counts":{"failed":0,"queued":1,"success":0,"upgrading":1},"device_type":"gateway","id":"ceef2c8a-e2e6-447a-8b27-cb4f3ec1adae","status":"upgrading","strategy":"serial","versions":{}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
