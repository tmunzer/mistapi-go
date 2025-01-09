package models

// AclTagTypeEnum is a string enum.
// enum: 
// * `any`: matching anything not identified
// * `dynamic_gbp`: from the gbp_tag received from RADIUS
// * `gbp_resource`: can only be used in `dst_tags`
// * `mac`
// * `network`
// * `radius_group`
// * `resource`: can only be used in `dst_tags`
// * `static_gbp`: applying gbp tag against matching conditions
// * `subnet`'
type AclTagTypeEnum string

const (
    AclTagTypeEnum_ANY         AclTagTypeEnum = "any"
    AclTagTypeEnum_DYNAMICGBP  AclTagTypeEnum = "dynamic_gbp"
    AclTagTypeEnum_GBPRESOURCE AclTagTypeEnum = "gbp_resource"
    AclTagTypeEnum_MAC         AclTagTypeEnum = "mac"
    AclTagTypeEnum_NETWORK     AclTagTypeEnum = "network"
    AclTagTypeEnum_RADIUSGROUP AclTagTypeEnum = "radius_group"
    AclTagTypeEnum_RESOURCE    AclTagTypeEnum = "resource"
    AclTagTypeEnum_STATICGBP   AclTagTypeEnum = "static_gbp"
    AclTagTypeEnum_SUBNET      AclTagTypeEnum = "subnet"
)

// AdminComplianceStatusEnum is a string enum.
// trade compliance status. enum: `blocked`, `restricted`
type AdminComplianceStatusEnum string

const (
    AdminComplianceStatusEnum_BLOCKED    AdminComplianceStatusEnum = "blocked"
    AdminComplianceStatusEnum_RESTRICTED AdminComplianceStatusEnum = "restricted"
)

// AdminPrivilegeRoleEnum is a string enum.
// access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
type AdminPrivilegeRoleEnum string

const (
    AdminPrivilegeRoleEnum_ADMIN     AdminPrivilegeRoleEnum = "admin"
    AdminPrivilegeRoleEnum_HELPDESK  AdminPrivilegeRoleEnum = "helpdesk"
    AdminPrivilegeRoleEnum_INSTALLER AdminPrivilegeRoleEnum = "installer"
    AdminPrivilegeRoleEnum_READ      AdminPrivilegeRoleEnum = "read"
    AdminPrivilegeRoleEnum_WRITE     AdminPrivilegeRoleEnum = "write"
)

// AdminPrivilegeScopeEnum is a string enum.
// enum: `msp`, `org`, `orggroup`, `site`, `sitegroup`
type AdminPrivilegeScopeEnum string

const (
    AdminPrivilegeScopeEnum_MSP       AdminPrivilegeScopeEnum = "msp"
    AdminPrivilegeScopeEnum_ORG       AdminPrivilegeScopeEnum = "org"
    AdminPrivilegeScopeEnum_ORGGROUP  AdminPrivilegeScopeEnum = "orggroup"
    AdminPrivilegeScopeEnum_SITE      AdminPrivilegeScopeEnum = "site"
    AdminPrivilegeScopeEnum_SITEGROUP AdminPrivilegeScopeEnum = "sitegroup"
)

// AdminPrivilegeViewEnum is a string enum.
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  
// You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  
// Below are the list of supported UI views. Note that this is UI only feature.  
// | UI View | Required Role | Description |
// | --- | --- | --- |
// | `reporting` | `read` | full access to all analytics tools |
// | `marketing` | `read` | can view analytics and location maps |
// | `super_observer` | `read` | can view all the organization except the subscription page |
// | `location` | `write` | can view and manage location maps, can view analytics |
// | `security` | `write` | can view and manage site labels, policies and security |
// | `switch_admin` | `helpdesk` | can view and manage Switch ports, can view wired clients |
// | `mxedge_admin` | `admin` | can view and manage Mist edges and Mist tunnels |
// | `lobby_admin` | `admin` | full access to Org and Site Pre-shared keys |
type AdminPrivilegeViewEnum string

const (
    AdminPrivilegeViewEnum_LOBBYADMIN    AdminPrivilegeViewEnum = "lobby_admin"
    AdminPrivilegeViewEnum_LOCATION      AdminPrivilegeViewEnum = "location"
    AdminPrivilegeViewEnum_MARKETING     AdminPrivilegeViewEnum = "marketing"
    AdminPrivilegeViewEnum_MXEDGEADMIN   AdminPrivilegeViewEnum = "mxedge_admin"
    AdminPrivilegeViewEnum_REPORTING     AdminPrivilegeViewEnum = "reporting"
    AdminPrivilegeViewEnum_SECURITY      AdminPrivilegeViewEnum = "security"
    AdminPrivilegeViewEnum_SUPEROBSERVER AdminPrivilegeViewEnum = "super_observer"
    AdminPrivilegeViewEnum_SWITCHADMIN   AdminPrivilegeViewEnum = "switch_admin"
)

// AlarmCountDisctinctEnum is a string enum.
// enum: `acked`, `group`, `severity`, `type`
type AlarmCountDisctinctEnum string

const (
    AlarmCountDisctinctEnum_ACKED    AlarmCountDisctinctEnum = "acked"
    AlarmCountDisctinctEnum_GROUP    AlarmCountDisctinctEnum = "group"
    AlarmCountDisctinctEnum_SEVERITY AlarmCountDisctinctEnum = "severity"
    AlarmCountDisctinctEnum_ENUMTYPE AlarmCountDisctinctEnum = "type"
)

// AllowDenyEnum is a string enum.
// enum: `allow`, `deny`
type AllowDenyEnum string

const (
    AllowDenyEnum_ALLOW AllowDenyEnum = "allow"
    AllowDenyEnum_DENY  AllowDenyEnum = "deny"
)

// ApClientBridgeAuthTypeEnum is a string enum.
// wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`
type ApClientBridgeAuthTypeEnum string

const (
    ApClientBridgeAuthTypeEnum_OPEN ApClientBridgeAuthTypeEnum = "open"
    ApClientBridgeAuthTypeEnum_PSK  ApClientBridgeAuthTypeEnum = "psk"
)

// ApEslTypeEnum is a string enum.
// note: ble_config will be ingored if esl_config is enabled and with native mode. enum: `hanshow`, `imagotag`, `native`, `solum`
type ApEslTypeEnum string

const (
    ApEslTypeEnum_HANSHOW  ApEslTypeEnum = "hanshow"
    ApEslTypeEnum_IMAGOTAG ApEslTypeEnum = "imagotag"
    ApEslTypeEnum_NATIVE   ApEslTypeEnum = "native"
    ApEslTypeEnum_SOLUM    ApEslTypeEnum = "solum"
)

// ApIotPullupEnum is a string enum.
// the type of pull-up the pin uses. enum: `external`, `internal`, `none`
type ApIotPullupEnum string

const (
    ApIotPullupEnum_EXTERNAL ApIotPullupEnum = "external"
    ApIotPullupEnum_INTERNAL ApIotPullupEnum = "internal"
    ApIotPullupEnum_NONE     ApIotPullupEnum = "none"
)

// ApMeshRoleEnum is a string enum.
// enum: `base`, `remote`
type ApMeshRoleEnum string

const (
    ApMeshRoleEnum_BASE   ApMeshRoleEnum = "base"
    ApMeshRoleEnum_REMOTE ApMeshRoleEnum = "remote"
)

// ApPortConfigForwardingEnum is a string enum.
// enum: `all`, `limited`, `mxtunnel`, `site_mxedge`, `wxtunnel`
type ApPortConfigForwardingEnum string

const (
    ApPortConfigForwardingEnum_ALL        ApPortConfigForwardingEnum = "all"
    ApPortConfigForwardingEnum_LIMITED    ApPortConfigForwardingEnum = "limited"
    ApPortConfigForwardingEnum_MXTUNNEL   ApPortConfigForwardingEnum = "mxtunnel"
    ApPortConfigForwardingEnum_SITEMXEDGE ApPortConfigForwardingEnum = "site_mxedge"
    ApPortConfigForwardingEnum_WXTUNNEL   ApPortConfigForwardingEnum = "wxtunnel"
)

// ApPortConfigMacAuthProtocolEnum is a string enum.
// if `enable_mac_auth`==`true`, allows user to select an authentication protocol. enum: `eap-md5`, `eap-peap`, `pap`
type ApPortConfigMacAuthProtocolEnum string

const (
    ApPortConfigMacAuthProtocolEnum_EAPMD5  ApPortConfigMacAuthProtocolEnum = "eap-md5"
    ApPortConfigMacAuthProtocolEnum_EAPPEAP ApPortConfigMacAuthProtocolEnum = "eap-peap"
    ApPortConfigMacAuthProtocolEnum_PAP     ApPortConfigMacAuthProtocolEnum = "pap"
)

// ApPortConfigPortAuthEnum is a string enum.
// When doing port auth. enum: `dot1x`, `none`
type ApPortConfigPortAuthEnum string

const (
    ApPortConfigPortAuthEnum_DOT1X ApPortConfigPortAuthEnum = "dot1x"
    ApPortConfigPortAuthEnum_NONE  ApPortConfigPortAuthEnum = "none"
)

// ApRadioAntennaModeEnum is a string enum.
// enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`
type ApRadioAntennaModeEnum string

const (
    ApRadioAntennaModeEnum_ENUM1X1     ApRadioAntennaModeEnum = "1x1"
    ApRadioAntennaModeEnum_ENUM2X2     ApRadioAntennaModeEnum = "2x2"
    ApRadioAntennaModeEnum_ENUM3X3     ApRadioAntennaModeEnum = "3x3"
    ApRadioAntennaModeEnum_ENUM4X4     ApRadioAntennaModeEnum = "4x4"
    ApRadioAntennaModeEnum_ENUMDEFAULT ApRadioAntennaModeEnum = "default"
)

// ApUsbTypeEnum is a string enum.
// usb config type. enum: `hanshow`, `imagotag`, `solum`
type ApUsbTypeEnum string

const (
    ApUsbTypeEnum_HANSHOW  ApUsbTypeEnum = "hanshow"
    ApUsbTypeEnum_IMAGOTAG ApUsbTypeEnum = "imagotag"
    ApUsbTypeEnum_SOLUM    ApUsbTypeEnum = "solum"
)

// AppProbingCustomAppProtocolEnum is a string enum.
// enum: `http`, `icmp`
type AppProbingCustomAppProtocolEnum string

const (
    AppProbingCustomAppProtocolEnum_HTTP AppProbingCustomAppProtocolEnum = "http"
    AppProbingCustomAppProtocolEnum_ICMP AppProbingCustomAppProtocolEnum = "icmp"
)

// AutoOrientationStateEnum is a string enum.
// The state of auto orient for a given map derived from an Enum. enum: `Enqueued`, `Not Started`, `Oriented`
type AutoOrientationStateEnum string

const (
    AutoOrientationStateEnum_ENQUEUED       AutoOrientationStateEnum = "Enqueued"
    AutoOrientationStateEnum_ENUMNOTSTARTED AutoOrientationStateEnum = "Not Started"
    AutoOrientationStateEnum_ORIENTED       AutoOrientationStateEnum = "Oriented"
)

// AutoPlacementInfoStatusEnum is a string enum.
// the status of autoplacement for a given map. enum: `done`, `error`, `inprogress`, `pending`
type AutoPlacementInfoStatusEnum string

const (
    AutoPlacementInfoStatusEnum_DONE       AutoPlacementInfoStatusEnum = "done"
    AutoPlacementInfoStatusEnum_ENUMERROR  AutoPlacementInfoStatusEnum = "error"
    AutoPlacementInfoStatusEnum_INPROGRESS AutoPlacementInfoStatusEnum = "inprogress"
    AutoPlacementInfoStatusEnum_PENDING    AutoPlacementInfoStatusEnum = "pending"
)

// AvprofileFallbackActionEnum is a string enum.
// enum: `block`, `permit`
type AvprofileFallbackActionEnum string

const (
    AvprofileFallbackActionEnum_BLOCK  AvprofileFallbackActionEnum = "block"
    AvprofileFallbackActionEnum_PERMIT AvprofileFallbackActionEnum = "permit"
)

// AvprofileProtocolsEnum is a string enum.
// enum: `ftp`, `http`, `imap`, `pop3`, `smtp`
type AvprofileProtocolsEnum string

const (
    AvprofileProtocolsEnum_FTP  AvprofileProtocolsEnum = "ftp"
    AvprofileProtocolsEnum_HTTP AvprofileProtocolsEnum = "http"
    AvprofileProtocolsEnum_IMAP AvprofileProtocolsEnum = "imap"
    AvprofileProtocolsEnum_POP3 AvprofileProtocolsEnum = "pop3"
    AvprofileProtocolsEnum_SMTP AvprofileProtocolsEnum = "smtp"
)

// BeaconTypeEnum is a string enum.
// enum: `eddystone-uid`, `eddystone-url`, `ibeacon`
type BeaconTypeEnum string

const (
    BeaconTypeEnum_EDDYSTONEUID BeaconTypeEnum = "eddystone-uid"
    BeaconTypeEnum_EDDYSTONEURL BeaconTypeEnum = "eddystone-url"
    BeaconTypeEnum_IBEACON      BeaconTypeEnum = "ibeacon"
)

// BgpConfigTypeEnum is a string enum.
// enum: `external`, `internal`
type BgpConfigTypeEnum string

const (
    BgpConfigTypeEnum_EXTERNAL BgpConfigTypeEnum = "external"
    BgpConfigTypeEnum_INTERNAL BgpConfigTypeEnum = "internal"
)

// BgpConfigViaEnum is a string enum.
// network name. enum: `lan`, `tunnel`, `vpn`, `wan`
type BgpConfigViaEnum string

const (
    BgpConfigViaEnum_LAN    BgpConfigViaEnum = "lan"
    BgpConfigViaEnum_TUNNEL BgpConfigViaEnum = "tunnel"
    BgpConfigViaEnum_VPN    BgpConfigViaEnum = "vpn"
    BgpConfigViaEnum_WAN    BgpConfigViaEnum = "wan"
)

// BgpStatsStateEnum is a string enum.
// enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent`
type BgpStatsStateEnum string

const (
    BgpStatsStateEnum_ACTIVE      BgpStatsStateEnum = "active"
    BgpStatsStateEnum_CONNECT     BgpStatsStateEnum = "connect"
    BgpStatsStateEnum_ESTABLISHED BgpStatsStateEnum = "established"
    BgpStatsStateEnum_IDLE        BgpStatsStateEnum = "idle"
    BgpStatsStateEnum_OPENCONFIG  BgpStatsStateEnum = "open_config"
    BgpStatsStateEnum_OPENSENT    BgpStatsStateEnum = "open_sent"
)

// BleConfigBeaconRateModeEnum is a string enum.
// enum: `custom`, `default`
type BleConfigBeaconRateModeEnum string

const (
    BleConfigBeaconRateModeEnum_CUSTOM      BleConfigBeaconRateModeEnum = "custom"
    BleConfigBeaconRateModeEnum_ENUMDEFAULT BleConfigBeaconRateModeEnum = "default"
)

// BleConfigPowerModeEnum is a string enum.
// enum: `custom`, `default`
type BleConfigPowerModeEnum string

const (
    BleConfigPowerModeEnum_CUSTOM      BleConfigPowerModeEnum = "custom"
    BleConfigPowerModeEnum_ENUMDEFAULT BleConfigPowerModeEnum = "default"
)

// CaptureClientTypeEnum is a string enum.
// enum: `client`
type CaptureClientTypeEnum string

const (
    CaptureClientTypeEnum_CLIENT CaptureClientTypeEnum = "client"
)

// CaptureGatewayFormatEnum is a string enum.
// enum: `stream`
type CaptureGatewayFormatEnum string

const (
    CaptureGatewayFormatEnum_STREAM CaptureGatewayFormatEnum = "stream"
)

// CaptureGatewayTypeEnum is a string enum.
// enum: `gateway`
type CaptureGatewayTypeEnum string

const (
    CaptureGatewayTypeEnum_GATEWAY CaptureGatewayTypeEnum = "gateway"
)

// CaptureMxedgeFormatEnum is a string enum.
// pcap format. enum: 
// * `stream`: to Mist cloud
// * `tzsp`: tream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)
type CaptureMxedgeFormatEnum string

const (
    CaptureMxedgeFormatEnum_STREAM CaptureMxedgeFormatEnum = "stream"
    CaptureMxedgeFormatEnum_TZSP   CaptureMxedgeFormatEnum = "tzsp"
)

// CaptureMxedgeTypeEnum is a string enum.
// enum: `mxedge`
type CaptureMxedgeTypeEnum string

const (
    CaptureMxedgeTypeEnum_MXEDGE CaptureMxedgeTypeEnum = "mxedge"
)

// CaptureNewAssocTypeEnum is a string enum.
// enum: `new_assoc`
type CaptureNewAssocTypeEnum string

const (
    CaptureNewAssocTypeEnum_NEWASSOC CaptureNewAssocTypeEnum = "new_assoc"
)

// CaptureRadiotapBandEnum is a string enum.
// enum: `24`, `24,5,6`, `5`, `6`
type CaptureRadiotapBandEnum string

const (
    CaptureRadiotapBandEnum_ENUM24   CaptureRadiotapBandEnum = "24"
    CaptureRadiotapBandEnum_ENUM2456 CaptureRadiotapBandEnum = "24,5,6"
    CaptureRadiotapBandEnum_ENUM5    CaptureRadiotapBandEnum = "5"
    CaptureRadiotapBandEnum_ENUM6    CaptureRadiotapBandEnum = "6"
)

// CaptureRadiotapFormatEnum is a string enum.
// enum: `pcap`, `stream`
type CaptureRadiotapFormatEnum string

const (
    CaptureRadiotapFormatEnum_PCAP   CaptureRadiotapFormatEnum = "pcap"
    CaptureRadiotapFormatEnum_STREAM CaptureRadiotapFormatEnum = "stream"
)

// CaptureRadiotapTypeEnum is a string enum.
// enum: `radiotap`
type CaptureRadiotapTypeEnum string

const (
    CaptureRadiotapTypeEnum_RADIOTAP CaptureRadiotapTypeEnum = "radiotap"
)

// CaptureRadiotapwiredBandEnum is a string enum.
// only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`
type CaptureRadiotapwiredBandEnum string

const (
    CaptureRadiotapwiredBandEnum_ENUM24   CaptureRadiotapwiredBandEnum = "24"
    CaptureRadiotapwiredBandEnum_ENUM2456 CaptureRadiotapwiredBandEnum = "24,5,6"
    CaptureRadiotapwiredBandEnum_ENUM5    CaptureRadiotapwiredBandEnum = "5"
    CaptureRadiotapwiredBandEnum_ENUM6    CaptureRadiotapwiredBandEnum = "6"
)

// CaptureRadiotapwiredFormatEnum is a string enum.
// enum: `pcap`, `stream`
type CaptureRadiotapwiredFormatEnum string

const (
    CaptureRadiotapwiredFormatEnum_PCAP   CaptureRadiotapwiredFormatEnum = "pcap"
    CaptureRadiotapwiredFormatEnum_STREAM CaptureRadiotapwiredFormatEnum = "stream"
)

// CaptureRadiotapwiredTypeEnum is a string enum.
// enum: `radiotap,wired`
type CaptureRadiotapwiredTypeEnum string

const (
    CaptureRadiotapwiredTypeEnum_ENUMRADIOTAPWIRED CaptureRadiotapwiredTypeEnum = "radiotap,wired"
)

// CaptureScanApsBandEnum is a string enum.
// Only Single value allowed. enum: `24`, `5`, `6`
type CaptureScanApsBandEnum string

const (
    CaptureScanApsBandEnum_ENUM24 CaptureScanApsBandEnum = "24"
    CaptureScanApsBandEnum_ENUM5  CaptureScanApsBandEnum = "5"
    CaptureScanApsBandEnum_ENUM6  CaptureScanApsBandEnum = "6"
)

// CaptureScanBandEnum is a string enum.
// Only Single value allowed, default value gets applied when user provides wrong values. enum: `24`, `5`, `6`
type CaptureScanBandEnum string

const (
    CaptureScanBandEnum_ENUM24 CaptureScanBandEnum = "24"
    CaptureScanBandEnum_ENUM5  CaptureScanBandEnum = "5"
    CaptureScanBandEnum_ENUM6  CaptureScanBandEnum = "6"
)

// CaptureScanFormatEnum is a string enum.
// enum: `pcap`, `stream`
type CaptureScanFormatEnum string

const (
    CaptureScanFormatEnum_PCAP   CaptureScanFormatEnum = "pcap"
    CaptureScanFormatEnum_STREAM CaptureScanFormatEnum = "stream"
)

// CaptureScanTypeEnum is a string enum.
// enum: `scan`
type CaptureScanTypeEnum string

const (
    CaptureScanTypeEnum_SCAN CaptureScanTypeEnum = "scan"
)

// CaptureSwitchFormatEnum is a string enum.
// enum: `stream`
type CaptureSwitchFormatEnum string

const (
    CaptureSwitchFormatEnum_STREAM CaptureSwitchFormatEnum = "stream"
)

// CaptureSwitchTypeEnum is a string enum.
// enum: `switch`
type CaptureSwitchTypeEnum string

const (
    CaptureSwitchTypeEnum_ENUMSWITCH CaptureSwitchTypeEnum = "switch"
)

// CaptureWiredFormatEnum is a string enum.
// pcap format. enum: `pcap`, `stream`
type CaptureWiredFormatEnum string

const (
    CaptureWiredFormatEnum_PCAP   CaptureWiredFormatEnum = "pcap"
    CaptureWiredFormatEnum_STREAM CaptureWiredFormatEnum = "stream"
)

// CaptureWiredTypeEnum is a string enum.
// enum: `wired`
type CaptureWiredTypeEnum string

const (
    CaptureWiredTypeEnum_WIRED CaptureWiredTypeEnum = "wired"
)

// CaptureWirelessBandEnum is a string enum.
// enum: `24`, `24,5,6`, `5`, `6`
type CaptureWirelessBandEnum string

const (
    CaptureWirelessBandEnum_ENUM24   CaptureWirelessBandEnum = "24"
    CaptureWirelessBandEnum_ENUM2456 CaptureWirelessBandEnum = "24,5,6"
    CaptureWirelessBandEnum_ENUM5    CaptureWirelessBandEnum = "5"
    CaptureWirelessBandEnum_ENUM6    CaptureWirelessBandEnum = "6"
)

// CaptureWirelessFormatEnum is a string enum.
// pcap format. enum: `pcap`, `stream`
type CaptureWirelessFormatEnum string

const (
    CaptureWirelessFormatEnum_PCAP   CaptureWirelessFormatEnum = "pcap"
    CaptureWirelessFormatEnum_STREAM CaptureWirelessFormatEnum = "stream"
)

// CaptureWirelessTypeEnum is a string enum.
// enum: `wireless`
type CaptureWirelessTypeEnum string

const (
    CaptureWirelessTypeEnum_WIRELESS CaptureWirelessTypeEnum = "wireless"
)

// ClaimTypeEnum is a string enum.
// what to claim. enum: `all`, `inventory`, `license`
type ClaimTypeEnum string

const (
    ClaimTypeEnum_ALL       ClaimTypeEnum = "all"
    ClaimTypeEnum_INVENTORY ClaimTypeEnum = "inventory"
    ClaimTypeEnum_LICENSE   ClaimTypeEnum = "license"
)

// ConfigSwitchLocalAccountsUserRoleEnum is a string enum.
// enum: `admin`, `helpdesk`, `none`, `read`
type ConfigSwitchLocalAccountsUserRoleEnum string

const (
    ConfigSwitchLocalAccountsUserRoleEnum_ADMIN    ConfigSwitchLocalAccountsUserRoleEnum = "admin"
    ConfigSwitchLocalAccountsUserRoleEnum_HELPDESK ConfigSwitchLocalAccountsUserRoleEnum = "helpdesk"
    ConfigSwitchLocalAccountsUserRoleEnum_NONE     ConfigSwitchLocalAccountsUserRoleEnum = "none"
    ConfigSwitchLocalAccountsUserRoleEnum_READ     ConfigSwitchLocalAccountsUserRoleEnum = "read"
)

// ConstDeviceApBand24UsageEnum is a string enum.
// enum: `24`, `5`, `6`
type ConstDeviceApBand24UsageEnum string

const (
    ConstDeviceApBand24UsageEnum_ENUM24 ConstDeviceApBand24UsageEnum = "24"
    ConstDeviceApBand24UsageEnum_ENUM5  ConstDeviceApBand24UsageEnum = "5"
    ConstDeviceApBand24UsageEnum_ENUM6  ConstDeviceApBand24UsageEnum = "6"
)

// ConstDeviceApExtiosDefaultDirEnum is a string enum.
// enum: `IN`, `OUT`
type ConstDeviceApExtiosDefaultDirEnum string

const (
    ConstDeviceApExtiosDefaultDirEnum_IN  ConstDeviceApExtiosDefaultDirEnum = "IN"
    ConstDeviceApExtiosDefaultDirEnum_OUT ConstDeviceApExtiosDefaultDirEnum = "OUT"
)

// ConstDeviceTypeApEnum is a string enum.
// Device Type. enum: `ap`
type ConstDeviceTypeApEnum string

const (
    ConstDeviceTypeApEnum_AP ConstDeviceTypeApEnum = "ap"
)

// ConstDeviceTypeGatewayEnum is a string enum.
// Device Type. enum: `gateway`
type ConstDeviceTypeGatewayEnum string

const (
    ConstDeviceTypeGatewayEnum_GATEWAY ConstDeviceTypeGatewayEnum = "gateway"
)

// ConstDeviceTypeSwitchEnum is a string enum.
// Device Type. enum: `switch`
type ConstDeviceTypeSwitchEnum string

const (
    ConstDeviceTypeSwitchEnum_ENUMSWITCH ConstDeviceTypeSwitchEnum = "switch"
)

// ConstInsightMetricsPropertyScopeEnum is a string enum.
// enum: `ap`, `client`, `device`, `mxedge`, `site`, `switch`
type ConstInsightMetricsPropertyScopeEnum string

const (
    ConstInsightMetricsPropertyScopeEnum_AP         ConstInsightMetricsPropertyScopeEnum = "ap"
    ConstInsightMetricsPropertyScopeEnum_CLIENT     ConstInsightMetricsPropertyScopeEnum = "client"
    ConstInsightMetricsPropertyScopeEnum_DEVICE     ConstInsightMetricsPropertyScopeEnum = "device"
    ConstInsightMetricsPropertyScopeEnum_MXEDGE     ConstInsightMetricsPropertyScopeEnum = "mxedge"
    ConstInsightMetricsPropertyScopeEnum_SITE       ConstInsightMetricsPropertyScopeEnum = "site"
    ConstInsightMetricsPropertyScopeEnum_ENUMSWITCH ConstInsightMetricsPropertyScopeEnum = "switch"
)

// CountOrgDevicesMxtunnelStatusEnum is a string enum.
// enum: `down`, `up`
type CountOrgDevicesMxtunnelStatusEnum string

const (
    CountOrgDevicesMxtunnelStatusEnum_DOWN CountOrgDevicesMxtunnelStatusEnum = "down"
    CountOrgDevicesMxtunnelStatusEnum_UP   CountOrgDevicesMxtunnelStatusEnum = "up"
)

// CountPortsAuthStateEnum is a string enum.
// enum: `authenticated`, `authenticating`, `held`, `init`
type CountPortsAuthStateEnum string

const (
    CountPortsAuthStateEnum_AUTHENTICATED  CountPortsAuthStateEnum = "authenticated"
    CountPortsAuthStateEnum_AUTHENTICATING CountPortsAuthStateEnum = "authenticating"
    CountPortsAuthStateEnum_HELD           CountPortsAuthStateEnum = "held"
    CountPortsAuthStateEnum_INIT           CountPortsAuthStateEnum = "init"
)

// CountPortsStpRoleEnum is a string enum.
// enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
type CountPortsStpRoleEnum string

const (
    CountPortsStpRoleEnum_ALTERNATE     CountPortsStpRoleEnum = "alternate"
    CountPortsStpRoleEnum_BACKUP        CountPortsStpRoleEnum = "backup"
    CountPortsStpRoleEnum_DESIGNATED    CountPortsStpRoleEnum = "designated"
    CountPortsStpRoleEnum_ROOT          CountPortsStpRoleEnum = "root"
    CountPortsStpRoleEnum_ROOTPREVENTED CountPortsStpRoleEnum = "root-prevented"
)

// CountPortsStpStateEnum is a string enum.
// enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
type CountPortsStpStateEnum string

const (
    CountPortsStpStateEnum_BLOCKING   CountPortsStpStateEnum = "blocking"
    CountPortsStpStateEnum_DISABLED   CountPortsStpStateEnum = "disabled"
    CountPortsStpStateEnum_FORWARDING CountPortsStpStateEnum = "forwarding"
    CountPortsStpStateEnum_LEARNING   CountPortsStpStateEnum = "learning"
    CountPortsStpStateEnum_LISTENING  CountPortsStpStateEnum = "listening"
)

// CountSiteCallsDistrinctEnum is a string enum.
// enum: `mac`
type CountSiteCallsDistrinctEnum string

const (
    CountSiteCallsDistrinctEnum_MAC CountSiteCallsDistrinctEnum = "mac"
)

// DayOfWeekEnum is a string enum.
// enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
type DayOfWeekEnum string

const (
    DayOfWeekEnum_ANY DayOfWeekEnum = "any"
    DayOfWeekEnum_FRI DayOfWeekEnum = "fri"
    DayOfWeekEnum_MON DayOfWeekEnum = "mon"
    DayOfWeekEnum_SAT DayOfWeekEnum = "sat"
    DayOfWeekEnum_SUN DayOfWeekEnum = "sun"
    DayOfWeekEnum_THU DayOfWeekEnum = "thu"
    DayOfWeekEnum_TUE DayOfWeekEnum = "tue"
    DayOfWeekEnum_WED DayOfWeekEnum = "wed"
)

// DeviceStatusEnum is a string enum.
// enum: `all`, `connected`, `disconnected`
type DeviceStatusEnum string

const (
    DeviceStatusEnum_ALL          DeviceStatusEnum = "all"
    DeviceStatusEnum_CONNECTED    DeviceStatusEnum = "connected"
    DeviceStatusEnum_DISCONNECTED DeviceStatusEnum = "disconnected"
)

// DeviceTypeEnum is a string enum.
// enum: `ap`, `gateway`, `switch`
type DeviceTypeEnum string

const (
    DeviceTypeEnum_AP         DeviceTypeEnum = "ap"
    DeviceTypeEnum_GATEWAY    DeviceTypeEnum = "gateway"
    DeviceTypeEnum_ENUMSWITCH DeviceTypeEnum = "switch"
)

// DeviceTypeApEnum is a string enum.
// Device Type. enum: `ap`
type DeviceTypeApEnum string

const (
    DeviceTypeApEnum_AP DeviceTypeApEnum = "ap"
)

// DeviceTypeGatewayEnum is a string enum.
// Device Type. enum: `gateway`
type DeviceTypeGatewayEnum string

const (
    DeviceTypeGatewayEnum_GATEWAY DeviceTypeGatewayEnum = "gateway"
)

// DeviceTypeSwitchEnum is a string enum.
// Device Type. enum: `switch`
type DeviceTypeSwitchEnum string

const (
    DeviceTypeSwitchEnum_ENUMSWITCH DeviceTypeSwitchEnum = "switch"
)

// DeviceTypeWithAllEnum is a string enum.
// enum: `all`, `ap`, `gateway`, `switch`
type DeviceTypeWithAllEnum string

const (
    DeviceTypeWithAllEnum_ALL        DeviceTypeWithAllEnum = "all"
    DeviceTypeWithAllEnum_AP         DeviceTypeWithAllEnum = "ap"
    DeviceTypeWithAllEnum_GATEWAY    DeviceTypeWithAllEnum = "gateway"
    DeviceTypeWithAllEnum_ENUMSWITCH DeviceTypeWithAllEnum = "switch"
)

// DeviceUpgradeRrmMeshUpgradeEnum is a string enum.
// For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`
type DeviceUpgradeRrmMeshUpgradeEnum string

const (
    DeviceUpgradeRrmMeshUpgradeEnum_PARALLEL   DeviceUpgradeRrmMeshUpgradeEnum = "parallel"
    DeviceUpgradeRrmMeshUpgradeEnum_SEQUENTIAL DeviceUpgradeRrmMeshUpgradeEnum = "sequential"
)

// DeviceUpgradeRrmNodeOrderEnum is a string enum.
// For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`
type DeviceUpgradeRrmNodeOrderEnum string

