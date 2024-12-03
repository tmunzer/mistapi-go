package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
)

// Client is an interface representing the main client for accessing configuration and controllers.
type ClientInterface interface {
    Configuration() Configuration
    CloneWithConfiguration(options ...ConfigurationOptions) ClientInterface
    Admins() *Admins
    AdminsLogin() *AdminsLogin
    AdminsLogout() *AdminsLogout
    AdminsRecoverPassword() *AdminsRecoverPassword
    AdminsLookup() *AdminsLookup
    AdminsLoginOAuth2() *AdminsLoginOAuth2
    Installer() *Installer
    MSPs() *MSPs
    MSPsAdmins() *MSPsAdmins
    OrgsSecIntelProfiles() *OrgsSecIntelProfiles
    SitesSecIntelProfiles() *SitesSecIntelProfiles
    MSPsInventory() *MSPsInventory
    MSPsLogo() *MSPsLogo
    MSPsLogs() *MSPsLogs
    MSPsLicenses() *MSPsLicenses
    MSPsMarvis() *MSPsMarvis
    MSPsOrgGroups() *MSPsOrgGroups
    MSPsOrgs() *MSPsOrgs
    MSPsSLEs() *MSPsSLEs
    MSPsSSORoles() *MSPsSSORoles
    MSPsSSO() *MSPsSSO
    MSPsTickets() *MSPsTickets
    Orgs() *Orgs
    OrgsAdmins() *OrgsAdmins
    OrgsAlarms() *OrgsAlarms
    OrgsAlarmTemplates() *OrgsAlarmTemplates
    OrgsAPTemplates() *OrgsAPTemplates
    OrgsAPITokens() *OrgsAPITokens
    OrgsAssets() *OrgsAssets
    OrgsAssetFilters() *OrgsAssetFilters
    OrgsCert() *OrgsCert
    OrgsAntivirusProfiles() *OrgsAntivirusProfiles
    OrgsClientsMarvis() *OrgsClientsMarvis
    OrgsClientsNAC() *OrgsClientsNAC
    OrgsClientsWan() *OrgsClientsWan
    OrgsClientsWired() *OrgsClientsWired
    OrgsClientsWireless() *OrgsClientsWireless
    OrgsClientsSDK() *OrgsClientsSDK
    OrgsCradlepoint() *OrgsCradlepoint
    OrgsCRL() *OrgsCRL
    OrgsDeviceProfiles() *OrgsDeviceProfiles
    OrgsDevices() *OrgsDevices
    OrgsDevicesSSR() *OrgsDevicesSSR
    OrgsEVPNTopologies() *OrgsEVPNTopologies
    OrgsDevicesOthers() *OrgsDevicesOthers
    OrgsSettingZscaler() *OrgsSettingZscaler
    OrgsSCEP() *OrgsSCEP
    OrgsGatewayTemplates() *OrgsGatewayTemplates
    OrgsGuests() *OrgsGuests
    OrgsIDPProfiles() *OrgsIDPProfiles
    OrgsInventory() *OrgsInventory
    OrgsJSE() *OrgsJSE
    OrgsJSI() *OrgsJSI
    OrgsLicenses() *OrgsLicenses
    OrgsLinkedApplications() *OrgsLinkedApplications
    OrgsLogs() *OrgsLogs
    OrgsMarvis() *OrgsMarvis
    OrgsMaps() *OrgsMaps
    OrgsMxClusters() *OrgsMxClusters
    OrgsMxEdges() *OrgsMxEdges
    OrgsMxTunnels() *OrgsMxTunnels
    OrgsNACIDP() *OrgsNACIDP
    OrgsNACTags() *OrgsNACTags
    OrgsNACPortals() *OrgsNACPortals
    OrgsNACCRL() *OrgsNACCRL
    OrgsStats() *OrgsStats
    OrgsStatsAssets() *OrgsStatsAssets
    OrgsStatsBGPPeers() *OrgsStatsBGPPeers
    OrgsStatsDevices() *OrgsStatsDevices
    OrgsStatsMxEdges() *OrgsStatsMxEdges
    OrgsStatsOtherDevices() *OrgsStatsOtherDevices
    OrgsStatsPorts() *OrgsStatsPorts
    OrgsStatsSites() *OrgsStatsSites
    OrgsStatsTunnels() *OrgsStatsTunnels
    OrgsStatsVPNPeers() *OrgsStatsVPNPeers
    OrgsNACRules() *OrgsNACRules
    OrgsNetworkTemplates() *OrgsNetworkTemplates
    OrgsNetworks() *OrgsNetworks
    OrgsPremiumAnalytics() *OrgsPremiumAnalytics
    OrgsPsks() *OrgsPsks
    OrgsPskPortals() *OrgsPskPortals
    OrgsRFTemplates() *OrgsRFTemplates
    OrgsSDKInvites() *OrgsSDKInvites
    OrgsSDKTemplates() *OrgsSDKTemplates
    OrgsSecPolicies() *OrgsSecPolicies
    OrgsServices() *OrgsServices
    OrgsServicePolicies() *OrgsServicePolicies
    OrgsSetting() *OrgsSetting
    OrgsSitegroups() *OrgsSitegroups
    OrgsSites() *OrgsSites
    OrgsSiteTemplates() *OrgsSiteTemplates
    OrgsSLEs() *OrgsSLEs
    OrgsSSORoles() *OrgsSSORoles
    OrgsSSO() *OrgsSSO
    OrgsSubscriptions() *OrgsSubscriptions
    OrgsWLANTemplates() *OrgsWLANTemplates
    OrgsTickets() *OrgsTickets
    OrgsUserMACs() *OrgsUserMACs
    OrgsVars() *OrgsVars
    OrgsVPNs() *OrgsVPNs
    OrgsWebhooks() *OrgsWebhooks
    OrgsWlans() *OrgsWlans
    OrgsWxRules() *OrgsWxRules
    OrgsWxTags() *OrgsWxTags
    OrgsWxTunnels() *OrgsWxTunnels
    Sites() *Sites
    SitesAlarms() *SitesAlarms
    SitesAPTemplates() *SitesAPTemplates
    SitesApplications() *SitesApplications
    SitesAnomaly() *SitesAnomaly
    SitesAssetFilters() *SitesAssetFilters
    SitesAssets() *SitesAssets
    SitesBeacons() *SitesBeacons
    SitesClientsNAC() *SitesClientsNAC
    SitesClientsWan() *SitesClientsWan
    SitesClientsWired() *SitesClientsWired
    SitesClientsWireless() *SitesClientsWireless
    SitesDevices() *SitesDevices
    SitesDevicesWireless() *SitesDevicesWireless
    SitesDevicesOthers() *SitesDevicesOthers
    SitesDevicesWired() *SitesDevicesWired
    SitesDevicesWiredVirtualChassis() *SitesDevicesWiredVirtualChassis
    SitesDevicesWANCluster() *SitesDevicesWANCluster
    SitesDeviceProfiles() *SitesDeviceProfiles
    SitesEvents() *SitesEvents
    SitesEVPNTopologies() *SitesEVPNTopologies
    SitesGatewayTemplates() *SitesGatewayTemplates
    SitesGuests() *SitesGuests
    SitesInsights() *SitesInsights
    SitesJSE() *SitesJSE
    SitesLicenses() *SitesLicenses
    SitesLocation() *SitesLocation
    SitesMaps() *SitesMaps
    SitesMapsAutoPlacement() *SitesMapsAutoPlacement
    SitesMapsAutoZone() *SitesMapsAutoZone
    SitesMxEdges() *SitesMxEdges
    SitesNetworkTemplates() *SitesNetworkTemplates
    SitesNetworks() *SitesNetworks
    SitesPsks() *SitesPsks
    SitesRFTemplates() *SitesRFTemplates
    SitesRfdiags() *SitesRfdiags
    SitesRogues() *SitesRogues
    SitesRRM() *SitesRRM
    SitesRSSIZones() *SitesRSSIZones
    SitesServices() *SitesServices
    SitesServicePolicies() *SitesServicePolicies
    SitesSetting() *SitesSetting
    SitesSiteTemplates() *SitesSiteTemplates
    SitesSkyatp() *SitesSkyatp
    SitesSLEs() *SitesSLEs
    SitesSyntheticTests() *SitesSyntheticTests
    SitesUISettings() *SitesUISettings
    SitesVBeacons() *SitesVBeacons
    SitesVPNs() *SitesVPNs
    SitesWANUsages() *SitesWANUsages
    SitesWebhooks() *SitesWebhooks
    SitesWlans() *SitesWlans
    SitesWxRules() *SitesWxRules
    SitesWxTags() *SitesWxTags
    SitesWxTunnels() *SitesWxTunnels
    SitesZones() *SitesZones
    SitesStats() *SitesStats
    SitesStatsApps() *SitesStatsApps
    SitesStatsAssets() *SitesStatsAssets
    SitesStatsBeacons() *SitesStatsBeacons
    SitesStatsBGPPeers() *SitesStatsBGPPeers
    SitesStatsCalls() *SitesStatsCalls
    SitesStatsClientsWireless() *SitesStatsClientsWireless
    SitesStatsClientsSDK() *SitesStatsClientsSDK
    SitesStatsDevices() *SitesStatsDevices
    SitesStatsMxEdges() *SitesStatsMxEdges
    SitesStatsPorts() *SitesStatsPorts
    SitesStatsWxRules() *SitesStatsWxRules
    SitesStatsZones() *SitesStatsZones
    SitesStatsDiscoveredSwitches() *SitesStatsDiscoveredSwitches
    ConstantsDefinitions() *ConstantsDefinitions
    ConstantsEvents() *ConstantsEvents
    ConstantsModels() *ConstantsModels
    SelfAccount() *SelfAccount
    SelfAPIToken() *SelfAPIToken
    SelfOAuth2() *SelfOAuth2
    SelfMFA() *SelfMFA
    SelfAlarms() *SelfAlarms
    SelfAuditLogs() *SelfAuditLogs
    UtilitiesCommon() *UtilitiesCommon
    UtilitiesWAN() *UtilitiesWAN
    UtilitiesLAN() *UtilitiesLAN
    UtilitiesWiFi() *UtilitiesWiFi
    UtilitiesPCAPs() *UtilitiesPCAPs
    UtilitiesLocation() *UtilitiesLocation
    UtilitiesMxEdge() *UtilitiesMxEdge
    UtilitiesUpgrade() *UtilitiesUpgrade
    UserAgent() *string
}

