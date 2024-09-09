package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsServicesTestListOrgServices tests the behavior of the OrgsServices
func TestOrgsServicesTestListOrgServices(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsServices.ListOrgServices(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"addresses":["string"],"apps":["string"],"dscp":8,"hostnames":["string"],"max_jitter":0,"max_latency":0,"max_loss":0,"name":"string","specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"default","type":"custom"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsServicesTestCreateOrgService tests the behavior of the OrgsServices
func TestOrgsServicesTestCreateOrgService(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Service
    errBody := json.Unmarshal([]byte(`{"app_key":"string","name":"string","network_id":"d6797cf4-42b9-4cad-8591-9dd91c3f0fc3","specs":[{"address":"string","port":0,"protocol":"any"}],"subnet":"string","type":"custom"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsServices.CreateOrgService(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"addresses":["string"],"app_categories":["string"],"apps":["string"],"created_time":0,"dscp":0,"failover_policy":"revertable","hostnames":["string"],"id":"497f6eca-6276-5004-bfeb-53cbbbba6f16","max_jitter":0,"max_latency":0,"max_loss":0,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","sle_enabled":false,"specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"data_best_effort","type":"custom","vpn_name":"addresses"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsServicesTestDeleteOrgService tests the behavior of the OrgsServices
func TestOrgsServicesTestDeleteOrgService(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    serviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsServices.DeleteOrgService(ctx, orgId, serviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsServicesTestGetOrgService tests the behavior of the OrgsServices
func TestOrgsServicesTestGetOrgService(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    serviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsServices.GetOrgService(ctx, orgId, serviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"addresses":["string"],"app_categories":["string"],"apps":["string"],"created_time":0,"dscp":0,"failover_policy":"revertable","hostnames":["string"],"id":"497f6eca-6276-5004-bfeb-53cbbbba6f16","max_jitter":0,"max_latency":0,"max_loss":0,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","sle_enabled":false,"specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"data_best_effort","type":"custom","vpn_name":"addresses"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsServicesTestUpdateOrgService tests the behavior of the OrgsServices
func TestOrgsServicesTestUpdateOrgService(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    serviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Service
    errBody := json.Unmarshal([]byte(`{"addresses":["string"],"app_categories":["string"],"apps":["string"],"dscp":0,"failover_policy":"revertable","hostnames":["string"],"max_jitter":0,"max_latency":0,"max_loss":0,"name":"string","sle_enabled":false,"specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"data_best_effort","type":"custom","vpn_name":"addresses"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsServices.UpdateOrgService(ctx, orgId, serviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"addresses":["string"],"app_categories":["string"],"apps":["string"],"created_time":0,"dscp":0,"failover_policy":"revertable","hostnames":["string"],"id":"497f6eca-6276-5004-bfeb-53cbbbba6f16","max_jitter":0,"max_latency":0,"max_loss":0,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","sle_enabled":false,"specs":[{"port_range":"0","protocol":"any"}],"traffic_class":"best_effort","traffic_type":"data_best_effort","type":"custom","vpn_name":"addresses"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