const (
    DeviceUpgradeRrmNodeOrderEnum_CENTERTOFRINGE DeviceUpgradeRrmNodeOrderEnum = "center_to_fringe"
    DeviceUpgradeRrmNodeOrderEnum_FRINGETOCENTER DeviceUpgradeRrmNodeOrderEnum = "fringe_to_center"
)

// DeviceUpgradeStatusEnum is a string enum.
// status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`
type DeviceUpgradeStatusEnum string

const (
    DeviceUpgradeStatusEnum_CANCELLED   DeviceUpgradeStatusEnum = "cancelled"
    DeviceUpgradeStatusEnum_COMPLETED   DeviceUpgradeStatusEnum = "completed"
    DeviceUpgradeStatusEnum_CREATED     DeviceUpgradeStatusEnum = "created"
    DeviceUpgradeStatusEnum_DOWNLOADED  DeviceUpgradeStatusEnum = "downloaded"
    DeviceUpgradeStatusEnum_DOWNLOADING DeviceUpgradeStatusEnum = "downloading"
    DeviceUpgradeStatusEnum_FAILED      DeviceUpgradeStatusEnum = "failed"
    DeviceUpgradeStatusEnum_UPGRADING   DeviceUpgradeStatusEnum = "upgrading"
)

// DeviceUpgradeStrategyEnum is a string enum.
// For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)
type DeviceUpgradeStrategyEnum string

const (
    DeviceUpgradeStrategyEnum_BIGBANG DeviceUpgradeStrategyEnum = "big_bang"
    DeviceUpgradeStrategyEnum_CANARY  DeviceUpgradeStrategyEnum = "canary"
    DeviceUpgradeStrategyEnum_RRM     DeviceUpgradeStrategyEnum = "rrm"
    DeviceUpgradeStrategyEnum_SERIAL  DeviceUpgradeStrategyEnum = "serial"
)

// DhcpdConfigOptionTypeEnum is a string enum.
// enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32`
type DhcpdConfigOptionTypeEnum string

const (
    DhcpdConfigOptionTypeEnum_BOOLEAN    DhcpdConfigOptionTypeEnum = "boolean"
    DhcpdConfigOptionTypeEnum_HEX        DhcpdConfigOptionTypeEnum = "hex"
    DhcpdConfigOptionTypeEnum_ENUMINT16  DhcpdConfigOptionTypeEnum = "int16"
    DhcpdConfigOptionTypeEnum_ENUMINT32  DhcpdConfigOptionTypeEnum = "int32"
    DhcpdConfigOptionTypeEnum_IP         DhcpdConfigOptionTypeEnum = "ip"
    DhcpdConfigOptionTypeEnum_ENUMSTRING DhcpdConfigOptionTypeEnum = "string"
    DhcpdConfigOptionTypeEnum_ENUMUINT16 DhcpdConfigOptionTypeEnum = "uint16"
    DhcpdConfigOptionTypeEnum_ENUMUINT32 DhcpdConfigOptionTypeEnum = "uint32"
)

// DhcpdConfigTypeEnum is a string enum.
// enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)
type DhcpdConfigTypeEnum string

const (
    DhcpdConfigTypeEnum_LOCAL DhcpdConfigTypeEnum = "local"
    DhcpdConfigTypeEnum_NONE  DhcpdConfigTypeEnum = "none"
    DhcpdConfigTypeEnum_RELAY DhcpdConfigTypeEnum = "relay"
)

// DhcpdConfigVendorOptionTypeEnum is a string enum.
// enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32`
type DhcpdConfigVendorOptionTypeEnum string

const (
    DhcpdConfigVendorOptionTypeEnum_BOOLEAN    DhcpdConfigVendorOptionTypeEnum = "boolean"
    DhcpdConfigVendorOptionTypeEnum_HEX        DhcpdConfigVendorOptionTypeEnum = "hex"
    DhcpdConfigVendorOptionTypeEnum_ENUMINT16  DhcpdConfigVendorOptionTypeEnum = "int16"
    DhcpdConfigVendorOptionTypeEnum_ENUMINT32  DhcpdConfigVendorOptionTypeEnum = "int32"
    DhcpdConfigVendorOptionTypeEnum_IP         DhcpdConfigVendorOptionTypeEnum = "ip"
    DhcpdConfigVendorOptionTypeEnum_ENUMSTRING DhcpdConfigVendorOptionTypeEnum = "string"
    DhcpdConfigVendorOptionTypeEnum_ENUMUINT16 DhcpdConfigVendorOptionTypeEnum = "uint16"
    DhcpdConfigVendorOptionTypeEnum_ENUMUINT32 DhcpdConfigVendorOptionTypeEnum = "uint32"
)

// DiscoveredSwitchMetricTypeEnum is a string enum.
// enum: `inactive_wired_vlans`, `poe_compliance`, `switch_ap_affinity`, `version_compliance`
type DiscoveredSwitchMetricTypeEnum string

const (
    DiscoveredSwitchMetricTypeEnum_INACTIVEWIREDVLANS DiscoveredSwitchMetricTypeEnum = "inactive_wired_vlans"
    DiscoveredSwitchMetricTypeEnum_POECOMPLIANCE      DiscoveredSwitchMetricTypeEnum = "poe_compliance"
    DiscoveredSwitchMetricTypeEnum_SWITCHAPAFFINITY   DiscoveredSwitchMetricTypeEnum = "switch_ap_affinity"
    DiscoveredSwitchMetricTypeEnum_VERSIONCOMPLIANCE  DiscoveredSwitchMetricTypeEnum = "version_compliance"
)

// DiscoveredSwitchesMetricScopeEnum is a string enum.
// enum: `site`, `switch`
type DiscoveredSwitchesMetricScopeEnum string

const (
    DiscoveredSwitchesMetricScopeEnum_SITE       DiscoveredSwitchesMetricScopeEnum = "site"
    DiscoveredSwitchesMetricScopeEnum_ENUMSWITCH DiscoveredSwitchesMetricScopeEnum = "switch"
)

// Dot11BandEnum is a string enum.
// enum: `24`, `5`, `6`
type Dot11BandEnum string

const (
    Dot11BandEnum_ENUM24 Dot11BandEnum = "24"
    Dot11BandEnum_ENUM5  Dot11BandEnum = "5"
    Dot11BandEnum_ENUM6  Dot11BandEnum = "6"
)

// Dot11BandwidthEnum is a int enum.
// channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
type Dot11BandwidthEnum int

const (
    Dot11BandwidthEnum_ENUM20  Dot11BandwidthEnum = 20
    Dot11BandwidthEnum_ENUM40  Dot11BandwidthEnum = 40
    Dot11BandwidthEnum_ENUM80  Dot11BandwidthEnum = 80
    Dot11BandwidthEnum_ENUM160 Dot11BandwidthEnum = 160
)

// Dot11Bandwidth24Enum is a int enum.
// channel width for the 2.4GHz band. enum: `20`, `40`
type Dot11Bandwidth24Enum int

const (
    Dot11Bandwidth24Enum_ENUM20 Dot11Bandwidth24Enum = 20
    Dot11Bandwidth24Enum_ENUM40 Dot11Bandwidth24Enum = 40
)

// Dot11Bandwidth5Enum is a int enum.
// channel width for the 5GHz band. enum: `20`, `40`, `80`
type Dot11Bandwidth5Enum int

const (
    Dot11Bandwidth5Enum_ENUM20 Dot11Bandwidth5Enum = 20
    Dot11Bandwidth5Enum_ENUM40 Dot11Bandwidth5Enum = 40
    Dot11Bandwidth5Enum_ENUM80 Dot11Bandwidth5Enum = 80
)

// Dot11Bandwidth6Enum is a int enum.
// channel width for the 6GHz band. enum: `20`, `40`, `80`, `160`
type Dot11Bandwidth6Enum int

const (
    Dot11Bandwidth6Enum_ENUM20  Dot11Bandwidth6Enum = 20
    Dot11Bandwidth6Enum_ENUM40  Dot11Bandwidth6Enum = 40
    Dot11Bandwidth6Enum_ENUM80  Dot11Bandwidth6Enum = 80
    Dot11Bandwidth6Enum_ENUM160 Dot11Bandwidth6Enum = 160
)

// Dot11ProtoEnum is a string enum.
// enum: `a`, `ac`, `ax`, `b`, `g`, `n`
type Dot11ProtoEnum string

const (
    Dot11ProtoEnum_A  Dot11ProtoEnum = "a"
    Dot11ProtoEnum_AC Dot11ProtoEnum = "ac"
    Dot11ProtoEnum_AX Dot11ProtoEnum = "ax"
    Dot11ProtoEnum_B  Dot11ProtoEnum = "b"
    Dot11ProtoEnum_G  Dot11ProtoEnum = "g"
    Dot11ProtoEnum_N  Dot11ProtoEnum = "n"
)

// DynamicPskSourceEnum is a string enum.
// enum: `cloud_psks`, `radius`
type DynamicPskSourceEnum string

const (
    DynamicPskSourceEnum_CLOUDPSKS DynamicPskSourceEnum = "cloud_psks"
    DynamicPskSourceEnum_RADIUS    DynamicPskSourceEnum = "radius"
)

// EventFastroamTypeEnum is a string enum.
// enum: `fail`, `none`, `pingpong`, `poor`, `slow`, `success`
type EventFastroamTypeEnum string

const (
    EventFastroamTypeEnum_FAIL     EventFastroamTypeEnum = "fail"
    EventFastroamTypeEnum_NONE     EventFastroamTypeEnum = "none"
    EventFastroamTypeEnum_PINGPONG EventFastroamTypeEnum = "pingpong"
    EventFastroamTypeEnum_POOR     EventFastroamTypeEnum = "poor"
    EventFastroamTypeEnum_SLOW     EventFastroamTypeEnum = "slow"
    EventFastroamTypeEnum_SUCCESS  EventFastroamTypeEnum = "success"
)

// EvpnConfigRoleEnum is a string enum.
// enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`
type EvpnConfigRoleEnum string

const (
    EvpnConfigRoleEnum_ACCESS        EvpnConfigRoleEnum = "access"
    EvpnConfigRoleEnum_COLLAPSEDCORE EvpnConfigRoleEnum = "collapsed-core"
    EvpnConfigRoleEnum_CORE          EvpnConfigRoleEnum = "core"
    EvpnConfigRoleEnum_DISTRIBUTION  EvpnConfigRoleEnum = "distribution"
    EvpnConfigRoleEnum_ESILAGACCESS  EvpnConfigRoleEnum = "esilag-access"
    EvpnConfigRoleEnum_NONE          EvpnConfigRoleEnum = "none"
)

// EvpnOptionsRoutedAtEnum is a string enum.
// optional, where virtual-gateway should reside. enum: `core`, `distribution`, `edge`
type EvpnOptionsRoutedAtEnum string

const (
    EvpnOptionsRoutedAtEnum_CORE         EvpnOptionsRoutedAtEnum = "core"
    EvpnOptionsRoutedAtEnum_DISTRIBUTION EvpnOptionsRoutedAtEnum = "distribution"
    EvpnOptionsRoutedAtEnum_EDGE         EvpnOptionsRoutedAtEnum = "edge"
)

// EvpnTopologySwitchRoleEnum is a string enum.
// use `role`==`none` to remove a switch from the topology. enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`
type EvpnTopologySwitchRoleEnum string

const (
    EvpnTopologySwitchRoleEnum_ACCESS        EvpnTopologySwitchRoleEnum = "access"
    EvpnTopologySwitchRoleEnum_COLLAPSEDCORE EvpnTopologySwitchRoleEnum = "collapsed-core"
    EvpnTopologySwitchRoleEnum_CORE          EvpnTopologySwitchRoleEnum = "core"
    EvpnTopologySwitchRoleEnum_DISTRIBUTION  EvpnTopologySwitchRoleEnum = "distribution"
    EvpnTopologySwitchRoleEnum_ESILAGACCESS  EvpnTopologySwitchRoleEnum = "esilag-access"
    EvpnTopologySwitchRoleEnum_NONE          EvpnTopologySwitchRoleEnum = "none"
)

// FastRoamResultEnum is a string enum.
// enum: `fail`, `none`, `success`
type FastRoamResultEnum string

const (
    FastRoamResultEnum_FAIL    FastRoamResultEnum = "fail"
    FastRoamResultEnum_NONE    FastRoamResultEnum = "none"
    FastRoamResultEnum_SUCCESS FastRoamResultEnum = "success"
)

// FwupdateStatStatusEnum is a string enum.
// enum: `inprogress`, `failed`, `upgraded`
type FwupdateStatStatusEnum string

const (
    FwupdateStatStatusEnum_INPROGRESS FwupdateStatStatusEnum = "inprogress"
    FwupdateStatStatusEnum_FAILED     FwupdateStatStatusEnum = "failed"
    FwupdateStatStatusEnum_UPGRADED   FwupdateStatStatusEnum = "upgraded"
)

// GatewayPathStrategyEnum is a string enum.
// enum: `ecmp`, `ordered`, `weighted`
type GatewayPathStrategyEnum string

const (
    GatewayPathStrategyEnum_ECMP     GatewayPathStrategyEnum = "ecmp"
    GatewayPathStrategyEnum_ORDERED  GatewayPathStrategyEnum = "ordered"
    GatewayPathStrategyEnum_WEIGHTED GatewayPathStrategyEnum = "weighted"
)

// GatewayPathTypeEnum is a string enum.
// enum: `local`, `tunnel`, `vpn`, `wan`
type GatewayPathTypeEnum string

const (
    GatewayPathTypeEnum_LOCAL  GatewayPathTypeEnum = "local"
    GatewayPathTypeEnum_TUNNEL GatewayPathTypeEnum = "tunnel"
    GatewayPathTypeEnum_VPN    GatewayPathTypeEnum = "vpn"
    GatewayPathTypeEnum_WAN    GatewayPathTypeEnum = "wan"
)

// GatewayPortDslTypeEnum is a string enum.
// if `wan_type`==`dsl`. enum: `adsl`, `vdsl`
type GatewayPortDslTypeEnum string

const (
    GatewayPortDslTypeEnum_ADSL GatewayPortDslTypeEnum = "adsl"
    GatewayPortDslTypeEnum_VDSL GatewayPortDslTypeEnum = "vdsl"
)

// GatewayPortDuplexEnum is a string enum.
// enum: `auto`, `full`, `half`
type GatewayPortDuplexEnum string

const (
    GatewayPortDuplexEnum_AUTO GatewayPortDuplexEnum = "auto"
    GatewayPortDuplexEnum_FULL GatewayPortDuplexEnum = "full"
    GatewayPortDuplexEnum_HALF GatewayPortDuplexEnum = "half"
)

// GatewayPortLteAuthEnum is a string enum.
// if `wan_type`==`lte`. enum: `chap`, `none`, `pap`
type GatewayPortLteAuthEnum string

const (
    GatewayPortLteAuthEnum_CHAP GatewayPortLteAuthEnum = "chap"
    GatewayPortLteAuthEnum_NONE GatewayPortLteAuthEnum = "none"
    GatewayPortLteAuthEnum_PAP  GatewayPortLteAuthEnum = "pap"
)

// GatewayPortUsageEnum is a string enum.
// port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`
type GatewayPortUsageEnum string

const (
    GatewayPortUsageEnum_HACONTROL GatewayPortUsageEnum = "ha_control"
    GatewayPortUsageEnum_HADATA    GatewayPortUsageEnum = "ha_data"
    GatewayPortUsageEnum_LAN       GatewayPortUsageEnum = "lan"
    GatewayPortUsageEnum_WAN       GatewayPortUsageEnum = "wan"
)

// GatewayPortVpnPathBfdProfileEnum is a string enum.
// Only if the VPN `type`==`hub_spoke`. enum: `broadband`, `lte`
type GatewayPortVpnPathBfdProfileEnum string

const (
    GatewayPortVpnPathBfdProfileEnum_BROADBAND GatewayPortVpnPathBfdProfileEnum = "broadband"
    GatewayPortVpnPathBfdProfileEnum_LTE       GatewayPortVpnPathBfdProfileEnum = "lte"
)

// GatewayPortVpnPathRoleEnum is a string enum.
// Only if the VPN `type`==`hub_spoke`. enum: `hub`, `spoke`
type GatewayPortVpnPathRoleEnum string

const (
    GatewayPortVpnPathRoleEnum_HUB   GatewayPortVpnPathRoleEnum = "hub"
    GatewayPortVpnPathRoleEnum_SPOKE GatewayPortVpnPathRoleEnum = "spoke"
)

// GatewayPortWanArpPolicerEnum is a string enum.
// Only when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`
type GatewayPortWanArpPolicerEnum string

const (
    GatewayPortWanArpPolicerEnum_ENUMDEFAULT GatewayPortWanArpPolicerEnum = "default"
    GatewayPortWanArpPolicerEnum_MAX         GatewayPortWanArpPolicerEnum = "max"
    GatewayPortWanArpPolicerEnum_RECOMMENDED GatewayPortWanArpPolicerEnum = "recommended"
)

// GatewayPortWanTypeEnum is a string enum.
// Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`
type GatewayPortWanTypeEnum string

const (
    GatewayPortWanTypeEnum_BROADBAND GatewayPortWanTypeEnum = "broadband"
    GatewayPortWanTypeEnum_DSL       GatewayPortWanTypeEnum = "dsl"
    GatewayPortWanTypeEnum_LTE       GatewayPortWanTypeEnum = "lte"
)

// GatewayTemplateTypeEnum is a string enum.
// enum: `spoke`, `standalone`
type GatewayTemplateTypeEnum string

const (
    GatewayTemplateTypeEnum_SPOKE      GatewayTemplateTypeEnum = "spoke"
    GatewayTemplateTypeEnum_STANDALONE GatewayTemplateTypeEnum = "standalone"
)

// GatewayWanPpoeAuthEnum is a string enum.
// if `type`==`pppoe`. enum: `chap`, `none`, `pap`
type GatewayWanPpoeAuthEnum string

const (
    GatewayWanPpoeAuthEnum_CHAP GatewayWanPpoeAuthEnum = "chap"
    GatewayWanPpoeAuthEnum_NONE GatewayWanPpoeAuthEnum = "none"
    GatewayWanPpoeAuthEnum_PAP  GatewayWanPpoeAuthEnum = "pap"
)

// GatewayWanProbeOverrideProbeProfileEnum is a string enum.
// enum: `broadband`, `lte`
type GatewayWanProbeOverrideProbeProfileEnum string

const (
    GatewayWanProbeOverrideProbeProfileEnum_BROADBAND GatewayWanProbeOverrideProbeProfileEnum = "broadband"
    GatewayWanProbeOverrideProbeProfileEnum_LTE       GatewayWanProbeOverrideProbeProfileEnum = "lte"
)

// GatewayWanTypeEnum is a string enum.
// enum: `dhcp`, `pppoe`, `static`
type GatewayWanTypeEnum string

const (
    GatewayWanTypeEnum_DHCP   GatewayWanTypeEnum = "dhcp"
    GatewayWanTypeEnum_PPPOE  GatewayWanTypeEnum = "pppoe"
    GatewayWanTypeEnum_STATIC GatewayWanTypeEnum = "static"
)

// GetOrgMxedgeUpgradeInfoChannelEnum is a string enum.
// enum: `alpha`, `beta`, `stable`
type GetOrgMxedgeUpgradeInfoChannelEnum string

const (
    GetOrgMxedgeUpgradeInfoChannelEnum_ALPHA  GetOrgMxedgeUpgradeInfoChannelEnum = "alpha"
    GetOrgMxedgeUpgradeInfoChannelEnum_BETA   GetOrgMxedgeUpgradeInfoChannelEnum = "beta"
    GetOrgMxedgeUpgradeInfoChannelEnum_STABLE GetOrgMxedgeUpgradeInfoChannelEnum = "stable"
)

// HaClusterNodeEnum is a string enum.
// only for HA. enum: `node0`, `node1`
type HaClusterNodeEnum string

const (
    HaClusterNodeEnum_NODE0 HaClusterNodeEnum = "node0"
    HaClusterNodeEnum_NODE1 HaClusterNodeEnum = "node1"
)

// IdpMachineCertLookupFieldEnum is a string enum.
// allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`
type IdpMachineCertLookupFieldEnum string

const (
    IdpMachineCertLookupFieldEnum_AUTOMATIC IdpMachineCertLookupFieldEnum = "automatic"
    IdpMachineCertLookupFieldEnum_CN        IdpMachineCertLookupFieldEnum = "cn"
    IdpMachineCertLookupFieldEnum_DNS       IdpMachineCertLookupFieldEnum = "dns"
)

// IdpProfileActionEnum is a string enum.
// enum:
// * alert (default)
// * drop: siliently dropping packets
// * close: notify client/server to close connection
type IdpProfileActionEnum string

const (
    IdpProfileActionEnum_ALERT IdpProfileActionEnum = "alert"
    IdpProfileActionEnum_CLOSE IdpProfileActionEnum = "close"
    IdpProfileActionEnum_DROP  IdpProfileActionEnum = "drop"
)

// IdpProfileBaseProfileEnum is a string enum.
// enum: `critical`, `standard`, `strict`
type IdpProfileBaseProfileEnum string

const (
    IdpProfileBaseProfileEnum_CRITICAL IdpProfileBaseProfileEnum = "critical"
    IdpProfileBaseProfileEnum_STANDARD IdpProfileBaseProfileEnum = "standard"
    IdpProfileBaseProfileEnum_STRICT   IdpProfileBaseProfileEnum = "strict"
)

// IdpProfileMatchingSeverityValueEnum is a string enum.
// enum: `critical`, `info`, `major`, `minor`
type IdpProfileMatchingSeverityValueEnum string

const (
    IdpProfileMatchingSeverityValueEnum_CRITICAL IdpProfileMatchingSeverityValueEnum = "critical"
    IdpProfileMatchingSeverityValueEnum_INFO     IdpProfileMatchingSeverityValueEnum = "info"
    IdpProfileMatchingSeverityValueEnum_MAJOR    IdpProfileMatchingSeverityValueEnum = "major"
    IdpProfileMatchingSeverityValueEnum_MINOR    IdpProfileMatchingSeverityValueEnum = "minor"
)

// IdpUserCertLookupFieldEnum is a string enum.
// allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`
type IdpUserCertLookupFieldEnum string

const (
    IdpUserCertLookupFieldEnum_AUTOMATIC IdpUserCertLookupFieldEnum = "automatic"
    IdpUserCertLookupFieldEnum_CN        IdpUserCertLookupFieldEnum = "cn"
    IdpUserCertLookupFieldEnum_EMAIL     IdpUserCertLookupFieldEnum = "email"
    IdpUserCertLookupFieldEnum_UPN       IdpUserCertLookupFieldEnum = "upn"
)

// ImportSiteAssetsUpsertEnum is a string enum.
// enum: `False`, `True`
type ImportSiteAssetsUpsertEnum string

const (
    ImportSiteAssetsUpsertEnum_FALSE ImportSiteAssetsUpsertEnum = "False"
    ImportSiteAssetsUpsertEnum_TRUE  ImportSiteAssetsUpsertEnum = "True"
)

// InventoryCountDistinctEnum is a string enum.
// enum: `model`, `status`, `site_id`, `sku`, `version`
type InventoryCountDistinctEnum string

const (
    InventoryCountDistinctEnum_MODEL   InventoryCountDistinctEnum = "model"
    InventoryCountDistinctEnum_STATUS  InventoryCountDistinctEnum = "status"
    InventoryCountDistinctEnum_SITEID  InventoryCountDistinctEnum = "site_id"
    InventoryCountDistinctEnum_SKU     InventoryCountDistinctEnum = "sku"
    InventoryCountDistinctEnum_VERSION InventoryCountDistinctEnum = "version"
)

// InventoryUpdateOperationEnum is a string enum.
// enum:
// * `upgrade_to_mist`: Upgrade to mist-managed
// * `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.
// * `assign`: Assign inventory to a site
// * `unassign`: Unassign inventory from a site
// * `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned
type InventoryUpdateOperationEnum string

const (
    InventoryUpdateOperationEnum_ASSIGN         InventoryUpdateOperationEnum = "assign"
    InventoryUpdateOperationEnum_DELETE         InventoryUpdateOperationEnum = "delete"
    InventoryUpdateOperationEnum_DOWNGRADETOJSI InventoryUpdateOperationEnum = "downgrade_to_jsi"
    InventoryUpdateOperationEnum_UNASSIGN       InventoryUpdateOperationEnum = "unassign"
    InventoryUpdateOperationEnum_UPGRADETOMIST  InventoryUpdateOperationEnum = "upgrade_to_mist"
)

// IpTypeEnum is a string enum.
// enum: `dhcp`, `static`
type IpTypeEnum string

const (
    IpTypeEnum_DHCP   IpTypeEnum = "dhcp"
    IpTypeEnum_STATIC IpTypeEnum = "static"
)

// IpType6Enum is a string enum.
// enum: `autoconf`, `dhcp`, `disabled`, `static`
type IpType6Enum string

const (
    IpType6Enum_AUTOCONF IpType6Enum = "autoconf"
    IpType6Enum_DHCP     IpType6Enum = "dhcp"
    IpType6Enum_DISABLED IpType6Enum = "disabled"
    IpType6Enum_STATIC   IpType6Enum = "static"
)

// JseInventoryItemTypeEnum is a string enum.
// enum: `ap`, `gateway`, `switch`
type JseInventoryItemTypeEnum string

const (
    JseInventoryItemTypeEnum_AP         JseInventoryItemTypeEnum = "ap"
    JseInventoryItemTypeEnum_GATEWAY    JseInventoryItemTypeEnum = "gateway"
    JseInventoryItemTypeEnum_ENUMSWITCH JseInventoryItemTypeEnum = "switch"
)

// JunosPortConfigDuplexEnum is a string enum.
// enum: `auto`, `full`, `half`
type JunosPortConfigDuplexEnum string

const (
    JunosPortConfigDuplexEnum_AUTO JunosPortConfigDuplexEnum = "auto"
    JunosPortConfigDuplexEnum_FULL JunosPortConfigDuplexEnum = "full"
    JunosPortConfigDuplexEnum_HALF JunosPortConfigDuplexEnum = "half"
)

// JunosPortConfigSpeedEnum is a string enum.
// enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
type JunosPortConfigSpeedEnum string

const (
    JunosPortConfigSpeedEnum_ENUM10M  JunosPortConfigSpeedEnum = "10m"
    JunosPortConfigSpeedEnum_ENUM100M JunosPortConfigSpeedEnum = "100m"
    JunosPortConfigSpeedEnum_ENUM1G   JunosPortConfigSpeedEnum = "1g"
    JunosPortConfigSpeedEnum_ENUM25G  JunosPortConfigSpeedEnum = "2.5g"
    JunosPortConfigSpeedEnum_ENUM5G   JunosPortConfigSpeedEnum = "5g"
    JunosPortConfigSpeedEnum_ENUM10G  JunosPortConfigSpeedEnum = "10g"
    JunosPortConfigSpeedEnum_ENUM25G1 JunosPortConfigSpeedEnum = "25g"
    JunosPortConfigSpeedEnum_ENUM40G  JunosPortConfigSpeedEnum = "40g"
    JunosPortConfigSpeedEnum_ENUM100G JunosPortConfigSpeedEnum = "100g"
    JunosPortConfigSpeedEnum_AUTO     JunosPortConfigSpeedEnum = "auto"
)

// L2tpStateEnum is a string enum.
// enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
type L2tpStateEnum string

const (
    L2tpStateEnum_ESTABLISHED            L2tpStateEnum = "established"
    L2tpStateEnum_ESTABLISHEDWITHSESSION L2tpStateEnum = "established_with_session"
    L2tpStateEnum_IDLE                   L2tpStateEnum = "idle"
    L2tpStateEnum_WAITCTRLCONN           L2tpStateEnum = "wait-ctrl-conn"
    L2tpStateEnum_WAITCTRLREPLY          L2tpStateEnum = "wait-ctrl-reply"
)

// LicenseTypeEnum is a string enum.
// enum: `SUB-AST`, `SUB-DATA`, `SUB-ENG`, `SUB-EX12`, `SUB-EX24`, `SUB-EX48`, `SUB-MAN`, `SUB-ME`, `SUB-PMA`, `SUB-SRX1`, `SUB-SRX2`, `SUB-SVNA`, `SUB-VNA`, `SUB-WAN1`, `SUB-WAN2`, `SUB-WVNA1`, `SUB-WVNA2`
type LicenseTypeEnum string

const (
    LicenseTypeEnum_SUBAST   LicenseTypeEnum = "SUB-AST"
    LicenseTypeEnum_SUBDATA  LicenseTypeEnum = "SUB-DATA"
    LicenseTypeEnum_SUBENG   LicenseTypeEnum = "SUB-ENG"
    LicenseTypeEnum_SUBEX12  LicenseTypeEnum = "SUB-EX12"
    LicenseTypeEnum_SUBEX24  LicenseTypeEnum = "SUB-EX24"
    LicenseTypeEnum_SUBEX48  LicenseTypeEnum = "SUB-EX48"
    LicenseTypeEnum_SUBMAN   LicenseTypeEnum = "SUB-MAN"
    LicenseTypeEnum_SUBME    LicenseTypeEnum = "SUB-ME"
    LicenseTypeEnum_SUBPMA   LicenseTypeEnum = "SUB-PMA"
    LicenseTypeEnum_SUBSRX1  LicenseTypeEnum = "SUB-SRX1"
    LicenseTypeEnum_SUBSRX2  LicenseTypeEnum = "SUB-SRX2"
    LicenseTypeEnum_SUBSVNA  LicenseTypeEnum = "SUB-SVNA"
    LicenseTypeEnum_SUBVNA   LicenseTypeEnum = "SUB-VNA"
    LicenseTypeEnum_SUBWAN1  LicenseTypeEnum = "SUB-WAN1"
    LicenseTypeEnum_SUBWAN2  LicenseTypeEnum = "SUB-WAN2"
    LicenseTypeEnum_SUBWVNA1 LicenseTypeEnum = "SUB-WVNA1"
    LicenseTypeEnum_SUBWVNA2 LicenseTypeEnum = "SUB-WVNA2"
)

// ListMspLogsSortEnum is a string enum.
// enum: `-timestamp`, `admin_id`, `site_id`, `timestamp`
type ListMspLogsSortEnum string

const (
    ListMspLogsSortEnum_ENUMTIMESTAMP ListMspLogsSortEnum = "-timestamp"
    ListMspLogsSortEnum_ADMINID       ListMspLogsSortEnum = "admin_id"
    ListMspLogsSortEnum_SITEID        ListMspLogsSortEnum = "site_id"
    ListMspLogsSortEnum_TIMESTAMP     ListMspLogsSortEnum = "timestamp"
)

// ListOrgLogsSortEnum is a string enum.
// enum: `-timestamp`, `admin_id`, `site_id`, `timestamp`
type ListOrgLogsSortEnum string

const (
    ListOrgLogsSortEnum_ENUMTIMESTAMP ListOrgLogsSortEnum = "-timestamp"
    ListOrgLogsSortEnum_ADMINID       ListOrgLogsSortEnum = "admin_id"
    ListOrgLogsSortEnum_SITEID        ListOrgLogsSortEnum = "site_id"
    ListOrgLogsSortEnum_TIMESTAMP     ListOrgLogsSortEnum = "timestamp"
)

// MapImportJsonVendorNameEnum is a string enum.
// enum: `ekahau`, `ibwave`
type MapImportJsonVendorNameEnum string

const (
    MapImportJsonVendorNameEnum_EKAHAU MapImportJsonVendorNameEnum = "ekahau"
    MapImportJsonVendorNameEnum_IBWAVE MapImportJsonVendorNameEnum = "ibwave"
)

// MapJibestreamVendorNameEnum is a string enum.
// the vendor ‘jibestream’. enum: `jibestream`
type MapJibestreamVendorNameEnum string

const (
    MapJibestreamVendorNameEnum_JIBESTREAM MapJibestreamVendorNameEnum = "jibestream"
)

// MapMicelloVendorNameEnum is a string enum.
// the vendor ‘micello’. enum: `micello`
type MapMicelloVendorNameEnum string

const (
    MapMicelloVendorNameEnum_MICELLO MapMicelloVendorNameEnum = "micello"
)

// MapOrgImportFileJsonVendorNameEnum is a string enum.
// enum: `ekahau`, `ibwave`
type MapOrgImportFileJsonVendorNameEnum string

const (
    MapOrgImportFileJsonVendorNameEnum_EKAHAU MapOrgImportFileJsonVendorNameEnum = "ekahau"
    MapOrgImportFileJsonVendorNameEnum_IBWAVE MapOrgImportFileJsonVendorNameEnum = "ibwave"
)

