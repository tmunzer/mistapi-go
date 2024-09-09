package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsGatewayTemplatesTestListOrgGatewayTemplates tests the behavior of the OrgsGatewayTemplates
func TestOrgsGatewayTemplatesTestListOrgGatewayTemplates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsGatewayTemplates.ListOrgGatewayTemplates(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsGatewayTemplatesTestCreateOrgGatewayTemplate tests the behavior of the OrgsGatewayTemplates
func TestOrgsGatewayTemplatesTestCreateOrgGatewayTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.GatewayTemplate
    errBody := json.Unmarshal([]byte(`{"dhcpd_config":{"Corp-Mgmt":{"dns_servers":["8.8.8.8"],"dns_suffix":["stag.one"],"gateway":"10.3.172.9","ip_end":"10.3.172.99","ip_start":"10.3.172.50","type":"local"},"Corp-lan":{"dns_servers":["8.8.8.8"],"dns_suffix":["stag.one"],"gateway":"10.3.171.9","ip_end":"10.3.171.99","ip_start":"10.3.171.50","type":"local"}},"dnsOverride":true,"dns_servers":["10.3.20.201","10.3.51.222","1.1.1.1"],"dns_suffix":["example.com"],"extra_routes":{"10.101.0.0/16":{"via":"10.3.100.10"}},"ip_configs":{"Corp-Core":{"ip":"10.3.100.9","netmask":"/24","type":"static"},"Corp-Mgmt":{"ip":"10.3.172.9","netmask":"/24","type":"static"},"Corp-lan":{"ip":"10.3.171.9","netmask":"/24","type":"static"}},"name":"ITParis","ntpOverride":true,"ntp_servers":["10.3.51.222"],"path_preferences":{"core":{"paths":[{"networks":["Corp-Core"],"type":"local"}],"strategy":"ordered"},"lab":{"paths":[{"networks":["Corp-lan"],"type":"local"}],"strategy":"ordered"},"mgmt":{"paths":[{"networks":["Corp-Mgmt"],"type":"local"}],"strategy":"ordered"},"untrust":{"paths":[{"name":"wan","type":"wan"}],"strategy":"ordered"}},"port_config":{"ge-0/0/0":{"aggregated":false,"ip_config":{"gateway":"192.168.1.1","ip":"192.168.1.9","netmask":"/24","type":"static"},"name":"wan","redundant":false,"traffic_shaping":{"enabled":false},"usage":"wan","wan_type":"broadband"},"ge-0/0/6-7":{"ae_disable_lacp":false,"ae_idx":"0","ae_lacp_force_up":true,"aggregated":true,"networks":["Corp-lan","Corp-Mgmt","Corp-Core"],"usage":"lan"}},"service_policies":[{"action":"allow","idp":{"enabled":false},"name":"ITParis-Internal","path_preference":"core","services":["internal_dns","drive"],"tenants":["ITParis"]},{"action":"allow","idp":{"enabled":false},"name":"ITParis-internet","path_preference":"untrust","services":["internet_any"],"tenants":["ITParis"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-core","path_preference":"core","services":["internal_dns","internal_ntp"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-mxe-tt-in","path_preference":"mxe-in","services":["internal_any"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-untrust","path_preference":"untrust","services":["mxedge-updates","radsec","icmp","internet_dns","internet_ntp"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mxe-data-0-to-untrust","path_preference":"untrust","services":["internet_any"],"tenants":["ITParis"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-mgt","path_preference":"mgmt","services":["mgmt"],"tenants":["domain.Corp-Core","lan.Corp-Core","servers.Corp-Core","Corp-Core"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-edge-in","path_preference":"mxe-in","services":["internal_any"],"tenants":["lan.Corp-Core","Corp-Core"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-iot","path_preference":"iot","services":["iot"],"tenants":["lan.Corp-Core","servers-hassio.Corp-Core","servers-kubes.Corp-Core"]},{"action":"allow","idp":{"enabled":false},"name":"tanker-to-cctv","path_preference":"iot","services":["rtsp"],"tenants":["servers-tanker.Corp-Core"]},{"action":"allow","idp":{"enabled":false},"name":"core-to-untrust","path_preference":"untrust","services":["internet_any"],"tenants":["lan.Corp-Core","domain.Corp-Core","servers.Corp-Core"]},{"action":"deny","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"iot-upgrade-cctv","path_preference":"untrust","services":["motioneye","nodejs","raspbian"],"tenants":["printer"]}],"type":"standalone"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsGatewayTemplates.CreateOrgGatewayTemplate(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsGatewayTemplatesTestDeleteOrgGatewayTemplate tests the behavior of the OrgsGatewayTemplates
func TestOrgsGatewayTemplatesTestDeleteOrgGatewayTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    gatewaytemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsGatewayTemplates.DeleteOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsGatewayTemplatesTestGetOrgGatewayTemplate tests the behavior of the OrgsGatewayTemplates
func TestOrgsGatewayTemplatesTestGetOrgGatewayTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    gatewaytemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsGatewayTemplates.GetOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsGatewayTemplatesTestUpdateOrgGatewayTemplate tests the behavior of the OrgsGatewayTemplates
func TestOrgsGatewayTemplatesTestUpdateOrgGatewayTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    gatewaytemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.GatewayTemplate
    errBody := json.Unmarshal([]byte(`{"dhcpd_config":{"Corp-Mgmt":{"dns_servers":["8.8.8.8"],"dns_suffix":["stag.one"],"gateway":"10.3.172.9","ip_end":"10.3.172.99","ip_start":"10.3.172.50","type":"local"},"Corp-lan":{"dns_servers":["8.8.8.8"],"dns_suffix":["stag.one"],"gateway":"10.3.171.9","ip_end":"10.3.171.99","ip_start":"10.3.171.50","type":"local"}},"dnsOverride":true,"dns_servers":["10.3.20.201","10.3.51.222","1.1.1.1"],"dns_suffix":["example.com"],"extra_routes":{"10.101.0.0/16":{"via":"10.3.100.10"}},"ip_configs":{"Corp-Core":{"ip":"10.3.100.9","netmask":"/24","type":"static"},"Corp-Mgmt":{"ip":"10.3.172.9","netmask":"/24","type":"static"},"Corp-lan":{"ip":"10.3.171.9","netmask":"/24","type":"static"}},"name":"ITParis","ntpOverride":true,"ntp_servers":["10.3.51.222"],"path_preferences":{"core":{"paths":[{"networks":["Corp-Core"],"type":"local"}],"strategy":"ordered"},"lab":{"paths":[{"networks":["Corp-lan"],"type":"local"}],"strategy":"ordered"},"mgmt":{"paths":[{"networks":["Corp-Mgmt"],"type":"local"}],"strategy":"ordered"},"untrust":{"paths":[{"name":"wan","type":"wan"}],"strategy":"ordered"}},"port_config":{"ge-0/0/0":{"aggregated":false,"ip_config":{"gateway":"192.168.1.1","ip":"192.168.1.9","netmask":"/24","type":"static"},"name":"wan","redundant":false,"traffic_shaping":{"enabled":false},"usage":"wan","wan_type":"broadband"},"ge-0/0/6-7":{"ae_disable_lacp":false,"ae_idx":"0","ae_lacp_force_up":true,"aggregated":true,"networks":["Corp-lan","Corp-Mgmt","Corp-Core"],"usage":"lan"}},"service_policies":[{"action":"allow","idp":{"enabled":false},"name":"ITParis-Internal","path_preference":"core","services":["internal_dns","drive"],"tenants":["ITParis"]},{"action":"allow","idp":{"enabled":false},"name":"ITParis-internet","path_preference":"untrust","services":["internet_any"],"tenants":["ITParis"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-core","path_preference":"core","services":["internal_dns","internal_ntp"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-mxe-tt-in","path_preference":"mxe-in","services":["internal_any"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mgmt-to-untrust","path_preference":"untrust","services":["mxedge-updates","radsec","icmp","internet_dns","internet_ntp"],"tenants":["Corp-Mgmt"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"mxe-data-0-to-untrust","path_preference":"untrust","services":["internet_any"],"tenants":["ITParis"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-mgt","path_preference":"mgmt","services":["mgmt"],"tenants":["domain.Corp-Core","lan.Corp-Core","servers.Corp-Core","Corp-Core"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-edge-in","path_preference":"mxe-in","services":["internal_any"],"tenants":["lan.Corp-Core","Corp-Core"]},{"action":"allow","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"core-to-iot","path_preference":"iot","services":["iot"],"tenants":["lan.Corp-Core","servers-hassio.Corp-Core","servers-kubes.Corp-Core"]},{"action":"allow","idp":{"enabled":false},"name":"tanker-to-cctv","path_preference":"iot","services":["rtsp"],"tenants":["servers-tanker.Corp-Core"]},{"action":"allow","idp":{"enabled":false},"name":"core-to-untrust","path_preference":"untrust","services":["internet_any"],"tenants":["lan.Corp-Core","domain.Corp-Core","servers.Corp-Core"]},{"action":"deny","idp":{"alert_only":true,"enabled":true,"profile":"standard"},"name":"iot-upgrade-cctv","path_preference":"untrust","services":["motioneye","nodejs","raspbian"],"tenants":["printer"]}],"type":"standalone"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsGatewayTemplates.UpdateOrgGatewayTemplate(ctx, orgId, gatewaytemplateId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