// client is an implementation of the Client interface.
type client struct {
    callBuilderFactory              https.CallBuilderFactory
    configuration                   Configuration
    userAgent                       string
    admins                          Admins
    adminsLogin                     AdminsLogin
    adminsLogout                    AdminsLogout
    adminsRecoverPassword           AdminsRecoverPassword
    adminsLookup                    AdminsLookup
    adminsLoginOAuth2               AdminsLoginOAuth2
    installer                       Installer
    mSPs                            MSPs
    mSPsAdmins                      MSPsAdmins
    orgsSecIntelProfiles            OrgsSecIntelProfiles
    sitesSecIntelProfiles           SitesSecIntelProfiles
    mSPsInventory                   MSPsInventory
    mSPsLogo                        MSPsLogo
    mSPsLogs                        MSPsLogs
    mSPsLicenses                    MSPsLicenses
    mSPsMarvis                      MSPsMarvis
    mSPsOrgGroups                   MSPsOrgGroups
    mSPsOrgs                        MSPsOrgs
    mSPsSLEs                        MSPsSLEs
    mSPsSSORoles                    MSPsSSORoles
    mSPsSSO                         MSPsSSO
    mSPsTickets                     MSPsTickets
    orgs                            Orgs
    orgsAdmins                      OrgsAdmins
    orgsAlarms                      OrgsAlarms
    orgsAlarmTemplates              OrgsAlarmTemplates
    orgsAPTemplates                 OrgsAPTemplates
    orgsAPITokens                   OrgsAPITokens
    orgsAssets                      OrgsAssets
    orgsAssetFilters                OrgsAssetFilters
    orgsCert                        OrgsCert
    orgsAntivirusProfiles           OrgsAntivirusProfiles
    orgsClientsMarvis               OrgsClientsMarvis
    orgsClientsNAC                  OrgsClientsNAC
    orgsClientsWan                  OrgsClientsWan
    orgsClientsWired                OrgsClientsWired
    orgsClientsWireless             OrgsClientsWireless
    orgsClientsSDK                  OrgsClientsSDK
    orgsCradlepoint                 OrgsCradlepoint
    orgsCRL                         OrgsCRL
    orgsDeviceProfiles              OrgsDeviceProfiles
    orgsDevices                     OrgsDevices
    orgsDevicesSSR                  OrgsDevicesSSR
    orgsEVPNTopologies              OrgsEVPNTopologies
    orgsDevicesOthers               OrgsDevicesOthers
    orgsSettingZscaler              OrgsSettingZscaler
    orgsSCEP                        OrgsSCEP
    orgsGatewayTemplates            OrgsGatewayTemplates
    orgsGuests                      OrgsGuests
    orgsIDPProfiles                 OrgsIDPProfiles
    orgsInventory                   OrgsInventory
    orgsJSE                         OrgsJSE
    orgsJSI                         OrgsJSI
    orgsLicenses                    OrgsLicenses
    orgsLinkedApplications          OrgsLinkedApplications
    orgsLogs                        OrgsLogs
    orgsMarvis                      OrgsMarvis
    orgsMaps                        OrgsMaps
    orgsMxClusters                  OrgsMxClusters
    orgsMxEdges                     OrgsMxEdges
    orgsMxTunnels                   OrgsMxTunnels
    orgsNACIDP                      OrgsNACIDP
    orgsNACTags                     OrgsNACTags
    orgsNACPortals                  OrgsNACPortals
    orgsNACCRL                      OrgsNACCRL
    orgsStats                       OrgsStats
    orgsStatsAssets                 OrgsStatsAssets
    orgsStatsBGPPeers               OrgsStatsBGPPeers
    orgsStatsDevices                OrgsStatsDevices
    orgsStatsMxEdges                OrgsStatsMxEdges
    orgsStatsOtherDevices           OrgsStatsOtherDevices
    orgsStatsPorts                  OrgsStatsPorts
    orgsStatsSites                  OrgsStatsSites
    orgsStatsTunnels                OrgsStatsTunnels
    orgsStatsVPNPeers               OrgsStatsVPNPeers
    orgsNACRules                    OrgsNACRules
    orgsNetworkTemplates            OrgsNetworkTemplates
    orgsNetworks                    OrgsNetworks
    orgsPremiumAnalytics            OrgsPremiumAnalytics
    orgsPsks                        OrgsPsks
    orgsPskPortals                  OrgsPskPortals
    orgsRFTemplates                 OrgsRFTemplates
    orgsSDKInvites                  OrgsSDKInvites
    orgsSDKTemplates                OrgsSDKTemplates
    orgsSecPolicies                 OrgsSecPolicies
    orgsServices                    OrgsServices
    orgsServicePolicies             OrgsServicePolicies
    orgsSetting                     OrgsSetting
    orgsSitegroups                  OrgsSitegroups
    orgsSites                       OrgsSites
    orgsSiteTemplates               OrgsSiteTemplates
    orgsSLEs                        OrgsSLEs
    orgsSSORoles                    OrgsSSORoles
    orgsSSO                         OrgsSSO
    orgsSubscriptions               OrgsSubscriptions
    orgsWLANTemplates               OrgsWLANTemplates
    orgsTickets                     OrgsTickets
    orgsUserMACs                    OrgsUserMACs
    orgsVars                        OrgsVars
    orgsVPNs                        OrgsVPNs
    orgsWebhooks                    OrgsWebhooks
    orgsWlans                       OrgsWlans
    orgsWxRules                     OrgsWxRules
    orgsWxTags                      OrgsWxTags
    orgsWxTunnels                   OrgsWxTunnels
    sites                           Sites
    sitesAlarms                     SitesAlarms
    sitesAPTemplates                SitesAPTemplates
    sitesApplications               SitesApplications
    sitesAnomaly                    SitesAnomaly
    sitesAssetFilters               SitesAssetFilters
    sitesAssets                     SitesAssets
    sitesBeacons                    SitesBeacons
    sitesClientsNAC                 SitesClientsNAC
    sitesClientsWan                 SitesClientsWan
    sitesClientsWired               SitesClientsWired
    sitesClientsWireless            SitesClientsWireless
    sitesDevices                    SitesDevices
    sitesDevicesWireless            SitesDevicesWireless
    sitesDevicesOthers              SitesDevicesOthers
    sitesDevicesWired               SitesDevicesWired
    sitesDevicesWiredVirtualChassis SitesDevicesWiredVirtualChassis
    sitesDevicesWANCluster          SitesDevicesWANCluster
    sitesDeviceProfiles             SitesDeviceProfiles
    sitesEvents                     SitesEvents
    sitesEVPNTopologies             SitesEVPNTopologies
    sitesGatewayTemplates           SitesGatewayTemplates
    sitesGuests                     SitesGuests
    sitesInsights                   SitesInsights
    sitesJSE                        SitesJSE
    sitesLicenses                   SitesLicenses
    sitesLocation                   SitesLocation
    sitesMaps                       SitesMaps
    sitesMapsAutoPlacement          SitesMapsAutoPlacement
    sitesMapsAutoZone               SitesMapsAutoZone
    sitesMxEdges                    SitesMxEdges
    sitesNetworkTemplates           SitesNetworkTemplates
    sitesNetworks                   SitesNetworks
    sitesPsks                       SitesPsks
    sitesRFTemplates                SitesRFTemplates
    sitesRfdiags                    SitesRfdiags
    sitesRogues                     SitesRogues
    sitesRRM                        SitesRRM
    sitesRSSIZones                  SitesRSSIZones
    sitesServices                   SitesServices
    sitesServicePolicies            SitesServicePolicies
    sitesSetting                    SitesSetting
    sitesSiteTemplates              SitesSiteTemplates
    sitesSkyatp                     SitesSkyatp
    sitesSLEs                       SitesSLEs
    sitesSyntheticTests             SitesSyntheticTests
    sitesUISettings                 SitesUISettings
    sitesVBeacons                   SitesVBeacons
    sitesVPNs                       SitesVPNs
    sitesWANUsages                  SitesWANUsages
    sitesWebhooks                   SitesWebhooks
    sitesWlans                      SitesWlans
    sitesWxRules                    SitesWxRules
    sitesWxTags                     SitesWxTags
    sitesWxTunnels                  SitesWxTunnels
    sitesZones                      SitesZones
    sitesStats                      SitesStats
    sitesStatsApps                  SitesStatsApps
    sitesStatsAssets                SitesStatsAssets
    sitesStatsBeacons               SitesStatsBeacons
    sitesStatsBGPPeers              SitesStatsBGPPeers
    sitesStatsCalls                 SitesStatsCalls
    sitesStatsClientsWireless       SitesStatsClientsWireless
    sitesStatsClientsSDK            SitesStatsClientsSDK
    sitesStatsDevices               SitesStatsDevices
    sitesStatsMxEdges               SitesStatsMxEdges
    sitesStatsPorts                 SitesStatsPorts
    sitesStatsWxRules               SitesStatsWxRules
    sitesStatsZones                 SitesStatsZones
    sitesStatsDiscoveredSwitches    SitesStatsDiscoveredSwitches
    constantsDefinitions            ConstantsDefinitions
    constantsEvents                 ConstantsEvents
    constantsModels                 ConstantsModels
    selfAccount                     SelfAccount
    selfAPIToken                    SelfAPIToken
    selfOAuth2                      SelfOAuth2
    selfMFA                         SelfMFA
    selfAlarms                      SelfAlarms
    selfAuditLogs                   SelfAuditLogs
    utilitiesCommon                 UtilitiesCommon
    utilitiesWAN                    UtilitiesWAN
    utilitiesLAN                    UtilitiesLAN
    utilitiesWiFi                   UtilitiesWiFi
    utilitiesPCAPs                  UtilitiesPCAPs
    utilitiesLocation               UtilitiesLocation
    utilitiesMxEdge                 UtilitiesMxEdge
    utilitiesUpgrade                UtilitiesUpgrade
}

