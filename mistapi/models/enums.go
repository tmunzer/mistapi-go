package models

// CountOrgDevicesMxtunnelStatusEnum is a string enum.
type CountOrgDevicesMxtunnelStatusEnum string

const (
    CountOrgDevicesMxtunnelStatusEnum_UP   CountOrgDevicesMxtunnelStatusEnum = "up"
    CountOrgDevicesMxtunnelStatusEnum_DOWN CountOrgDevicesMxtunnelStatusEnum = "down"
)

// CountOrgSwitchPortsAuthStateEnum is a string enum.
type CountOrgSwitchPortsAuthStateEnum string

const (
    CountOrgSwitchPortsAuthStateEnum_INIT           CountOrgSwitchPortsAuthStateEnum = "init"
    CountOrgSwitchPortsAuthStateEnum_AUTHENTICATED  CountOrgSwitchPortsAuthStateEnum = "authenticated"
    CountOrgSwitchPortsAuthStateEnum_AUTHENTICATING CountOrgSwitchPortsAuthStateEnum = "authenticating"
    CountOrgSwitchPortsAuthStateEnum_HELD           CountOrgSwitchPortsAuthStateEnum = "held"
)

// CountOrgSwitchPortsStpRoleEnum is a string enum.
type CountOrgSwitchPortsStpRoleEnum string

const (
    CountOrgSwitchPortsStpRoleEnum_DESIGNATED    CountOrgSwitchPortsStpRoleEnum = "designated"
    CountOrgSwitchPortsStpRoleEnum_BACKUP        CountOrgSwitchPortsStpRoleEnum = "backup"
    CountOrgSwitchPortsStpRoleEnum_ALTERNATE     CountOrgSwitchPortsStpRoleEnum = "alternate"
    CountOrgSwitchPortsStpRoleEnum_ROOT          CountOrgSwitchPortsStpRoleEnum = "root"
    CountOrgSwitchPortsStpRoleEnum_ROOTPREVENTED CountOrgSwitchPortsStpRoleEnum = "root-prevented"
)

// CountOrgSwitchPortsStpStateEnum is a string enum.
type CountOrgSwitchPortsStpStateEnum string

const (
    CountOrgSwitchPortsStpStateEnum_FORWARDING CountOrgSwitchPortsStpStateEnum = "forwarding"
    CountOrgSwitchPortsStpStateEnum_BLOCKING   CountOrgSwitchPortsStpStateEnum = "blocking"
    CountOrgSwitchPortsStpStateEnum_LEARNING   CountOrgSwitchPortsStpStateEnum = "learning"
    CountOrgSwitchPortsStpStateEnum_LISTENING  CountOrgSwitchPortsStpStateEnum = "listening"
    CountOrgSwitchPortsStpStateEnum_DISABLED   CountOrgSwitchPortsStpStateEnum = "disabled"
)

// DeviceStatusEnum is a string enum.
type DeviceStatusEnum string

const (
    DeviceStatusEnum_ALL          DeviceStatusEnum = "all"
    DeviceStatusEnum_CONNECTED    DeviceStatusEnum = "connected"
    DeviceStatusEnum_DISCONNECTED DeviceStatusEnum = "disconnected"
)

// DeviceTypeEnum is a string enum.
type DeviceTypeEnum string

const (
    DeviceTypeEnum_AP         DeviceTypeEnum = "ap"
    DeviceTypeEnum_ENUMSWITCH DeviceTypeEnum = "switch"
    DeviceTypeEnum_GATEWAY    DeviceTypeEnum = "gateway"
)

// DeviceTypeWithAllEnum is a string enum.
type DeviceTypeWithAllEnum string

const (
    DeviceTypeWithAllEnum_AP         DeviceTypeWithAllEnum = "ap"
    DeviceTypeWithAllEnum_ENUMSWITCH DeviceTypeWithAllEnum = "switch"
    DeviceTypeWithAllEnum_GATEWAY    DeviceTypeWithAllEnum = "gateway"
    DeviceTypeWithAllEnum_ALL        DeviceTypeWithAllEnum = "all"
)

// GetOrgMxedgeUpgradeInfoChannelEnum is a string enum.
type GetOrgMxedgeUpgradeInfoChannelEnum string

const (
    GetOrgMxedgeUpgradeInfoChannelEnum_STABLE GetOrgMxedgeUpgradeInfoChannelEnum = "stable"
    GetOrgMxedgeUpgradeInfoChannelEnum_BETA   GetOrgMxedgeUpgradeInfoChannelEnum = "beta"
    GetOrgMxedgeUpgradeInfoChannelEnum_ALPHA  GetOrgMxedgeUpgradeInfoChannelEnum = "alpha"
)

// ListOrgLogsSortEnum is a string enum.
type ListOrgLogsSortEnum string

const (
    ListOrgLogsSortEnum_TIMESTAMP     ListOrgLogsSortEnum = "timestamp"
    ListOrgLogsSortEnum_ENUMTIMESTAMP ListOrgLogsSortEnum = "-timestamp"
    ListOrgLogsSortEnum_SITEID        ListOrgLogsSortEnum = "site_id"
    ListOrgLogsSortEnum_ADMINID       ListOrgLogsSortEnum = "admin_id"
)

// MxedgeForSiteEnum is a string enum.
type MxedgeForSiteEnum string

const (
    MxedgeForSiteEnum_ANY   MxedgeForSiteEnum = "any"
    MxedgeForSiteEnum_TRUE  MxedgeForSiteEnum = "true"
    MxedgeForSiteEnum_FALSE MxedgeForSiteEnum = "false"
)

// MxedgeServiceActionEnum is a string enum.
type MxedgeServiceActionEnum string

const (
    MxedgeServiceActionEnum_RESTART MxedgeServiceActionEnum = "restart"
    MxedgeServiceActionEnum_START   MxedgeServiceActionEnum = "start"
    MxedgeServiceActionEnum_STOP    MxedgeServiceActionEnum = "stop"
)

// MxedgeServiceNameEnum is a string enum.
type MxedgeServiceNameEnum string

const (
    MxedgeServiceNameEnum_TUNTERM     MxedgeServiceNameEnum = "tunterm"
    MxedgeServiceNameEnum_RADSECPROXY MxedgeServiceNameEnum = "radsecproxy"
    MxedgeServiceNameEnum_MXAGENT     MxedgeServiceNameEnum = "mxagent"
    MxedgeServiceNameEnum_MXOCPROXY   MxedgeServiceNameEnum = "mxocproxy"
    MxedgeServiceNameEnum_MXDAS       MxedgeServiceNameEnum = "mxdas"
)

// OauthAppNameEnum is a string enum.
type OauthAppNameEnum string

const (
    OauthAppNameEnum_ZOOM   OauthAppNameEnum = "zoom"
    OauthAppNameEnum_TEAMS  OauthAppNameEnum = "teams"
    OauthAppNameEnum_INTUNE OauthAppNameEnum = "intune"
    OauthAppNameEnum_JAMF   OauthAppNameEnum = "jamf"
    OauthAppNameEnum_VMWARE OauthAppNameEnum = "vmware"
)

// OrgAssetCountDistinctEnum is a string enum.
type OrgAssetCountDistinctEnum string

const (
    OrgAssetCountDistinctEnum_SITEID       OrgAssetCountDistinctEnum = "site_id"
    OrgAssetCountDistinctEnum_MAC          OrgAssetCountDistinctEnum = "mac"
    OrgAssetCountDistinctEnum_MAPID        OrgAssetCountDistinctEnum = "map_id"
    OrgAssetCountDistinctEnum_IBEACONUUID  OrgAssetCountDistinctEnum = "ibeacon_uuid"
    OrgAssetCountDistinctEnum_IBEACONMAJOR OrgAssetCountDistinctEnum = "ibeacon_major"
    OrgAssetCountDistinctEnum_IBEACONMINOR OrgAssetCountDistinctEnum = "ibeacon_minor"
)

// OrgClientSessionsCountDistinctEnum is a string enum.
type OrgClientSessionsCountDistinctEnum string

const (
    OrgClientSessionsCountDistinctEnum_SSID     OrgClientSessionsCountDistinctEnum = "ssid"
    OrgClientSessionsCountDistinctEnum_AP       OrgClientSessionsCountDistinctEnum = "ap"
    OrgClientSessionsCountDistinctEnum_IP       OrgClientSessionsCountDistinctEnum = "ip"
    OrgClientSessionsCountDistinctEnum_VLAN     OrgClientSessionsCountDistinctEnum = "vlan"
    OrgClientSessionsCountDistinctEnum_HOSTNAME OrgClientSessionsCountDistinctEnum = "hostname"
    OrgClientSessionsCountDistinctEnum_OS       OrgClientSessionsCountDistinctEnum = "os"
    OrgClientSessionsCountDistinctEnum_MODEL    OrgClientSessionsCountDistinctEnum = "model"
    OrgClientSessionsCountDistinctEnum_DEVICE   OrgClientSessionsCountDistinctEnum = "device"
)

// OrgClientsCountDistinctEnum is a string enum.
type OrgClientsCountDistinctEnum string

const (
    OrgClientsCountDistinctEnum_MAC      OrgClientsCountDistinctEnum = "mac"
    OrgClientsCountDistinctEnum_HOSTNAME OrgClientsCountDistinctEnum = "hostname"
    OrgClientsCountDistinctEnum_DEVICE   OrgClientsCountDistinctEnum = "device"
    OrgClientsCountDistinctEnum_OS       OrgClientsCountDistinctEnum = "os"
    OrgClientsCountDistinctEnum_MODEL    OrgClientsCountDistinctEnum = "model"
    OrgClientsCountDistinctEnum_AP       OrgClientsCountDistinctEnum = "ap"
    OrgClientsCountDistinctEnum_VLAN     OrgClientsCountDistinctEnum = "vlan"
    OrgClientsCountDistinctEnum_SSID     OrgClientsCountDistinctEnum = "ssid"
    OrgClientsCountDistinctEnum_IP       OrgClientsCountDistinctEnum = "ip"
)

// OrgDevicesCountDistinctEnum is a string enum.
type OrgDevicesCountDistinctEnum string

const (
    OrgDevicesCountDistinctEnum_HOSTNAME       OrgDevicesCountDistinctEnum = "hostname"
    OrgDevicesCountDistinctEnum_SITEID         OrgDevicesCountDistinctEnum = "site_id"
    OrgDevicesCountDistinctEnum_MODEL          OrgDevicesCountDistinctEnum = "model"
    OrgDevicesCountDistinctEnum_MAC            OrgDevicesCountDistinctEnum = "mac"
    OrgDevicesCountDistinctEnum_VERSION        OrgDevicesCountDistinctEnum = "version"
    OrgDevicesCountDistinctEnum_IP             OrgDevicesCountDistinctEnum = "ip"
    OrgDevicesCountDistinctEnum_MXTUNNELSTATUS OrgDevicesCountDistinctEnum = "mxtunnel_status"
    OrgDevicesCountDistinctEnum_MXEDGEID       OrgDevicesCountDistinctEnum = "mxedge_id"
    OrgDevicesCountDistinctEnum_LLDPSYSTEMNAME OrgDevicesCountDistinctEnum = "lldp_system_name"
    OrgDevicesCountDistinctEnum_LLDPSYSTEMDESC OrgDevicesCountDistinctEnum = "lldp_system_desc"
    OrgDevicesCountDistinctEnum_LLDPPORTID     OrgDevicesCountDistinctEnum = "lldp_port_id"
    OrgDevicesCountDistinctEnum_LLDPMGMTADDR   OrgDevicesCountDistinctEnum = "lldp_mgmt_addr"
)

// OrgDevicesEventsCountDistinctEnum is a string enum.
type OrgDevicesEventsCountDistinctEnum string

const (
    OrgDevicesEventsCountDistinctEnum_ORGID     OrgDevicesEventsCountDistinctEnum = "org_id"
    OrgDevicesEventsCountDistinctEnum_SITEID    OrgDevicesEventsCountDistinctEnum = "site_id"
    OrgDevicesEventsCountDistinctEnum_AP        OrgDevicesEventsCountDistinctEnum = "ap"
    OrgDevicesEventsCountDistinctEnum_APFW      OrgDevicesEventsCountDistinctEnum = "apfw"
    OrgDevicesEventsCountDistinctEnum_MODEL     OrgDevicesEventsCountDistinctEnum = "model"
    OrgDevicesEventsCountDistinctEnum_TEXT      OrgDevicesEventsCountDistinctEnum = "text"
    OrgDevicesEventsCountDistinctEnum_TIMESTAMP OrgDevicesEventsCountDistinctEnum = "timestamp"
    OrgDevicesEventsCountDistinctEnum_ENUMTYPE  OrgDevicesEventsCountDistinctEnum = "type"
)

// OrgDevicesLastConfigsCountDistinctEnum is a string enum.
type OrgDevicesLastConfigsCountDistinctEnum string

const (
    OrgDevicesLastConfigsCountDistinctEnum_MAC     OrgDevicesLastConfigsCountDistinctEnum = "mac"
    OrgDevicesLastConfigsCountDistinctEnum_VERSION OrgDevicesLastConfigsCountDistinctEnum = "version"
    OrgDevicesLastConfigsCountDistinctEnum_NAME    OrgDevicesLastConfigsCountDistinctEnum = "name"
    OrgDevicesLastConfigsCountDistinctEnum_SITEID  OrgDevicesLastConfigsCountDistinctEnum = "site_id"
)

// OrgGuestsCountDistinctEnum is a string enum.
type OrgGuestsCountDistinctEnum string

const (
    OrgGuestsCountDistinctEnum_AUTHMETHOD OrgGuestsCountDistinctEnum = "auth_method"
    OrgGuestsCountDistinctEnum_SSID       OrgGuestsCountDistinctEnum = "ssid"
    OrgGuestsCountDistinctEnum_COMPANY    OrgGuestsCountDistinctEnum = "company"
)

// OrgLogsCountDistinctEnum is a string enum.
type OrgLogsCountDistinctEnum string

const (
    OrgLogsCountDistinctEnum_ADMINID   OrgLogsCountDistinctEnum = "admin_id"
    OrgLogsCountDistinctEnum_ADMINNAME OrgLogsCountDistinctEnum = "admin_name"
    OrgLogsCountDistinctEnum_MESSAGE   OrgLogsCountDistinctEnum = "message"
    OrgLogsCountDistinctEnum_SITEID    OrgLogsCountDistinctEnum = "site_id"
)

// OrgMxedgeCountDistinctEnum is a string enum.
type OrgMxedgeCountDistinctEnum string

const (
    OrgMxedgeCountDistinctEnum_MODEL          OrgMxedgeCountDistinctEnum = "model"
    OrgMxedgeCountDistinctEnum_MXCLUSTERID    OrgMxedgeCountDistinctEnum = "mxcluster_id"
    OrgMxedgeCountDistinctEnum_DISTRO         OrgMxedgeCountDistinctEnum = "distro"
    OrgMxedgeCountDistinctEnum_TUNTERMVERSION OrgMxedgeCountDistinctEnum = "tunterm_version"
    OrgMxedgeCountDistinctEnum_SITEID         OrgMxedgeCountDistinctEnum = "site_id"
)

// OrgMxedgeEventsCountDistinctEnum is a string enum.
type OrgMxedgeEventsCountDistinctEnum string

const (
    OrgMxedgeEventsCountDistinctEnum_MXEDGEID    OrgMxedgeEventsCountDistinctEnum = "mxedge_id"
    OrgMxedgeEventsCountDistinctEnum_ENUMTYPE    OrgMxedgeEventsCountDistinctEnum = "type"
    OrgMxedgeEventsCountDistinctEnum_MXCLUSTERID OrgMxedgeEventsCountDistinctEnum = "mxcluster_id"
    OrgMxedgeEventsCountDistinctEnum_ENUMPACKAGE OrgMxedgeEventsCountDistinctEnum = "package"
)

// OrgNacClientEventsCountDistinctEnum is a string enum.
type OrgNacClientEventsCountDistinctEnum string

const (
    OrgNacClientEventsCountDistinctEnum_ENUMTYPE        OrgNacClientEventsCountDistinctEnum = "type"
    OrgNacClientEventsCountDistinctEnum_NACRULEID       OrgNacClientEventsCountDistinctEnum = "nacrule_id"
    OrgNacClientEventsCountDistinctEnum_DRYRUNNACRULEID OrgNacClientEventsCountDistinctEnum = "dryrun_nacrule_id"
    OrgNacClientEventsCountDistinctEnum_AUTHTYPE        OrgNacClientEventsCountDistinctEnum = "auth_type"
    OrgNacClientEventsCountDistinctEnum_VLAN            OrgNacClientEventsCountDistinctEnum = "vlan"
    OrgNacClientEventsCountDistinctEnum_NASVENDOR       OrgNacClientEventsCountDistinctEnum = "nas_vendor"
    OrgNacClientEventsCountDistinctEnum_USERNAME        OrgNacClientEventsCountDistinctEnum = "username"
    OrgNacClientEventsCountDistinctEnum_AP              OrgNacClientEventsCountDistinctEnum = "ap"
    OrgNacClientEventsCountDistinctEnum_MAC             OrgNacClientEventsCountDistinctEnum = "mac"
    OrgNacClientEventsCountDistinctEnum_SSID            OrgNacClientEventsCountDistinctEnum = "ssid"
)

// OrgNacClientsCountDistinctEnum is a string enum.
type OrgNacClientsCountDistinctEnum string

const (
    OrgNacClientsCountDistinctEnum_ENUMTYPE      OrgNacClientsCountDistinctEnum = "type"
    OrgNacClientsCountDistinctEnum_LASTNACRULEID OrgNacClientsCountDistinctEnum = "last_nacrule_id"
    OrgNacClientsCountDistinctEnum_AUTHTYPE      OrgNacClientsCountDistinctEnum = "auth_type"
    OrgNacClientsCountDistinctEnum_LASTVLAN      OrgNacClientsCountDistinctEnum = "last_vlan"
    OrgNacClientsCountDistinctEnum_LASTNASVENDOR OrgNacClientsCountDistinctEnum = "last_nas_vendor"
    OrgNacClientsCountDistinctEnum_LASTUSERNAME  OrgNacClientsCountDistinctEnum = "last_username"
    OrgNacClientsCountDistinctEnum_LASTAP        OrgNacClientsCountDistinctEnum = "last_ap"
    OrgNacClientsCountDistinctEnum_MAC           OrgNacClientsCountDistinctEnum = "mac"
    OrgNacClientsCountDistinctEnum_LASTSSID      OrgNacClientsCountDistinctEnum = "last_ssid"
    OrgNacClientsCountDistinctEnum_LASTSTATUS    OrgNacClientsCountDistinctEnum = "last_status"
    OrgNacClientsCountDistinctEnum_MDMCOMPLIANCE OrgNacClientsCountDistinctEnum = "mdm_compliance"
    OrgNacClientsCountDistinctEnum_MDMPROVIDER   OrgNacClientsCountDistinctEnum = "mdm_provider"
)

// OrgOtherdevicesEventsCountDistinctEnum is a string enum.
type OrgOtherdevicesEventsCountDistinctEnum string

const (
    OrgOtherdevicesEventsCountDistinctEnum_MAC      OrgOtherdevicesEventsCountDistinctEnum = "mac"
    OrgOtherdevicesEventsCountDistinctEnum_ENUMTYPE OrgOtherdevicesEventsCountDistinctEnum = "type"
    OrgOtherdevicesEventsCountDistinctEnum_VENDOR   OrgOtherdevicesEventsCountDistinctEnum = "vendor"
    OrgOtherdevicesEventsCountDistinctEnum_SITEID   OrgOtherdevicesEventsCountDistinctEnum = "site_id"
)

// OrgPskPortalLogsCountDistinctEnum is a string enum.
type OrgPskPortalLogsCountDistinctEnum string

const (
    OrgPskPortalLogsCountDistinctEnum_ADMINID     OrgPskPortalLogsCountDistinctEnum = "admin_id"
    OrgPskPortalLogsCountDistinctEnum_ADMINNAME   OrgPskPortalLogsCountDistinctEnum = "admin_name"
    OrgPskPortalLogsCountDistinctEnum_PSKNAME     OrgPskPortalLogsCountDistinctEnum = "psk_name"
    OrgPskPortalLogsCountDistinctEnum_PSKID       OrgPskPortalLogsCountDistinctEnum = "psk_id"
    OrgPskPortalLogsCountDistinctEnum_PSKPORTALID OrgPskPortalLogsCountDistinctEnum = "pskportal_id"
    OrgPskPortalLogsCountDistinctEnum_USERID      OrgPskPortalLogsCountDistinctEnum = "user_id"
)

// OrgSiteSleTypeEnum is a string enum.
type OrgSiteSleTypeEnum string

const (
    OrgSiteSleTypeEnum_WAN   OrgSiteSleTypeEnum = "wan"
    OrgSiteSleTypeEnum_WIRED OrgSiteSleTypeEnum = "wired"
    OrgSiteSleTypeEnum_WIFI  OrgSiteSleTypeEnum = "wifi"
)

// OrgSitesCountDistinctEnum is a string enum.
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
    OrgSitesCountDistinctEnum_ROGUEENABLED        OrgSitesCountDistinctEnum = "rogue_enabled"
    OrgSitesCountDistinctEnum_REMOTESYSLOGENABLED OrgSitesCountDistinctEnum = "remote_syslog_enabled"
    OrgSitesCountDistinctEnum_RTSAENABLED         OrgSitesCountDistinctEnum = "rtsa_enabled"
    OrgSitesCountDistinctEnum_VNAENABLED          OrgSitesCountDistinctEnum = "vna_enabled"
    OrgSitesCountDistinctEnum_WIFIENABLED         OrgSitesCountDistinctEnum = "wifi_enabled"
)

// OrgSwitchPortCountDistinctEnum is a string enum.
type OrgSwitchPortCountDistinctEnum string

const (
    OrgSwitchPortCountDistinctEnum_PORTID             OrgSwitchPortCountDistinctEnum = "port_id"
    OrgSwitchPortCountDistinctEnum_PORTMAC            OrgSwitchPortCountDistinctEnum = "port_mac"
    OrgSwitchPortCountDistinctEnum_FULLDUPLEX         OrgSwitchPortCountDistinctEnum = "full_duplex"
    OrgSwitchPortCountDistinctEnum_MAC                OrgSwitchPortCountDistinctEnum = "mac"
    OrgSwitchPortCountDistinctEnum_NEIGHBORMAC        OrgSwitchPortCountDistinctEnum = "neighbor_mac"
    OrgSwitchPortCountDistinctEnum_NEIGHBORPORTDESC   OrgSwitchPortCountDistinctEnum = "neighbor_port_desc"
    OrgSwitchPortCountDistinctEnum_NEIGHBORSYSTEMNAME OrgSwitchPortCountDistinctEnum = "neighbor_system_name"
    OrgSwitchPortCountDistinctEnum_POEDISABLED        OrgSwitchPortCountDistinctEnum = "poe_disabled"
    OrgSwitchPortCountDistinctEnum_POEMODE            OrgSwitchPortCountDistinctEnum = "poe_mode"
    OrgSwitchPortCountDistinctEnum_POEON              OrgSwitchPortCountDistinctEnum = "poe_on"
    OrgSwitchPortCountDistinctEnum_SPEED              OrgSwitchPortCountDistinctEnum = "speed"
    OrgSwitchPortCountDistinctEnum_UP                 OrgSwitchPortCountDistinctEnum = "up"
)

// OrgTicketsCountDistinctEnum is a string enum.
type OrgTicketsCountDistinctEnum string

const (
    OrgTicketsCountDistinctEnum_STATUS   OrgTicketsCountDistinctEnum = "status"
    OrgTicketsCountDistinctEnum_ENUMTYPE OrgTicketsCountDistinctEnum = "type"
)

// OrgTunnelCountDistinctEnum is a string enum.
type OrgTunnelCountDistinctEnum string

const (
    OrgTunnelCountDistinctEnum_AUTHALGO     OrgTunnelCountDistinctEnum = "auth_algo"
    OrgTunnelCountDistinctEnum_WXTUNNELID   OrgTunnelCountDistinctEnum = "wxtunnel_id"
    OrgTunnelCountDistinctEnum_AP           OrgTunnelCountDistinctEnum = "ap"
    OrgTunnelCountDistinctEnum_REMOTEIP     OrgTunnelCountDistinctEnum = "remote_ip"
    OrgTunnelCountDistinctEnum_REMOTEPORT   OrgTunnelCountDistinctEnum = "remote_port"
    OrgTunnelCountDistinctEnum_STATE        OrgTunnelCountDistinctEnum = "state"
    OrgTunnelCountDistinctEnum_MXEDGEID     OrgTunnelCountDistinctEnum = "mxedge_id"
    OrgTunnelCountDistinctEnum_MXCLUSTERID  OrgTunnelCountDistinctEnum = "mxcluster_id"
    OrgTunnelCountDistinctEnum_SITEID       OrgTunnelCountDistinctEnum = "site_id"
    OrgTunnelCountDistinctEnum_PEERMXEDGEID OrgTunnelCountDistinctEnum = "peer_mxedge_id"
    OrgTunnelCountDistinctEnum_MAC          OrgTunnelCountDistinctEnum = "mac"
    OrgTunnelCountDistinctEnum_NODE         OrgTunnelCountDistinctEnum = "node"
    OrgTunnelCountDistinctEnum_PEERIP       OrgTunnelCountDistinctEnum = "peer_ip"
    OrgTunnelCountDistinctEnum_PEERHOST     OrgTunnelCountDistinctEnum = "peer_host"
    OrgTunnelCountDistinctEnum_IP           OrgTunnelCountDistinctEnum = "ip"
    OrgTunnelCountDistinctEnum_TUNNELNAME   OrgTunnelCountDistinctEnum = "tunnel_name"
    OrgTunnelCountDistinctEnum_PROTOCOL     OrgTunnelCountDistinctEnum = "protocol"
    OrgTunnelCountDistinctEnum_ENCRYPTALGO  OrgTunnelCountDistinctEnum = "encrypt_algo"
    OrgTunnelCountDistinctEnum_IKEVERSION   OrgTunnelCountDistinctEnum = "ike_version"
    OrgTunnelCountDistinctEnum_LASTEVENT    OrgTunnelCountDistinctEnum = "last_event"
    OrgTunnelCountDistinctEnum_UP           OrgTunnelCountDistinctEnum = "up"
)

// OrgTunnelTypeCountEnum is a string enum.
type OrgTunnelTypeCountEnum string

const (
    OrgTunnelTypeCountEnum_WXTUNNEL OrgTunnelTypeCountEnum = "wxtunnel"
    OrgTunnelTypeCountEnum_WAN      OrgTunnelTypeCountEnum = "wan"
)

// OrgWanClientsCountDistinctEnum is a string enum.
type OrgWanClientsCountDistinctEnum string

const (
    OrgWanClientsCountDistinctEnum_HOSTNAME OrgWanClientsCountDistinctEnum = "hostname"
    OrgWanClientsCountDistinctEnum_IP       OrgWanClientsCountDistinctEnum = "ip"
    OrgWanClientsCountDistinctEnum_MFG      OrgWanClientsCountDistinctEnum = "mfg"
    OrgWanClientsCountDistinctEnum_MAC      OrgWanClientsCountDistinctEnum = "mac"
)

// OrgWanClientsEventsCountDistinctEnum is a string enum.
type OrgWanClientsEventsCountDistinctEnum string

const (
    OrgWanClientsEventsCountDistinctEnum_ENUMTYPE OrgWanClientsEventsCountDistinctEnum = "type"
    OrgWanClientsEventsCountDistinctEnum_HOSTNAME OrgWanClientsEventsCountDistinctEnum = "hostname"
    OrgWanClientsEventsCountDistinctEnum_IP       OrgWanClientsEventsCountDistinctEnum = "ip"
    OrgWanClientsEventsCountDistinctEnum_MFG      OrgWanClientsEventsCountDistinctEnum = "mfg"
    OrgWanClientsEventsCountDistinctEnum_MAC      OrgWanClientsEventsCountDistinctEnum = "mac"
)

// OrgWiredClientsCountDistinctEnum is a string enum.
type OrgWiredClientsCountDistinctEnum string

