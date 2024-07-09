package mistapigo

import (
	"context"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
)

// Client is an interface representing the main client for accessing configuration and controllers.
type ClientInterface interface {
	Configuration() Configuration
	CloneWithConfiguration(options ...ConfigurationOptions) ClientInterface
	Orgs() *Orgs
	OrgsAPTemplates() *OrgsAPTemplates
	OrgsAPITokens() *OrgsAPITokens
	OrgsAdmins() *OrgsAdmins
	OrgsAlarmTemplates() *OrgsAlarmTemplates
	OrgsAlarms() *OrgsAlarms
	OrgsAssetFilters() *OrgsAssetFilters
	OrgsAssets() *OrgsAssets
	OrgsCRL() *OrgsCRL
	OrgsCert() *OrgsCert
	OrgsClientsMarvis() *OrgsClientsMarvis
	OrgsClientsNAC() *OrgsClientsNAC
	OrgsClientsSDK() *OrgsClientsSDK
	OrgsClientsWan() *OrgsClientsWan
	OrgsClientsWired() *OrgsClientsWired
	OrgsClientsWireless() *OrgsClientsWireless
	OrgsCradlepoint() *OrgsCradlepoint
	OrgsDeviceProfiles() *OrgsDeviceProfiles
	OrgsDevices() *OrgsDevices
	OrgsDevicesOthers() *OrgsDevicesOthers
	OrgsDevicesSSR() *OrgsDevicesSSR
	OrgsDevicesWANCluster() *OrgsDevicesWANCluster
	OrgsDevicesStats() *OrgsDevicesStats
	OrgsDevicesZscaler() *OrgsDevicesZscaler
	OrgsEVPNTopologies() *OrgsEVPNTopologies
	OrgsGatewayTemplates() *OrgsGatewayTemplates
	OrgsGuests() *OrgsGuests
	OrgsIDPProfiles() *OrgsIDPProfiles
	OrgsInventory() *OrgsInventory
	OrgsJSE() *OrgsJSE
	OrgsJSI() *OrgsJSI
	OrgsLicenses() *OrgsLicenses
	OrgsLinkedApplications() *OrgsLinkedApplications
	OrgsLogs() *OrgsLogs
	OrgsMaps() *OrgsMaps
	OrgsMarvis() *OrgsMarvis
	OrgsMxClusters() *OrgsMxClusters
	OrgsMxEdges() *OrgsMxEdges
	OrgsMxTunnels() *OrgsMxTunnels
	OrgsNACCRL() *OrgsNACCRL
	OrgsNACIDP() *OrgsNACIDP
	OrgsNACPortals() *OrgsNACPortals
	OrgsNACRules() *OrgsNACRules
	OrgsNACTags() *OrgsNACTags
	OrgsNetworkTemplates() *OrgsNetworkTemplates
	OrgsNetworks() *OrgsNetworks
	OrgsPremiumAnalytics() *OrgsPremiumAnalytics
	OrgsPskPortals() *OrgsPskPortals
	OrgsPsks() *OrgsPsks
	OrgsRFTemplates() *OrgsRFTemplates
	OrgsSDKInvites() *OrgsSDKInvites
	OrgsSDKTemplates() *OrgsSDKTemplates
	OrgsSLEs() *OrgsSLEs
	OrgsSSO() *OrgsSSO
	OrgsSSORoles() *OrgsSSORoles
	OrgsSecPolicies() *OrgsSecPolicies
	OrgsServicePolicies() *OrgsServicePolicies
	OrgsServices() *OrgsServices
	OrgsSetting() *OrgsSetting
	OrgsSiteTemplates() *OrgsSiteTemplates
	OrgsSitegroups() *OrgsSitegroups
	OrgsSites() *OrgsSites
	OrgsSubscriptions() *OrgsSubscriptions
	OrgsTickets() *OrgsTickets
	OrgsTunnels() *OrgsTunnels
	OrgsUserMACs() *OrgsUserMACs
	OrgsVPNs() *OrgsVPNs
	OrgsVars() *OrgsVars
	OrgsWLANTemplates() *OrgsWLANTemplates
	OrgsWebhooks() *OrgsWebhooks
	OrgsWlans() *OrgsWlans
	OrgsWxRules() *OrgsWxRules
	OrgsWxTags() *OrgsWxTags
	OrgsWxTunnels() *OrgsWxTunnels
	UtilitiesPCAPs() *UtilitiesPCAPs
	UtilitiesUpgrade() *UtilitiesUpgrade
	UtilitiesWiFi() *UtilitiesWiFi
	Installer() *Installer
	Admins() *Admins
	AdminsLogin() *AdminsLogin
	AdminsLoginOAuth2() *AdminsLoginOAuth2
	AdminsLogout() *AdminsLogout
	AdminsLookup() *AdminsLookup
	AdminsRecoverPassword() *AdminsRecoverPassword
	ConstantsDefinitions() *ConstantsDefinitions
	ConstantsEvents() *ConstantsEvents
	ConstantsMisc() *ConstantsMisc
	ConstantsModels() *ConstantsModels
	SelfAPIToken() *SelfAPIToken
	SelfAccount() *SelfAccount
	SelfAlarms() *SelfAlarms
	SelfAuditLogs() *SelfAuditLogs
	SelfMFA() *SelfMFA
	SelfOAuth2() *SelfOAuth2
	MSPs() *MSPs
	MSPsAdmins() *MSPsAdmins
	MSPsInventory() *MSPsInventory
	MSPsLicenses() *MSPsLicenses
	MSPsLogo() *MSPsLogo
	MSPsLogs() *MSPsLogs
	MSPsMarvis() *MSPsMarvis
	MSPsOrgGroups() *MSPsOrgGroups
	MSPsOrgs() *MSPsOrgs
	MSPsSLEs() *MSPsSLEs
	MSPsSSO() *MSPsSSO
	MSPsSSORoles() *MSPsSSORoles
	MSPsTickets() *MSPsTickets
	Sites() *Sites
	SitesAPTemplates() *SitesAPTemplates
	SitesAlarms() *SitesAlarms
	SitesAnomaly() *SitesAnomaly
	SitesApplications() *SitesApplications
	SitesAssetFilters() *SitesAssetFilters
	SitesAssets() *SitesAssets
	SitesBeacons() *SitesBeacons
	SitesCalls() *SitesCalls
	SitesClientsNAC() *SitesClientsNAC
	SitesClientsSDK() *SitesClientsSDK
	SitesClientsWan() *SitesClientsWan
	SitesClientsWired() *SitesClientsWired
	SitesClientsWireless() *SitesClientsWireless
	SitesDeviceProfiles() *SitesDeviceProfiles
	SitesDevices() *SitesDevices
	SitesDevicesDiscoveredSwitches() *SitesDevicesDiscoveredSwitches
	SitesDevicesOthers() *SitesDevicesOthers
	SitesDevicesWANCluster() *SitesDevicesWANCluster
	SitesDevicesWired() *SitesDevicesWired
	SitesDevicesWiredVirtualChassis() *SitesDevicesWiredVirtualChassis
	SitesDevicesWireless() *SitesDevicesWireless
	SitesDevicesStats() *SitesDevicesStats
	SitesEVPNTopologies() *SitesEVPNTopologies
	SitesEvents() *SitesEvents
	SitesGatewayTemplates() *SitesGatewayTemplates
	SitesGuests() *SitesGuests
	SitesInsights() *SitesInsights
	SitesJSE() *SitesJSE
	SitesLicenses() *SitesLicenses
	SitesLocation() *SitesLocation
	SitesMaps() *SitesMaps
	SitesMapsAutoOrientation() *SitesMapsAutoOrientation
	SitesMapsAutoPlacement() *SitesMapsAutoPlacement
	SitesMxEdges() *SitesMxEdges
	SitesNetworkTemplates() *SitesNetworkTemplates
	SitesNetworks() *SitesNetworks
	SitesPsks() *SitesPsks
	SitesRFTemplates() *SitesRFTemplates
	SitesRRM() *SitesRRM
	SitesRSSIZones() *SitesRSSIZones
	SitesRfdiags() *SitesRfdiags
	SitesRogues() *SitesRogues
	SitesSLEs() *SitesSLEs
	SitesServicePolicies() *SitesServicePolicies
	SitesServices() *SitesServices
	SitesSetting() *SitesSetting
	SitesSiteTemplates() *SitesSiteTemplates
	SitesSkyatp() *SitesSkyatp
	SitesSyntheticTests() *SitesSyntheticTests
	SitesUISettings() *SitesUISettings
	SitesVPNs() *SitesVPNs
	SitesWANUsages() *SitesWANUsages
	SitesWebhooks() *SitesWebhooks
	SitesWlans() *SitesWlans
	SitesWxRules() *SitesWxRules
	SitesWxTags() *SitesWxTags
	SitesWxTunnels() *SitesWxTunnels
	SitesZones() *SitesZones
	SitesVBeacons() *SitesVBeacons
	UtilitiesCommon() *UtilitiesCommon
	UtilitiesLAN() *UtilitiesLAN
	UtilitiesLocation() *UtilitiesLocation
	UtilitiesMxEdge() *UtilitiesMxEdge
	UtilitiesWAN() *UtilitiesWAN
	UserAgent() *string
}