// NewClient is the constructor for creating a new client instance.
// It takes a Configuration object as a parameter and returns the Client interface.
func NewClient(configuration Configuration) ClientInterface {
    client := &client{
        configuration: configuration,
    }
    
    client.userAgent = utilities.UpdateUserAgent("SDK 2410.1.18")
    client.callBuilderFactory = callBuilderHandler(
    	func(server string) string {
    		if server == "" {
    			server = "API Host"
    		}
    		return getBaseUri(Server(server), client.configuration)
    	},
    	createAuthenticationFromConfig(configuration),
    	https.NewHttpClient(configuration.HttpConfiguration()),
    	configuration.httpConfiguration.RetryConfiguration(),
    	https.Indexed,
        withUserAgent(client.userAgent),
        withLogger(configuration.loggerConfiguration),
    )
    
    baseController := NewBaseController(client)
    client.admins = *NewAdmins(*baseController)
    client.adminsLogin = *NewAdminsLogin(*baseController)
    client.adminsLogout = *NewAdminsLogout(*baseController)
    client.adminsRecoverPassword = *NewAdminsRecoverPassword(*baseController)
    client.adminsLookup = *NewAdminsLookup(*baseController)
    client.adminsLoginOAuth2 = *NewAdminsLoginOAuth2(*baseController)
    client.installer = *NewInstaller(*baseController)
    client.mSPs = *NewMSPs(*baseController)
    client.mSPsAdmins = *NewMSPsAdmins(*baseController)
    client.orgsSecIntelProfiles = *NewOrgsSecIntelProfiles(*baseController)
    client.sitesSecIntelProfiles = *NewSitesSecIntelProfiles(*baseController)
    client.mSPsInventory = *NewMSPsInventory(*baseController)
    client.mSPsLogo = *NewMSPsLogo(*baseController)
    client.mSPsLogs = *NewMSPsLogs(*baseController)
    client.mSPsLicenses = *NewMSPsLicenses(*baseController)
    client.mSPsMarvis = *NewMSPsMarvis(*baseController)
    client.mSPsOrgGroups = *NewMSPsOrgGroups(*baseController)
    client.mSPsOrgs = *NewMSPsOrgs(*baseController)
    client.mSPsSLEs = *NewMSPsSLEs(*baseController)
    client.mSPsSSORoles = *NewMSPsSSORoles(*baseController)
    client.mSPsSSO = *NewMSPsSSO(*baseController)
    client.mSPsTickets = *NewMSPsTickets(*baseController)
    client.orgs = *NewOrgs(*baseController)
    client.orgsAdmins = *NewOrgsAdmins(*baseController)
    client.orgsAlarms = *NewOrgsAlarms(*baseController)
    client.orgsAlarmTemplates = *NewOrgsAlarmTemplates(*baseController)
    client.orgsAPTemplates = *NewOrgsAPTemplates(*baseController)
    client.orgsAPITokens = *NewOrgsAPITokens(*baseController)
    client.orgsAssets = *NewOrgsAssets(*baseController)
    client.orgsAssetFilters = *NewOrgsAssetFilters(*baseController)
    client.orgsCert = *NewOrgsCert(*baseController)
    client.orgsAntivirusProfiles = *NewOrgsAntivirusProfiles(*baseController)
    client.orgsClientsMarvis = *NewOrgsClientsMarvis(*baseController)
    client.orgsClientsNAC = *NewOrgsClientsNAC(*baseController)
    client.orgsClientsWan = *NewOrgsClientsWan(*baseController)
    client.orgsClientsWired = *NewOrgsClientsWired(*baseController)
    client.orgsClientsWireless = *NewOrgsClientsWireless(*baseController)
    client.orgsClientsSDK = *NewOrgsClientsSDK(*baseController)
    client.orgsCradlepoint = *NewOrgsCradlepoint(*baseController)
    client.orgsCRL = *NewOrgsCRL(*baseController)
    client.orgsDeviceProfiles = *NewOrgsDeviceProfiles(*baseController)
    client.orgsDevices = *NewOrgsDevices(*baseController)
    client.orgsDevicesSSR = *NewOrgsDevicesSSR(*baseController)
    client.orgsEVPNTopologies = *NewOrgsEVPNTopologies(*baseController)
    client.orgsDevicesOthers = *NewOrgsDevicesOthers(*baseController)
    client.orgsSettingZscaler = *NewOrgsSettingZscaler(*baseController)
    client.orgsSCEP = *NewOrgsSCEP(*baseController)
    client.orgsGatewayTemplates = *NewOrgsGatewayTemplates(*baseController)
    client.orgsGuests = *NewOrgsGuests(*baseController)
    client.orgsIDPProfiles = *NewOrgsIDPProfiles(*baseController)
    client.orgsInventory = *NewOrgsInventory(*baseController)
    client.orgsJSE = *NewOrgsJSE(*baseController)
    client.orgsJSI = *NewOrgsJSI(*baseController)
    client.orgsLicenses = *NewOrgsLicenses(*baseController)
    client.orgsLinkedApplications = *NewOrgsLinkedApplications(*baseController)
    client.orgsLogs = *NewOrgsLogs(*baseController)
    client.orgsMarvis = *NewOrgsMarvis(*baseController)
    client.orgsMaps = *NewOrgsMaps(*baseController)
    client.orgsMxClusters = *NewOrgsMxClusters(*baseController)
    client.orgsMxEdges = *NewOrgsMxEdges(*baseController)
    client.orgsMxTunnels = *NewOrgsMxTunnels(*baseController)
    client.orgsNACIDP = *NewOrgsNACIDP(*baseController)
    client.orgsNACTags = *NewOrgsNACTags(*baseController)
    client.orgsNACPortals = *NewOrgsNACPortals(*baseController)
    client.orgsNACCRL = *NewOrgsNACCRL(*baseController)
    client.orgsStats = *NewOrgsStats(*baseController)
    client.orgsStatsAssets = *NewOrgsStatsAssets(*baseController)
    client.orgsStatsBGPPeers = *NewOrgsStatsBGPPeers(*baseController)
    client.orgsStatsDevices = *NewOrgsStatsDevices(*baseController)
    client.orgsStatsMxEdges = *NewOrgsStatsMxEdges(*baseController)
    client.orgsStatsOtherDevices = *NewOrgsStatsOtherDevices(*baseController)
    client.orgsStatsPorts = *NewOrgsStatsPorts(*baseController)
    client.orgsStatsSites = *NewOrgsStatsSites(*baseController)
    client.orgsStatsTunnels = *NewOrgsStatsTunnels(*baseController)
    client.orgsStatsVPNPeers = *NewOrgsStatsVPNPeers(*baseController)
    client.orgsNACRules = *NewOrgsNACRules(*baseController)
    client.orgsNetworkTemplates = *NewOrgsNetworkTemplates(*baseController)
    client.orgsNetworks = *NewOrgsNetworks(*baseController)
    client.orgsPremiumAnalytics = *NewOrgsPremiumAnalytics(*baseController)
    client.orgsPsks = *NewOrgsPsks(*baseController)
    client.orgsPskPortals = *NewOrgsPskPortals(*baseController)
    client.orgsRFTemplates = *NewOrgsRFTemplates(*baseController)
    client.orgsSDKInvites = *NewOrgsSDKInvites(*baseController)
    client.orgsSDKTemplates = *NewOrgsSDKTemplates(*baseController)
    client.orgsSecPolicies = *NewOrgsSecPolicies(*baseController)
    client.orgsServices = *NewOrgsServices(*baseController)
    client.orgsServicePolicies = *NewOrgsServicePolicies(*baseController)
    client.orgsSetting = *NewOrgsSetting(*baseController)
    client.orgsSitegroups = *NewOrgsSitegroups(*baseController)
    client.orgsSites = *NewOrgsSites(*baseController)
    client.orgsSiteTemplates = *NewOrgsSiteTemplates(*baseController)
    client.orgsSLEs = *NewOrgsSLEs(*baseController)
    client.orgsSSORoles = *NewOrgsSSORoles(*baseController)
    client.orgsSSO = *NewOrgsSSO(*baseController)
    client.orgsSubscriptions = *NewOrgsSubscriptions(*baseController)
    client.orgsWLANTemplates = *NewOrgsWLANTemplates(*baseController)
    client.orgsTickets = *NewOrgsTickets(*baseController)
    client.orgsUserMACs = *NewOrgsUserMACs(*baseController)
    client.orgsVars = *NewOrgsVars(*baseController)
    client.orgsVPNs = *NewOrgsVPNs(*baseController)
    client.orgsWebhooks = *NewOrgsWebhooks(*baseController)
    client.orgsWlans = *NewOrgsWlans(*baseController)
    client.orgsWxRules = *NewOrgsWxRules(*baseController)
    client.orgsWxTags = *NewOrgsWxTags(*baseController)
    client.orgsWxTunnels = *NewOrgsWxTunnels(*baseController)
    client.sites = *NewSites(*baseController)
    client.sitesAlarms = *NewSitesAlarms(*baseController)
    client.sitesAPTemplates = *NewSitesAPTemplates(*baseController)
    client.sitesApplications = *NewSitesApplications(*baseController)
    client.sitesAnomaly = *NewSitesAnomaly(*baseController)
    client.sitesAssetFilters = *NewSitesAssetFilters(*baseController)
    client.sitesAssets = *NewSitesAssets(*baseController)
    client.sitesBeacons = *NewSitesBeacons(*baseController)
    client.sitesClientsNAC = *NewSitesClientsNAC(*baseController)
    client.sitesClientsWan = *NewSitesClientsWan(*baseController)
    client.sitesClientsWired = *NewSitesClientsWired(*baseController)
    client.sitesClientsWireless = *NewSitesClientsWireless(*baseController)
    client.sitesDevices = *NewSitesDevices(*baseController)
    client.sitesDevicesWireless = *NewSitesDevicesWireless(*baseController)
    client.sitesDevicesOthers = *NewSitesDevicesOthers(*baseController)
    client.sitesDevicesWired = *NewSitesDevicesWired(*baseController)
    client.sitesDevicesWiredVirtualChassis = *NewSitesDevicesWiredVirtualChassis(*baseController)
    client.sitesDevicesWANCluster = *NewSitesDevicesWANCluster(*baseController)
    client.sitesDeviceProfiles = *NewSitesDeviceProfiles(*baseController)
    client.sitesEvents = *NewSitesEvents(*baseController)
    client.sitesEVPNTopologies = *NewSitesEVPNTopologies(*baseController)
    client.sitesGatewayTemplates = *NewSitesGatewayTemplates(*baseController)
    client.sitesGuests = *NewSitesGuests(*baseController)
    client.sitesInsights = *NewSitesInsights(*baseController)
    client.sitesJSE = *NewSitesJSE(*baseController)
    client.sitesLicenses = *NewSitesLicenses(*baseController)
    client.sitesLocation = *NewSitesLocation(*baseController)
    client.sitesMaps = *NewSitesMaps(*baseController)
    client.sitesMapsAutoPlacement = *NewSitesMapsAutoPlacement(*baseController)
    client.sitesMapsAutoZone = *NewSitesMapsAutoZone(*baseController)
    client.sitesMxEdges = *NewSitesMxEdges(*baseController)
    client.sitesNetworkTemplates = *NewSitesNetworkTemplates(*baseController)
    client.sitesNetworks = *NewSitesNetworks(*baseController)
    client.sitesPsks = *NewSitesPsks(*baseController)
    client.sitesRFTemplates = *NewSitesRFTemplates(*baseController)
    client.sitesRfdiags = *NewSitesRfdiags(*baseController)
    client.sitesRogues = *NewSitesRogues(*baseController)
    client.sitesRRM = *NewSitesRRM(*baseController)
    client.sitesRSSIZones = *NewSitesRSSIZones(*baseController)
    client.sitesServices = *NewSitesServices(*baseController)
    client.sitesServicePolicies = *NewSitesServicePolicies(*baseController)
    client.sitesSetting = *NewSitesSetting(*baseController)
    client.sitesSiteTemplates = *NewSitesSiteTemplates(*baseController)
    client.sitesSkyatp = *NewSitesSkyatp(*baseController)
    client.sitesSLEs = *NewSitesSLEs(*baseController)
    client.sitesSyntheticTests = *NewSitesSyntheticTests(*baseController)
    client.sitesUISettings = *NewSitesUISettings(*baseController)
    client.sitesVBeacons = *NewSitesVBeacons(*baseController)
    client.sitesVPNs = *NewSitesVPNs(*baseController)
    client.sitesWANUsages = *NewSitesWANUsages(*baseController)
    client.sitesWebhooks = *NewSitesWebhooks(*baseController)
    client.sitesWlans = *NewSitesWlans(*baseController)
    client.sitesWxRules = *NewSitesWxRules(*baseController)
    client.sitesWxTags = *NewSitesWxTags(*baseController)
    client.sitesWxTunnels = *NewSitesWxTunnels(*baseController)
    client.sitesZones = *NewSitesZones(*baseController)
    client.sitesStats = *NewSitesStats(*baseController)
    client.sitesStatsApps = *NewSitesStatsApps(*baseController)
    client.sitesStatsAssets = *NewSitesStatsAssets(*baseController)
    client.sitesStatsBeacons = *NewSitesStatsBeacons(*baseController)
    client.sitesStatsBGPPeers = *NewSitesStatsBGPPeers(*baseController)
    client.sitesStatsCalls = *NewSitesStatsCalls(*baseController)
    client.sitesStatsClientsWireless = *NewSitesStatsClientsWireless(*baseController)
    client.sitesStatsClientsSDK = *NewSitesStatsClientsSDK(*baseController)
    client.sitesStatsDevices = *NewSitesStatsDevices(*baseController)
    client.sitesStatsMxEdges = *NewSitesStatsMxEdges(*baseController)
    client.sitesStatsPorts = *NewSitesStatsPorts(*baseController)
    client.sitesStatsWxRules = *NewSitesStatsWxRules(*baseController)
    client.sitesStatsZones = *NewSitesStatsZones(*baseController)
    client.sitesStatsDiscoveredSwitches = *NewSitesStatsDiscoveredSwitches(*baseController)
    client.constantsDefinitions = *NewConstantsDefinitions(*baseController)
    client.constantsEvents = *NewConstantsEvents(*baseController)
    client.constantsModels = *NewConstantsModels(*baseController)
    client.selfAccount = *NewSelfAccount(*baseController)
    client.selfAPIToken = *NewSelfAPIToken(*baseController)
    client.selfOAuth2 = *NewSelfOAuth2(*baseController)
    client.selfMFA = *NewSelfMFA(*baseController)
    client.selfAlarms = *NewSelfAlarms(*baseController)
    client.selfAuditLogs = *NewSelfAuditLogs(*baseController)
    client.utilitiesCommon = *NewUtilitiesCommon(*baseController)
    client.utilitiesWAN = *NewUtilitiesWAN(*baseController)
    client.utilitiesLAN = *NewUtilitiesLAN(*baseController)
    client.utilitiesWiFi = *NewUtilitiesWiFi(*baseController)
    client.utilitiesPCAPs = *NewUtilitiesPCAPs(*baseController)
    client.utilitiesLocation = *NewUtilitiesLocation(*baseController)
    client.utilitiesMxEdge = *NewUtilitiesMxEdge(*baseController)
    client.utilitiesUpgrade = *NewUtilitiesUpgrade(*baseController)
    return client
}