const (
    OrgWiredClientsCountDistinctEnum_PORTID    OrgWiredClientsCountDistinctEnum = "port_id"
    OrgWiredClientsCountDistinctEnum_VLAN      OrgWiredClientsCountDistinctEnum = "vlan"
    OrgWiredClientsCountDistinctEnum_MAC       OrgWiredClientsCountDistinctEnum = "mac"
    OrgWiredClientsCountDistinctEnum_DEVICEMAC OrgWiredClientsCountDistinctEnum = "device_mac"
    OrgWiredClientsCountDistinctEnum_SITEID    OrgWiredClientsCountDistinctEnum = "site_id"
    OrgWiredClientsCountDistinctEnum_ENUMTYPE  OrgWiredClientsCountDistinctEnum = "type"
)

// SearchOrgDevicesMxtunnelStatusEnum is a string enum.
type SearchOrgDevicesMxtunnelStatusEnum string

const (
    SearchOrgDevicesMxtunnelStatusEnum_UP   SearchOrgDevicesMxtunnelStatusEnum = "up"
    SearchOrgDevicesMxtunnelStatusEnum_DOWN SearchOrgDevicesMxtunnelStatusEnum = "down"
)

// SearchOrgSwOrGwPortsAuthStateEnum is a string enum.
type SearchOrgSwOrGwPortsAuthStateEnum string

const (
    SearchOrgSwOrGwPortsAuthStateEnum_INIT           SearchOrgSwOrGwPortsAuthStateEnum = "init"
    SearchOrgSwOrGwPortsAuthStateEnum_AUTHENTICATED  SearchOrgSwOrGwPortsAuthStateEnum = "authenticated"
    SearchOrgSwOrGwPortsAuthStateEnum_AUTHENTICATING SearchOrgSwOrGwPortsAuthStateEnum = "authenticating"
    SearchOrgSwOrGwPortsAuthStateEnum_HELD           SearchOrgSwOrGwPortsAuthStateEnum = "held"
)

// SearchOrgSwOrGwPortsStpRoleEnum is a string enum.
type SearchOrgSwOrGwPortsStpRoleEnum string

const (
    SearchOrgSwOrGwPortsStpRoleEnum_DESIGNATED    SearchOrgSwOrGwPortsStpRoleEnum = "designated"
    SearchOrgSwOrGwPortsStpRoleEnum_BACKUP        SearchOrgSwOrGwPortsStpRoleEnum = "backup"
    SearchOrgSwOrGwPortsStpRoleEnum_ALTERNATE     SearchOrgSwOrGwPortsStpRoleEnum = "alternate"
    SearchOrgSwOrGwPortsStpRoleEnum_ROOT          SearchOrgSwOrGwPortsStpRoleEnum = "root"
    SearchOrgSwOrGwPortsStpRoleEnum_ROOTPREVENTED SearchOrgSwOrGwPortsStpRoleEnum = "root-prevented"
)

// SearchOrgSwOrGwPortsStpStateEnum is a string enum.
type SearchOrgSwOrGwPortsStpStateEnum string

const (
    SearchOrgSwOrGwPortsStpStateEnum_FORWARDING SearchOrgSwOrGwPortsStpStateEnum = "forwarding"
    SearchOrgSwOrGwPortsStpStateEnum_BLOCKING   SearchOrgSwOrGwPortsStpStateEnum = "blocking"
    SearchOrgSwOrGwPortsStpStateEnum_LEARNING   SearchOrgSwOrGwPortsStpStateEnum = "learning"
    SearchOrgSwOrGwPortsStpStateEnum_LISTENING  SearchOrgSwOrGwPortsStpStateEnum = "listening"
    SearchOrgSwOrGwPortsStpStateEnum_DISABLED   SearchOrgSwOrGwPortsStpStateEnum = "disabled"
)

// SuppressedAlarmScopeEnum is a string enum.
// level of scope
type SuppressedAlarmScopeEnum string

const (
    SuppressedAlarmScopeEnum_ORG  SuppressedAlarmScopeEnum = "org"
    SuppressedAlarmScopeEnum_SITE SuppressedAlarmScopeEnum = "site"
)

// TroubleshootTypeEnum is a string enum.
type TroubleshootTypeEnum string

const (
    TroubleshootTypeEnum_WIRELESS TroubleshootTypeEnum = "wireless"
    TroubleshootTypeEnum_WIRED    TroubleshootTypeEnum = "wired"
    TroubleshootTypeEnum_WAN      TroubleshootTypeEnum = "wan"
)

// TunnelTypeEnum is a string enum.
type TunnelTypeEnum string

const (
    TunnelTypeEnum_WXTUNNEL TunnelTypeEnum = "wxtunnel"
    TunnelTypeEnum_WAN      TunnelTypeEnum = "wan"
)

// VarSourceEnum is a string enum.
type VarSourceEnum string

const (
    VarSourceEnum_SITE          VarSourceEnum = "site"
    VarSourceEnum_DEVICEPROFILE VarSourceEnum = "deviceprofile"
)

// WebhookDeliveryDistinctEnum is a string enum.
// webhook topic
type WebhookDeliveryDistinctEnum string

const (
    WebhookDeliveryDistinctEnum_STATUS     WebhookDeliveryDistinctEnum = "status"
    WebhookDeliveryDistinctEnum_TOPIC      WebhookDeliveryDistinctEnum = "topic"
    WebhookDeliveryDistinctEnum_STATUSCODE WebhookDeliveryDistinctEnum = "status_code"
    WebhookDeliveryDistinctEnum_WEBHOOKID  WebhookDeliveryDistinctEnum = "webhook_id"
)

// WebhookDeliveryStatusEnum is a string enum.
// webhook delivery status
type WebhookDeliveryStatusEnum string

const (
    WebhookDeliveryStatusEnum_SUCCESS WebhookDeliveryStatusEnum = "success"
    WebhookDeliveryStatusEnum_FAILURE WebhookDeliveryStatusEnum = "failure"
)

// WebhookDeliveryTopicEnum is a string enum.
// webhook topic
type WebhookDeliveryTopicEnum string

const (
    WebhookDeliveryTopicEnum_ALARMS          WebhookDeliveryTopicEnum = "alarms"
    WebhookDeliveryTopicEnum_AUDITS          WebhookDeliveryTopicEnum = "audits"
    WebhookDeliveryTopicEnum_DEVICEUPDOWNS   WebhookDeliveryTopicEnum = "device-updowns"
    WebhookDeliveryTopicEnum_OCCUPANCYALERTS WebhookDeliveryTopicEnum = "occupancy-alerts"
    WebhookDeliveryTopicEnum_PING            WebhookDeliveryTopicEnum = "ping"
)

// AclTagTypeEnum is a string enum.
type AclTagTypeEnum string

const (
    AclTagTypeEnum_MAC         AclTagTypeEnum = "mac"
    AclTagTypeEnum_SUBNET      AclTagTypeEnum = "subnet"
    AclTagTypeEnum_NETWORK     AclTagTypeEnum = "network"
    AclTagTypeEnum_RADIUSGROUP AclTagTypeEnum = "radius_group"
    AclTagTypeEnum_ANY         AclTagTypeEnum = "any"
    AclTagTypeEnum_RESOURCE    AclTagTypeEnum = "resource"
    AclTagTypeEnum_DYNAMICGBP  AclTagTypeEnum = "dynamic_gbp"
    AclTagTypeEnum_STATICGBP   AclTagTypeEnum = "static_gbp"
)

// AdminComplianceStatusEnum is a string enum.
// trade compliance status
type AdminComplianceStatusEnum string

const (
    AdminComplianceStatusEnum_RESTRICTED AdminComplianceStatusEnum = "restricted"
    AdminComplianceStatusEnum_BLOCKED    AdminComplianceStatusEnum = "blocked"
)

// AllowDenyEnum is a string enum.
type AllowDenyEnum string

const (
    AllowDenyEnum_ALLOW AllowDenyEnum = "allow"
    AllowDenyEnum_DENY  AllowDenyEnum = "deny"
)

// ApEslTypeEnum is a string enum.
// note: ble_config will be ingored if esl_config is enabled and with native mode.
type ApEslTypeEnum string

const (
    ApEslTypeEnum_IMAGOTAG ApEslTypeEnum = "imagotag"
    ApEslTypeEnum_HANSHOW  ApEslTypeEnum = "hanshow"
    ApEslTypeEnum_SOLUM    ApEslTypeEnum = "solum"
    ApEslTypeEnum_NATIVE   ApEslTypeEnum = "native"
)

// ApIotInputPullupEnum is a string enum.
// the type of pull-up the pin uses (internal, external, none), default none
type ApIotInputPullupEnum string

const (
    ApIotInputPullupEnum_INTERNAL ApIotInputPullupEnum = "internal"
    ApIotInputPullupEnum_EXTERNAL ApIotInputPullupEnum = "external"
    ApIotInputPullupEnum_NONE     ApIotInputPullupEnum = "none"
)

// ApIotOutputPullupEnum is a string enum.
// the type of pull-up the pin uses (internal, external, none), default none
type ApIotOutputPullupEnum string

const (
    ApIotOutputPullupEnum_INTERNAL ApIotOutputPullupEnum = "internal"
    ApIotOutputPullupEnum_EXTERNAL ApIotOutputPullupEnum = "external"
    ApIotOutputPullupEnum_NON      ApIotOutputPullupEnum = "non"
)

// ApMeshRoleEnum is a string enum.
type ApMeshRoleEnum string

const (
    ApMeshRoleEnum_BASE   ApMeshRoleEnum = "base"
    ApMeshRoleEnum_REMOTE ApMeshRoleEnum = "remote"
)

// ApPortConfigForwardingEnum is a string enum.
type ApPortConfigForwardingEnum string

const (
    ApPortConfigForwardingEnum_ALL        ApPortConfigForwardingEnum = "all"
    ApPortConfigForwardingEnum_LIMITED    ApPortConfigForwardingEnum = "limited"
    ApPortConfigForwardingEnum_WXTUNNEL   ApPortConfigForwardingEnum = "wxtunnel"
    ApPortConfigForwardingEnum_MXTUNNEL   ApPortConfigForwardingEnum = "mxtunnel"
    ApPortConfigForwardingEnum_SITEMXEDGE ApPortConfigForwardingEnum = "site_mxedge"
)

// ApPortConfigMacAuthProtocolEnum is a string enum.
// if `enable_mac_auth`==`true`, allows user to select an authentication protocol
type ApPortConfigMacAuthProtocolEnum string

const (
    ApPortConfigMacAuthProtocolEnum_PAP     ApPortConfigMacAuthProtocolEnum = "pap"
    ApPortConfigMacAuthProtocolEnum_EAPPEAP ApPortConfigMacAuthProtocolEnum = "eap-peap"
    ApPortConfigMacAuthProtocolEnum_EAPMD5  ApPortConfigMacAuthProtocolEnum = "eap-md5"
)

// ApPortConfigPortAuthEnum is a string enum.
// When doing port auth
type ApPortConfigPortAuthEnum string

const (
    ApPortConfigPortAuthEnum_NONE  ApPortConfigPortAuthEnum = "none"
    ApPortConfigPortAuthEnum_DOT1X ApPortConfigPortAuthEnum = "dot1x"
)

// ApRadioAntennaModeEnum is a string enum.
type ApRadioAntennaModeEnum string

const (
    ApRadioAntennaModeEnum_ENUMDEFAULT ApRadioAntennaModeEnum = "default"
    ApRadioAntennaModeEnum_ENUM1X1     ApRadioAntennaModeEnum = "1x1"
    ApRadioAntennaModeEnum_ENUM2X2     ApRadioAntennaModeEnum = "2x2"
    ApRadioAntennaModeEnum_ENUM3X3     ApRadioAntennaModeEnum = "3x3"
    ApRadioAntennaModeEnum_ENUM4X4     ApRadioAntennaModeEnum = "4x4"
)

// ApUsbTypeEnum is a string enum.
// usb config type
type ApUsbTypeEnum string

const (
    ApUsbTypeEnum_IMAGOTAG ApUsbTypeEnum = "imagotag"
    ApUsbTypeEnum_SOLUM    ApUsbTypeEnum = "solum"
    ApUsbTypeEnum_HANSHOW  ApUsbTypeEnum = "hanshow"
)

// BgpConfigTypeEnum is a string enum.
type BgpConfigTypeEnum string

const (
    BgpConfigTypeEnum_INTERNAL BgpConfigTypeEnum = "internal"
    BgpConfigTypeEnum_EXTERNAL BgpConfigTypeEnum = "external"
)

// BgpConfigViaEnum is a string enum.
// network name
type BgpConfigViaEnum string

const (
    BgpConfigViaEnum_LAN BgpConfigViaEnum = "lan"
    BgpConfigViaEnum_VPN BgpConfigViaEnum = "vpn"
    BgpConfigViaEnum_WAN BgpConfigViaEnum = "wan"
)

// BgpStatsStateEnum is a string enum.
type BgpStatsStateEnum string

const (
    BgpStatsStateEnum_IDLE        BgpStatsStateEnum = "idle"
    BgpStatsStateEnum_CONNECT     BgpStatsStateEnum = "connect"
    BgpStatsStateEnum_ACTIVE      BgpStatsStateEnum = "active"
    BgpStatsStateEnum_OPENSENT    BgpStatsStateEnum = "open_sent"
    BgpStatsStateEnum_OPENCONFIG  BgpStatsStateEnum = "open_config"
    BgpStatsStateEnum_ESTABLISHED BgpStatsStateEnum = "established"
)

// BleConfigBeaconRateModeEnum is a string enum.
type BleConfigBeaconRateModeEnum string

const (
    BleConfigBeaconRateModeEnum_ENUMDEFAULT BleConfigBeaconRateModeEnum = "default"
    BleConfigBeaconRateModeEnum_CUSTOM      BleConfigBeaconRateModeEnum = "custom"
)

// BleConfigPowerModeEnum is a string enum.
type BleConfigPowerModeEnum string

const (
    BleConfigPowerModeEnum_ENUMDEFAULT BleConfigPowerModeEnum = "default"
    BleConfigPowerModeEnum_CUSTOM      BleConfigPowerModeEnum = "custom"
)

// CaptureMxedgeFormatEnum is a string enum.
// pcap format
type CaptureMxedgeFormatEnum string

const (
    CaptureMxedgeFormatEnum_STREAM CaptureMxedgeFormatEnum = "stream"
)

// ClaimTypeEnum is a string enum.
// what to claim
type ClaimTypeEnum string

const (
    ClaimTypeEnum_ALL       ClaimTypeEnum = "all"
    ClaimTypeEnum_LICENSE   ClaimTypeEnum = "license"
    ClaimTypeEnum_INVENTORY ClaimTypeEnum = "inventory"
)

// ConfigSwitchLocalAccountsUserRoleEnum is a string enum.
type ConfigSwitchLocalAccountsUserRoleEnum string

const (
    ConfigSwitchLocalAccountsUserRoleEnum_NONE     ConfigSwitchLocalAccountsUserRoleEnum = "none"
    ConfigSwitchLocalAccountsUserRoleEnum_ADMIN    ConfigSwitchLocalAccountsUserRoleEnum = "admin"
    ConfigSwitchLocalAccountsUserRoleEnum_READ     ConfigSwitchLocalAccountsUserRoleEnum = "read"
    ConfigSwitchLocalAccountsUserRoleEnum_HELPDESK ConfigSwitchLocalAccountsUserRoleEnum = "helpdesk"
)

// DayOfWeekEnum is a string enum.
type DayOfWeekEnum string

const (
    DayOfWeekEnum_ANY DayOfWeekEnum = "any"
    DayOfWeekEnum_MON DayOfWeekEnum = "mon"
    DayOfWeekEnum_TUE DayOfWeekEnum = "tue"
    DayOfWeekEnum_WED DayOfWeekEnum = "wed"
    DayOfWeekEnum_THU DayOfWeekEnum = "thu"
    DayOfWeekEnum_FRI DayOfWeekEnum = "fri"
    DayOfWeekEnum_SAT DayOfWeekEnum = "sat"
    DayOfWeekEnum_SUN DayOfWeekEnum = "sun"
)

// DeviceTypeApEnum is a string enum.
// Device Type
type DeviceTypeApEnum string

const (
    DeviceTypeApEnum_AP DeviceTypeApEnum = "ap"
)

// DeviceTypeGatewayEnum is a string enum.
// Device Type
type DeviceTypeGatewayEnum string

const (
    DeviceTypeGatewayEnum_GATEWAY DeviceTypeGatewayEnum = "gateway"
)

// DeviceTypeSwitchEnum is a string enum.
// Device Type
type DeviceTypeSwitchEnum string

const (
    DeviceTypeSwitchEnum_ENUMSWITCH DeviceTypeSwitchEnum = "switch"
)

// DeviceUpgradeRrmMeshUpgradeEnum is a string enum.
// sequential or parallel (default parallel). Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade
type DeviceUpgradeRrmMeshUpgradeEnum string

const (
    DeviceUpgradeRrmMeshUpgradeEnum_SEQUENTIAL DeviceUpgradeRrmMeshUpgradeEnum = "sequential"
    DeviceUpgradeRrmMeshUpgradeEnum_PARALLEL   DeviceUpgradeRrmMeshUpgradeEnum = "parallel"
)

// DeviceUpgradeRrmNodeOrderEnum is a string enum.
// Used in rrm to determine whether to start upgrade from fringe or center AP’s
type DeviceUpgradeRrmNodeOrderEnum string

const (
    DeviceUpgradeRrmNodeOrderEnum_FRINGETOCENTER DeviceUpgradeRrmNodeOrderEnum = "fringe_to_center"
    DeviceUpgradeRrmNodeOrderEnum_CENTERTOFRINGE DeviceUpgradeRrmNodeOrderEnum = "center_to_fringe"
)

// DeviceUpgradeStatusEnum is a string enum.
// status upgrade is in
type DeviceUpgradeStatusEnum string

const (
    DeviceUpgradeStatusEnum_DOWNLOADING DeviceUpgradeStatusEnum = "downloading"
    DeviceUpgradeStatusEnum_COMPLETED   DeviceUpgradeStatusEnum = "completed"
    DeviceUpgradeStatusEnum_CREATED     DeviceUpgradeStatusEnum = "created"
    DeviceUpgradeStatusEnum_DOWNLOADED  DeviceUpgradeStatusEnum = "downloaded"
    DeviceUpgradeStatusEnum_UPGRADING   DeviceUpgradeStatusEnum = "upgrading"
    DeviceUpgradeStatusEnum_CANCELLED   DeviceUpgradeStatusEnum = "cancelled"
    DeviceUpgradeStatusEnum_FAILED      DeviceUpgradeStatusEnum = "failed"
)

// DeviceUpgradeStrategyEnum is a string enum.
// `big_bang` (upgrade all at once), `serial` (one at a time), `canary` or `rrm`
type DeviceUpgradeStrategyEnum string

const (
    DeviceUpgradeStrategyEnum_BIGBANG DeviceUpgradeStrategyEnum = "big_bang"
    DeviceUpgradeStrategyEnum_CANARY  DeviceUpgradeStrategyEnum = "canary"
    DeviceUpgradeStrategyEnum_SERIAL  DeviceUpgradeStrategyEnum = "serial"
    DeviceUpgradeStrategyEnum_RRM     DeviceUpgradeStrategyEnum = "rrm"
)

// DhcpdConfigOptionTypeEnum is a string enum.
type DhcpdConfigOptionTypeEnum string

const (
    DhcpdConfigOptionTypeEnum_ENUMSTRING DhcpdConfigOptionTypeEnum = "string"
    DhcpdConfigOptionTypeEnum_BOOLEAN    DhcpdConfigOptionTypeEnum = "boolean"
    DhcpdConfigOptionTypeEnum_IP         DhcpdConfigOptionTypeEnum = "ip"
    DhcpdConfigOptionTypeEnum_HEX        DhcpdConfigOptionTypeEnum = "hex"
    DhcpdConfigOptionTypeEnum_ENUMINT16  DhcpdConfigOptionTypeEnum = "int16"
    DhcpdConfigOptionTypeEnum_ENUMINT32  DhcpdConfigOptionTypeEnum = "int32"
    DhcpdConfigOptionTypeEnum_ENUMUINT16 DhcpdConfigOptionTypeEnum = "uint16"
    DhcpdConfigOptionTypeEnum_ENUMUINT32 DhcpdConfigOptionTypeEnum = "uint32"
)

// DhcpdConfigTypeEnum is a string enum.
// DHCP Server (local) or DHCP Relay (relay)
type DhcpdConfigTypeEnum string

const (
    DhcpdConfigTypeEnum_LOCAL DhcpdConfigTypeEnum = "local"
    DhcpdConfigTypeEnum_RELAY DhcpdConfigTypeEnum = "relay"
    DhcpdConfigTypeEnum_NONE  DhcpdConfigTypeEnum = "none"
)

// DhcpdConfigVendorOptionTypeEnum is a string enum.
type DhcpdConfigVendorOptionTypeEnum string

const (
    DhcpdConfigVendorOptionTypeEnum_ENUMSTRING DhcpdConfigVendorOptionTypeEnum = "string"
    DhcpdConfigVendorOptionTypeEnum_BOOLEAN    DhcpdConfigVendorOptionTypeEnum = "boolean"
    DhcpdConfigVendorOptionTypeEnum_IP         DhcpdConfigVendorOptionTypeEnum = "ip"
    DhcpdConfigVendorOptionTypeEnum_HEX        DhcpdConfigVendorOptionTypeEnum = "hex"
    DhcpdConfigVendorOptionTypeEnum_ENUMINT16  DhcpdConfigVendorOptionTypeEnum = "int16"
    DhcpdConfigVendorOptionTypeEnum_ENUMINT32  DhcpdConfigVendorOptionTypeEnum = "int32"
    DhcpdConfigVendorOptionTypeEnum_ENUMUINT16 DhcpdConfigVendorOptionTypeEnum = "uint16"
    DhcpdConfigVendorOptionTypeEnum_ENUMUINT32 DhcpdConfigVendorOptionTypeEnum = "uint32"
)

// Dot11BandEnum is a string enum.
type Dot11BandEnum string

const (
    Dot11BandEnum_ENUM24 Dot11BandEnum = "24"
    Dot11BandEnum_ENUM5  Dot11BandEnum = "5"
    Dot11BandEnum_ENUM6  Dot11BandEnum = "6"
)

// Dot11BandwidthEnum is a int enum.
// channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6
type Dot11BandwidthEnum int

const (
    Dot11BandwidthEnum_ENUM20  Dot11BandwidthEnum = 20
    Dot11BandwidthEnum_ENUM40  Dot11BandwidthEnum = 40
    Dot11BandwidthEnum_ENUM80  Dot11BandwidthEnum = 80
    Dot11BandwidthEnum_ENUM160 Dot11BandwidthEnum = 160
)

// Dot11ProtoEnum is a string enum.
type Dot11ProtoEnum string

const (
    Dot11ProtoEnum_A  Dot11ProtoEnum = "a"
    Dot11ProtoEnum_B  Dot11ProtoEnum = "b"
    Dot11ProtoEnum_G  Dot11ProtoEnum = "g"
    Dot11ProtoEnum_N  Dot11ProtoEnum = "n"
    Dot11ProtoEnum_AC Dot11ProtoEnum = "ac"
    Dot11ProtoEnum_AX Dot11ProtoEnum = "ax"
)

// Dot11Bandwidth5Enum is a int enum.
// channel width for the 5GHz band
type Dot11Bandwidth5Enum int

const (
    Dot11Bandwidth5Enum_ENUM20 Dot11Bandwidth5Enum = 20
    Dot11Bandwidth5Enum_ENUM40 Dot11Bandwidth5Enum = 40
    Dot11Bandwidth5Enum_ENUM80 Dot11Bandwidth5Enum = 80
)

// Dot11Bandwidth6Enum is a int enum.
// channel width for the 6GHz band
type Dot11Bandwidth6Enum int

const (
    Dot11Bandwidth6Enum_ENUM20  Dot11Bandwidth6Enum = 20
    Dot11Bandwidth6Enum_ENUM40  Dot11Bandwidth6Enum = 40
    Dot11Bandwidth6Enum_ENUM80  Dot11Bandwidth6Enum = 80
    Dot11Bandwidth6Enum_ENUM160 Dot11Bandwidth6Enum = 160
)

// Dot11Bandwidth24Enum is a int enum.
// channel width for the 2.4GHz band
type Dot11Bandwidth24Enum int

const (
    Dot11Bandwidth24Enum_ENUM20 Dot11Bandwidth24Enum = 20
    Dot11Bandwidth24Enum_ENUM40 Dot11Bandwidth24Enum = 40
)

// DynamicPskSourceEnum is a string enum.
type DynamicPskSourceEnum string

const (
    DynamicPskSourceEnum_RADIUS    DynamicPskSourceEnum = "radius"
    DynamicPskSourceEnum_CLOUDPSKS DynamicPskSourceEnum = "cloud_psks"
)

// EvpnConfigRoleEnum is a string enum.
type EvpnConfigRoleEnum string

const (
    EvpnConfigRoleEnum_CORE         EvpnConfigRoleEnum = "core"
    EvpnConfigRoleEnum_DISTRIBUTION EvpnConfigRoleEnum = "distribution"
    EvpnConfigRoleEnum_ACCESS       EvpnConfigRoleEnum = "access"
)

// EvpnTopologySwitchRoleEnum is a string enum.
// use `role`==`none` to remove a switch from the topology
type EvpnTopologySwitchRoleEnum string

const (
    EvpnTopologySwitchRoleEnum_CORE          EvpnTopologySwitchRoleEnum = "core"
    EvpnTopologySwitchRoleEnum_DISTRIBUTION  EvpnTopologySwitchRoleEnum = "distribution"
    EvpnTopologySwitchRoleEnum_ACCESS        EvpnTopologySwitchRoleEnum = "access"
    EvpnTopologySwitchRoleEnum_COLLAPSEDCORE EvpnTopologySwitchRoleEnum = "collapsed-core"
    EvpnTopologySwitchRoleEnum_NONE          EvpnTopologySwitchRoleEnum = "none"
    EvpnTopologySwitchRoleEnum_ESILAGACCESS  EvpnTopologySwitchRoleEnum = "esilag-access"
)

// GatewayPathStrategyEnum is a string enum.
type GatewayPathStrategyEnum string

const (
    GatewayPathStrategyEnum_ORDERED  GatewayPathStrategyEnum = "ordered"
    GatewayPathStrategyEnum_WEIGHTED GatewayPathStrategyEnum = "weighted"
    GatewayPathStrategyEnum_ECMP     GatewayPathStrategyEnum = "ecmp"
)

// GatewayPathTypeEnum is a string enum.
type GatewayPathTypeEnum string

const (
    GatewayPathTypeEnum_LOCAL  GatewayPathTypeEnum = "local"
    GatewayPathTypeEnum_WAN    GatewayPathTypeEnum = "wan"
    GatewayPathTypeEnum_VPN    GatewayPathTypeEnum = "vpn"
    GatewayPathTypeEnum_TUNNEL GatewayPathTypeEnum = "tunnel"
)

// GatewayPortDslTypeEnum is a string enum.
// if `wan_type`==`lte`
type GatewayPortDslTypeEnum string

const (
    GatewayPortDslTypeEnum_VDSL GatewayPortDslTypeEnum = "vdsl"
    GatewayPortDslTypeEnum_ADSL GatewayPortDslTypeEnum = "adsl"
)

// GatewayPortDuplexEnum is a string enum.
type GatewayPortDuplexEnum string

const (
    GatewayPortDuplexEnum_HALF GatewayPortDuplexEnum = "half"
    GatewayPortDuplexEnum_FULL GatewayPortDuplexEnum = "full"
    GatewayPortDuplexEnum_AUTO GatewayPortDuplexEnum = "auto"
)

// GatewayPortLteAuthEnum is a string enum.
// if `wan_type`==`lte`
type GatewayPortLteAuthEnum string

const (
    GatewayPortLteAuthEnum_NONE GatewayPortLteAuthEnum = "none"
    GatewayPortLteAuthEnum_CHAP GatewayPortLteAuthEnum = "chap"
    GatewayPortLteAuthEnum_PAP  GatewayPortLteAuthEnum = "pap"
)

// GatewayPortUsageEnum is a string enum.
// port usage name
type GatewayPortUsageEnum string