// MapTypeEnum is a string enum.
// enum: `google`, `image`
type MapTypeEnum string

const (
    MapTypeEnum_GOOGLE MapTypeEnum = "google"
    MapTypeEnum_IMAGE  MapTypeEnum = "image"
)

// MapViewEnum is a string enum.
// if `type`==`google`. enum: `hybrid`, `roadmap`, `satellite`, `terrain`
type MapViewEnum string

const (
    MapViewEnum_HYBRID    MapViewEnum = "hybrid"
    MapViewEnum_ROADMAP   MapViewEnum = "roadmap"
    MapViewEnum_SATELLITE MapViewEnum = "satellite"
    MapViewEnum_TERRAIN   MapViewEnum = "terrain"
)

// MfaSecretTypeEnum is a string enum.
// enum: `qrcode`
type MfaSecretTypeEnum string

const (
    MfaSecretTypeEnum_QRCODE MfaSecretTypeEnum = "qrcode"
)

// MspLicenseActionOperationEnum is a string enum.
// enum: `amend`, `annotate`, `delete`, `unamend`
type MspLicenseActionOperationEnum string

const (
    MspLicenseActionOperationEnum_AMEND    MspLicenseActionOperationEnum = "amend"
    MspLicenseActionOperationEnum_ANNOTATE MspLicenseActionOperationEnum = "annotate"
    MspLicenseActionOperationEnum_DELETE   MspLicenseActionOperationEnum = "delete"
    MspLicenseActionOperationEnum_UNAMEND  MspLicenseActionOperationEnum = "unamend"
)

// MspLogsCountDistinctEnum is a string enum.
// enum: `admin_id`, `admin_name`, `message`, `org_id`
type MspLogsCountDistinctEnum string

const (
    MspLogsCountDistinctEnum_ADMINID   MspLogsCountDistinctEnum = "admin_id"
    MspLogsCountDistinctEnum_ADMINNAME MspLogsCountDistinctEnum = "admin_name"
    MspLogsCountDistinctEnum_MESSAGE   MspLogsCountDistinctEnum = "message"
    MspLogsCountDistinctEnum_ORGID     MspLogsCountDistinctEnum = "org_id"
)

// MspMarvisSuggestionsCountDistinctEnum is a string enum.
// enum: `org_id`, `status`
type MspMarvisSuggestionsCountDistinctEnum string

const (
    MspMarvisSuggestionsCountDistinctEnum_ORGID  MspMarvisSuggestionsCountDistinctEnum = "org_id"
    MspMarvisSuggestionsCountDistinctEnum_STATUS MspMarvisSuggestionsCountDistinctEnum = "status"
)

// MspOrgChangeOperationEnum is a string enum.
// enum: `assign`, `unassign`
type MspOrgChangeOperationEnum string

const (
    MspOrgChangeOperationEnum_ASSIGN   MspOrgChangeOperationEnum = "assign"
    MspOrgChangeOperationEnum_UNASSIGN MspOrgChangeOperationEnum = "unassign"
)

// MspSearchTypeEnum is a string enum.
// enum: `orgs`
type MspSearchTypeEnum string

const (
    MspSearchTypeEnum_ORGS MspSearchTypeEnum = "orgs"
)

// MspTicketsCountDistinctEnum is a string enum.
// enum: `org_id`, `status`, `type`
type MspTicketsCountDistinctEnum string

const (
    MspTicketsCountDistinctEnum_ORGID    MspTicketsCountDistinctEnum = "org_id"
    MspTicketsCountDistinctEnum_STATUS   MspTicketsCountDistinctEnum = "status"
    MspTicketsCountDistinctEnum_ENUMTYPE MspTicketsCountDistinctEnum = "type"
)

// MspTierEnum is a string enum.
// enum: `advanced`, `base`
type MspTierEnum string

const (
    MspTierEnum_ADVANCED MspTierEnum = "advanced"
    MspTierEnum_BASE     MspTierEnum = "base"
)

// MxclusterNacClientVendorEnum is a string enum.
// convention to be followed is : "<vendor>-<variant>", <variant> could be an os/platform/model/company. For ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc. enum: `aruba`, `cisco-aironet`, `cisco-ios`, `cisco-meraki`, `generic`, `juniper`, `paloalto`
type MxclusterNacClientVendorEnum string

const (
    MxclusterNacClientVendorEnum_ARUBA        MxclusterNacClientVendorEnum = "aruba"
    MxclusterNacClientVendorEnum_CISCOAIRONET MxclusterNacClientVendorEnum = "cisco-aironet"
    MxclusterNacClientVendorEnum_CISCOIOS     MxclusterNacClientVendorEnum = "cisco-ios"
    MxclusterNacClientVendorEnum_CISCOMERAKI  MxclusterNacClientVendorEnum = "cisco-meraki"
    MxclusterNacClientVendorEnum_GENERIC      MxclusterNacClientVendorEnum = "generic"
    MxclusterNacClientVendorEnum_JUNIPER      MxclusterNacClientVendorEnum = "juniper"
    MxclusterNacClientVendorEnum_PALOALTO     MxclusterNacClientVendorEnum = "paloalto"
)

// MxclusterRadAuthServerKeywrapFormatEnum is a string enum.
// if used for Mist APs. enum: `ascii`, `hex`
type MxclusterRadAuthServerKeywrapFormatEnum string

const (
    MxclusterRadAuthServerKeywrapFormatEnum_ASCII MxclusterRadAuthServerKeywrapFormatEnum = "ascii"
    MxclusterRadAuthServerKeywrapFormatEnum_HEX   MxclusterRadAuthServerKeywrapFormatEnum = "hex"
)

// MxclusterRadsecNasIpSourceEnum is a string enum.
// SSpecify NAS-IP-ADDRESS, NAS-IPv6-ADDRESS to use with auth_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`
type MxclusterRadsecNasIpSourceEnum string

const (
    MxclusterRadsecNasIpSourceEnum_ANY     MxclusterRadsecNasIpSourceEnum = "any"
    MxclusterRadsecNasIpSourceEnum_OOB     MxclusterRadsecNasIpSourceEnum = "oob"
    MxclusterRadsecNasIpSourceEnum_OOB6    MxclusterRadsecNasIpSourceEnum = "oob6"
    MxclusterRadsecNasIpSourceEnum_TUNNEL  MxclusterRadsecNasIpSourceEnum = "tunnel"
    MxclusterRadsecNasIpSourceEnum_TUNNEL6 MxclusterRadsecNasIpSourceEnum = "tunnel6"
)

// MxclusterRadsecServerSelectionEnum is a string enum.
// When ordered, Mist Edge will prefer and go back to the first radius server if possible. enum: `ordered`, `unordered`
type MxclusterRadsecServerSelectionEnum string

const (
    MxclusterRadsecServerSelectionEnum_ORDERED   MxclusterRadsecServerSelectionEnum = "ordered"
    MxclusterRadsecServerSelectionEnum_UNORDERED MxclusterRadsecServerSelectionEnum = "unordered"
)

// MxclusterRadsecSrcIpSourceEnum is a string enum.
// Specify IP address to connect to auth_servers and acct_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`
type MxclusterRadsecSrcIpSourceEnum string

const (
    MxclusterRadsecSrcIpSourceEnum_ANY     MxclusterRadsecSrcIpSourceEnum = "any"
    MxclusterRadsecSrcIpSourceEnum_OOB     MxclusterRadsecSrcIpSourceEnum = "oob"
    MxclusterRadsecSrcIpSourceEnum_OOB6    MxclusterRadsecSrcIpSourceEnum = "oob6"
    MxclusterRadsecSrcIpSourceEnum_TUNNEL  MxclusterRadsecSrcIpSourceEnum = "tunnel"
    MxclusterRadsecSrcIpSourceEnum_TUNNEL6 MxclusterRadsecSrcIpSourceEnum = "tunnel6"
)

// MxclusterTuntermHostsSelectionEnum is a string enum.
// Ordering of tunterm_hosts for mxedge within the same mxcluster. enum:
// * `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC
// * `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)
// * `ordered`: order decided by tunterm_hosts_order
type MxclusterTuntermHostsSelectionEnum string

const (
    MxclusterTuntermHostsSelectionEnum_ORDERED       MxclusterTuntermHostsSelectionEnum = "ordered"
    MxclusterTuntermHostsSelectionEnum_SHUFFLE       MxclusterTuntermHostsSelectionEnum = "shuffle"
    MxclusterTuntermHostsSelectionEnum_SHUFFLEBYSITE MxclusterTuntermHostsSelectionEnum = "shuffle-by-site"
)

// MxedgeForSiteEnum is a string enum.
// enum: `any`, `false`, `true`
type MxedgeForSiteEnum string

const (
    MxedgeForSiteEnum_ANY   MxedgeForSiteEnum = "any"
    MxedgeForSiteEnum_FALSE MxedgeForSiteEnum = "false"
    MxedgeForSiteEnum_TRUE  MxedgeForSiteEnum = "true"
)

// MxedgeMgmtOobIpTypeEnum is a string enum.
// enum: `dhcp`, `disabled`, `static`
type MxedgeMgmtOobIpTypeEnum string

const (
    MxedgeMgmtOobIpTypeEnum_DHCP     MxedgeMgmtOobIpTypeEnum = "dhcp"
    MxedgeMgmtOobIpTypeEnum_DISABLED MxedgeMgmtOobIpTypeEnum = "disabled"
    MxedgeMgmtOobIpTypeEnum_STATIC   MxedgeMgmtOobIpTypeEnum = "static"
)

// MxedgeMgmtOobIpType6Enum is a string enum.
// enum: `autoconf`, `dhcp`, `disabled`, `static`
type MxedgeMgmtOobIpType6Enum string

const (
    MxedgeMgmtOobIpType6Enum_AUTOCONF MxedgeMgmtOobIpType6Enum = "autoconf"
    MxedgeMgmtOobIpType6Enum_DHCP     MxedgeMgmtOobIpType6Enum = "dhcp"
    MxedgeMgmtOobIpType6Enum_DISABLED MxedgeMgmtOobIpType6Enum = "disabled"
    MxedgeMgmtOobIpType6Enum_STATIC   MxedgeMgmtOobIpType6Enum = "static"
)

// MxedgeServiceActionEnum is a string enum.
// enum: `restart`, `start`, `stop`
type MxedgeServiceActionEnum string

const (
    MxedgeServiceActionEnum_RESTART MxedgeServiceActionEnum = "restart"
    MxedgeServiceActionEnum_START   MxedgeServiceActionEnum = "start"
    MxedgeServiceActionEnum_STOP    MxedgeServiceActionEnum = "stop"
)

// MxedgeServiceNameEnum is a string enum.
// enum: `mxagent`, `mxdas`, `mxocproxy`, `radsecproxy`, `tunterm`
type MxedgeServiceNameEnum string

const (
    MxedgeServiceNameEnum_MXAGENT     MxedgeServiceNameEnum = "mxagent"
    MxedgeServiceNameEnum_MXDAS       MxedgeServiceNameEnum = "mxdas"
    MxedgeServiceNameEnum_MXNACEDGE   MxedgeServiceNameEnum = "mxnacedge"
    MxedgeServiceNameEnum_MXOCPROXY   MxedgeServiceNameEnum = "mxocproxy"
    MxedgeServiceNameEnum_RADSECPROXY MxedgeServiceNameEnum = "radsecproxy"
    MxedgeServiceNameEnum_TUNTERM     MxedgeServiceNameEnum = "tunterm"
)

// MxedgeTuntermDhcpdConfigTypeEnum is a string enum.
// enum: `relay`
type MxedgeTuntermDhcpdConfigTypeEnum string

const (
    MxedgeTuntermDhcpdConfigTypeEnum_RELAY MxedgeTuntermDhcpdConfigTypeEnum = "relay"
)

// MxedgeTuntermDhcpdTypeEnum is a string enum.
// enum: `relay`
type MxedgeTuntermDhcpdTypeEnum string

const (
    MxedgeTuntermDhcpdTypeEnum_RELAY MxedgeTuntermDhcpdTypeEnum = "relay"
)

// MxedgeUpgradeChannelEnum is a string enum.
// upgrade channel to follow. enum: `alpha`, `beta`, `stable`
type MxedgeUpgradeChannelEnum string

const (
    MxedgeUpgradeChannelEnum_ALPHA  MxedgeUpgradeChannelEnum = "alpha"
    MxedgeUpgradeChannelEnum_BETA   MxedgeUpgradeChannelEnum = "beta"
    MxedgeUpgradeChannelEnum_STABLE MxedgeUpgradeChannelEnum = "stable"
)

// MxedgeUpgradeStrategyEnum is a string enum.
// enum:
// * `big_bang`: upgrade all at once, no orchestration
// * `serial`: one at a time'
// * `canary`: upgrade in phases
type MxedgeUpgradeStrategyEnum string

const (
    MxedgeUpgradeStrategyEnum_CANARY  MxedgeUpgradeStrategyEnum = "canary"
    MxedgeUpgradeStrategyEnum_BIGBANG MxedgeUpgradeStrategyEnum = "big_bang"
    MxedgeUpgradeStrategyEnum_SERIAL  MxedgeUpgradeStrategyEnum = "serial"
)

// MxtunnelProtocolEnum is a string enum.
// enum: `ip`, `udp`
type MxtunnelProtocolEnum string

const (
    MxtunnelProtocolEnum_IP  MxtunnelProtocolEnum = "ip"
    MxtunnelProtocolEnum_UDP MxtunnelProtocolEnum = "udp"
)

// NacPortalAccessTypeEnum is a string enum.
// if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`
type NacPortalAccessTypeEnum string

const (
    NacPortalAccessTypeEnum_WIRELESS          NacPortalAccessTypeEnum = "wireless"
    NacPortalAccessTypeEnum_ENUMWIRELESSWIRED NacPortalAccessTypeEnum = "wireless+wired"
)

// NacPortalEapTypeEnum is a string enum.
// enum: `wpa2`, `wpa3`
type NacPortalEapTypeEnum string

const (
    NacPortalEapTypeEnum_WPA2 NacPortalEapTypeEnum = "wpa2"
    NacPortalEapTypeEnum_WPA3 NacPortalEapTypeEnum = "wpa3"
)

// NacPortalSsoIdpSignAlgoEnum is a string enum.
// Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`.
type NacPortalSsoIdpSignAlgoEnum string

const (
    NacPortalSsoIdpSignAlgoEnum_SHA1   NacPortalSsoIdpSignAlgoEnum = "sha1"
    NacPortalSsoIdpSignAlgoEnum_SHA256 NacPortalSsoIdpSignAlgoEnum = "sha256"
    NacPortalSsoIdpSignAlgoEnum_SHA384 NacPortalSsoIdpSignAlgoEnum = "sha384"
    NacPortalSsoIdpSignAlgoEnum_SHA512 NacPortalSsoIdpSignAlgoEnum = "sha512"
)

// NacPortalTypeEnum is a string enum.
// enum: `guest`, `marvis_client`
type NacPortalTypeEnum string

const (
    NacPortalTypeEnum_GUEST        NacPortalTypeEnum = "guest"
    NacPortalTypeEnum_MARVISCLIENT NacPortalTypeEnum = "marvis_client"
)

// NacRuleActionEnum is a string enum.
// enum: `allow`, `block`
type NacRuleActionEnum string

const (
    NacRuleActionEnum_ALLOW NacRuleActionEnum = "allow"
    NacRuleActionEnum_BLOCK NacRuleActionEnum = "block"
)

// NacRuleMatchingAuthTypeEnum is a string enum.
// enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk`
type NacRuleMatchingAuthTypeEnum string

const (
    NacRuleMatchingAuthTypeEnum_CERT       NacRuleMatchingAuthTypeEnum = "cert"
    NacRuleMatchingAuthTypeEnum_DEVICEAUTH NacRuleMatchingAuthTypeEnum = "device-auth"
    NacRuleMatchingAuthTypeEnum_EAPTEAP    NacRuleMatchingAuthTypeEnum = "eap-teap"
    NacRuleMatchingAuthTypeEnum_EAPTLS     NacRuleMatchingAuthTypeEnum = "eap-tls"
    NacRuleMatchingAuthTypeEnum_EAPTTLS    NacRuleMatchingAuthTypeEnum = "eap-ttls"
    NacRuleMatchingAuthTypeEnum_IDP        NacRuleMatchingAuthTypeEnum = "idp"
    NacRuleMatchingAuthTypeEnum_MAB        NacRuleMatchingAuthTypeEnum = "mab"
    NacRuleMatchingAuthTypeEnum_PEAPTLS    NacRuleMatchingAuthTypeEnum = "peap-tls"
    NacRuleMatchingAuthTypeEnum_PSK        NacRuleMatchingAuthTypeEnum = "psk"
)

// NacRuleMatchingPortTypeEnum is a string enum.
// enum: `wired`, `wireless`
type NacRuleMatchingPortTypeEnum string

const (
    NacRuleMatchingPortTypeEnum_WIRED    NacRuleMatchingPortTypeEnum = "wired"
    NacRuleMatchingPortTypeEnum_WIRELESS NacRuleMatchingPortTypeEnum = "wireless"
)

// NacTagMatchEnum is a string enum.
// if `type`==`match`. enum: `cert_cn`, `cert_issuer`, `cert_san`, `cert_serial`, `cert_sub`, `cert_template`, `client_mac`, `idp_role`, `ingress_vlan`, `mdm_status`, `nas_ip`, `radius_group`, `realm`, `ssid`, `user_name`, `usermac_label`
type NacTagMatchEnum string

const (
    NacTagMatchEnum_CERTCN       NacTagMatchEnum = "cert_cn"
    NacTagMatchEnum_CERTISSUER   NacTagMatchEnum = "cert_issuer"
    NacTagMatchEnum_CERTSAN      NacTagMatchEnum = "cert_san"
    NacTagMatchEnum_CERTSERIAL   NacTagMatchEnum = "cert_serial"
    NacTagMatchEnum_CERTSUB      NacTagMatchEnum = "cert_sub"
    NacTagMatchEnum_CERTTEMPLATE NacTagMatchEnum = "cert_template"
    NacTagMatchEnum_CLIENTMAC    NacTagMatchEnum = "client_mac"
    NacTagMatchEnum_IDPROLE      NacTagMatchEnum = "idp_role"
    NacTagMatchEnum_INGRESSVLAN  NacTagMatchEnum = "ingress_vlan"
    NacTagMatchEnum_MDMSTATUS    NacTagMatchEnum = "mdm_status"
    NacTagMatchEnum_NASIP        NacTagMatchEnum = "nas_ip"
    NacTagMatchEnum_RADIUSGROUP  NacTagMatchEnum = "radius_group"
    NacTagMatchEnum_REALM        NacTagMatchEnum = "realm"
    NacTagMatchEnum_SSID         NacTagMatchEnum = "ssid"
    NacTagMatchEnum_USERNAME     NacTagMatchEnum = "user_name"
    NacTagMatchEnum_USERMACLABEL NacTagMatchEnum = "usermac_label"
)

// NacTagTypeEnum is a string enum.
// enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `session_timeout`, `username_attr`, `vlan`
type NacTagTypeEnum string

const (
    NacTagTypeEnum_EGRESSVLANNAMES   NacTagTypeEnum = "egress_vlan_names"
    NacTagTypeEnum_GBPTAG            NacTagTypeEnum = "gbp_tag"
    NacTagTypeEnum_MATCH             NacTagTypeEnum = "match"
    NacTagTypeEnum_RADIUSATTRS       NacTagTypeEnum = "radius_attrs"
    NacTagTypeEnum_RADIUSGROUP       NacTagTypeEnum = "radius_group"
    NacTagTypeEnum_RADIUSVENDORATTRS NacTagTypeEnum = "radius_vendor_attrs"
    NacTagTypeEnum_SESSIONTIMEOUT    NacTagTypeEnum = "session_timeout"
    NacTagTypeEnum_USERNAMEATTR      NacTagTypeEnum = "username_attr"
    NacTagTypeEnum_VLAN              NacTagTypeEnum = "vlan"
)

// NacTagUsernameAttrEnum is a string enum.
// enum: `automatic`, `cn`, `dns`, `email`, `upn`
type NacTagUsernameAttrEnum string

const (
    NacTagUsernameAttrEnum_AUTOMATIC NacTagUsernameAttrEnum = "automatic"
    NacTagUsernameAttrEnum_CN        NacTagUsernameAttrEnum = "cn"
    NacTagUsernameAttrEnum_DNS       NacTagUsernameAttrEnum = "dns"
    NacTagUsernameAttrEnum_EMAIL     NacTagUsernameAttrEnum = "email"
    NacTagUsernameAttrEnum_UPN       NacTagUsernameAttrEnum = "upn"
)

// OauthAppNameEnum is a string enum.
// enum: `crowdstrike`, `intune`, `jamf`, `mobicontrol`, `teams`, `vmware`, `zdx`, `zoom`
type OauthAppNameEnum string

const (
    OauthAppNameEnum_CROWDSTRIKE OauthAppNameEnum = "crowdstrike"
    OauthAppNameEnum_INTUNE      OauthAppNameEnum = "intune"
    OauthAppNameEnum_JAMF        OauthAppNameEnum = "jamf"
    OauthAppNameEnum_MOBICONTROL OauthAppNameEnum = "mobicontrol"
    OauthAppNameEnum_TEAMS       OauthAppNameEnum = "teams"
    OauthAppNameEnum_VMWARE      OauthAppNameEnum = "vmware"
    OauthAppNameEnum_ZDX         OauthAppNameEnum = "zdx"
    OauthAppNameEnum_ZOOM        OauthAppNameEnum = "zoom"
)

// OauthPingIdentityRegionEnum is a string enum.
// enum: `us` (United States, default), `ca` (Canada), `eu` (Europe), `asia` (Asia), `au` (Australia)
type OauthPingIdentityRegionEnum string

const (
    OauthPingIdentityRegionEnum_ASIA OauthPingIdentityRegionEnum = "asia"
    OauthPingIdentityRegionEnum_AU   OauthPingIdentityRegionEnum = "au"
    OauthPingIdentityRegionEnum_CA   OauthPingIdentityRegionEnum = "ca"
    OauthPingIdentityRegionEnum_EU   OauthPingIdentityRegionEnum = "eu"
    OauthPingIdentityRegionEnum_US   OauthPingIdentityRegionEnum = "us"
)

// OrgAssetCountDistinctEnum is a string enum.
// enum: `ibeacon_major`, `ibeacon_minor`, `ibeacon_uuid`, `mac`, `map_id`, `site_id`
type OrgAssetCountDistinctEnum string

const (
    OrgAssetCountDistinctEnum_IBEACONMAJOR OrgAssetCountDistinctEnum = "ibeacon_major"
    OrgAssetCountDistinctEnum_IBEACONMINOR OrgAssetCountDistinctEnum = "ibeacon_minor"
    OrgAssetCountDistinctEnum_IBEACONUUID  OrgAssetCountDistinctEnum = "ibeacon_uuid"
    OrgAssetCountDistinctEnum_MAC          OrgAssetCountDistinctEnum = "mac"
    OrgAssetCountDistinctEnum_MAPID        OrgAssetCountDistinctEnum = "map_id"
    OrgAssetCountDistinctEnum_SITEID       OrgAssetCountDistinctEnum = "site_id"
)

// OrgAutoRulesMatchDeviceTypeEnum is a string enum.
// optional/additional filter. enum: `ap`, `gateway`, `other`, `switch`
type OrgAutoRulesMatchDeviceTypeEnum string

const (
    OrgAutoRulesMatchDeviceTypeEnum_AP         OrgAutoRulesMatchDeviceTypeEnum = "ap"
    OrgAutoRulesMatchDeviceTypeEnum_GATEWAY    OrgAutoRulesMatchDeviceTypeEnum = "gateway"
    OrgAutoRulesMatchDeviceTypeEnum_OTHER      OrgAutoRulesMatchDeviceTypeEnum = "other"
    OrgAutoRulesMatchDeviceTypeEnum_ENUMSWITCH OrgAutoRulesMatchDeviceTypeEnum = "switch"
)

// OrgAutoRulesSrcEnum is a string enum.
// enum: `dns_suffix`, `geoip`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet`
type OrgAutoRulesSrcEnum string

const (
    OrgAutoRulesSrcEnum_DNSSUFFIX      OrgAutoRulesSrcEnum = "dns_suffix"
    OrgAutoRulesSrcEnum_GEOIP          OrgAutoRulesSrcEnum = "geoip"
    OrgAutoRulesSrcEnum_LLDPPORTDESC   OrgAutoRulesSrcEnum = "lldp_port_desc"
    OrgAutoRulesSrcEnum_LLDPSYSTEMNAME OrgAutoRulesSrcEnum = "lldp_system_name"
    OrgAutoRulesSrcEnum_MODEL          OrgAutoRulesSrcEnum = "model"
    OrgAutoRulesSrcEnum_NAME           OrgAutoRulesSrcEnum = "name"
    OrgAutoRulesSrcEnum_SUBNET         OrgAutoRulesSrcEnum = "subnet"
)

// OrgClientSessionsCountDistinctEnum is a string enum.
// enum: `ap`, `device`, `hostname`, `ip`, `model`, `os`, `ssid`, `vlan`
type OrgClientSessionsCountDistinctEnum string

const (
    OrgClientSessionsCountDistinctEnum_AP       OrgClientSessionsCountDistinctEnum = "ap"
    OrgClientSessionsCountDistinctEnum_DEVICE   OrgClientSessionsCountDistinctEnum = "device"
    OrgClientSessionsCountDistinctEnum_HOSTNAME OrgClientSessionsCountDistinctEnum = "hostname"
    OrgClientSessionsCountDistinctEnum_IP       OrgClientSessionsCountDistinctEnum = "ip"
    OrgClientSessionsCountDistinctEnum_MODEL    OrgClientSessionsCountDistinctEnum = "model"
    OrgClientSessionsCountDistinctEnum_OS       OrgClientSessionsCountDistinctEnum = "os"
    OrgClientSessionsCountDistinctEnum_SSID     OrgClientSessionsCountDistinctEnum = "ssid"
    OrgClientSessionsCountDistinctEnum_VLAN     OrgClientSessionsCountDistinctEnum = "vlan"
)

// OrgClientsCountDistinctEnum is a string enum.
// enum: `ap`, `device`, `hostname`, `ip`, `mac`, `model`, `os`, `ssid`, `vlan`
type OrgClientsCountDistinctEnum string

const (
    OrgClientsCountDistinctEnum_AP       OrgClientsCountDistinctEnum = "ap"
    OrgClientsCountDistinctEnum_DEVICE   OrgClientsCountDistinctEnum = "device"
    OrgClientsCountDistinctEnum_HOSTNAME OrgClientsCountDistinctEnum = "hostname"
    OrgClientsCountDistinctEnum_IP       OrgClientsCountDistinctEnum = "ip"
    OrgClientsCountDistinctEnum_MAC      OrgClientsCountDistinctEnum = "mac"
    OrgClientsCountDistinctEnum_MODEL    OrgClientsCountDistinctEnum = "model"
    OrgClientsCountDistinctEnum_OS       OrgClientsCountDistinctEnum = "os"
    OrgClientsCountDistinctEnum_SSID     OrgClientsCountDistinctEnum = "ssid"
    OrgClientsCountDistinctEnum_VLAN     OrgClientsCountDistinctEnum = "vlan"
)

// OrgDevicesCountDistinctEnum is a string enum.
// enum: `hostname`, `ip`, `lldp_mgmt_addr`, `lldp_port_id`, `lldp_system_desc`, `lldp_system_name`, `mac`, `model`, `mxedge_id`, `mxtunnel_status`, `site_id`, `version`
type OrgDevicesCountDistinctEnum string

const (
    OrgDevicesCountDistinctEnum_HOSTNAME       OrgDevicesCountDistinctEnum = "hostname"
    OrgDevicesCountDistinctEnum_IP             OrgDevicesCountDistinctEnum = "ip"
    OrgDevicesCountDistinctEnum_LLDPMGMTADDR   OrgDevicesCountDistinctEnum = "lldp_mgmt_addr"
    OrgDevicesCountDistinctEnum_LLDPPORTID     OrgDevicesCountDistinctEnum = "lldp_port_id"
    OrgDevicesCountDistinctEnum_LLDPSYSTEMDESC OrgDevicesCountDistinctEnum = "lldp_system_desc"
    OrgDevicesCountDistinctEnum_LLDPSYSTEMNAME OrgDevicesCountDistinctEnum = "lldp_system_name"
    OrgDevicesCountDistinctEnum_MAC            OrgDevicesCountDistinctEnum = "mac"
    OrgDevicesCountDistinctEnum_MODEL          OrgDevicesCountDistinctEnum = "model"
    OrgDevicesCountDistinctEnum_MXEDGEID       OrgDevicesCountDistinctEnum = "mxedge_id"
    OrgDevicesCountDistinctEnum_MXTUNNELSTATUS OrgDevicesCountDistinctEnum = "mxtunnel_status"
    OrgDevicesCountDistinctEnum_SITEID         OrgDevicesCountDistinctEnum = "site_id"
    OrgDevicesCountDistinctEnum_VERSION        OrgDevicesCountDistinctEnum = "version"
)

// OrgDevicesEventsCountDistinctEnum is a string enum.
// enum: `ap`, `apfw`, `model`, `org_id`, `site_id`, `text`, `timestamp`, `type`
type OrgDevicesEventsCountDistinctEnum string

const (
    OrgDevicesEventsCountDistinctEnum_AP        OrgDevicesEventsCountDistinctEnum = "ap"
    OrgDevicesEventsCountDistinctEnum_APFW      OrgDevicesEventsCountDistinctEnum = "apfw"
    OrgDevicesEventsCountDistinctEnum_MODEL     OrgDevicesEventsCountDistinctEnum = "model"
    OrgDevicesEventsCountDistinctEnum_ORGID     OrgDevicesEventsCountDistinctEnum = "org_id"
    OrgDevicesEventsCountDistinctEnum_SITEID    OrgDevicesEventsCountDistinctEnum = "site_id"
    OrgDevicesEventsCountDistinctEnum_TEXT      OrgDevicesEventsCountDistinctEnum = "text"
    OrgDevicesEventsCountDistinctEnum_TIMESTAMP OrgDevicesEventsCountDistinctEnum = "timestamp"
    OrgDevicesEventsCountDistinctEnum_ENUMTYPE  OrgDevicesEventsCountDistinctEnum = "type"
)

// OrgDevicesLastConfigsCountDistinctEnum is a string enum.
// enum: `mac`, `name`, `site_id`, `version`
type OrgDevicesLastConfigsCountDistinctEnum string

const (
    OrgDevicesLastConfigsCountDistinctEnum_MAC     OrgDevicesLastConfigsCountDistinctEnum = "mac"
    OrgDevicesLastConfigsCountDistinctEnum_NAME    OrgDevicesLastConfigsCountDistinctEnum = "name"
    OrgDevicesLastConfigsCountDistinctEnum_SITEID  OrgDevicesLastConfigsCountDistinctEnum = "site_id"
    OrgDevicesLastConfigsCountDistinctEnum_VERSION OrgDevicesLastConfigsCountDistinctEnum = "version"
)

// OrgGuestsCountDistinctEnum is a string enum.
// enum: `auth_method`, `company`, `ssid`
type OrgGuestsCountDistinctEnum string

const (
    OrgGuestsCountDistinctEnum_AUTHMETHOD OrgGuestsCountDistinctEnum = "auth_method"
    OrgGuestsCountDistinctEnum_COMPANY    OrgGuestsCountDistinctEnum = "company"
    OrgGuestsCountDistinctEnum_SSID       OrgGuestsCountDistinctEnum = "ssid"
)

// OrgLicenseActionOperationEnum is a string enum.
// to move a license, use the `amend` operation. enum: `amend`, `annotate`, `delete`, `unamend`
type OrgLicenseActionOperationEnum string

const (
    OrgLicenseActionOperationEnum_AMEND    OrgLicenseActionOperationEnum = "amend"
    OrgLicenseActionOperationEnum_ANNOTATE OrgLicenseActionOperationEnum = "annotate"
    OrgLicenseActionOperationEnum_DELETE   OrgLicenseActionOperationEnum = "delete"
    OrgLicenseActionOperationEnum_UNAMEND  OrgLicenseActionOperationEnum = "unamend"
)

// OrgLogsCountDistinctEnum is a string enum.
// enum: `admin_id`, `admin_name`, `message`, `site_id`
type OrgLogsCountDistinctEnum string

const (
    OrgLogsCountDistinctEnum_ADMINID   OrgLogsCountDistinctEnum = "admin_id"
    OrgLogsCountDistinctEnum_ADMINNAME OrgLogsCountDistinctEnum = "admin_name"
    OrgLogsCountDistinctEnum_MESSAGE   OrgLogsCountDistinctEnum = "message"
    OrgLogsCountDistinctEnum_SITEID    OrgLogsCountDistinctEnum = "site_id"
)

