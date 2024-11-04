package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestConstantsEventsTestListClientEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListClientEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListClientEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"11r Association","key":"CLIENT_AUTH_ASSOCIATION_11R"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsEventsTestListDeviceEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListDeviceEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListDeviceEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"description":"AP was assigned to a site","display":"AP Assigned","example":{"ap":"5c5b35000001","audit_id":"e9a88814-fa81-5bdc-34b0-84e8735420e5","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1552408871,"type":"AP_ASSIGNED"},"key":"AP_ASSIGNED"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsEventsTestListMxEdgeEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListMxEdgeEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListMxEdgeEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"description":"Config change on ME was triggered as a result of change made by user","display":"ME Config changed by user","example":{"audit_id":"e9a88814-fa81-5bdc-34b0-84e8735420e5","mxcluster_id":"ed4665ed-c9ad-4835-8ca5-dda642765ad3","mxedge_id":"387804a7-3474-85ce-15a2-f9a9684c9c9","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","service":"mxagent","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1552408871,"type":"ME_CONFIG_CHANGED_BY_USER"},"key":"ME_CONFIG_CHANGED_BY_USER"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsEventsTestListNacEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListNacEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListNacEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"ap":"5c5b355008c0","bssid":"5c5b35548892","cert_cn":"suriyas","cert_expiry":1711557441,"cert_issuer":"/DC=net/DC=jnpr/CN=Juniper Networks Issuing AWS1 CA","cert_san_upn":["suriyas@juniper.net"],"cert_serial":"1300103d29e56ef083797bedc2000100103d29","cert_subject":"/CN=suriyas/emailAddress=suriyas@juniper.net","eap_type":"EAP-TLS","nas_vendor":"Mist","org_id":"94de66e8-556a-4d56-8780-a114620a5c42","random_mac":true,"site_id":"b5a005ab-47d4-41f7-97bf-733f9cc252dd","ssid":"Test_Suriya-SSID","timestamp":1685658478.4389949,"type":"NAC_CLIENT_CERT_CHECK_SUCCESS","username":"suriyas@juniper.net","wcid":"b43637b0-f0d9-0a1d-1ec2-73c394a9f679"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsEventsTestListOtherDeviceEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListOtherDeviceEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListOtherDeviceEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"Connected to NCM","example":{"device_mac":"5c5b351e13b5","mac":"0030447771c0","org_id":"c080ce4d-4e35-4373-bdc4-08df15d257f5","site_id":"1df889ad-9111-4c0e-a00b-8a008b83eb68","text":"Connected to NCM","timestamp":1675827825.765,"type":"CELLULAR_EDGE_CONNECTED_TO_NCM","vendor":"cradlepoint"},"key":"CELLULAR_EDGE_CONNECTED_TO_NCM"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsEventsTestListSystemEventsDefinitions tests the behavior of the ConstantsEvents
func TestConstantsEventsTestListSystemEventsDefinitions(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := constantsEvents.ListSystemEventsDefinitions(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"display":"AP Disconnect","group":"ap_health","key":"ap_disconnected"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