const (
    GatewayPortUsageEnum_LAN       GatewayPortUsageEnum = "lan"
    GatewayPortUsageEnum_WAN       GatewayPortUsageEnum = "wan"
    GatewayPortUsageEnum_HADATA    GatewayPortUsageEnum = "ha_data"
    GatewayPortUsageEnum_HACONTROL GatewayPortUsageEnum = "ha_control"
)

// GatewayPortVpnPathBfdProfileEnum is a string enum.
type GatewayPortVpnPathBfdProfileEnum string

const (
    GatewayPortVpnPathBfdProfileEnum_LTE       GatewayPortVpnPathBfdProfileEnum = "lte"
    GatewayPortVpnPathBfdProfileEnum_BROADBAND GatewayPortVpnPathBfdProfileEnum = "broadband"
)

// GatewayPortVpnPathRoleEnum is a string enum.
type GatewayPortVpnPathRoleEnum string

const (
    GatewayPortVpnPathRoleEnum_SPOKE GatewayPortVpnPathRoleEnum = "spoke"
    GatewayPortVpnPathRoleEnum_HUB   GatewayPortVpnPathRoleEnum = "hub"
)

// GatewayPortWanArpPolicerEnum is a string enum.
// when `wan_type`==`broadband`
type GatewayPortWanArpPolicerEnum string

const (
    GatewayPortWanArpPolicerEnum_RECOMMENDED GatewayPortWanArpPolicerEnum = "recommended"
    GatewayPortWanArpPolicerEnum_ENUMDEFAULT GatewayPortWanArpPolicerEnum = "default"
    GatewayPortWanArpPolicerEnum_MAX         GatewayPortWanArpPolicerEnum = "max"
)

// GatewayPortWanTypeEnum is a string enum.
// if `usage`==`wan`
type GatewayPortWanTypeEnum string

const (
    GatewayPortWanTypeEnum_BROADBAND GatewayPortWanTypeEnum = "broadband"
    GatewayPortWanTypeEnum_DSL       GatewayPortWanTypeEnum = "dsl"
    GatewayPortWanTypeEnum_LTE       GatewayPortWanTypeEnum = "lte"
)

// GatewayTemplateProbeTypeEnum is a string enum.
type GatewayTemplateProbeTypeEnum string

const (
    GatewayTemplateProbeTypeEnum_ICMP GatewayTemplateProbeTypeEnum = "icmp"
    GatewayTemplateProbeTypeEnum_HTTP GatewayTemplateProbeTypeEnum = "http"
)

// GatewayTemplateTunnelIkeDhGroupEnum is a string enum.
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
type GatewayTemplateTunnelIkeDhGroupEnum string

const (
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM1  GatewayTemplateTunnelIkeDhGroupEnum = "1"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM2  GatewayTemplateTunnelIkeDhGroupEnum = "2"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM5  GatewayTemplateTunnelIkeDhGroupEnum = "5"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM14 GatewayTemplateTunnelIkeDhGroupEnum = "14"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM15 GatewayTemplateTunnelIkeDhGroupEnum = "15"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM16 GatewayTemplateTunnelIkeDhGroupEnum = "16"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM19 GatewayTemplateTunnelIkeDhGroupEnum = "19"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM20 GatewayTemplateTunnelIkeDhGroupEnum = "20"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM21 GatewayTemplateTunnelIkeDhGroupEnum = "21"
    GatewayTemplateTunnelIkeDhGroupEnum_ENUM24 GatewayTemplateTunnelIkeDhGroupEnum = "24"
)

// GatewayTemplateTunnelIkeModeEnum is a string enum.
// Only if:
// * `provider`== `custom-ipsec`
type GatewayTemplateTunnelIkeModeEnum string

const (
    GatewayTemplateTunnelIkeModeEnum_MAIN       GatewayTemplateTunnelIkeModeEnum = "main"
    GatewayTemplateTunnelIkeModeEnum_AGGRESSIVE GatewayTemplateTunnelIkeModeEnum = "aggressive"
)

// GatewayTemplateTunnelModeEnum is a string enum.
type GatewayTemplateTunnelModeEnum string

const (
    GatewayTemplateTunnelModeEnum_ACTIVESTANDBY GatewayTemplateTunnelModeEnum = "active-standby"
    GatewayTemplateTunnelModeEnum_ACTIVEACTIVE  GatewayTemplateTunnelModeEnum = "active-active"
)

// GatewayTemplateTunnelProtocolEnum is a string enum.
// Only if:
// * `provider`== `custom-ipsec`
type GatewayTemplateTunnelProtocolEnum string

const (
    GatewayTemplateTunnelProtocolEnum_IPSEC GatewayTemplateTunnelProtocolEnum = "ipsec"
    GatewayTemplateTunnelProtocolEnum_GRE   GatewayTemplateTunnelProtocolEnum = "gre"
)

// GatewayTemplateTunnelVersionEnum is a string enum.
// Only if:
// * `provider`== `custom-gre` 
// * `provider`== `custom-ipsec`
type GatewayTemplateTunnelVersionEnum string

const (
    GatewayTemplateTunnelVersionEnum_ENUM1 GatewayTemplateTunnelVersionEnum = "1"
    GatewayTemplateTunnelVersionEnum_ENUM2 GatewayTemplateTunnelVersionEnum = "2"
)

// GatewayTemplateTypeEnum is a string enum.
type GatewayTemplateTypeEnum string

const (
    GatewayTemplateTypeEnum_STANDALONE GatewayTemplateTypeEnum = "standalone"
    GatewayTemplateTypeEnum_SPOKE      GatewayTemplateTypeEnum = "spoke"
)

// GatewayWanPpoeAuthEnum is a string enum.
// if `type`==`pppoe`
type GatewayWanPpoeAuthEnum string

const (
    GatewayWanPpoeAuthEnum_NONE GatewayWanPpoeAuthEnum = "none"
    GatewayWanPpoeAuthEnum_CHAP GatewayWanPpoeAuthEnum = "chap"
    GatewayWanPpoeAuthEnum_PAP  GatewayWanPpoeAuthEnum = "pap"
)

// GatewayWanTypeEnum is a string enum.
type GatewayWanTypeEnum string

const (
    GatewayWanTypeEnum_DHCP   GatewayWanTypeEnum = "dhcp"
    GatewayWanTypeEnum_STATIC GatewayWanTypeEnum = "static"
    GatewayWanTypeEnum_PPPOE  GatewayWanTypeEnum = "pppoe"
)

// IdpProfileActionEnum is a string enum.
// - alert (default) 
// - drop: siliently dropping packets
// - close: notify client/server to close connection
type IdpProfileActionEnum string

const (
    IdpProfileActionEnum_ALERT IdpProfileActionEnum = "alert"
    IdpProfileActionEnum_DROP  IdpProfileActionEnum = "drop"
    IdpProfileActionEnum_CLOSE IdpProfileActionEnum = "close"
)

// IdpProfileBaseProfileEnum is a string enum.
type IdpProfileBaseProfileEnum string

const (
    IdpProfileBaseProfileEnum_CRITICAL IdpProfileBaseProfileEnum = "critical"
    IdpProfileBaseProfileEnum_STRICT   IdpProfileBaseProfileEnum = "strict"
    IdpProfileBaseProfileEnum_STANDARD IdpProfileBaseProfileEnum = "standard"
)

// IdpProfileMatchingSeverityValueEnum is a string enum.
type IdpProfileMatchingSeverityValueEnum string

const (
    IdpProfileMatchingSeverityValueEnum_CRITICAL IdpProfileMatchingSeverityValueEnum = "critical"
    IdpProfileMatchingSeverityValueEnum_MAJOR    IdpProfileMatchingSeverityValueEnum = "major"
    IdpProfileMatchingSeverityValueEnum_MINOR    IdpProfileMatchingSeverityValueEnum = "minor"
    IdpProfileMatchingSeverityValueEnum_INFO     IdpProfileMatchingSeverityValueEnum = "info"
)

// InventoryUpdateOperationEnum is a string enum.
// * if `op`== `upgrade_to_mist`: Upgrade to mist-managed
// * if `op`== `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist.
// * if `op`== `assign`: Assign inventory to a site
// * if `op`== `unassign`: Unassign inventory from a site
// * if `op`== `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned.
type InventoryUpdateOperationEnum string

const (
    InventoryUpdateOperationEnum_ASSIGN         InventoryUpdateOperationEnum = "assign"
    InventoryUpdateOperationEnum_UNASSIGN       InventoryUpdateOperationEnum = "unassign"
    InventoryUpdateOperationEnum_DELETE         InventoryUpdateOperationEnum = "delete"
    InventoryUpdateOperationEnum_UPGRADETOMIST  InventoryUpdateOperationEnum = "upgrade_to_mist"
    InventoryUpdateOperationEnum_DOWNGRADETOJSI InventoryUpdateOperationEnum = "downgrade_to_jsi"
)

// IpTypeEnum is a string enum.
type IpTypeEnum string

const (
    IpTypeEnum_STATIC IpTypeEnum = "static"
    IpTypeEnum_DHCP   IpTypeEnum = "dhcp"
)

// IpType6Enum is a string enum.
type IpType6Enum string

const (
    IpType6Enum_DISABLED IpType6Enum = "disabled"
    IpType6Enum_STATIC   IpType6Enum = "static"
    IpType6Enum_DHCP     IpType6Enum = "dhcp"
    IpType6Enum_AUTOCONF IpType6Enum = "autoconf"
)

// JunosPortConfigDuplexEnum is a string enum.
type JunosPortConfigDuplexEnum string

const (
    JunosPortConfigDuplexEnum_AUTO JunosPortConfigDuplexEnum = "auto"
    JunosPortConfigDuplexEnum_FULL JunosPortConfigDuplexEnum = "full"
    JunosPortConfigDuplexEnum_HALF JunosPortConfigDuplexEnum = "half"
)

// JunosPortConfigSpeedEnum is a string enum.
type JunosPortConfigSpeedEnum string

const (
    JunosPortConfigSpeedEnum_AUTO     JunosPortConfigSpeedEnum = "auto"
    JunosPortConfigSpeedEnum_ENUM10M  JunosPortConfigSpeedEnum = "10m"
    JunosPortConfigSpeedEnum_ENUM100M JunosPortConfigSpeedEnum = "100m"
    JunosPortConfigSpeedEnum_ENUM1G   JunosPortConfigSpeedEnum = "1g"
    JunosPortConfigSpeedEnum_ENUM25G  JunosPortConfigSpeedEnum = "2.5g"
    JunosPortConfigSpeedEnum_ENUM5G   JunosPortConfigSpeedEnum = "5g"
)

// L2TpStateEnum is a string enum.
type L2TpStateEnum string

const (
    L2TpStateEnum_IDLE                   L2TpStateEnum = "idle"
    L2TpStateEnum_WAITCTRLREPLY          L2TpStateEnum = "wait-ctrl-reply"
    L2TpStateEnum_WAITCTRLCONN           L2TpStateEnum = "wait-ctrl-conn"
    L2TpStateEnum_ESTABLISHED            L2TpStateEnum = "established"
    L2TpStateEnum_ESTABLISHEDWITHSESSION L2TpStateEnum = "established_with_session"
)

// LicenseTypeEnum is a string enum.
type LicenseTypeEnum string

const (
    LicenseTypeEnum_SUBMAN   LicenseTypeEnum = "SUB-MAN"
    LicenseTypeEnum_SUBAST   LicenseTypeEnum = "SUB-AST"
    LicenseTypeEnum_SUBVNA   LicenseTypeEnum = "SUB-VNA"
    LicenseTypeEnum_SUBDATA  LicenseTypeEnum = "SUB-DATA"
    LicenseTypeEnum_SUBENG   LicenseTypeEnum = "SUB-ENG"
    LicenseTypeEnum_SUBPMA   LicenseTypeEnum = "SUB-PMA"
    LicenseTypeEnum_SUBEX12  LicenseTypeEnum = "SUB-EX12"
    LicenseTypeEnum_SUBEX24  LicenseTypeEnum = "SUB-EX24"
    LicenseTypeEnum_SUBEX48  LicenseTypeEnum = "SUB-EX48"
    LicenseTypeEnum_SUBSVNA  LicenseTypeEnum = "SUB-SVNA"
    LicenseTypeEnum_SUBME    LicenseTypeEnum = "SUB-ME"
    LicenseTypeEnum_SUBWAN1  LicenseTypeEnum = "SUB-WAN1"
    LicenseTypeEnum_SUBWAN2  LicenseTypeEnum = "SUB-WAN2"
    LicenseTypeEnum_SUBWVNA1 LicenseTypeEnum = "SUB-WVNA1"
    LicenseTypeEnum_SUBWVNA2 LicenseTypeEnum = "SUB-WVNA2"
    LicenseTypeEnum_SUBSRX1  LicenseTypeEnum = "SUB-SRX1"
    LicenseTypeEnum_SUBSRX2  LicenseTypeEnum = "SUB-SRX2"
)

// MapImportJsonVendorNameEnum is a string enum.
type MapImportJsonVendorNameEnum string

const (
    MapImportJsonVendorNameEnum_EKAHAU MapImportJsonVendorNameEnum = "ekahau"
    MapImportJsonVendorNameEnum_IBWAVE MapImportJsonVendorNameEnum = "ibwave"
)

// MapOrgImportFileJsonVendorNameEnum is a string enum.
type MapOrgImportFileJsonVendorNameEnum string

const (
    MapOrgImportFileJsonVendorNameEnum_EKAHAU MapOrgImportFileJsonVendorNameEnum = "ekahau"
    MapOrgImportFileJsonVendorNameEnum_IBWAVE MapOrgImportFileJsonVendorNameEnum = "ibwave"
)

// MxclusterNacClientVendorEnum is a string enum.
// convention to be followed is : "<vendor>-<variant>"
// <variant> could be an os/platform/model/company
// for ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc.
type MxclusterNacClientVendorEnum string

const (
    MxclusterNacClientVendorEnum_GENERIC      MxclusterNacClientVendorEnum = "generic"
    MxclusterNacClientVendorEnum_JUNIPER      MxclusterNacClientVendorEnum = "juniper"
    MxclusterNacClientVendorEnum_ARUBA        MxclusterNacClientVendorEnum = "aruba"
    MxclusterNacClientVendorEnum_PALOALTO     MxclusterNacClientVendorEnum = "paloalto"
    MxclusterNacClientVendorEnum_CISCOAIRONET MxclusterNacClientVendorEnum = "cisco-aironet"
    MxclusterNacClientVendorEnum_CISCOIOS     MxclusterNacClientVendorEnum = "cisco-ios"
    MxclusterNacClientVendorEnum_CISCOMERAKI  MxclusterNacClientVendorEnum = "cisco-meraki"
)

// MxclusterRadAuthServerKeywrapFormatEnum is a string enum.
// if used for Mist APs
type MxclusterRadAuthServerKeywrapFormatEnum string

const (
    MxclusterRadAuthServerKeywrapFormatEnum_HEX   MxclusterRadAuthServerKeywrapFormatEnum = "hex"
    MxclusterRadAuthServerKeywrapFormatEnum_ASCII MxclusterRadAuthServerKeywrapFormatEnum = "ascii"
)

// MxclusterRadsecServerSelectionEnum is a string enum.
// ordered (default) / unordered. When ordered, Mist Edge will prefer and go back to the first radius server if possible
type MxclusterRadsecServerSelectionEnum string

const (
    MxclusterRadsecServerSelectionEnum_ORDERED   MxclusterRadsecServerSelectionEnum = "ordered"
    MxclusterRadsecServerSelectionEnum_UNORDERED MxclusterRadsecServerSelectionEnum = "unordered"
)

// MxclusterRadsecSourceEnum is a string enum.
// Specify source address to use when connecting to RADIUS servers
type MxclusterRadsecSourceEnum string

const (
    MxclusterRadsecSourceEnum_ANY     MxclusterRadsecSourceEnum = "any"
    MxclusterRadsecSourceEnum_OOB     MxclusterRadsecSourceEnum = "oob"
    MxclusterRadsecSourceEnum_OOB6    MxclusterRadsecSourceEnum = "oob6"
    MxclusterRadsecSourceEnum_TUNNEL  MxclusterRadsecSourceEnum = "tunnel"
    MxclusterRadsecSourceEnum_TUNNEL6 MxclusterRadsecSourceEnum = "tunnel6"
)

// MxclusterTuntermHostsSelectionEnum is a string enum.
// Ordering of tunterm_hosts for mxedge within the same mxcluster. 
// * When `shuffle`, the ordering of tunterm_hosts is randomized by the device’s MAC. 
// * When `shuffle-by-site`, we shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels). 
// * When `ordered`, the order is decided by tunterm_hosts_order
type MxclusterTuntermHostsSelectionEnum string

const (
    MxclusterTuntermHostsSelectionEnum_SHUFFLE       MxclusterTuntermHostsSelectionEnum = "shuffle"
    MxclusterTuntermHostsSelectionEnum_SHUFFLEBYSITE MxclusterTuntermHostsSelectionEnum = "shuffle-by-site"
    MxclusterTuntermHostsSelectionEnum_ORDERED       MxclusterTuntermHostsSelectionEnum = "ordered"
)

// MxedgeMgmtOobIpTypeEnum is a string enum.
type MxedgeMgmtOobIpTypeEnum string

const (
    MxedgeMgmtOobIpTypeEnum_DHCP     MxedgeMgmtOobIpTypeEnum = "dhcp"
    MxedgeMgmtOobIpTypeEnum_STATIC   MxedgeMgmtOobIpTypeEnum = "static"
    MxedgeMgmtOobIpTypeEnum_DISABLED MxedgeMgmtOobIpTypeEnum = "disabled"
)

// MxedgeMgmtOobIpType6Enum is a string enum.
type MxedgeMgmtOobIpType6Enum string

const (
    MxedgeMgmtOobIpType6Enum_AUTOCONF MxedgeMgmtOobIpType6Enum = "autoconf"
    MxedgeMgmtOobIpType6Enum_DHCP     MxedgeMgmtOobIpType6Enum = "dhcp"
    MxedgeMgmtOobIpType6Enum_STATIC   MxedgeMgmtOobIpType6Enum = "static"
    MxedgeMgmtOobIpType6Enum_DISABLED MxedgeMgmtOobIpType6Enum = "disabled"
)

// MxedgeTuntermDhcpdConfigTypeEnum is a string enum.
type MxedgeTuntermDhcpdConfigTypeEnum string

const (
    MxedgeTuntermDhcpdConfigTypeEnum_RELAY MxedgeTuntermDhcpdConfigTypeEnum = "relay"
)

// MxedgeTuntermDhcpdTypeEnum is a string enum.
type MxedgeTuntermDhcpdTypeEnum string

const (
    MxedgeTuntermDhcpdTypeEnum_RELAY MxedgeTuntermDhcpdTypeEnum = "relay"
)

// MxedgeUpgradeChannelEnum is a string enum.
// upgrade channel to follow, stable (default) / beta / alpha
type MxedgeUpgradeChannelEnum string

const (
    MxedgeUpgradeChannelEnum_STABLE MxedgeUpgradeChannelEnum = "stable"
    MxedgeUpgradeChannelEnum_BETA   MxedgeUpgradeChannelEnum = "beta"
    MxedgeUpgradeChannelEnum_ALPHA  MxedgeUpgradeChannelEnum = "alpha"
)

// MxedgeUpgradeStrategyEnum is a string enum.
// * `big_bang`: upgrade all at once
// * `serial`: one at a time
type MxedgeUpgradeStrategyEnum string

const (
    MxedgeUpgradeStrategyEnum_BIGBANG MxedgeUpgradeStrategyEnum = "big_bang"
    MxedgeUpgradeStrategyEnum_SERIAL  MxedgeUpgradeStrategyEnum = "serial"
)

// MxtunnelProtocolEnum is a string enum.
type MxtunnelProtocolEnum string

const (
    MxtunnelProtocolEnum_UDP MxtunnelProtocolEnum = "udp"
    MxtunnelProtocolEnum_IP  MxtunnelProtocolEnum = "ip"
)

// MxtunnelStatsStateEnum is a string enum.
type MxtunnelStatsStateEnum string

const (
    MxtunnelStatsStateEnum_IDLE                   MxtunnelStatsStateEnum = "idle"
    MxtunnelStatsStateEnum_WAITCTRLREPLY          MxtunnelStatsStateEnum = "wait-ctrl-reply"
    MxtunnelStatsStateEnum_WAITCTRLCONN           MxtunnelStatsStateEnum = "wait-ctrl-conn"
    MxtunnelStatsStateEnum_ESTABLISHED            MxtunnelStatsStateEnum = "established"
    MxtunnelStatsStateEnum_ESTABLISHEDWITHSESSION MxtunnelStatsStateEnum = "established_with_session"
)

// NacPortalAccessTypeEnum is a string enum.
type NacPortalAccessTypeEnum string

const (
    NacPortalAccessTypeEnum_WIRELESS          NacPortalAccessTypeEnum = "wireless"
    NacPortalAccessTypeEnum_ENUMWIRELESSWIRED NacPortalAccessTypeEnum = "wireless+wired"
)

// NacRuleActionEnum is a string enum.
type NacRuleActionEnum string

const (
    NacRuleActionEnum_ALLOW NacRuleActionEnum = "allow"
    NacRuleActionEnum_BLOCK NacRuleActionEnum = "block"
)

// NacRuleMatchingAuthTypeEnum is a string enum.
type NacRuleMatchingAuthTypeEnum string

const (
    NacRuleMatchingAuthTypeEnum_CERT       NacRuleMatchingAuthTypeEnum = "cert"
    NacRuleMatchingAuthTypeEnum_IDP        NacRuleMatchingAuthTypeEnum = "idp"
    NacRuleMatchingAuthTypeEnum_MAB        NacRuleMatchingAuthTypeEnum = "mab"
    NacRuleMatchingAuthTypeEnum_PSK        NacRuleMatchingAuthTypeEnum = "psk"
    NacRuleMatchingAuthTypeEnum_DEVICEAUTH NacRuleMatchingAuthTypeEnum = "device-auth"
    NacRuleMatchingAuthTypeEnum_EAPTLS     NacRuleMatchingAuthTypeEnum = "eap-tls"
    NacRuleMatchingAuthTypeEnum_EAPTTLS    NacRuleMatchingAuthTypeEnum = "eap-ttls"
    NacRuleMatchingAuthTypeEnum_EAPTEAP    NacRuleMatchingAuthTypeEnum = "eap-teap"
)

// NacRuleMatchingPortTypeEnum is a string enum.
type NacRuleMatchingPortTypeEnum string

const (
    NacRuleMatchingPortTypeEnum_WIRELESS NacRuleMatchingPortTypeEnum = "wireless"
    NacRuleMatchingPortTypeEnum_WIRED    NacRuleMatchingPortTypeEnum = "wired"
)

// NacTagMatchEnum is a string enum.
// if `type`==`match`
type NacTagMatchEnum string

const (
    NacTagMatchEnum_CERTCN       NacTagMatchEnum = "cert_cn"
    NacTagMatchEnum_CERTISSUER   NacTagMatchEnum = "cert_issuer"
    NacTagMatchEnum_CERTSAN      NacTagMatchEnum = "cert_san"
    NacTagMatchEnum_CERTSERIAL   NacTagMatchEnum = "cert_serial"
    NacTagMatchEnum_CERTSUB      NacTagMatchEnum = "cert_sub"
    NacTagMatchEnum_CLIENTMAC    NacTagMatchEnum = "client_mac"
    NacTagMatchEnum_IDPROLE      NacTagMatchEnum = "idp_role"
    NacTagMatchEnum_MDMSTATUS    NacTagMatchEnum = "mdm_status"
    NacTagMatchEnum_RADIUSGROUP  NacTagMatchEnum = "radius_group"
    NacTagMatchEnum_REALM        NacTagMatchEnum = "realm"
    NacTagMatchEnum_SSID         NacTagMatchEnum = "ssid"
    NacTagMatchEnum_USERNAME     NacTagMatchEnum = "user_name"
    NacTagMatchEnum_USERMACLABEL NacTagMatchEnum = "usermac_label"
)

// NacTagTypeEnum is a string enum.
type NacTagTypeEnum string

const (
    NacTagTypeEnum_EGRESSVLANNAMES   NacTagTypeEnum = "egress_vlan_names"
    NacTagTypeEnum_MATCH             NacTagTypeEnum = "match"
    NacTagTypeEnum_VLAN              NacTagTypeEnum = "vlan"
    NacTagTypeEnum_GBPTAG            NacTagTypeEnum = "gbp_tag"
    NacTagTypeEnum_RADIUSATTRS       NacTagTypeEnum = "radius_attrs"
    NacTagTypeEnum_RADIUSGROUP       NacTagTypeEnum = "radius_group"
    NacTagTypeEnum_RADIUSVENDORATTRS NacTagTypeEnum = "radius_vendor_attrs"
    NacTagTypeEnum_SESSIONTIMEOUT    NacTagTypeEnum = "session_timeout"
)

// OrgAutoRulesMatchDeviceTypeEnum is a string enum.
// optional/additional filter
type OrgAutoRulesMatchDeviceTypeEnum string

const (
    OrgAutoRulesMatchDeviceTypeEnum_AP         OrgAutoRulesMatchDeviceTypeEnum = "ap"
    OrgAutoRulesMatchDeviceTypeEnum_ENUMSWITCH OrgAutoRulesMatchDeviceTypeEnum = "switch"
    OrgAutoRulesMatchDeviceTypeEnum_GATEWAY    OrgAutoRulesMatchDeviceTypeEnum = "gateway"
    OrgAutoRulesMatchDeviceTypeEnum_OTHER      OrgAutoRulesMatchDeviceTypeEnum = "other"
)

// OrgAutoRulesSrcEnum is a string enum.
type OrgAutoRulesSrcEnum string

const (
    OrgAutoRulesSrcEnum_NAME           OrgAutoRulesSrcEnum = "name"
    OrgAutoRulesSrcEnum_SUBNET         OrgAutoRulesSrcEnum = "subnet"
    OrgAutoRulesSrcEnum_LLDPSYSTEMNAME OrgAutoRulesSrcEnum = "lldp_system_name"
    OrgAutoRulesSrcEnum_DNSSUFFIX      OrgAutoRulesSrcEnum = "dns_suffix"
    OrgAutoRulesSrcEnum_MODEL          OrgAutoRulesSrcEnum = "model"
    OrgAutoRulesSrcEnum_LLDPPORTDESC   OrgAutoRulesSrcEnum = "lldp_port_desc"
)

// OrgLicenseActionOperationEnum is a string enum.
// to move a license, use the `amend` operation
type OrgLicenseActionOperationEnum string

const (
    OrgLicenseActionOperationEnum_AMEND    OrgLicenseActionOperationEnum = "amend"
    OrgLicenseActionOperationEnum_UNAMEND  OrgLicenseActionOperationEnum = "unamend"
    OrgLicenseActionOperationEnum_DELETE   OrgLicenseActionOperationEnum = "delete"
    OrgLicenseActionOperationEnum_ANNOTATE OrgLicenseActionOperationEnum = "annotate"
)

// OrgSettingMistNacIpVersionEnum is a string enum.
// by default NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4
type OrgSettingMistNacIpVersionEnum string

const (
    OrgSettingMistNacIpVersionEnum_V4 OrgSettingMistNacIpVersionEnum = "v4"
    OrgSettingMistNacIpVersionEnum_V6 OrgSettingMistNacIpVersionEnum = "v6"
)

// OtherDeviceUpdateOperationEnum is a string enum.
// The operation being performed
type OtherDeviceUpdateOperationEnum string

const (
    OtherDeviceUpdateOperationEnum_ASSIGN   OtherDeviceUpdateOperationEnum = "assign"
    OtherDeviceUpdateOperationEnum_UNASSIGN OtherDeviceUpdateOperationEnum = "unassign"
)

// PcapTypeEnum is a string enum.
type PcapTypeEnum string

const (
    PcapTypeEnum_NEWASSOC          PcapTypeEnum = "new_assoc"
    PcapTypeEnum_CLIENT            PcapTypeEnum = "client"
    PcapTypeEnum_WIRED             PcapTypeEnum = "wired"
    PcapTypeEnum_WIRELESS          PcapTypeEnum = "wireless"
    PcapTypeEnum_RADIOTAP          PcapTypeEnum = "radiotap"
    PcapTypeEnum_ENUMRADIOTAPWIRED PcapTypeEnum = "radiotap,wired"
    PcapTypeEnum_GATEWAY           PcapTypeEnum = "gateway"
)