// OrgMxedgeCountDistinctEnum is a string enum.
// enum: `distro`, `model`, `mxcluster_id`, `site_id`, `tunterm_version`
type OrgMxedgeCountDistinctEnum string

const (
    OrgMxedgeCountDistinctEnum_DISTRO         OrgMxedgeCountDistinctEnum = "distro"
    OrgMxedgeCountDistinctEnum_MODEL          OrgMxedgeCountDistinctEnum = "model"
    OrgMxedgeCountDistinctEnum_MXCLUSTERID    OrgMxedgeCountDistinctEnum = "mxcluster_id"
    OrgMxedgeCountDistinctEnum_SITEID         OrgMxedgeCountDistinctEnum = "site_id"
    OrgMxedgeCountDistinctEnum_TUNTERMVERSION OrgMxedgeCountDistinctEnum = "tunterm_version"
)

// OrgMxedgeEventsCountDistinctEnum is a string enum.
// enum: `mxcluster_id`, `mxedge_id`, `package`, `type`
type OrgMxedgeEventsCountDistinctEnum string

const (
    OrgMxedgeEventsCountDistinctEnum_MXCLUSTERID OrgMxedgeEventsCountDistinctEnum = "mxcluster_id"
    OrgMxedgeEventsCountDistinctEnum_MXEDGEID    OrgMxedgeEventsCountDistinctEnum = "mxedge_id"
    OrgMxedgeEventsCountDistinctEnum_ENUMPACKAGE OrgMxedgeEventsCountDistinctEnum = "package"
    OrgMxedgeEventsCountDistinctEnum_ENUMTYPE    OrgMxedgeEventsCountDistinctEnum = "type"
)

// OrgNacClientEventsCountDistinctEnum is a string enum.
// enum: `ap`, `auth_type`, `dryrun_nacrule_id`, `mac`, `nacrule_id`, `nas_vendor`, `ssid`, `type`, `username`, `vlan`
type OrgNacClientEventsCountDistinctEnum string

const (
    OrgNacClientEventsCountDistinctEnum_AP              OrgNacClientEventsCountDistinctEnum = "ap"
    OrgNacClientEventsCountDistinctEnum_AUTHTYPE        OrgNacClientEventsCountDistinctEnum = "auth_type"
    OrgNacClientEventsCountDistinctEnum_DRYRUNNACRULEID OrgNacClientEventsCountDistinctEnum = "dryrun_nacrule_id"
    OrgNacClientEventsCountDistinctEnum_MAC             OrgNacClientEventsCountDistinctEnum = "mac"
    OrgNacClientEventsCountDistinctEnum_NACRULEID       OrgNacClientEventsCountDistinctEnum = "nacrule_id"
    OrgNacClientEventsCountDistinctEnum_NASVENDOR       OrgNacClientEventsCountDistinctEnum = "nas_vendor"
    OrgNacClientEventsCountDistinctEnum_SSID            OrgNacClientEventsCountDistinctEnum = "ssid"
    OrgNacClientEventsCountDistinctEnum_ENUMTYPE        OrgNacClientEventsCountDistinctEnum = "type"
    OrgNacClientEventsCountDistinctEnum_USERNAME        OrgNacClientEventsCountDistinctEnum = "username"
    OrgNacClientEventsCountDistinctEnum_VLAN            OrgNacClientEventsCountDistinctEnum = "vlan"
)

// OrgNacClientsCountDistinctEnum is a string enum.
// enum: `auth_type`, `last_ap`, `last_nacrule_id`, `last_nas_vendor`, `last_ssid`, `last_status`, `last_username`, `last_vlan`, `mac`, `mdm_compliance`, `mdm_provider`, `type`
type OrgNacClientsCountDistinctEnum string

const (
    OrgNacClientsCountDistinctEnum_AUTHTYPE      OrgNacClientsCountDistinctEnum = "auth_type"
    OrgNacClientsCountDistinctEnum_LASTAP        OrgNacClientsCountDistinctEnum = "last_ap"
    OrgNacClientsCountDistinctEnum_LASTNACRULEID OrgNacClientsCountDistinctEnum = "last_nacrule_id"
    OrgNacClientsCountDistinctEnum_LASTNASVENDOR OrgNacClientsCountDistinctEnum = "last_nas_vendor"
    OrgNacClientsCountDistinctEnum_LASTSSID      OrgNacClientsCountDistinctEnum = "last_ssid"
    OrgNacClientsCountDistinctEnum_LASTSTATUS    OrgNacClientsCountDistinctEnum = "last_status"
    OrgNacClientsCountDistinctEnum_LASTUSERNAME  OrgNacClientsCountDistinctEnum = "last_username"
    OrgNacClientsCountDistinctEnum_LASTVLAN      OrgNacClientsCountDistinctEnum = "last_vlan"
    OrgNacClientsCountDistinctEnum_MAC           OrgNacClientsCountDistinctEnum = "mac"
    OrgNacClientsCountDistinctEnum_MDMCOMPLIANCE OrgNacClientsCountDistinctEnum = "mdm_compliance"
    OrgNacClientsCountDistinctEnum_MDMPROVIDER   OrgNacClientsCountDistinctEnum = "mdm_provider"
    OrgNacClientsCountDistinctEnum_ENUMTYPE      OrgNacClientsCountDistinctEnum = "type"
)

// OrgOtherdevicesEventsCountDistinctEnum is a string enum.
// enum: `mac`, `site_id`, `type`, `vendor`
type OrgOtherdevicesEventsCountDistinctEnum string

const (
    OrgOtherdevicesEventsCountDistinctEnum_MAC      OrgOtherdevicesEventsCountDistinctEnum = "mac"
    OrgOtherdevicesEventsCountDistinctEnum_SITEID   OrgOtherdevicesEventsCountDistinctEnum = "site_id"
    OrgOtherdevicesEventsCountDistinctEnum_ENUMTYPE OrgOtherdevicesEventsCountDistinctEnum = "type"
    OrgOtherdevicesEventsCountDistinctEnum_VENDOR   OrgOtherdevicesEventsCountDistinctEnum = "vendor"
)

// OrgPskPortalLogsCountDistinctEnum is a string enum.
// enum: `admin_id`, `admin_name`, `psk_id`, `psk_name`, `pskportal_id`, `user_id`
type OrgPskPortalLogsCountDistinctEnum string

const (
    OrgPskPortalLogsCountDistinctEnum_ADMINID     OrgPskPortalLogsCountDistinctEnum = "admin_id"
    OrgPskPortalLogsCountDistinctEnum_ADMINNAME   OrgPskPortalLogsCountDistinctEnum = "admin_name"
    OrgPskPortalLogsCountDistinctEnum_PSKID       OrgPskPortalLogsCountDistinctEnum = "psk_id"
    OrgPskPortalLogsCountDistinctEnum_PSKNAME     OrgPskPortalLogsCountDistinctEnum = "psk_name"
    OrgPskPortalLogsCountDistinctEnum_PSKPORTALID OrgPskPortalLogsCountDistinctEnum = "pskportal_id"
    OrgPskPortalLogsCountDistinctEnum_USERID      OrgPskPortalLogsCountDistinctEnum = "user_id"
)

// OrgSettingMistNacIpVersionEnum is a string enum.
// by default NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4. enum: `v4`, `v6`
type OrgSettingMistNacIpVersionEnum string

const (
    OrgSettingMistNacIpVersionEnum_V4 OrgSettingMistNacIpVersionEnum = "v4"
    OrgSettingMistNacIpVersionEnum_V6 OrgSettingMistNacIpVersionEnum = "v6"
)

// OrgSiteSleTypeEnum is a string enum.
// enum: `wan`, `wifi`, `wired`
type OrgSiteSleTypeEnum string

const (
    OrgSiteSleTypeEnum_WAN   OrgSiteSleTypeEnum = "wan"
    OrgSiteSleTypeEnum_WIFI  OrgSiteSleTypeEnum = "wifi"
    OrgSiteSleTypeEnum_WIRED OrgSiteSleTypeEnum = "wired"
)

// OrgSitesCountDistinctEnum is a string enum.
// enum: `analytic_enabled`, `app_waking`, `asset_enabled`, `auto_upgrade_enabled`, `auto_upgrade_version`, `country_code`, `honeypot_enabled`, `id`, `locate_unconnected`, `mesh_enabled`, `name`, `remote_syslog_enabled`, `rogue_enabled`, `rtsa_enabled`, `vna_enabled`, `wifi_enabled`
type OrgSitesCountDistinctEnum string

const (
    OrgSitesCountDistinctEnum_ANALYTICENABLED     OrgSitesCountDistinctEnum = "analytic_enabled"
    OrgSitesCountDistinctEnum_APPWAKING           OrgSitesCountDistinctEnum = "app_waking"
    OrgSitesCountDistinctEnum_ASSETENABLED        OrgSitesCountDistinctEnum = "asset_enabled"
    OrgSitesCountDistinctEnum_AUTOUPGRADEENABLED  OrgSitesCountDistinctEnum = "auto_upgrade_enabled"
    OrgSitesCountDistinctEnum_AUTOUPGRADEVERSION  OrgSitesCountDistinctEnum = "auto_upgrade_version"
    OrgSitesCountDistinctEnum_COUNTRYCODE         OrgSitesCountDistinctEnum = "country_code"
    OrgSitesCountDistinctEnum_HONEYPOTENABLED     OrgSitesCountDistinctEnum = "honeypot_enabled"
    OrgSitesCountDistinctEnum_ID                  OrgSitesCountDistinctEnum = "id"
    OrgSitesCountDistinctEnum_LOCATEUNCONNECTED   OrgSitesCountDistinctEnum = "locate_unconnected"
    OrgSitesCountDistinctEnum_MESHENABLED         OrgSitesCountDistinctEnum = "mesh_enabled"
    OrgSitesCountDistinctEnum_NAME                OrgSitesCountDistinctEnum = "name"
    OrgSitesCountDistinctEnum_REMOTESYSLOGENABLED OrgSitesCountDistinctEnum = "remote_syslog_enabled"
    OrgSitesCountDistinctEnum_ROGUEENABLED        OrgSitesCountDistinctEnum = "rogue_enabled"
    OrgSitesCountDistinctEnum_RTSAENABLED         OrgSitesCountDistinctEnum = "rtsa_enabled"
    OrgSitesCountDistinctEnum_VNAENABLED          OrgSitesCountDistinctEnum = "vna_enabled"
    OrgSitesCountDistinctEnum_WIFIENABLED         OrgSitesCountDistinctEnum = "wifi_enabled"
)

// OrgTicketsCountDistinctEnum is a string enum.
// enum: `status`, `type`
type OrgTicketsCountDistinctEnum string

const (
    OrgTicketsCountDistinctEnum_STATUS   OrgTicketsCountDistinctEnum = "status"
    OrgTicketsCountDistinctEnum_ENUMTYPE OrgTicketsCountDistinctEnum = "type"
)

// OrgTunnelCountDistinctEnum is a string enum.
// enum: `ap`, `auth_algo`, `encrypt_algo`, `ike_version`, `ip`, `last_event`, `mac`, `mxcluster_id`, `mxedge_id`, `node`, `peer_host`, `peer_ip`, `peer_mxedge_id`, `protocol`, `remote_ip`, `remote_port`, `site_id`, `state`, `tunnel_name`, `up`, `wxtunnel_id`
type OrgTunnelCountDistinctEnum string

const (
    OrgTunnelCountDistinctEnum_AP           OrgTunnelCountDistinctEnum = "ap"
    OrgTunnelCountDistinctEnum_AUTHALGO     OrgTunnelCountDistinctEnum = "auth_algo"
    OrgTunnelCountDistinctEnum_ENCRYPTALGO  OrgTunnelCountDistinctEnum = "encrypt_algo"
    OrgTunnelCountDistinctEnum_IKEVERSION   OrgTunnelCountDistinctEnum = "ike_version"
    OrgTunnelCountDistinctEnum_IP           OrgTunnelCountDistinctEnum = "ip"
    OrgTunnelCountDistinctEnum_LASTEVENT    OrgTunnelCountDistinctEnum = "last_event"
    OrgTunnelCountDistinctEnum_MAC          OrgTunnelCountDistinctEnum = "mac"
    OrgTunnelCountDistinctEnum_MXCLUSTERID  OrgTunnelCountDistinctEnum = "mxcluster_id"
    OrgTunnelCountDistinctEnum_MXEDGEID     OrgTunnelCountDistinctEnum = "mxedge_id"
    OrgTunnelCountDistinctEnum_NODE         OrgTunnelCountDistinctEnum = "node"
    OrgTunnelCountDistinctEnum_PEERHOST     OrgTunnelCountDistinctEnum = "peer_host"
    OrgTunnelCountDistinctEnum_PEERIP       OrgTunnelCountDistinctEnum = "peer_ip"
    OrgTunnelCountDistinctEnum_PEERMXEDGEID OrgTunnelCountDistinctEnum = "peer_mxedge_id"
    OrgTunnelCountDistinctEnum_PROTOCOL     OrgTunnelCountDistinctEnum = "protocol"
    OrgTunnelCountDistinctEnum_REMOTEIP     OrgTunnelCountDistinctEnum = "remote_ip"
    OrgTunnelCountDistinctEnum_REMOTEPORT   OrgTunnelCountDistinctEnum = "remote_port"
    OrgTunnelCountDistinctEnum_SITEID       OrgTunnelCountDistinctEnum = "site_id"
    OrgTunnelCountDistinctEnum_STATE        OrgTunnelCountDistinctEnum = "state"
    OrgTunnelCountDistinctEnum_TUNNELNAME   OrgTunnelCountDistinctEnum = "tunnel_name"
    OrgTunnelCountDistinctEnum_UP           OrgTunnelCountDistinctEnum = "up"
    OrgTunnelCountDistinctEnum_WXTUNNELID   OrgTunnelCountDistinctEnum = "wxtunnel_id"
)

// OrgTunnelTypeCountEnum is a string enum.
// enum: `wan`, `wxtunnel`
type OrgTunnelTypeCountEnum string

const (
    OrgTunnelTypeCountEnum_WAN      OrgTunnelTypeCountEnum = "wan"
    OrgTunnelTypeCountEnum_WXTUNNEL OrgTunnelTypeCountEnum = "wxtunnel"
)

// OrgWanClientsCountDistinctEnum is a string enum.
// enum: `hostname`, `ip`, `mac`, `mfg`, `network`
type OrgWanClientsCountDistinctEnum string

const (
    OrgWanClientsCountDistinctEnum_HOSTNAME OrgWanClientsCountDistinctEnum = "hostname"
    OrgWanClientsCountDistinctEnum_IP       OrgWanClientsCountDistinctEnum = "ip"
    OrgWanClientsCountDistinctEnum_MAC      OrgWanClientsCountDistinctEnum = "mac"
    OrgWanClientsCountDistinctEnum_MFG      OrgWanClientsCountDistinctEnum = "mfg"
    OrgWanClientsCountDistinctEnum_NETWORK  OrgWanClientsCountDistinctEnum = "network"
)

// OrgWanClientsEventsCountDistinctEnum is a string enum.
// enum: `hostname`, `ip`, `mac`, `mfg`, `type`
type OrgWanClientsEventsCountDistinctEnum string

const (
    OrgWanClientsEventsCountDistinctEnum_HOSTNAME OrgWanClientsEventsCountDistinctEnum = "hostname"
    OrgWanClientsEventsCountDistinctEnum_IP       OrgWanClientsEventsCountDistinctEnum = "ip"
    OrgWanClientsEventsCountDistinctEnum_MAC      OrgWanClientsEventsCountDistinctEnum = "mac"
    OrgWanClientsEventsCountDistinctEnum_MFG      OrgWanClientsEventsCountDistinctEnum = "mfg"
    OrgWanClientsEventsCountDistinctEnum_ENUMTYPE OrgWanClientsEventsCountDistinctEnum = "type"
)

// OrgWiredClientsCountDistinctEnum is a string enum.
// enum: `device_mac`, `mac`, `port_id`, `site_id`, `type`, `vlan`
type OrgWiredClientsCountDistinctEnum string

const (
    OrgWiredClientsCountDistinctEnum_DEVICEMAC OrgWiredClientsCountDistinctEnum = "device_mac"
    OrgWiredClientsCountDistinctEnum_MAC       OrgWiredClientsCountDistinctEnum = "mac"
    OrgWiredClientsCountDistinctEnum_PORTID    OrgWiredClientsCountDistinctEnum = "port_id"
    OrgWiredClientsCountDistinctEnum_SITEID    OrgWiredClientsCountDistinctEnum = "site_id"
    OrgWiredClientsCountDistinctEnum_ENUMTYPE  OrgWiredClientsCountDistinctEnum = "type"
    OrgWiredClientsCountDistinctEnum_VLAN      OrgWiredClientsCountDistinctEnum = "vlan"
)

// OspfAreaNetworkAuthTypeEnum is a string enum.
// auth type. enum: `md5`, `none`, `password`
type OspfAreaNetworkAuthTypeEnum string

const (
    OspfAreaNetworkAuthTypeEnum_MD5      OspfAreaNetworkAuthTypeEnum = "md5"
    OspfAreaNetworkAuthTypeEnum_NONE     OspfAreaNetworkAuthTypeEnum = "none"
    OspfAreaNetworkAuthTypeEnum_PASSWORD OspfAreaNetworkAuthTypeEnum = "password"
)

// OspfAreaNetworkInterfaceTypeEnum is a string enum.
// interface type (nbma = non-broadcast multi-access). enum: `broadcast`, `nbma`, `p2mp`, `p2p`
type OspfAreaNetworkInterfaceTypeEnum string

const (
    OspfAreaNetworkInterfaceTypeEnum_BROADCAST OspfAreaNetworkInterfaceTypeEnum = "broadcast"
    OspfAreaNetworkInterfaceTypeEnum_NBMA      OspfAreaNetworkInterfaceTypeEnum = "nbma"
    OspfAreaNetworkInterfaceTypeEnum_P2MP      OspfAreaNetworkInterfaceTypeEnum = "p2mp"
    OspfAreaNetworkInterfaceTypeEnum_P2P       OspfAreaNetworkInterfaceTypeEnum = "p2p"
)

// OspfAreaTypeEnum is a string enum.
// OSPF type. enum: `default`, `nssa`, `stub`
type OspfAreaTypeEnum string

const (
    OspfAreaTypeEnum_ENUMDEFAULT OspfAreaTypeEnum = "default"
    OspfAreaTypeEnum_NSSA        OspfAreaTypeEnum = "nssa"
    OspfAreaTypeEnum_STUB        OspfAreaTypeEnum = "stub"
)

// OtherDeviceUpdateOperationEnum is a string enum.
// The operation being performed. enum: `assign`, `unassign`
type OtherDeviceUpdateOperationEnum string

const (
    OtherDeviceUpdateOperationEnum_ASSIGN   OtherDeviceUpdateOperationEnum = "assign"
    OtherDeviceUpdateOperationEnum_UNASSIGN OtherDeviceUpdateOperationEnum = "unassign"
)

// PcapTypeEnum is a string enum.
// enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless`
type PcapTypeEnum string

const (
    PcapTypeEnum_CLIENT            PcapTypeEnum = "client"
    PcapTypeEnum_GATEWAY           PcapTypeEnum = "gateway"
    PcapTypeEnum_NEWASSOC          PcapTypeEnum = "new_assoc"
    PcapTypeEnum_RADIOTAP          PcapTypeEnum = "radiotap"
    PcapTypeEnum_ENUMRADIOTAPWIRED PcapTypeEnum = "radiotap,wired"
    PcapTypeEnum_WIRED             PcapTypeEnum = "wired"
    PcapTypeEnum_WIRELESS          PcapTypeEnum = "wireless"
)

// PortalTemplateAlignmentEnum is a string enum.
// defines alignment on portal. enum: `center`, `left`, `right`
type PortalTemplateAlignmentEnum string

const (
    PortalTemplateAlignmentEnum_CENTER PortalTemplateAlignmentEnum = "center"
    PortalTemplateAlignmentEnum_LEFT   PortalTemplateAlignmentEnum = "left"
    PortalTemplateAlignmentEnum_RIGHT  PortalTemplateAlignmentEnum = "right"
)

// PrivilegeMspRoleEnum is a string enum.
// access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
type PrivilegeMspRoleEnum string

const (
    PrivilegeMspRoleEnum_ADMIN     PrivilegeMspRoleEnum = "admin"
    PrivilegeMspRoleEnum_HELPDESK  PrivilegeMspRoleEnum = "helpdesk"
    PrivilegeMspRoleEnum_INSTALLER PrivilegeMspRoleEnum = "installer"
    PrivilegeMspRoleEnum_READ      PrivilegeMspRoleEnum = "read"
    PrivilegeMspRoleEnum_WRITE     PrivilegeMspRoleEnum = "write"
)

// PrivilegeMspScopeEnum is a string enum.
// enum: `msp`, `org`, `orggroup`
type PrivilegeMspScopeEnum string

const (
    PrivilegeMspScopeEnum_MSP      PrivilegeMspScopeEnum = "msp"
    PrivilegeMspScopeEnum_ORG      PrivilegeMspScopeEnum = "org"
    PrivilegeMspScopeEnum_ORGGROUP PrivilegeMspScopeEnum = "orggroup"
)

// PrivilegeOrgRoleEnum is a string enum.
// access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
type PrivilegeOrgRoleEnum string

const (
    PrivilegeOrgRoleEnum_ADMIN     PrivilegeOrgRoleEnum = "admin"
    PrivilegeOrgRoleEnum_HELPDESK  PrivilegeOrgRoleEnum = "helpdesk"
    PrivilegeOrgRoleEnum_INSTALLER PrivilegeOrgRoleEnum = "installer"
    PrivilegeOrgRoleEnum_READ      PrivilegeOrgRoleEnum = "read"
    PrivilegeOrgRoleEnum_WRITE     PrivilegeOrgRoleEnum = "write"
)

// PrivilegeOrgScopeEnum is a string enum.
// enum: `org`, `site`, `sitegroup`
type PrivilegeOrgScopeEnum string

const (
    PrivilegeOrgScopeEnum_ORG       PrivilegeOrgScopeEnum = "org"
    PrivilegeOrgScopeEnum_SITE      PrivilegeOrgScopeEnum = "site"
    PrivilegeOrgScopeEnum_SITEGROUP PrivilegeOrgScopeEnum = "sitegroup"
)

// ProtectReAllowedServiceEnum is a string enum.
// enum: `icmp`, `ssh`
type ProtectReAllowedServiceEnum string

const (
    ProtectReAllowedServiceEnum_ICMP ProtectReAllowedServiceEnum = "icmp"
    ProtectReAllowedServiceEnum_SSH  ProtectReAllowedServiceEnum = "ssh"
)

// ProtectReCustomProtocolEnum is a string enum.
// enum: `any`, `icmp`, `tcp`, `udp`
type ProtectReCustomProtocolEnum string

const (
    ProtectReCustomProtocolEnum_ANY  ProtectReCustomProtocolEnum = "any"
    ProtectReCustomProtocolEnum_ICMP ProtectReCustomProtocolEnum = "icmp"
    ProtectReCustomProtocolEnum_TCP  ProtectReCustomProtocolEnum = "tcp"
    ProtectReCustomProtocolEnum_UDP  ProtectReCustomProtocolEnum = "udp"
)

// PskPortalAuthEnum is a string enum.
// enum: `sponsor`, `sso`
type PskPortalAuthEnum string

const (
    PskPortalAuthEnum_SPONSOR PskPortalAuthEnum = "sponsor"
    PskPortalAuthEnum_SSO     PskPortalAuthEnum = "sso"
)

// PskPortalSsoIdpSignAlgoEnum is a string enum.
// Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`. enum: `sha1`, `sha256`, `sha384`, `sha512`
type PskPortalSsoIdpSignAlgoEnum string

const (
    PskPortalSsoIdpSignAlgoEnum_SHA1   PskPortalSsoIdpSignAlgoEnum = "sha1"
    PskPortalSsoIdpSignAlgoEnum_SHA256 PskPortalSsoIdpSignAlgoEnum = "sha256"
    PskPortalSsoIdpSignAlgoEnum_SHA384 PskPortalSsoIdpSignAlgoEnum = "sha384"
    PskPortalSsoIdpSignAlgoEnum_SHA512 PskPortalSsoIdpSignAlgoEnum = "sha512"
)

// PskPortalTypeEnum is a string enum.
// for personal psk portal. enum: `admin`, `byod`
type PskPortalTypeEnum string

const (
    PskPortalTypeEnum_ADMIN PskPortalTypeEnum = "admin"
    PskPortalTypeEnum_BYOD  PskPortalTypeEnum = "byod"
)

// PskUsageEnum is a string enum.
// enum: `macs`, `multi`, `single`
type PskUsageEnum string

const (
    PskUsageEnum_MACS   PskUsageEnum = "macs"
    PskUsageEnum_MULTI  PskUsageEnum = "multi"
    PskUsageEnum_SINGLE PskUsageEnum = "single"
)

// RadioBand24UsageEnum is a string enum.
// enum: `24`, `5`, `6`, `auto`
type RadioBand24UsageEnum string

const (
    RadioBand24UsageEnum_ENUM24 RadioBand24UsageEnum = "24"
    RadioBand24UsageEnum_ENUM5  RadioBand24UsageEnum = "5"
    RadioBand24UsageEnum_ENUM6  RadioBand24UsageEnum = "6"
    RadioBand24UsageEnum_AUTO   RadioBand24UsageEnum = "auto"
)

// RadioBandAntennaModeEnum is a string enum.
// enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`
type RadioBandAntennaModeEnum string

const (
    RadioBandAntennaModeEnum_ENUM1X1     RadioBandAntennaModeEnum = "1x1"
    RadioBandAntennaModeEnum_ENUM2X2     RadioBandAntennaModeEnum = "2x2"
    RadioBandAntennaModeEnum_ENUM3X3     RadioBandAntennaModeEnum = "3x3"
    RadioBandAntennaModeEnum_ENUM4X4     RadioBandAntennaModeEnum = "4x4"
    RadioBandAntennaModeEnum_ENUMDEFAULT RadioBandAntennaModeEnum = "default"
)

// RadioBandPreambleEnum is a string enum.
// enum: `auto`, `long`, `short`
type RadioBandPreambleEnum string

const (
    RadioBandPreambleEnum_AUTO  RadioBandPreambleEnum = "auto"
    RadioBandPreambleEnum_LONG  RadioBandPreambleEnum = "long"
    RadioBandPreambleEnum_SHORT RadioBandPreambleEnum = "short"
)

// RadiusKeywrapFormatEnum is a string enum.
// enum: `ascii`, `hex`
type RadiusKeywrapFormatEnum string

const (
    RadiusKeywrapFormatEnum_ASCII RadiusKeywrapFormatEnum = "ascii"
    RadiusKeywrapFormatEnum_HEX   RadiusKeywrapFormatEnum = "hex"
)

// RecaptchaFlavorEnum is a string enum.
// flavor of the captcha. enum: `google`, `hcaptcha`
type RecaptchaFlavorEnum string

const (
    RecaptchaFlavorEnum_GOOGLE   RecaptchaFlavorEnum = "google"
    RecaptchaFlavorEnum_HCAPTCHA RecaptchaFlavorEnum = "hcaptcha"
)

// RemoteSyslogFacilityEnum is a string enum.
// enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`
type RemoteSyslogFacilityEnum string

const (
    RemoteSyslogFacilityEnum_ANY                 RemoteSyslogFacilityEnum = "any"
    RemoteSyslogFacilityEnum_AUTHORIZATION       RemoteSyslogFacilityEnum = "authorization"
    RemoteSyslogFacilityEnum_CHANGELOG           RemoteSyslogFacilityEnum = "change-log"
    RemoteSyslogFacilityEnum_CONFIG              RemoteSyslogFacilityEnum = "config"
    RemoteSyslogFacilityEnum_CONFLICTLOG         RemoteSyslogFacilityEnum = "conflict-log"
    RemoteSyslogFacilityEnum_DAEMON              RemoteSyslogFacilityEnum = "daemon"
    RemoteSyslogFacilityEnum_DFC                 RemoteSyslogFacilityEnum = "dfc"
    RemoteSyslogFacilityEnum_EXTERNAL            RemoteSyslogFacilityEnum = "external"
    RemoteSyslogFacilityEnum_FIREWALL            RemoteSyslogFacilityEnum = "firewall"
    RemoteSyslogFacilityEnum_FTP                 RemoteSyslogFacilityEnum = "ftp"
    RemoteSyslogFacilityEnum_INTERACTIVECOMMANDS RemoteSyslogFacilityEnum = "interactive-commands"
    RemoteSyslogFacilityEnum_KERNEL              RemoteSyslogFacilityEnum = "kernel"
    RemoteSyslogFacilityEnum_NTP                 RemoteSyslogFacilityEnum = "ntp"
    RemoteSyslogFacilityEnum_PFE                 RemoteSyslogFacilityEnum = "pfe"
    RemoteSyslogFacilityEnum_SECURITY            RemoteSyslogFacilityEnum = "security"
    RemoteSyslogFacilityEnum_USER                RemoteSyslogFacilityEnum = "user"
)

// RemoteSyslogServerProtocolEnum is a string enum.
// enum: `tcp`, `udp`
type RemoteSyslogServerProtocolEnum string

const (
    RemoteSyslogServerProtocolEnum_TCP RemoteSyslogServerProtocolEnum = "tcp"
    RemoteSyslogServerProtocolEnum_UDP RemoteSyslogServerProtocolEnum = "udp"
)

// RemoteSyslogSeverityEnum is a string enum.
// enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`
type RemoteSyslogSeverityEnum string

const (
    RemoteSyslogSeverityEnum_ALERT     RemoteSyslogSeverityEnum = "alert"
    RemoteSyslogSeverityEnum_ANY       RemoteSyslogSeverityEnum = "any"
    RemoteSyslogSeverityEnum_CRITICAL  RemoteSyslogSeverityEnum = "critical"
    RemoteSyslogSeverityEnum_EMERGENCY RemoteSyslogSeverityEnum = "emergency"
    RemoteSyslogSeverityEnum_ENUMERROR RemoteSyslogSeverityEnum = "error"
    RemoteSyslogSeverityEnum_INFO      RemoteSyslogSeverityEnum = "info"
    RemoteSyslogSeverityEnum_NOTICE    RemoteSyslogSeverityEnum = "notice"
    RemoteSyslogSeverityEnum_WARNING   RemoteSyslogSeverityEnum = "warning"
)

// RemoteSyslogTimeFormatEnum is a string enum.
// enum: `millisecond`, `year`, `year millisecond`
type RemoteSyslogTimeFormatEnum string

const (
    RemoteSyslogTimeFormatEnum_MILLISECOND         RemoteSyslogTimeFormatEnum = "millisecond"
    RemoteSyslogTimeFormatEnum_YEAR                RemoteSyslogTimeFormatEnum = "year"
    RemoteSyslogTimeFormatEnum_ENUMYEARMILLISECOND RemoteSyslogTimeFormatEnum = "year millisecond"
)

// ResolutionEnum is a string enum.
// enum: `default`, `fine`
type ResolutionEnum string

const (
    ResolutionEnum_ENUMDEFAULT ResolutionEnum = "default"
    ResolutionEnum_FINE        ResolutionEnum = "fine"
)

// ResponseAsyncLicenseStatusEnum is a string enum.
// processing status of async. enum: `prepared`, `ongoing`, `done`
type ResponseAsyncLicenseStatusEnum string

const (
    ResponseAsyncLicenseStatusEnum_PREPARED ResponseAsyncLicenseStatusEnum = "prepared"
    ResponseAsyncLicenseStatusEnum_ONGOING  ResponseAsyncLicenseStatusEnum = "ongoing"
    ResponseAsyncLicenseStatusEnum_DONE     ResponseAsyncLicenseStatusEnum = "done"
)

// ResponseAutoZoneStatusEnum is a string enum.
// The status for the auto zones service for a given map. enum:
// * not_started: The auto zones service has not been run on this map or the results were cleared by the user
// * in_progress: The auto zones service is currently in progress
// * awaiting_review: The auto zones service has completed and suggested location zones to be added to the map
// * error: There was an error with the auto zones service
type ResponseAutoZoneStatusEnum string

