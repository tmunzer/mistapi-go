package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

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