// PrivilegeMspRoleEnum is a string enum.
// access permissions
type PrivilegeMspRoleEnum string

const (
    PrivilegeMspRoleEnum_ADMIN     PrivilegeMspRoleEnum = "admin"
    PrivilegeMspRoleEnum_WRITE     PrivilegeMspRoleEnum = "write"
    PrivilegeMspRoleEnum_READ      PrivilegeMspRoleEnum = "read"
    PrivilegeMspRoleEnum_HELPDESK  PrivilegeMspRoleEnum = "helpdesk"
    PrivilegeMspRoleEnum_INSTALLER PrivilegeMspRoleEnum = "installer"
)

// PrivilegeMspScopeEnum is a string enum.
type PrivilegeMspScopeEnum string

const (
    PrivilegeMspScopeEnum_ORG      PrivilegeMspScopeEnum = "org"
    PrivilegeMspScopeEnum_MSP      PrivilegeMspScopeEnum = "msp"
    PrivilegeMspScopeEnum_ORGGROUP PrivilegeMspScopeEnum = "orggroup"
)

// PrivilegeMspViewEnum is a string enum.
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// | UI View | Description |
// | --- | --- |
// | `reporting` | full access to all analytics tools |
// | `marketing` | can view analytics and location maps |
// | `location` | can view and manage location maps |
// | `security` | can view and manage WLAN, rogues and authentication |
// | `switch_admin` | can view and manage Switch ports |
// | `mxedge_admin` | can view and manage Mist edges and Mist tunnels |
// | `lobby_admin` | full access to Org and Site Pre-shared keys |
type PrivilegeMspViewEnum string

const (
    PrivilegeMspViewEnum_REPORTING   PrivilegeMspViewEnum = "reporting"
    PrivilegeMspViewEnum_MARKETING   PrivilegeMspViewEnum = "marketing"
    PrivilegeMspViewEnum_LOCATION    PrivilegeMspViewEnum = "location"
    PrivilegeMspViewEnum_SECURITY    PrivilegeMspViewEnum = "security"
    PrivilegeMspViewEnum_SWITCHADMIN PrivilegeMspViewEnum = "switch_admin"
    PrivilegeMspViewEnum_MXEDGEADMIN PrivilegeMspViewEnum = "mxedge_admin"
    PrivilegeMspViewEnum_LOBBYADMIN  PrivilegeMspViewEnum = "lobby_admin"
)

// PrivilegeOrgRoleEnum is a string enum.
// access permissions
type PrivilegeOrgRoleEnum string

const (
    PrivilegeOrgRoleEnum_ADMIN     PrivilegeOrgRoleEnum = "admin"
    PrivilegeOrgRoleEnum_WRITE     PrivilegeOrgRoleEnum = "write"
    PrivilegeOrgRoleEnum_READ      PrivilegeOrgRoleEnum = "read"
    PrivilegeOrgRoleEnum_HELPDESK  PrivilegeOrgRoleEnum = "helpdesk"
    PrivilegeOrgRoleEnum_INSTALLER PrivilegeOrgRoleEnum = "installer"
)

// PrivilegeOrgScopeEnum is a string enum.
type PrivilegeOrgScopeEnum string

const (
    PrivilegeOrgScopeEnum_ORG       PrivilegeOrgScopeEnum = "org"
    PrivilegeOrgScopeEnum_SITE      PrivilegeOrgScopeEnum = "site"
    PrivilegeOrgScopeEnum_SITEGROUP PrivilegeOrgScopeEnum = "sitegroup"
)

// PrivilegeOrgViewsEnum is a string enum.
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// | UI View | Description |
// | --- | --- |
// | `reporting` | full access to all analytics tools |
// | `marketing` | can view analytics and location maps |
// | `location` | can view and manage location maps |
// | `security` | can view and manage WLAN, rogues and authentication |
// | `switch_admin` | can view and manage Switch ports |
// | `mxedge_admin` | can view and manage Mist edges and Mist tunnels |
// | `lobby_admin` | full access to Org and Site Pre-shared keys |
type PrivilegeOrgViewsEnum string

const (
    PrivilegeOrgViewsEnum_REPORTING   PrivilegeOrgViewsEnum = "reporting"
    PrivilegeOrgViewsEnum_MARKETING   PrivilegeOrgViewsEnum = "marketing"
    PrivilegeOrgViewsEnum_LOCATION    PrivilegeOrgViewsEnum = "location"
    PrivilegeOrgViewsEnum_SECURITY    PrivilegeOrgViewsEnum = "security"
    PrivilegeOrgViewsEnum_SWITCHADMIN PrivilegeOrgViewsEnum = "switch_admin"
    PrivilegeOrgViewsEnum_MXEDGEADMIN PrivilegeOrgViewsEnum = "mxedge_admin"
    PrivilegeOrgViewsEnum_LOBBYADMIN  PrivilegeOrgViewsEnum = "lobby_admin"
)

// PrivilegeSelfRoleEnum is a string enum.
// access permissions
type PrivilegeSelfRoleEnum string

const (
    PrivilegeSelfRoleEnum_ADMIN     PrivilegeSelfRoleEnum = "admin"
    PrivilegeSelfRoleEnum_WRITE     PrivilegeSelfRoleEnum = "write"
    PrivilegeSelfRoleEnum_READ      PrivilegeSelfRoleEnum = "read"
    PrivilegeSelfRoleEnum_HELPDESK  PrivilegeSelfRoleEnum = "helpdesk"
    PrivilegeSelfRoleEnum_INSTALLER PrivilegeSelfRoleEnum = "installer"
)

// PrivilegeSelfScopeEnum is a string enum.
type PrivilegeSelfScopeEnum string

const (
    PrivilegeSelfScopeEnum_MSP       PrivilegeSelfScopeEnum = "msp"
    PrivilegeSelfScopeEnum_ORG       PrivilegeSelfScopeEnum = "org"
    PrivilegeSelfScopeEnum_ORGGROUP  PrivilegeSelfScopeEnum = "orggroup"
    PrivilegeSelfScopeEnum_SITE      PrivilegeSelfScopeEnum = "site"
    PrivilegeSelfScopeEnum_SITEGROUP PrivilegeSelfScopeEnum = "sitegroup"
)

// PrivilegeSelfViewsEnum is a string enum.
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.
// You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.
// Below are the list of supported UI views. Note that this is UI only feature
// | UI View | Description |
// | --- | --- |
// | `reporting` | full access to all analytics tools |
// | `marketing` | can view analytics and location maps |
// | `location` | can view and manage location maps |
// | `security` | can view and manage WLAN, rogues and authentication |
// | `switch_admin` | can view and manage Switch ports |
// | `mxedge_admin` | can view and manage Mist edges and Mist tunnels |
// | `lobby_admin` | full access to Org and Site Pre-shared keys |
type PrivilegeSelfViewsEnum string

const (
    PrivilegeSelfViewsEnum_REPORTING   PrivilegeSelfViewsEnum = "reporting"
    PrivilegeSelfViewsEnum_MARKETING   PrivilegeSelfViewsEnum = "marketing"
    PrivilegeSelfViewsEnum_LOCATION    PrivilegeSelfViewsEnum = "location"
    PrivilegeSelfViewsEnum_SECURITY    PrivilegeSelfViewsEnum = "security"
    PrivilegeSelfViewsEnum_SWITCHADMIN PrivilegeSelfViewsEnum = "switch_admin"
    PrivilegeSelfViewsEnum_MXEDGEADMIN PrivilegeSelfViewsEnum = "mxedge_admin"
    PrivilegeSelfViewsEnum_LOBBYADMIN  PrivilegeSelfViewsEnum = "lobby_admin"
)

// ProtectReCustomProtocolEnum is a string enum.
type ProtectReCustomProtocolEnum string

const (
    ProtectReCustomProtocolEnum_TCP  ProtectReCustomProtocolEnum = "tcp"
    ProtectReCustomProtocolEnum_UDP  ProtectReCustomProtocolEnum = "udp"
    ProtectReCustomProtocolEnum_ICMP ProtectReCustomProtocolEnum = "icmp"
    ProtectReCustomProtocolEnum_ANY  ProtectReCustomProtocolEnum = "any"
)

// PskPortalAuthEnum is a string enum.
// Note: `sponsor` not yet available
type PskPortalAuthEnum string

const (
    PskPortalAuthEnum_SSO     PskPortalAuthEnum = "sso"
    PskPortalAuthEnum_SPONSOR PskPortalAuthEnum = "sponsor"
)

// PskPortalTemplateAlignmentEnum is a string enum.
// defines alignment on portal
type PskPortalTemplateAlignmentEnum string

const (
    PskPortalTemplateAlignmentEnum_LEFT   PskPortalTemplateAlignmentEnum = "left"
    PskPortalTemplateAlignmentEnum_CENTER PskPortalTemplateAlignmentEnum = "center"
    PskPortalTemplateAlignmentEnum_RIGHT  PskPortalTemplateAlignmentEnum = "right"
)

// PskPortalTypeEnum is a string enum.
// for personal psk portal
type PskPortalTypeEnum string

const (
    PskPortalTypeEnum_BYOD  PskPortalTypeEnum = "byod"
    PskPortalTypeEnum_ADMIN PskPortalTypeEnum = "admin"
)

// PskUsageEnum is a string enum.
type PskUsageEnum string

const (
    PskUsageEnum_MULTI  PskUsageEnum = "multi"
    PskUsageEnum_SINGLE PskUsageEnum = "single"
    PskUsageEnum_MACS   PskUsageEnum = "macs"
)

// RadioBand24UsageEnum is a string enum.
type RadioBand24UsageEnum string

const (
    RadioBand24UsageEnum_ENUM24 RadioBand24UsageEnum = "24"
    RadioBand24UsageEnum_ENUM5  RadioBand24UsageEnum = "5"
    RadioBand24UsageEnum_ENUM6  RadioBand24UsageEnum = "6"
    RadioBand24UsageEnum_AUTO   RadioBand24UsageEnum = "auto"
)

// RadioBandAntennaModeEnum is a string enum.
type RadioBandAntennaModeEnum string

const (
    RadioBandAntennaModeEnum_ENUMDEFAULT RadioBandAntennaModeEnum = "default"
    RadioBandAntennaModeEnum_ENUM1X1     RadioBandAntennaModeEnum = "1x1"
    RadioBandAntennaModeEnum_ENUM2X2     RadioBandAntennaModeEnum = "2x2"
    RadioBandAntennaModeEnum_ENUM3X3     RadioBandAntennaModeEnum = "3x3"
    RadioBandAntennaModeEnum_ENUM4X4     RadioBandAntennaModeEnum = "4x4"
)

// RadioBandPreambleEnum is a string enum.
type RadioBandPreambleEnum string

const (
    RadioBandPreambleEnum_SHORT RadioBandPreambleEnum = "short"
    RadioBandPreambleEnum_LONG  RadioBandPreambleEnum = "long"
    RadioBandPreambleEnum_AUTO  RadioBandPreambleEnum = "auto"
)

// RadiusKeywrapFormatEnum is a string enum.
type RadiusKeywrapFormatEnum string

const (
    RadiusKeywrapFormatEnum_HEX   RadiusKeywrapFormatEnum = "hex"
    RadiusKeywrapFormatEnum_ASCII RadiusKeywrapFormatEnum = "ascii"
)

// RemoteSyslogFacilityEnum is a string enum.
type RemoteSyslogFacilityEnum string

const (
    RemoteSyslogFacilityEnum_ANY                 RemoteSyslogFacilityEnum = "any"
    RemoteSyslogFacilityEnum_AUTHORIZATION       RemoteSyslogFacilityEnum = "authorization"
    RemoteSyslogFacilityEnum_CONFLICTLOG         RemoteSyslogFacilityEnum = "conflict-log"
    RemoteSyslogFacilityEnum_CHANGELOG           RemoteSyslogFacilityEnum = "change-log"
    RemoteSyslogFacilityEnum_CONFIG              RemoteSyslogFacilityEnum = "config"
    RemoteSyslogFacilityEnum_DAEMON              RemoteSyslogFacilityEnum = "daemon"
    RemoteSyslogFacilityEnum_DFC                 RemoteSyslogFacilityEnum = "dfc"
    RemoteSyslogFacilityEnum_KERNEL              RemoteSyslogFacilityEnum = "kernel"
    RemoteSyslogFacilityEnum_INTERACTIVECOMMANDS RemoteSyslogFacilityEnum = "interactive-commands"
    RemoteSyslogFacilityEnum_FTP                 RemoteSyslogFacilityEnum = "ftp"
    RemoteSyslogFacilityEnum_FIREWALL            RemoteSyslogFacilityEnum = "firewall"
    RemoteSyslogFacilityEnum_EXTERNAL            RemoteSyslogFacilityEnum = "external"
    RemoteSyslogFacilityEnum_PFE                 RemoteSyslogFacilityEnum = "pfe"
    RemoteSyslogFacilityEnum_NTP                 RemoteSyslogFacilityEnum = "ntp"
    RemoteSyslogFacilityEnum_SECURITY            RemoteSyslogFacilityEnum = "security"
    RemoteSyslogFacilityEnum_USER                RemoteSyslogFacilityEnum = "user"
)

// RemoteSyslogServerProtocolEnum is a string enum.
type RemoteSyslogServerProtocolEnum string

const (
    RemoteSyslogServerProtocolEnum_UDP RemoteSyslogServerProtocolEnum = "udp"
    RemoteSyslogServerProtocolEnum_TCP RemoteSyslogServerProtocolEnum = "tcp"
)

// RemoteSyslogSeverityEnum is a string enum.
type RemoteSyslogSeverityEnum string

const (
    RemoteSyslogSeverityEnum_ANY       RemoteSyslogSeverityEnum = "any"
    RemoteSyslogSeverityEnum_ALERT     RemoteSyslogSeverityEnum = "alert"
    RemoteSyslogSeverityEnum_EMERGENCY RemoteSyslogSeverityEnum = "emergency"
    RemoteSyslogSeverityEnum_CRITICAL  RemoteSyslogSeverityEnum = "critical"
    RemoteSyslogSeverityEnum_WARNING   RemoteSyslogSeverityEnum = "warning"
    RemoteSyslogSeverityEnum_INFO      RemoteSyslogSeverityEnum = "info"
    RemoteSyslogSeverityEnum_NOTICE    RemoteSyslogSeverityEnum = "notice"
    RemoteSyslogSeverityEnum_ENUMERROR RemoteSyslogSeverityEnum = "error"
)

// RemoteSyslogTimeFormatEnum is a string enum.
type RemoteSyslogTimeFormatEnum string

const (
    RemoteSyslogTimeFormatEnum_MILLISECOND         RemoteSyslogTimeFormatEnum = "millisecond"
    RemoteSyslogTimeFormatEnum_YEAR                RemoteSyslogTimeFormatEnum = "year"
    RemoteSyslogTimeFormatEnum_ENUMYEARMILLISECOND RemoteSyslogTimeFormatEnum = "year millisecond"
)

// ResponseMapImportApActionEnum is a string enum.
type ResponseMapImportApActionEnum string

const (
    ResponseMapImportApActionEnum_PLACED              ResponseMapImportApActionEnum = "placed"
    ResponseMapImportApActionEnum_ASSIGNEDPLACED      ResponseMapImportApActionEnum = "assigned-placed"
    ResponseMapImportApActionEnum_NAMEDPLACED         ResponseMapImportApActionEnum = "named-placed"
    ResponseMapImportApActionEnum_ASSIGNEDNAMEDPLACED ResponseMapImportApActionEnum = "assigned-named-placed"
    ResponseMapImportApActionEnum_IGNORED             ResponseMapImportApActionEnum = "ignored"
)

// ResponseOrgInventoryChangeOpEnum is a string enum.
type ResponseOrgInventoryChangeOpEnum string

const (
    ResponseOrgInventoryChangeOpEnum_ASSIGN         ResponseOrgInventoryChangeOpEnum = "assign"
    ResponseOrgInventoryChangeOpEnum_UNASSIGN       ResponseOrgInventoryChangeOpEnum = "unassign"
    ResponseOrgInventoryChangeOpEnum_DELETE         ResponseOrgInventoryChangeOpEnum = "delete"
    ResponseOrgInventoryChangeOpEnum_UPGRADETOMIST  ResponseOrgInventoryChangeOpEnum = "upgrade_to_mist"
    ResponseOrgInventoryChangeOpEnum_DOWNGRADETOJSI ResponseOrgInventoryChangeOpEnum = "downgrade_to_jsi"
)

// ServiceFailoverPolicyEnum is a string enum.
type ServiceFailoverPolicyEnum string

const (
    ServiceFailoverPolicyEnum_REVERTABLE    ServiceFailoverPolicyEnum = "revertable"
    ServiceFailoverPolicyEnum_NONREVERTABLE ServiceFailoverPolicyEnum = "non_revertable"
    ServiceFailoverPolicyEnum_NONE          ServiceFailoverPolicyEnum = "none"
)

// ServicePolicyEwfRuleProfileEnum is a string enum.
type ServicePolicyEwfRuleProfileEnum string

const (
    ServicePolicyEwfRuleProfileEnum_CRITICAL ServicePolicyEwfRuleProfileEnum = "critical"
    ServicePolicyEwfRuleProfileEnum_STRICT   ServicePolicyEwfRuleProfileEnum = "strict"
    ServicePolicyEwfRuleProfileEnum_STANDARD ServicePolicyEwfRuleProfileEnum = "standard"
)

// ServiceTrafficClassEnum is a string enum.
// when `traffic_type`==`custom`
type ServiceTrafficClassEnum string

const (
    ServiceTrafficClassEnum_BESTEFFORT ServiceTrafficClassEnum = "best_effort"
    ServiceTrafficClassEnum_HIGH       ServiceTrafficClassEnum = "high"
    ServiceTrafficClassEnum_MEDIUM     ServiceTrafficClassEnum = "medium"
    ServiceTrafficClassEnum_LOW        ServiceTrafficClassEnum = "low"
)

// ServiceTypeEnum is a string enum.
type ServiceTypeEnum string

const (
    ServiceTypeEnum_APPS          ServiceTypeEnum = "apps"
    ServiceTypeEnum_APPCATEGORIES ServiceTypeEnum = "app_categories"
    ServiceTypeEnum_CUSTOM        ServiceTypeEnum = "custom"
    ServiceTypeEnum_URLS          ServiceTypeEnum = "urls"
)

// SnmpConfigEngineIdEnum is a string enum.
type SnmpConfigEngineIdEnum string

const (
    SnmpConfigEngineIdEnum_LOCAL               SnmpConfigEngineIdEnum = "local"
    SnmpConfigEngineIdEnum_ENGINEIDSUFFIX      SnmpConfigEngineIdEnum = "engine-id-suffix"
    SnmpConfigEngineIdEnum_USEDEFAULTIPADDRESS SnmpConfigEngineIdEnum = "use-default-ip-address"
    SnmpConfigEngineIdEnum_USEMACADDRESS       SnmpConfigEngineIdEnum = "use_mac-address"
)

// SnmpConfigTrapVerionEnum is a string enum.
type SnmpConfigTrapVerionEnum string

const (
    SnmpConfigTrapVerionEnum_V1  SnmpConfigTrapVerionEnum = "v1"
    SnmpConfigTrapVerionEnum_V2  SnmpConfigTrapVerionEnum = "v2"
    SnmpConfigTrapVerionEnum_ALL SnmpConfigTrapVerionEnum = "all"
)

// SnmpUsmEngineTypeEnum is a string enum.
type SnmpUsmEngineTypeEnum string

const (
    SnmpUsmEngineTypeEnum_REMOTEENGINE SnmpUsmEngineTypeEnum = "remote_engine"
    SnmpUsmEngineTypeEnum_LOCALENGINE  SnmpUsmEngineTypeEnum = "local_engine"
)

// SnmpUsmpUserAuthenticationTypeEnum is a string enum.
// sha224, sha256, sha384, sha512 are supported in 21.1 and newer release
type SnmpUsmpUserAuthenticationTypeEnum string

const (
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONMD5    SnmpUsmpUserAuthenticationTypeEnum = "authentication_md5"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA    SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA224 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha224"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA256 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha256"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA384 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha384"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONSHA512 SnmpUsmpUserAuthenticationTypeEnum = "authentication_sha512"
    SnmpUsmpUserAuthenticationTypeEnum_AUTHENTICATIONNONE   SnmpUsmpUserAuthenticationTypeEnum = "authentication_none"
)

// SnmpUsmpUserEncryptionTypeEnum is a string enum.
type SnmpUsmpUserEncryptionTypeEnum string

const (
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYAES128 SnmpUsmpUserEncryptionTypeEnum = "privacy-aes128"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYDES    SnmpUsmpUserEncryptionTypeEnum = "privacy-des"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACY3DES   SnmpUsmpUserEncryptionTypeEnum = "privacy-3des"
    SnmpUsmpUserEncryptionTypeEnum_PRIVACYNONE   SnmpUsmpUserEncryptionTypeEnum = "privacy-none"
)

// SnmpVacmAccessItemPrefixListItemLevelEnum is a string enum.
type SnmpVacmAccessItemPrefixListItemLevelEnum string

const (
    SnmpVacmAccessItemPrefixListItemLevelEnum_PRIVACY        SnmpVacmAccessItemPrefixListItemLevelEnum = "privacy"
    SnmpVacmAccessItemPrefixListItemLevelEnum_AUTHENTICATION SnmpVacmAccessItemPrefixListItemLevelEnum = "authentication"
    SnmpVacmAccessItemPrefixListItemLevelEnum_NONE           SnmpVacmAccessItemPrefixListItemLevelEnum = "none"
)

// SnmpVacmAccessItemPrefixListItemModelEnum is a string enum.
type SnmpVacmAccessItemPrefixListItemModelEnum string

const (
    SnmpVacmAccessItemPrefixListItemModelEnum_ANY SnmpVacmAccessItemPrefixListItemModelEnum = "any"
    SnmpVacmAccessItemPrefixListItemModelEnum_USM SnmpVacmAccessItemPrefixListItemModelEnum = "usm"
    SnmpVacmAccessItemPrefixListItemModelEnum_V1  SnmpVacmAccessItemPrefixListItemModelEnum = "v1"
    SnmpVacmAccessItemPrefixListItemModelEnum_V2C SnmpVacmAccessItemPrefixListItemModelEnum = "v2c"
)

// SnmpVacmAccessItemTypeEnum is a string enum.
type SnmpVacmAccessItemTypeEnum string

const (
    SnmpVacmAccessItemTypeEnum_CONTEXTPREFIX        SnmpVacmAccessItemTypeEnum = "context_prefix"
    SnmpVacmAccessItemTypeEnum_DEFAULTCONTEXTPREFIX SnmpVacmAccessItemTypeEnum = "default_context_prefix"
)

// SnmpVacmSecurityModelEnum is a string enum.
type SnmpVacmSecurityModelEnum string

const (
    SnmpVacmSecurityModelEnum_USM SnmpVacmSecurityModelEnum = "usm"
    SnmpVacmSecurityModelEnum_V1  SnmpVacmSecurityModelEnum = "v1"
    SnmpVacmSecurityModelEnum_V2C SnmpVacmSecurityModelEnum = "v2c"
)

// Snmpv3ConfigNotifyTypeEnum is a string enum.
type Snmpv3ConfigNotifyTypeEnum string

const (
    Snmpv3ConfigNotifyTypeEnum_TRAP   Snmpv3ConfigNotifyTypeEnum = "trap"
    Snmpv3ConfigNotifyTypeEnum_INFORM Snmpv3ConfigNotifyTypeEnum = "inform"
)

// Snmpv3ConfigTargetParamMessProcessModelEnum is a string enum.
type Snmpv3ConfigTargetParamMessProcessModelEnum string

const (
    Snmpv3ConfigTargetParamMessProcessModelEnum_V1  Snmpv3ConfigTargetParamMessProcessModelEnum = "v1"
    Snmpv3ConfigTargetParamMessProcessModelEnum_V2C Snmpv3ConfigTargetParamMessProcessModelEnum = "v2c"
    Snmpv3ConfigTargetParamMessProcessModelEnum_V3  Snmpv3ConfigTargetParamMessProcessModelEnum = "v3"
)

// Snmpv3ConfigTargetParamSecurityLevelEnum is a string enum.
type Snmpv3ConfigTargetParamSecurityLevelEnum string

const (
    Snmpv3ConfigTargetParamSecurityLevelEnum_AUTHENTICATION Snmpv3ConfigTargetParamSecurityLevelEnum = "authentication"
    Snmpv3ConfigTargetParamSecurityLevelEnum_NONE           Snmpv3ConfigTargetParamSecurityLevelEnum = "none"
    Snmpv3ConfigTargetParamSecurityLevelEnum_PRIVACY        Snmpv3ConfigTargetParamSecurityLevelEnum = "privacy"
)

// Snmpv3ConfigTargetParamSecurityModelEnum is a string enum.
type Snmpv3ConfigTargetParamSecurityModelEnum string

const (
    Snmpv3ConfigTargetParamSecurityModelEnum_USM Snmpv3ConfigTargetParamSecurityModelEnum = "usm"
    Snmpv3ConfigTargetParamSecurityModelEnum_V1  Snmpv3ConfigTargetParamSecurityModelEnum = "v1"
    Snmpv3ConfigTargetParamSecurityModelEnum_V2C Snmpv3ConfigTargetParamSecurityModelEnum = "v2c"
)

// SsoIdpTypeEnum is a string enum.
type SsoIdpTypeEnum string

const (
    SsoIdpTypeEnum_SAML        SsoIdpTypeEnum = "saml"
    SsoIdpTypeEnum_LDAP        SsoIdpTypeEnum = "ldap"
    SsoIdpTypeEnum_OAUTH       SsoIdpTypeEnum = "oauth"
    SsoIdpTypeEnum_MXEDGEPROXY SsoIdpTypeEnum = "mxedge_proxy"
)

// SsoLdapTypeEnum is a string enum.
// if `idp_type`==`ldap`
type SsoLdapTypeEnum string

const (
    SsoLdapTypeEnum_AZURE  SsoLdapTypeEnum = "azure"
    SsoLdapTypeEnum_OKTA   SsoLdapTypeEnum = "okta"
    SsoLdapTypeEnum_GOOGLE SsoLdapTypeEnum = "google"
    SsoLdapTypeEnum_CUSTOM SsoLdapTypeEnum = "custom"
)

// SsoNameidFormatEnum is a string enum.
// if `idp_type`==`saml`
type SsoNameidFormatEnum string

const (
    SsoNameidFormatEnum_EMAIL       SsoNameidFormatEnum = "email"
    SsoNameidFormatEnum_UNSPECIFIED SsoNameidFormatEnum = "unspecified"
)

// SsoOauthTypeEnum is a string enum.
type SsoOauthTypeEnum string

const (
    SsoOauthTypeEnum_AZURE    SsoOauthTypeEnum = "azure"
    SsoOauthTypeEnum_AZUREGOV SsoOauthTypeEnum = "azure-gov"
    SsoOauthTypeEnum_OKTA     SsoOauthTypeEnum = "okta"
)

// SsrUpgradeChannelEnum is a string enum.
// upgrade channel to follow
type SsrUpgradeChannelEnum string

const (
    SsrUpgradeChannelEnum_STABLE SsrUpgradeChannelEnum = "stable"
    SsrUpgradeChannelEnum_BETA   SsrUpgradeChannelEnum = "beta"
    SsrUpgradeChannelEnum_ALPHA  SsrUpgradeChannelEnum = "alpha"
)

// SsrUpgradeStrategyEnum is a string enum.
// * `big_bang`: upgrade all at once
// * `serial`: one at a time
type SsrUpgradeStrategyEnum string

const (
    SsrUpgradeStrategyEnum_BIGBANG SsrUpgradeStrategyEnum = "big_bang"
    SsrUpgradeStrategyEnum_SERIAL  SsrUpgradeStrategyEnum = "serial"
)

// SwitchPortStatsAuthStateEnum is a string enum.
// if `up`==`true` and has Authenticator role
type SwitchPortStatsAuthStateEnum string