// Configuration returns the configuration instance of the client.
func (c *client) Configuration() Configuration {
    return c.configuration
}

// CloneWithConfiguration returns a new copy with the provided options of the configuration instance of the client.
func (c *client) CloneWithConfiguration(options ...ConfigurationOptions) ClientInterface {
    return NewClient(c.configuration.cloneWithOptions(options...))
}

// Admins returns the admins instance of the client.
func (c *client) Admins() *Admins {
    return &c.admins
}

// AdminsLogin returns the adminsLogin instance of the client.
func (c *client) AdminsLogin() *AdminsLogin {
    return &c.adminsLogin
}

// AdminsLogout returns the adminsLogout instance of the client.
func (c *client) AdminsLogout() *AdminsLogout {
    return &c.adminsLogout
}

// AdminsRecoverPassword returns the adminsRecoverPassword instance of the client.
func (c *client) AdminsRecoverPassword() *AdminsRecoverPassword {
    return &c.adminsRecoverPassword
}

// AdminsLookup returns the adminsLookup instance of the client.
func (c *client) AdminsLookup() *AdminsLookup {
    return &c.adminsLookup
}

// AdminsLoginOAuth2 returns the adminsLoginOAuth2 instance of the client.
func (c *client) AdminsLoginOAuth2() *AdminsLoginOAuth2 {
    return &c.adminsLoginOAuth2
}