// client is an implementation of the Client interface.
type client struct {
	callBuilderFactory              https.CallBuilderFactory
	configuration                   Configuration
	userAgent                       string
	orgs                            Orgs
	orgsAPTemplates                 OrgsAPTemplates
	orgsAPITokens                   OrgsAPITokens
	orgsAdmins                      OrgsAdmins
	orgsAlarmTemplates              OrgsAlarmTemplates
	orgsAlarms                      OrgsAlarms
	orgsAssetFilters                OrgsAssetFilters
	orgsAssets                      OrgsAssets
	orgsCRL                         OrgsCRL
	orgsCert                        OrgsCert
	orgsClientsMarvis               OrgsClientsMarvis
	orgsClientsNAC                  OrgsClientsNAC
	orgsClientsSDK                  OrgsClientsSDK
	orgsClientsWan                  OrgsClientsWan
	orgsClientsWired                OrgsClientsWired
	orgsClientsWireless             OrgsClientsWireless
	orgsCradlepoint                 OrgsCradlepoint
	orgsDeviceProfiles              OrgsDeviceProfiles
	orgsDevices                     OrgsDevices
	orgsDevicesOthers               OrgsDevicesOthers
	orgsDevicesSSR                  OrgsDevicesSSR
	orgsDevicesWANCluster           OrgsDevicesWANCluster
	orgsDevicesStats                OrgsDevicesStats
	orgsDevicesZscaler              OrgsDevicesZscaler
	orgsEVPNTopologies              OrgsEVPNTopologies
	orgsGatewayTemplates            OrgsGatewayTemplates
	orgsGuests                      OrgsGuests
	orgsIDPProfiles                 OrgsIDPProfiles
	orgsInventory                   OrgsInventory
	orgsJSE                         OrgsJSE
	orgsJSI                         OrgsJSI
	orgsLicenses                    OrgsLicenses
	orgsLinkedApplications          OrgsLinkedApplications
	orgsLogs                        OrgsLogs
	orgsMaps                        OrgsMaps
	orgsMarvis                      OrgsMarvis
	orgsMxClusters                  OrgsMxClusters
	orgsMxEdges                     OrgsMxEdges
	orgsMxTunnels                   OrgsMxTunnels
	orgsNACCRL                      OrgsNACCRL
	orgsNACIDP                      OrgsNACIDP
	orgsNACPortals                  OrgsNACPortals
	orgsNACRules                    OrgsNACRules
	orgsNACTags                     OrgsNACTags
	orgsNetworkTemplates            OrgsNetworkTemplates
	orgsNetworks                    OrgsNetworks
	orgsPremiumAnalytics            OrgsPremiumAnalytics
	orgsPskPortals                  OrgsPskPortals
	orgsPsks                        OrgsPsks
	orgsRFTemplates                 OrgsRFTemplates
	orgsSDKInvites                  OrgsSDKInvites
	orgsSDKTemplates                OrgsSDKTemplates
	orgsSLEs                        OrgsSLEs
	orgsSSO                         OrgsSSO
	orgsSSORoles                    OrgsSSORoles
	orgsSecPolicies                 OrgsSecPolicies
	orgsServicePolicies             OrgsServicePolicies
	orgsServices                    OrgsServices
	orgsSetting                     OrgsSetting
	orgsSiteTemplates               OrgsSiteTemplates
	orgsSitegroups                  OrgsSitegroups
	orgsSites                       OrgsSites
	orgsSubscriptions               OrgsSubscriptions
	orgsTickets                     OrgsTickets
	orgsTunnels                     OrgsTunnels
	orgsUserMACs                    OrgsUserMACs
	orgsVPNs                        OrgsVPNs
	orgsVars                        OrgsVars
	orgsWLANTemplates               OrgsWLANTemplates
	orgsWebhooks                    OrgsWebhooks
	orgsWlans                       OrgsWlans
	orgsWxRules                     OrgsWxRules
	orgsWxTags                      OrgsWxTags
	orgsWxTunnels                   OrgsWxTunnels
	utilitiesPCAPs                  UtilitiesPCAPs
	utilitiesUpgrade                UtilitiesUpgrade
	utilitiesWiFi                   UtilitiesWiFi
	installer                       Installer
	admins                          Admins
	adminsLogin                     AdminsLogin
	adminsLoginOAuth2               AdminsLoginOAuth2
	adminsLogout                    AdminsLogout
	adminsLookup                    AdminsLookup
	adminsRecoverPassword           AdminsRecoverPassword
	constantsDefinitions            ConstantsDefinitions
	constantsEvents                 ConstantsEvents
	constantsMisc                   ConstantsMisc
	constantsModels                 ConstantsModels
	selfAPIToken                    SelfAPIToken
	selfAccount                     SelfAccount
	selfAlarms                      SelfAlarms
	selfAuditLogs                   SelfAuditLogs
	selfMFA                         SelfMFA
	selfOAuth2                      SelfOAuth2
	mSPs                            MSPs
	mSPsAdmins                      MSPsAdmins
	mSPsInventory                   MSPsInventory
	mSPsLicenses                    MSPsLicenses
	mSPsLogo                        MSPsLogo
	mSPsLogs                        MSPsLogs
	mSPsMarvis                      MSPsMarvis
	mSPsOrgGroups                   MSPsOrgGroups
	mSPsOrgs                        MSPsOrgs
	mSPsSLEs                        MSPsSLEs
	mSPsSSO                         MSPsSSO
	mSPsSSORoles                    MSPsSSORoles
	mSPsTickets                     MSPsTickets
	sites                           Sites
	sitesAPTemplates                SitesAPTemplates
	sitesAlarms                     SitesAlarms
	sitesAnomaly                    SitesAnomaly
	sitesApplications               SitesApplications
	sitesAssetFilters               SitesAssetFilters
	sitesAssets                     SitesAssets
	sitesBeacons                    SitesBeacons
	sitesCalls                      SitesCalls
	sitesClientsNAC                 SitesClientsNAC
	sitesClientsSDK                 SitesClientsSDK
	sitesClientsWan                 SitesClientsWan
	sitesClientsWired               SitesClientsWired
	sitesClientsWireless            SitesClientsWireless
	sitesDeviceProfiles             SitesDeviceProfiles
	sitesDevices                    SitesDevices
	sitesDevicesDiscoveredSwitches  SitesDevicesDiscoveredSwitches
	sitesDevicesOthers              SitesDevicesOthers
	sitesDevicesWANCluster          SitesDevicesWANCluster
	sitesDevicesWired               SitesDevicesWired
	sitesDevicesWiredVirtualChassis SitesDevicesWiredVirtualChassis
	sitesDevicesWireless            SitesDevicesWireless
	sitesDevicesStats               SitesDevicesStats
	sitesEVPNTopologies             SitesEVPNTopologies
	sitesEvents                     SitesEvents
	sitesGatewayTemplates           SitesGatewayTemplates
	sitesGuests                     SitesGuests
	sitesInsights                   SitesInsights
	sitesJSE                        SitesJSE
	sitesLicenses                   SitesLicenses
	sitesLocation                   SitesLocation
	sitesMaps                       SitesMaps
	sitesMapsAutoOrientation        SitesMapsAutoOrientation
	sitesMapsAutoPlacement          SitesMapsAutoPlacement
	sitesMxEdges                    SitesMxEdges
	sitesNetworkTemplates           SitesNetworkTemplates
	sitesNetworks                   SitesNetworks
	sitesPsks                       SitesPsks
	sitesRFTemplates                SitesRFTemplates
	sitesRRM                        SitesRRM
	sitesRSSIZones                  SitesRSSIZones
	sitesRfdiags                    SitesRfdiags
	sitesRogues                     SitesRogues
	sitesSLEs                       SitesSLEs
	sitesServicePolicies            SitesServicePolicies
	sitesServices                   SitesServices
	sitesSetting                    SitesSetting
	sitesSiteTemplates              SitesSiteTemplates
	sitesSkyatp                     SitesSkyatp
	sitesSyntheticTests             SitesSyntheticTests
	sitesUISettings                 SitesUISettings
	sitesVPNs                       SitesVPNs
	sitesWANUsages                  SitesWANUsages
	sitesWebhooks                   SitesWebhooks
	sitesWlans                      SitesWlans
	sitesWxRules                    SitesWxRules
	sitesWxTags                     SitesWxTags
	sitesWxTunnels                  SitesWxTunnels
	sitesZones                      SitesZones
	sitesVBeacons                   SitesVBeacons
	utilitiesCommon                 UtilitiesCommon
	utilitiesLAN                    UtilitiesLAN
	utilitiesLocation               UtilitiesLocation
	utilitiesMxEdge                 UtilitiesMxEdge
	utilitiesWAN                    UtilitiesWAN
}