const (
    SwitchPortStatsAuthStateEnum_INIT           SwitchPortStatsAuthStateEnum = "init"
    SwitchPortStatsAuthStateEnum_AUTHENTICATED  SwitchPortStatsAuthStateEnum = "authenticated"
    SwitchPortStatsAuthStateEnum_AUTHENTICATING SwitchPortStatsAuthStateEnum = "authenticating"
    SwitchPortStatsAuthStateEnum_HELD           SwitchPortStatsAuthStateEnum = "held"
)

// SwitchPortStatsPoeModeEnum is a string enum.
type SwitchPortStatsPoeModeEnum string

const (
    SwitchPortStatsPoeModeEnum_ENUM8023AF SwitchPortStatsPoeModeEnum = "802.3af"
    SwitchPortStatsPoeModeEnum_ENUM8023AT SwitchPortStatsPoeModeEnum = "802.3at"
    SwitchPortStatsPoeModeEnum_ENUM8023BT SwitchPortStatsPoeModeEnum = "802.3bt"
)

// SwitchPortStatsPortUsageEnum is a string enum.
// gateway port usage
type SwitchPortStatsPortUsageEnum string

const (
    SwitchPortStatsPortUsageEnum_LAN SwitchPortStatsPortUsageEnum = "lan"
)

// SwitchPortStatsStpRoleEnum is a string enum.
// if `up`==`true`
type SwitchPortStatsStpRoleEnum string

const (
    SwitchPortStatsStpRoleEnum_DESIGNATED    SwitchPortStatsStpRoleEnum = "designated"
    SwitchPortStatsStpRoleEnum_BACKUP        SwitchPortStatsStpRoleEnum = "backup"
    SwitchPortStatsStpRoleEnum_ALTERNATE     SwitchPortStatsStpRoleEnum = "alternate"
    SwitchPortStatsStpRoleEnum_ROOT          SwitchPortStatsStpRoleEnum = "root"
    SwitchPortStatsStpRoleEnum_ROOTPREVENTED SwitchPortStatsStpRoleEnum = "root-prevented"
)

// SwitchPortStatsStpStateEnum is a string enum.
// if `up`==`true`
type SwitchPortStatsStpStateEnum string

const (
    SwitchPortStatsStpStateEnum_FORWARDING SwitchPortStatsStpStateEnum = "forwarding"
    SwitchPortStatsStpStateEnum_BLOCKING   SwitchPortStatsStpStateEnum = "blocking"
    SwitchPortStatsStpStateEnum_LEARNING   SwitchPortStatsStpStateEnum = "learning"
    SwitchPortStatsStpStateEnum_LISTENING  SwitchPortStatsStpStateEnum = "listening"
    SwitchPortStatsStpStateEnum_DISABLED   SwitchPortStatsStpStateEnum = "disabled"
)

// SwitchPortStatsTypeEnum is a string enum.
// device type
type SwitchPortStatsTypeEnum string

const (
    SwitchPortStatsTypeEnum_AP         SwitchPortStatsTypeEnum = "ap"
    SwitchPortStatsTypeEnum_BLE        SwitchPortStatsTypeEnum = "ble"
    SwitchPortStatsTypeEnum_ENUMSWITCH SwitchPortStatsTypeEnum = "switch"
    SwitchPortStatsTypeEnum_GATEWAY    SwitchPortStatsTypeEnum = "gateway"
    SwitchPortStatsTypeEnum_MXEDGE     SwitchPortStatsTypeEnum = "mxedge"
    SwitchPortStatsTypeEnum_NAC        SwitchPortStatsTypeEnum = "nac"
)

// SwitchPortUsageDot1XEnum is a string enum.
// Only if `mode`!=`dynamic` if dot1x is desired, set to dot1x
type SwitchPortUsageDot1XEnum string

const (
    SwitchPortUsageDot1XEnum_DOT1X SwitchPortUsageDot1XEnum = "dot1x"
)

// SwitchPortUsageDuplexEnum is a string enum.
// Only if `mode`!=`dynamic` link connection mode
type SwitchPortUsageDuplexEnum string

const (
    SwitchPortUsageDuplexEnum_HALF SwitchPortUsageDuplexEnum = "half"
    SwitchPortUsageDuplexEnum_FULL SwitchPortUsageDuplexEnum = "full"
    SwitchPortUsageDuplexEnum_AUTO SwitchPortUsageDuplexEnum = "auto"
)

// SwitchPortUsageDynamicResetDefaultWhenEnum is a string enum.
// Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage
// Configuring to none will let the DPC port keep at the current port usage.
type SwitchPortUsageDynamicResetDefaultWhenEnum string

const (
    SwitchPortUsageDynamicResetDefaultWhenEnum_NONE     SwitchPortUsageDynamicResetDefaultWhenEnum = "none"
    SwitchPortUsageDynamicResetDefaultWhenEnum_LINKDOWN SwitchPortUsageDynamicResetDefaultWhenEnum = "link_down"
)

// SwitchPortUsageDynamicRuleSrcEnum is a string enum.
type SwitchPortUsageDynamicRuleSrcEnum string

const (
    SwitchPortUsageDynamicRuleSrcEnum_LLDPCHASSISID        SwitchPortUsageDynamicRuleSrcEnum = "lldp_chassis_id"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPSYSTEMNAME       SwitchPortUsageDynamicRuleSrcEnum = "lldp_system_name"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPSERIALNUMBER     SwitchPortUsageDynamicRuleSrcEnum = "lldp_serial_number"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPHARDWAREREVISION SwitchPortUsageDynamicRuleSrcEnum = "lldp_hardware_revision"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPMANUFACTURERNAME SwitchPortUsageDynamicRuleSrcEnum = "lldp_manufacturer_name"
    SwitchPortUsageDynamicRuleSrcEnum_LLDPOUI              SwitchPortUsageDynamicRuleSrcEnum = "lldp_oui"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSUSERNAME       SwitchPortUsageDynamicRuleSrcEnum = "radius_username"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSUSERMAC        SwitchPortUsageDynamicRuleSrcEnum = "radius_usermac"
    SwitchPortUsageDynamicRuleSrcEnum_RADIUSDYNAMICFILTER  SwitchPortUsageDynamicRuleSrcEnum = "radius_dynamicfilter"
    SwitchPortUsageDynamicRuleSrcEnum_LINKPEERMAC          SwitchPortUsageDynamicRuleSrcEnum = "link_peermac"
)

// SwitchPortUsageMacAuthProtocolEnum is a string enum.
// Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled.
type SwitchPortUsageMacAuthProtocolEnum string

const (
    SwitchPortUsageMacAuthProtocolEnum_PAP     SwitchPortUsageMacAuthProtocolEnum = "pap"
    SwitchPortUsageMacAuthProtocolEnum_EAPPEAP SwitchPortUsageMacAuthProtocolEnum = "eap-peap"
    SwitchPortUsageMacAuthProtocolEnum_EAPMD5  SwitchPortUsageMacAuthProtocolEnum = "eap-md5"
)

// SwitchPortUsageModeEnum is a string enum.
// `mode`==`dynamic` must only be used with the port usage with the name `dynamic`
type SwitchPortUsageModeEnum string

const (
    SwitchPortUsageModeEnum_ACCESS  SwitchPortUsageModeEnum = "access"
    SwitchPortUsageModeEnum_TRUNK   SwitchPortUsageModeEnum = "trunk"
    SwitchPortUsageModeEnum_INET    SwitchPortUsageModeEnum = "inet"
    SwitchPortUsageModeEnum_DYNAMIC SwitchPortUsageModeEnum = "dynamic"
)

// SwitchStpConfigTypeEnum is a string enum.
type SwitchStpConfigTypeEnum string

const (
    SwitchStpConfigTypeEnum_RSTP SwitchStpConfigTypeEnum = "rstp"
    SwitchStpConfigTypeEnum_VSTP SwitchStpConfigTypeEnum = "vstp"
)

// SwitchVirtualChassisMemberVcRoleEnum is a string enum.
// Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config
type SwitchVirtualChassisMemberVcRoleEnum string

const (
    SwitchVirtualChassisMemberVcRoleEnum_MASTER   SwitchVirtualChassisMemberVcRoleEnum = "master"
    SwitchVirtualChassisMemberVcRoleEnum_LINECARD SwitchVirtualChassisMemberVcRoleEnum = "linecard"
    SwitchVirtualChassisMemberVcRoleEnum_BACKUP   SwitchVirtualChassisMemberVcRoleEnum = "backup"
)

// TacacsDefaultRoleEnum is a string enum.
type TacacsDefaultRoleEnum string

const (
    TacacsDefaultRoleEnum_ADMIN    TacacsDefaultRoleEnum = "admin"
    TacacsDefaultRoleEnum_NONE     TacacsDefaultRoleEnum = "none"
    TacacsDefaultRoleEnum_READ     TacacsDefaultRoleEnum = "read"
    TacacsDefaultRoleEnum_HELPDESK TacacsDefaultRoleEnum = "helpdesk"
)

// TicketStatusEnum is a string enum.
// Status open: ticket is open, Mist is working on it 
// * pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information) 
// * solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it) 
// * closed: ticket is archived and cannot be changed
type TicketStatusEnum string

const (
    TicketStatusEnum_OPEN    TicketStatusEnum = "open"
    TicketStatusEnum_PENDING TicketStatusEnum = "pending"
    TicketStatusEnum_SOLVED  TicketStatusEnum = "solved"
    TicketStatusEnum_CLOSED  TicketStatusEnum = "closed"
)

// TunnelConfigsAuthAlgoEnum is a string enum.
type TunnelConfigsAuthAlgoEnum string

const (
    TunnelConfigsAuthAlgoEnum_SHA1 TunnelConfigsAuthAlgoEnum = "sha1"
    TunnelConfigsAuthAlgoEnum_SHA2 TunnelConfigsAuthAlgoEnum = "sha2"
    TunnelConfigsAuthAlgoEnum_MD5  TunnelConfigsAuthAlgoEnum = "md5"
)

// TunnelConfigsAutoProvisionRegionEnum is a string enum.
type TunnelConfigsAutoProvisionRegionEnum string

const (
    TunnelConfigsAutoProvisionRegionEnum_AUTO     TunnelConfigsAutoProvisionRegionEnum = "auto"
    TunnelConfigsAutoProvisionRegionEnum_EMEA     TunnelConfigsAutoProvisionRegionEnum = "EMEA"
    TunnelConfigsAutoProvisionRegionEnum_AMERICAS TunnelConfigsAutoProvisionRegionEnum = "Americas"
    TunnelConfigsAutoProvisionRegionEnum_APAC     TunnelConfigsAutoProvisionRegionEnum = "APAC"
)

// TunnelConfigsDhGroupEnum is a string enum.
// Only if:
// * `provider`== `custom-ipsec`
// Values:
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
type TunnelConfigsDhGroupEnum string

const (
    TunnelConfigsDhGroupEnum_ENUM1  TunnelConfigsDhGroupEnum = "1"
    TunnelConfigsDhGroupEnum_ENUM2  TunnelConfigsDhGroupEnum = "2"
    TunnelConfigsDhGroupEnum_ENUM5  TunnelConfigsDhGroupEnum = "5"
    TunnelConfigsDhGroupEnum_ENUM14 TunnelConfigsDhGroupEnum = "14"
    TunnelConfigsDhGroupEnum_ENUM15 TunnelConfigsDhGroupEnum = "15"
    TunnelConfigsDhGroupEnum_ENUM16 TunnelConfigsDhGroupEnum = "16"
    TunnelConfigsDhGroupEnum_ENUM19 TunnelConfigsDhGroupEnum = "19"
    TunnelConfigsDhGroupEnum_ENUM20 TunnelConfigsDhGroupEnum = "20"
    TunnelConfigsDhGroupEnum_ENUM21 TunnelConfigsDhGroupEnum = "21"
    TunnelConfigsDhGroupEnum_ENUM24 TunnelConfigsDhGroupEnum = "24"
)

// TunnelConfigsEncAlgoEnum is a string enum.
type TunnelConfigsEncAlgoEnum string

const (
    TunnelConfigsEncAlgoEnum_AES256    TunnelConfigsEncAlgoEnum = "aes256"
    TunnelConfigsEncAlgoEnum_ENUM3DES  TunnelConfigsEncAlgoEnum = "3des"
    TunnelConfigsEncAlgoEnum_AES128    TunnelConfigsEncAlgoEnum = "aes128"
    TunnelConfigsEncAlgoEnum_AESGCM256 TunnelConfigsEncAlgoEnum = "aes_gcm256"
    TunnelConfigsEncAlgoEnum_AESGCM128 TunnelConfigsEncAlgoEnum = "aes_gcm128"
)

// TunnelProviderOptionsNameEnum is a string enum.
type TunnelProviderOptionsNameEnum string

const (
    TunnelProviderOptionsNameEnum_ZSCALERIPSEC TunnelProviderOptionsNameEnum = "zscaler-ipsec"
    TunnelProviderOptionsNameEnum_ZSCALERGRE   TunnelProviderOptionsNameEnum = "zscaler-gre"
    TunnelProviderOptionsNameEnum_CUSTOMERGRE  TunnelProviderOptionsNameEnum = "customer-gre"
    TunnelProviderOptionsNameEnum_JSEIPSEC     TunnelProviderOptionsNameEnum = "jse-ipsec"
    TunnelProviderOptionsNameEnum_CUSTOMIPSEC  TunnelProviderOptionsNameEnum = "custom-ipsec"
)

// TuntermDhcpdTypeEnum is a string enum.
type TuntermDhcpdTypeEnum string

const (
    TuntermDhcpdTypeEnum_RELAY TuntermDhcpdTypeEnum = "relay"
)

// TunternMonitoringProtocolEnum is a string enum.
type TunternMonitoringProtocolEnum string

const (
    TunternMonitoringProtocolEnum_ARP  TunternMonitoringProtocolEnum = "arp"
    TunternMonitoringProtocolEnum_PING TunternMonitoringProtocolEnum = "ping"
    TunternMonitoringProtocolEnum_TCP  TunternMonitoringProtocolEnum = "tcp"
)

// VpnPathBfdProfileEnum is a string enum.
type VpnPathBfdProfileEnum string

const (
    VpnPathBfdProfileEnum_BROADBAND VpnPathBfdProfileEnum = "broadband"
    VpnPathBfdProfileEnum_LTE       VpnPathBfdProfileEnum = "lte"
)

// WanTunnelProtocolEnum is a string enum.
type WanTunnelProtocolEnum string

const (
    WanTunnelProtocolEnum_IPSEC WanTunnelProtocolEnum = "ipsec"
    WanTunnelProtocolEnum_GRE   WanTunnelProtocolEnum = "gre"
)

// WanTunnelStatsPriorityEnum is a string enum.
type WanTunnelStatsPriorityEnum string

const (
    WanTunnelStatsPriorityEnum_PRIMARY   WanTunnelStatsPriorityEnum = "primary"
    WanTunnelStatsPriorityEnum_SECONDARY WanTunnelStatsPriorityEnum = "secondary"
)

// WebhookOauth2GrantTypeEnum is a string enum.
// required when `type`==`oauth2`
type WebhookOauth2GrantTypeEnum string

const (
    WebhookOauth2GrantTypeEnum_CLIENTCREDENTIALS WebhookOauth2GrantTypeEnum = "client_credentials"
    WebhookOauth2GrantTypeEnum_PASSWORD          WebhookOauth2GrantTypeEnum = "password"
)

// WebhookTopicEnum is a string enum.
type WebhookTopicEnum string

const (
    WebhookTopicEnum_ALARMS            WebhookTopicEnum = "alarms"
    WebhookTopicEnum_AUDITS            WebhookTopicEnum = "audits"
    WebhookTopicEnum_CLIENTINFO        WebhookTopicEnum = "client-info"
    WebhookTopicEnum_CLIENTJOIN        WebhookTopicEnum = "client-join"
    WebhookTopicEnum_CLIENTLATENCY     WebhookTopicEnum = "client-latency"
    WebhookTopicEnum_CLIENTSESSIONS    WebhookTopicEnum = "client-sessions"
    WebhookTopicEnum_DEVICEEVENTS      WebhookTopicEnum = "device_events"
    WebhookTopicEnum_DEVICEUPDOWNS     WebhookTopicEnum = "device-updowns"
    WebhookTopicEnum_MXEDGEEVENTS      WebhookTopicEnum = "mxedge_events"
    WebhookTopicEnum_NACACCOUNTING     WebhookTopicEnum = "nac-accounting"
    WebhookTopicEnum_NACEVENTS         WebhookTopicEnum = "nac_events"
    WebhookTopicEnum_OCCUPANCYALERTS   WebhookTopicEnum = "occupancy-alerts"
    WebhookTopicEnum_RSSIZONE          WebhookTopicEnum = "rssizone"
    WebhookTopicEnum_SDKCLIENTSCANDATA WebhookTopicEnum = "sdkclient_scan_data"
    WebhookTopicEnum_SITESLE           WebhookTopicEnum = "site_sle"
    WebhookTopicEnum_VBEACON           WebhookTopicEnum = "vbeacon"
    WebhookTopicEnum_ZONE              WebhookTopicEnum = "zone"
    WebhookTopicEnum_LOCATION          WebhookTopicEnum = "location"
    WebhookTopicEnum_LOCATIONASSET     WebhookTopicEnum = "location_asset"
    WebhookTopicEnum_LOCATIONCENTRAK   WebhookTopicEnum = "location_centrak"
    WebhookTopicEnum_LOCATIONCLIENT    WebhookTopicEnum = "location_client"
    WebhookTopicEnum_LOCATIONUNCLIENT  WebhookTopicEnum = "location_unclient"
    WebhookTopicEnum_LOCATIONSDK       WebhookTopicEnum = "location_sdk"
    WebhookTopicEnum_ASSETRAW          WebhookTopicEnum = "asset-raw"
    WebhookTopicEnum_ASSETRAWRSSI      WebhookTopicEnum = "asset-raw-rssi"
    WebhookTopicEnum_DISCOVEREDRAWRSSI WebhookTopicEnum = "discovered-raw-rssi"
    WebhookTopicEnum_WIFICONNRAW       WebhookTopicEnum = "wifi-conn-raw"
    WebhookTopicEnum_WIFIUNCONNRAW     WebhookTopicEnum = "wifi-unconn-raw"
)

// WebhookTypeEnum is a string enum.
type WebhookTypeEnum string

const (
    WebhookTypeEnum_HTTPPOST     WebhookTypeEnum = "http-post"
    WebhookTypeEnum_SPLUNK       WebhookTypeEnum = "splunk"
    WebhookTypeEnum_GOOGLEPUBSUB WebhookTypeEnum = "google-pubsub"
    WebhookTypeEnum_AWSSNS       WebhookTypeEnum = "aws-sns"
    WebhookTypeEnum_OAUTH2       WebhookTypeEnum = "oauth2"
)

// WlanApplyToEnum is a string enum.
type WlanApplyToEnum string

const (
    WlanApplyToEnum_SITE   WlanApplyToEnum = "site"
    WlanApplyToEnum_WXTAGS WlanApplyToEnum = "wxtags"
    WlanApplyToEnum_APS    WlanApplyToEnum = "aps"
)

// WlanAuthOweEnum is a string enum.
// `enabled` means transition mode
type WlanAuthOweEnum string

const (
    WlanAuthOweEnum_ENABLED  WlanAuthOweEnum = "enabled"
    WlanAuthOweEnum_DISABLED WlanAuthOweEnum = "disabled"
    WlanAuthOweEnum_REQUIRED WlanAuthOweEnum = "required"
)

// WlanAuthPairwiseItemEnum is a string enum.
type WlanAuthPairwiseItemEnum string

const (
    WlanAuthPairwiseItemEnum_WPA1CCMP WlanAuthPairwiseItemEnum = "wpa1-ccmp"
    WlanAuthPairwiseItemEnum_WPA2TKIP WlanAuthPairwiseItemEnum = "wpa2-tkip"
    WlanAuthPairwiseItemEnum_WPA1TKIP WlanAuthPairwiseItemEnum = "wpa1-tkip"
    WlanAuthPairwiseItemEnum_WPA2CCMP WlanAuthPairwiseItemEnum = "wpa2-ccmp"
    WlanAuthPairwiseItemEnum_WPA3     WlanAuthPairwiseItemEnum = "wpa3"
)

// WlanAuthServerSelectionEnum is a string enum.
// When ordered, AP will prefer and go back to the first server if possible
type WlanAuthServerSelectionEnum string

const (
    WlanAuthServerSelectionEnum_ORDERED   WlanAuthServerSelectionEnum = "ordered"
    WlanAuthServerSelectionEnum_UNORDERED WlanAuthServerSelectionEnum = "unordered"
)

// WlanAuthTypeEnum is a string enum.
type WlanAuthTypeEnum string

const (
    WlanAuthTypeEnum_OPEN        WlanAuthTypeEnum = "open"
    WlanAuthTypeEnum_PSK         WlanAuthTypeEnum = "psk"
    WlanAuthTypeEnum_WEP         WlanAuthTypeEnum = "wep"
    WlanAuthTypeEnum_EAP         WlanAuthTypeEnum = "eap"
    WlanAuthTypeEnum_EAP192      WlanAuthTypeEnum = "eap192"
    WlanAuthTypeEnum_PSKTKIP     WlanAuthTypeEnum = "psk-tkip"
    WlanAuthTypeEnum_PSKWPA2TKIP WlanAuthTypeEnum = "psk-wpa2-tkip"
)

// WlanBonjourServicePropertiesScopeEnum is a string enum.
// how bonjour services should be discovered for the same WLAN
type WlanBonjourServicePropertiesScopeEnum string

const (
    WlanBonjourServicePropertiesScopeEnum_SAMESITE WlanBonjourServicePropertiesScopeEnum = "same_site"
    WlanBonjourServicePropertiesScopeEnum_SAMEMAP  WlanBonjourServicePropertiesScopeEnum = "same_map"
    WlanBonjourServicePropertiesScopeEnum_SAMEAP   WlanBonjourServicePropertiesScopeEnum = "same_ap"
)

// WlanDataratesLegacyItemEnum is a string enum.
type WlanDataratesLegacyItemEnum string

const (
    WlanDataratesLegacyItemEnum_ENUM1   WlanDataratesLegacyItemEnum = "1"
    WlanDataratesLegacyItemEnum_ENUM1B  WlanDataratesLegacyItemEnum = "1b"
    WlanDataratesLegacyItemEnum_ENUM2   WlanDataratesLegacyItemEnum = "2"
    WlanDataratesLegacyItemEnum_ENUM2B  WlanDataratesLegacyItemEnum = "2b"
    WlanDataratesLegacyItemEnum_ENUM55  WlanDataratesLegacyItemEnum = "5.5"
    WlanDataratesLegacyItemEnum_ENUM55B WlanDataratesLegacyItemEnum = "5.5b"
    WlanDataratesLegacyItemEnum_ENUM11  WlanDataratesLegacyItemEnum = "11"
    WlanDataratesLegacyItemEnum_ENUM11B WlanDataratesLegacyItemEnum = "11b"
    WlanDataratesLegacyItemEnum_ENUM6   WlanDataratesLegacyItemEnum = "6"
    WlanDataratesLegacyItemEnum_ENUM6B  WlanDataratesLegacyItemEnum = "6b"
    WlanDataratesLegacyItemEnum_ENUM9   WlanDataratesLegacyItemEnum = "9"
    WlanDataratesLegacyItemEnum_ENUM9B  WlanDataratesLegacyItemEnum = "9b"
    WlanDataratesLegacyItemEnum_ENUM12  WlanDataratesLegacyItemEnum = "12"
    WlanDataratesLegacyItemEnum_ENUM12B WlanDataratesLegacyItemEnum = "12b"
    WlanDataratesLegacyItemEnum_ENUM18  WlanDataratesLegacyItemEnum = "18"
    WlanDataratesLegacyItemEnum_ENUM18B WlanDataratesLegacyItemEnum = "18b"
    WlanDataratesLegacyItemEnum_ENUM24  WlanDataratesLegacyItemEnum = "24"
    WlanDataratesLegacyItemEnum_ENUM24B WlanDataratesLegacyItemEnum = "24b"
    WlanDataratesLegacyItemEnum_ENUM36  WlanDataratesLegacyItemEnum = "36"
    WlanDataratesLegacyItemEnum_ENUM36B WlanDataratesLegacyItemEnum = "36b"
    WlanDataratesLegacyItemEnum_ENUM48  WlanDataratesLegacyItemEnum = "48"
    WlanDataratesLegacyItemEnum_ENUM48B WlanDataratesLegacyItemEnum = "48b"
    WlanDataratesLegacyItemEnum_ENUM54  WlanDataratesLegacyItemEnum = "54"
    WlanDataratesLegacyItemEnum_ENUM54B WlanDataratesLegacyItemEnum = "54b"
)

// WlanDynamicVlanTypeEnum is a string enum.
// standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco)
type WlanDynamicVlanTypeEnum string

const (
    WlanDynamicVlanTypeEnum_STANDARD               WlanDynamicVlanTypeEnum = "standard"
    WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME WlanDynamicVlanTypeEnum = "airespace-interface-name"
)

// WlanHotspot20OperatorsItemEnum is a string enum.
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
// where this WLAN will be connected to
type WlanInterfaceEnum string

const (
    WlanInterfaceEnum_ALL        WlanInterfaceEnum = "all"
    WlanInterfaceEnum_ETH0       WlanInterfaceEnum = "eth0"
    WlanInterfaceEnum_ETH1       WlanInterfaceEnum = "eth1"
    WlanInterfaceEnum_ETH2       WlanInterfaceEnum = "eth2"
    WlanInterfaceEnum_ETH3       WlanInterfaceEnum = "eth3"
    WlanInterfaceEnum_WXTUNNEL   WlanInterfaceEnum = "wxtunnel"
    WlanInterfaceEnum_MXTUNNEL   WlanInterfaceEnum = "mxtunnel"
    WlanInterfaceEnum_SITEMXEDGE WlanInterfaceEnum = "site_mxedge"
)

// WlanPortalAuthEnum is a string enum.
// authentication scheme
type WlanPortalAuthEnum string

const (
    WlanPortalAuthEnum_NONE     WlanPortalAuthEnum = "none"
    WlanPortalAuthEnum_EXTERNAL WlanPortalAuthEnum = "external"
    WlanPortalAuthEnum_SSO      WlanPortalAuthEnum = "sso"
)

// WlanPortalSmsProviderEnum is a string enum.
type WlanPortalSmsProviderEnum string

const (
    WlanPortalSmsProviderEnum_MANUAL     WlanPortalSmsProviderEnum = "manual"
    WlanPortalSmsProviderEnum_TWILIO     WlanPortalSmsProviderEnum = "twilio"
    WlanPortalSmsProviderEnum_BROADNET   WlanPortalSmsProviderEnum = "broadnet"
    WlanPortalSmsProviderEnum_CLICKATELL WlanPortalSmsProviderEnum = "clickatell"
    WlanPortalSmsProviderEnum_PUZZEL     WlanPortalSmsProviderEnum = "puzzel"
    WlanPortalSmsProviderEnum_GUPSHUP    WlanPortalSmsProviderEnum = "gupshup"
    WlanPortalSmsProviderEnum_TELSTRA    WlanPortalSmsProviderEnum = "telstra"
)

// WlanPortalSsoNameidFormatEnum is a string enum.
type WlanPortalSsoNameidFormatEnum string

const (
    WlanPortalSsoNameidFormatEnum_EMAIL       WlanPortalSsoNameidFormatEnum = "email"
    WlanPortalSsoNameidFormatEnum_UNSPECIFIED WlanPortalSsoNameidFormatEnum = "unspecified"
)

// WlanQosClassEnum is a string enum.
type WlanQosClassEnum string

const (
    WlanQosClassEnum_BACKGROUND WlanQosClassEnum = "background"
    WlanQosClassEnum_BESTEFFORT WlanQosClassEnum = "best_effort"
    WlanQosClassEnum_VIDEO      WlanQosClassEnum = "video"
    WlanQosClassEnum_VOICE      WlanQosClassEnum = "voice"
)

// WlanRoamModeEnum is a string enum.
type WlanRoamModeEnum string

const (
    WlanRoamModeEnum_NONE    WlanRoamModeEnum = "none"
    WlanRoamModeEnum_OKC     WlanRoamModeEnum = "OKC"
    WlanRoamModeEnum_ENUM11R WlanRoamModeEnum = "11r"
)