// Installer returns the installer instance of the client.
func (c *client) Installer() *Installer {
    return &c.installer
}

// MSPs returns the mSPs instance of the client.
func (c *client) MSPs() *MSPs {
    return &c.mSPs
}

// MSPsAdmins returns the mSPsAdmins instance of the client.
func (c *client) MSPsAdmins() *MSPsAdmins {
    return &c.mSPsAdmins
}

// OrgsSecIntelProfiles returns the orgsSecIntelProfiles instance of the client.
func (c *client) OrgsSecIntelProfiles() *OrgsSecIntelProfiles {
    return &c.orgsSecIntelProfiles
}

// SitesSecIntelProfiles returns the sitesSecIntelProfiles instance of the client.
func (c *client) SitesSecIntelProfiles() *SitesSecIntelProfiles {
    return &c.sitesSecIntelProfiles
}

// MSPsInventory returns the mSPsInventory instance of the client.
func (c *client) MSPsInventory() *MSPsInventory {
    return &c.mSPsInventory
}

// MSPsLogo returns the mSPsLogo instance of the client.
func (c *client) MSPsLogo() *MSPsLogo {
    return &c.mSPsLogo
}

// MSPsLogs returns the mSPsLogs instance of the client.
func (c *client) MSPsLogs() *MSPsLogs {
    return &c.mSPsLogs
}

// MSPsLicenses returns the mSPsLicenses instance of the client.
func (c *client) MSPsLicenses() *MSPsLicenses {
    return &c.mSPsLicenses
}

// MSPsMarvis returns the mSPsMarvis instance of the client.
func (c *client) MSPsMarvis() *MSPsMarvis {
    return &c.mSPsMarvis
}

// MSPsOrgGroups returns the mSPsOrgGroups instance of the client.
func (c *client) MSPsOrgGroups() *MSPsOrgGroups {
    return &c.mSPsOrgGroups
}

// MSPsOrgs returns the mSPsOrgs instance of the client.
func (c *client) MSPsOrgs() *MSPsOrgs {
    return &c.mSPsOrgs
}

// MSPsSLEs returns the mSPsSLEs instance of the client.
func (c *client) MSPsSLEs() *MSPsSLEs {
    return &c.mSPsSLEs
}

// MSPsSSORoles returns the mSPsSSORoles instance of the client.
func (c *client) MSPsSSORoles() *MSPsSSORoles {
    return &c.mSPsSSORoles
}

// MSPsSSO returns the mSPsSSO instance of the client.
func (c *client) MSPsSSO() *MSPsSSO {
    return &c.mSPsSSO
}

// MSPsTickets returns the mSPsTickets instance of the client.
func (c *client) MSPsTickets() *MSPsTickets {
    return &c.mSPsTickets
}

// Orgs returns the orgs instance of the client.
func (c *client) Orgs() *Orgs {
    return &c.orgs
}

// OrgsAdmins returns the orgsAdmins instance of the client.
func (c *client) OrgsAdmins() *OrgsAdmins {
    return &c.orgsAdmins
}

// OrgsAlarms returns the orgsAlarms instance of the client.
func (c *client) OrgsAlarms() *OrgsAlarms {
    return &c.orgsAlarms
}

// OrgsAlarmTemplates returns the orgsAlarmTemplates instance of the client.
func (c *client) OrgsAlarmTemplates() *OrgsAlarmTemplates {
    return &c.orgsAlarmTemplates
}

// OrgsAPTemplates returns the orgsAPTemplates instance of the client.
func (c *client) OrgsAPTemplates() *OrgsAPTemplates {
    return &c.orgsAPTemplates
}

// OrgsAPITokens returns the orgsAPITokens instance of the client.
func (c *client) OrgsAPITokens() *OrgsAPITokens {
    return &c.orgsAPITokens
}

// OrgsAssets returns the orgsAssets instance of the client.
func (c *client) OrgsAssets() *OrgsAssets {
    return &c.orgsAssets
}

// OrgsAssetFilters returns the orgsAssetFilters instance of the client.
func (c *client) OrgsAssetFilters() *OrgsAssetFilters {
    return &c.orgsAssetFilters
}

// OrgsCert returns the orgsCert instance of the client.
func (c *client) OrgsCert() *OrgsCert {
    return &c.orgsCert
}

// OrgsAntivirusProfiles returns the orgsAntivirusProfiles instance of the client.
func (c *client) OrgsAntivirusProfiles() *OrgsAntivirusProfiles {
    return &c.orgsAntivirusProfiles
}

// OrgsClientsMarvis returns the orgsClientsMarvis instance of the client.
func (c *client) OrgsClientsMarvis() *OrgsClientsMarvis {
    return &c.orgsClientsMarvis
}

// OrgsClientsNAC returns the orgsClientsNAC instance of the client.
func (c *client) OrgsClientsNAC() *OrgsClientsNAC {
    return &c.orgsClientsNAC
}

// OrgsClientsWan returns the orgsClientsWan instance of the client.
func (c *client) OrgsClientsWan() *OrgsClientsWan {
    return &c.orgsClientsWan
}

// OrgsClientsWired returns the orgsClientsWired instance of the client.
func (c *client) OrgsClientsWired() *OrgsClientsWired {
    return &c.orgsClientsWired
}

// OrgsClientsWireless returns the orgsClientsWireless instance of the client.
func (c *client) OrgsClientsWireless() *OrgsClientsWireless {
    return &c.orgsClientsWireless
}

// OrgsClientsSDK returns the orgsClientsSDK instance of the client.
func (c *client) OrgsClientsSDK() *OrgsClientsSDK {
    return &c.orgsClientsSDK
}

// OrgsCradlepoint returns the orgsCradlepoint instance of the client.
func (c *client) OrgsCradlepoint() *OrgsCradlepoint {
    return &c.orgsCradlepoint
}

// OrgsCRL returns the orgsCRL instance of the client.
func (c *client) OrgsCRL() *OrgsCRL {
    return &c.orgsCRL
}

// OrgsDeviceProfiles returns the orgsDeviceProfiles instance of the client.
func (c *client) OrgsDeviceProfiles() *OrgsDeviceProfiles {
    return &c.orgsDeviceProfiles
}

// OrgsDevices returns the orgsDevices instance of the client.
func (c *client) OrgsDevices() *OrgsDevices {
    return &c.orgsDevices
}

// OrgsDevicesSSR returns the orgsDevicesSSR instance of the client.
func (c *client) OrgsDevicesSSR() *OrgsDevicesSSR {
    return &c.orgsDevicesSSR
}

// OrgsEVPNTopologies returns the orgsEVPNTopologies instance of the client.
func (c *client) OrgsEVPNTopologies() *OrgsEVPNTopologies {
    return &c.orgsEVPNTopologies
}

// OrgsDevicesOthers returns the orgsDevicesOthers instance of the client.
func (c *client) OrgsDevicesOthers() *OrgsDevicesOthers {
    return &c.orgsDevicesOthers
}

// OrgsSettingZscaler returns the orgsSettingZscaler instance of the client.
func (c *client) OrgsSettingZscaler() *OrgsSettingZscaler {
    return &c.orgsSettingZscaler
}

// OrgsSCEP returns the orgsSCEP instance of the client.
func (c *client) OrgsSCEP() *OrgsSCEP {
    return &c.orgsSCEP
}

// OrgsGatewayTemplates returns the orgsGatewayTemplates instance of the client.
func (c *client) OrgsGatewayTemplates() *OrgsGatewayTemplates {
    return &c.orgsGatewayTemplates
}

// OrgsGuests returns the orgsGuests instance of the client.
func (c *client) OrgsGuests() *OrgsGuests {
    return &c.orgsGuests
}

// OrgsIDPProfiles returns the orgsIDPProfiles instance of the client.
func (c *client) OrgsIDPProfiles() *OrgsIDPProfiles {
    return &c.orgsIDPProfiles
}

// OrgsInventory returns the orgsInventory instance of the client.
func (c *client) OrgsInventory() *OrgsInventory {
    return &c.orgsInventory
}

// OrgsJSE returns the orgsJSE instance of the client.
func (c *client) OrgsJSE() *OrgsJSE {
    return &c.orgsJSE
}

// OrgsJSI returns the orgsJSI instance of the client.
func (c *client) OrgsJSI() *OrgsJSI {
    return &c.orgsJSI
}

// OrgsLicenses returns the orgsLicenses instance of the client.
func (c *client) OrgsLicenses() *OrgsLicenses {
    return &c.orgsLicenses
}

// OrgsLinkedApplications returns the orgsLinkedApplications instance of the client.
func (c *client) OrgsLinkedApplications() *OrgsLinkedApplications {
    return &c.orgsLinkedApplications
}

