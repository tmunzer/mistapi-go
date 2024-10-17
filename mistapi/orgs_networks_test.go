package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsNetworksTestListOrgNetworks tests the behavior of the OrgsNetworks
func TestOrgsNetworksTestListOrgNetworks(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsNetworks.ListOrgNetworks(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f13","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.45"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworksTestCreateOrgNetwork tests the behavior of the OrgsNetworks
func TestOrgsNetworksTestCreateOrgNetwork(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Network
	errBody := json.Unmarshal([]byte(`{"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"name":"string","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10."]},"property2":{"addresses":["10.10.10.52"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsNetworks.CreateOrgNetwork(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f12","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.52"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworksTestDeleteOrgNetwork tests the behavior of the OrgsNetworks
func TestOrgsNetworksTestDeleteOrgNetwork(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	networkId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsNetworks.DeleteOrgNetwork(ctx, orgId, networkId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNetworksTestGetOrgNetwork tests the behavior of the OrgsNetworks
func TestOrgsNetworksTestGetOrgNetwork(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	networkId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsNetworks.GetOrgNetwork(ctx, orgId, networkId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f12","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.52"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworksTestUpdateOrgNetwork tests the behavior of the OrgsNetworks
func TestOrgsNetworksTestUpdateOrgNetwork(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	networkId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Network
	errBody := json.Unmarshal([]byte(`{"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"name":"string","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.52"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsNetworks.UpdateOrgNetwork(ctx, orgId, networkId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f12","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":443},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":443}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.52"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":443,"to":"192.168.70.5/30"},"property2":{"name":"web server","port":443,"to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