// WxlanRuleActionEnum is a string enum.
// type of action, allow / block
type WxlanRuleActionEnum string

const (
    WxlanRuleActionEnum_ALLOW WxlanRuleActionEnum = "allow"
    WxlanRuleActionEnum_BLOCK WxlanRuleActionEnum = "block"
)

// WxlanTagMatchEnum is a string enum.
type WxlanTagMatchEnum string

const (
    WxlanTagMatchEnum_APID           WxlanTagMatchEnum = "ap_id"
    WxlanTagMatchEnum_APP            WxlanTagMatchEnum = "app"
    WxlanTagMatchEnum_ASSETMAC       WxlanTagMatchEnum = "asset_mac"
    WxlanTagMatchEnum_CLIENTMAC      WxlanTagMatchEnum = "client_mac"
    WxlanTagMatchEnum_HOSTNAME       WxlanTagMatchEnum = "hostname"
    WxlanTagMatchEnum_IPRANGESUBNET  WxlanTagMatchEnum = "ip_range_subnet"
    WxlanTagMatchEnum_PORT           WxlanTagMatchEnum = "port"
    WxlanTagMatchEnum_RADIUSATTR     WxlanTagMatchEnum = "radius_attr"
    WxlanTagMatchEnum_RADIUSGROUP    WxlanTagMatchEnum = "radius_group"
    WxlanTagMatchEnum_RADIUSUSERNAME WxlanTagMatchEnum = "radius_username"
    WxlanTagMatchEnum_WLANID         WxlanTagMatchEnum = "wlan_id"
    WxlanTagMatchEnum_PSKNAME        WxlanTagMatchEnum = "psk_name"
    WxlanTagMatchEnum_PSKROLE        WxlanTagMatchEnum = "psk_role"
)

// WxlanTagOperationEnum is a string enum.
type WxlanTagOperationEnum string

const (
    WxlanTagOperationEnum_IN    WxlanTagOperationEnum = "in"
    WxlanTagOperationEnum_NOTIN WxlanTagOperationEnum = "not_in"
)

// WxlanTagTypeEnum is a string enum.
type WxlanTagTypeEnum string

const (
    WxlanTagTypeEnum_MATCH    WxlanTagTypeEnum = "match"
    WxlanTagTypeEnum_CLIENT   WxlanTagTypeEnum = "client"
    WxlanTagTypeEnum_RESOURCE WxlanTagTypeEnum = "resource"
    WxlanTagTypeEnum_SUBNET   WxlanTagTypeEnum = "subnet"
    WxlanTagTypeEnum_SPEC     WxlanTagTypeEnum = "spec"
    WxlanTagTypeEnum_VLAN     WxlanTagTypeEnum = "vlan"
)

// WxlanTunnelSessionEthertypeEnum is a string enum.
type WxlanTunnelSessionEthertypeEnum string

const (
    WxlanTunnelSessionEthertypeEnum_ETHERNET WxlanTunnelSessionEthertypeEnum = "ethernet"
    WxlanTunnelSessionEthertypeEnum_VLAN     WxlanTunnelSessionEthertypeEnum = "vlan"
)

// MapTypeEnum is a string enum.
type MapTypeEnum string

const (
    MapTypeEnum_IMAGE  MapTypeEnum = "image"
    MapTypeEnum_GOOGLE MapTypeEnum = "google"
)

// MapViewEnum is a string enum.
// when type=google
type MapViewEnum string

const (
    MapViewEnum_ROADMAP   MapViewEnum = "roadmap"
    MapViewEnum_SATELLITE MapViewEnum = "satellite"
    MapViewEnum_HYBRID    MapViewEnum = "hybrid"
    MapViewEnum_TERRAIN   MapViewEnum = "terrain"
)

// VirtualChassisConfigMemberVcRoleEnum is a string enum.
type VirtualChassisConfigMemberVcRoleEnum string

const (
    VirtualChassisConfigMemberVcRoleEnum_MASTER   VirtualChassisConfigMemberVcRoleEnum = "master"
    VirtualChassisConfigMemberVcRoleEnum_BACKUP   VirtualChassisConfigMemberVcRoleEnum = "backup"
    VirtualChassisConfigMemberVcRoleEnum_LINECARD VirtualChassisConfigMemberVcRoleEnum = "linecard"
)

// VirtualChassisMemberUpdateVcRoleEnum is a string enum.
// Required if `op`==`add` or `op`==`preprovision`
type VirtualChassisMemberUpdateVcRoleEnum string

const (
    VirtualChassisMemberUpdateVcRoleEnum_MASTER   VirtualChassisMemberUpdateVcRoleEnum = "master"
    VirtualChassisMemberUpdateVcRoleEnum_BACKUP   VirtualChassisMemberUpdateVcRoleEnum = "backup"
    VirtualChassisMemberUpdateVcRoleEnum_LINECARD VirtualChassisMemberUpdateVcRoleEnum = "linecard"
)

// VirtualChassisUpdateOpEnum is a string enum.
type VirtualChassisUpdateOpEnum string

const (
    VirtualChassisUpdateOpEnum_ADD          VirtualChassisUpdateOpEnum = "add"
    VirtualChassisUpdateOpEnum_REMOVE       VirtualChassisUpdateOpEnum = "remove"
    VirtualChassisUpdateOpEnum_RENUMBER     VirtualChassisUpdateOpEnum = "renumber"
    VirtualChassisUpdateOpEnum_PREPROVISION VirtualChassisUpdateOpEnum = "preprovision"
)

// MfaSecretTypeEnum is a string enum.
type MfaSecretTypeEnum string

const (
    MfaSecretTypeEnum_QRCODE MfaSecretTypeEnum = "qrcode"
)

// RecaptchaFlavorEnum is a string enum.
// flavor of the captcha
type RecaptchaFlavorEnum string

const (
    RecaptchaFlavorEnum_HCAPTCHA RecaptchaFlavorEnum = "hcaptcha"
    RecaptchaFlavorEnum_GOOGLE   RecaptchaFlavorEnum = "google"
)

// ConstDeviceApExtiosDefaultDirEnum is a string enum.
type ConstDeviceApExtiosDefaultDirEnum string

const (
    ConstDeviceApExtiosDefaultDirEnum_IN  ConstDeviceApExtiosDefaultDirEnum = "IN"
    ConstDeviceApExtiosDefaultDirEnum_OUT ConstDeviceApExtiosDefaultDirEnum = "OUT"
)

// ConstInsightMetricsPropertyScopeEnum is a string enum.
type ConstInsightMetricsPropertyScopeEnum string

const (
    ConstInsightMetricsPropertyScopeEnum_SITE       ConstInsightMetricsPropertyScopeEnum = "site"
    ConstInsightMetricsPropertyScopeEnum_AP         ConstInsightMetricsPropertyScopeEnum = "ap"
    ConstInsightMetricsPropertyScopeEnum_CLIENT     ConstInsightMetricsPropertyScopeEnum = "client"
    ConstInsightMetricsPropertyScopeEnum_ENUMSWITCH ConstInsightMetricsPropertyScopeEnum = "switch"
    ConstInsightMetricsPropertyScopeEnum_DEVICE     ConstInsightMetricsPropertyScopeEnum = "device"
    ConstInsightMetricsPropertyScopeEnum_MXEDGE     ConstInsightMetricsPropertyScopeEnum = "mxedge"
)

// HaClusterNodeEnum is a string enum.
// only for HA
type HaClusterNodeEnum string

const (
    HaClusterNodeEnum_NODE0 HaClusterNodeEnum = "node0"
    HaClusterNodeEnum_NODE1 HaClusterNodeEnum = "node1"
)

// ScanDataItemBandEnum is a string enum.
// 5GHz or 2.4GHz band, associated with the BSSID scanned
type ScanDataItemBandEnum string

const (
    ScanDataItemBandEnum_ENUM24 ScanDataItemBandEnum = "2.4"
    ScanDataItemBandEnum_ENUM5  ScanDataItemBandEnum = "5"
)

// WebhookClientInfoTopicEnum is a string enum.
type WebhookClientInfoTopicEnum string

const (
    WebhookClientInfoTopicEnum_CLIENTINFO WebhookClientInfoTopicEnum = "client-info"
)

// WebhookDeviceEventsEventDeviceTypeEnum is a string enum.
type WebhookDeviceEventsEventDeviceTypeEnum string

const (
    WebhookDeviceEventsEventDeviceTypeEnum_AP         WebhookDeviceEventsEventDeviceTypeEnum = "ap"
    WebhookDeviceEventsEventDeviceTypeEnum_ENUMSWITCH WebhookDeviceEventsEventDeviceTypeEnum = "switch"
    WebhookDeviceEventsEventDeviceTypeEnum_GATEWAY    WebhookDeviceEventsEventDeviceTypeEnum = "gateway"
)

// WebhookDeviceEventsEventEvTypeEnum is a string enum.
// (optional) event advisory (notice/warn)
type WebhookDeviceEventsEventEvTypeEnum string

const (
    WebhookDeviceEventsEventEvTypeEnum_NOTICE WebhookDeviceEventsEventEvTypeEnum = "notice"
    WebhookDeviceEventsEventEvTypeEnum_WARN   WebhookDeviceEventsEventEvTypeEnum = "warn"
)

// WebhookOccupancyAlertTypeEnum is a string enum.
// event type (COMPLIANCE-VIOLATION / COMPLIANCE-OK)
type WebhookOccupancyAlertTypeEnum string

const (
    WebhookOccupancyAlertTypeEnum_COMPLIANCEVIOLATION WebhookOccupancyAlertTypeEnum = "COMPLIANCE-VIOLATION"
    WebhookOccupancyAlertTypeEnum_COMPLIANCEOK        WebhookOccupancyAlertTypeEnum = "COMPLIANCE-OK"
)

// WebhookZoneEventTriggerEnum is a string enum.
// enter / exit
type WebhookZoneEventTriggerEnum string

const (
    WebhookZoneEventTriggerEnum_ENTER WebhookZoneEventTriggerEnum = "enter"
    WebhookZoneEventTriggerEnum_EXIT  WebhookZoneEventTriggerEnum = "exit"
)

// ListMspLogsSortEnum is a string enum.
type ListMspLogsSortEnum string

const (
    ListMspLogsSortEnum_TIMESTAMP     ListMspLogsSortEnum = "timestamp"
    ListMspLogsSortEnum_ENUMTIMESTAMP ListMspLogsSortEnum = "-timestamp"
    ListMspLogsSortEnum_SITEID        ListMspLogsSortEnum = "site_id"
    ListMspLogsSortEnum_ADMINID       ListMspLogsSortEnum = "admin_id"
)

// MspLogsCountDistinctEnum is a string enum.
type MspLogsCountDistinctEnum string

const (
    MspLogsCountDistinctEnum_ADMINID   MspLogsCountDistinctEnum = "admin_id"
    MspLogsCountDistinctEnum_ADMINNAME MspLogsCountDistinctEnum = "admin_name"
    MspLogsCountDistinctEnum_MESSAGE   MspLogsCountDistinctEnum = "message"
    MspLogsCountDistinctEnum_ORGID     MspLogsCountDistinctEnum = "org_id"
)

// MspMarvisSuggestionsCountDistinctEnum is a string enum.
type MspMarvisSuggestionsCountDistinctEnum string

const (
    MspMarvisSuggestionsCountDistinctEnum_ORGID  MspMarvisSuggestionsCountDistinctEnum = "org_id"
    MspMarvisSuggestionsCountDistinctEnum_STATUS MspMarvisSuggestionsCountDistinctEnum = "status"
)

// MspSearchTypeEnum is a string enum.
type MspSearchTypeEnum string

const (
    MspSearchTypeEnum_ORGS MspSearchTypeEnum = "orgs"
)

// MspTicketsCountDistinctEnum is a string enum.
type MspTicketsCountDistinctEnum string

const (
    MspTicketsCountDistinctEnum_STATUS   MspTicketsCountDistinctEnum = "status"
    MspTicketsCountDistinctEnum_ENUMTYPE MspTicketsCountDistinctEnum = "type"
    MspTicketsCountDistinctEnum_ORGID    MspTicketsCountDistinctEnum = "org_id"
)

// MspLicenseActionOperationEnum is a string enum.
type MspLicenseActionOperationEnum string

const (
    MspLicenseActionOperationEnum_AMEND    MspLicenseActionOperationEnum = "amend"
    MspLicenseActionOperationEnum_UNAMEND  MspLicenseActionOperationEnum = "unamend"
    MspLicenseActionOperationEnum_DELETE   MspLicenseActionOperationEnum = "delete"
    MspLicenseActionOperationEnum_ANNOTATE MspLicenseActionOperationEnum = "annotate"
)

// MspOrgChangeOperationEnum is a string enum.
type MspOrgChangeOperationEnum string

const (
    MspOrgChangeOperationEnum_ASSIGN   MspOrgChangeOperationEnum = "assign"
    MspOrgChangeOperationEnum_UNASSIGN MspOrgChangeOperationEnum = "unassign"
)

// MspTierEnum is a string enum.
type MspTierEnum string

const (
    MspTierEnum_BASE     MspTierEnum = "base"
    MspTierEnum_ADVANCED MspTierEnum = "advanced"
)

// AlarmCountDisctinctEnum is a string enum.
type AlarmCountDisctinctEnum string

const (
    AlarmCountDisctinctEnum_ENUMTYPE AlarmCountDisctinctEnum = "type"
    AlarmCountDisctinctEnum_ACKED    AlarmCountDisctinctEnum = "acked"
    AlarmCountDisctinctEnum_SEVERITY AlarmCountDisctinctEnum = "severity"
    AlarmCountDisctinctEnum_GROUP    AlarmCountDisctinctEnum = "group"
)

// CountSiteCallsDistrinctEnum is a string enum.
type CountSiteCallsDistrinctEnum string

const (
    CountSiteCallsDistrinctEnum_MAC CountSiteCallsDistrinctEnum = "mac"
)

// CountSiteSwOrGwPortsAuthStateEnum is a string enum.
type CountSiteSwOrGwPortsAuthStateEnum string

const (
    CountSiteSwOrGwPortsAuthStateEnum_INIT           CountSiteSwOrGwPortsAuthStateEnum = "init"
    CountSiteSwOrGwPortsAuthStateEnum_AUTHENTICATED  CountSiteSwOrGwPortsAuthStateEnum = "authenticated"
    CountSiteSwOrGwPortsAuthStateEnum_AUTHENTICATING CountSiteSwOrGwPortsAuthStateEnum = "authenticating"
    CountSiteSwOrGwPortsAuthStateEnum_HELD           CountSiteSwOrGwPortsAuthStateEnum = "held"
)

// CountSiteSwOrGwPortsStpRoleEnum is a string enum.
type CountSiteSwOrGwPortsStpRoleEnum string

const (
    CountSiteSwOrGwPortsStpRoleEnum_DESIGNATED    CountSiteSwOrGwPortsStpRoleEnum = "designated"
    CountSiteSwOrGwPortsStpRoleEnum_BACKUP        CountSiteSwOrGwPortsStpRoleEnum = "backup"
    CountSiteSwOrGwPortsStpRoleEnum_ALTERNATE     CountSiteSwOrGwPortsStpRoleEnum = "alternate"
    CountSiteSwOrGwPortsStpRoleEnum_ROOT          CountSiteSwOrGwPortsStpRoleEnum = "root"
    CountSiteSwOrGwPortsStpRoleEnum_ROOTPREVENTED CountSiteSwOrGwPortsStpRoleEnum = "root-prevented"
)

// CountSiteSwOrGwPortsStpStateEnum is a string enum.
type CountSiteSwOrGwPortsStpStateEnum string

const (
    CountSiteSwOrGwPortsStpStateEnum_FORWARDING CountSiteSwOrGwPortsStpStateEnum = "forwarding"
    CountSiteSwOrGwPortsStpStateEnum_BLOCKING   CountSiteSwOrGwPortsStpStateEnum = "blocking"
    CountSiteSwOrGwPortsStpStateEnum_LEARNING   CountSiteSwOrGwPortsStpStateEnum = "learning"
    CountSiteSwOrGwPortsStpStateEnum_LISTENING  CountSiteSwOrGwPortsStpStateEnum = "listening"
    CountSiteSwOrGwPortsStpStateEnum_DISABLED   CountSiteSwOrGwPortsStpStateEnum = "disabled"
)

// CountSiteSwitchPortsAuthStateEnum is a string enum.
type CountSiteSwitchPortsAuthStateEnum string

const (
    CountSiteSwitchPortsAuthStateEnum_INIT           CountSiteSwitchPortsAuthStateEnum = "init"
    CountSiteSwitchPortsAuthStateEnum_AUTHENTICATED  CountSiteSwitchPortsAuthStateEnum = "authenticated"
    CountSiteSwitchPortsAuthStateEnum_AUTHENTICATING CountSiteSwitchPortsAuthStateEnum = "authenticating"
    CountSiteSwitchPortsAuthStateEnum_HELD           CountSiteSwitchPortsAuthStateEnum = "held"
)

// CountSiteSwitchPortsStpRoleEnum is a string enum.
type CountSiteSwitchPortsStpRoleEnum string

const (
    CountSiteSwitchPortsStpRoleEnum_DESIGNATED    CountSiteSwitchPortsStpRoleEnum = "designated"
    CountSiteSwitchPortsStpRoleEnum_BACKUP        CountSiteSwitchPortsStpRoleEnum = "backup"
    CountSiteSwitchPortsStpRoleEnum_ALTERNATE     CountSiteSwitchPortsStpRoleEnum = "alternate"
    CountSiteSwitchPortsStpRoleEnum_ROOT          CountSiteSwitchPortsStpRoleEnum = "root"
    CountSiteSwitchPortsStpRoleEnum_ROOTPREVENTED CountSiteSwitchPortsStpRoleEnum = "root-prevented"
)

// CountSiteSwitchPortsStpStateEnum is a string enum.
type CountSiteSwitchPortsStpStateEnum string

const (
    CountSiteSwitchPortsStpStateEnum_FORWARDING CountSiteSwitchPortsStpStateEnum = "forwarding"
    CountSiteSwitchPortsStpStateEnum_BLOCKING   CountSiteSwitchPortsStpStateEnum = "blocking"
    CountSiteSwitchPortsStpStateEnum_LEARNING   CountSiteSwitchPortsStpStateEnum = "learning"
    CountSiteSwitchPortsStpStateEnum_LISTENING  CountSiteSwitchPortsStpStateEnum = "listening"
    CountSiteSwitchPortsStpStateEnum_DISABLED   CountSiteSwitchPortsStpStateEnum = "disabled"
)

// DiscoveredSwitchMetricTypeEnum is a string enum.
type DiscoveredSwitchMetricTypeEnum string

const (
    DiscoveredSwitchMetricTypeEnum_INACTIVEWIREDVLANS DiscoveredSwitchMetricTypeEnum = "inactive_wired_vlans"
    DiscoveredSwitchMetricTypeEnum_SWITCHAPAFFINITY   DiscoveredSwitchMetricTypeEnum = "switch_ap_affinity"
    DiscoveredSwitchMetricTypeEnum_POECOMPLIANCE      DiscoveredSwitchMetricTypeEnum = "poe_compliance"
    DiscoveredSwitchMetricTypeEnum_VERSIONCOMPLIANCE  DiscoveredSwitchMetricTypeEnum = "version_compliance"
)

// DiscoveredSwitchesMetricScopeEnum is a string enum.
type DiscoveredSwitchesMetricScopeEnum string

const (
    DiscoveredSwitchesMetricScopeEnum_SITE       DiscoveredSwitchesMetricScopeEnum = "site"
    DiscoveredSwitchesMetricScopeEnum_ENUMSWITCH DiscoveredSwitchesMetricScopeEnum = "switch"
)

// FastRoamResultEnum is a string enum.
type FastRoamResultEnum string

const (
    FastRoamResultEnum_SUCCESS FastRoamResultEnum = "success"
    FastRoamResultEnum_FAIL    FastRoamResultEnum = "fail"
    FastRoamResultEnum_NONE    FastRoamResultEnum = "none"
)

// ImportSiteAssetsUpsertEnum is a string enum.
type ImportSiteAssetsUpsertEnum string

const (
    ImportSiteAssetsUpsertEnum_TRUE  ImportSiteAssetsUpsertEnum = "True"
    ImportSiteAssetsUpsertEnum_FALSE ImportSiteAssetsUpsertEnum = "False"
)

// RfClientTypeEnum is a string enum.
type RfClientTypeEnum string

const (
    RfClientTypeEnum_SDKCLIENT RfClientTypeEnum = "sdkclient"
    RfClientTypeEnum_CLIENT    RfClientTypeEnum = "client"
    RfClientTypeEnum_ASSET     RfClientTypeEnum = "asset"
)

// SearchSiteDevicesDescSortEnum is a string enum.
type SearchSiteDevicesDescSortEnum string

const (
    SearchSiteDevicesDescSortEnum_TIMESTAMP SearchSiteDevicesDescSortEnum = "timestamp"
    SearchSiteDevicesDescSortEnum_MAC       SearchSiteDevicesDescSortEnum = "mac"
    SearchSiteDevicesDescSortEnum_MODEL     SearchSiteDevicesDescSortEnum = "model"
    SearchSiteDevicesDescSortEnum_SKU       SearchSiteDevicesDescSortEnum = "sku"
)

// SearchSiteDevicesMxtunnelStatusEnum is a string enum.
type SearchSiteDevicesMxtunnelStatusEnum string

const (
    SearchSiteDevicesMxtunnelStatusEnum_UP   SearchSiteDevicesMxtunnelStatusEnum = "up"
    SearchSiteDevicesMxtunnelStatusEnum_DOWN SearchSiteDevicesMxtunnelStatusEnum = "down"
)

// SearchSiteDevicesSortEnum is a string enum.
type SearchSiteDevicesSortEnum string

const (
    SearchSiteDevicesSortEnum_TIMESTAMP SearchSiteDevicesSortEnum = "timestamp"
    SearchSiteDevicesSortEnum_MAC       SearchSiteDevicesSortEnum = "mac"
    SearchSiteDevicesSortEnum_MODEL     SearchSiteDevicesSortEnum = "model"
    SearchSiteDevicesSortEnum_SKU       SearchSiteDevicesSortEnum = "sku"
)

// SearchSiteSwOrGwPortsAuthStateEnum is a string enum.
type SearchSiteSwOrGwPortsAuthStateEnum string

const (
    SearchSiteSwOrGwPortsAuthStateEnum_INIT           SearchSiteSwOrGwPortsAuthStateEnum = "init"
    SearchSiteSwOrGwPortsAuthStateEnum_AUTHENTICATED  SearchSiteSwOrGwPortsAuthStateEnum = "authenticated"
    SearchSiteSwOrGwPortsAuthStateEnum_AUTHENTICATING SearchSiteSwOrGwPortsAuthStateEnum = "authenticating"
    SearchSiteSwOrGwPortsAuthStateEnum_HELD           SearchSiteSwOrGwPortsAuthStateEnum = "held"
)

// SearchSiteSwOrGwPortsDeviceTypeEnum is a string enum.
type SearchSiteSwOrGwPortsDeviceTypeEnum string

const (
    SearchSiteSwOrGwPortsDeviceTypeEnum_AP         SearchSiteSwOrGwPortsDeviceTypeEnum = "ap"
    SearchSiteSwOrGwPortsDeviceTypeEnum_BLE        SearchSiteSwOrGwPortsDeviceTypeEnum = "ble"
    SearchSiteSwOrGwPortsDeviceTypeEnum_ENUMSWITCH SearchSiteSwOrGwPortsDeviceTypeEnum = "switch"
    SearchSiteSwOrGwPortsDeviceTypeEnum_GATEWAY    SearchSiteSwOrGwPortsDeviceTypeEnum = "gateway"
    SearchSiteSwOrGwPortsDeviceTypeEnum_MXEDGE     SearchSiteSwOrGwPortsDeviceTypeEnum = "mxedge"
    SearchSiteSwOrGwPortsDeviceTypeEnum_NAC        SearchSiteSwOrGwPortsDeviceTypeEnum = "nac"
)

// SearchSiteSwOrGwPortsStpRoleEnum is a string enum.
type SearchSiteSwOrGwPortsStpRoleEnum string

const (
    SearchSiteSwOrGwPortsStpRoleEnum_DESIGNATED    SearchSiteSwOrGwPortsStpRoleEnum = "designated"
    SearchSiteSwOrGwPortsStpRoleEnum_BACKUP        SearchSiteSwOrGwPortsStpRoleEnum = "backup"
    SearchSiteSwOrGwPortsStpRoleEnum_ALTERNATE     SearchSiteSwOrGwPortsStpRoleEnum = "alternate"
    SearchSiteSwOrGwPortsStpRoleEnum_ROOT          SearchSiteSwOrGwPortsStpRoleEnum = "root"
    SearchSiteSwOrGwPortsStpRoleEnum_ROOTPREVENTED SearchSiteSwOrGwPortsStpRoleEnum = "root-prevented"
)

// SearchSiteSwOrGwPortsStpStateEnum is a string enum.
type SearchSiteSwOrGwPortsStpStateEnum string

const (
    SearchSiteSwOrGwPortsStpStateEnum_FORWARDING SearchSiteSwOrGwPortsStpStateEnum = "forwarding"
    SearchSiteSwOrGwPortsStpStateEnum_BLOCKING   SearchSiteSwOrGwPortsStpStateEnum = "blocking"
    SearchSiteSwOrGwPortsStpStateEnum_LEARNING   SearchSiteSwOrGwPortsStpStateEnum = "learning"
    SearchSiteSwOrGwPortsStpStateEnum_LISTENING  SearchSiteSwOrGwPortsStpStateEnum = "listening"
    SearchSiteSwOrGwPortsStpStateEnum_DISABLED   SearchSiteSwOrGwPortsStpStateEnum = "disabled"
)

// SearchSiteSwitchPortsAuthStateEnum is a string enum.
type SearchSiteSwitchPortsAuthStateEnum string

const (
    SearchSiteSwitchPortsAuthStateEnum_INIT           SearchSiteSwitchPortsAuthStateEnum = "init"
    SearchSiteSwitchPortsAuthStateEnum_AUTHENTICATED  SearchSiteSwitchPortsAuthStateEnum = "authenticated"
    SearchSiteSwitchPortsAuthStateEnum_AUTHENTICATING SearchSiteSwitchPortsAuthStateEnum = "authenticating"
    SearchSiteSwitchPortsAuthStateEnum_HELD           SearchSiteSwitchPortsAuthStateEnum = "held"
)

// SearchSiteSwitchPortsStpRoleEnum is a string enum.
type SearchSiteSwitchPortsStpRoleEnum string

const (
    SearchSiteSwitchPortsStpRoleEnum_DESIGNATED    SearchSiteSwitchPortsStpRoleEnum = "designated"
    SearchSiteSwitchPortsStpRoleEnum_BACKUP        SearchSiteSwitchPortsStpRoleEnum = "backup"
    SearchSiteSwitchPortsStpRoleEnum_ALTERNATE     SearchSiteSwitchPortsStpRoleEnum = "alternate"
    SearchSiteSwitchPortsStpRoleEnum_ROOT          SearchSiteSwitchPortsStpRoleEnum = "root"
    SearchSiteSwitchPortsStpRoleEnum_ROOTPREVENTED SearchSiteSwitchPortsStpRoleEnum = "root-prevented"
)

// SearchSiteSwitchPortsStpStateEnum is a string enum.
type SearchSiteSwitchPortsStpStateEnum string

const (
    SearchSiteSwitchPortsStpStateEnum_FORWARDING SearchSiteSwitchPortsStpStateEnum = "forwarding"
    SearchSiteSwitchPortsStpStateEnum_BLOCKING   SearchSiteSwitchPortsStpStateEnum = "blocking"
    SearchSiteSwitchPortsStpStateEnum_LEARNING   SearchSiteSwitchPortsStpStateEnum = "learning"
    SearchSiteSwitchPortsStpStateEnum_LISTENING  SearchSiteSwitchPortsStpStateEnum = "listening"
    SearchSiteSwitchPortsStpStateEnum_DISABLED   SearchSiteSwitchPortsStpStateEnum = "disabled"
)

// SiteAppsCountDistinctEnum is a string enum.
type SiteAppsCountDistinctEnum string

