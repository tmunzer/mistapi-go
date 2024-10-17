package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsMxEdgesTestListOrgMxEdges tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestListOrgMxEdges(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	forSites := models.MxedgeForSiteEnum("any")
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsMxEdges.ListOrgMxEdges(ctx, orgId, &forSites, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"cpu_stat":{"cpus":{"cpu0":{"idle":79,"interrupt":0,"system":4,"usage":20,"user":16},"cpu1":{"idle":93,"interrupt":0,"system":4,"usage":6,"user":1}},"idle":87,"interrupt":0,"system":5,"usage":12,"user":7},"ext_ip":"116.187.144.16","id":"387804a7-3474-85ce-15a2-f9a9684c9c90","ip_stat":{"ip":"172.16.5.3","ips":{"ens192":"172.16.5.3/24,fe81::20c:29ff:fef8:d18e/64"}},"lag_stat":{"lag0":{"active_ports":["0","1"]}},"last_seen":1547437078,"magic":"ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","memory_stats":{"active":1061085184,"available":4124860416,"buffers":789495808,"cached":718016512,"free":2818838528,"inactive":458158080,"swap_cached":0,"swap_free":8161062912,"swap_total":8161062912,"total":7947616256,"usage":65},"model":"ME-S2019","mxagent_registered":false,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","name":"Guest","num_tunnels":31,"port_stat":{"eth0":{"full_duplex":true,"lldp_stats":{"mgmt_addr":"122.16.3.11","port_desc":"GigabitEthernet4/0/16","port_id":"\u0005Gi4/0/16","system_desc":"Cisco IOS Software","system_name":"ME-DC2-DIS-SW"},"rx_bytes":2056,"rx_errors":0,"rx_pkts":670,"speed":1000,"tx_bytes":2056,"tx_pkts":670,"up":true},"eth1":{"up":false},"module":{"up":false}},"status":"connected","tunterm_registered":false,"tunterm_stat":{"monitoring_failed":false},"uptime":884221,"version":"0.1.2","virtualization_type":"VirtualizationVMware"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestCreateOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestCreateOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Mxedge
	errBody := json.Unmarshal([]byte(`{"model":"ME-100","mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{},"name":"Guest","ntp_servers":[],"oob_ip_config":{},"services":["tunterm"],"tunterm_ip_config":{"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"},"tunterm_port_config":{"downstream_ports":["0","1","2","3"],"separate_upstream_downstream":true,"upstream_port_vlan_id":1,"upstream_ports":["0","1","2","3"]},"tunterm_switch_config":{"0":{"port_vlan_id":1,"vlan_ids":[5,2,3]},"enabled":true}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsMxEdges.CreateOrgMxEdge(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","model":"ME-100","mxagent_registered":true,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"name":"Guest","ntp_servers":[],"oob_ip_config":{"dns":["8.8.8.8","4.4.4.4"],"gateway":"10.2.1.254","ip":"10.2.1.10","netmask":"255.255.255.0","type":"static"},"tunterm_dhcpd_config":{"2":{"enabled":true,"servers":["11.2.3.44"]},"enabled":false,"servers":["11.2.3.4"]},"tunterm_extra_routes":{"11.0.0.0/8":{"via":"10.3.3.1"}},"tunterm_ip_config":{"dns":["8.8.8.8"],"dns_suffix":[".mist.local"],"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestAssignOrgMxEdgeToSite tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestAssignOrgMxEdgeToSite(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MxedgesAssign
	errBody := json.Unmarshal([]byte(`{"mxedge_ids":["387804a7-3474-85ce-15a2-f9a9684c9c90"],"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsMxEdges.AssignOrgMxEdgeToSite(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestClaimOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestClaimOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.CodeString
	errBody := json.Unmarshal([]byte(`{"code":"135-546-673"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsMxEdges.ClaimOrgMxEdge(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestCountOrgMxEdges tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestCountOrgMxEdges(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgMxedgeCountDistinctEnum("model")

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsMxEdges.CountOrgMxEdges(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestCountOrgSiteMxEdgeEvents tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestCountOrgSiteMxEdgeEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgMxedgeEventsCountDistinctEnum("mxedge_id")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsMxEdges.CountOrgSiteMxEdgeEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestSearchOrgMistEdgeEvents tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestSearchOrgMistEdgeEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsMxEdges.SearchOrgMistEdgeEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1694708579,"limit":10,"page":3,"results":[{"mxcluster_id":"2815c917-58e7-472f-a190-bfd44fb58d05","mxedge_id":"00000000-0000-0000-1000-020000dc585c","org_id":"f2695c32-0e83-4936-b1b2-96fc88051213","service":"tunterm","timestamp":1694678225.927,"type":"ME_SERVICE_STOPPED"}],"start":1694622179}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestSearchOrgMxEdges tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestSearchOrgMxEdges(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsMxEdges.SearchOrgMxEdges(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1694708579,"limit":10,"results":[{"distro":"buster","last_seen":1695151551.833,"model":"ME-X5","mxedge_id":"00000000-0000-0000-1000-d420b0f0025d","org_id":"35d96b1a-1a13-4ba8-90f5-1e78dd2a10c5","tunterm_version":"0.1.2813","uptime":5662632}],"start":1694622179,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestUnassignOrgMxEdgeFromSite tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestUnassignOrgMxEdgeFromSite(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MxedgesUnassign
	errBody := json.Unmarshal([]byte(`{"mxedge_ids":["387804a7-3474-85ce-15a2-f9a9684c9c90"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsMxEdges.UnassignOrgMxEdgeFromSite(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestGetOrgMxEdgeUpgradeInfo tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestGetOrgMxEdgeUpgradeInfo(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	channel := models.GetOrgMxedgeUpgradeInfoChannelEnum("stable")
	apiResponse, err := orgsMxEdges.GetOrgMxEdgeUpgradeInfo(ctx, orgId, &channel)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"default":true,"package":"mxagent","version":"2.4.100"},{"package":"tunterm","version":"1.0.0"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestDeleteOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestDeleteOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsMxEdges.DeleteOrgMxEdge(ctx, orgId, mxedgeId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxEdgesTestGetOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestGetOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsMxEdges.GetOrgMxEdge(ctx, orgId, mxedgeId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","model":"ME-100","mxagent_registered":true,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"name":"Guest","ntp_servers":[],"oob_ip_config":{"dns":["8.8.8.8","4.4.4.4"],"gateway":"10.2.1.254","ip":"10.2.1.10","netmask":"255.255.255.0","type":"static"},"tunterm_dhcpd_config":{"2":{"enabled":true,"servers":["11.2.3.44"]},"enabled":false,"servers":["11.2.3.4"]},"tunterm_extra_routes":{"11.0.0.0/8":{"via":"10.3.3.1"}},"tunterm_ip_config":{"dns":["8.8.8.8"],"dns_suffix":[".mist.local"],"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestUpdateOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestUpdateOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Mxedge
	errBody := json.Unmarshal([]byte(`{"model":"ME-X1","name":"string","ntp_servers":["string"],"services":["tunterm"],"tunterm_ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"string","ip":"string","netmask":"string"},"tunterm_port_config":{"downstream_ports":["string"],"separate_upstream_downstream":true,"upstream_port_vlan_id":1,"upstream_ports":["string"]}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsMxEdges.UpdateOrgMxEdge(ctx, orgId, mxedgeId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","model":"ME-100","mxagent_registered":true,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"name":"Guest","ntp_servers":[],"oob_ip_config":{"dns":["8.8.8.8","4.4.4.4"],"gateway":"10.2.1.254","ip":"10.2.1.10","netmask":"255.255.255.0","type":"static"},"tunterm_dhcpd_config":{"2":{"enabled":true,"servers":["11.2.3.44"]},"enabled":false,"servers":["11.2.3.4"]},"tunterm_extra_routes":{"11.0.0.0/8":{"via":"10.3.3.1"}},"tunterm_ip_config":{"dns":["8.8.8.8"],"dns_suffix":[".mist.local"],"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxEdgesTestRestartOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestRestartOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsMxEdges.RestartOrgMxEdge(ctx, orgId, mxedgeId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxEdgesTestBounceOrgMxEdgeDataPorts tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestBounceOrgMxEdgeDataPorts(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UtilsTuntermBouncePort
	errBody := json.Unmarshal([]byte(`{"ports":["0","2"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := orgsMxEdges.BounceOrgMxEdgeDataPorts(ctx, orgId, mxedgeId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxEdgesTestUploadOrgMxEdgeSupportFiles tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestUploadOrgMxEdgeSupportFiles(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsMxEdges.UploadOrgMxEdgeSupportFiles(ctx, orgId, mxedgeId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxEdgesTestUnregisterOrgMxEdge tests the behavior of the OrgsMxEdges
func TestOrgsMxEdgesTestUnregisterOrgMxEdge(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsMxEdges.UnregisterOrgMxEdge(ctx, orgId, mxedgeId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
