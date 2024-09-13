package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgSetting represents a OrgSetting struct.
// Org Settings
type OrgSetting struct {
    // enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
    ApUpdownThreshold           Optional[int]                          `json:"ap_updown_threshold"`
    ApiPolicy                   *OrgSettingApiPolicy                   `json:"api_policy,omitempty"`
    AutoDeviceNaming            *OrgSettingAutoDeviceNaming            `json:"auto_device_naming,omitempty"`
    AutoDeviceprofileAssignment *OrgSettingAutoDeviceprofileAssignment `json:"auto_deviceprofile_assignment,omitempty"`
    AutoSiteAssignment          *OrgSettingAutoSiteAssignment          `json:"auto_site_assignment,omitempty"`
    BlacklistUrl                *string                                `json:"blacklist_url,omitempty"`
    // list of PEM-encoded ca certs
    Cacerts                     []string                               `json:"cacerts,omitempty"`
    Celona                      *OrgSettingCelona                      `json:"celona,omitempty"`
    Cloudshark                  *OrgSettingCloudshark                  `json:"cloudshark,omitempty"`
    Cradlepoint                 *OrgSettingCradlepoint                 `json:"cradlepoint,omitempty"`
    CreatedTime                 *float64                               `json:"created_time,omitempty"`
    // common device cert, optional
    DeviceCert                  *OrgSettingDeviceCert                  `json:"device_cert,omitempty"`
    // enable threshold-based device down delivery via
    // * device-updowns webhooks topic,
    // * Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)
    DeviceUpdownThreshold       Optional[int]                          `json:"device_updown_threshold"`
    // whether to disallow Mist to analyze pcap files (this is required for marvis pcap)
    DisablePcap                 *bool                                  `json:"disable_pcap,omitempty"`
    // whether to disable remote shell access for an entire org
    DisableRemoteShell          *bool                                  `json:"disable_remote_shell,omitempty"`
    ForSite                     *bool                                  `json:"for_site,omitempty"`
    GatewayMgmt                 *OrgSettingGatewayMgmt                 `json:"gateway_mgmt,omitempty"`
    // enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
    GatewayUpdownThreshold      Optional[int]                          `json:"gateway_updown_threshold"`
    Id                          *uuid.UUID                             `json:"id,omitempty"`
    Installer                   *OrgSettingInstaller                   `json:"installer,omitempty"`
    Jcloud                      *OrgSettingJcloud                      `json:"jcloud,omitempty"`
    Juniper                     *AccountJuniperInfo                    `json:"juniper,omitempty"`
    // management-related properties
    Mgmt                        *OrgSettingMgmt                        `json:"mgmt,omitempty"`
    MistNac                     *OrgSettingMistNac                     `json:"mist_nac,omitempty"`
    ModifiedTime                *float64                               `json:"modified_time,omitempty"`
    MspId                       *uuid.UUID                             `json:"msp_id,omitempty"`
    MxedgeFipsEnabled           *bool                                  `json:"mxedge_fips_enabled,omitempty"`
    MxedgeMgmt                  *MxedgeMgmt                            `json:"mxedge_mgmt,omitempty"`
    OrgId                       *uuid.UUID                             `json:"org_id,omitempty"`
    // password policy
    PasswordPolicy              *OrgSettingPasswordPolicy              `json:"password_policy,omitempty"`
    Pcap                        *OrgSettingPcap                        `json:"pcap,omitempty"`
    PcapBucketVerified          *bool                                  `json:"pcap_bucket_verified,omitempty"`
    PortChannelization          *PortChannelization                    `json:"port_channelization,omitempty"`
    Security                    *OrgSettingSecurity                    `json:"security,omitempty"`
    // Set of heuristic rules will be enabled when marvis subscription is not available.
    // It triggers when, in a Z minute window, there are more than Y distinct client encountring over X failures
    SimpleAlert                 *SimpleAlert                           `json:"simple_alert,omitempty"`
    SwitchMgmt                  *OrgSettingSwitchMgmt                  `json:"switch_mgmt,omitempty"`
    // enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
    SwitchUpdownThreshold       Optional[int]                          `json:"switch_updown_threshold"`
    SyntheticTest               *SynthetictestConfig                   `json:"synthetic_test,omitempty"`
    // list of tags
    Tags                        []string                               `json:"tags,omitempty"`
    // automatically logout the user when UI session is inactive. `0` means disabled
    UiIdleTimeout               *int                                   `json:"ui_idle_timeout,omitempty"`
    VpnOptions                  *OrgSettingVpnOptions                  `json:"vpn_options,omitempty"`
    WanPma                      *OrgSettingWanPma                      `json:"wan_pma,omitempty"`
    WiredPma                    *OrgSettingWiredPma                    `json:"wired_pma,omitempty"`
    WirelessPma                 *OrgSettingWirelessPma                 `json:"wireless_pma,omitempty"`
    AdditionalProperties        map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSetting.
// It customizes the JSON marshaling process for OrgSetting objects.
func (o OrgSetting) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSetting object to a map representation for JSON marshaling.
func (o OrgSetting) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ApUpdownThreshold.IsValueSet() {
        if o.ApUpdownThreshold.Value() != nil {
            structMap["ap_updown_threshold"] = o.ApUpdownThreshold.Value()
        } else {
            structMap["ap_updown_threshold"] = nil
        }
    }
    if o.ApiPolicy != nil {
        structMap["api_policy"] = o.ApiPolicy.toMap()
    }
    if o.AutoDeviceNaming != nil {
        structMap["auto_device_naming"] = o.AutoDeviceNaming.toMap()
    }
    if o.AutoDeviceprofileAssignment != nil {
        structMap["auto_deviceprofile_assignment"] = o.AutoDeviceprofileAssignment.toMap()
    }
    if o.AutoSiteAssignment != nil {
        structMap["auto_site_assignment"] = o.AutoSiteAssignment.toMap()
    }
    if o.BlacklistUrl != nil {
        structMap["blacklist_url"] = o.BlacklistUrl
    }
    if o.Cacerts != nil {
        structMap["cacerts"] = o.Cacerts
    }
    if o.Celona != nil {
        structMap["celona"] = o.Celona.toMap()
    }
    if o.Cloudshark != nil {
        structMap["cloudshark"] = o.Cloudshark.toMap()
    }
    if o.Cradlepoint != nil {
        structMap["cradlepoint"] = o.Cradlepoint.toMap()
    }
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.DeviceCert != nil {
        structMap["device_cert"] = o.DeviceCert.toMap()
    }
    if o.DeviceUpdownThreshold.IsValueSet() {
        if o.DeviceUpdownThreshold.Value() != nil {
            structMap["device_updown_threshold"] = o.DeviceUpdownThreshold.Value()
        } else {
            structMap["device_updown_threshold"] = nil
        }
    }
    if o.DisablePcap != nil {
        structMap["disable_pcap"] = o.DisablePcap
    }
    if o.DisableRemoteShell != nil {
        structMap["disable_remote_shell"] = o.DisableRemoteShell
    }
    if o.ForSite != nil {
        structMap["for_site"] = o.ForSite
    }
    if o.GatewayMgmt != nil {
        structMap["gateway_mgmt"] = o.GatewayMgmt.toMap()
    }
    if o.GatewayUpdownThreshold.IsValueSet() {
        if o.GatewayUpdownThreshold.Value() != nil {
            structMap["gateway_updown_threshold"] = o.GatewayUpdownThreshold.Value()
        } else {
            structMap["gateway_updown_threshold"] = nil
        }
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.Installer != nil {
        structMap["installer"] = o.Installer.toMap()
    }
    if o.Jcloud != nil {
        structMap["jcloud"] = o.Jcloud.toMap()
    }
    if o.Juniper != nil {
        structMap["juniper"] = o.Juniper.toMap()
    }
    if o.Mgmt != nil {
        structMap["mgmt"] = o.Mgmt.toMap()
    }
    if o.MistNac != nil {
        structMap["mist_nac"] = o.MistNac.toMap()
    }
    if o.ModifiedTime != nil {
        structMap["modified_time"] = o.ModifiedTime
    }
    if o.MspId != nil {
        structMap["msp_id"] = o.MspId
    }
    if o.MxedgeFipsEnabled != nil {
        structMap["mxedge_fips_enabled"] = o.MxedgeFipsEnabled
    }
    if o.MxedgeMgmt != nil {
        structMap["mxedge_mgmt"] = o.MxedgeMgmt.toMap()
    }
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.PasswordPolicy != nil {
        structMap["password_policy"] = o.PasswordPolicy.toMap()
    }
    if o.Pcap != nil {
        structMap["pcap"] = o.Pcap.toMap()
    }
    if o.PcapBucketVerified != nil {
        structMap["pcap_bucket_verified"] = o.PcapBucketVerified
    }
    if o.PortChannelization != nil {
        structMap["port_channelization"] = o.PortChannelization.toMap()
    }
    if o.Security != nil {
        structMap["security"] = o.Security.toMap()
    }
    if o.SimpleAlert != nil {
        structMap["simple_alert"] = o.SimpleAlert.toMap()
    }
    if o.SwitchMgmt != nil {
        structMap["switch_mgmt"] = o.SwitchMgmt.toMap()
    }
    if o.SwitchUpdownThreshold.IsValueSet() {
        if o.SwitchUpdownThreshold.Value() != nil {
            structMap["switch_updown_threshold"] = o.SwitchUpdownThreshold.Value()
        } else {
            structMap["switch_updown_threshold"] = nil
        }
    }
    if o.SyntheticTest != nil {
        structMap["synthetic_test"] = o.SyntheticTest.toMap()
    }
    if o.Tags != nil {
        structMap["tags"] = o.Tags
    }
    if o.UiIdleTimeout != nil {
        structMap["ui_idle_timeout"] = o.UiIdleTimeout
    }
    if o.VpnOptions != nil {
        structMap["vpn_options"] = o.VpnOptions.toMap()
    }
    if o.WanPma != nil {
        structMap["wan_pma"] = o.WanPma.toMap()
    }
    if o.WiredPma != nil {
        structMap["wired_pma"] = o.WiredPma.toMap()
    }
    if o.WirelessPma != nil {
        structMap["wireless_pma"] = o.WirelessPma.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSetting.
// It customizes the JSON unmarshaling process for OrgSetting objects.
func (o *OrgSetting) UnmarshalJSON(input []byte) error {
    var temp tempOrgSetting
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_updown_threshold", "api_policy", "auto_device_naming", "auto_deviceprofile_assignment", "auto_site_assignment", "blacklist_url", "cacerts", "celona", "cloudshark", "cradlepoint", "created_time", "device_cert", "device_updown_threshold", "disable_pcap", "disable_remote_shell", "for_site", "gateway_mgmt", "gateway_updown_threshold", "id", "installer", "jcloud", "juniper", "mgmt", "mist_nac", "modified_time", "msp_id", "mxedge_fips_enabled", "mxedge_mgmt", "org_id", "password_policy", "pcap", "pcap_bucket_verified", "port_channelization", "security", "simple_alert", "switch_mgmt", "switch_updown_threshold", "synthetic_test", "tags", "ui_idle_timeout", "vpn_options", "wan_pma", "wired_pma", "wireless_pma")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.ApUpdownThreshold = temp.ApUpdownThreshold
    o.ApiPolicy = temp.ApiPolicy
    o.AutoDeviceNaming = temp.AutoDeviceNaming
    o.AutoDeviceprofileAssignment = temp.AutoDeviceprofileAssignment
    o.AutoSiteAssignment = temp.AutoSiteAssignment
    o.BlacklistUrl = temp.BlacklistUrl
    o.Cacerts = temp.Cacerts
    o.Celona = temp.Celona
    o.Cloudshark = temp.Cloudshark
    o.Cradlepoint = temp.Cradlepoint
    o.CreatedTime = temp.CreatedTime
    o.DeviceCert = temp.DeviceCert
    o.DeviceUpdownThreshold = temp.DeviceUpdownThreshold
    o.DisablePcap = temp.DisablePcap
    o.DisableRemoteShell = temp.DisableRemoteShell
    o.ForSite = temp.ForSite
    o.GatewayMgmt = temp.GatewayMgmt
    o.GatewayUpdownThreshold = temp.GatewayUpdownThreshold
    o.Id = temp.Id
    o.Installer = temp.Installer
    o.Jcloud = temp.Jcloud
    o.Juniper = temp.Juniper
    o.Mgmt = temp.Mgmt
    o.MistNac = temp.MistNac
    o.ModifiedTime = temp.ModifiedTime
    o.MspId = temp.MspId
    o.MxedgeFipsEnabled = temp.MxedgeFipsEnabled
    o.MxedgeMgmt = temp.MxedgeMgmt
    o.OrgId = temp.OrgId
    o.PasswordPolicy = temp.PasswordPolicy
    o.Pcap = temp.Pcap
    o.PcapBucketVerified = temp.PcapBucketVerified
    o.PortChannelization = temp.PortChannelization
    o.Security = temp.Security
    o.SimpleAlert = temp.SimpleAlert
    o.SwitchMgmt = temp.SwitchMgmt
    o.SwitchUpdownThreshold = temp.SwitchUpdownThreshold
    o.SyntheticTest = temp.SyntheticTest
    o.Tags = temp.Tags
    o.UiIdleTimeout = temp.UiIdleTimeout
    o.VpnOptions = temp.VpnOptions
    o.WanPma = temp.WanPma
    o.WiredPma = temp.WiredPma
    o.WirelessPma = temp.WirelessPma
    return nil
}

// tempOrgSetting is a temporary struct used for validating the fields of OrgSetting.
type tempOrgSetting  struct {
    ApUpdownThreshold           Optional[int]                          `json:"ap_updown_threshold"`
    ApiPolicy                   *OrgSettingApiPolicy                   `json:"api_policy,omitempty"`
    AutoDeviceNaming            *OrgSettingAutoDeviceNaming            `json:"auto_device_naming,omitempty"`
    AutoDeviceprofileAssignment *OrgSettingAutoDeviceprofileAssignment `json:"auto_deviceprofile_assignment,omitempty"`
    AutoSiteAssignment          *OrgSettingAutoSiteAssignment          `json:"auto_site_assignment,omitempty"`
    BlacklistUrl                *string                                `json:"blacklist_url,omitempty"`
    Cacerts                     []string                               `json:"cacerts,omitempty"`
    Celona                      *OrgSettingCelona                      `json:"celona,omitempty"`
    Cloudshark                  *OrgSettingCloudshark                  `json:"cloudshark,omitempty"`
    Cradlepoint                 *OrgSettingCradlepoint                 `json:"cradlepoint,omitempty"`
    CreatedTime                 *float64                               `json:"created_time,omitempty"`
    DeviceCert                  *OrgSettingDeviceCert                  `json:"device_cert,omitempty"`
    DeviceUpdownThreshold       Optional[int]                          `json:"device_updown_threshold"`
    DisablePcap                 *bool                                  `json:"disable_pcap,omitempty"`
    DisableRemoteShell          *bool                                  `json:"disable_remote_shell,omitempty"`
    ForSite                     *bool                                  `json:"for_site,omitempty"`
    GatewayMgmt                 *OrgSettingGatewayMgmt                 `json:"gateway_mgmt,omitempty"`
    GatewayUpdownThreshold      Optional[int]                          `json:"gateway_updown_threshold"`
    Id                          *uuid.UUID                             `json:"id,omitempty"`
    Installer                   *OrgSettingInstaller                   `json:"installer,omitempty"`
    Jcloud                      *OrgSettingJcloud                      `json:"jcloud,omitempty"`
    Juniper                     *AccountJuniperInfo                    `json:"juniper,omitempty"`
    Mgmt                        *OrgSettingMgmt                        `json:"mgmt,omitempty"`
    MistNac                     *OrgSettingMistNac                     `json:"mist_nac,omitempty"`
    ModifiedTime                *float64                               `json:"modified_time,omitempty"`
    MspId                       *uuid.UUID                             `json:"msp_id,omitempty"`
    MxedgeFipsEnabled           *bool                                  `json:"mxedge_fips_enabled,omitempty"`
    MxedgeMgmt                  *MxedgeMgmt                            `json:"mxedge_mgmt,omitempty"`
    OrgId                       *uuid.UUID                             `json:"org_id,omitempty"`
    PasswordPolicy              *OrgSettingPasswordPolicy              `json:"password_policy,omitempty"`
    Pcap                        *OrgSettingPcap                        `json:"pcap,omitempty"`
    PcapBucketVerified          *bool                                  `json:"pcap_bucket_verified,omitempty"`
    PortChannelization          *PortChannelization                    `json:"port_channelization,omitempty"`
    Security                    *OrgSettingSecurity                    `json:"security,omitempty"`
    SimpleAlert                 *SimpleAlert                           `json:"simple_alert,omitempty"`
    SwitchMgmt                  *OrgSettingSwitchMgmt                  `json:"switch_mgmt,omitempty"`
    SwitchUpdownThreshold       Optional[int]                          `json:"switch_updown_threshold"`
    SyntheticTest               *SynthetictestConfig                   `json:"synthetic_test,omitempty"`
    Tags                        []string                               `json:"tags,omitempty"`
    UiIdleTimeout               *int                                   `json:"ui_idle_timeout,omitempty"`
    VpnOptions                  *OrgSettingVpnOptions                  `json:"vpn_options,omitempty"`
    WanPma                      *OrgSettingWanPma                      `json:"wan_pma,omitempty"`
    WiredPma                    *OrgSettingWiredPma                    `json:"wired_pma,omitempty"`
    WirelessPma                 *OrgSettingWirelessPma                 `json:"wireless_pma,omitempty"`
}
