package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsDeviceProfilesTestListOrgDeviceProfiles tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestListOrgDeviceProfiles(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsDeviceProfiles.ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"aeroscout":{"enabled":false,"host":"aero.pvt.net","locate_connected":true},"ble_config":{"beacon_enabled":false,"beacon_rate":3,"beacon_rate_mode":"custom","beam_disabled":[1,3,6],"custom_ble_packet_enabled":false,"custom_ble_packet_frame":"0x........","custom_ble_packet_freq_msec":300,"eddystone_uid_adv_power":-65,"eddystone_uid_beams":"2-4,7","eddystone_uid_enabled":false,"eddystone_uid_freq_msec":200,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_adv_power":-65,"eddystone_url_beams":"2-4,7","eddystone_url_enabled":true,"eddystone_url_freq_msec":1000,"eddystone_url_url":"https://www.abc.com","ibeacon_adv_power":-65,"ibeacon_beams":"2-4,7","ibeacon_enabled":false,"ibeacon_freq_msec":0,"ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","power":7,"power_mode":"custom"},"created_time":0,"disable_eth1":false,"disable_eth2":false,"disable_eth3":false,"disable_module":false,"for_site":true,"height":0,"id":"497f6eca-6276-4993-bfeb-53cbbbba6108","ip_config":{"dns":["8.8.8.8","4.4.4.4"],"dns_suffix":[".mist.local",".mist.com"],"gateway":"10.2.1.254","gateway6":"2607:f8b0:4005:808::1","ip":"10.2.1.1","ip6":"2607:f8b0:4005:808::2004","netmask":"255.255.255.0","netmask6":"/32","type":"static","type6":"static","vlan_id":1},"led":{"brightness":255,"enabled":true},"map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mesh":{"enabled":false,"group":1,"role":"base"},"modified_time":0,"name":"string","notes":"string","ntp_servers":["string"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","orientation":0,"poe_passthrough":false,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","type":"ap","x":0,"y":0}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestCreateOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestCreateOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsDeviceProfiles.CreateOrgDeviceProfile(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":0,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestDeleteOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestDeleteOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsDeviceProfiles.DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsDeviceProfilesTestGetOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestGetOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsDeviceProfiles.GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":0,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsDeviceProfiles.UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":0,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestAssignOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestAssignOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MacAddresses
    errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsDeviceProfiles.AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"success":["5c5b350e0001"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.MacAddresses
    errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsDeviceProfiles.UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"success":["5c5b350e0001"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