// OrgsLogs returns the orgsLogs instance of the client.
func (c *client) OrgsLogs() *OrgsLogs {
    return &c.orgsLogs
}

// OrgsMarvis returns the orgsMarvis instance of the client.
func (c *client) OrgsMarvis() *OrgsMarvis {
    return &c.orgsMarvis
}

// OrgsMaps returns the orgsMaps instance of the client.
func (c *client) OrgsMaps() *OrgsMaps {
    return &c.orgsMaps
}

// OrgsMxClusters returns the orgsMxClusters instance of the client.
func (c *client) OrgsMxClusters() *OrgsMxClusters {
    return &c.orgsMxClusters
}

// OrgsMxEdges returns the orgsMxEdges instance of the client.
func (c *client) OrgsMxEdges() *OrgsMxEdges {
    return &c.orgsMxEdges
}

// OrgsMxTunnels returns the orgsMxTunnels instance of the client.
func (c *client) OrgsMxTunnels() *OrgsMxTunnels {
    return &c.orgsMxTunnels
}

// OrgsNACIDP returns the orgsNACIDP instance of the client.
func (c *client) OrgsNACIDP() *OrgsNACIDP {
    return &c.orgsNACIDP
}

// OrgsNACTags returns the orgsNACTags instance of the client.
func (c *client) OrgsNACTags() *OrgsNACTags {
    return &c.orgsNACTags
}

// OrgsNACPortals returns the orgsNACPortals instance of the client.
func (c *client) OrgsNACPortals() *OrgsNACPortals {
    return &c.orgsNACPortals
}

// OrgsNACCRL returns the orgsNACCRL instance of the client.
func (c *client) OrgsNACCRL() *OrgsNACCRL {
    return &c.orgsNACCRL
}

// OrgsStats returns the orgsStats instance of the client.
func (c *client) OrgsStats() *OrgsStats {
    return &c.orgsStats
}

// OrgsStatsAssets returns the orgsStatsAssets instance of the client.
func (c *client) OrgsStatsAssets() *OrgsStatsAssets {
    return &c.orgsStatsAssets
}

// OrgsStatsBGPPeers returns the orgsStatsBGPPeers instance of the client.
func (c *client) OrgsStatsBGPPeers() *OrgsStatsBGPPeers {
    return &c.orgsStatsBGPPeers
}

// OrgsStatsDevices returns the orgsStatsDevices instance of the client.
func (c *client) OrgsStatsDevices() *OrgsStatsDevices {
    return &c.orgsStatsDevices
}

// OrgsStatsMxEdges returns the orgsStatsMxEdges instance of the client.
func (c *client) OrgsStatsMxEdges() *OrgsStatsMxEdges {
    return &c.orgsStatsMxEdges
}

// OrgsStatsOtherDevices returns the orgsStatsOtherDevices instance of the client.
func (c *client) OrgsStatsOtherDevices() *OrgsStatsOtherDevices {
    return &c.orgsStatsOtherDevices
}

// OrgsStatsPorts returns the orgsStatsPorts instance of the client.
func (c *client) OrgsStatsPorts() *OrgsStatsPorts {
    return &c.orgsStatsPorts
}

// OrgsStatsSites returns the orgsStatsSites instance of the client.
func (c *client) OrgsStatsSites() *OrgsStatsSites {
    return &c.orgsStatsSites
}

// OrgsStatsTunnels returns the orgsStatsTunnels instance of the client.
func (c *client) OrgsStatsTunnels() *OrgsStatsTunnels {
    return &c.orgsStatsTunnels
}

// OrgsStatsVPNPeers returns the orgsStatsVPNPeers instance of the client.
func (c *client) OrgsStatsVPNPeers() *OrgsStatsVPNPeers {
    return &c.orgsStatsVPNPeers
}

// OrgsNACRules returns the orgsNACRules instance of the client.
func (c *client) OrgsNACRules() *OrgsNACRules {
    return &c.orgsNACRules
}

// OrgsNetworkTemplates returns the orgsNetworkTemplates instance of the client.
func (c *client) OrgsNetworkTemplates() *OrgsNetworkTemplates {
    return &c.orgsNetworkTemplates
}

// OrgsNetworks returns the orgsNetworks instance of the client.
func (c *client) OrgsNetworks() *OrgsNetworks {
    return &c.orgsNetworks
}

// OrgsPremiumAnalytics returns the orgsPremiumAnalytics instance of the client.
func (c *client) OrgsPremiumAnalytics() *OrgsPremiumAnalytics {
    return &c.orgsPremiumAnalytics
}

// OrgsPsks returns the orgsPsks instance of the client.
func (c *client) OrgsPsks() *OrgsPsks {
    return &c.orgsPsks
}

// OrgsPskPortals returns the orgsPskPortals instance of the client.
func (c *client) OrgsPskPortals() *OrgsPskPortals {
    return &c.orgsPskPortals
}

// OrgsRFTemplates returns the orgsRFTemplates instance of the client.
func (c *client) OrgsRFTemplates() *OrgsRFTemplates {
    return &c.orgsRFTemplates
}

// OrgsSDKInvites returns the orgsSDKInvites instance of the client.
func (c *client) OrgsSDKInvites() *OrgsSDKInvites {
    return &c.orgsSDKInvites
}

// OrgsSDKTemplates returns the orgsSDKTemplates instance of the client.
func (c *client) OrgsSDKTemplates() *OrgsSDKTemplates {
    return &c.orgsSDKTemplates
}

// OrgsSecPolicies returns the orgsSecPolicies instance of the client.
func (c *client) OrgsSecPolicies() *OrgsSecPolicies {
    return &c.orgsSecPolicies
}

// OrgsServices returns the orgsServices instance of the client.
func (c *client) OrgsServices() *OrgsServices {
    return &c.orgsServices
}

// OrgsServicePolicies returns the orgsServicePolicies instance of the client.
func (c *client) OrgsServicePolicies() *OrgsServicePolicies {
    return &c.orgsServicePolicies
}

// OrgsSetting returns the orgsSetting instance of the client.
func (c *client) OrgsSetting() *OrgsSetting {
    return &c.orgsSetting
}

// OrgsSitegroups returns the orgsSitegroups instance of the client.
func (c *client) OrgsSitegroups() *OrgsSitegroups {
    return &c.orgsSitegroups
}

// OrgsSites returns the orgsSites instance of the client.
func (c *client) OrgsSites() *OrgsSites {
    return &c.orgsSites
}

// OrgsSiteTemplates returns the orgsSiteTemplates instance of the client.
func (c *client) OrgsSiteTemplates() *OrgsSiteTemplates {
    return &c.orgsSiteTemplates
}

// OrgsSLEs returns the orgsSLEs instance of the client.
func (c *client) OrgsSLEs() *OrgsSLEs {
    return &c.orgsSLEs
}

// OrgsSSORoles returns the orgsSSORoles instance of the client.
func (c *client) OrgsSSORoles() *OrgsSSORoles {
    return &c.orgsSSORoles
}

// OrgsSSO returns the orgsSSO instance of the client.
func (c *client) OrgsSSO() *OrgsSSO {
    return &c.orgsSSO
}

// OrgsSubscriptions returns the orgsSubscriptions instance of the client.
func (c *client) OrgsSubscriptions() *OrgsSubscriptions {
    return &c.orgsSubscriptions
}

// OrgsWLANTemplates returns the orgsWLANTemplates instance of the client.
func (c *client) OrgsWLANTemplates() *OrgsWLANTemplates {
    return &c.orgsWLANTemplates
}

// OrgsTickets returns the orgsTickets instance of the client.
func (c *client) OrgsTickets() *OrgsTickets {
    return &c.orgsTickets
}

// OrgsUserMACs returns the orgsUserMACs instance of the client.
func (c *client) OrgsUserMACs() *OrgsUserMACs {
    return &c.orgsUserMACs
}

// OrgsVars returns the orgsVars instance of the client.
func (c *client) OrgsVars() *OrgsVars {
    return &c.orgsVars
}

// OrgsVPNs returns the orgsVPNs instance of the client.
func (c *client) OrgsVPNs() *OrgsVPNs {
    return &c.orgsVPNs
}

// OrgsWebhooks returns the orgsWebhooks instance of the client.
func (c *client) OrgsWebhooks() *OrgsWebhooks {
    return &c.orgsWebhooks
}

// OrgsWlans returns the orgsWlans instance of the client.
func (c *client) OrgsWlans() *OrgsWlans {
    return &c.orgsWlans
}

// OrgsWxRules returns the orgsWxRules instance of the client.
func (c *client) OrgsWxRules() *OrgsWxRules {
    return &c.orgsWxRules
}

// OrgsWxTags returns the orgsWxTags instance of the client.
func (c *client) OrgsWxTags() *OrgsWxTags {
    return &c.orgsWxTags
}

// OrgsWxTunnels returns the orgsWxTunnels instance of the client.
func (c *client) OrgsWxTunnels() *OrgsWxTunnels {
    return &c.orgsWxTunnels
}

// Sites returns the sites instance of the client.
func (c *client) Sites() *Sites {
    return &c.sites
}

