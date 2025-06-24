package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsMxClustersTestListOrgMxEdgeClusters tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestListOrgMxEdgeClusters(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsMxClusters.ListOrgMxEdgeClusters(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"for_site":true,"id":"467f6eca-6276-4993-bfeb-53cbbbba6f78","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestListOrgMxEdgeClusters1 tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestListOrgMxEdgeClusters1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsMxClusters.ListOrgMxEdgeClusters(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"for_site":true,"id":"467f6eca-6276-4993-bfeb-53cbbbba6f78","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestCreateOrgMxEdgeCluster tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestCreateOrgMxEdgeCluster(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxcluster
    errBody := json.Unmarshal([]byte(`{"name":"string","radsec":{"enabled":true,"server_name":"string","servers":[{"host":"string","port":0}]},"tunterm_ap_subnets":["string"],"tunterm_hosts":["string"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxClusters.CreateOrgMxEdgeCluster(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestCreateOrgMxEdgeCluster1 tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestCreateOrgMxEdgeCluster1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxcluster
    errBody := json.Unmarshal([]byte(`{"name":"string","radsec":{"enabled":true,"server_name":"string","servers":[{"host":"string","port":0}]},"tunterm_ap_subnets":["string"],"tunterm_hosts":["string"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxClusters.CreateOrgMxEdgeCluster(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestDeleteOrgMxEdgeCluster tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestDeleteOrgMxEdgeCluster(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxclusterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsMxClusters.DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxClustersTestGetOrgMxEdgeCluster tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestGetOrgMxEdgeCluster(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxclusterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMxClusters.GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestGetOrgMxEdgeCluster1 tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestGetOrgMxEdgeCluster1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxclusterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMxClusters.GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestUpdateOrgMxEdgeCluster tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestUpdateOrgMxEdgeCluster(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxclusterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxcluster
    errBody := json.Unmarshal([]byte(`{"name":"string","radsec":{"enabled":true,"server_name":"string","servers":[{"host":"string","port":0}]},"tunterm_ap_subnets":["string"],"tunterm_hosts":["string"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxClusters.UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxClustersTestUpdateOrgMxEdgeCluster1 tests the behavior of the OrgsMxClusters
func TestOrgsMxClustersTestUpdateOrgMxEdgeCluster1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxclusterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxcluster
    errBody := json.Unmarshal([]byte(`{"name":"string","radsec":{"enabled":true,"server_name":"string","servers":[{"host":"string","port":0}]},"tunterm_ap_subnets":["string"],"tunterm_hosts":["string"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxClusters.UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"for_site":true,"id":"468f6eca-6276-4993-bfeb-53cbbbba6f68","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string"}],"auth_servers":[{"host":"string","port":1812,"secret":"string"}],"enabled":true,"server_selection":"ordered"},"radsec_tls":{"keypair":"string"},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","tunterm_ap_subnets":["string"],"tunterm_dhcpd_config":{"enabled":false,"property1":{"enabled":false,"servers":["string"],"type":"relay"},"property2":{"enabled":false,"servers":["string"],"type":"relay"},"servers":["string"],"type":"relay"},"tunterm_extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"tunterm_hosts":["string"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