const (
    ResponseAutoZoneStatusEnum_INPROGRESS     ResponseAutoZoneStatusEnum = "in_progress"
    ResponseAutoZoneStatusEnum_AWAITINGREVIEW ResponseAutoZoneStatusEnum = "awaiting_review"
    ResponseAutoZoneStatusEnum_NOTSTARTED     ResponseAutoZoneStatusEnum = "not_started"
    ResponseAutoZoneStatusEnum_ENUMERROR      ResponseAutoZoneStatusEnum = "error"
)

// ResponseDeviceSnapshotStatusEnum is a string enum.
// enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
type ResponseDeviceSnapshotStatusEnum string

const (
    ResponseDeviceSnapshotStatusEnum_ENUMERROR  ResponseDeviceSnapshotStatusEnum = "error"
    ResponseDeviceSnapshotStatusEnum_INPROGRESS ResponseDeviceSnapshotStatusEnum = "inprogress"
    ResponseDeviceSnapshotStatusEnum_SCHEDULED  ResponseDeviceSnapshotStatusEnum = "scheduled"
    ResponseDeviceSnapshotStatusEnum_STARTING   ResponseDeviceSnapshotStatusEnum = "starting"
    ResponseDeviceSnapshotStatusEnum_SUCCESS    ResponseDeviceSnapshotStatusEnum = "success"
)

// ResponseMapImportApActionEnum is a string enum.
// enum: `assigned-named-placed`, `assigned-placed`, `ignored`, `named-placed`, `placed`
type ResponseMapImportApActionEnum string

const (
    ResponseMapImportApActionEnum_ASSIGNEDNAMEDPLACED ResponseMapImportApActionEnum = "assigned-named-placed"
    ResponseMapImportApActionEnum_ASSIGNEDPLACED      ResponseMapImportApActionEnum = "assigned-placed"
    ResponseMapImportApActionEnum_IGNORED             ResponseMapImportApActionEnum = "ignored"
    ResponseMapImportApActionEnum_NAMEDPLACED         ResponseMapImportApActionEnum = "named-placed"
    ResponseMapImportApActionEnum_PLACED              ResponseMapImportApActionEnum = "placed"
)

// ResponseOrgInventoryChangeOpEnum is a string enum.
// enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist`
type ResponseOrgInventoryChangeOpEnum string

const (
    ResponseOrgInventoryChangeOpEnum_ASSIGN         ResponseOrgInventoryChangeOpEnum = "assign"
    ResponseOrgInventoryChangeOpEnum_DELETE         ResponseOrgInventoryChangeOpEnum = "delete"
    ResponseOrgInventoryChangeOpEnum_DOWNGRADETOJSI ResponseOrgInventoryChangeOpEnum = "downgrade_to_jsi"
    ResponseOrgInventoryChangeOpEnum_UNASSIGN       ResponseOrgInventoryChangeOpEnum = "unassign"
    ResponseOrgInventoryChangeOpEnum_UPGRADETOMIST  ResponseOrgInventoryChangeOpEnum = "upgrade_to_mist"
)

// RfClientTypeEnum is a string enum.
// enum: `asset`, `client`, `sdkclient`
type RfClientTypeEnum string

const (
    RfClientTypeEnum_ASSET     RfClientTypeEnum = "asset"
    RfClientTypeEnum_CLIENT    RfClientTypeEnum = "client"
    RfClientTypeEnum_SDKCLIENT RfClientTypeEnum = "sdkclient"
)

// RogueTypeEnum is a string enum.
// enum: `honeypot`, `lan`, `others`, `spoof`
type RogueTypeEnum string

const (
    RogueTypeEnum_HONEYPOT RogueTypeEnum = "honeypot"
    RogueTypeEnum_LAN      RogueTypeEnum = "lan"
    RogueTypeEnum_OTHERS   RogueTypeEnum = "others"
    RogueTypeEnum_SPOOF    RogueTypeEnum = "spoof"
)

// RrmEventPreBandwidthEnum is a int enum.
// (previously) channel width for the band , 0 means no previously available. enum: `0`, `20`, `40`, `80`, `160`
type RrmEventPreBandwidthEnum int

const (
    RrmEventPreBandwidthEnum_ENUM0   RrmEventPreBandwidthEnum = 0
    RrmEventPreBandwidthEnum_ENUM20  RrmEventPreBandwidthEnum = 20
    RrmEventPreBandwidthEnum_ENUM40  RrmEventPreBandwidthEnum = 40
    RrmEventPreBandwidthEnum_ENUM80  RrmEventPreBandwidthEnum = 80
    RrmEventPreBandwidthEnum_ENUM160 RrmEventPreBandwidthEnum = 160
)

// RrmEventTypeEnum is a string enum.
// enum: `interference-ap-co-channel`, `interference-ap-non-wifi`, `neighbor-ap-down`, `neighbor-ap-recovered`, `radar-detected`, `rrm-radar`, `scheduled-site_rrm`, `triggered-site_rrm`
type RrmEventTypeEnum string

const (
    RrmEventTypeEnum_INTERFERENCEAPCOCHANNEL RrmEventTypeEnum = "interference-ap-co-channel"
    RrmEventTypeEnum_INTERFERENCEAPNONWIFI   RrmEventTypeEnum = "interference-ap-non-wifi"
    RrmEventTypeEnum_NEIGHBORAPDOWN          RrmEventTypeEnum = "neighbor-ap-down"
    RrmEventTypeEnum_NEIGHBORAPRECOVERED     RrmEventTypeEnum = "neighbor-ap-recovered"
    RrmEventTypeEnum_RADARDETECTED           RrmEventTypeEnum = "radar-detected"
    RrmEventTypeEnum_RRMRADAR                RrmEventTypeEnum = "rrm-radar"
    RrmEventTypeEnum_SCHEDULEDSITERRM        RrmEventTypeEnum = "scheduled-site_rrm"
    RrmEventTypeEnum_TRIGGEREDSITERRM        RrmEventTypeEnum = "triggered-site_rrm"
)

// RrmStatusEnum is a string enum.
// enum: `ready`, `unknown`, `updating`
type RrmStatusEnum string

const (
    RrmStatusEnum_READY    RrmStatusEnum = "ready"
    RrmStatusEnum_UNKNOWN  RrmStatusEnum = "unknown"
    RrmStatusEnum_UPDATING RrmStatusEnum = "updating"
)

// ScanDataItemBandEnum is a string enum.
// 5GHz or 2.4GHz band, associated with the BSSID scanned. enum: `2.4`, `5`
type ScanDataItemBandEnum string

const (
    ScanDataItemBandEnum_ENUM24 ScanDataItemBandEnum = "2.4"
    ScanDataItemBandEnum_ENUM5  ScanDataItemBandEnum = "5"
)

// SearchOrgDevicesMxtunnelStatusEnum is a string enum.
// enum: `down`, `up`
type SearchOrgDevicesMxtunnelStatusEnum string

const (
    SearchOrgDevicesMxtunnelStatusEnum_DOWN SearchOrgDevicesMxtunnelStatusEnum = "down"
    SearchOrgDevicesMxtunnelStatusEnum_UP   SearchOrgDevicesMxtunnelStatusEnum = "up"
)

// SearchOrgSwOrGwPortsAuthStateEnum is a string enum.
// enum: `authenticated`, `authenticating`, `held`, `init`
type SearchOrgSwOrGwPortsAuthStateEnum string

const (
    SearchOrgSwOrGwPortsAuthStateEnum_AUTHENTICATED  SearchOrgSwOrGwPortsAuthStateEnum = "authenticated"
    SearchOrgSwOrGwPortsAuthStateEnum_AUTHENTICATING SearchOrgSwOrGwPortsAuthStateEnum = "authenticating"
    SearchOrgSwOrGwPortsAuthStateEnum_HELD           SearchOrgSwOrGwPortsAuthStateEnum = "held"
    SearchOrgSwOrGwPortsAuthStateEnum_INIT           SearchOrgSwOrGwPortsAuthStateEnum = "init"
)

// SearchOrgSwOrGwPortsStpRoleEnum is a string enum.
// enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
type SearchOrgSwOrGwPortsStpRoleEnum string

const (
    SearchOrgSwOrGwPortsStpRoleEnum_ALTERNATE     SearchOrgSwOrGwPortsStpRoleEnum = "alternate"
    SearchOrgSwOrGwPortsStpRoleEnum_BACKUP        SearchOrgSwOrGwPortsStpRoleEnum = "backup"
    SearchOrgSwOrGwPortsStpRoleEnum_DESIGNATED    SearchOrgSwOrGwPortsStpRoleEnum = "designated"
    SearchOrgSwOrGwPortsStpRoleEnum_ROOT          SearchOrgSwOrGwPortsStpRoleEnum = "root"
    SearchOrgSwOrGwPortsStpRoleEnum_ROOTPREVENTED SearchOrgSwOrGwPortsStpRoleEnum = "root-prevented"
)

// SearchOrgSwOrGwPortsStpStateEnum is a string enum.
// enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
type SearchOrgSwOrGwPortsStpStateEnum string

const (
    SearchOrgSwOrGwPortsStpStateEnum_BLOCKING   SearchOrgSwOrGwPortsStpStateEnum = "blocking"
    SearchOrgSwOrGwPortsStpStateEnum_DISABLED   SearchOrgSwOrGwPortsStpStateEnum = "disabled"
    SearchOrgSwOrGwPortsStpStateEnum_FORWARDING SearchOrgSwOrGwPortsStpStateEnum = "forwarding"
    SearchOrgSwOrGwPortsStpStateEnum_LEARNING   SearchOrgSwOrGwPortsStpStateEnum = "learning"
    SearchOrgSwOrGwPortsStpStateEnum_LISTENING  SearchOrgSwOrGwPortsStpStateEnum = "listening"
)

// SearchSiteDevicesDescSortEnum is a string enum.
// enum: `mac`, `model`, `sku`, `timestamp`
type SearchSiteDevicesDescSortEnum string

const (
    SearchSiteDevicesDescSortEnum_MAC       SearchSiteDevicesDescSortEnum = "mac"
    SearchSiteDevicesDescSortEnum_MODEL     SearchSiteDevicesDescSortEnum = "model"
    SearchSiteDevicesDescSortEnum_SKU       SearchSiteDevicesDescSortEnum = "sku"
    SearchSiteDevicesDescSortEnum_TIMESTAMP SearchSiteDevicesDescSortEnum = "timestamp"
)

// SearchSiteDevicesMxtunnelStatusEnum is a string enum.
// enum: `down`, `up`
type SearchSiteDevicesMxtunnelStatusEnum string

const (
    SearchSiteDevicesMxtunnelStatusEnum_DOWN SearchSiteDevicesMxtunnelStatusEnum = "down"
    SearchSiteDevicesMxtunnelStatusEnum_UP   SearchSiteDevicesMxtunnelStatusEnum = "up"
)

// SearchSiteDevicesSortEnum is a string enum.
// enum: `mac`, `model`, `sku`, `timestamp`
type SearchSiteDevicesSortEnum string

const (
    SearchSiteDevicesSortEnum_MAC       SearchSiteDevicesSortEnum = "mac"
    SearchSiteDevicesSortEnum_MODEL     SearchSiteDevicesSortEnum = "model"
    SearchSiteDevicesSortEnum_SKU       SearchSiteDevicesSortEnum = "sku"
    SearchSiteDevicesSortEnum_TIMESTAMP SearchSiteDevicesSortEnum = "timestamp"
)

// SearchSiteSwOrGwPortsAuthStateEnum is a string enum.
// enum: `authenticated`, `authenticating`, `held`, `init`
type SearchSiteSwOrGwPortsAuthStateEnum string

const (
    SearchSiteSwOrGwPortsAuthStateEnum_AUTHENTICATED  SearchSiteSwOrGwPortsAuthStateEnum = "authenticated"
    SearchSiteSwOrGwPortsAuthStateEnum_AUTHENTICATING SearchSiteSwOrGwPortsAuthStateEnum = "authenticating"
    SearchSiteSwOrGwPortsAuthStateEnum_HELD           SearchSiteSwOrGwPortsAuthStateEnum = "held"
    SearchSiteSwOrGwPortsAuthStateEnum_INIT           SearchSiteSwOrGwPortsAuthStateEnum = "init"
)

// SearchSiteSwOrGwPortsDeviceTypeEnum is a string enum.
// enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch`
type SearchSiteSwOrGwPortsDeviceTypeEnum string

const (
    SearchSiteSwOrGwPortsDeviceTypeEnum_AP         SearchSiteSwOrGwPortsDeviceTypeEnum = "ap"
    SearchSiteSwOrGwPortsDeviceTypeEnum_BLE        SearchSiteSwOrGwPortsDeviceTypeEnum = "ble"
    SearchSiteSwOrGwPortsDeviceTypeEnum_GATEWAY    SearchSiteSwOrGwPortsDeviceTypeEnum = "gateway"
    SearchSiteSwOrGwPortsDeviceTypeEnum_MXEDGE     SearchSiteSwOrGwPortsDeviceTypeEnum = "mxedge"
    SearchSiteSwOrGwPortsDeviceTypeEnum_NAC        SearchSiteSwOrGwPortsDeviceTypeEnum = "nac"
    SearchSiteSwOrGwPortsDeviceTypeEnum_ENUMSWITCH SearchSiteSwOrGwPortsDeviceTypeEnum = "switch"
)

// SearchSiteSwOrGwPortsStpRoleEnum is a string enum.
// enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
type SearchSiteSwOrGwPortsStpRoleEnum string

const (
    SearchSiteSwOrGwPortsStpRoleEnum_ALTERNATE     SearchSiteSwOrGwPortsStpRoleEnum = "alternate"
    SearchSiteSwOrGwPortsStpRoleEnum_BACKUP        SearchSiteSwOrGwPortsStpRoleEnum = "backup"
    SearchSiteSwOrGwPortsStpRoleEnum_DESIGNATED    SearchSiteSwOrGwPortsStpRoleEnum = "designated"
    SearchSiteSwOrGwPortsStpRoleEnum_ROOT          SearchSiteSwOrGwPortsStpRoleEnum = "root"
    SearchSiteSwOrGwPortsStpRoleEnum_ROOTPREVENTED SearchSiteSwOrGwPortsStpRoleEnum = "root-prevented"
)

// SearchSiteSwOrGwPortsStpStateEnum is a string enum.
// enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
type SearchSiteSwOrGwPortsStpStateEnum string

const (
    SearchSiteSwOrGwPortsStpStateEnum_BLOCKING   SearchSiteSwOrGwPortsStpStateEnum = "blocking"
    SearchSiteSwOrGwPortsStpStateEnum_DISABLED   SearchSiteSwOrGwPortsStpStateEnum = "disabled"
    SearchSiteSwOrGwPortsStpStateEnum_FORWARDING SearchSiteSwOrGwPortsStpStateEnum = "forwarding"
    SearchSiteSwOrGwPortsStpStateEnum_LEARNING   SearchSiteSwOrGwPortsStpStateEnum = "learning"
    SearchSiteSwOrGwPortsStpStateEnum_LISTENING  SearchSiteSwOrGwPortsStpStateEnum = "listening"
)

// SecintelProfileProfileActionEnum is a string enum.
// enum: `default`, `standard`, `strict`
type SecintelProfileProfileActionEnum string

const (
    SecintelProfileProfileActionEnum_ENUMDEFAULT SecintelProfileProfileActionEnum = "default"
    SecintelProfileProfileActionEnum_STANDARD    SecintelProfileProfileActionEnum = "standard"
    SecintelProfileProfileActionEnum_STRICT      SecintelProfileProfileActionEnum = "strict"
)

// SecintelProfileProfileCategoryEnum is a string enum.
// enum: `CC`, `IH` (Infected Host), `DNS`
type SecintelProfileProfileCategoryEnum string

const (
    SecintelProfileProfileCategoryEnum_CC  SecintelProfileProfileCategoryEnum = "CC"
    SecintelProfileProfileCategoryEnum_DNS SecintelProfileProfileCategoryEnum = "DNS"
    SecintelProfileProfileCategoryEnum_IH  SecintelProfileProfileCategoryEnum = "IH"
)

// ServiceFailoverPolicyEnum is a string enum.
// enum: `non_revertable`, `none`, `revertable`
type ServiceFailoverPolicyEnum string

const (
    ServiceFailoverPolicyEnum_NONREVERTABLE ServiceFailoverPolicyEnum = "non_revertable"
    ServiceFailoverPolicyEnum_NONE          ServiceFailoverPolicyEnum = "none"
    ServiceFailoverPolicyEnum_REVERTABLE    ServiceFailoverPolicyEnum = "revertable"
)

// ServicePolicyEwfRuleProfileEnum is a string enum.
// enum: `critical`, `standard`, `strict`
type ServicePolicyEwfRuleProfileEnum string

const (
    ServicePolicyEwfRuleProfileEnum_CRITICAL ServicePolicyEwfRuleProfileEnum = "critical"
    ServicePolicyEwfRuleProfileEnum_STANDARD ServicePolicyEwfRuleProfileEnum = "standard"
    ServicePolicyEwfRuleProfileEnum_STRICT   ServicePolicyEwfRuleProfileEnum = "strict"
)

// ServiceTrafficClassEnum is a string enum.
// when `traffic_type`==`custom`. enum: `best_effort`, `high`, `low`, `medium`
type ServiceTrafficClassEnum string

const (
    ServiceTrafficClassEnum_BESTEFFORT ServiceTrafficClassEnum = "best_effort"
    ServiceTrafficClassEnum_HIGH       ServiceTrafficClassEnum = "high"
    ServiceTrafficClassEnum_LOW        ServiceTrafficClassEnum = "low"
    ServiceTrafficClassEnum_MEDIUM     ServiceTrafficClassEnum = "medium"
)

// ServiceTypeEnum is a string enum.
// enum: `app_categories`, `apps`, `custom`, `urls`
type ServiceTypeEnum string

const (
    ServiceTypeEnum_APPCATEGORIES ServiceTypeEnum = "app_categories"
    ServiceTypeEnum_APPS          ServiceTypeEnum = "apps"
    ServiceTypeEnum_CUSTOM        ServiceTypeEnum = "custom"
    ServiceTypeEnum_URLS          ServiceTypeEnum = "urls"
)

// SiteAppsCountDistinctEnum is a string enum.
// enum: `ap`, `app`, `category`, `device_mac`, `port_id`, `service`, `src_ip`, `ssid`, `wcid`, `wlan_id app`
type SiteAppsCountDistinctEnum string

const (
    SiteAppsCountDistinctEnum_AP            SiteAppsCountDistinctEnum = "ap"
    SiteAppsCountDistinctEnum_APP           SiteAppsCountDistinctEnum = "app"
    SiteAppsCountDistinctEnum_CATEGORY      SiteAppsCountDistinctEnum = "category"
    SiteAppsCountDistinctEnum_DEVICEMAC     SiteAppsCountDistinctEnum = "device_mac"
    SiteAppsCountDistinctEnum_PORTID        SiteAppsCountDistinctEnum = "port_id"
    SiteAppsCountDistinctEnum_SERVICE       SiteAppsCountDistinctEnum = "service"
    SiteAppsCountDistinctEnum_SRCIP         SiteAppsCountDistinctEnum = "src_ip"
    SiteAppsCountDistinctEnum_SSID          SiteAppsCountDistinctEnum = "ssid"
    SiteAppsCountDistinctEnum_WCID          SiteAppsCountDistinctEnum = "wcid"
    SiteAppsCountDistinctEnum_ENUMWLANIDAPP SiteAppsCountDistinctEnum = "wlan_id app"
)

// SiteAssetsCountDistinctEnum is a string enum.
// enum: `by`, `device_name`, `eddystone_uid_instance`, `eddystone_uid_namespace`, `eddystone_url`, `ibeacon_major`, `ibeacon_minor`, `ibeacon_uuid`, `mac`, `map_id`, `name`
type SiteAssetsCountDistinctEnum string

const (
    SiteAssetsCountDistinctEnum_BY                    SiteAssetsCountDistinctEnum = "by"
    SiteAssetsCountDistinctEnum_DEVICENAME            SiteAssetsCountDistinctEnum = "device_name"
    SiteAssetsCountDistinctEnum_EDDYSTONEUIDINSTANCE  SiteAssetsCountDistinctEnum = "eddystone_uid_instance"
    SiteAssetsCountDistinctEnum_EDDYSTONEUIDNAMESPACE SiteAssetsCountDistinctEnum = "eddystone_uid_namespace"
    SiteAssetsCountDistinctEnum_EDDYSTONEURL          SiteAssetsCountDistinctEnum = "eddystone_url"
    SiteAssetsCountDistinctEnum_IBEACONMAJOR          SiteAssetsCountDistinctEnum = "ibeacon_major"
    SiteAssetsCountDistinctEnum_IBEACONMINOR          SiteAssetsCountDistinctEnum = "ibeacon_minor"
    SiteAssetsCountDistinctEnum_IBEACONUUID           SiteAssetsCountDistinctEnum = "ibeacon_uuid"
    SiteAssetsCountDistinctEnum_MAC                   SiteAssetsCountDistinctEnum = "mac"
    SiteAssetsCountDistinctEnum_MAPID                 SiteAssetsCountDistinctEnum = "map_id"
    SiteAssetsCountDistinctEnum_NAME                  SiteAssetsCountDistinctEnum = "name"
)

// SiteAutoUpgradeVersionEnum is a string enum.
// desired version. enum: `beta`, `custom`, `stable`
type SiteAutoUpgradeVersionEnum string

const (
    SiteAutoUpgradeVersionEnum_BETA   SiteAutoUpgradeVersionEnum = "beta"
    SiteAutoUpgradeVersionEnum_CUSTOM SiteAutoUpgradeVersionEnum = "custom"
    SiteAutoUpgradeVersionEnum_STABLE SiteAutoUpgradeVersionEnum = "stable"
)

// SiteClientEventsCountDistinctEnum is a string enum.
// enum: `band`, `channel`, `proto`, `ssid`, `type`, `wlan_id`
type SiteClientEventsCountDistinctEnum string

const (
    SiteClientEventsCountDistinctEnum_BAND     SiteClientEventsCountDistinctEnum = "band"
    SiteClientEventsCountDistinctEnum_CHANNEL  SiteClientEventsCountDistinctEnum = "channel"
    SiteClientEventsCountDistinctEnum_PROTO    SiteClientEventsCountDistinctEnum = "proto"
    SiteClientEventsCountDistinctEnum_SSID     SiteClientEventsCountDistinctEnum = "ssid"
    SiteClientEventsCountDistinctEnum_ENUMTYPE SiteClientEventsCountDistinctEnum = "type"
    SiteClientEventsCountDistinctEnum_WLANID   SiteClientEventsCountDistinctEnum = "wlan_id"
)

// SiteClientSessionsCountDistinctEnum is a string enum.
// enum: `ap`, `client_family`, `client_manufacture`, `client_model`, `client_os`, `mac`, `ssid`, `wlan_id`
type SiteClientSessionsCountDistinctEnum string

const (
    SiteClientSessionsCountDistinctEnum_AP                SiteClientSessionsCountDistinctEnum = "ap"
    SiteClientSessionsCountDistinctEnum_CLIENTFAMILY      SiteClientSessionsCountDistinctEnum = "client_family"
    SiteClientSessionsCountDistinctEnum_CLIENTMANUFACTURE SiteClientSessionsCountDistinctEnum = "client_manufacture"
    SiteClientSessionsCountDistinctEnum_CLIENTMODEL       SiteClientSessionsCountDistinctEnum = "client_model"
    SiteClientSessionsCountDistinctEnum_CLIENTOS          SiteClientSessionsCountDistinctEnum = "client_os"
    SiteClientSessionsCountDistinctEnum_MAC               SiteClientSessionsCountDistinctEnum = "mac"
    SiteClientSessionsCountDistinctEnum_SSID              SiteClientSessionsCountDistinctEnum = "ssid"
    SiteClientSessionsCountDistinctEnum_WLANID            SiteClientSessionsCountDistinctEnum = "wlan_id"
)

// SiteClientsCountDistinctEnum is a string enum.
// enum: `ap`, `device`, `hostname`, `ip`, `model`, `os`, `ssid`, `vlan`
type SiteClientsCountDistinctEnum string

const (
    SiteClientsCountDistinctEnum_AP       SiteClientsCountDistinctEnum = "ap"
    SiteClientsCountDistinctEnum_DEVICE   SiteClientsCountDistinctEnum = "device"
    SiteClientsCountDistinctEnum_HOSTNAME SiteClientsCountDistinctEnum = "hostname"
    SiteClientsCountDistinctEnum_IP       SiteClientsCountDistinctEnum = "ip"
    SiteClientsCountDistinctEnum_MODEL    SiteClientsCountDistinctEnum = "model"
    SiteClientsCountDistinctEnum_OS       SiteClientsCountDistinctEnum = "os"
    SiteClientsCountDistinctEnum_SSID     SiteClientsCountDistinctEnum = "ssid"
    SiteClientsCountDistinctEnum_VLAN     SiteClientsCountDistinctEnum = "vlan"
)

// SiteDeviceEventsCountDistinctEnum is a string enum.
// enum: `mac`, `model`, `type`, `type_code`
type SiteDeviceEventsCountDistinctEnum string

const (
    SiteDeviceEventsCountDistinctEnum_MAC      SiteDeviceEventsCountDistinctEnum = "mac"
    SiteDeviceEventsCountDistinctEnum_MODEL    SiteDeviceEventsCountDistinctEnum = "model"
    SiteDeviceEventsCountDistinctEnum_ENUMTYPE SiteDeviceEventsCountDistinctEnum = "type"
    SiteDeviceEventsCountDistinctEnum_TYPECODE SiteDeviceEventsCountDistinctEnum = "type_code"
)

// SiteDeviceLastConfigCountDistinctEnum is a string enum.
// enum: `mac`, `name`, `site_id`, `version`
type SiteDeviceLastConfigCountDistinctEnum string

const (
    SiteDeviceLastConfigCountDistinctEnum_MAC     SiteDeviceLastConfigCountDistinctEnum = "mac"
    SiteDeviceLastConfigCountDistinctEnum_NAME    SiteDeviceLastConfigCountDistinctEnum = "name"
    SiteDeviceLastConfigCountDistinctEnum_SITEID  SiteDeviceLastConfigCountDistinctEnum = "site_id"
    SiteDeviceLastConfigCountDistinctEnum_VERSION SiteDeviceLastConfigCountDistinctEnum = "version"
)

// SiteDevicesCountDistinctEnum is a string enum.
// enum: `hostname`, `lldp_mgmt_addr`, `lldp_port_id`, `lldp_system_desc`, `lldp_system_name`, `map_id`, `model`, `mxedge_id`, `mxtunnel_status`, `version`
type SiteDevicesCountDistinctEnum string

const (
    SiteDevicesCountDistinctEnum_HOSTNAME       SiteDevicesCountDistinctEnum = "hostname"
    SiteDevicesCountDistinctEnum_LLDPMGMTADDR   SiteDevicesCountDistinctEnum = "lldp_mgmt_addr"
    SiteDevicesCountDistinctEnum_LLDPPORTID     SiteDevicesCountDistinctEnum = "lldp_port_id"
    SiteDevicesCountDistinctEnum_LLDPSYSTEMDESC SiteDevicesCountDistinctEnum = "lldp_system_desc"
    SiteDevicesCountDistinctEnum_LLDPSYSTEMNAME SiteDevicesCountDistinctEnum = "lldp_system_name"
    SiteDevicesCountDistinctEnum_MAPID          SiteDevicesCountDistinctEnum = "map_id"
    SiteDevicesCountDistinctEnum_MODEL          SiteDevicesCountDistinctEnum = "model"
    SiteDevicesCountDistinctEnum_MXEDGEID       SiteDevicesCountDistinctEnum = "mxedge_id"
    SiteDevicesCountDistinctEnum_MXTUNNELSTATUS SiteDevicesCountDistinctEnum = "mxtunnel_status"
    SiteDevicesCountDistinctEnum_VERSION        SiteDevicesCountDistinctEnum = "version"
)

// SiteDiscoveredSwitchesCountDistinctEnum is a string enum.
// enum: `mgmt_addr`, `model`, `system_name`, `version`
type SiteDiscoveredSwitchesCountDistinctEnum string

const (
    SiteDiscoveredSwitchesCountDistinctEnum_MGMTADDR   SiteDiscoveredSwitchesCountDistinctEnum = "mgmt_addr"
    SiteDiscoveredSwitchesCountDistinctEnum_MODEL      SiteDiscoveredSwitchesCountDistinctEnum = "model"
    SiteDiscoveredSwitchesCountDistinctEnum_SYSTEMNAME SiteDiscoveredSwitchesCountDistinctEnum = "system_name"
    SiteDiscoveredSwitchesCountDistinctEnum_VERSION    SiteDiscoveredSwitchesCountDistinctEnum = "version"
)

// SiteGuestsCountDistinctEnum is a string enum.
// enum: `auth_method`, `company`, `ssid`
type SiteGuestsCountDistinctEnum string

const (
    SiteGuestsCountDistinctEnum_AUTHMETHOD SiteGuestsCountDistinctEnum = "auth_method"
    SiteGuestsCountDistinctEnum_COMPANY    SiteGuestsCountDistinctEnum = "company"
    SiteGuestsCountDistinctEnum_SSID       SiteGuestsCountDistinctEnum = "ssid"
)

// SiteMxedgeEventsCountDistinctEnum is a string enum.
// enum: `mxcluster_id`, `mxedge_id`, `package`, `type`
type SiteMxedgeEventsCountDistinctEnum string

const (
    SiteMxedgeEventsCountDistinctEnum_MXCLUSTERID SiteMxedgeEventsCountDistinctEnum = "mxcluster_id"
    SiteMxedgeEventsCountDistinctEnum_MXEDGEID    SiteMxedgeEventsCountDistinctEnum = "mxedge_id"
    SiteMxedgeEventsCountDistinctEnum_ENUMPACKAGE SiteMxedgeEventsCountDistinctEnum = "package"
    SiteMxedgeEventsCountDistinctEnum_ENUMTYPE    SiteMxedgeEventsCountDistinctEnum = "type"
)

// SiteMxtunnelProtocolEnum is a string enum.
// enum: `ip`, `udp`
type SiteMxtunnelProtocolEnum string

const (
    SiteMxtunnelProtocolEnum_IP  SiteMxtunnelProtocolEnum = "ip"
    SiteMxtunnelProtocolEnum_UDP SiteMxtunnelProtocolEnum = "udp"
)

// SiteNacClientEventsCountDistinctEnum is a string enum.
// enum: `ap`, `auth_type`, `dryrun_nacrule_id`, `mac`, `nacrule_id`, `nas_vendor`, `ssid`, `type`, `username`, `vlan`
type SiteNacClientEventsCountDistinctEnum string

const (
    SiteNacClientEventsCountDistinctEnum_AP              SiteNacClientEventsCountDistinctEnum = "ap"
    SiteNacClientEventsCountDistinctEnum_AUTHTYPE        SiteNacClientEventsCountDistinctEnum = "auth_type"
    SiteNacClientEventsCountDistinctEnum_DRYRUNNACRULEID SiteNacClientEventsCountDistinctEnum = "dryrun_nacrule_id"
    SiteNacClientEventsCountDistinctEnum_MAC             SiteNacClientEventsCountDistinctEnum = "mac"
    SiteNacClientEventsCountDistinctEnum_NACRULEID       SiteNacClientEventsCountDistinctEnum = "nacrule_id"
    SiteNacClientEventsCountDistinctEnum_NASVENDOR       SiteNacClientEventsCountDistinctEnum = "nas_vendor"
    SiteNacClientEventsCountDistinctEnum_SSID            SiteNacClientEventsCountDistinctEnum = "ssid"
    SiteNacClientEventsCountDistinctEnum_ENUMTYPE        SiteNacClientEventsCountDistinctEnum = "type"
    SiteNacClientEventsCountDistinctEnum_USERNAME        SiteNacClientEventsCountDistinctEnum = "username"
    SiteNacClientEventsCountDistinctEnum_VLAN            SiteNacClientEventsCountDistinctEnum = "vlan"
)

// SiteNacClientsCountDistinctEnum is a string enum.
// enum: `auth_type`, `last_ap`, `last_nacrule_id`, `last_nas_vendor`, `last_ssid`, `last_status`, `last_username`, `last_vlan`, `mac`, `mdm_compliance`, `mdm_provider`, `type`
type SiteNacClientsCountDistinctEnum string