// SitesAlarms returns the sitesAlarms instance of the client.
func (c *client) SitesAlarms() *SitesAlarms {
    return &c.sitesAlarms
}

// SitesAPTemplates returns the sitesAPTemplates instance of the client.
func (c *client) SitesAPTemplates() *SitesAPTemplates {
    return &c.sitesAPTemplates
}

// SitesApplications returns the sitesApplications instance of the client.
func (c *client) SitesApplications() *SitesApplications {
    return &c.sitesApplications
}

// SitesAnomaly returns the sitesAnomaly instance of the client.
func (c *client) SitesAnomaly() *SitesAnomaly {
    return &c.sitesAnomaly
}

// SitesAssetFilters returns the sitesAssetFilters instance of the client.
func (c *client) SitesAssetFilters() *SitesAssetFilters {
    return &c.sitesAssetFilters
}

// SitesAssets returns the sitesAssets instance of the client.
func (c *client) SitesAssets() *SitesAssets {
    return &c.sitesAssets
}

// SitesBeacons returns the sitesBeacons instance of the client.
func (c *client) SitesBeacons() *SitesBeacons {
    return &c.sitesBeacons
}

// SitesClientsNAC returns the sitesClientsNAC instance of the client.
func (c *client) SitesClientsNAC() *SitesClientsNAC {
    return &c.sitesClientsNAC
}

// SitesClientsWan returns the sitesClientsWan instance of the client.
func (c *client) SitesClientsWan() *SitesClientsWan {
    return &c.sitesClientsWan
}

// SitesClientsWired returns the sitesClientsWired instance of the client.
func (c *client) SitesClientsWired() *SitesClientsWired {
    return &c.sitesClientsWired
}

// SitesClientsWireless returns the sitesClientsWireless instance of the client.
func (c *client) SitesClientsWireless() *SitesClientsWireless {
    return &c.sitesClientsWireless
}

// SitesDevices returns the sitesDevices instance of the client.
func (c *client) SitesDevices() *SitesDevices {
    return &c.sitesDevices
}

// SitesDevicesWireless returns the sitesDevicesWireless instance of the client.
func (c *client) SitesDevicesWireless() *SitesDevicesWireless {
    return &c.sitesDevicesWireless
}

// SitesDevicesOthers returns the sitesDevicesOthers instance of the client.
func (c *client) SitesDevicesOthers() *SitesDevicesOthers {
    return &c.sitesDevicesOthers
}

// SitesDevicesWired returns the sitesDevicesWired instance of the client.
func (c *client) SitesDevicesWired() *SitesDevicesWired {
    return &c.sitesDevicesWired
}

// SitesDevicesWiredVirtualChassis returns the sitesDevicesWiredVirtualChassis instance of the client.
func (c *client) SitesDevicesWiredVirtualChassis() *SitesDevicesWiredVirtualChassis {
    return &c.sitesDevicesWiredVirtualChassis
}

// SitesDevicesWANCluster returns the sitesDevicesWANCluster instance of the client.
func (c *client) SitesDevicesWANCluster() *SitesDevicesWANCluster {
    return &c.sitesDevicesWANCluster
}

// SitesDeviceProfiles returns the sitesDeviceProfiles instance of the client.
func (c *client) SitesDeviceProfiles() *SitesDeviceProfiles {
    return &c.sitesDeviceProfiles
}

// SitesEvents returns the sitesEvents instance of the client.
func (c *client) SitesEvents() *SitesEvents {
    return &c.sitesEvents
}

// SitesEVPNTopologies returns the sitesEVPNTopologies instance of the client.
func (c *client) SitesEVPNTopologies() *SitesEVPNTopologies {
    return &c.sitesEVPNTopologies
}

// SitesGatewayTemplates returns the sitesGatewayTemplates instance of the client.
func (c *client) SitesGatewayTemplates() *SitesGatewayTemplates {
    return &c.sitesGatewayTemplates
}

// SitesGuests returns the sitesGuests instance of the client.
func (c *client) SitesGuests() *SitesGuests {
    return &c.sitesGuests
}

// SitesInsights returns the sitesInsights instance of the client.
func (c *client) SitesInsights() *SitesInsights {
    return &c.sitesInsights
}

// SitesJSE returns the sitesJSE instance of the client.
func (c *client) SitesJSE() *SitesJSE {
    return &c.sitesJSE
}

// SitesLicenses returns the sitesLicenses instance of the client.
func (c *client) SitesLicenses() *SitesLicenses {
    return &c.sitesLicenses
}

// SitesLocation returns the sitesLocation instance of the client.
func (c *client) SitesLocation() *SitesLocation {
    return &c.sitesLocation
}

// SitesMaps returns the sitesMaps instance of the client.
func (c *client) SitesMaps() *SitesMaps {
    return &c.sitesMaps
}

// SitesMapsAutoPlacement returns the sitesMapsAutoPlacement instance of the client.
func (c *client) SitesMapsAutoPlacement() *SitesMapsAutoPlacement {
    return &c.sitesMapsAutoPlacement
}

// SitesMapsAutoZone returns the sitesMapsAutoZone instance of the client.
func (c *client) SitesMapsAutoZone() *SitesMapsAutoZone {
    return &c.sitesMapsAutoZone
}

// SitesMxEdges returns the sitesMxEdges instance of the client.
func (c *client) SitesMxEdges() *SitesMxEdges {
    return &c.sitesMxEdges
}

// SitesNetworkTemplates returns the sitesNetworkTemplates instance of the client.
func (c *client) SitesNetworkTemplates() *SitesNetworkTemplates {
    return &c.sitesNetworkTemplates
}

// SitesNetworks returns the sitesNetworks instance of the client.
func (c *client) SitesNetworks() *SitesNetworks {
    return &c.sitesNetworks
}

// SitesPsks returns the sitesPsks instance of the client.
func (c *client) SitesPsks() *SitesPsks {
    return &c.sitesPsks
}

// SitesRFTemplates returns the sitesRFTemplates instance of the client.
func (c *client) SitesRFTemplates() *SitesRFTemplates {
    return &c.sitesRFTemplates
}

// SitesRfdiags returns the sitesRfdiags instance of the client.
func (c *client) SitesRfdiags() *SitesRfdiags {
    return &c.sitesRfdiags
}

// SitesRogues returns the sitesRogues instance of the client.
func (c *client) SitesRogues() *SitesRogues {
    return &c.sitesRogues
}

// SitesRRM returns the sitesRRM instance of the client.
func (c *client) SitesRRM() *SitesRRM {
    return &c.sitesRRM
}

// SitesRSSIZones returns the sitesRSSIZones instance of the client.
func (c *client) SitesRSSIZones() *SitesRSSIZones {
    return &c.sitesRSSIZones
}

// SitesServices returns the sitesServices instance of the client.
func (c *client) SitesServices() *SitesServices {
    return &c.sitesServices
}

// SitesServicePolicies returns the sitesServicePolicies instance of the client.
func (c *client) SitesServicePolicies() *SitesServicePolicies {
    return &c.sitesServicePolicies
}

// SitesSetting returns the sitesSetting instance of the client.
func (c *client) SitesSetting() *SitesSetting {
    return &c.sitesSetting
}

// SitesSiteTemplates returns the sitesSiteTemplates instance of the client.
func (c *client) SitesSiteTemplates() *SitesSiteTemplates {
    return &c.sitesSiteTemplates
}

// SitesSkyatp returns the sitesSkyatp instance of the client.
func (c *client) SitesSkyatp() *SitesSkyatp {
    return &c.sitesSkyatp
}

// SitesSLEs returns the sitesSLEs instance of the client.
func (c *client) SitesSLEs() *SitesSLEs {
    return &c.sitesSLEs
}

// SitesSyntheticTests returns the sitesSyntheticTests instance of the client.
func (c *client) SitesSyntheticTests() *SitesSyntheticTests {
    return &c.sitesSyntheticTests
}

// SitesUISettings returns the sitesUISettings instance of the client.
func (c *client) SitesUISettings() *SitesUISettings {
    return &c.sitesUISettings
}

// SitesVBeacons returns the sitesVBeacons instance of the client.
func (c *client) SitesVBeacons() *SitesVBeacons {
    return &c.sitesVBeacons
}

// SitesVPNs returns the sitesVPNs instance of the client.
func (c *client) SitesVPNs() *SitesVPNs {
    return &c.sitesVPNs
}

// SitesWANUsages returns the sitesWANUsages instance of the client.
func (c *client) SitesWANUsages() *SitesWANUsages {
    return &c.sitesWANUsages
}

// SitesWebhooks returns the sitesWebhooks instance of the client.
func (c *client) SitesWebhooks() *SitesWebhooks {
    return &c.sitesWebhooks
}

// SitesWlans returns the sitesWlans instance of the client.
func (c *client) SitesWlans() *SitesWlans {
    return &c.sitesWlans
}

// SitesWxRules returns the sitesWxRules instance of the client.
func (c *client) SitesWxRules() *SitesWxRules {
    return &c.sitesWxRules
}

// SitesWxTags returns the sitesWxTags instance of the client.
func (c *client) SitesWxTags() *SitesWxTags {
    return &c.sitesWxTags
}

