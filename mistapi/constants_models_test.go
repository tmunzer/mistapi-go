package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestConstantsModelsTestListDeviceModels tests the behavior of the ConstantsModels
func TestConstantsModelsTestListDeviceModels(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsModels.ListDeviceModels(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"ap_type":"jewel","band24":{"band5_channels_op":"low","max_clients":128,"max_power":19,"min_power":8},"band5":{"max_clients":128,"max_power":17,"min_power":8},"band6":{"max_clients":128,"max_power":17,"min_power":8},"band_24_usages":["5"],"ce_dfs_ok":true,"cisco_pace":true,"description":"AP-45","disallowed_channels":{"property1":{"property1":[0],"property2":[0]},"property2":{"property1":[0],"property2":[0]}},"display":"AP45","fcc_dfs_ok":true,"has_11ax":true,"has_compass":false,"has_ext_ant":true,"has_extio":false,"has_height":false,"has_module_port":true,"has_poe_out":true,"has_scanning_radio":true,"has_selectable_radio":true,"has_usb":true,"has_vble":true,"has_wifi_band24":true,"has_wifi_band5":true,"has_wifi_band6":true,"max_poe_out":15400,"max_wlans":0,"model":"AP45","other_dfs_ok":true,"outdoor":false,"radios":{"r0":"6","r1":"5","r2":"24"},"shared_scanning_radio":true,"type":"ap","unmanaged":true,"vble":{"beacon_rate":4,"beams":9,"power":8}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListMxEdgeModels tests the behavior of the ConstantsModels
func TestConstantsModelsTestListMxEdgeModels(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsModels.ListMxEdgeModels(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"X10","model":"ME-X10","ports":{"0":{"display":"xe0","speed":10000},"1":{"display":"xe1","speed":10000},"2":{"display":"xe2","speed":10000},"3":{"display":"xe3","speed":10000}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListSupportedOtherDeviceModels tests the behavior of the ConstantsModels
func TestConstantsModelsTestListSupportedOtherDeviceModels(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsModels.ListSupportedOtherDeviceModels(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"_vendor_model_id":"65","display":"W1850","model":"W1850","type":"router","vendor":"cradlepoint"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