const (
    SiteAppsCountDistinctEnum_AP            SiteAppsCountDistinctEnum = "ap"
    SiteAppsCountDistinctEnum_WCID          SiteAppsCountDistinctEnum = "wcid"
    SiteAppsCountDistinctEnum_SSID          SiteAppsCountDistinctEnum = "ssid"
    SiteAppsCountDistinctEnum_ENUMWLANIDAPP SiteAppsCountDistinctEnum = "wlan_id app"
    SiteAppsCountDistinctEnum_DEVICEMAC     SiteAppsCountDistinctEnum = "device_mac"
    SiteAppsCountDistinctEnum_SRCIP         SiteAppsCountDistinctEnum = "src_ip"
    SiteAppsCountDistinctEnum_PORTID        SiteAppsCountDistinctEnum = "port_id"
    SiteAppsCountDistinctEnum_APP           SiteAppsCountDistinctEnum = "app"
    SiteAppsCountDistinctEnum_CATEGORY      SiteAppsCountDistinctEnum = "category"
    SiteAppsCountDistinctEnum_SERVICE       SiteAppsCountDistinctEnum = "service"
)

// SiteAssetsCountDistinctEnum is a string enum.
type SiteAssetsCountDistinctEnum string

const (
    SiteAssetsCountDistinctEnum_MAC                   SiteAssetsCountDistinctEnum = "mac"
    SiteAssetsCountDistinctEnum_MAPID                 SiteAssetsCountDistinctEnum = "map_id"
    SiteAssetsCountDistinctEnum_IBEACONUUID           SiteAssetsCountDistinctEnum = "ibeacon_uuid"
    SiteAssetsCountDistinctEnum_IBEACONMAJOR          SiteAssetsCountDistinctEnum = "ibeacon_major"
    SiteAssetsCountDistinctEnum_IBEACONMINOR          SiteAssetsCountDistinctEnum = "ibeacon_minor"
    SiteAssetsCountDistinctEnum_EDDYSTONEUIDNAMESPACE SiteAssetsCountDistinctEnum = "eddystone_uid_namespace"
    SiteAssetsCountDistinctEnum_EDDYSTONEUIDINSTANCE  SiteAssetsCountDistinctEnum = "eddystone_uid_instance"
    SiteAssetsCountDistinctEnum_EDDYSTONEURL          SiteAssetsCountDistinctEnum = "eddystone_url"
    SiteAssetsCountDistinctEnum_BY                    SiteAssetsCountDistinctEnum = "by"
    SiteAssetsCountDistinctEnum_NAME                  SiteAssetsCountDistinctEnum = "name"
    SiteAssetsCountDistinctEnum_DEVICENAME            SiteAssetsCountDistinctEnum = "device_name"
)

// SiteClientEventsCountDistinctEnum is a string enum.
type SiteClientEventsCountDistinctEnum string

const (
    SiteClientEventsCountDistinctEnum_ENUMTYPE SiteClientEventsCountDistinctEnum = "type"
    SiteClientEventsCountDistinctEnum_PROTO    SiteClientEventsCountDistinctEnum = "proto"
    SiteClientEventsCountDistinctEnum_BAND     SiteClientEventsCountDistinctEnum = "band"
    SiteClientEventsCountDistinctEnum_CHANNEL  SiteClientEventsCountDistinctEnum = "channel"
    SiteClientEventsCountDistinctEnum_WLANID   SiteClientEventsCountDistinctEnum = "wlan_id"
    SiteClientEventsCountDistinctEnum_SSID     SiteClientEventsCountDistinctEnum = "ssid"
)

// SiteClientSessionsCountDistinctEnum is a string enum.
type SiteClientSessionsCountDistinctEnum string

const (
    SiteClientSessionsCountDistinctEnum_SSID              SiteClientSessionsCountDistinctEnum = "ssid"
    SiteClientSessionsCountDistinctEnum_WLANID            SiteClientSessionsCountDistinctEnum = "wlan_id"
    SiteClientSessionsCountDistinctEnum_AP                SiteClientSessionsCountDistinctEnum = "ap"
    SiteClientSessionsCountDistinctEnum_MAC               SiteClientSessionsCountDistinctEnum = "mac"
    SiteClientSessionsCountDistinctEnum_CLIENTFAMILY      SiteClientSessionsCountDistinctEnum = "client_family"
    SiteClientSessionsCountDistinctEnum_CLIENTMANUFACTURE SiteClientSessionsCountDistinctEnum = "client_manufacture"
    SiteClientSessionsCountDistinctEnum_CLIENTMODEL       SiteClientSessionsCountDistinctEnum = "client_model"
    SiteClientSessionsCountDistinctEnum_CLIENTOS          SiteClientSessionsCountDistinctEnum = "client_os"
)

// SiteClientsCountDistinctEnum is a string enum.
type SiteClientsCountDistinctEnum string

const (
    SiteClientsCountDistinctEnum_SSID     SiteClientsCountDistinctEnum = "ssid"
    SiteClientsCountDistinctEnum_AP       SiteClientsCountDistinctEnum = "ap"
    SiteClientsCountDistinctEnum_IP       SiteClientsCountDistinctEnum = "ip"
    SiteClientsCountDistinctEnum_VLAN     SiteClientsCountDistinctEnum = "vlan"
    SiteClientsCountDistinctEnum_HOSTNAME SiteClientsCountDistinctEnum = "hostname"
    SiteClientsCountDistinctEnum_OS       SiteClientsCountDistinctEnum = "os"
    SiteClientsCountDistinctEnum_MODEL    SiteClientsCountDistinctEnum = "model"
    SiteClientsCountDistinctEnum_DEVICE   SiteClientsCountDistinctEnum = "device"
)

// SiteDeviceEventsCountDistinctEnum is a string enum.
type SiteDeviceEventsCountDistinctEnum string

const (
    SiteDeviceEventsCountDistinctEnum_MODEL    SiteDeviceEventsCountDistinctEnum = "model"
    SiteDeviceEventsCountDistinctEnum_ENUMTYPE SiteDeviceEventsCountDistinctEnum = "type"
    SiteDeviceEventsCountDistinctEnum_TYPECODE SiteDeviceEventsCountDistinctEnum = "type_code"
    SiteDeviceEventsCountDistinctEnum_MAC      SiteDeviceEventsCountDistinctEnum = "mac"
)

// SiteDeviceLastConfigCountDistinctEnum is a string enum.
type SiteDeviceLastConfigCountDistinctEnum string

const (
    SiteDeviceLastConfigCountDistinctEnum_VERSION SiteDeviceLastConfigCountDistinctEnum = "version"
    SiteDeviceLastConfigCountDistinctEnum_NAME    SiteDeviceLastConfigCountDistinctEnum = "name"
    SiteDeviceLastConfigCountDistinctEnum_SITEID  SiteDeviceLastConfigCountDistinctEnum = "site_id"
    SiteDeviceLastConfigCountDistinctEnum_MAC     SiteDeviceLastConfigCountDistinctEnum = "mac"
)

// SiteDevicesCountDistinctEnum is a string enum.
type SiteDevicesCountDistinctEnum string

const (
    SiteDevicesCountDistinctEnum_MODEL          SiteDevicesCountDistinctEnum = "model"
    SiteDevicesCountDistinctEnum_VERSION        SiteDevicesCountDistinctEnum = "version"
    SiteDevicesCountDistinctEnum_MAPID          SiteDevicesCountDistinctEnum = "map_id"
    SiteDevicesCountDistinctEnum_HOSTNAME       SiteDevicesCountDistinctEnum = "hostname"
    SiteDevicesCountDistinctEnum_MXTUNNELSTATUS SiteDevicesCountDistinctEnum = "mxtunnel_status"
    SiteDevicesCountDistinctEnum_MXEDGEID       SiteDevicesCountDistinctEnum = "mxedge_id"
    SiteDevicesCountDistinctEnum_LLDPSYSTEMNAME SiteDevicesCountDistinctEnum = "lldp_system_name"
    SiteDevicesCountDistinctEnum_LLDPSYSTEMDESC SiteDevicesCountDistinctEnum = "lldp_system_desc"
    SiteDevicesCountDistinctEnum_LLDPPORTID     SiteDevicesCountDistinctEnum = "lldp_port_id"
    SiteDevicesCountDistinctEnum_LLDPMGMTADDR   SiteDevicesCountDistinctEnum = "lldp_mgmt_addr"
)

// SiteDiscoveredSwitchesCountDistinctEnum is a string enum.
type SiteDiscoveredSwitchesCountDistinctEnum string

const (
    SiteDiscoveredSwitchesCountDistinctEnum_SYSTEMNAME SiteDiscoveredSwitchesCountDistinctEnum = "system_name"
    SiteDiscoveredSwitchesCountDistinctEnum_VERSION    SiteDiscoveredSwitchesCountDistinctEnum = "version"
    SiteDiscoveredSwitchesCountDistinctEnum_MODEL      SiteDiscoveredSwitchesCountDistinctEnum = "model"
    SiteDiscoveredSwitchesCountDistinctEnum_MGMTADDR   SiteDiscoveredSwitchesCountDistinctEnum = "mgmt_addr"
)

// SiteGuestsCountDistinctEnum is a string enum.
type SiteGuestsCountDistinctEnum string

const (
    SiteGuestsCountDistinctEnum_AUTHMETHOD SiteGuestsCountDistinctEnum = "auth_method"
    SiteGuestsCountDistinctEnum_SSID       SiteGuestsCountDistinctEnum = "ssid"
    SiteGuestsCountDistinctEnum_COMPANY    SiteGuestsCountDistinctEnum = "company"
)

// SiteMxedgeEventsCountDistinctEnum is a string enum.
type SiteMxedgeEventsCountDistinctEnum string

const (
    SiteMxedgeEventsCountDistinctEnum_MXEDGEID    SiteMxedgeEventsCountDistinctEnum = "mxedge_id"
    SiteMxedgeEventsCountDistinctEnum_ENUMTYPE    SiteMxedgeEventsCountDistinctEnum = "type"
    SiteMxedgeEventsCountDistinctEnum_MXCLUSTERID SiteMxedgeEventsCountDistinctEnum = "mxcluster_id"
    SiteMxedgeEventsCountDistinctEnum_ENUMPACKAGE SiteMxedgeEventsCountDistinctEnum = "package"
)

// SiteNacClientEventsCountDistinctEnum is a string enum.
type SiteNacClientEventsCountDistinctEnum string

const (
    SiteNacClientEventsCountDistinctEnum_ENUMTYPE        SiteNacClientEventsCountDistinctEnum = "type"
    SiteNacClientEventsCountDistinctEnum_NACRULEID       SiteNacClientEventsCountDistinctEnum = "nacrule_id"
    SiteNacClientEventsCountDistinctEnum_DRYRUNNACRULEID SiteNacClientEventsCountDistinctEnum = "dryrun_nacrule_id"
    SiteNacClientEventsCountDistinctEnum_AUTHTYPE        SiteNacClientEventsCountDistinctEnum = "auth_type"
    SiteNacClientEventsCountDistinctEnum_VLAN            SiteNacClientEventsCountDistinctEnum = "vlan"
    SiteNacClientEventsCountDistinctEnum_NASVENDOR       SiteNacClientEventsCountDistinctEnum = "nas_vendor"
    SiteNacClientEventsCountDistinctEnum_USERNAME        SiteNacClientEventsCountDistinctEnum = "username"
    SiteNacClientEventsCountDistinctEnum_AP              SiteNacClientEventsCountDistinctEnum = "ap"
    SiteNacClientEventsCountDistinctEnum_MAC             SiteNacClientEventsCountDistinctEnum = "mac"
    SiteNacClientEventsCountDistinctEnum_SSID            SiteNacClientEventsCountDistinctEnum = "ssid"
)

// SiteNacClientsCountDistinctEnum is a string enum.
type SiteNacClientsCountDistinctEnum string

const (
    SiteNacClientsCountDistinctEnum_ENUMTYPE      SiteNacClientsCountDistinctEnum = "type"
    SiteNacClientsCountDistinctEnum_LASTNACRULEID SiteNacClientsCountDistinctEnum = "last_nacrule_id"
    SiteNacClientsCountDistinctEnum_AUTHTYPE      SiteNacClientsCountDistinctEnum = "auth_type"
    SiteNacClientsCountDistinctEnum_LASTVLAN      SiteNacClientsCountDistinctEnum = "last_vlan"
    SiteNacClientsCountDistinctEnum_LASTNASVENDOR SiteNacClientsCountDistinctEnum = "last_nas_vendor"
    SiteNacClientsCountDistinctEnum_LASTUSERNAME  SiteNacClientsCountDistinctEnum = "last_username"
    SiteNacClientsCountDistinctEnum_LASTAP        SiteNacClientsCountDistinctEnum = "last_ap"
    SiteNacClientsCountDistinctEnum_MAC           SiteNacClientsCountDistinctEnum = "mac"
    SiteNacClientsCountDistinctEnum_LASTSSID      SiteNacClientsCountDistinctEnum = "last_ssid"
    SiteNacClientsCountDistinctEnum_LASTSTATUS    SiteNacClientsCountDistinctEnum = "last_status"
    SiteNacClientsCountDistinctEnum_MDMCOMPLIANCE SiteNacClientsCountDistinctEnum = "mdm_compliance"
    SiteNacClientsCountDistinctEnum_MDMPROVIDER   SiteNacClientsCountDistinctEnum = "mdm_provider"
)

// SiteOtherDeviceEventsCountDistinctEnum is a string enum.
type SiteOtherDeviceEventsCountDistinctEnum string

const (
    SiteOtherDeviceEventsCountDistinctEnum_MAC      SiteOtherDeviceEventsCountDistinctEnum = "mac"
    SiteOtherDeviceEventsCountDistinctEnum_ENUMTYPE SiteOtherDeviceEventsCountDistinctEnum = "type"
    SiteOtherDeviceEventsCountDistinctEnum_VENDOR   SiteOtherDeviceEventsCountDistinctEnum = "vendor"
    SiteOtherDeviceEventsCountDistinctEnum_SITEID   SiteOtherDeviceEventsCountDistinctEnum = "site_id"
)

// SitePortsCountDistinctEnum is a string enum.
type SitePortsCountDistinctEnum string

const (
    SitePortsCountDistinctEnum_PORTID             SitePortsCountDistinctEnum = "port_id"
    SitePortsCountDistinctEnum_PORTMAC            SitePortsCountDistinctEnum = "port_mac"
    SitePortsCountDistinctEnum_FULLDUPLEX         SitePortsCountDistinctEnum = "full_duplex"
    SitePortsCountDistinctEnum_MAC                SitePortsCountDistinctEnum = "mac"
    SitePortsCountDistinctEnum_NEIGHBORMAC        SitePortsCountDistinctEnum = "neighbor_mac"
    SitePortsCountDistinctEnum_NEIGHBORPORTDESC   SitePortsCountDistinctEnum = "neighbor_port_desc"
    SitePortsCountDistinctEnum_NEIGHBORSYSTEMNAME SitePortsCountDistinctEnum = "neighbor_system_name"
    SitePortsCountDistinctEnum_POEDISABLED        SitePortsCountDistinctEnum = "poe_disabled"
    SitePortsCountDistinctEnum_POEMODE            SitePortsCountDistinctEnum = "poe_mode"
    SitePortsCountDistinctEnum_POEON              SitePortsCountDistinctEnum = "poe_on"
    SitePortsCountDistinctEnum_SPEED              SitePortsCountDistinctEnum = "speed"
    SitePortsCountDistinctEnum_UP                 SitePortsCountDistinctEnum = "up"
)

// SiteRogueEventsCountDistinctEnum is a string enum.
type SiteRogueEventsCountDistinctEnum string

const (
    SiteRogueEventsCountDistinctEnum_BSSID    SiteRogueEventsCountDistinctEnum = "bssid"
    SiteRogueEventsCountDistinctEnum_SSID     SiteRogueEventsCountDistinctEnum = "ssid"
    SiteRogueEventsCountDistinctEnum_AP       SiteRogueEventsCountDistinctEnum = "ap"
    SiteRogueEventsCountDistinctEnum_ENUMTYPE SiteRogueEventsCountDistinctEnum = "type"
)

// SiteServiceEventsCountDistinctEnum is a string enum.
type SiteServiceEventsCountDistinctEnum string

const (
    SiteServiceEventsCountDistinctEnum_ENUMTYPE SiteServiceEventsCountDistinctEnum = "type"
    SiteServiceEventsCountDistinctEnum_MAC      SiteServiceEventsCountDistinctEnum = "mac"
    SiteServiceEventsCountDistinctEnum_VPNNAME  SiteServiceEventsCountDistinctEnum = "vpn_name"
    SiteServiceEventsCountDistinctEnum_VPNPATH  SiteServiceEventsCountDistinctEnum = "vpn_path"
    SiteServiceEventsCountDistinctEnum_POLICY   SiteServiceEventsCountDistinctEnum = "policy"
    SiteServiceEventsCountDistinctEnum_PORTID   SiteServiceEventsCountDistinctEnum = "port_id"
    SiteServiceEventsCountDistinctEnum_MODEL    SiteServiceEventsCountDistinctEnum = "model"
    SiteServiceEventsCountDistinctEnum_SITEID   SiteServiceEventsCountDistinctEnum = "site_id"
)

// SiteSkyAtpEventsCountDistinctEnum is a string enum.
type SiteSkyAtpEventsCountDistinctEnum string

const (
    SiteSkyAtpEventsCountDistinctEnum_ENUMTYPE    SiteSkyAtpEventsCountDistinctEnum = "type"
    SiteSkyAtpEventsCountDistinctEnum_MAC         SiteSkyAtpEventsCountDistinctEnum = "mac"
    SiteSkyAtpEventsCountDistinctEnum_DEVICEMAC   SiteSkyAtpEventsCountDistinctEnum = "device_mac"
    SiteSkyAtpEventsCountDistinctEnum_THREATLEVEL SiteSkyAtpEventsCountDistinctEnum = "threat_level"
)

// SiteSleHistogramScopeParametersEnum is a string enum.
type SiteSleHistogramScopeParametersEnum string

const (
    SiteSleHistogramScopeParametersEnum_SITE       SiteSleHistogramScopeParametersEnum = "site"
    SiteSleHistogramScopeParametersEnum_AP         SiteSleHistogramScopeParametersEnum = "ap"
    SiteSleHistogramScopeParametersEnum_ENUMSWITCH SiteSleHistogramScopeParametersEnum = "switch"
    SiteSleHistogramScopeParametersEnum_GATEWAY    SiteSleHistogramScopeParametersEnum = "gateway"
    SiteSleHistogramScopeParametersEnum_CLIENT     SiteSleHistogramScopeParametersEnum = "client"
)

// SiteSleImpactSummaryFieldsParameterEnum is a string enum.
type SiteSleImpactSummaryFieldsParameterEnum string

const (
    SiteSleImpactSummaryFieldsParameterEnum_WLAN          SiteSleImpactSummaryFieldsParameterEnum = "wlan"
    SiteSleImpactSummaryFieldsParameterEnum_DEVICETYPE    SiteSleImpactSummaryFieldsParameterEnum = "device_type"
    SiteSleImpactSummaryFieldsParameterEnum_DEVICEOS      SiteSleImpactSummaryFieldsParameterEnum = "device_os"
    SiteSleImpactSummaryFieldsParameterEnum_BAND          SiteSleImpactSummaryFieldsParameterEnum = "band"
    SiteSleImpactSummaryFieldsParameterEnum_AP            SiteSleImpactSummaryFieldsParameterEnum = "ap"
    SiteSleImpactSummaryFieldsParameterEnum_SERVER        SiteSleImpactSummaryFieldsParameterEnum = "server"
    SiteSleImpactSummaryFieldsParameterEnum_MXEDGE        SiteSleImpactSummaryFieldsParameterEnum = "mxedge"
    SiteSleImpactSummaryFieldsParameterEnum_ENUMSWITCH    SiteSleImpactSummaryFieldsParameterEnum = "switch"
    SiteSleImpactSummaryFieldsParameterEnum_CLIENT        SiteSleImpactSummaryFieldsParameterEnum = "client"
    SiteSleImpactSummaryFieldsParameterEnum_VLAN          SiteSleImpactSummaryFieldsParameterEnum = "vlan"
    SiteSleImpactSummaryFieldsParameterEnum_ENUMINTERFACE SiteSleImpactSummaryFieldsParameterEnum = "interface"
    SiteSleImpactSummaryFieldsParameterEnum_CHASSIS       SiteSleImpactSummaryFieldsParameterEnum = "chassis"
    SiteSleImpactSummaryFieldsParameterEnum_GATEWAY       SiteSleImpactSummaryFieldsParameterEnum = "gateway"
    SiteSleImpactSummaryFieldsParameterEnum_PEERPATH      SiteSleImpactSummaryFieldsParameterEnum = "peer_path"
    SiteSleImpactSummaryFieldsParameterEnum_GATEWAYZONES  SiteSleImpactSummaryFieldsParameterEnum = "gateway_zones"
)

// SiteSleImpactSummaryScopeParametersEnum is a string enum.
type SiteSleImpactSummaryScopeParametersEnum string

const (
    SiteSleImpactSummaryScopeParametersEnum_SITE       SiteSleImpactSummaryScopeParametersEnum = "site"
    SiteSleImpactSummaryScopeParametersEnum_AP         SiteSleImpactSummaryScopeParametersEnum = "ap"
    SiteSleImpactSummaryScopeParametersEnum_ENUMSWITCH SiteSleImpactSummaryScopeParametersEnum = "switch"
    SiteSleImpactSummaryScopeParametersEnum_GATEWAY    SiteSleImpactSummaryScopeParametersEnum = "gateway"
    SiteSleImpactSummaryScopeParametersEnum_CLIENT     SiteSleImpactSummaryScopeParametersEnum = "client"
)

// SiteSleImpactedApsScopeParametersEnum is a string enum.
type SiteSleImpactedApsScopeParametersEnum string

const (
    SiteSleImpactedApsScopeParametersEnum_SITE SiteSleImpactedApsScopeParametersEnum = "site"
)

// SiteSleImpactedChassisScopeParametersEnum is a string enum.
type SiteSleImpactedChassisScopeParametersEnum string

const (
    SiteSleImpactedChassisScopeParametersEnum_SITE       SiteSleImpactedChassisScopeParametersEnum = "site"
    SiteSleImpactedChassisScopeParametersEnum_ENUMSWITCH SiteSleImpactedChassisScopeParametersEnum = "switch"
    SiteSleImpactedChassisScopeParametersEnum_GATEWAY    SiteSleImpactedChassisScopeParametersEnum = "gateway"
)

// SiteSleImpactedClientsScopeParametersEnum is a string enum.
type SiteSleImpactedClientsScopeParametersEnum string

const (
    SiteSleImpactedClientsScopeParametersEnum_SITE       SiteSleImpactedClientsScopeParametersEnum = "site"
    SiteSleImpactedClientsScopeParametersEnum_ENUMSWITCH SiteSleImpactedClientsScopeParametersEnum = "switch"
    SiteSleImpactedClientsScopeParametersEnum_GATEWAY    SiteSleImpactedClientsScopeParametersEnum = "gateway"
)

// SiteSleImpactedGatewaysScopeParametersEnum is a string enum.
type SiteSleImpactedGatewaysScopeParametersEnum string

const (
    SiteSleImpactedGatewaysScopeParametersEnum_SITE SiteSleImpactedGatewaysScopeParametersEnum = "site"
)

// SiteSleImpactedInterfacesScopeParametersEnum is a string enum.
type SiteSleImpactedInterfacesScopeParametersEnum string

const (
    SiteSleImpactedInterfacesScopeParametersEnum_SITE       SiteSleImpactedInterfacesScopeParametersEnum = "site"
    SiteSleImpactedInterfacesScopeParametersEnum_ENUMSWITCH SiteSleImpactedInterfacesScopeParametersEnum = "switch"
    SiteSleImpactedInterfacesScopeParametersEnum_GATEWAY    SiteSleImpactedInterfacesScopeParametersEnum = "gateway"
)

// SiteSleImpactedSwitchesScopeParametersEnum is a string enum.
type SiteSleImpactedSwitchesScopeParametersEnum string

const (
    SiteSleImpactedSwitchesScopeParametersEnum_SITE SiteSleImpactedSwitchesScopeParametersEnum = "site"
)

// SiteSleImpactedUsersScopeParameterEnum is a string enum.
type SiteSleImpactedUsersScopeParameterEnum string

const (
    SiteSleImpactedUsersScopeParameterEnum_SITE SiteSleImpactedUsersScopeParameterEnum = "site"
    SiteSleImpactedUsersScopeParameterEnum_AP   SiteSleImpactedUsersScopeParameterEnum = "ap"
)

// SiteSleMetricClassifiersScopeParametersEnum is a string enum.
type SiteSleMetricClassifiersScopeParametersEnum string

const (
    SiteSleMetricClassifiersScopeParametersEnum_SITE       SiteSleMetricClassifiersScopeParametersEnum = "site"
    SiteSleMetricClassifiersScopeParametersEnum_AP         SiteSleMetricClassifiersScopeParametersEnum = "ap"
    SiteSleMetricClassifiersScopeParametersEnum_ENUMSWITCH SiteSleMetricClassifiersScopeParametersEnum = "switch"
    SiteSleMetricClassifiersScopeParametersEnum_GATEWAY    SiteSleMetricClassifiersScopeParametersEnum = "gateway"
    SiteSleMetricClassifiersScopeParametersEnum_CLIENT     SiteSleMetricClassifiersScopeParametersEnum = "client"
)

// SiteSleMetricSummaryScopeParametersEnum is a string enum.
type SiteSleMetricSummaryScopeParametersEnum string

const (
    SiteSleMetricSummaryScopeParametersEnum_SITE       SiteSleMetricSummaryScopeParametersEnum = "site"
    SiteSleMetricSummaryScopeParametersEnum_AP         SiteSleMetricSummaryScopeParametersEnum = "ap"
    SiteSleMetricSummaryScopeParametersEnum_ENUMSWITCH SiteSleMetricSummaryScopeParametersEnum = "switch"
    SiteSleMetricSummaryScopeParametersEnum_GATEWAY    SiteSleMetricSummaryScopeParametersEnum = "gateway"
    SiteSleMetricSummaryScopeParametersEnum_CLIENT     SiteSleMetricSummaryScopeParametersEnum = "client"
)

// SiteSleMetricsScopeParametersEnum is a string enum.
type SiteSleMetricsScopeParametersEnum string

const (
    SiteSleMetricsScopeParametersEnum_SITE       SiteSleMetricsScopeParametersEnum = "site"
    SiteSleMetricsScopeParametersEnum_AP         SiteSleMetricsScopeParametersEnum = "ap"
    SiteSleMetricsScopeParametersEnum_ENUMSWITCH SiteSleMetricsScopeParametersEnum = "switch"
    SiteSleMetricsScopeParametersEnum_GATEWAY    SiteSleMetricsScopeParametersEnum = "gateway"
    SiteSleMetricsScopeParametersEnum_CLIENT     SiteSleMetricsScopeParametersEnum = "client"
)

// SiteSleScopeEnum is a string enum.
type SiteSleScopeEnum string

const (
    SiteSleScopeEnum_SITE       SiteSleScopeEnum = "site"
    SiteSleScopeEnum_ENUMSWITCH SiteSleScopeEnum = "switch"
    SiteSleScopeEnum_GATEWAY    SiteSleScopeEnum = "gateway"
)

// SiteSleThresholdScopeParameterEnum is a string enum.
type SiteSleThresholdScopeParameterEnum string

const (
    SiteSleThresholdScopeParameterEnum_SITE       SiteSleThresholdScopeParameterEnum = "site"
    SiteSleThresholdScopeParameterEnum_AP         SiteSleThresholdScopeParameterEnum = "ap"
    SiteSleThresholdScopeParameterEnum_ENUMSWITCH SiteSleThresholdScopeParameterEnum = "switch"
    SiteSleThresholdScopeParameterEnum_GATEWAY    SiteSleThresholdScopeParameterEnum = "gateway"
    SiteSleThresholdScopeParameterEnum_CLIENT     SiteSleThresholdScopeParameterEnum = "client"
)

// SiteSwitchPortsCountDistinctEnum is a string enum.
type SiteSwitchPortsCountDistinctEnum string