// SitesWxTunnels returns the sitesWxTunnels instance of the client.
func (c *client) SitesWxTunnels() *SitesWxTunnels {
    return &c.sitesWxTunnels
}

// SitesZones returns the sitesZones instance of the client.
func (c *client) SitesZones() *SitesZones {
    return &c.sitesZones
}

// SitesStats returns the sitesStats instance of the client.
func (c *client) SitesStats() *SitesStats {
    return &c.sitesStats
}

// SitesStatsApps returns the sitesStatsApps instance of the client.
func (c *client) SitesStatsApps() *SitesStatsApps {
    return &c.sitesStatsApps
}

// SitesStatsAssets returns the sitesStatsAssets instance of the client.
func (c *client) SitesStatsAssets() *SitesStatsAssets {
    return &c.sitesStatsAssets
}

// SitesStatsBeacons returns the sitesStatsBeacons instance of the client.
func (c *client) SitesStatsBeacons() *SitesStatsBeacons {
    return &c.sitesStatsBeacons
}

// SitesStatsBGPPeers returns the sitesStatsBGPPeers instance of the client.
func (c *client) SitesStatsBGPPeers() *SitesStatsBGPPeers {
    return &c.sitesStatsBGPPeers
}

// SitesStatsCalls returns the sitesStatsCalls instance of the client.
func (c *client) SitesStatsCalls() *SitesStatsCalls {
    return &c.sitesStatsCalls
}

// SitesStatsClientsWireless returns the sitesStatsClientsWireless instance of the client.
func (c *client) SitesStatsClientsWireless() *SitesStatsClientsWireless {
    return &c.sitesStatsClientsWireless
}

// SitesStatsClientsSDK returns the sitesStatsClientsSDK instance of the client.
func (c *client) SitesStatsClientsSDK() *SitesStatsClientsSDK {
    return &c.sitesStatsClientsSDK
}

// SitesStatsDevices returns the sitesStatsDevices instance of the client.
func (c *client) SitesStatsDevices() *SitesStatsDevices {
    return &c.sitesStatsDevices
}

// SitesStatsMxEdges returns the sitesStatsMxEdges instance of the client.
func (c *client) SitesStatsMxEdges() *SitesStatsMxEdges {
    return &c.sitesStatsMxEdges
}

// SitesStatsPorts returns the sitesStatsPorts instance of the client.
func (c *client) SitesStatsPorts() *SitesStatsPorts {
    return &c.sitesStatsPorts
}

// SitesStatsWxRules returns the sitesStatsWxRules instance of the client.
func (c *client) SitesStatsWxRules() *SitesStatsWxRules {
    return &c.sitesStatsWxRules
}

// SitesStatsZones returns the sitesStatsZones instance of the client.
func (c *client) SitesStatsZones() *SitesStatsZones {
    return &c.sitesStatsZones
}

// SitesStatsDiscoveredSwitches returns the sitesStatsDiscoveredSwitches instance of the client.
func (c *client) SitesStatsDiscoveredSwitches() *SitesStatsDiscoveredSwitches {
    return &c.sitesStatsDiscoveredSwitches
}

// ConstantsDefinitions returns the constantsDefinitions instance of the client.
func (c *client) ConstantsDefinitions() *ConstantsDefinitions {
    return &c.constantsDefinitions
}

// ConstantsEvents returns the constantsEvents instance of the client.
func (c *client) ConstantsEvents() *ConstantsEvents {
    return &c.constantsEvents
}

// ConstantsModels returns the constantsModels instance of the client.
func (c *client) ConstantsModels() *ConstantsModels {
    return &c.constantsModels
}

// SelfAccount returns the selfAccount instance of the client.
func (c *client) SelfAccount() *SelfAccount {
    return &c.selfAccount
}

// SelfAPIToken returns the selfAPIToken instance of the client.
func (c *client) SelfAPIToken() *SelfAPIToken {
    return &c.selfAPIToken
}

// SelfOAuth2 returns the selfOAuth2 instance of the client.
func (c *client) SelfOAuth2() *SelfOAuth2 {
    return &c.selfOAuth2
}

// SelfMFA returns the selfMFA instance of the client.
func (c *client) SelfMFA() *SelfMFA {
    return &c.selfMFA
}

// SelfAlarms returns the selfAlarms instance of the client.
func (c *client) SelfAlarms() *SelfAlarms {
    return &c.selfAlarms
}

// SelfAuditLogs returns the selfAuditLogs instance of the client.
func (c *client) SelfAuditLogs() *SelfAuditLogs {
    return &c.selfAuditLogs
}


// UtilitiesCommon returns the utilitiesCommon instance of the client.
func (c *client) UtilitiesCommon() *UtilitiesCommon {
    return &c.utilitiesCommon
}

// UtilitiesWAN returns the utilitiesWAN instance of the client.
func (c *client) UtilitiesWAN() *UtilitiesWAN {
    return &c.utilitiesWAN
}

// UtilitiesLAN returns the utilitiesLAN instance of the client.
func (c *client) UtilitiesLAN() *UtilitiesLAN {
    return &c.utilitiesLAN
}

// UtilitiesWiFi returns the utilitiesWiFi instance of the client.
func (c *client) UtilitiesWiFi() *UtilitiesWiFi {
    return &c.utilitiesWiFi
}

// UtilitiesPCAPs returns the utilitiesPCAPs instance of the client.
func (c *client) UtilitiesPCAPs() *UtilitiesPCAPs {
    return &c.utilitiesPCAPs
}

// UtilitiesLocation returns the utilitiesLocation instance of the client.
func (c *client) UtilitiesLocation() *UtilitiesLocation {
    return &c.utilitiesLocation
}

// UtilitiesMxEdge returns the utilitiesMxEdge instance of the client.
func (c *client) UtilitiesMxEdge() *UtilitiesMxEdge {
    return &c.utilitiesMxEdge
}

// UtilitiesUpgrade returns the utilitiesUpgrade instance of the client.
func (c *client) UtilitiesUpgrade() *UtilitiesUpgrade {
    return &c.utilitiesUpgrade
}

// UserAgent returns the userAgent instance of the client.
func (c *client) UserAgent() *string {
    return &c.userAgent
}

// GetCallBuilder returns the CallBuilderFactory used by the client.
func (c *client) GetCallBuilder() https.CallBuilderFactory {
    return c.callBuilderFactory
}

// getBaseUri returns the base URI based on the server and configuration.
func getBaseUri(
    server Server,
    configuration Configuration) string {
    if configuration.Environment() == Environment(AWS_STAGING) {
        if server == Server(APIHOST) {
            return "https://api.mistsys.com"
        }
    }
    if configuration.Environment() == Environment(MIST_GLOBAL_01) {
        if server == Server(APIHOST) {
            return "https://api.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_GLOBAL_02) {
        if server == Server(APIHOST) {
            return "https://api.gc1.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_GLOBAL_03) {
        if server == Server(APIHOST) {
            return "https://api.ac2.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_GLOBAL_04) {
        if server == Server(APIHOST) {
            return "https://api.gc2.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_EMEA_01) {
        if server == Server(APIHOST) {
            return "https://api.eu.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_EMEA_02) {
        if server == Server(APIHOST) {
            return "https://api.gc3.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_EMEA_03) {
        if server == Server(APIHOST) {
            return "https://api.ac6.mist.com"
        }
    }
    if configuration.Environment() == Environment(MIST_APAC_01) {
        if server == Server(APIHOST) {
            return "https://api.ac5.mist.com"
        }
    }
    return "TODO: Select a valid server."
}

// clientOptions is a function type representing options for the client.
type clientOptions func(cb https.CallBuilder)

// callBuilderHandler creates the call builder factory with various options.
func callBuilderHandler(
    baseUrlProvider func(server string) string,
    auth map[string]https.AuthInterface,
    httpClient https.HttpClient,
    retryConfig RetryConfiguration,
    arraySerializationOption https.ArraySerializationOption,
    opts ...clientOptions) https.CallBuilderFactory {
    callBuilderFactory := https.CreateCallBuilderFactory(baseUrlProvider, auth, httpClient, retryConfig, arraySerializationOption)
    return tap(callBuilderFactory, opts...)
}

// tap is a utility function to apply client options to the call builder factory.
func tap(
    callBuilderFactory https.CallBuilderFactory,
    opts ...clientOptions) https.CallBuilderFactory {
    return func(ctx context.Context, httpMethod, path string) https.CallBuilder {
    	callBuilder := callBuilderFactory(ctx, httpMethod, path)
    	for _, opt := range opts {
    		opt(callBuilder)
    	}
    	return callBuilder
    }
}

// withUserAgent is an option to add a user agent header to the HTTP request.
func withUserAgent(userAgent string) clientOptions {
    f := func(request *http.Request) *http.Request {
    	request.Header.Set("user-agent", userAgent)
    	return request
    }
    return func(cb https.CallBuilder) {
    	cb.InterceptRequest(f)
    }
}

// withLogger will add a new instance of SdkLoggerInterface to callBuilder.
func withLogger(loggerConfiguration LoggerConfiguration) clientOptions {
    return func(cb https.CallBuilder) {
        cb.Logger(NewSdkLogger(loggerConfiguration))
    }
}