const (
    SiteNacClientsCountDistinctEnum_AUTHTYPE      SiteNacClientsCountDistinctEnum = "auth_type"
    SiteNacClientsCountDistinctEnum_LASTAP        SiteNacClientsCountDistinctEnum = "last_ap"
    SiteNacClientsCountDistinctEnum_LASTNACRULEID SiteNacClientsCountDistinctEnum = "last_nacrule_id"
    SiteNacClientsCountDistinctEnum_LASTNASVENDOR SiteNacClientsCountDistinctEnum = "last_nas_vendor"
    SiteNacClientsCountDistinctEnum_LASTSSID      SiteNacClientsCountDistinctEnum = "last_ssid"
    SiteNacClientsCountDistinctEnum_LASTSTATUS    SiteNacClientsCountDistinctEnum = "last_status"
    SiteNacClientsCountDistinctEnum_LASTUSERNAME  SiteNacClientsCountDistinctEnum = "last_username"
    SiteNacClientsCountDistinctEnum_LASTVLAN      SiteNacClientsCountDistinctEnum = "last_vlan"
    SiteNacClientsCountDistinctEnum_MAC           SiteNacClientsCountDistinctEnum = "mac"
    SiteNacClientsCountDistinctEnum_MDMCOMPLIANCE SiteNacClientsCountDistinctEnum = "mdm_compliance"
    SiteNacClientsCountDistinctEnum_MDMPROVIDER   SiteNacClientsCountDistinctEnum = "mdm_provider"
    SiteNacClientsCountDistinctEnum_ENUMTYPE      SiteNacClientsCountDistinctEnum = "type"
)

// SiteOtherDeviceEventsCountDistinctEnum is a string enum.
// enum: `mac`, `site_id`, `type`, `vendor`
type SiteOtherDeviceEventsCountDistinctEnum string

const (
    SiteOtherDeviceEventsCountDistinctEnum_MAC      SiteOtherDeviceEventsCountDistinctEnum = "mac"
    SiteOtherDeviceEventsCountDistinctEnum_SITEID   SiteOtherDeviceEventsCountDistinctEnum = "site_id"
    SiteOtherDeviceEventsCountDistinctEnum_ENUMTYPE SiteOtherDeviceEventsCountDistinctEnum = "type"
    SiteOtherDeviceEventsCountDistinctEnum_VENDOR   SiteOtherDeviceEventsCountDistinctEnum = "vendor"
)

// SitePortsCountDistinctEnum is a string enum.
// enum: `full_duplex`, `mac`, `neighbor_mac`, `neighbor_port_desc`, `neighbor_system_name`, `poe_disabled`, `poe_mode`, `poe_on`, `port_id`, `port_mac`, `speed`, `up`
type SitePortsCountDistinctEnum string

const (
    SitePortsCountDistinctEnum_FULLDUPLEX         SitePortsCountDistinctEnum = "full_duplex"
    SitePortsCountDistinctEnum_MAC                SitePortsCountDistinctEnum = "mac"
    SitePortsCountDistinctEnum_NEIGHBORMAC        SitePortsCountDistinctEnum = "neighbor_mac"
    SitePortsCountDistinctEnum_NEIGHBORPORTDESC   SitePortsCountDistinctEnum = "neighbor_port_desc"
    SitePortsCountDistinctEnum_NEIGHBORSYSTEMNAME SitePortsCountDistinctEnum = "neighbor_system_name"
    SitePortsCountDistinctEnum_POEDISABLED        SitePortsCountDistinctEnum = "poe_disabled"
    SitePortsCountDistinctEnum_POEMODE            SitePortsCountDistinctEnum = "poe_mode"
    SitePortsCountDistinctEnum_POEON              SitePortsCountDistinctEnum = "poe_on"
    SitePortsCountDistinctEnum_PORTID             SitePortsCountDistinctEnum = "port_id"
    SitePortsCountDistinctEnum_PORTMAC            SitePortsCountDistinctEnum = "port_mac"
    SitePortsCountDistinctEnum_SPEED              SitePortsCountDistinctEnum = "speed"
    SitePortsCountDistinctEnum_UP                 SitePortsCountDistinctEnum = "up"
)

// SiteRogueEventsCountDistinctEnum is a string enum.
// enum: `ap`, `bssid`, `ssid`, `type`
type SiteRogueEventsCountDistinctEnum string

const (
    SiteRogueEventsCountDistinctEnum_AP       SiteRogueEventsCountDistinctEnum = "ap"
    SiteRogueEventsCountDistinctEnum_BSSID    SiteRogueEventsCountDistinctEnum = "bssid"
    SiteRogueEventsCountDistinctEnum_SSID     SiteRogueEventsCountDistinctEnum = "ssid"
    SiteRogueEventsCountDistinctEnum_ENUMTYPE SiteRogueEventsCountDistinctEnum = "type"
)

// SiteServiceEventsCountDistinctEnum is a string enum.
// enum: `mac`, `model`, `policy`, `port_id`, `site_id`, `type`, `vpn_name`, `vpn_path`
type SiteServiceEventsCountDistinctEnum string

const (
    SiteServiceEventsCountDistinctEnum_MAC      SiteServiceEventsCountDistinctEnum = "mac"
    SiteServiceEventsCountDistinctEnum_MODEL    SiteServiceEventsCountDistinctEnum = "model"
    SiteServiceEventsCountDistinctEnum_POLICY   SiteServiceEventsCountDistinctEnum = "policy"
    SiteServiceEventsCountDistinctEnum_PORTID   SiteServiceEventsCountDistinctEnum = "port_id"
    SiteServiceEventsCountDistinctEnum_SITEID   SiteServiceEventsCountDistinctEnum = "site_id"
    SiteServiceEventsCountDistinctEnum_ENUMTYPE SiteServiceEventsCountDistinctEnum = "type"
    SiteServiceEventsCountDistinctEnum_VPNNAME  SiteServiceEventsCountDistinctEnum = "vpn_name"
    SiteServiceEventsCountDistinctEnum_VPNPATH  SiteServiceEventsCountDistinctEnum = "vpn_path"
)

// SystemDefinedPortUsagesEnum is a string enum.
// system-default port usages. enum: `ap`, `iot`, `uplink``
type SystemDefinedPortUsagesEnum string

const (
    SystemDefinedPortUsagesEnum_AP     SystemDefinedPortUsagesEnum = "ap"
    SystemDefinedPortUsagesEnum_IOT    SystemDefinedPortUsagesEnum = "iot"
    SystemDefinedPortUsagesEnum_UPLINK SystemDefinedPortUsagesEnum = "uplink"
)

// SiteSkyAtpEventsCountDistinctEnum is a string enum.
// enum: `device_mac`, `mac`, `threat_level`, `type`
type SiteSkyAtpEventsCountDistinctEnum string

const (
    SiteSkyAtpEventsCountDistinctEnum_DEVICEMAC   SiteSkyAtpEventsCountDistinctEnum = "device_mac"
    SiteSkyAtpEventsCountDistinctEnum_MAC         SiteSkyAtpEventsCountDistinctEnum = "mac"
    SiteSkyAtpEventsCountDistinctEnum_THREATLEVEL SiteSkyAtpEventsCountDistinctEnum = "threat_level"
    SiteSkyAtpEventsCountDistinctEnum_ENUMTYPE    SiteSkyAtpEventsCountDistinctEnum = "type"
)

// SiteSleHistogramScopeParametersEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleHistogramScopeParametersEnum string

const (
    SiteSleHistogramScopeParametersEnum_AP         SiteSleHistogramScopeParametersEnum = "ap"
    SiteSleHistogramScopeParametersEnum_CLIENT     SiteSleHistogramScopeParametersEnum = "client"
    SiteSleHistogramScopeParametersEnum_GATEWAY    SiteSleHistogramScopeParametersEnum = "gateway"
    SiteSleHistogramScopeParametersEnum_SITE       SiteSleHistogramScopeParametersEnum = "site"
    SiteSleHistogramScopeParametersEnum_ENUMSWITCH SiteSleHistogramScopeParametersEnum = "switch"
)

// SiteSleImpactSummaryFieldsParameterEnum is a string enum.
// enum: `ap`, `band`, `chassis`, `client`, `device_os`, `device_type`, `gateway`, `gateway_zones`, `interface`, `mxedge`, `peer_path`, `server`, `switch`, `vlan`, `wlan`
type SiteSleImpactSummaryFieldsParameterEnum string

const (
    SiteSleImpactSummaryFieldsParameterEnum_AP            SiteSleImpactSummaryFieldsParameterEnum = "ap"
    SiteSleImpactSummaryFieldsParameterEnum_BAND          SiteSleImpactSummaryFieldsParameterEnum = "band"
    SiteSleImpactSummaryFieldsParameterEnum_CHASSIS       SiteSleImpactSummaryFieldsParameterEnum = "chassis"
    SiteSleImpactSummaryFieldsParameterEnum_CLIENT        SiteSleImpactSummaryFieldsParameterEnum = "client"
    SiteSleImpactSummaryFieldsParameterEnum_DEVICEOS      SiteSleImpactSummaryFieldsParameterEnum = "device_os"
    SiteSleImpactSummaryFieldsParameterEnum_DEVICETYPE    SiteSleImpactSummaryFieldsParameterEnum = "device_type"
    SiteSleImpactSummaryFieldsParameterEnum_GATEWAY       SiteSleImpactSummaryFieldsParameterEnum = "gateway"
    SiteSleImpactSummaryFieldsParameterEnum_GATEWAYZONES  SiteSleImpactSummaryFieldsParameterEnum = "gateway_zones"
    SiteSleImpactSummaryFieldsParameterEnum_ENUMINTERFACE SiteSleImpactSummaryFieldsParameterEnum = "interface"
    SiteSleImpactSummaryFieldsParameterEnum_MXEDGE        SiteSleImpactSummaryFieldsParameterEnum = "mxedge"
    SiteSleImpactSummaryFieldsParameterEnum_PEERPATH      SiteSleImpactSummaryFieldsParameterEnum = "peer_path"
    SiteSleImpactSummaryFieldsParameterEnum_SERVER        SiteSleImpactSummaryFieldsParameterEnum = "server"
    SiteSleImpactSummaryFieldsParameterEnum_ENUMSWITCH    SiteSleImpactSummaryFieldsParameterEnum = "switch"
    SiteSleImpactSummaryFieldsParameterEnum_VLAN          SiteSleImpactSummaryFieldsParameterEnum = "vlan"
    SiteSleImpactSummaryFieldsParameterEnum_WLAN          SiteSleImpactSummaryFieldsParameterEnum = "wlan"
)

// SiteSleImpactSummaryScopeParametersEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleImpactSummaryScopeParametersEnum string

const (
    SiteSleImpactSummaryScopeParametersEnum_AP         SiteSleImpactSummaryScopeParametersEnum = "ap"
    SiteSleImpactSummaryScopeParametersEnum_CLIENT     SiteSleImpactSummaryScopeParametersEnum = "client"
    SiteSleImpactSummaryScopeParametersEnum_GATEWAY    SiteSleImpactSummaryScopeParametersEnum = "gateway"
    SiteSleImpactSummaryScopeParametersEnum_SITE       SiteSleImpactSummaryScopeParametersEnum = "site"
    SiteSleImpactSummaryScopeParametersEnum_ENUMSWITCH SiteSleImpactSummaryScopeParametersEnum = "switch"
)

// SiteSleImpactedApsScopeParametersEnum is a string enum.
// enum: `site`
type SiteSleImpactedApsScopeParametersEnum string

const (
    SiteSleImpactedApsScopeParametersEnum_SITE SiteSleImpactedApsScopeParametersEnum = "site"
)

// SiteSleImpactedChassisScopeParametersEnum is a string enum.
// enum: `gateway`, `site`, `switch`
type SiteSleImpactedChassisScopeParametersEnum string

const (
    SiteSleImpactedChassisScopeParametersEnum_GATEWAY    SiteSleImpactedChassisScopeParametersEnum = "gateway"
    SiteSleImpactedChassisScopeParametersEnum_SITE       SiteSleImpactedChassisScopeParametersEnum = "site"
    SiteSleImpactedChassisScopeParametersEnum_ENUMSWITCH SiteSleImpactedChassisScopeParametersEnum = "switch"
)

// SiteSleImpactedClientsScopeParametersEnum is a string enum.
// enum: `gateway`, `site`, `switch`
type SiteSleImpactedClientsScopeParametersEnum string

const (
    SiteSleImpactedClientsScopeParametersEnum_GATEWAY    SiteSleImpactedClientsScopeParametersEnum = "gateway"
    SiteSleImpactedClientsScopeParametersEnum_SITE       SiteSleImpactedClientsScopeParametersEnum = "site"
    SiteSleImpactedClientsScopeParametersEnum_ENUMSWITCH SiteSleImpactedClientsScopeParametersEnum = "switch"
)

// SiteSleImpactedGatewaysScopeParametersEnum is a string enum.
// enum: `site`
type SiteSleImpactedGatewaysScopeParametersEnum string

const (
    SiteSleImpactedGatewaysScopeParametersEnum_SITE SiteSleImpactedGatewaysScopeParametersEnum = "site"
)

// SiteSleImpactedInterfacesScopeParametersEnum is a string enum.
// enum: `gateway`, `site`, `switch`
type SiteSleImpactedInterfacesScopeParametersEnum string

const (
    SiteSleImpactedInterfacesScopeParametersEnum_GATEWAY    SiteSleImpactedInterfacesScopeParametersEnum = "gateway"
    SiteSleImpactedInterfacesScopeParametersEnum_SITE       SiteSleImpactedInterfacesScopeParametersEnum = "site"
    SiteSleImpactedInterfacesScopeParametersEnum_ENUMSWITCH SiteSleImpactedInterfacesScopeParametersEnum = "switch"
)

// SiteSleImpactedSwitchesScopeParametersEnum is a string enum.
// enum: `site`
type SiteSleImpactedSwitchesScopeParametersEnum string

const (
    SiteSleImpactedSwitchesScopeParametersEnum_SITE SiteSleImpactedSwitchesScopeParametersEnum = "site"
)

// SiteSleImpactedUsersScopeParameterEnum is a string enum.
// enum: `ap`, `site`
type SiteSleImpactedUsersScopeParameterEnum string

const (
    SiteSleImpactedUsersScopeParameterEnum_AP   SiteSleImpactedUsersScopeParameterEnum = "ap"
    SiteSleImpactedUsersScopeParameterEnum_SITE SiteSleImpactedUsersScopeParameterEnum = "site"
)

// SiteSleMetricClassifiersScopeParametersEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleMetricClassifiersScopeParametersEnum string

const (
    SiteSleMetricClassifiersScopeParametersEnum_AP         SiteSleMetricClassifiersScopeParametersEnum = "ap"
    SiteSleMetricClassifiersScopeParametersEnum_CLIENT     SiteSleMetricClassifiersScopeParametersEnum = "client"
    SiteSleMetricClassifiersScopeParametersEnum_GATEWAY    SiteSleMetricClassifiersScopeParametersEnum = "gateway"
    SiteSleMetricClassifiersScopeParametersEnum_SITE       SiteSleMetricClassifiersScopeParametersEnum = "site"
    SiteSleMetricClassifiersScopeParametersEnum_ENUMSWITCH SiteSleMetricClassifiersScopeParametersEnum = "switch"
)

// SiteSleMetricSummaryScopeParametersEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleMetricSummaryScopeParametersEnum string

const (
    SiteSleMetricSummaryScopeParametersEnum_AP         SiteSleMetricSummaryScopeParametersEnum = "ap"
    SiteSleMetricSummaryScopeParametersEnum_CLIENT     SiteSleMetricSummaryScopeParametersEnum = "client"
    SiteSleMetricSummaryScopeParametersEnum_GATEWAY    SiteSleMetricSummaryScopeParametersEnum = "gateway"
    SiteSleMetricSummaryScopeParametersEnum_SITE       SiteSleMetricSummaryScopeParametersEnum = "site"
    SiteSleMetricSummaryScopeParametersEnum_ENUMSWITCH SiteSleMetricSummaryScopeParametersEnum = "switch"
)

// SiteSleMetricsScopeParametersEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleMetricsScopeParametersEnum string

const (
    SiteSleMetricsScopeParametersEnum_AP         SiteSleMetricsScopeParametersEnum = "ap"
    SiteSleMetricsScopeParametersEnum_CLIENT     SiteSleMetricsScopeParametersEnum = "client"
    SiteSleMetricsScopeParametersEnum_GATEWAY    SiteSleMetricsScopeParametersEnum = "gateway"
    SiteSleMetricsScopeParametersEnum_SITE       SiteSleMetricsScopeParametersEnum = "site"
    SiteSleMetricsScopeParametersEnum_ENUMSWITCH SiteSleMetricsScopeParametersEnum = "switch"
)

// SiteSleScopeEnum is a string enum.
// enum: `gateway`, `site`, `switch`
type SiteSleScopeEnum string

const (
    SiteSleScopeEnum_GATEWAY    SiteSleScopeEnum = "gateway"
    SiteSleScopeEnum_SITE       SiteSleScopeEnum = "site"
    SiteSleScopeEnum_ENUMSWITCH SiteSleScopeEnum = "switch"
)

// SiteSleThresholdScopeParameterEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SiteSleThresholdScopeParameterEnum string

const (
    SiteSleThresholdScopeParameterEnum_AP         SiteSleThresholdScopeParameterEnum = "ap"
    SiteSleThresholdScopeParameterEnum_CLIENT     SiteSleThresholdScopeParameterEnum = "client"
    SiteSleThresholdScopeParameterEnum_GATEWAY    SiteSleThresholdScopeParameterEnum = "gateway"
    SiteSleThresholdScopeParameterEnum_SITE       SiteSleThresholdScopeParameterEnum = "site"
    SiteSleThresholdScopeParameterEnum_ENUMSWITCH SiteSleThresholdScopeParameterEnum = "switch"
)

// SiteSystemEventsCountDistinctEnum is a string enum.
// enum: `type`
type SiteSystemEventsCountDistinctEnum string

const (
    SiteSystemEventsCountDistinctEnum_ENUMTYPE SiteSystemEventsCountDistinctEnum = "type"
)

// SiteWanClientEventsDistinctEnum is a string enum.
// enum: `hostname`, `ip`, `mac`, `mfg`, `type`
type SiteWanClientEventsDistinctEnum string

const (
    SiteWanClientEventsDistinctEnum_HOSTNAME SiteWanClientEventsDistinctEnum = "hostname"
    SiteWanClientEventsDistinctEnum_IP       SiteWanClientEventsDistinctEnum = "ip"
    SiteWanClientEventsDistinctEnum_MAC      SiteWanClientEventsDistinctEnum = "mac"
    SiteWanClientEventsDistinctEnum_MFG      SiteWanClientEventsDistinctEnum = "mfg"
    SiteWanClientEventsDistinctEnum_ENUMTYPE SiteWanClientEventsDistinctEnum = "type"
)

// SiteWanClientsCountDistinctEnum is a string enum.
// enum: `hostname`, `ip`, `mac`, `mfg`
type SiteWanClientsCountDistinctEnum string

const (
    SiteWanClientsCountDistinctEnum_HOSTNAME SiteWanClientsCountDistinctEnum = "hostname"
    SiteWanClientsCountDistinctEnum_IP       SiteWanClientsCountDistinctEnum = "ip"
    SiteWanClientsCountDistinctEnum_MAC      SiteWanClientsCountDistinctEnum = "mac"
    SiteWanClientsCountDistinctEnum_MFG      SiteWanClientsCountDistinctEnum = "mfg"
)

// SiteWifiProxyArpEnum is a string enum.
// enum: `default`, `disabled`, `enabled`
type SiteWifiProxyArpEnum string

const (
    SiteWifiProxyArpEnum_ENUMDEFAULT SiteWifiProxyArpEnum = "default"
    SiteWifiProxyArpEnum_DISABLED    SiteWifiProxyArpEnum = "disabled"
    SiteWifiProxyArpEnum_ENABLED     SiteWifiProxyArpEnum = "enabled"
)

// SiteWiredClientsCountDistinctEnum is a string enum.
// enum: `mac`, `port_id`, `vlan`
type SiteWiredClientsCountDistinctEnum string

const (
    SiteWiredClientsCountDistinctEnum_MAC    SiteWiredClientsCountDistinctEnum = "mac"
    SiteWiredClientsCountDistinctEnum_PORTID SiteWiredClientsCountDistinctEnum = "port_id"
    SiteWiredClientsCountDistinctEnum_VLAN   SiteWiredClientsCountDistinctEnum = "vlan"
)

// SiteZoneCountDistinctEnum is a string enum.
// enum: `scope`, `scope_id`, `user`, `user_type`
type SiteZoneCountDistinctEnum string

const (
    SiteZoneCountDistinctEnum_SCOPE    SiteZoneCountDistinctEnum = "scope"
    SiteZoneCountDistinctEnum_SCOPEID  SiteZoneCountDistinctEnum = "scope_id"
    SiteZoneCountDistinctEnum_USER     SiteZoneCountDistinctEnum = "user"
    SiteZoneCountDistinctEnum_USERTYPE SiteZoneCountDistinctEnum = "user_type"
)

// SleSummaryScopeEnum is a string enum.
// enum: `ap`, `client`, `gateway`, `site`, `switch`
type SleSummaryScopeEnum string

const (
    SleSummaryScopeEnum_AP         SleSummaryScopeEnum = "ap"
    SleSummaryScopeEnum_CLIENT     SleSummaryScopeEnum = "client"
    SleSummaryScopeEnum_GATEWAY    SleSummaryScopeEnum = "gateway"
    SleSummaryScopeEnum_SITE       SleSummaryScopeEnum = "site"
    SleSummaryScopeEnum_ENUMSWITCH SleSummaryScopeEnum = "switch"
)

// SnmpConfigEngineIdEnum is a string enum.
// enum: `engine-id-suffix`, `local`, `use-default-ip-address`, `use_mac-address`
type SnmpConfigEngineIdEnum string

const (
    SnmpConfigEngineIdEnum_ENGINEIDSUFFIX      SnmpConfigEngineIdEnum = "engine-id-suffix"
    SnmpConfigEngineIdEnum_LOCAL               SnmpConfigEngineIdEnum = "local"
    SnmpConfigEngineIdEnum_USEDEFAULTIPADDRESS SnmpConfigEngineIdEnum = "use-default-ip-address"
    SnmpConfigEngineIdEnum_USEMACADDRESS       SnmpConfigEngineIdEnum = "use_mac-address"
)

// SnmpConfigTrapVerionEnum is a string enum.
// enum: `all`, `v1`, `v2`
type SnmpConfigTrapVerionEnum string

const (
    SnmpConfigTrapVerionEnum_ALL SnmpConfigTrapVerionEnum = "all"
    SnmpConfigTrapVerionEnum_V1  SnmpConfigTrapVerionEnum = "v1"
    SnmpConfigTrapVerionEnum_V2  SnmpConfigTrapVerionEnum = "v2"
)

// SnmpUsmEngineTypeEnum is a string enum.
// enum: `local_engine`, `remote_engine`
type SnmpUsmEngineTypeEnum string

const (
    SnmpUsmEngineTypeEnum_LOCALENGINE  SnmpUsmEngineTypeEnum = "local_engine"
    SnmpUsmEngineTypeEnum_REMOTEENGINE SnmpUsmEngineTypeEnum = "remote_engine"
)

// SnmpUsmpUserAuthenticationTypeEnum is a string enum.
// sha224, sha256, sha384, sha512 are supported in 21.1 and newer release. enum: `authentication_md5`, `authentication_none`, `authentication_sha`, `authentication_sha224`, `authentication_sha256`, `authentication_sha384`, `authentication_sha512`
type SnmpUsmpUserAuthenticationTypeEnum string

const (
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONMD5    SnmpUsmpUserAuthenticationTypeEnum = "authentication_md5"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONNONE   SnmpUsmpUserAuthenticationTypeEnum = "authentication_none"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA    SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA224 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha224"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA256 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha256"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA384 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha384"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA512 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha512"
)

// SnmpUsmpUserEncryptionTypeEnum is a string enum.
// enum: `privacy-3des`, `privacy-aes128`, `privacy-des`, `privacy-none`
type SnmpUsmpUserEncryptionTypeEnum string

const (
    SnmpUsmpUserEncryptionTypeEnum_PRIVACY3DES   SnmpUsmpUserEncryptionTypeEnum = "privacy-3des"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYAES128 SnmpUsmpUserEncryptionTypeEnum = "privacy-aes128"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYDES    SnmpUsmpUserEncryptionTypeEnum = "privacy-des"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYNONE   SnmpUsmpUserEncryptionTypeEnum = "privacy-none"
)

// SnmpVacmAccessItemPrefixListItemLevelEnum is a string enum.
// enum: `authentication`, `none`, `privacy`
type SnmpVacmAccessItemPrefixListItemLevelEnum string

const (
    SnmpVacmAccessItemPrefixListItemLevelEnum_AUTHENTICATION SnmpVacmAccessItemPrefixListItemLevelEnum = "authentication"
    SnmpVacmAccessItemPrefixListItemLevelEnum_NONE           SnmpVacmAccessItemPrefixListItemLevelEnum = "none"
    SnmpVacmAccessItemPrefixListItemLevelEnum_PRIVACY        SnmpVacmAccessItemPrefixListItemLevelEnum = "privacy"
)

// SnmpVacmAccessItemPrefixListItemModelEnum is a string enum.
// enum: `any`, `usm`, `v1`, `v2c`
type SnmpVacmAccessItemPrefixListItemModelEnum string

const (
    SnmpVacmAccessItemPrefixListItemModelEnum_ANY SnmpVacmAccessItemPrefixListItemModelEnum = "any"
    SnmpVacmAccessItemPrefixListItemModelEnum_USM SnmpVacmAccessItemPrefixListItemModelEnum = "usm"
    SnmpVacmAccessItemPrefixListItemModelEnum_V1  SnmpVacmAccessItemPrefixListItemModelEnum = "v1"
    SnmpVacmAccessItemPrefixListItemModelEnum_V2C SnmpVacmAccessItemPrefixListItemModelEnum = "v2c"
)

// SnmpVacmAccessItemTypeEnum is a string enum.
// enum: `context_prefix`, `default_context_prefix`
type SnmpVacmAccessItemTypeEnum string

const (
    SnmpVacmAccessItemTypeEnum_CONTEXTPREFIX        SnmpVacmAccessItemTypeEnum = "context_prefix"
    SnmpVacmAccessItemTypeEnum_DEFAULTCONTEXTPREFIX SnmpVacmAccessItemTypeEnum = "default_context_prefix"
)

// SnmpVacmSecurityModelEnum is a string enum.
// enum: `usm`, `v1`, `v2c`
type SnmpVacmSecurityModelEnum string

const (
    SnmpVacmSecurityModelEnum_USM SnmpVacmSecurityModelEnum = "usm"
    SnmpVacmSecurityModelEnum_V1  SnmpVacmSecurityModelEnum = "v1"
    SnmpVacmSecurityModelEnum_V2C SnmpVacmSecurityModelEnum = "v2c"
)

// Snmpv3ConfigNotifyTypeEnum is a string enum.
// enum: `inform`, `trap`
type Snmpv3ConfigNotifyTypeEnum string

const (
    Snmpv3ConfigNotifyTypeEnum_INFORM Snmpv3ConfigNotifyTypeEnum = "inform"
    Snmpv3ConfigNotifyTypeEnum_TRAP   Snmpv3ConfigNotifyTypeEnum = "trap"
)

// Snmpv3ConfigTargetParamMessProcessModelEnum is a string enum.
// enum: `v1`, `v2c`, `v3`
type Snmpv3ConfigTargetParamMessProcessModelEnum string

const (
    Snmpv3ConfigTargetParamMessProcessModelEnum_V1  Snmpv3ConfigTargetParamMessProcessModelEnum = "v1"
    Snmpv3ConfigTargetParamMessProcessModelEnum_V2C Snmpv3ConfigTargetParamMessProcessModelEnum = "v2c"
    Snmpv3ConfigTargetParamMessProcessModelEnum_V3  Snmpv3ConfigTargetParamMessProcessModelEnum = "v3"
)

// Snmpv3ConfigTargetParamSecurityLevelEnum is a string enum.
// enum: `authentication`, `none`, `privacy`
type Snmpv3ConfigTargetParamSecurityLevelEnum string

const (
    Snmpv3ConfigTargetParamSecurityLevelEnum_AUTHENTICATION Snmpv3ConfigTargetParamSecurityLevelEnum = "authentication"
    Snmpv3ConfigTargetParamSecurityLevelEnum_NONE           Snmpv3ConfigTargetParamSecurityLevelEnum = "none"
    Snmpv3ConfigTargetParamSecurityLevelEnum_PRIVACY        Snmpv3ConfigTargetParamSecurityLevelEnum = "privacy"
)

// Snmpv3ConfigTargetParamSecurityModelEnum is a string enum.
// enum: `usm`, `v1`, `v2c`
type Snmpv3ConfigTargetParamSecurityModelEnum string

const (
    Snmpv3ConfigTargetParamSecurityModelEnum_USM Snmpv3ConfigTargetParamSecurityModelEnum = "usm"
    Snmpv3ConfigTargetParamSecurityModelEnum_V1  Snmpv3ConfigTargetParamSecurityModelEnum = "v1"
    Snmpv3ConfigTargetParamSecurityModelEnum_V2C Snmpv3ConfigTargetParamSecurityModelEnum = "v2c"
)

// SslProxyCiphersCatagoryEnum is a string enum.
// enum: `medium`, `strong`, `weak`
type SslProxyCiphersCatagoryEnum string

const (
    SslProxyCiphersCatagoryEnum_MEDIUM SslProxyCiphersCatagoryEnum = "medium"
    SslProxyCiphersCatagoryEnum_STRONG SslProxyCiphersCatagoryEnum = "strong"
    SslProxyCiphersCatagoryEnum_WEAK   SslProxyCiphersCatagoryEnum = "weak"
)

// SsoIdpSignAlgoEnum is a string enum.
// Required if `idp_type`==`saml`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`
type SsoIdpSignAlgoEnum string

const (
    SsoIdpSignAlgoEnum_SHA1   SsoIdpSignAlgoEnum = "sha1"
    SsoIdpSignAlgoEnum_SHA256 SsoIdpSignAlgoEnum = "sha256"
    SsoIdpSignAlgoEnum_SHA384 SsoIdpSignAlgoEnum = "sha384"
    SsoIdpSignAlgoEnum_SHA512 SsoIdpSignAlgoEnum = "sha512"
)

// SsoIdpTypeEnum is a string enum.
// SSO IDP Type:
// * For Admin SSO, enum: `saml`
// * For NAC SSO, enum: `ldap`, `mxedge_proxy`, `oauth`
type SsoIdpTypeEnum string

const (
    SsoIdpTypeEnum_LDAP        SsoIdpTypeEnum = "ldap"
    SsoIdpTypeEnum_MXEDGEPROXY SsoIdpTypeEnum = "mxedge_proxy"
    SsoIdpTypeEnum_OAUTH       SsoIdpTypeEnum = "oauth"
    SsoIdpTypeEnum_SAML        SsoIdpTypeEnum = "saml"
)

// SsoLdapTypeEnum is a string enum.
// if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`
type SsoLdapTypeEnum string

const (
    SsoLdapTypeEnum_AZURE  SsoLdapTypeEnum = "azure"
    SsoLdapTypeEnum_CUSTOM SsoLdapTypeEnum = "custom"
    SsoLdapTypeEnum_GOOGLE SsoLdapTypeEnum = "google"
    SsoLdapTypeEnum_OKTA   SsoLdapTypeEnum = "okta"
)

// SsoNameidFormatEnum is a string enum.
// if `idp_type`==`saml`. enum: `email`, `unspecified`
type SsoNameidFormatEnum string

const (
    SsoNameidFormatEnum_EMAIL       SsoNameidFormatEnum = "email"
    SsoNameidFormatEnum_UNSPECIFIED SsoNameidFormatEnum = "unspecified"
)

// SsoOauthTypeEnum is a string enum.
// if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`
type SsoOauthTypeEnum string

const (
    SsoOauthTypeEnum_AZURE        SsoOauthTypeEnum = "azure"
    SsoOauthTypeEnum_AZUREGOV     SsoOauthTypeEnum = "azure-gov"
    SsoOauthTypeEnum_STANDARDS    SsoOauthTypeEnum = "standards"
    SsoOauthTypeEnum_OKTA         SsoOauthTypeEnum = "okta"
    SsoOauthTypeEnum_PINGIDENTITY SsoOauthTypeEnum = "ping_identity"
)

// SsrUpgradeChannelEnum is a string enum.
// upgrade channel to follow. enum: `alpha`, `beta`, `stable`
type SsrUpgradeChannelEnum string

const (
    SsrUpgradeChannelEnum_ALPHA  SsrUpgradeChannelEnum = "alpha"
    SsrUpgradeChannelEnum_BETA   SsrUpgradeChannelEnum = "beta"
    SsrUpgradeChannelEnum_STABLE SsrUpgradeChannelEnum = "stable"
)