const (
    SiteSwitchPortsCountDistinctEnum_PORTID             SiteSwitchPortsCountDistinctEnum = "port_id"
    SiteSwitchPortsCountDistinctEnum_PORTMAC            SiteSwitchPortsCountDistinctEnum = "port_mac"
    SiteSwitchPortsCountDistinctEnum_FULLDUPLEX         SiteSwitchPortsCountDistinctEnum = "full_duplex"
    SiteSwitchPortsCountDistinctEnum_MAC                SiteSwitchPortsCountDistinctEnum = "mac"
    SiteSwitchPortsCountDistinctEnum_NEIGHBORMAC        SiteSwitchPortsCountDistinctEnum = "neighbor_mac"
    SiteSwitchPortsCountDistinctEnum_NEIGHBORPORTDESC   SiteSwitchPortsCountDistinctEnum = "neighbor_port_desc"
    SiteSwitchPortsCountDistinctEnum_NEIGHBORSYSTEMNAME SiteSwitchPortsCountDistinctEnum = "neighbor_system_name"
    SiteSwitchPortsCountDistinctEnum_POEDISABLED        SiteSwitchPortsCountDistinctEnum = "poe_disabled"
    SiteSwitchPortsCountDistinctEnum_POEMODE            SiteSwitchPortsCountDistinctEnum = "poe_mode"
    SiteSwitchPortsCountDistinctEnum_POEON              SiteSwitchPortsCountDistinctEnum = "poe_on"
    SiteSwitchPortsCountDistinctEnum_SPEED              SiteSwitchPortsCountDistinctEnum = "speed"
    SiteSwitchPortsCountDistinctEnum_UP                 SiteSwitchPortsCountDistinctEnum = "up"
)

// SiteSystemEventsCountDistinctEnum is a string enum.
type SiteSystemEventsCountDistinctEnum string

const (
    SiteSystemEventsCountDistinctEnum_ENUMTYPE SiteSystemEventsCountDistinctEnum = "type"
)

// SiteWanClientEventsDistinctEnum is a string enum.
type SiteWanClientEventsDistinctEnum string

const (
    SiteWanClientEventsDistinctEnum_ENUMTYPE SiteWanClientEventsDistinctEnum = "type"
    SiteWanClientEventsDistinctEnum_HOSTNAME SiteWanClientEventsDistinctEnum = "hostname"
    SiteWanClientEventsDistinctEnum_IP       SiteWanClientEventsDistinctEnum = "ip"
    SiteWanClientEventsDistinctEnum_MFG      SiteWanClientEventsDistinctEnum = "mfg"
    SiteWanClientEventsDistinctEnum_MAC      SiteWanClientEventsDistinctEnum = "mac"
)

// SiteWanClientsCountDistinctEnum is a string enum.
type SiteWanClientsCountDistinctEnum string

const (
    SiteWanClientsCountDistinctEnum_HOSTNAME SiteWanClientsCountDistinctEnum = "hostname"
    SiteWanClientsCountDistinctEnum_IP       SiteWanClientsCountDistinctEnum = "ip"
    SiteWanClientsCountDistinctEnum_MFG      SiteWanClientsCountDistinctEnum = "mfg"
    SiteWanClientsCountDistinctEnum_MAC      SiteWanClientsCountDistinctEnum = "mac"
)

// SiteWiredClientsCountDistinctEnum is a string enum.
type SiteWiredClientsCountDistinctEnum string

const (
    SiteWiredClientsCountDistinctEnum_PORTID SiteWiredClientsCountDistinctEnum = "port_id"
    SiteWiredClientsCountDistinctEnum_VLAN   SiteWiredClientsCountDistinctEnum = "vlan"
    SiteWiredClientsCountDistinctEnum_MAC    SiteWiredClientsCountDistinctEnum = "mac"
)

// SiteZoneCountDistinctEnum is a string enum.
type SiteZoneCountDistinctEnum string

const (
    SiteZoneCountDistinctEnum_USERTYPE SiteZoneCountDistinctEnum = "user_type"
    SiteZoneCountDistinctEnum_USER     SiteZoneCountDistinctEnum = "user"
    SiteZoneCountDistinctEnum_SCOPEID  SiteZoneCountDistinctEnum = "scope_id"
    SiteZoneCountDistinctEnum_SCOPE    SiteZoneCountDistinctEnum = "scope"
)

// SleSummaryScopeEnum is a string enum.
type SleSummaryScopeEnum string

const (
    SleSummaryScopeEnum_SITE       SleSummaryScopeEnum = "site"
    SleSummaryScopeEnum_AP         SleSummaryScopeEnum = "ap"
    SleSummaryScopeEnum_ENUMSWITCH SleSummaryScopeEnum = "switch"
    SleSummaryScopeEnum_GATEWAY    SleSummaryScopeEnum = "gateway"
    SleSummaryScopeEnum_CLIENT     SleSummaryScopeEnum = "client"
)

// StatDeviceStatusFilterEnum is a string enum.
type StatDeviceStatusFilterEnum string

const (
    StatDeviceStatusFilterEnum_ALL          StatDeviceStatusFilterEnum = "all"
    StatDeviceStatusFilterEnum_CONNECTED    StatDeviceStatusFilterEnum = "connected"
    StatDeviceStatusFilterEnum_DISCONNECTED StatDeviceStatusFilterEnum = "disconnected"
)

// SwitchMetricScopeEnum is a string enum.
type SwitchMetricScopeEnum string

const (
    SwitchMetricScopeEnum_SITE       SwitchMetricScopeEnum = "site"
    SwitchMetricScopeEnum_ENUMSWITCH SwitchMetricScopeEnum = "switch"
)

// SwitchMetricTypeEnum is a string enum.
type SwitchMetricTypeEnum string

const (
    SwitchMetricTypeEnum_ACTIVEPORTSSUMMARY SwitchMetricTypeEnum = "active_ports_summary"
)

// SynthetictestTypeEnum is a string enum.
type SynthetictestTypeEnum string

const (
    SynthetictestTypeEnum_DNS       SynthetictestTypeEnum = "dns"
    SynthetictestTypeEnum_ARP       SynthetictestTypeEnum = "arp"
    SynthetictestTypeEnum_DHCP      SynthetictestTypeEnum = "dhcp"
    SynthetictestTypeEnum_CURL      SynthetictestTypeEnum = "curl"
    SynthetictestTypeEnum_RADIUS    SynthetictestTypeEnum = "radius"
    SynthetictestTypeEnum_SPEEDTEST SynthetictestTypeEnum = "speedtest"
    SynthetictestTypeEnum_DHCP6     SynthetictestTypeEnum = "dhcp6"
)

// VisitsScopeEnum is a string enum.
type VisitsScopeEnum string

const (
    VisitsScopeEnum_SITE     VisitsScopeEnum = "site"
    VisitsScopeEnum_ENUMMAP  VisitsScopeEnum = "map"
    VisitsScopeEnum_ZONE     VisitsScopeEnum = "zone"
    VisitsScopeEnum_RSSIZONE VisitsScopeEnum = "rssizone"
)

// WanUsagesCountDisctinctEnum is a string enum.
type WanUsagesCountDisctinctEnum string

const (
    WanUsagesCountDisctinctEnum_MAC        WanUsagesCountDisctinctEnum = "mac"
    WanUsagesCountDisctinctEnum_PEERMAC    WanUsagesCountDisctinctEnum = "peer_mac"
    WanUsagesCountDisctinctEnum_PORTID     WanUsagesCountDisctinctEnum = "port_id"
    WanUsagesCountDisctinctEnum_PEERPORTID WanUsagesCountDisctinctEnum = "peer_port_id"
    WanUsagesCountDisctinctEnum_POLICY     WanUsagesCountDisctinctEnum = "policy"
    WanUsagesCountDisctinctEnum_TENANT     WanUsagesCountDisctinctEnum = "tenant"
    WanUsagesCountDisctinctEnum_PATHTYPE   WanUsagesCountDisctinctEnum = "path_type"
)

// ZoneScopeEnum is a string enum.
type ZoneScopeEnum string

const (
    ZoneScopeEnum_SITE     ZoneScopeEnum = "site"
    ZoneScopeEnum_ENUMMAP  ZoneScopeEnum = "map"
    ZoneScopeEnum_ZONE     ZoneScopeEnum = "zone"
    ZoneScopeEnum_RSSIZONE ZoneScopeEnum = "rssizone"
)

// ApClientBridgeAuthTypeEnum is a string enum.
// wpa2-AES/CCMPp is assumed when `type`==`psk`
type ApClientBridgeAuthTypeEnum string

const (
    ApClientBridgeAuthTypeEnum_OPEN ApClientBridgeAuthTypeEnum = "open"
    ApClientBridgeAuthTypeEnum_PSK  ApClientBridgeAuthTypeEnum = "psk"
)

// AppProbingCustomAppProtocolEnum is a string enum.
type AppProbingCustomAppProtocolEnum string

const (
    AppProbingCustomAppProtocolEnum_HTTP AppProbingCustomAppProtocolEnum = "http"
    AppProbingCustomAppProtocolEnum_UDP  AppProbingCustomAppProtocolEnum = "udp"
)

// AutoOrientationStateEnum is a string enum.
// The state of auto orient for a given map derived from an Enum
type AutoOrientationStateEnum string

const (
    AutoOrientationStateEnum_ENUMNOTSTARTED AutoOrientationStateEnum = "Not Started"
    AutoOrientationStateEnum_ENQUEUED       AutoOrientationStateEnum = "Enqueued"
    AutoOrientationStateEnum_ORIENTED       AutoOrientationStateEnum = "Oriented"
)

// AutoPlacementInfoStatusEnum is a string enum.
// the status of autoplacement for a given map
type AutoPlacementInfoStatusEnum string

const (
    AutoPlacementInfoStatusEnum_PENDING    AutoPlacementInfoStatusEnum = "pending"
    AutoPlacementInfoStatusEnum_INPROGRESS AutoPlacementInfoStatusEnum = "inprogress"
    AutoPlacementInfoStatusEnum_DONE       AutoPlacementInfoStatusEnum = "done"
    AutoPlacementInfoStatusEnum_ENUMERROR  AutoPlacementInfoStatusEnum = "error"
)

// BeaconTypeEnum is a string enum.
type BeaconTypeEnum string

const (
    BeaconTypeEnum_IBEACON      BeaconTypeEnum = "ibeacon"
    BeaconTypeEnum_EDDYSTONEUID BeaconTypeEnum = "eddystone-uid"
    BeaconTypeEnum_EDDYSTONEURL BeaconTypeEnum = "eddystone-url"
)

// CaptureGatewayFormatEnum is a string enum.
type CaptureGatewayFormatEnum string

const (
    CaptureGatewayFormatEnum_STREAM CaptureGatewayFormatEnum = "stream"
)

// CaptureRadiotapBandEnum is a string enum.
type CaptureRadiotapBandEnum string

const (
    CaptureRadiotapBandEnum_ENUM24   CaptureRadiotapBandEnum = "24"
    CaptureRadiotapBandEnum_ENUM5    CaptureRadiotapBandEnum = "5"
    CaptureRadiotapBandEnum_ENUM6    CaptureRadiotapBandEnum = "6"
    CaptureRadiotapBandEnum_ENUM2456 CaptureRadiotapBandEnum = "24,5,6"
)

// CaptureRadiotapFormatEnum is a string enum.
type CaptureRadiotapFormatEnum string

const (
    CaptureRadiotapFormatEnum_PCAP   CaptureRadiotapFormatEnum = "pcap"
    CaptureRadiotapFormatEnum_STREAM CaptureRadiotapFormatEnum = "stream"
)

// CaptureRadiotapwiredBandEnum is a string enum.
// only used for radiotap
type CaptureRadiotapwiredBandEnum string

const (
    CaptureRadiotapwiredBandEnum_ENUM24   CaptureRadiotapwiredBandEnum = "24"
    CaptureRadiotapwiredBandEnum_ENUM5    CaptureRadiotapwiredBandEnum = "5"
    CaptureRadiotapwiredBandEnum_ENUM6    CaptureRadiotapwiredBandEnum = "6"
    CaptureRadiotapwiredBandEnum_ENUM2456 CaptureRadiotapwiredBandEnum = "24,5,6"
)

// CaptureRadiotapwiredFormatEnum is a string enum.
type CaptureRadiotapwiredFormatEnum string

const (
    CaptureRadiotapwiredFormatEnum_PCAP   CaptureRadiotapwiredFormatEnum = "pcap"
    CaptureRadiotapwiredFormatEnum_STREAM CaptureRadiotapwiredFormatEnum = "stream"
)

// CaptureScanApsBandEnum is a string enum.
// Only Single value allowed
type CaptureScanApsBandEnum string

const (
    CaptureScanApsBandEnum_ENUM24 CaptureScanApsBandEnum = "24"
    CaptureScanApsBandEnum_ENUM5  CaptureScanApsBandEnum = "5"
    CaptureScanApsBandEnum_ENUM6  CaptureScanApsBandEnum = "6"
)

// CaptureScanBandEnum is a string enum.
// Only Single value allowed, default value gets applied when user provides wrong values
type CaptureScanBandEnum string

const (
    CaptureScanBandEnum_ENUM24 CaptureScanBandEnum = "24"
    CaptureScanBandEnum_ENUM5  CaptureScanBandEnum = "5"
    CaptureScanBandEnum_ENUM6  CaptureScanBandEnum = "6"
)

// CaptureScanFormatEnum is a string enum.
type CaptureScanFormatEnum string

const (
    CaptureScanFormatEnum_PCAP   CaptureScanFormatEnum = "pcap"
    CaptureScanFormatEnum_STREAM CaptureScanFormatEnum = "stream"
)

// CaptureSwitchFormatEnum is a string enum.
type CaptureSwitchFormatEnum string

const (
    CaptureSwitchFormatEnum_STREAM CaptureSwitchFormatEnum = "stream"
)

// CaptureWiredFormatEnum is a string enum.
// pcap format
type CaptureWiredFormatEnum string

const (
    CaptureWiredFormatEnum_PCAP   CaptureWiredFormatEnum = "pcap"
    CaptureWiredFormatEnum_STREAM CaptureWiredFormatEnum = "stream"
)

// CaptureWirelessBandEnum is a string enum.
type CaptureWirelessBandEnum string

const (
    CaptureWirelessBandEnum_ENUM24   CaptureWirelessBandEnum = "24"
    CaptureWirelessBandEnum_ENUM5    CaptureWirelessBandEnum = "5"
    CaptureWirelessBandEnum_ENUM6    CaptureWirelessBandEnum = "6"
    CaptureWirelessBandEnum_ENUM2456 CaptureWirelessBandEnum = "24,5,6"
)

// CaptureWirelessFormatEnum is a string enum.
// pcap format
type CaptureWirelessFormatEnum string

const (
    CaptureWirelessFormatEnum_PCAP   CaptureWirelessFormatEnum = "pcap"
    CaptureWirelessFormatEnum_STREAM CaptureWirelessFormatEnum = "stream"
)

// EventFastroamTypeEnum is a string enum.
type EventFastroamTypeEnum string

const (
    EventFastroamTypeEnum_SUCCESS  EventFastroamTypeEnum = "success"
    EventFastroamTypeEnum_FAIL     EventFastroamTypeEnum = "fail"
    EventFastroamTypeEnum_NONE     EventFastroamTypeEnum = "none"
    EventFastroamTypeEnum_POOR     EventFastroamTypeEnum = "poor"
    EventFastroamTypeEnum_PINGPONG EventFastroamTypeEnum = "pingpong"
    EventFastroamTypeEnum_SLOW     EventFastroamTypeEnum = "slow"
)

// EvpnOptionsRoutedAtEnum is a string enum.
// optional, where virtual-gateway should reside
type EvpnOptionsRoutedAtEnum string

const (
    EvpnOptionsRoutedAtEnum_EDGE         EvpnOptionsRoutedAtEnum = "edge"
    EvpnOptionsRoutedAtEnum_CORE         EvpnOptionsRoutedAtEnum = "core"
    EvpnOptionsRoutedAtEnum_DISTRIBUTION EvpnOptionsRoutedAtEnum = "distribution"
)

// GatewayPortUsage1Enum is a string enum.
// port usage name. 
// If EVPN is used, use `evpn_uplink`or `evpn_downlink`
type GatewayPortUsage1Enum string

const (
    GatewayPortUsage1Enum_LAN       GatewayPortUsage1Enum = "lan"
    GatewayPortUsage1Enum_WAN       GatewayPortUsage1Enum = "wan"
    GatewayPortUsage1Enum_HADATA    GatewayPortUsage1Enum = "ha_data"
    GatewayPortUsage1Enum_HACONTROL GatewayPortUsage1Enum = "ha_control"
)

// OspfAreasNetworkAuthTypeEnum is a string enum.
// auth type
type OspfAreasNetworkAuthTypeEnum string

const (
    OspfAreasNetworkAuthTypeEnum_NONE     OspfAreasNetworkAuthTypeEnum = "none"
    OspfAreasNetworkAuthTypeEnum_MD5      OspfAreasNetworkAuthTypeEnum = "md5"
    OspfAreasNetworkAuthTypeEnum_PASSWORD OspfAreasNetworkAuthTypeEnum = "password"
)

// OspfAreasNetworkInterfaceTypeEnum is a string enum.
// interface type (nbma = non-broadcast multi-access)
type OspfAreasNetworkInterfaceTypeEnum string

const (
    OspfAreasNetworkInterfaceTypeEnum_BROADCAST OspfAreasNetworkInterfaceTypeEnum = "broadcast"
    OspfAreasNetworkInterfaceTypeEnum_NBMA      OspfAreasNetworkInterfaceTypeEnum = "nbma"
    OspfAreasNetworkInterfaceTypeEnum_P2P       OspfAreasNetworkInterfaceTypeEnum = "p2p"
    OspfAreasNetworkInterfaceTypeEnum_P2MP      OspfAreasNetworkInterfaceTypeEnum = "p2mp"
)

// OspfAreasTypeEnum is a string enum.
// OSPF type, default (default) / stub / nssa
type OspfAreasTypeEnum string

const (
    OspfAreasTypeEnum_ENUMDEFAULT OspfAreasTypeEnum = "default"
    OspfAreasTypeEnum_STUB        OspfAreasTypeEnum = "stub"
    OspfAreasTypeEnum_NSSA        OspfAreasTypeEnum = "nssa"
)

// ResolutionEnum is a string enum.
type ResolutionEnum string

const (
    ResolutionEnum_ENUMDEFAULT ResolutionEnum = "default"
    ResolutionEnum_FINE        ResolutionEnum = "fine"
)

// ResponseDeviceSnapshotStatusEnum is a string enum.
type ResponseDeviceSnapshotStatusEnum string

const (
    ResponseDeviceSnapshotStatusEnum_STARTING   ResponseDeviceSnapshotStatusEnum = "starting"
    ResponseDeviceSnapshotStatusEnum_INPROGRESS ResponseDeviceSnapshotStatusEnum = "inprogress"
    ResponseDeviceSnapshotStatusEnum_SUCCESS    ResponseDeviceSnapshotStatusEnum = "success"
    ResponseDeviceSnapshotStatusEnum_ENUMERROR  ResponseDeviceSnapshotStatusEnum = "error"
    ResponseDeviceSnapshotStatusEnum_SCHEDULED  ResponseDeviceSnapshotStatusEnum = "scheduled"
)

// RogueTypeEnum is a string enum.
type RogueTypeEnum string

const (
    RogueTypeEnum_HONEYPOT RogueTypeEnum = "honeypot"
    RogueTypeEnum_LAN      RogueTypeEnum = "lan"
    RogueTypeEnum_OTHERS   RogueTypeEnum = "others"
    RogueTypeEnum_SPOOF    RogueTypeEnum = "spoof"
)

// RrmEventPreBandwidthEnum is a int enum.
// (previously) channel width for the band , 0 means no previously available
type RrmEventPreBandwidthEnum int

const (
    RrmEventPreBandwidthEnum_ENUM0   RrmEventPreBandwidthEnum = 0
    RrmEventPreBandwidthEnum_ENUM20  RrmEventPreBandwidthEnum = 20
    RrmEventPreBandwidthEnum_ENUM40  RrmEventPreBandwidthEnum = 40
    RrmEventPreBandwidthEnum_ENUM80  RrmEventPreBandwidthEnum = 80
    RrmEventPreBandwidthEnum_ENUM160 RrmEventPreBandwidthEnum = 160
)

// RrmEventTypeEnum is a string enum.
// schedule-site_rrm / triggered-site_rrm / interference-ap-co-channel / rrm-radar
type RrmEventTypeEnum string

const (
    RrmEventTypeEnum_TRIGGEREDSITERRM        RrmEventTypeEnum = "triggered-site_rrm"
    RrmEventTypeEnum_INTERFERENCEAPCOCHANNEL RrmEventTypeEnum = "interference-ap-co-channel"
    RrmEventTypeEnum_RRMRADAR                RrmEventTypeEnum = "rrm-radar"
    RrmEventTypeEnum_SCHEDULEDSITERRM        RrmEventTypeEnum = "scheduled-site_rrm"
    RrmEventTypeEnum_INTERFERENCEAPNONWIFI   RrmEventTypeEnum = "interference-ap-non-wifi"
    RrmEventTypeEnum_RADARDETECTED           RrmEventTypeEnum = "radar-detected"
    RrmEventTypeEnum_NEIGHBORAPDOWN          RrmEventTypeEnum = "neighbor-ap-down"
    RrmEventTypeEnum_NEIGHBORAPRECOVERED     RrmEventTypeEnum = "neighbor-ap-recovered"
)

// RrmStatusEnum is a string enum.
type RrmStatusEnum string

const (
    RrmStatusEnum_UNKNOWN  RrmStatusEnum = "unknown"
    RrmStatusEnum_UPDATING RrmStatusEnum = "updating"
    RrmStatusEnum_READY    RrmStatusEnum = "ready"
)

// SiteAutoUpgradeVersionEnum is a string enum.
// desired version
type SiteAutoUpgradeVersionEnum string

const (
    SiteAutoUpgradeVersionEnum_BETA   SiteAutoUpgradeVersionEnum = "beta"
    SiteAutoUpgradeVersionEnum_STABLE SiteAutoUpgradeVersionEnum = "stable"
    SiteAutoUpgradeVersionEnum_CUSTOM SiteAutoUpgradeVersionEnum = "custom"
)

// SiteMxtunnelProtocolEnum is a string enum.
type SiteMxtunnelProtocolEnum string

const (
    SiteMxtunnelProtocolEnum_UDP SiteMxtunnelProtocolEnum = "udp"
    SiteMxtunnelProtocolEnum_IP  SiteMxtunnelProtocolEnum = "ip"
)

// SiteWifiProxyArpEnum is a string enum.
// default / enabled / disabled
type SiteWifiProxyArpEnum string

const (
    SiteWifiProxyArpEnum_ENUMDEFAULT SiteWifiProxyArpEnum = "default"
    SiteWifiProxyArpEnum_ENABLED     SiteWifiProxyArpEnum = "enabled"
    SiteWifiProxyArpEnum_DISABLED    SiteWifiProxyArpEnum = "disabled"
)

// UpgradeInfoStatusEnum is a string enum.
type UpgradeInfoStatusEnum string

const (
    UpgradeInfoStatusEnum_STARTING   UpgradeInfoStatusEnum = "starting"
    UpgradeInfoStatusEnum_INPROGRESS UpgradeInfoStatusEnum = "inprogress"
    UpgradeInfoStatusEnum_SUCCESS    UpgradeInfoStatusEnum = "success"
    UpgradeInfoStatusEnum_ENUMERROR  UpgradeInfoStatusEnum = "error"
    UpgradeInfoStatusEnum_SCHEDULED  UpgradeInfoStatusEnum = "scheduled"
)

// UpgradeSiteDevicesRrmNodeOrderEnum is a string enum.
// Used in rrm to determine whether to start upgrade from fringe or center AP’s
type UpgradeSiteDevicesRrmNodeOrderEnum string

const (
    UpgradeSiteDevicesRrmNodeOrderEnum_FRINGETOCENTER UpgradeSiteDevicesRrmNodeOrderEnum = "fringe_to_center"
    UpgradeSiteDevicesRrmNodeOrderEnum_CENTERTOFRINGE UpgradeSiteDevicesRrmNodeOrderEnum = "center_to_fringe"
)

// UseAutoApValuesForEnum is a string enum.
// The selector to choose auto placement or auto orientation.
type UseAutoApValuesForEnum string

const (
    UseAutoApValuesForEnum_PLACEMENT   UseAutoApValuesForEnum = "placement"
    UseAutoApValuesForEnum_ORIENTATION UseAutoApValuesForEnum = "orientation"
)

// UtilsClearBgpTypeEnum is a string enum.
type UtilsClearBgpTypeEnum string

const (
    UtilsClearBgpTypeEnum_HARD UtilsClearBgpTypeEnum = "hard"
    UtilsClearBgpTypeEnum_SOFT UtilsClearBgpTypeEnum = "soft"
    UtilsClearBgpTypeEnum_IN   UtilsClearBgpTypeEnum = "in"
    UtilsClearBgpTypeEnum_OUT  UtilsClearBgpTypeEnum = "out"
)

// UtilsSendSupportLogsInfoEnum is a string enum.
// optional choice: process, outbound-ssh, and full (default)
type UtilsSendSupportLogsInfoEnum string

const (
    UtilsSendSupportLogsInfoEnum_PROCESS     UtilsSendSupportLogsInfoEnum = "process"
    UtilsSendSupportLogsInfoEnum_OUTBOUNDSSH UtilsSendSupportLogsInfoEnum = "outbound-ssh"
    UtilsSendSupportLogsInfoEnum_MESSAGES    UtilsSendSupportLogsInfoEnum = "messages"
    UtilsSendSupportLogsInfoEnum_CODEDUMPS   UtilsSendSupportLogsInfoEnum = "code-dumps"
    UtilsSendSupportLogsInfoEnum_FULL        UtilsSendSupportLogsInfoEnum = "full"
    UtilsSendSupportLogsInfoEnum_VARLOGS     UtilsSendSupportLogsInfoEnum = "var-logs"
    UtilsSendSupportLogsInfoEnum_JMALOGS     UtilsSendSupportLogsInfoEnum = "jma-logs"
)

// UtilsShowRouteProtocolEnum is a string enum.
type UtilsShowRouteProtocolEnum string

const (
    UtilsShowRouteProtocolEnum_BGP    UtilsShowRouteProtocolEnum = "bgp"
    UtilsShowRouteProtocolEnum_ANY    UtilsShowRouteProtocolEnum = "any"
    UtilsShowRouteProtocolEnum_OSPF   UtilsShowRouteProtocolEnum = "ospf"
    UtilsShowRouteProtocolEnum_STATIC UtilsShowRouteProtocolEnum = "static"
    UtilsShowRouteProtocolEnum_DIRECT UtilsShowRouteProtocolEnum = "direct"
    UtilsShowRouteProtocolEnum_EVPN   UtilsShowRouteProtocolEnum = "evpn"
)

// UtilsTracerouteProtocolEnum is a string enum.
type UtilsTracerouteProtocolEnum string

const (
    UtilsTracerouteProtocolEnum_UDP UtilsTracerouteProtocolEnum = "udp"
)

// VirtualChassisPortOperationEnum is a string enum.
type VirtualChassisPortOperationEnum string

const (
    VirtualChassisPortOperationEnum_SET    VirtualChassisPortOperationEnum = "set"
    VirtualChassisPortOperationEnum_DELETE VirtualChassisPortOperationEnum = "delete"
)

// VrrpGroupAuthTypeEnum is a string enum.
type VrrpGroupAuthTypeEnum string

const (
    VrrpGroupAuthTypeEnum_MD5    VrrpGroupAuthTypeEnum = "md5"
    VrrpGroupAuthTypeEnum_SIMPLE VrrpGroupAuthTypeEnum = "simple"
)

// WxruleStatsActionEnum is a string enum.
type WxruleStatsActionEnum string

const (
    WxruleStatsActionEnum_ALLOW WxruleStatsActionEnum = "allow"
    WxruleStatsActionEnum_BLOCK WxruleStatsActionEnum = "block"
)

// ZoneTypeEnum is a string enum.
type ZoneTypeEnum string

const (
    ZoneTypeEnum_ZONES     ZoneTypeEnum = "zones"
    ZoneTypeEnum_RSSIZONES ZoneTypeEnum = "rssizones"
)
