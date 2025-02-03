package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestConstantsDefinitionsTestListAlarmDefinitions tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListAlarmDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListAlarmDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"Device offline","example":{"aps":["d420b02000fa"],"count":1,"group":"infrastructure","hostnames":["Vendor_AP2"],"id":"e70c308f-7007-4866-9ecd-0d01842979ea","last_seen":1629753888,"org_id":"09dac91f-6e73-4100-89f7-698e0fafbb1b","severity":"warn","site_id":"dcfb31a1-d615-4361-8c95-b9dde05aa704","timestamp":1629753888,"type":"device_down"},"fields":["aps","hostnames"],"group":"infrastructure","key":"device_down","marvis_suggestion_category":"string","severity":"warn"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApChannels tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApChannels(t *testing.T) {
    ctx := context.Background()
    countryCode := "US"
    apiResponse, err := constantsDefinitions.ListApChannels(ctx, &countryCode)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"band24_40mhz_allowed":true,"band24_channels":{"20":[1,2,3,4,5,6,7,8,9,10,11],"40":[1,2,3,4,5,6,7,8,9,10,11]},"band24_enabled":true,"band5_channels":{"20":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165],"40":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"80":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"dfs":[52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144],"outdoor":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165]},"band5_enabled":true,"band6_channels":{"160":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"20":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229,233],"40":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229],"80":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"psc":[5,21,37,53,69,85,101,117,133,149,165,181,197,213,229]},"band6_enabled":true,"certified":true,"code":840,"dfs_ok":true,"key":"US","name":"United States","uses":"US_FCC"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApLedDefinition tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApLedDefinition(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListApLedDefinition(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"code":"01","description":"LED not working","key":"LED_FAILURE","name":"LED Failure"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListAppCategoryDefinitions tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListAppCategoryDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListAppCategoryDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"Images","filters":{"srx":["Enhanced_Images_Media","Enhanced_Web_Images","Enhanced_Image_Servers"]},"key":"Images"},{"display":"Standard","includes":["Adult","FileSharing","Games","Images","Malware","NewsAndReference","Recreation","Religion","Security","Sports","Technology","Violence"],"key":"Standard"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListAppSubCategoryDefinitions tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListAppSubCategoryDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListAppSubCategoryDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"Office Documents","key":"Office_Documents","traffic_type":"data_interactive"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApplications tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApplications(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListApplications(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"app_id":true,"app_image_url":"","app_probe":true,"category":"FileSharing","group":"File Sharing","key":"dropbox","name":"Dropbox","signature_based":true,"ssr_app_id":true}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListCountryCodes tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListCountryCodes(t *testing.T) {
    ctx := context.Background()
    extend := bool(false)
    apiResponse, err := constantsDefinitions.ListCountryCodes(ctx, &extend)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"alpha2":"FR","certified":true,"name":"France","numeric":250}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListGatewayApplications tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListGatewayApplications(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListGatewayApplications(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"app_id":true,"key":"4shared","name":"4shared","ssr_app_id":true}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListInsightMetrics tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListInsightMetrics(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListInsightMetrics(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bytes":{"description":"Aggregated bytes over time","example":[185,197,250],"intervals":{"10m":{"interval":600,"max_age":86400},"1h":{"interval":3600,"max_age":1209600}},"report_durations":{"1d":{"duration":86400,"interval":3600},"1w":{"duration":604800,"interval":3600}},"report_scopes":["site","org"],"scopes":["site","ap","client"],"type":"timeseries","unit":"byte"},"num_clients":{"description":"Number of client over time","example":[18,null,15],"intervals":{"10m":{"interval":600,"max_age":86400},"1h":{"interval":3600,"max_age":1209600}},"report_durations":{"1d":{"duration":86400,"interval":3600},"1w":{"duration":604800,"interval":3600}},"report_scopes":["site","org"],"scopes":["site","ap","device"],"type":"timeseries","unit":""}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListSiteLanguages tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListSiteLanguages(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListSiteLanguages(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"English (US)","display_native":"English (US)","key":"en-US"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestGetLicenseTypes tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestGetLicenseTypes(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.GetLicenseTypes(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"description":"Wired Assurance 12","includes":["sub_ex12a","sub_ex12p"],"key":"sub_ex12","name":"SUB-EX12"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListMarvisClientVersions tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListMarvisClientVersions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListMarvisClientVersions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"label":"default","notes":"","os":"android","url":"https://mobile.mist.com/installers/marvisclient/android/1.1.9/marvisclient-installer.apk","version":"1.1.9"},{"label":"default","notes":"","os":"macos","url":"https://mobile.mist.com/installers/marvisclient/macos/0.100.29/marvisclient-installer.dmg","version":"0.100.29"},{"label":"default","notes":"","os":"windows","url":"https://mobile.mist.com/installers/marvisclient/windows/0.100.26/marvisclient-installer.zip","version":"0.100.26"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListTrafficTypes tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListTrafficTypes(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListTrafficTypes(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"VoIP Video","dscp":32,"failover_policy":"non_revertible","max_jitter":250,"max_latency":1500,"max_loss":35,"name":"voip_video","traffic_class":"medium"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListWebhookTopics tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListWebhookTopics(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsDefinitions.ListWebhookTopics(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"for_org":true,"has_delivery_results":true,"key":"alarms"},{"key":"asset-raw"},{"key":"asset-raw-rssi"},{"for_org":true,"has_delivery_results":true,"key":"audits"},{"for_org":true,"key":"client-info"},{"for_org":true,"key":"client-join"},{"key":"client-latency"},{"for_org":true,"key":"client-sessions"},{"for_org":true,"key":"device-events"},{"for_org":true,"has_delivery_results":true,"key":"device-updowns"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