// SsrUpgradeStrategyEnum is a string enum.
// enum:
// * `big_bang`: upgrade all at once
// * `serial`: one at a time
type SsrUpgradeStrategyEnum string

const (
    SsrUpgradeStrategyEnum_BIGBANG SsrUpgradeStrategyEnum = "big_bang"
    SsrUpgradeStrategyEnum_SERIAL  SsrUpgradeStrategyEnum = "serial"
)

// StatDeviceStatusFilterEnum is a string enum.
// enum: `all`, `connected`, `disconnected`
type StatDeviceStatusFilterEnum string

const (
    StatDeviceStatusFilterEnum_ALL          StatDeviceStatusFilterEnum = "all"
    StatDeviceStatusFilterEnum_CONNECTED    StatDeviceStatusFilterEnum = "connected"
    StatDeviceStatusFilterEnum_DISCONNECTED StatDeviceStatusFilterEnum = "disconnected"
)

// StatsApGpsStatSrcEnum is a string enum.
// The origin of the GPS data. enum:
// * `gps`: from this device’s GPS estimates
// * `other_ap` from neighboring device GPS estimates
type StatsApGpsStatSrcEnum string

const (
    StatsApGpsStatSrcEnum_GPS     StatsApGpsStatSrcEnum = "gps"
    StatsApGpsStatSrcEnum_OTHERAP StatsApGpsStatSrcEnum = "other_ap"
)

// StatsMxtunnelStateEnum is a string enum.
// enum: `established`, `established_with_sessions`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
type StatsMxtunnelStateEnum string

const (
    StatsMxtunnelStateEnum_ESTABLISHED             StatsMxtunnelStateEnum = "established"
    StatsMxtunnelStateEnum_ESTABLISHEDWITHSESSIONS StatsMxtunnelStateEnum = "established_with_sessions"
    StatsMxtunnelStateEnum_IDLE                    StatsMxtunnelStateEnum = "idle"
    StatsMxtunnelStateEnum_WAITCTRLCONN            StatsMxtunnelStateEnum = "wait-ctrl-conn"
    StatsMxtunnelStateEnum_WAITCTRLREPLY           StatsMxtunnelStateEnum = "wait-ctrl-reply"
)

// StatsSwitchPortAuthStateEnum is a string enum.
// if `up`==`true` and has Authenticator role. enum: `authenticated`, `authenticating`, `held`, `init`
type StatsSwitchPortAuthStateEnum string

const (
    StatsSwitchPortAuthStateEnum_AUTHENTICATED  StatsSwitchPortAuthStateEnum = "authenticated"
    StatsSwitchPortAuthStateEnum_AUTHENTICATING StatsSwitchPortAuthStateEnum = "authenticating"
    StatsSwitchPortAuthStateEnum_HELD           StatsSwitchPortAuthStateEnum = "held"
    StatsSwitchPortAuthStateEnum_INIT           StatsSwitchPortAuthStateEnum = "init"
)

// StatsSwitchPortPoeModeEnum is a string enum.
// enum: `802.3af`, `802.3at`, `802.3bt`
type StatsSwitchPortPoeModeEnum string

const (
    StatsSwitchPortPoeModeEnum_ENUM8023AF StatsSwitchPortPoeModeEnum = "802.3af"
    StatsSwitchPortPoeModeEnum_ENUM8023AT StatsSwitchPortPoeModeEnum = "802.3at"
    StatsSwitchPortPoeModeEnum_ENUM8023BT StatsSwitchPortPoeModeEnum = "802.3bt"
)

// StatsSwitchPortPortUsageEnum is a string enum.
// gateway port usage. enum: `lan`
type StatsSwitchPortPortUsageEnum string

const (
    StatsSwitchPortPortUsageEnum_LAN StatsSwitchPortPortUsageEnum = "lan"
)

// StatsSwitchPortStpRoleEnum is a string enum.
// if `up`==`true`. enum: `alternate`, `backup`, `designated`, `root`, `root-prevented`
type StatsSwitchPortStpRoleEnum string

const (
    StatsSwitchPortStpRoleEnum_ALTERNATE     StatsSwitchPortStpRoleEnum = "alternate"
    StatsSwitchPortStpRoleEnum_BACKUP        StatsSwitchPortStpRoleEnum = "backup"
    StatsSwitchPortStpRoleEnum_DESIGNATED    StatsSwitchPortStpRoleEnum = "designated"
    StatsSwitchPortStpRoleEnum_ROOT          StatsSwitchPortStpRoleEnum = "root"
    StatsSwitchPortStpRoleEnum_ROOTPREVENTED StatsSwitchPortStpRoleEnum = "root-prevented"
)

// StatsSwitchPortStpStateEnum is a string enum.
// if `up`==`true`. enum: `blocking`, `disabled`, `forwarding`, `learning`, `listening`
type StatsSwitchPortStpStateEnum string

const (
    StatsSwitchPortStpStateEnum_BLOCKING   StatsSwitchPortStpStateEnum = "blocking"
    StatsSwitchPortStpStateEnum_DISABLED   StatsSwitchPortStpStateEnum = "disabled"
    StatsSwitchPortStpStateEnum_FORWARDING StatsSwitchPortStpStateEnum = "forwarding"
    StatsSwitchPortStpStateEnum_LEARNING   StatsSwitchPortStpStateEnum = "learning"
    StatsSwitchPortStpStateEnum_LISTENING  StatsSwitchPortStpStateEnum = "listening"
)

// StatsSwitchPortTypeEnum is a string enum.
// device type. enum: `ap`, `ble`, `gateway`, `mxedge`, `nac`, `switch`
type StatsSwitchPortTypeEnum string

const (
    StatsSwitchPortTypeEnum_AP         StatsSwitchPortTypeEnum = "ap"
    StatsSwitchPortTypeEnum_BLE        StatsSwitchPortTypeEnum = "ble"
    StatsSwitchPortTypeEnum_GATEWAY    StatsSwitchPortTypeEnum = "gateway"
    StatsSwitchPortTypeEnum_MXEDGE     StatsSwitchPortTypeEnum = "mxedge"
    StatsSwitchPortTypeEnum_NAC        StatsSwitchPortTypeEnum = "nac"
    StatsSwitchPortTypeEnum_ENUMSWITCH StatsSwitchPortTypeEnum = "switch"
)

// StatsWanTunnelPriorityEnum is a string enum.
// enum: `primary`, `secondary`
type StatsWanTunnelPriorityEnum string

const (
    StatsWanTunnelPriorityEnum_PRIMARY   StatsWanTunnelPriorityEnum = "primary"
    StatsWanTunnelPriorityEnum_SECONDARY StatsWanTunnelPriorityEnum = "secondary"
)

// StatsWxruleActionEnum is a string enum.
// enum: `allow`, `block`
type StatsWxruleActionEnum string

const (
    StatsWxruleActionEnum_ALLOW StatsWxruleActionEnum = "allow"
    StatsWxruleActionEnum_BLOCK StatsWxruleActionEnum = "block"
)

// SuppressedAlarmScopeEnum is a string enum.
// level of scope. enum: `org`, `site`
type SuppressedAlarmScopeEnum string

const (
    SuppressedAlarmScopeEnum_ORG  SuppressedAlarmScopeEnum = "org"
    SuppressedAlarmScopeEnum_SITE SuppressedAlarmScopeEnum = "site"
)

// SwitchDhcpdConfigTypeEnum is a string enum.
// enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server)
type SwitchDhcpdConfigTypeEnum string

const (
    SwitchDhcpdConfigTypeEnum_NONE   SwitchDhcpdConfigTypeEnum = "none"
    SwitchDhcpdConfigTypeEnum_RELAY  SwitchDhcpdConfigTypeEnum = "relay"
    SwitchDhcpdConfigTypeEnum_SERVER SwitchDhcpdConfigTypeEnum = "server"
)

// SwitchMetricScopeEnum is a string enum.
// enum: `site`, `switch`
type SwitchMetricScopeEnum string

const (
    SwitchMetricScopeEnum_SITE       SwitchMetricScopeEnum = "site"
    SwitchMetricScopeEnum_ENUMSWITCH SwitchMetricScopeEnum = "switch"
)

// SwitchMetricTypeEnum is a string enum.
// enum: `active_ports_summary`
type SwitchMetricTypeEnum string

const (
    SwitchMetricTypeEnum_ACTIVEPORTSSUMMARY SwitchMetricTypeEnum = "active_ports_summary"
)

// SwitchPortLocalUsageDot1xEnum is a string enum.
// if dot1x is desired, set to dot1x. enum: `dot1x`
type SwitchPortLocalUsageDot1xEnum string

const (
    SwitchPortLocalUsageDot1xEnum_DOT1X SwitchPortLocalUsageDot1xEnum = "dot1x"
)

// SwitchPortLocalUsageDuplexEnum is a string enum.
// link connection mode. enum: `auto`, `full`, `half`
type SwitchPortLocalUsageDuplexEnum string

const (
    SwitchPortLocalUsageDuplexEnum_AUTO SwitchPortLocalUsageDuplexEnum = "auto"
    SwitchPortLocalUsageDuplexEnum_FULL SwitchPortLocalUsageDuplexEnum = "full"
    SwitchPortLocalUsageDuplexEnum_HALF SwitchPortLocalUsageDuplexEnum = "half"
)

// SwitchPortLocalUsageMacAuthProtocolEnum is a string enum.
// Only if `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`
type SwitchPortLocalUsageMacAuthProtocolEnum string

const (
    SwitchPortLocalUsageMacAuthProtocolEnum_EAPMD5  SwitchPortLocalUsageMacAuthProtocolEnum = "eap-md5"
    SwitchPortLocalUsageMacAuthProtocolEnum_EAPPEAP SwitchPortLocalUsageMacAuthProtocolEnum = "eap-peap"
    SwitchPortLocalUsageMacAuthProtocolEnum_PAP     SwitchPortLocalUsageMacAuthProtocolEnum = "pap"
)

// SwitchPortLocalUsageModeEnum is a string enum.
// enum: `access`, `inet`, `trunk`
type SwitchPortLocalUsageModeEnum string

const (
    SwitchPortLocalUsageModeEnum_ACCESS SwitchPortLocalUsageModeEnum = "access"
    SwitchPortLocalUsageModeEnum_INET   SwitchPortLocalUsageModeEnum = "inet"
    SwitchPortLocalUsageModeEnum_TRUNK  SwitchPortLocalUsageModeEnum = "trunk"
)

// SwitchPortUsageDot1xEnum is a string enum.
// Only if `mode`!=`dynamic` if dot1x is desired, set to dot1x. enum: `dot1x`
type SwitchPortUsageDot1xEnum string

const (
    SwitchPortUsageDot1xEnum_DOT1X SwitchPortUsageDot1xEnum = "dot1x"
)

// SwitchPortUsageDuplexEnum is a string enum.
// Only if `mode`!=`dynamic` link connection mode. enum: `auto`, `full`, `half`
type SwitchPortUsageDuplexEnum string

const (
    SwitchPortUsageDuplexEnum_AUTO SwitchPortUsageDuplexEnum = "auto"
    SwitchPortUsageDuplexEnum_FULL SwitchPortUsageDuplexEnum = "full"
    SwitchPortUsageDuplexEnum_HALF SwitchPortUsageDuplexEnum = "half"
)

// SwitchPortUsageDynamicResetDefaultWhenEnum is a string enum.
// Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)
type SwitchPortUsageDynamicResetDefaultWhenEnum string

const (
    SwitchPortUsageDynamicResetDefaultWhenEnum_LINKDOWN SwitchPortUsageDynamicResetDefaultWhenEnum = "link_down"
    SwitchPortUsageDynamicResetDefaultWhenEnum_NONE     SwitchPortUsageDynamicResetDefaultWhenEnum = "none"
)

// SwitchPortUsageDynamicRuleSrcEnum is a string enum.
// enum: `link_peermac`, `lldp_chassis_id`, `lldp_hardware_revision`, `lldp_manufacturer_name`, `lldp_oui`, `lldp_serial_number`, `lldp_system_name`, `radius_dynamicfilter`, `radius_usermac`, `radius_username`
type SwitchPortUsageDynamicRuleSrcEnum string

const (
    SwitchPortUsageDynamicRuleSrcEnum_LINKPEERMAC          SwitchPortUsageDynamicRuleSrcEnum = "link_peermac"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPCHASSISID        SwitchPortUsageDynamicRuleSrcEnum = "lldp_chassis_id"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPHARDWAREREVISION SwitchPortUsageDynamicRuleSrcEnum = "lldp_hardware_revision"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPMANUFACTURERNAME SwitchPortUsageDynamicRuleSrcEnum = "lldp_manufacturer_name"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPOUI              SwitchPortUsageDynamicRuleSrcEnum = "lldp_oui"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPSERIALNUMBER     SwitchPortUsageDynamicRuleSrcEnum = "lldp_serial_number"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPSYSTEMNAME       SwitchPortUsageDynamicRuleSrcEnum = "lldp_system_name"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSDYNAMICFILTER  SwitchPortUsageDynamicRuleSrcEnum = "radius_dynamicfilter"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSUSERMAC        SwitchPortUsageDynamicRuleSrcEnum = "radius_usermac"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSUSERNAME       SwitchPortUsageDynamicRuleSrcEnum = "radius_username"
)

// SwitchPortUsageMacAuthProtocolEnum is a string enum.
// Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`
type SwitchPortUsageMacAuthProtocolEnum string

const (
    SwitchPortUsageMacAuthProtocolEnum_EAPMD5  SwitchPortUsageMacAuthProtocolEnum = "eap-md5"
    SwitchPortUsageMacAuthProtocolEnum_EAPPEAP SwitchPortUsageMacAuthProtocolEnum = "eap-peap"
    SwitchPortUsageMacAuthProtocolEnum_PAP     SwitchPortUsageMacAuthProtocolEnum = "pap"
)

// SwitchPortUsageModeEnum is a string enum.
// `mode`==`dynamic` must only be used if the port usage name is `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk`
type SwitchPortUsageModeEnum string

const (
    SwitchPortUsageModeEnum_ACCESS  SwitchPortUsageModeEnum = "access"
    SwitchPortUsageModeEnum_DYNAMIC SwitchPortUsageModeEnum = "dynamic"
    SwitchPortUsageModeEnum_INET    SwitchPortUsageModeEnum = "inet"
    SwitchPortUsageModeEnum_TRUNK   SwitchPortUsageModeEnum = "trunk"
)

// SwitchPortUsageSpeedEnum is a string enum.
// Only if `mode`!=`dynamic` speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
type SwitchPortUsageSpeedEnum string

const (
    SwitchPortUsageSpeedEnum_ENUM10M  SwitchPortUsageSpeedEnum = "10m"
    SwitchPortUsageSpeedEnum_ENUM100M SwitchPortUsageSpeedEnum = "100m"
    SwitchPortUsageSpeedEnum_ENUM1G   SwitchPortUsageSpeedEnum = "1g"
    SwitchPortUsageSpeedEnum_ENUM25G  SwitchPortUsageSpeedEnum = "2.5g"
    SwitchPortUsageSpeedEnum_ENUM5G   SwitchPortUsageSpeedEnum = "5g"
    SwitchPortUsageSpeedEnum_ENUM10G  SwitchPortUsageSpeedEnum = "10g"
    SwitchPortUsageSpeedEnum_ENUM25G1 SwitchPortUsageSpeedEnum = "25g"
    SwitchPortUsageSpeedEnum_ENUM40G  SwitchPortUsageSpeedEnum = "40g"
    SwitchPortUsageSpeedEnum_ENUM100G SwitchPortUsageSpeedEnum = "100g"
    SwitchPortUsageSpeedEnum_AUTO     SwitchPortUsageSpeedEnum = "auto"
)

// SwitchVirtualChassisMemberVcRoleEnum is a string enum.
// Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config. enum: `backup`, `linecard`, `master`
type SwitchVirtualChassisMemberVcRoleEnum string

const (
    SwitchVirtualChassisMemberVcRoleEnum_BACKUP   SwitchVirtualChassisMemberVcRoleEnum = "backup"
    SwitchVirtualChassisMemberVcRoleEnum_LINECARD SwitchVirtualChassisMemberVcRoleEnum = "linecard"
    SwitchVirtualChassisMemberVcRoleEnum_MASTER   SwitchVirtualChassisMemberVcRoleEnum = "master"
)

// SynthetictestDeviceProtocolEnum is a string enum.
// if `type`==`lan_connectivity`. enum: `ping`, `traceroute`, `ping+traceroute`
type SynthetictestDeviceProtocolEnum string

const (
    SynthetictestDeviceProtocolEnum_PING               SynthetictestDeviceProtocolEnum = "ping"
    SynthetictestDeviceProtocolEnum_ENUMPINGTRACEROUTE SynthetictestDeviceProtocolEnum = "ping+traceroute"
    SynthetictestDeviceProtocolEnum_TRACEROUTE         SynthetictestDeviceProtocolEnum = "traceroute"
)

// SynthetictestInfoDeviceTypeEnum is a string enum.
// enum: `ap`, `gateway`, `switch`
type SynthetictestInfoDeviceTypeEnum string

const (
    SynthetictestInfoDeviceTypeEnum_AP         SynthetictestInfoDeviceTypeEnum = "ap"
    SynthetictestInfoDeviceTypeEnum_GATEWAY    SynthetictestInfoDeviceTypeEnum = "gateway"
    SynthetictestInfoDeviceTypeEnum_ENUMSWITCH SynthetictestInfoDeviceTypeEnum = "switch"
)

// SynthetictestProtocolEnum is a string enum.
// enum: `ping`, `traceroute`
type SynthetictestProtocolEnum string

const (
    SynthetictestProtocolEnum_PING       SynthetictestProtocolEnum = "ping"
    SynthetictestProtocolEnum_TRACEROUTE SynthetictestProtocolEnum = "traceroute"
)

// SynthetictestTypeEnum is a string enum.
// enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest`
type SynthetictestTypeEnum string

const (
    SynthetictestTypeEnum_ARP             SynthetictestTypeEnum = "arp"
    SynthetictestTypeEnum_CURL            SynthetictestTypeEnum = "curl"
    SynthetictestTypeEnum_DHCP            SynthetictestTypeEnum = "dhcp"
    SynthetictestTypeEnum_DHCP6           SynthetictestTypeEnum = "dhcp6"
    SynthetictestTypeEnum_DNS             SynthetictestTypeEnum = "dns"
    SynthetictestTypeEnum_LANCONNECTIVITY SynthetictestTypeEnum = "lan_connectivity"
    SynthetictestTypeEnum_RADIUS          SynthetictestTypeEnum = "radius"
    SynthetictestTypeEnum_SPEEDTEST       SynthetictestTypeEnum = "speedtest"
)

// TacacsDefaultRoleEnum is a string enum.
// enum: `admin`, `helpdesk`, `none`, `read`
type TacacsDefaultRoleEnum string

const (
    TacacsDefaultRoleEnum_ADMIN    TacacsDefaultRoleEnum = "admin"
    TacacsDefaultRoleEnum_HELPDESK TacacsDefaultRoleEnum = "helpdesk"
    TacacsDefaultRoleEnum_NONE     TacacsDefaultRoleEnum = "none"
    TacacsDefaultRoleEnum_READ     TacacsDefaultRoleEnum = "read"
)

// TicketStatusEnum is a string enum.
// Ticket status. enum: 
// * open: ticket is open, Mist is working on it
// * pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information)
// * solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it)
// * closed: ticket is archived and cannot be changed.
type TicketStatusEnum string

const (
    TicketStatusEnum_CLOSED  TicketStatusEnum = "closed"
    TicketStatusEnum_OPEN    TicketStatusEnum = "open"
    TicketStatusEnum_PENDING TicketStatusEnum = "pending"
    TicketStatusEnum_SOLVED  TicketStatusEnum = "solved"
)

// TroubleshootTypeEnum is a string enum.
// enum: `wan`, `wired`, `wireless`
type TroubleshootTypeEnum string

const (
    TroubleshootTypeEnum_WAN      TroubleshootTypeEnum = "wan"
    TroubleshootTypeEnum_WIRED    TroubleshootTypeEnum = "wired"
    TroubleshootTypeEnum_WIRELESS TroubleshootTypeEnum = "wireless"
)

// TunnelConfigAuthAlgoEnum is a string enum.
// enum: `md5`, `sha1`, `sha2`
type TunnelConfigAuthAlgoEnum string

const (
    TunnelConfigAuthAlgoEnum_MD5  TunnelConfigAuthAlgoEnum = "md5"
    TunnelConfigAuthAlgoEnum_SHA1 TunnelConfigAuthAlgoEnum = "sha1"
    TunnelConfigAuthAlgoEnum_SHA2 TunnelConfigAuthAlgoEnum = "sha2"
)

// TunnelConfigAutoProvisionProviderEnum is a string enum.
// enum: `jse-ipsec`, `zscaler-ipsec`
type TunnelConfigAutoProvisionProviderEnum string

const (
    TunnelConfigAutoProvisionProviderEnum_JSEIPSEC     TunnelConfigAutoProvisionProviderEnum = "jse-ipsec"
    TunnelConfigAutoProvisionProviderEnum_ZSCALERIPSEC TunnelConfigAutoProvisionProviderEnum = "zscaler-ipsec"
)

// TunnelConfigDhGroupEnum is a string enum.
// Only if `provider`==`custom-ipsec`. enum:
// * 1
// * 2 (1024-bit)
// * 5
// * 14 (default, 2048-bit)
// * 15 (3072-bit)
// * 16 (4096-bit)
// * 19 (256-bit ECP)
// * 20 (384-bit ECP)
// * 21 (521-bit ECP)
// * 24 (2048-bit ECP)
type TunnelConfigDhGroupEnum string

const (
    TunnelConfigDhGroupEnum_ENUM1  TunnelConfigDhGroupEnum = "1"
    TunnelConfigDhGroupEnum_ENUM14 TunnelConfigDhGroupEnum = "14"
    TunnelConfigDhGroupEnum_ENUM15 TunnelConfigDhGroupEnum = "15"
    TunnelConfigDhGroupEnum_ENUM16 TunnelConfigDhGroupEnum = "16"
    TunnelConfigDhGroupEnum_ENUM19 TunnelConfigDhGroupEnum = "19"
    TunnelConfigDhGroupEnum_ENUM2  TunnelConfigDhGroupEnum = "2"
    TunnelConfigDhGroupEnum_ENUM20 TunnelConfigDhGroupEnum = "20"
    TunnelConfigDhGroupEnum_ENUM21 TunnelConfigDhGroupEnum = "21"
    TunnelConfigDhGroupEnum_ENUM24 TunnelConfigDhGroupEnum = "24"
    TunnelConfigDhGroupEnum_ENUM5  TunnelConfigDhGroupEnum = "5"
)

// TunnelConfigEncAlgoEnum is a string enum.
// enum: `3des`, `aes128`, `aes256`, `aes_gcm128`, `aes_gcm256`
type TunnelConfigEncAlgoEnum string

const (
    TunnelConfigEncAlgoEnum_ENUM3DES  TunnelConfigEncAlgoEnum = "3des"
    TunnelConfigEncAlgoEnum_AES128    TunnelConfigEncAlgoEnum = "aes128"
    TunnelConfigEncAlgoEnum_AES256    TunnelConfigEncAlgoEnum = "aes256"
    TunnelConfigEncAlgoEnum_AESGCM128 TunnelConfigEncAlgoEnum = "aes_gcm128"
    TunnelConfigEncAlgoEnum_AESGCM256 TunnelConfigEncAlgoEnum = "aes_gcm256"
)

// TunnelConfigIkeDhGroupEnum is a string enum.
// enum:
// * 1
// * 2 (1024-bit)
// * 5
// * 14 (default, 2048-bit)
// * 15 (3072-bit)
// * 16 (4096-bit)
// * 19 (256-bit ECP)
// * 20 (384-bit ECP)
// * 21 (521-bit ECP)
// * 24 (2048-bit ECP)
type TunnelConfigIkeDhGroupEnum string

const (
    TunnelConfigIkeDhGroupEnum_ENUM1  TunnelConfigIkeDhGroupEnum = "1"
    TunnelConfigIkeDhGroupEnum_ENUM14 TunnelConfigIkeDhGroupEnum = "14"
    TunnelConfigIkeDhGroupEnum_ENUM15 TunnelConfigIkeDhGroupEnum = "15"
    TunnelConfigIkeDhGroupEnum_ENUM16 TunnelConfigIkeDhGroupEnum = "16"
    TunnelConfigIkeDhGroupEnum_ENUM19 TunnelConfigIkeDhGroupEnum = "19"
    TunnelConfigIkeDhGroupEnum_ENUM2  TunnelConfigIkeDhGroupEnum = "2"
    TunnelConfigIkeDhGroupEnum_ENUM20 TunnelConfigIkeDhGroupEnum = "20"
    TunnelConfigIkeDhGroupEnum_ENUM21 TunnelConfigIkeDhGroupEnum = "21"
    TunnelConfigIkeDhGroupEnum_ENUM24 TunnelConfigIkeDhGroupEnum = "24"
    TunnelConfigIkeDhGroupEnum_ENUM5  TunnelConfigIkeDhGroupEnum = "5"
)

// TunnelConfigIkeModeEnum is a string enum.
// Only if `provider`==`custom-ipsec`. enum: `aggressive`, `main`
type TunnelConfigIkeModeEnum string

const (
    TunnelConfigIkeModeEnum_AGGRESSIVE TunnelConfigIkeModeEnum = "aggressive"
    TunnelConfigIkeModeEnum_MAIN       TunnelConfigIkeModeEnum = "main"
)

// TunnelConfigProbeTypeEnum is a string enum.
// enum: `http`, `icmp`
type TunnelConfigProbeTypeEnum string

const (
    TunnelConfigProbeTypeEnum_HTTP TunnelConfigProbeTypeEnum = "http"
    TunnelConfigProbeTypeEnum_ICMP TunnelConfigProbeTypeEnum = "icmp"
)

// TunnelConfigProtocolEnum is a string enum.
// Only if `provider`==`custom-ipsec`. enum: `gre`, `ipsec`
type TunnelConfigProtocolEnum string

const (
    TunnelConfigProtocolEnum_GRE   TunnelConfigProtocolEnum = "gre"
    TunnelConfigProtocolEnum_IPSEC TunnelConfigProtocolEnum = "ipsec"
)

// TunnelConfigProviderEnum is a string enum.
// Only if `auto_provision.enabled`==`false`. enum: `custom-ipsec`, `customer-gre`, `jse-ipsec`, `zscaler-gre`, `zscaler-ipsec`
type TunnelConfigProviderEnum string

const (
    TunnelConfigProviderEnum_CUSTOMIPSEC  TunnelConfigProviderEnum = "custom-ipsec"
    TunnelConfigProviderEnum_CUSTOMERGRE  TunnelConfigProviderEnum = "customer-gre"
    TunnelConfigProviderEnum_JSEIPSEC     TunnelConfigProviderEnum = "jse-ipsec"
    TunnelConfigProviderEnum_ZSCALERGRE   TunnelConfigProviderEnum = "zscaler-gre"
    TunnelConfigProviderEnum_ZSCALERIPSEC TunnelConfigProviderEnum = "zscaler-ipsec"
)

// TunnelConfigTunnelModeEnum is a string enum.
// Required if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`. enum: `active-active`, `active-standby`
type TunnelConfigTunnelModeEnum string

const (
    TunnelConfigTunnelModeEnum_ACTIVEACTIVE  TunnelConfigTunnelModeEnum = "active-active"
    TunnelConfigTunnelModeEnum_ACTIVESTANDBY TunnelConfigTunnelModeEnum = "active-standby"
)

// TunnelConfigVersionEnum is a string enum.
// Only if `provider`==`custom-gre` or `provider`==`custom-ipsec`. enum: `1`, `2`
type TunnelConfigVersionEnum string

const (
    TunnelConfigVersionEnum_ENUM1 TunnelConfigVersionEnum = "1"
    TunnelConfigVersionEnum_ENUM2 TunnelConfigVersionEnum = "2"
)

// TunnelTypeEnum is a string enum.
// enum: `wan`, `wxtunnel`
type TunnelTypeEnum string

const (
    TunnelTypeEnum_WAN      TunnelTypeEnum = "wan"
    TunnelTypeEnum_WXTUNNEL TunnelTypeEnum = "wxtunnel"
)

// TuntermDhcpdTypeEnum is a string enum.
// enum: `relay`
type TuntermDhcpdTypeEnum string

const (
    TuntermDhcpdTypeEnum_RELAY TuntermDhcpdTypeEnum = "relay"
)

// TunternMonitoringProtocolEnum is a string enum.
// enum: `arp`, `ping`, `tcp`
type TunternMonitoringProtocolEnum string

const (
    TunternMonitoringProtocolEnum_ARP  TunternMonitoringProtocolEnum = "arp"
    TunternMonitoringProtocolEnum_PING TunternMonitoringProtocolEnum = "ping"
    TunternMonitoringProtocolEnum_TCP  TunternMonitoringProtocolEnum = "tcp"
)

// UpgradeInfoStatusEnum is a string enum.
// enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
type UpgradeInfoStatusEnum string

const (
    UpgradeInfoStatusEnum_ENUMERROR  UpgradeInfoStatusEnum = "error"
    UpgradeInfoStatusEnum_INPROGRESS UpgradeInfoStatusEnum = "inprogress"
    UpgradeInfoStatusEnum_SCHEDULED  UpgradeInfoStatusEnum = "scheduled"
    UpgradeInfoStatusEnum_STARTING   UpgradeInfoStatusEnum = "starting"
    UpgradeInfoStatusEnum_SUCCESS    UpgradeInfoStatusEnum = "success"
)

// UseAutoApValuesForEnum is a string enum.
// The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`
type UseAutoApValuesForEnum string

const (
    UseAutoApValuesForEnum_ORIENTATION UseAutoApValuesForEnum = "orientation"
    UseAutoApValuesForEnum_PLACEMENT   UseAutoApValuesForEnum = "placement"
)

// UtilsClearBgpTypeEnum is a string enum.
// enum: `hard`, `in`, `out`, `soft`
type UtilsClearBgpTypeEnum string

const (
    UtilsClearBgpTypeEnum_HARD UtilsClearBgpTypeEnum = "hard"
    UtilsClearBgpTypeEnum_IN   UtilsClearBgpTypeEnum = "in"
    UtilsClearBgpTypeEnum_OUT  UtilsClearBgpTypeEnum = "out"
    UtilsClearBgpTypeEnum_SOFT UtilsClearBgpTypeEnum = "soft"
)

// UtilsDevicesRestartNodeEnum is a string enum.
// only for SRX/SSR: if node is not present, both nodes are restarted. For other devices: node should not be present
type UtilsDevicesRestartNodeEnum string

const (
    UtilsDevicesRestartNodeEnum_NODE0 UtilsDevicesRestartNodeEnum = "node0"
    UtilsDevicesRestartNodeEnum_NODE1 UtilsDevicesRestartNodeEnum = "node1"
)

// UtilsSendSupportLogsInfoEnum is a string enum.
// optional, enum: 
// * `code-dumps`: Upload all core dump files, if any found
// * `full`: Upload 1 file with output of `request support information`, 1 file that concatenates all `/var/log/outbound-ssh.log*` files, all core dump files, the 5 most recent `/var/log/messages*` files, and Mist agent logs
// * `messages`: Upload 1 to 10 `/var/log/messages*` files
// * `outbound-ssh`: Upload 1 file that concatenates all `/var/log/outbound-ssh.log*` files
// * `process`: Upload 1 file with output of show `system processes extensive``
// * `var-logs`: Upload all non-empty files in the `/var/log/` directory
type UtilsSendSupportLogsInfoEnum string

const (
    UtilsSendSupportLogsInfoEnum_CODEDUMPS   UtilsSendSupportLogsInfoEnum = "code-dumps"
    UtilsSendSupportLogsInfoEnum_FULL        UtilsSendSupportLogsInfoEnum = "full"
    UtilsSendSupportLogsInfoEnum_MESSAGES    UtilsSendSupportLogsInfoEnum = "messages"
    UtilsSendSupportLogsInfoEnum_OUTBOUNDSSH UtilsSendSupportLogsInfoEnum = "outbound-ssh"
    UtilsSendSupportLogsInfoEnum_PROCESS     UtilsSendSupportLogsInfoEnum = "process"
    UtilsSendSupportLogsInfoEnum_VARLOGS     UtilsSendSupportLogsInfoEnum = "var-logs"
)

// UtilsShowRouteProtocolEnum is a string enum.
// enum: `any`, `bgp`, `direct`, `evpn`, `ospf`, `static`
type UtilsShowRouteProtocolEnum string