// NewClient is the constructor for creating a new client instance.
// It takes a Configuration object as a parameter and returns the Client interface.
func NewClient(configuration Configuration) ClientInterface {
	client := &client{
		configuration: configuration,
	}

	client.userAgent = utilities.UpdateUserAgent("SDK 2024.2.1")
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
	client.orgs = *NewOrgs(*baseController)
	client.orgsAPTemplates = *NewOrgsAPTemplates(*baseController)
	client.orgsAPITokens = *NewOrgsAPITokens(*baseController)
	client.orgsAdmins = *NewOrgsAdmins(*baseController)
	client.orgsAlarmTemplates = *NewOrgsAlarmTemplates(*baseController)
	client.orgsAlarms = *NewOrgsAlarms(*baseController)
	client.orgsAssetFilters = *NewOrgsAssetFilters(*baseController)
	client.orgsAssets = *NewOrgsAssets(*baseController)
	client.orgsCRL = *NewOrgsCRL(*baseController)
	client.orgsCert = *NewOrgsCert(*baseController)
	client.orgsClientsMarvis = *NewOrgsClientsMarvis(*baseController)
	client.orgsClientsNAC = *NewOrgsClientsNAC(*baseController)
	client.orgsClientsSDK = *NewOrgsClientsSDK(*baseController)
	client.orgsClientsWan = *NewOrgsClientsWan(*baseController)
	client.orgsClientsWired = *NewOrgsClientsWired(*baseController)
	client.orgsClientsWireless = *NewOrgsClientsWireless(*baseController)
	client.orgsCradlepoint = *NewOrgsCradlepoint(*baseController)
	client.orgsDeviceProfiles = *NewOrgsDeviceProfiles(*baseController)
	client.orgsDevices = *NewOrgsDevices(*baseController)
	client.orgsDevicesOthers = *NewOrgsDevicesOthers(*baseController)
	client.orgsDevicesSSR = *NewOrgsDevicesSSR(*baseController)
	client.orgsDevicesWANCluster = *NewOrgsDevicesWANCluster(*baseController)
	client.orgsDevicesStats = *NewOrgsDevicesStats(*baseController)
	client.orgsDevicesZscaler = *NewOrgsDevicesZscaler(*baseController)
	client.orgsEVPNTopologies = *NewOrgsEVPNTopologies(*baseController)
	client.orgsGatewayTemplates = *NewOrgsGatewayTemplates(*baseController)
	client.orgsGuests = *NewOrgsGuests(*baseController)
	client.orgsIDPProfiles = *NewOrgsIDPProfiles(*baseController)
	client.orgsInventory = *NewOrgsInventory(*baseController)
	client.orgsJSE = *NewOrgsJSE(*baseController)
	client.orgsJSI = *NewOrgsJSI(*baseController)
	client.orgsLicenses = *NewOrgsLicenses(*baseController)
	client.orgsLinkedApplications = *NewOrgsLinkedApplications(*baseController)
	client.orgsLogs = *NewOrgsLogs(*baseController)
	client.orgsMaps = *NewOrgsMaps(*baseController)
	client.orgsMarvis = *NewOrgsMarvis(*baseController)
	client.orgsMxClusters = *NewOrgsMxClusters(*baseController)
	client.orgsMxEdges = *NewOrgsMxEdges(*baseController)
	client.orgsMxTunnels = *NewOrgsMxTunnels(*baseController)
	client.orgsNACCRL = *NewOrgsNACCRL(*baseController)
	client.orgsNACIDP = *NewOrgsNACIDP(*baseController)
	client.orgsNACPortals = *NewOrgsNACPortals(*baseController)
	client.orgsNACRules = *NewOrgsNACRules(*baseController)
	client.orgsNACTags = *NewOrgsNACTags(*baseController)
	client.orgsNetworkTemplates = *NewOrgsNetworkTemplates(*baseController)
	client.orgsNetworks = *NewOrgsNetworks(*baseController)
	client.orgsPremiumAnalytics = *NewOrgsPremiumAnalytics(*baseController)
	client.orgsPskPortals = *NewOrgsPskPortals(*baseController)
	client.orgsPsks = *NewOrgsPsks(*baseController)
	client.orgsRFTemplates = *NewOrgsRFTemplates(*baseController)
	client.orgsSDKInvites = *NewOrgsSDKInvites(*baseController)
	client.orgsSDKTemplates = *NewOrgsSDKTemplates(*baseController)
	client.orgsSLEs = *NewOrgsSLEs(*baseController)
	client.orgsSSO = *NewOrgsSSO(*baseController)
	client.orgsSSORoles = *NewOrgsSSORoles(*baseController)
	client.orgsSecPolicies = *NewOrgsSecPolicies(*baseController)
	client.orgsServicePolicies = *NewOrgsServicePolicies(*baseController)
	client.orgsServices = *NewOrgsServices(*baseController)
	client.orgsSetting = *NewOrgsSetting(*baseController)
	client.orgsSiteTemplates = *NewOrgsSiteTemplates(*baseController)
	client.orgsSitegroups = *NewOrgsSitegroups(*baseController)
	client.orgsSites = *NewOrgsSites(*baseController)
	client.orgsSubscriptions = *NewOrgsSubscriptions(*baseController)
	client.orgsTickets = *NewOrgsTickets(*baseController)
	client.orgsTunnels = *NewOrgsTunnels(*baseController)
	client.orgsUserMACs = *NewOrgsUserMACs(*baseController)
	client.orgsVPNs = *NewOrgsVPNs(*baseController)
	client.orgsVars = *NewOrgsVars(*baseController)
	client.orgsWLANTemplates = *NewOrgsWLANTemplates(*baseController)
	client.orgsWebhooks = *NewOrgsWebhooks(*baseController)
	client.orgsWlans = *NewOrgsWlans(*baseController)
	client.orgsWxRules = *NewOrgsWxRules(*baseController)
	client.orgsWxTags = *NewOrgsWxTags(*baseController)
	client.orgsWxTunnels = *NewOrgsWxTunnels(*baseController)
	client.utilitiesPCAPs = *NewUtilitiesPCAPs(*baseController)
	client.utilitiesUpgrade = *NewUtilitiesUpgrade(*baseController)
	client.utilitiesWiFi = *NewUtilitiesWiFi(*baseController)
	client.installer = *NewInstaller(*baseController)
	client.admins = *NewAdmins(*baseController)
	client.adminsLogin = *NewAdminsLogin(*baseController)
	client.adminsLoginOAuth2 = *NewAdminsLoginOAuth2(*baseController)
	client.adminsLogout = *NewAdminsLogout(*baseController)
	client.adminsLookup = *NewAdminsLookup(*baseController)
	client.adminsRecoverPassword = *NewAdminsRecoverPassword(*baseController)
	client.constantsDefinitions = *NewConstantsDefinitions(*baseController)
	client.constantsEvents = *NewConstantsEvents(*baseController)
	client.constantsMisc = *NewConstantsMisc(*baseController)
	client.constantsModels = *NewConstantsModels(*baseController)
	client.selfAPIToken = *NewSelfAPIToken(*baseController)
	client.selfAccount = *NewSelfAccount(*baseController)
	client.selfAlarms = *NewSelfAlarms(*baseController)
	client.selfAuditLogs = *NewSelfAuditLogs(*baseController)
	client.selfMFA = *NewSelfMFA(*baseController)
	client.selfOAuth2 = *NewSelfOAuth2(*baseController)
	client.mSPs = *NewMSPs(*baseController)
	client.mSPsAdmins = *NewMSPsAdmins(*baseController)
	client.mSPsInventory = *NewMSPsInventory(*baseController)
	client.mSPsLicenses = *NewMSPsLicenses(*baseController)
	client.mSPsLogo = *NewMSPsLogo(*baseController)
	client.mSPsLogs = *NewMSPsLogs(*baseController)
	client.mSPsMarvis = *NewMSPsMarvis(*baseController)
	client.mSPsOrgGroups = *NewMSPsOrgGroups(*baseController)
	client.mSPsOrgs = *NewMSPsOrgs(*baseController)
	client.mSPsSLEs = *NewMSPsSLEs(*baseController)
	client.mSPsSSO = *NewMSPsSSO(*baseController)
	client.mSPsSSORoles = *NewMSPsSSORoles(*baseController)
	client.mSPsTickets = *NewMSPsTickets(*baseController)
	client.sites = *NewSites(*baseController)
	client.sitesAPTemplates = *NewSitesAPTemplates(*baseController)
	client.sitesAlarms = *NewSitesAlarms(*baseController)
	client.sitesAnomaly = *NewSitesAnomaly(*baseController)
	client.sitesApplications = *NewSitesApplications(*baseController)
	client.sitesAssetFilters = *NewSitesAssetFilters(*baseController)
	client.sitesAssets = *NewSitesAssets(*baseController)
	client.sitesBeacons = *NewSitesBeacons(*baseController)
	client.sitesCalls = *NewSitesCalls(*baseController)
	client.sitesClientsNAC = *NewSitesClientsNAC(*baseController)
	client.sitesClientsSDK = *NewSitesClientsSDK(*baseController)
	client.sitesClientsWan = *NewSitesClientsWan(*baseController)
	client.sitesClientsWired = *NewSitesClientsWired(*baseController)
	client.sitesClientsWireless = *NewSitesClientsWireless(*baseController)
	client.sitesDeviceProfiles = *NewSitesDeviceProfiles(*baseController)
	client.sitesDevices = *NewSitesDevices(*baseController)
	client.sitesDevicesDiscoveredSwitches = *NewSitesDevicesDiscoveredSwitches(*baseController)
	client.sitesDevicesOthers = *NewSitesDevicesOthers(*baseController)
	client.sitesDevicesWANCluster = *NewSitesDevicesWANCluster(*baseController)
	client.sitesDevicesWired = *NewSitesDevicesWired(*baseController)
	client.sitesDevicesWiredVirtualChassis = *NewSitesDevicesWiredVirtualChassis(*baseController)
	client.sitesDevicesWireless = *NewSitesDevicesWireless(*baseController)
	client.sitesDevicesStats = *NewSitesDevicesStats(*baseController)
	client.sitesEVPNTopologies = *NewSitesEVPNTopologies(*baseController)
	client.sitesEvents = *NewSitesEvents(*baseController)
	client.sitesGatewayTemplates = *NewSitesGatewayTemplates(*baseController)
	client.sitesGuests = *NewSitesGuests(*baseController)
	client.sitesInsights = *NewSitesInsights(*baseController)
	client.sitesJSE = *NewSitesJSE(*baseController)
	client.sitesLicenses = *NewSitesLicenses(*baseController)
	client.sitesLocation = *NewSitesLocation(*baseController)
	client.sitesMaps = *NewSitesMaps(*baseController)
	client.sitesMapsAutoOrientation = *NewSitesMapsAutoOrientation(*baseController)
	client.sitesMapsAutoPlacement = *NewSitesMapsAutoPlacement(*baseController)
	client.sitesMxEdges = *NewSitesMxEdges(*baseController)
	client.sitesNetworkTemplates = *NewSitesNetworkTemplates(*baseController)
	client.sitesNetworks = *NewSitesNetworks(*baseController)
	client.sitesPsks = *NewSitesPsks(*baseController)
	client.sitesRFTemplates = *NewSitesRFTemplates(*baseController)
	client.sitesRRM = *NewSitesRRM(*baseController)
	client.sitesRSSIZones = *NewSitesRSSIZones(*baseController)
	client.sitesRfdiags = *NewSitesRfdiags(*baseController)
	client.sitesRogues = *NewSitesRogues(*baseController)
	client.sitesSLEs = *NewSitesSLEs(*baseController)
	client.sitesServicePolicies = *NewSitesServicePolicies(*baseController)
	client.sitesServices = *NewSitesServices(*baseController)
	client.sitesSetting = *NewSitesSetting(*baseController)
	client.sitesSiteTemplates = *NewSitesSiteTemplates(*baseController)
	client.sitesSkyatp = *NewSitesSkyatp(*baseController)
	client.sitesSyntheticTests = *NewSitesSyntheticTests(*baseController)
	client.sitesUISettings = *NewSitesUISettings(*baseController)
	client.sitesVPNs = *NewSitesVPNs(*baseController)
	client.sitesWANUsages = *NewSitesWANUsages(*baseController)
	client.sitesWebhooks = *NewSitesWebhooks(*baseController)
	client.sitesWlans = *NewSitesWlans(*baseController)
	client.sitesWxRules = *NewSitesWxRules(*baseController)
	client.sitesWxTags = *NewSitesWxTags(*baseController)
	client.sitesWxTunnels = *NewSitesWxTunnels(*baseController)
	client.sitesZones = *NewSitesZones(*baseController)
	client.sitesVBeacons = *NewSitesVBeacons(*baseController)
	client.utilitiesCommon = *NewUtilitiesCommon(*baseController)
	client.utilitiesLAN = *NewUtilitiesLAN(*baseController)
	client.utilitiesLocation = *NewUtilitiesLocation(*baseController)
	client.utilitiesMxEdge = *NewUtilitiesMxEdge(*baseController)
	client.utilitiesWAN = *NewUtilitiesWAN(*baseController)
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

// Orgs returns the orgs instance of the client.
func (c *client) Orgs() *Orgs {
	return &c.orgs
}

// OrgsAPTemplates returns the orgsAPTemplates instance of the client.
func (c *client) OrgsAPTemplates() *OrgsAPTemplates {
	return &c.orgsAPTemplates
}

// OrgsAPITokens returns the orgsAPITokens instance of the client.
func (c *client) OrgsAPITokens() *OrgsAPITokens {
	return &c.orgsAPITokens
}

// OrgsAdmins returns the orgsAdmins instance of the client.
func (c *client) OrgsAdmins() *OrgsAdmins {
	return &c.orgsAdmins
}

// OrgsAlarmTemplates returns the orgsAlarmTemplates instance of the client.
func (c *client) OrgsAlarmTemplates() *OrgsAlarmTemplates {
	return &c.orgsAlarmTemplates
}

// OrgsAlarms returns the orgsAlarms instance of the client.
func (c *client) OrgsAlarms() *OrgsAlarms {
	return &c.orgsAlarms
}

// OrgsAssetFilters returns the orgsAssetFilters instance of the client.
func (c *client) OrgsAssetFilters() *OrgsAssetFilters {
	return &c.orgsAssetFilters
}

// OrgsAssets returns the orgsAssets instance of the client.
func (c *client) OrgsAssets() *OrgsAssets {
	return &c.orgsAssets
}

// OrgsCRL returns the orgsCRL instance of the client.
func (c *client) OrgsCRL() *OrgsCRL {
	return &c.orgsCRL
}

// OrgsCert returns the orgsCert instance of the client.
func (c *client) OrgsCert() *OrgsCert {
	return &c.orgsCert
}

// OrgsClientsMarvis returns the orgsClientsMarvis instance of the client.
func (c *client) OrgsClientsMarvis() *OrgsClientsMarvis {
	return &c.orgsClientsMarvis
}

// OrgsClientsNAC returns the orgsClientsNAC instance of the client.
func (c *client) OrgsClientsNAC() *OrgsClientsNAC {
	return &c.orgsClientsNAC
}

// OrgsClientsSDK returns the orgsClientsSDK instance of the client.
func (c *client) OrgsClientsSDK() *OrgsClientsSDK {
	return &c.orgsClientsSDK
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

// OrgsCradlepoint returns the orgsCradlepoint instance of the client.
func (c *client) OrgsCradlepoint() *OrgsCradlepoint {
	return &c.orgsCradlepoint
}

// OrgsDeviceProfiles returns the orgsDeviceProfiles instance of the client.
func (c *client) OrgsDeviceProfiles() *OrgsDeviceProfiles {
	return &c.orgsDeviceProfiles
}

// OrgsDevices returns the orgsDevices instance of the client.
func (c *client) OrgsDevices() *OrgsDevices {
	return &c.orgsDevices
}

// OrgsDevicesOthers returns the orgsDevicesOthers instance of the client.
func (c *client) OrgsDevicesOthers() *OrgsDevicesOthers {
	return &c.orgsDevicesOthers
}

// OrgsDevicesSSR returns the orgsDevicesSSR instance of the client.
func (c *client) OrgsDevicesSSR() *OrgsDevicesSSR {
	return &c.orgsDevicesSSR
}

// OrgsDevicesWANCluster returns the orgsDevicesWANCluster instance of the client.
func (c *client) OrgsDevicesWANCluster() *OrgsDevicesWANCluster {
	return &c.orgsDevicesWANCluster
}

// OrgsDevicesStats returns the orgsDevicesStats instance of the client.
func (c *client) OrgsDevicesStats() *OrgsDevicesStats {
	return &c.orgsDevicesStats
}

// OrgsDevicesZscaler returns the orgsDevicesZscaler instance of the client.
func (c *client) OrgsDevicesZscaler() *OrgsDevicesZscaler {
	return &c.orgsDevicesZscaler
}

// OrgsEVPNTopologies returns the orgsEVPNTopologies instance of the client.
func (c *client) OrgsEVPNTopologies() *OrgsEVPNTopologies {
	return &c.orgsEVPNTopologies
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

// OrgsMaps returns the orgsMaps instance of the client.
func (c *client) OrgsMaps() *OrgsMaps {
	return &c.orgsMaps
}

// OrgsMarvis returns the orgsMarvis instance of the client.
func (c *client) OrgsMarvis() *OrgsMarvis {
	return &c.orgsMarvis
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

// OrgsNACCRL returns the orgsNACCRL instance of the client.
func (c *client) OrgsNACCRL() *OrgsNACCRL {
	return &c.orgsNACCRL
}

// OrgsNACIDP returns the orgsNACIDP instance of the client.
func (c *client) OrgsNACIDP() *OrgsNACIDP {
	return &c.orgsNACIDP
}

// OrgsNACPortals returns the orgsNACPortals instance of the client.
func (c *client) OrgsNACPortals() *OrgsNACPortals {
	return &c.orgsNACPortals
}

// OrgsNACRules returns the orgsNACRules instance of the client.
func (c *client) OrgsNACRules() *OrgsNACRules {
	return &c.orgsNACRules
}

// OrgsNACTags returns the orgsNACTags instance of the client.
func (c *client) OrgsNACTags() *OrgsNACTags {
	return &c.orgsNACTags
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

// OrgsPskPortals returns the orgsPskPortals instance of the client.
func (c *client) OrgsPskPortals() *OrgsPskPortals {
	return &c.orgsPskPortals
}

// OrgsPsks returns the orgsPsks instance of the client.
func (c *client) OrgsPsks() *OrgsPsks {
	return &c.orgsPsks
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

// OrgsSLEs returns the orgsSLEs instance of the client.
func (c *client) OrgsSLEs() *OrgsSLEs {
	return &c.orgsSLEs
}

// OrgsSSO returns the orgsSSO instance of the client.
func (c *client) OrgsSSO() *OrgsSSO {
	return &c.orgsSSO
}

// OrgsSSORoles returns the orgsSSORoles instance of the client.
func (c *client) OrgsSSORoles() *OrgsSSORoles {
	return &c.orgsSSORoles
}

// OrgsSecPolicies returns the orgsSecPolicies instance of the client.
func (c *client) OrgsSecPolicies() *OrgsSecPolicies {
	return &c.orgsSecPolicies
}

// OrgsServicePolicies returns the orgsServicePolicies instance of the client.
func (c *client) OrgsServicePolicies() *OrgsServicePolicies {
	return &c.orgsServicePolicies
}

// OrgsServices returns the orgsServices instance of the client.
func (c *client) OrgsServices() *OrgsServices {
	return &c.orgsServices
}

// OrgsSetting returns the orgsSetting instance of the client.
func (c *client) OrgsSetting() *OrgsSetting {
	return &c.orgsSetting
}

// OrgsSiteTemplates returns the orgsSiteTemplates instance of the client.
func (c *client) OrgsSiteTemplates() *OrgsSiteTemplates {
	return &c.orgsSiteTemplates
}

// OrgsSitegroups returns the orgsSitegroups instance of the client.
func (c *client) OrgsSitegroups() *OrgsSitegroups {
	return &c.orgsSitegroups
}

// OrgsSites returns the orgsSites instance of the client.
func (c *client) OrgsSites() *OrgsSites {
	return &c.orgsSites
}

// OrgsSubscriptions returns the orgsSubscriptions instance of the client.
func (c *client) OrgsSubscriptions() *OrgsSubscriptions {
	return &c.orgsSubscriptions
}

// OrgsTickets returns the orgsTickets instance of the client.
func (c *client) OrgsTickets() *OrgsTickets {
	return &c.orgsTickets
}

// OrgsTunnels returns the orgsTunnels instance of the client.
func (c *client) OrgsTunnels() *OrgsTunnels {
	return &c.orgsTunnels
}

// OrgsUserMACs returns the orgsUserMACs instance of the client.
func (c *client) OrgsUserMACs() *OrgsUserMACs {
	return &c.orgsUserMACs
}

// OrgsVPNs returns the orgsVPNs instance of the client.
func (c *client) OrgsVPNs() *OrgsVPNs {
	return &c.orgsVPNs
}

// OrgsVars returns the orgsVars instance of the client.
func (c *client) OrgsVars() *OrgsVars {
	return &c.orgsVars
}

// OrgsWLANTemplates returns the orgsWLANTemplates instance of the client.
func (c *client) OrgsWLANTemplates() *OrgsWLANTemplates {
	return &c.orgsWLANTemplates
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

// UtilitiesPCAPs returns the utilitiesPCAPs instance of the client.
func (c *client) UtilitiesPCAPs() *UtilitiesPCAPs {
	return &c.utilitiesPCAPs
}

// UtilitiesUpgrade returns the utilitiesUpgrade instance of the client.
func (c *client) UtilitiesUpgrade() *UtilitiesUpgrade {
	return &c.utilitiesUpgrade
}

// UtilitiesWiFi returns the utilitiesWiFi instance of the client.
func (c *client) UtilitiesWiFi() *UtilitiesWiFi {
	return &c.utilitiesWiFi
}

// Installer returns the installer instance of the client.
func (c *client) Installer() *Installer {
	return &c.installer
}

// Admins returns the admins instance of the client.
func (c *client) Admins() *Admins {
	return &c.admins
}

// AdminsLogin returns the adminsLogin instance of the client.
func (c *client) AdminsLogin() *AdminsLogin {
	return &c.adminsLogin
}

// AdminsLoginOAuth2 returns the adminsLoginOAuth2 instance of the client.
func (c *client) AdminsLoginOAuth2() *AdminsLoginOAuth2 {
	return &c.adminsLoginOAuth2
}

// AdminsLogout returns the adminsLogout instance of the client.
func (c *client) AdminsLogout() *AdminsLogout {
	return &c.adminsLogout
}

// AdminsLookup returns the adminsLookup instance of the client.
func (c *client) AdminsLookup() *AdminsLookup {
	return &c.adminsLookup
}

// AdminsRecoverPassword returns the adminsRecoverPassword instance of the client.
func (c *client) AdminsRecoverPassword() *AdminsRecoverPassword {
	return &c.adminsRecoverPassword
}

// ConstantsDefinitions returns the constantsDefinitions instance of the client.
func (c *client) ConstantsDefinitions() *ConstantsDefinitions {
	return &c.constantsDefinitions
}

// ConstantsEvents returns the constantsEvents instance of the client.
func (c *client) ConstantsEvents() *ConstantsEvents {
	return &c.constantsEvents
}

// ConstantsMisc returns the constantsMisc instance of the client.
func (c *client) ConstantsMisc() *ConstantsMisc {
	return &c.constantsMisc
}

// ConstantsModels returns the constantsModels instance of the client.
func (c *client) ConstantsModels() *ConstantsModels {
	return &c.constantsModels
}

// SelfAPIToken returns the selfAPIToken instance of the client.
func (c *client) SelfAPIToken() *SelfAPIToken {
	return &c.selfAPIToken
}

// SelfAccount returns the selfAccount instance of the client.
func (c *client) SelfAccount() *SelfAccount {
	return &c.selfAccount
}

// SelfAlarms returns the selfAlarms instance of the client.
func (c *client) SelfAlarms() *SelfAlarms {
	return &c.selfAlarms
}

// SelfAuditLogs returns the selfAuditLogs instance of the client.
func (c *client) SelfAuditLogs() *SelfAuditLogs {
	return &c.selfAuditLogs
}

// SelfMFA returns the selfMFA instance of the client.
func (c *client) SelfMFA() *SelfMFA {
	return &c.selfMFA
}

// SelfOAuth2 returns the selfOAuth2 instance of the client.
func (c *client) SelfOAuth2() *SelfOAuth2 {
	return &c.selfOAuth2
}

// MSPs returns the mSPs instance of the client.
func (c *client) MSPs() *MSPs {
	return &c.mSPs
}

// MSPsAdmins returns the mSPsAdmins instance of the client.
func (c *client) MSPsAdmins() *MSPsAdmins {
	return &c.mSPsAdmins
}

// MSPsInventory returns the mSPsInventory instance of the client.
func (c *client) MSPsInventory() *MSPsInventory {
	return &c.mSPsInventory
}

// MSPsLicenses returns the mSPsLicenses instance of the client.
func (c *client) MSPsLicenses() *MSPsLicenses {
	return &c.mSPsLicenses
}

// MSPsLogo returns the mSPsLogo instance of the client.
func (c *client) MSPsLogo() *MSPsLogo {
	return &c.mSPsLogo
}

// MSPsLogs returns the mSPsLogs instance of the client.
func (c *client) MSPsLogs() *MSPsLogs {
	return &c.mSPsLogs
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

// MSPsSSO returns the mSPsSSO instance of the client.
func (c *client) MSPsSSO() *MSPsSSO {
	return &c.mSPsSSO
}

// MSPsSSORoles returns the mSPsSSORoles instance of the client.
func (c *client) MSPsSSORoles() *MSPsSSORoles {
	return &c.mSPsSSORoles
}

// MSPsTickets returns the mSPsTickets instance of the client.
func (c *client) MSPsTickets() *MSPsTickets {
	return &c.mSPsTickets
}

// Sites returns the sites instance of the client.
func (c *client) Sites() *Sites {
	return &c.sites
}

// SitesAPTemplates returns the sitesAPTemplates instance of the client.
func (c *client) SitesAPTemplates() *SitesAPTemplates {
	return &c.sitesAPTemplates
}

// SitesAlarms returns the sitesAlarms instance of the client.
func (c *client) SitesAlarms() *SitesAlarms {
	return &c.sitesAlarms
}

// SitesAnomaly returns the sitesAnomaly instance of the client.
func (c *client) SitesAnomaly() *SitesAnomaly {
	return &c.sitesAnomaly
}

// SitesApplications returns the sitesApplications instance of the client.
func (c *client) SitesApplications() *SitesApplications {
	return &c.sitesApplications
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

// SitesCalls returns the sitesCalls instance of the client.
func (c *client) SitesCalls() *SitesCalls {
	return &c.sitesCalls
}

// SitesClientsNAC returns the sitesClientsNAC instance of the client.
func (c *client) SitesClientsNAC() *SitesClientsNAC {
	return &c.sitesClientsNAC
}

// SitesClientsSDK returns the sitesClientsSDK instance of the client.
func (c *client) SitesClientsSDK() *SitesClientsSDK {
	return &c.sitesClientsSDK
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

// SitesDeviceProfiles returns the sitesDeviceProfiles instance of the client.
func (c *client) SitesDeviceProfiles() *SitesDeviceProfiles {
	return &c.sitesDeviceProfiles
}

// SitesDevices returns the sitesDevices instance of the client.
func (c *client) SitesDevices() *SitesDevices {
	return &c.sitesDevices
}

// SitesDevicesDiscoveredSwitches returns the sitesDevicesDiscoveredSwitches instance of the client.
func (c *client) SitesDevicesDiscoveredSwitches() *SitesDevicesDiscoveredSwitches {
	return &c.sitesDevicesDiscoveredSwitches
}

// SitesDevicesOthers returns the sitesDevicesOthers instance of the client.
func (c *client) SitesDevicesOthers() *SitesDevicesOthers {
	return &c.sitesDevicesOthers
}

// SitesDevicesWANCluster returns the sitesDevicesWANCluster instance of the client.
func (c *client) SitesDevicesWANCluster() *SitesDevicesWANCluster {
	return &c.sitesDevicesWANCluster
}

// SitesDevicesWired returns the sitesDevicesWired instance of the client.
func (c *client) SitesDevicesWired() *SitesDevicesWired {
	return &c.sitesDevicesWired
}

// SitesDevicesWiredVirtualChassis returns the sitesDevicesWiredVirtualChassis instance of the client.
func (c *client) SitesDevicesWiredVirtualChassis() *SitesDevicesWiredVirtualChassis {
	return &c.sitesDevicesWiredVirtualChassis
}

// SitesDevicesWireless returns the sitesDevicesWireless instance of the client.
func (c *client) SitesDevicesWireless() *SitesDevicesWireless {
	return &c.sitesDevicesWireless
}

// SitesDevicesStats returns the sitesDevicesStats instance of the client.
func (c *client) SitesDevicesStats() *SitesDevicesStats {
	return &c.sitesDevicesStats
}

// SitesEVPNTopologies returns the sitesEVPNTopologies instance of the client.
func (c *client) SitesEVPNTopologies() *SitesEVPNTopologies {
	return &c.sitesEVPNTopologies
}

// SitesEvents returns the sitesEvents instance of the client.
func (c *client) SitesEvents() *SitesEvents {
	return &c.sitesEvents
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

// SitesMapsAutoOrientation returns the sitesMapsAutoOrientation instance of the client.
func (c *client) SitesMapsAutoOrientation() *SitesMapsAutoOrientation {
	return &c.sitesMapsAutoOrientation
}

// SitesMapsAutoPlacement returns the sitesMapsAutoPlacement instance of the client.
func (c *client) SitesMapsAutoPlacement() *SitesMapsAutoPlacement {
	return &c.sitesMapsAutoPlacement
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

// SitesRRM returns the sitesRRM instance of the client.
func (c *client) SitesRRM() *SitesRRM {
	return &c.sitesRRM
}

// SitesRSSIZones returns the sitesRSSIZones instance of the client.
func (c *client) SitesRSSIZones() *SitesRSSIZones {
	return &c.sitesRSSIZones
}

// SitesRfdiags returns the sitesRfdiags instance of the client.
func (c *client) SitesRfdiags() *SitesRfdiags {
	return &c.sitesRfdiags
}

// SitesRogues returns the sitesRogues instance of the client.
func (c *client) SitesRogues() *SitesRogues {
	return &c.sitesRogues
}

// SitesSLEs returns the sitesSLEs instance of the client.
func (c *client) SitesSLEs() *SitesSLEs {
	return &c.sitesSLEs
}

// SitesServicePolicies returns the sitesServicePolicies instance of the client.
func (c *client) SitesServicePolicies() *SitesServicePolicies {
	return &c.sitesServicePolicies
}

// SitesServices returns the sitesServices instance of the client.
func (c *client) SitesServices() *SitesServices {
	return &c.sitesServices
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

// SitesSyntheticTests returns the sitesSyntheticTests instance of the client.
func (c *client) SitesSyntheticTests() *SitesSyntheticTests {
	return &c.sitesSyntheticTests
}

// SitesUISettings returns the sitesUISettings instance of the client.
func (c *client) SitesUISettings() *SitesUISettings {
	return &c.sitesUISettings
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

// SitesVBeacons returns the sitesVBeacons instance of the client.
func (c *client) SitesVBeacons() *SitesVBeacons {
	return &c.sitesVBeacons
}

// UtilitiesCommon returns the utilitiesCommon instance of the client.
func (c *client) UtilitiesCommon() *UtilitiesCommon {
	return &c.utilitiesCommon
}

// UtilitiesLAN returns the utilitiesLAN instance of the client.
func (c *client) UtilitiesLAN() *UtilitiesLAN {
	return &c.utilitiesLAN
}

// UtilitiesLocation returns the utilitiesLocation instance of the client.
func (c *client) UtilitiesLocation() *UtilitiesLocation {
	return &c.utilitiesLocation
}

// UtilitiesMxEdge returns the utilitiesMxEdge instance of the client.
func (c *client) UtilitiesMxEdge() *UtilitiesMxEdge {
	return &c.utilitiesMxEdge
}

// UtilitiesWAN returns the utilitiesWAN instance of the client.
func (c *client) UtilitiesWAN() *UtilitiesWAN {
	return &c.utilitiesWAN
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
