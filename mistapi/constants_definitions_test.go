// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"testing"
)

// TestConstantsDefinitionsTestListApChannels tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApChannels(t *testing.T) {
	ctx := context.Background()
	countryCode := "US"
	apiResponse, err := constantsDefinitions.ListApChannels(ctx, &countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band24_40mhz_allowed":true,"band24_channels":{"20":[1,2,3,4,5,6,7,8,9,10,11],"40":[1,2,3,4,5,6,7,8,9,10,11]},"band24_enabled":true,"band5_channels":{"20":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165],"40":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"80":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"dfs":[52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144],"outdoor":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165]},"band5_enabled":true,"band6_channels":{"160":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"20":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229,233],"40":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229],"80":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"psc":[5,21,37,53,69,85,101,117,133,149,165,181,197,213,229]},"band6_enabled":true,"certified":true,"code":840,"dfs_ok":true,"key":"US","name":"United States","uses":"US_FCC"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApChannels1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApChannels1(t *testing.T) {
	ctx := context.Background()
	countryCode := "US"
	apiResponse, err := constantsDefinitions.ListApChannels(ctx, &countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"band24_40mhz_allowed":true,"band24_channels":{"20":[1,2,3,4,5,6,7,8,9,10,11],"40":[1,2,3,4,5,6,7,8,9,10,11]},"band24_enabled":true,"band5_channels":{"20":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165],"40":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"80":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161],"dfs":[52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144],"outdoor":[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165]},"band5_enabled":true,"band6_channels":{"160":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"20":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229,233],"40":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229],"80":[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221],"psc":[5,21,37,53,69,85,101,117,133,149,165,181,197,213,229]},"band6_enabled":true,"certified":true,"code":840,"dfs_ok":true,"key":"US","name":"United States","uses":"US_FCC"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApLEslVersions tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApLEslVersions(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListApLEslVersions(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"esl_version":"2.5.1","model":"AP34"},{"esl_version":"2.5.0","model":"AP43"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApLEslVersions1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApLEslVersions1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListApLEslVersions(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"esl_version":"2.5.1","model":"AP34"},{"esl_version":"2.5.0","model":"AP43"}]`
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"code":"01","description":"LED not working","key":"LED_FAILURE","name":"LED Failure"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApLedDefinition1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApLedDefinition1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListApLedDefinition(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"Images","filters":{"srx":["Enhanced_Images_Media","Enhanced_Web_Images","Enhanced_Image_Servers"]},"key":"Images"},{"display":"Standard","includes":["Adult","FileSharing","Games","Images","Malware","NewsAndReference","Recreation","Religion","Security","Sports","Technology","Violence"],"key":"Standard"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListAppCategoryDefinitions1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListAppCategoryDefinitions1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListAppCategoryDefinitions(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"Office Documents","key":"Office_Documents","traffic_type":"data_interactive"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListAppSubCategoryDefinitions1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListAppSubCategoryDefinitions1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListAppSubCategoryDefinitions(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"app_id":true,"app_image_url":"","app_probe":true,"category":"FileSharing","group":"File Sharing","key":"dropbox","name":"Dropbox","signature_based":true,"ssr_app_id":true}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListApplications1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListApplications1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListApplications(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"alpha2":"FR","certified":true,"name":"France","numeric":250}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListCountryCodes1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListCountryCodes1(t *testing.T) {
	ctx := context.Background()
	extend := bool(false)
	apiResponse, err := constantsDefinitions.ListCountryCodes(ctx, &extend)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"alpha2":"FR","certified":true,"name":"France","numeric":250}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListFingerprintTypes tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListFingerprintTypes(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListFingerprintTypes(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"family":["2Wire Router","3Com Switches","ACTi Corporation Network Camera","APC Video Equipment","APC-Schneider UPS","Aastra VoIP","Acer","Actiontec Wireless Router","Aerohive Access Point","Alcatel","Alcatel VoIP","Amazon Echo"],"mfg":["100fio Networks Technology llc","10NET COMMUNICATIONS/DCA","11wave Technonlogy Co.,Ltd","12Sided Technology, LLC","1Net Corporation","1Verge Internet Technology (Beijing) Co., Ltd."],"model":["10T Lite","10T Pro","10th Gen","11 Lite","11 Pro","11 Pro Max"],"os":["Android","Apple OS","Asha Platform OS"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListFingerprintTypes1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListFingerprintTypes1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListFingerprintTypes(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"family":["2Wire Router","3Com Switches","ACTi Corporation Network Camera","APC Video Equipment","APC-Schneider UPS","Aastra VoIP","Acer","Actiontec Wireless Router","Aerohive Access Point","Alcatel","Alcatel VoIP","Amazon Echo"],"mfg":["100fio Networks Technology llc","10NET COMMUNICATIONS/DCA","11wave Technonlogy Co.,Ltd","12Sided Technology, LLC","1Net Corporation","1Verge Internet Technology (Beijing) Co., Ltd."],"model":["10T Lite","10T Pro","10th Gen","11 Lite","11 Pro","11 Pro Max"],"os":["Android","Apple OS","Asha Platform OS"]}`
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"app_id":true,"key":"4shared","name":"4shared","ssr_app_id":true}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListGatewayApplications1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListGatewayApplications1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListGatewayApplications(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"app_id":true,"key":"4shared","name":"4shared","ssr_app_id":true}]`
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"English (US)","display_native":"English (US)","key":"en-US"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListSiteLanguages1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListSiteLanguages1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListSiteLanguages(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"English (US)","display_native":"English (US)","key":"en-US"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListLicenseTypes tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListLicenseTypes(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListLicenseTypes(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"description":"Wired Assurance 12","includes":["sub_ex12a","sub_ex12p"],"key":"sub_ex12","name":"SUB-EX12"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListLicenseTypes1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListLicenseTypes1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListLicenseTypes(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"label":"default","notes":"","os":"android","url":"https://mobile.mist.com/installers/marvisclient/android/1.1.9/marvisclient-installer.apk","version":"1.1.9"},{"label":"default","notes":"","os":"macos","url":"https://mobile.mist.com/installers/marvisclient/macos/0.100.29/marvisclient-installer.dmg","version":"0.100.29"},{"label":"default","notes":"","os":"windows","url":"https://mobile.mist.com/installers/marvisclient/windows/0.100.26/marvisclient-installer.zip","version":"0.100.26"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListMarvisClientVersions1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListMarvisClientVersions1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListMarvisClientVersions(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"label":"default","notes":"","os":"android","url":"https://mobile.mist.com/installers/marvisclient/android/1.1.9/marvisclient-installer.apk","version":"1.1.9"},{"label":"default","notes":"","os":"macos","url":"https://mobile.mist.com/installers/marvisclient/macos/0.100.29/marvisclient-installer.dmg","version":"0.100.29"},{"label":"default","notes":"","os":"windows","url":"https://mobile.mist.com/installers/marvisclient/windows/0.100.26/marvisclient-installer.zip","version":"0.100.26"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListStates tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListStates(t *testing.T) {
	ctx := context.Background()
	countryCode := "US"
	apiResponse, err := constantsDefinitions.ListStates(ctx, countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"iso_code":"AK","name":"Alaska"},{"iso_code":"AL","name":"Alabama"},{"iso_code":"AS","name":"American Samoa"},{"iso_code":"AZ","name":"Arizona"},{"iso_code":"CA","name":"California"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListStates1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListStates1(t *testing.T) {
	ctx := context.Background()
	countryCode := "US"
	apiResponse, err := constantsDefinitions.ListStates(ctx, countryCode)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"iso_code":"AK","name":"Alaska"},{"iso_code":"AL","name":"Alabama"},{"iso_code":"AS","name":"American Samoa"},{"iso_code":"AZ","name":"Arizona"},{"iso_code":"CA","name":"California"}]`
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"VoIP Video","dscp":32,"failover_policy":"non_revertible","max_jitter":250,"max_latency":1500,"max_loss":35,"name":"voip_video","traffic_class":"medium"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListTrafficTypes1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListTrafficTypes1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListTrafficTypes(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"for_org":true,"has_delivery_results":true,"key":"alarms"},{"key":"asset-raw-rssi"},{"for_org":true,"has_delivery_results":true,"key":"audits"},{"for_org":true,"key":"client-info"},{"for_org":true,"key":"client-join"},{"key":"client-latency"},{"for_org":true,"key":"client-sessions"},{"allows_single_event_per_message":true,"for_org":true,"key":"device-events"},{"for_org":true,"has_delivery_results":true,"key":"device-updowns"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsDefinitionsTestListWebhookTopics1 tests the behavior of the ConstantsDefinitions
func TestConstantsDefinitionsTestListWebhookTopics1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsDefinitions.ListWebhookTopics(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"for_org":true,"has_delivery_results":true,"key":"alarms"},{"key":"asset-raw-rssi"},{"for_org":true,"has_delivery_results":true,"key":"audits"},{"for_org":true,"key":"client-info"},{"for_org":true,"key":"client-join"},{"key":"client-latency"},{"for_org":true,"key":"client-sessions"},{"allows_single_event_per_message":true,"for_org":true,"key":"device-events"},{"for_org":true,"has_delivery_results":true,"key":"device-updowns"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