const (
    UtilsShowRouteProtocolEnum_ANY    UtilsShowRouteProtocolEnum = "any"
    UtilsShowRouteProtocolEnum_BGP    UtilsShowRouteProtocolEnum = "bgp"
    UtilsShowRouteProtocolEnum_DIRECT UtilsShowRouteProtocolEnum = "direct"
    UtilsShowRouteProtocolEnum_EVPN   UtilsShowRouteProtocolEnum = "evpn"
    UtilsShowRouteProtocolEnum_OSPF   UtilsShowRouteProtocolEnum = "ospf"
    UtilsShowRouteProtocolEnum_STATIC UtilsShowRouteProtocolEnum = "static"
)

// UtilsTracerouteProtocolEnum is a string enum.
// enum: `icmp` (Only suported by AP/MxEdge), `udp`
type UtilsTracerouteProtocolEnum string

const (
    UtilsTracerouteProtocolEnum_ICMP UtilsTracerouteProtocolEnum = "icmp"
    UtilsTracerouteProtocolEnum_UDP  UtilsTracerouteProtocolEnum = "udp"
)

// VarSourceEnum is a string enum.
// enum: `deviceprofile`, `site`
type VarSourceEnum string

const (
    VarSourceEnum_DEVICEPROFILE VarSourceEnum = "deviceprofile"
    VarSourceEnum_SITE          VarSourceEnum = "site"
)

// VirtualChassisConfigMemberVcRoleEnum is a string enum.
// enum: `backup`, `linecard`, `master`
type VirtualChassisConfigMemberVcRoleEnum string

const (
    VirtualChassisConfigMemberVcRoleEnum_BACKUP   VirtualChassisConfigMemberVcRoleEnum = "backup"
    VirtualChassisConfigMemberVcRoleEnum_LINECARD VirtualChassisConfigMemberVcRoleEnum = "linecard"
    VirtualChassisConfigMemberVcRoleEnum_MASTER   VirtualChassisConfigMemberVcRoleEnum = "master"
)

// VirtualChassisMemberUpdateVcRoleEnum is a string enum.
// Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master`
type VirtualChassisMemberUpdateVcRoleEnum string

const (
    VirtualChassisMemberUpdateVcRoleEnum_BACKUP   VirtualChassisMemberUpdateVcRoleEnum = "backup"
    VirtualChassisMemberUpdateVcRoleEnum_LINECARD VirtualChassisMemberUpdateVcRoleEnum = "linecard"
    VirtualChassisMemberUpdateVcRoleEnum_MASTER   VirtualChassisMemberUpdateVcRoleEnum = "master"
)

// VirtualChassisPortOperationEnum is a string enum.
// enum: `delete`, `set`
type VirtualChassisPortOperationEnum string

const (
    VirtualChassisPortOperationEnum_DELETE VirtualChassisPortOperationEnum = "delete"
    VirtualChassisPortOperationEnum_SET    VirtualChassisPortOperationEnum = "set"
)

// VirtualChassisUpdateOpEnum is a string enum.
// enum: `add`, `preprovision`, `remove`, `renumber`
type VirtualChassisUpdateOpEnum string

const (
    VirtualChassisUpdateOpEnum_ADD          VirtualChassisUpdateOpEnum = "add"
    VirtualChassisUpdateOpEnum_PREPROVISION VirtualChassisUpdateOpEnum = "preprovision"
    VirtualChassisUpdateOpEnum_REMOVE       VirtualChassisUpdateOpEnum = "remove"
    VirtualChassisUpdateOpEnum_RENUMBER     VirtualChassisUpdateOpEnum = "renumber"
)

// VisitsScopeEnum is a string enum.
// enum: `map`, `rssizone`, `site`, `zone`
type VisitsScopeEnum string

const (
    VisitsScopeEnum_ENUMMAP  VisitsScopeEnum = "map"
    VisitsScopeEnum_RSSIZONE VisitsScopeEnum = "rssizone"
    VisitsScopeEnum_SITE     VisitsScopeEnum = "site"
    VisitsScopeEnum_ZONE     VisitsScopeEnum = "zone"
)

// VpnPathBfdProfileEnum is a string enum.
// enum: `broadband`, `lte`
type VpnPathBfdProfileEnum string

const (
    VpnPathBfdProfileEnum_BROADBAND VpnPathBfdProfileEnum = "broadband"
    VpnPathBfdProfileEnum_LTE       VpnPathBfdProfileEnum = "lte"
)

// VpnPathSelectionStrategyEnum is a string enum.
// enum: `disabled`, `simple`, `manual`
type VpnPathSelectionStrategyEnum string

const (
    VpnPathSelectionStrategyEnum_DISABLED VpnPathSelectionStrategyEnum = "disabled"
    VpnPathSelectionStrategyEnum_SIMPLE   VpnPathSelectionStrategyEnum = "simple"
    VpnPathSelectionStrategyEnum_MANUAL   VpnPathSelectionStrategyEnum = "manual"
)

// VpnTypeEnum is a string enum.
// enum: `hub_spoke`, `mesh`
type VpnTypeEnum string

const (
    VpnTypeEnum_HUBSPOKE VpnTypeEnum = "hub_spoke"
    VpnTypeEnum_MESH     VpnTypeEnum = "mesh"
)

// VrrpGroupAuthTypeEnum is a string enum.
// enum: `md5`, `simple`
type VrrpGroupAuthTypeEnum string

const (
    VrrpGroupAuthTypeEnum_MD5    VrrpGroupAuthTypeEnum = "md5"
    VrrpGroupAuthTypeEnum_SIMPLE VrrpGroupAuthTypeEnum = "simple"
)

// WanTunnelProtocolEnum is a string enum.
// enum: `gre`, `ipsec`
type WanTunnelProtocolEnum string

const (
    WanTunnelProtocolEnum_GRE   WanTunnelProtocolEnum = "gre"
    WanTunnelProtocolEnum_IPSEC WanTunnelProtocolEnum = "ipsec"
)

// WanUsagesCountDisctinctEnum is a string enum.
// enum: `mac`, `path_type`, `peer_mac`, `peer_port_id`, `policy`, `port_id`, `tenant`
type WanUsagesCountDisctinctEnum string

const (
    WanUsagesCountDisctinctEnum_MAC        WanUsagesCountDisctinctEnum = "mac"
    WanUsagesCountDisctinctEnum_PATHTYPE   WanUsagesCountDisctinctEnum = "path_type"
    WanUsagesCountDisctinctEnum_PEERMAC    WanUsagesCountDisctinctEnum = "peer_mac"
    WanUsagesCountDisctinctEnum_PEERPORTID WanUsagesCountDisctinctEnum = "peer_port_id"
    WanUsagesCountDisctinctEnum_POLICY     WanUsagesCountDisctinctEnum = "policy"
    WanUsagesCountDisctinctEnum_PORTID     WanUsagesCountDisctinctEnum = "port_id"
    WanUsagesCountDisctinctEnum_TENANT     WanUsagesCountDisctinctEnum = "tenant"
)

// WebhookClientInfoTopicEnum is a string enum.
// enum: `client-info`
type WebhookClientInfoTopicEnum string

const (
    WebhookClientInfoTopicEnum_CLIENTINFO WebhookClientInfoTopicEnum = "client-info"
)

// WebhookDeliveryDistinctEnum is a string enum.
// webhook topic. enum: `status`, `status_code`, `topic`, `webhook_id`
type WebhookDeliveryDistinctEnum string

const (
    WebhookDeliveryDistinctEnum_STATUS     WebhookDeliveryDistinctEnum = "status"
    WebhookDeliveryDistinctEnum_STATUSCODE WebhookDeliveryDistinctEnum = "status_code"
    WebhookDeliveryDistinctEnum_TOPIC      WebhookDeliveryDistinctEnum = "topic"
    WebhookDeliveryDistinctEnum_WEBHOOKID  WebhookDeliveryDistinctEnum = "webhook_id"
)

// WebhookDeliveryStatusEnum is a string enum.
// webhook delivery status. enum: `failure`, `success`
type WebhookDeliveryStatusEnum string

const (
    WebhookDeliveryStatusEnum_FAILURE WebhookDeliveryStatusEnum = "failure"
    WebhookDeliveryStatusEnum_SUCCESS WebhookDeliveryStatusEnum = "success"
)

// WebhookDeliveryTopicEnum is a string enum.
// webhook topic. enum: `alarms`, `audits`, `device-updowns`, `occupancy-alerts`, `ping`
type WebhookDeliveryTopicEnum string

const (
    WebhookDeliveryTopicEnum_ALARMS          WebhookDeliveryTopicEnum = "alarms"
    WebhookDeliveryTopicEnum_AUDITS          WebhookDeliveryTopicEnum = "audits"
    WebhookDeliveryTopicEnum_DEVICEUPDOWNS   WebhookDeliveryTopicEnum = "device-updowns"
    WebhookDeliveryTopicEnum_OCCUPANCYALERTS WebhookDeliveryTopicEnum = "occupancy-alerts"
    WebhookDeliveryTopicEnum_PING            WebhookDeliveryTopicEnum = "ping"
)

// WebhookDeviceEventsEventDeviceTypeEnum is a string enum.
// enum: `ap`, `gateway`, `switch`
type WebhookDeviceEventsEventDeviceTypeEnum string

const (
    WebhookDeviceEventsEventDeviceTypeEnum_AP         WebhookDeviceEventsEventDeviceTypeEnum = "ap"
    WebhookDeviceEventsEventDeviceTypeEnum_GATEWAY    WebhookDeviceEventsEventDeviceTypeEnum = "gateway"
    WebhookDeviceEventsEventDeviceTypeEnum_ENUMSWITCH WebhookDeviceEventsEventDeviceTypeEnum = "switch"
)

// WebhookDeviceEventsEventEvTypeEnum is a string enum.
// (optional) event advisory. enum: `notice`, `warn`
type WebhookDeviceEventsEventEvTypeEnum string

const (
    WebhookDeviceEventsEventEvTypeEnum_NOTICE WebhookDeviceEventsEventEvTypeEnum = "notice"
    WebhookDeviceEventsEventEvTypeEnum_WARN   WebhookDeviceEventsEventEvTypeEnum = "warn"
)

// WebhookOauth2GrantTypeEnum is a string enum.
// required when `type`==`oauth2`. enum: `client_credentials`, `password`
type WebhookOauth2GrantTypeEnum string

const (
    WebhookOauth2GrantTypeEnum_CLIENTCREDENTIALS WebhookOauth2GrantTypeEnum = "client_credentials"
    WebhookOauth2GrantTypeEnum_PASSWORD          WebhookOauth2GrantTypeEnum = "password"
)

// WebhookOccupancyAlertTypeEnum is a string enum.
// enum: `COMPLIANCE-OK`, `COMPLIANCE-VIOLATION`
type WebhookOccupancyAlertTypeEnum string

const (
    WebhookOccupancyAlertTypeEnum_COMPLIANCEOK        WebhookOccupancyAlertTypeEnum = "COMPLIANCE-OK"
    WebhookOccupancyAlertTypeEnum_COMPLIANCEVIOLATION WebhookOccupancyAlertTypeEnum = "COMPLIANCE-VIOLATION"
)

// WebhookSdkclientScanDataTopicEnum is a string enum.
// enum: `sdkclient_scan_data`
type WebhookSdkclientScanDataTopicEnum string

const (
    WebhookSdkclientScanDataTopicEnum_SDKCLIENTSCANDATA WebhookSdkclientScanDataTopicEnum = "sdkclient_scan_data"
)

// WebhookTypeEnum is a string enum.
// enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`
type WebhookTypeEnum string

const (
    WebhookTypeEnum_AWSSNS       WebhookTypeEnum = "aws-sns"
    WebhookTypeEnum_GOOGLEPUBSUB WebhookTypeEnum = "google-pubsub"
    WebhookTypeEnum_HTTPPOST     WebhookTypeEnum = "http-post"
    WebhookTypeEnum_OAUTH2       WebhookTypeEnum = "oauth2"
    WebhookTypeEnum_SPLUNK       WebhookTypeEnum = "splunk"
)

// WebhookZoneEventTriggerEnum is a string enum.
// enum: `enter`, `exit`
type WebhookZoneEventTriggerEnum string

const (
    WebhookZoneEventTriggerEnum_ENTER WebhookZoneEventTriggerEnum = "enter"
    WebhookZoneEventTriggerEnum_EXIT  WebhookZoneEventTriggerEnum = "exit"
)

// WlanApplyToEnum is a string enum.
// enum: `aps`, `site`, `wxtags`
type WlanApplyToEnum string

const (
    WlanApplyToEnum_APS    WlanApplyToEnum = "aps"
    WlanApplyToEnum_SITE   WlanApplyToEnum = "site"
    WlanApplyToEnum_WXTAGS WlanApplyToEnum = "wxtags"
)

// WlanAuthOweEnum is a string enum.
// if `type`==`open`. enum: `disabled`, `enabled` (means transition mode), `required`
type WlanAuthOweEnum string

const (
    WlanAuthOweEnum_DISABLED WlanAuthOweEnum = "disabled"
    WlanAuthOweEnum_ENABLED  WlanAuthOweEnum = "enabled"
    WlanAuthOweEnum_REQUIRED WlanAuthOweEnum = "required"
)

// WlanAuthPairwiseItemEnum is a string enum.
// enum: `wpa1-ccmp`, `wpa1-tkip`, `wpa2-ccmp`, `wpa2-tkip`, `wpa3`
type WlanAuthPairwiseItemEnum string

const (
    WlanAuthPairwiseItemEnum_WPA1CCMP WlanAuthPairwiseItemEnum = "wpa1-ccmp"
    WlanAuthPairwiseItemEnum_WPA1TKIP WlanAuthPairwiseItemEnum = "wpa1-tkip"
    WlanAuthPairwiseItemEnum_WPA2CCMP WlanAuthPairwiseItemEnum = "wpa2-ccmp"
    WlanAuthPairwiseItemEnum_WPA2TKIP WlanAuthPairwiseItemEnum = "wpa2-tkip"
    WlanAuthPairwiseItemEnum_WPA3     WlanAuthPairwiseItemEnum = "wpa3"
)

// WlanAuthServerSelectionEnum is a string enum.
// When ordered, AP will prefer and go back to the first server if possible. enum: `ordered`, `unordered`
type WlanAuthServerSelectionEnum string

const (
    WlanAuthServerSelectionEnum_ORDERED   WlanAuthServerSelectionEnum = "ordered"
    WlanAuthServerSelectionEnum_UNORDERED WlanAuthServerSelectionEnum = "unordered"
)

// WlanAuthTypeEnum is a string enum.
// enum: `eap`, `eap192`, `open`, `psk`, `psk-tkip`, `psk-wpa2-tkip`, `wep`
type WlanAuthTypeEnum string

const (
    WlanAuthTypeEnum_EAP         WlanAuthTypeEnum = "eap"
    WlanAuthTypeEnum_EAP192      WlanAuthTypeEnum = "eap192"
    WlanAuthTypeEnum_OPEN        WlanAuthTypeEnum = "open"
    WlanAuthTypeEnum_PSK         WlanAuthTypeEnum = "psk"
    WlanAuthTypeEnum_PSKTKIP     WlanAuthTypeEnum = "psk-tkip"
    WlanAuthTypeEnum_PSKWPA2TKIP WlanAuthTypeEnum = "psk-wpa2-tkip"
    WlanAuthTypeEnum_WEP         WlanAuthTypeEnum = "wep"
)

// WlanBonjourServicePropertiesScopeEnum is a string enum.
// how bonjour services should be discovered for the same WLAN. enum: `same_ap`, `same_map`, `same_site`
type WlanBonjourServicePropertiesScopeEnum string

const (
    WlanBonjourServicePropertiesScopeEnum_SAMEAP   WlanBonjourServicePropertiesScopeEnum = "same_ap"
    WlanBonjourServicePropertiesScopeEnum_SAMEMAP  WlanBonjourServicePropertiesScopeEnum = "same_map"
    WlanBonjourServicePropertiesScopeEnum_SAMESITE WlanBonjourServicePropertiesScopeEnum = "same_site"
)

// WlanDataratesLegacyItemEnum is a string enum.
// enum: `1`, `11`, `11b`, `12`, `12b`, `18`, `18b`, `1b`, `2`, `24`, `24b`, `2b`, `36`, `36b`, `48`, `48b`, `5.5`, `5.5b`, `54`, `54b`, `6`, `6b`, `9`, `9b`
type WlanDataratesLegacyItemEnum string

const (
    WlanDataratesLegacyItemEnum_ENUM1   WlanDataratesLegacyItemEnum = "1"
    WlanDataratesLegacyItemEnum_ENUM11  WlanDataratesLegacyItemEnum = "11"
    WlanDataratesLegacyItemEnum_ENUM11B WlanDataratesLegacyItemEnum = "11b"
    WlanDataratesLegacyItemEnum_ENUM12  WlanDataratesLegacyItemEnum = "12"
    WlanDataratesLegacyItemEnum_ENUM12B WlanDataratesLegacyItemEnum = "12b"
    WlanDataratesLegacyItemEnum_ENUM18  WlanDataratesLegacyItemEnum = "18"
    WlanDataratesLegacyItemEnum_ENUM18B WlanDataratesLegacyItemEnum = "18b"
    WlanDataratesLegacyItemEnum_ENUM1B  WlanDataratesLegacyItemEnum = "1b"
    WlanDataratesLegacyItemEnum_ENUM2   WlanDataratesLegacyItemEnum = "2"
    WlanDataratesLegacyItemEnum_ENUM24  WlanDataratesLegacyItemEnum = "24"
    WlanDataratesLegacyItemEnum_ENUM24B WlanDataratesLegacyItemEnum = "24b"
    WlanDataratesLegacyItemEnum_ENUM2B  WlanDataratesLegacyItemEnum = "2b"
    WlanDataratesLegacyItemEnum_ENUM36  WlanDataratesLegacyItemEnum = "36"
    WlanDataratesLegacyItemEnum_ENUM36B WlanDataratesLegacyItemEnum = "36b"
    WlanDataratesLegacyItemEnum_ENUM48  WlanDataratesLegacyItemEnum = "48"
    WlanDataratesLegacyItemEnum_ENUM48B WlanDataratesLegacyItemEnum = "48b"
    WlanDataratesLegacyItemEnum_ENUM55  WlanDataratesLegacyItemEnum = "5.5"
    WlanDataratesLegacyItemEnum_ENUM55B WlanDataratesLegacyItemEnum = "5.5b"
    WlanDataratesLegacyItemEnum_ENUM54  WlanDataratesLegacyItemEnum = "54"
    WlanDataratesLegacyItemEnum_ENUM54B WlanDataratesLegacyItemEnum = "54b"
    WlanDataratesLegacyItemEnum_ENUM6   WlanDataratesLegacyItemEnum = "6"
    WlanDataratesLegacyItemEnum_ENUM6B  WlanDataratesLegacyItemEnum = "6b"
    WlanDataratesLegacyItemEnum_ENUM9   WlanDataratesLegacyItemEnum = "9"
    WlanDataratesLegacyItemEnum_ENUM9B  WlanDataratesLegacyItemEnum = "9b"
)

// WlanDataratesTemplateEnum is a string enum.
// Data Rates template to apply. enum: 
// * `no-legacy`: no 11b
// * `compatible`: all, like before, default setting that Broadcom/Atheros used
// * `legacy-only`: disable 802.11n and 802.11ac
// * `high-density`: no 11b, no low rates
// * `custom`: user defined
type WlanDataratesTemplateEnum string

const (
    WlanDataratesTemplateEnum_COMPATIBLE  WlanDataratesTemplateEnum = "compatible"
    WlanDataratesTemplateEnum_LEGACYONLY  WlanDataratesTemplateEnum = "legacy-only"
    WlanDataratesTemplateEnum_CUSTOM      WlanDataratesTemplateEnum = "custom"
    WlanDataratesTemplateEnum_NOLEGACY    WlanDataratesTemplateEnum = "no-legacy"
    WlanDataratesTemplateEnum_HIGHDENSITY WlanDataratesTemplateEnum = "high-density"
)

// WlanDynamicVlanTypeEnum is a string enum.
// standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco). enum: `airespace-interface-name`, `standard`
type WlanDynamicVlanTypeEnum string

const (
    WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME WlanDynamicVlanTypeEnum = "airespace-interface-name"
    WlanDynamicVlanTypeEnum_STANDARD               WlanDynamicVlanTypeEnum = "standard"
)

// WlanHotspot20OperatorsItemEnum is a string enum.
// enum: `ameriband`, `att`, `boingo`, `charter`, `eduroam`, `global_reach`, `google`, `hughes_systique`, `openroaming_legacy`, `openroaming_settled`, `openroaming_settlement_free`, `single_digits`, `tmobile`, `verizon`
type WlanHotspot20OperatorsItemEnum string

const (
    WlanHotspot20OperatorsItemEnum_AMERIBAND                 WlanHotspot20OperatorsItemEnum = "ameriband"
    WlanHotspot20OperatorsItemEnum_ATT                       WlanHotspot20OperatorsItemEnum = "att"
    WlanHotspot20OperatorsItemEnum_BOINGO                    WlanHotspot20OperatorsItemEnum = "boingo"
    WlanHotspot20OperatorsItemEnum_CHARTER                   WlanHotspot20OperatorsItemEnum = "charter"
    WlanHotspot20OperatorsItemEnum_EDUROAM                   WlanHotspot20OperatorsItemEnum = "eduroam"
    WlanHotspot20OperatorsItemEnum_GLOBALREACH               WlanHotspot20OperatorsItemEnum = "global_reach"
    WlanHotspot20OperatorsItemEnum_GOOGLE                    WlanHotspot20OperatorsItemEnum = "google"
    WlanHotspot20OperatorsItemEnum_HUGHESSYSTIQUE            WlanHotspot20OperatorsItemEnum = "hughes_systique"
    WlanHotspot20OperatorsItemEnum_OPENROAMINGLEGACY         WlanHotspot20OperatorsItemEnum = "openroaming_legacy"
    WlanHotspot20OperatorsItemEnum_OPENROAMINGSETTLED        WlanHotspot20OperatorsItemEnum = "openroaming_settled"
    WlanHotspot20OperatorsItemEnum_OPENROAMINGSETTLEMENTFREE WlanHotspot20OperatorsItemEnum = "openroaming_settlement_free"
    WlanHotspot20OperatorsItemEnum_SINGLEDIGITS              WlanHotspot20OperatorsItemEnum = "single_digits"
    WlanHotspot20OperatorsItemEnum_TMOBILE                   WlanHotspot20OperatorsItemEnum = "tmobile"
    WlanHotspot20OperatorsItemEnum_VERIZON                   WlanHotspot20OperatorsItemEnum = "verizon"
)

// WlanInterfaceEnum is a string enum.
// where this WLAN will be connected to. enum: `all`, `eth0`, `eth1`, `eth2`, `eth3`, `mxtunnel`, `site_mxedge`, `wxtunnel`
type WlanInterfaceEnum string

const (
    WlanInterfaceEnum_ALL        WlanInterfaceEnum = "all"
    WlanInterfaceEnum_ETH0       WlanInterfaceEnum = "eth0"
    WlanInterfaceEnum_ETH1       WlanInterfaceEnum = "eth1"
    WlanInterfaceEnum_ETH2       WlanInterfaceEnum = "eth2"
    WlanInterfaceEnum_ETH3       WlanInterfaceEnum = "eth3"
    WlanInterfaceEnum_MXTUNNEL   WlanInterfaceEnum = "mxtunnel"
    WlanInterfaceEnum_SITEMXEDGE WlanInterfaceEnum = "site_mxedge"
    WlanInterfaceEnum_WXTUNNEL   WlanInterfaceEnum = "wxtunnel"
)

// WlanPortalAuthEnum is a string enum.
// authentication scheme. enum: `amazon`, `azure`, `email`, `external`, `facebook`, `google`, `microsoft`, `multi`, `none`, `password`, `sponsor`, `sso`
type WlanPortalAuthEnum string

const (
    WlanPortalAuthEnum_AMAZON    WlanPortalAuthEnum = "amazon"
    WlanPortalAuthEnum_AZURE     WlanPortalAuthEnum = "azure"
    WlanPortalAuthEnum_EMAIL     WlanPortalAuthEnum = "email"
    WlanPortalAuthEnum_EXTERNAL  WlanPortalAuthEnum = "external"
    WlanPortalAuthEnum_FACEBOOK  WlanPortalAuthEnum = "facebook"
    WlanPortalAuthEnum_GOOGLE    WlanPortalAuthEnum = "google"
    WlanPortalAuthEnum_MICROSOFT WlanPortalAuthEnum = "microsoft"
    WlanPortalAuthEnum_MULTI     WlanPortalAuthEnum = "multi"
    WlanPortalAuthEnum_NONE      WlanPortalAuthEnum = "none"
    WlanPortalAuthEnum_PASSWORD  WlanPortalAuthEnum = "password"
    WlanPortalAuthEnum_SPONSOR   WlanPortalAuthEnum = "sponsor"
    WlanPortalAuthEnum_SSO       WlanPortalAuthEnum = "sso"
)

// WlanPortalIdpSignAlgoEnum is a string enum.
// Optioanl if `wlan_portal_auth`==`sso`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`
type WlanPortalIdpSignAlgoEnum string

const (
    WlanPortalIdpSignAlgoEnum_SHA1   WlanPortalIdpSignAlgoEnum = "sha1"
    WlanPortalIdpSignAlgoEnum_SHA256 WlanPortalIdpSignAlgoEnum = "sha256"
    WlanPortalIdpSignAlgoEnum_SHA384 WlanPortalIdpSignAlgoEnum = "sha384"
    WlanPortalIdpSignAlgoEnum_SHA512 WlanPortalIdpSignAlgoEnum = "sha512"
)

// WlanPortalSmsProviderEnum is a string enum.
// Optioanl if `sms_enabled`==`true`. enum: `broadnet`, `clickatell`, `gupshup`, `manual`, `puzzel`, `telstra`, `twilio`
type WlanPortalSmsProviderEnum string

const (
    WlanPortalSmsProviderEnum_BROADNET   WlanPortalSmsProviderEnum = "broadnet"
    WlanPortalSmsProviderEnum_CLICKATELL WlanPortalSmsProviderEnum = "clickatell"
    WlanPortalSmsProviderEnum_GUPSHUP    WlanPortalSmsProviderEnum = "gupshup"
    WlanPortalSmsProviderEnum_MANUAL     WlanPortalSmsProviderEnum = "manual"
    WlanPortalSmsProviderEnum_PUZZEL     WlanPortalSmsProviderEnum = "puzzel"
    WlanPortalSmsProviderEnum_TELSTRA    WlanPortalSmsProviderEnum = "telstra"
    WlanPortalSmsProviderEnum_TWILIO     WlanPortalSmsProviderEnum = "twilio"
)

// WlanPortalSsoNameidFormatEnum is a string enum.
// Optional if `wlan_portal_auth`==`sso`. enum: `email`, `unspecified`
type WlanPortalSsoNameidFormatEnum string

const (
    WlanPortalSsoNameidFormatEnum_EMAIL       WlanPortalSsoNameidFormatEnum = "email"
    WlanPortalSsoNameidFormatEnum_UNSPECIFIED WlanPortalSsoNameidFormatEnum = "unspecified"
)

// WlanQosClassEnum is a string enum.
// enum: `background`, `best_effort`, `video`, `voice`
type WlanQosClassEnum string

const (
    WlanQosClassEnum_BACKGROUND WlanQosClassEnum = "background"
    WlanQosClassEnum_BESTEFFORT WlanQosClassEnum = "best_effort"
    WlanQosClassEnum_VIDEO      WlanQosClassEnum = "video"
    WlanQosClassEnum_VOICE      WlanQosClassEnum = "voice"
)

// WlanRoamModeEnum is a string enum.
// enum: `11r`, `OKC`, `NONE`
type WlanRoamModeEnum string

const (
    WlanRoamModeEnum_ENUM11R WlanRoamModeEnum = "11r"
    WlanRoamModeEnum_NONE    WlanRoamModeEnum = "NONE"
    WlanRoamModeEnum_OKC     WlanRoamModeEnum = "OKC"
)

// WxlanRuleActionEnum is a string enum.
// type of action, allow / block. enum: `allow`, `block`
type WxlanRuleActionEnum string

const (
    WxlanRuleActionEnum_ALLOW WxlanRuleActionEnum = "allow"
    WxlanRuleActionEnum_BLOCK WxlanRuleActionEnum = "block"
)

// WxlanTagMatchEnum is a string enum.
// required if `type`==`match`. enum: `ap_id`, `app`, `asset_mac`, `client_mac`, `hostname`, `ip_range_subnet`, `port`, `psk_name`, `psk_role`, `radius_attr`, `radius_class`, `radius_group`, `radius_username`, `sdkclient_uuid`, `wlan_id`
type WxlanTagMatchEnum string

const (
    WxlanTagMatchEnum_APID           WxlanTagMatchEnum = "ap_id"
    WxlanTagMatchEnum_APP            WxlanTagMatchEnum = "app"
    WxlanTagMatchEnum_ASSETMAC       WxlanTagMatchEnum = "asset_mac"
    WxlanTagMatchEnum_CLIENTMAC      WxlanTagMatchEnum = "client_mac"
    WxlanTagMatchEnum_HOSTNAME       WxlanTagMatchEnum = "hostname"
    WxlanTagMatchEnum_IPRANGESUBNET  WxlanTagMatchEnum = "ip_range_subnet"
    WxlanTagMatchEnum_PORT           WxlanTagMatchEnum = "port"
    WxlanTagMatchEnum_PSKNAME        WxlanTagMatchEnum = "psk_name"
    WxlanTagMatchEnum_PSKROLE        WxlanTagMatchEnum = "psk_role"
    WxlanTagMatchEnum_RADIUSATTR     WxlanTagMatchEnum = "radius_attr"
    WxlanTagMatchEnum_RADIUSCLASS    WxlanTagMatchEnum = "radius_class"
    WxlanTagMatchEnum_RADIUSGROUP    WxlanTagMatchEnum = "radius_group"
    WxlanTagMatchEnum_RADIUSUSERNAME WxlanTagMatchEnum = "radius_username"
    WxlanTagMatchEnum_SDKCLIENTUUID  WxlanTagMatchEnum = "sdkclient_uuid"
    WxlanTagMatchEnum_WLANID         WxlanTagMatchEnum = "wlan_id"
)

// WxlanTagOperationEnum is a string enum.
// required if `type`==`match`, type of tag (inclusive/exclusive). enum: `in`, `not_in`
type WxlanTagOperationEnum string

const (
    WxlanTagOperationEnum_IN    WxlanTagOperationEnum = "in"
    WxlanTagOperationEnum_NOTIN WxlanTagOperationEnum = "not_in"
)

// WxlanTagTypeEnum is a string enum.
// enum: `client`, `match`, `resource`, `spec`, `subnet`, `vlan`
type WxlanTagTypeEnum string

const (
    WxlanTagTypeEnum_CLIENT   WxlanTagTypeEnum = "client"
    WxlanTagTypeEnum_MATCH    WxlanTagTypeEnum = "match"
    WxlanTagTypeEnum_RESOURCE WxlanTagTypeEnum = "resource"
    WxlanTagTypeEnum_SPEC     WxlanTagTypeEnum = "spec"
    WxlanTagTypeEnum_SUBNET   WxlanTagTypeEnum = "subnet"
    WxlanTagTypeEnum_VLAN     WxlanTagTypeEnum = "vlan"
)

// WxlanTunnelSessionEthertypeEnum is a string enum.
// enum: `ethernet`, `vlan`
type WxlanTunnelSessionEthertypeEnum string

const (
    WxlanTunnelSessionEthertypeEnum_ETHERNET WxlanTunnelSessionEthertypeEnum = "ethernet"
    WxlanTunnelSessionEthertypeEnum_VLAN     WxlanTunnelSessionEthertypeEnum = "vlan"
)

// ZoneScopeEnum is a string enum.
// enum: `map`, `rssizone`, `site`, `zone`
type ZoneScopeEnum string

const (
    ZoneScopeEnum_ENUMMAP  ZoneScopeEnum = "map"
    ZoneScopeEnum_RSSIZONE ZoneScopeEnum = "rssizone"
    ZoneScopeEnum_SITE     ZoneScopeEnum = "site"
    ZoneScopeEnum_ZONE     ZoneScopeEnum = "zone"
)

// ZoneTypeEnum is a string enum.
// enum: `rssizones`, `zones`
type ZoneTypeEnum string

const (
    ZoneTypeEnum_RSSIZONES ZoneTypeEnum = "rssizones"
    ZoneTypeEnum_ZONES     ZoneTypeEnum = "zones"
)
