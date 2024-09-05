
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.MIST_GLOBAL_01`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `loggerConfiguration` | [`LoggerConfiguration`](logger-configuration.md) | Represents the logger configurations for API calls |
| `apiTokenCredentials` | [`ApiTokenCredentials`](auth/custom-header-signature.md) | The Credentials Setter for Custom Header Signature |
| `basicAuthCredentials` | [`BasicAuthCredentials`](auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |
| `csrfTokenCredentials` | [`CsrfTokenCredentials`](auth/custom-header-signature-1.md) | The Credentials Setter for Custom Header Signature |

The API client can be initialized as follows:

```go
client := mistapi.NewClient(
    mistapi.CreateConfiguration(
        mistapi.WithHttpConfiguration(
            mistapi.CreateHttpConfiguration(
                mistapi.WithTimeout(0),
            ),
        ),
        mistapi.WithEnvironment(mistapi.MIST_GLOBAL_01),
        mistapi.WithApiTokenCredentials(
            mistapi.NewApiTokenCredentials("Authorization"),
        ),
        mistapi.WithBasicAuthCredentials(
            mistapi.NewBasicAuthCredentials(
                "Username",
                "Password",
            ),
        ),
        mistapi.WithCsrfTokenCredentials(
            mistapi.NewCsrfTokenCredentials("X-CSRFToken"),
        ),
        mistapi.WithLoggerConfiguration(
            mistapi.WithLevel("info"),
            mistapi.WithRequestConfiguration(
                mistapi.WithRequestBody(true),
            ),
            mistapi.WithResponseConfiguration(
                mistapi.WithResponseHeaders(true),
            ),
        ),
    ),
)
```

## Mist API Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| Admins() | Gets Admins |
| AdminsLogin() | Gets AdminsLogin |
| AdminsLogout() | Gets AdminsLogout |
| AdminsRecoverPassword() | Gets AdminsRecoverPassword |
| AdminsLookup() | Gets AdminsLookup |
| AdminsLoginOAuth2() | Gets AdminsLoginOAuth2 |
| Installer() | Gets Installer |
| MSPs() | Gets MSPs |
| MSPsAdmins() | Gets MSPsAdmins |
| MSPsInventory() | Gets MSPsInventory |
| MSPsLogo() | Gets MSPsLogo |
| MSPsLogs() | Gets MSPsLogs |
| MSPsLicenses() | Gets MSPsLicenses |
| MSPsMarvis() | Gets MSPsMarvis |
| MSPsOrgGroups() | Gets MSPsOrgGroups |
| MSPsOrgs() | Gets MSPsOrgs |
| MSPsSLEs() | Gets MSPsSLEs |
| MSPsSSORoles() | Gets MSPsSSORoles |
| MSPsSSO() | Gets MSPsSSO |
| MSPsTickets() | Gets MSPsTickets |
| Orgs() | Gets Orgs |
| OrgsAdmins() | Gets OrgsAdmins |
| OrgsAlarms() | Gets OrgsAlarms |
| OrgsAlarmTemplates() | Gets OrgsAlarmTemplates |
| OrgsAPTemplates() | Gets OrgsAPTemplates |
| OrgsAPITokens() | Gets OrgsAPITokens |
| OrgsAssets() | Gets OrgsAssets |
| OrgsAssetFilters() | Gets OrgsAssetFilters |
| OrgsCert() | Gets OrgsCert |
| OrgsAntivirusProfiles() | Gets OrgsAntivirusProfiles |
| OrgsClientsMarvis() | Gets OrgsClientsMarvis |
| OrgsClientsNAC() | Gets OrgsClientsNAC |
| OrgsClientsWan() | Gets OrgsClientsWan |
| OrgsClientsWired() | Gets OrgsClientsWired |
| OrgsClientsWireless() | Gets OrgsClientsWireless |
| OrgsClientsSDK() | Gets OrgsClientsSDK |
| OrgsCradlepoint() | Gets OrgsCradlepoint |
| OrgsCRL() | Gets OrgsCRL |
| OrgsDeviceProfiles() | Gets OrgsDeviceProfiles |
| OrgsDevices() | Gets OrgsDevices |
| OrgsDevicesSSR() | Gets OrgsDevicesSSR |
| OrgsEVPNTopologies() | Gets OrgsEVPNTopologies |
| OrgsDevicesOthers() | Gets OrgsDevicesOthers |
| OrgsSettingZscaler() | Gets OrgsSettingZscaler |
| OrgsSCEP() | Gets OrgsSCEP |
| OrgsGatewayTemplates() | Gets OrgsGatewayTemplates |
| OrgsGuests() | Gets OrgsGuests |
| OrgsIDPProfiles() | Gets OrgsIDPProfiles |
| OrgsInventory() | Gets OrgsInventory |
| OrgsJSE() | Gets OrgsJSE |
| OrgsJSI() | Gets OrgsJSI |
| OrgsLicenses() | Gets OrgsLicenses |
| OrgsLinkedApplications() | Gets OrgsLinkedApplications |
| OrgsLogs() | Gets OrgsLogs |
| OrgsMarvis() | Gets OrgsMarvis |
| OrgsMaps() | Gets OrgsMaps |
| OrgsMxClusters() | Gets OrgsMxClusters |
| OrgsMxEdges() | Gets OrgsMxEdges |
| OrgsMxTunnels() | Gets OrgsMxTunnels |
| OrgsNACIDP() | Gets OrgsNACIDP |
| OrgsNACTags() | Gets OrgsNACTags |
| OrgsNACPortals() | Gets OrgsNACPortals |
| OrgsNACCRL() | Gets OrgsNACCRL |
| OrgsStats() | Gets OrgsStats |
| OrgsStatsAssets() | Gets OrgsStatsAssets |
| OrgsStatsBGPPeers() | Gets OrgsStatsBGPPeers |
| OrgsStatsDevices() | Gets OrgsStatsDevices |
| OrgsStatsMxEdges() | Gets OrgsStatsMxEdges |
| OrgsStatsOtherDevices() | Gets OrgsStatsOtherDevices |
| OrgsStatsPorts() | Gets OrgsStatsPorts |
| OrgsStatsTunnels() | Gets OrgsStatsTunnels |
| OrgsStatsVPNPeers() | Gets OrgsStatsVPNPeers |
| OrgsNACRules() | Gets OrgsNACRules |
| OrgsNetworkTemplates() | Gets OrgsNetworkTemplates |
| OrgsNetworks() | Gets OrgsNetworks |
| OrgsPremiumAnalytics() | Gets OrgsPremiumAnalytics |
| OrgsPsks() | Gets OrgsPsks |
| OrgsPskPortals() | Gets OrgsPskPortals |
| OrgsRFTemplates() | Gets OrgsRFTemplates |
| OrgsSDKInvites() | Gets OrgsSDKInvites |
| OrgsSDKTemplates() | Gets OrgsSDKTemplates |
| OrgsSecPolicies() | Gets OrgsSecPolicies |
| OrgsServices() | Gets OrgsServices |
| OrgsServicePolicies() | Gets OrgsServicePolicies |
| OrgsSetting() | Gets OrgsSetting |
| OrgsSitegroups() | Gets OrgsSitegroups |
| OrgsSites() | Gets OrgsSites |
| OrgsSiteTemplates() | Gets OrgsSiteTemplates |
| OrgsSLEs() | Gets OrgsSLEs |
| OrgsSSORoles() | Gets OrgsSSORoles |
| OrgsSSO() | Gets OrgsSSO |
| OrgsSubscriptions() | Gets OrgsSubscriptions |
| OrgsWLANTemplates() | Gets OrgsWLANTemplates |
| OrgsTickets() | Gets OrgsTickets |
| OrgsUserMACs() | Gets OrgsUserMACs |
| OrgsVars() | Gets OrgsVars |
| OrgsVPNs() | Gets OrgsVPNs |
| OrgsWebhooks() | Gets OrgsWebhooks |
| OrgsWlans() | Gets OrgsWlans |
| OrgsWxRules() | Gets OrgsWxRules |
| OrgsWxTags() | Gets OrgsWxTags |
| OrgsWxTunnels() | Gets OrgsWxTunnels |
| Sites() | Gets Sites |
| SitesAlarms() | Gets SitesAlarms |
| SitesAPTemplates() | Gets SitesAPTemplates |
| SitesApplications() | Gets SitesApplications |
| SitesAnomaly() | Gets SitesAnomaly |
| SitesAssetFilters() | Gets SitesAssetFilters |
| SitesAssets() | Gets SitesAssets |
| SitesBeacons() | Gets SitesBeacons |
| SitesClientsNAC() | Gets SitesClientsNAC |
| SitesClientsWan() | Gets SitesClientsWan |
| SitesClientsWired() | Gets SitesClientsWired |
| SitesClientsWireless() | Gets SitesClientsWireless |
| SitesDevices() | Gets SitesDevices |
| SitesDevicesWireless() | Gets SitesDevicesWireless |
| SitesDevicesOthers() | Gets SitesDevicesOthers |
| SitesDevicesWired() | Gets SitesDevicesWired |
| SitesDevicesWiredVirtualChassis() | Gets SitesDevicesWiredVirtualChassis |
| SitesDevicesWANCluster() | Gets SitesDevicesWANCluster |
| SitesDeviceProfiles() | Gets SitesDeviceProfiles |
| SitesEvents() | Gets SitesEvents |
| SitesEVPNTopologies() | Gets SitesEVPNTopologies |
| SitesGatewayTemplates() | Gets SitesGatewayTemplates |
| SitesGuests() | Gets SitesGuests |
| SitesInsights() | Gets SitesInsights |
| SitesJSE() | Gets SitesJSE |
| SitesLicenses() | Gets SitesLicenses |
| SitesLocation() | Gets SitesLocation |
| SitesMaps() | Gets SitesMaps |
| SitesMapsAutoPlacement() | Gets SitesMapsAutoPlacement |
| SitesMxEdges() | Gets SitesMxEdges |
| SitesNetworkTemplates() | Gets SitesNetworkTemplates |
| SitesNetworks() | Gets SitesNetworks |
| SitesPsks() | Gets SitesPsks |
| SitesRFTemplates() | Gets SitesRFTemplates |
| SitesRfdiags() | Gets SitesRfdiags |
| SitesRogues() | Gets SitesRogues |
| SitesRRM() | Gets SitesRRM |
| SitesRSSIZones() | Gets SitesRSSIZones |
| SitesServices() | Gets SitesServices |
| SitesServicePolicies() | Gets SitesServicePolicies |
| SitesSetting() | Gets SitesSetting |
| SitesSiteTemplates() | Gets SitesSiteTemplates |
| SitesSkyatp() | Gets SitesSkyatp |
| SitesSLEs() | Gets SitesSLEs |
| SitesSyntheticTests() | Gets SitesSyntheticTests |
| SitesUISettings() | Gets SitesUISettings |
| SitesVBeacons() | Gets SitesVBeacons |
| SitesVPNs() | Gets SitesVPNs |
| SitesWANUsages() | Gets SitesWANUsages |
| SitesWebhooks() | Gets SitesWebhooks |
| SitesWlans() | Gets SitesWlans |
| SitesWxRules() | Gets SitesWxRules |
| SitesWxTags() | Gets SitesWxTags |
| SitesWxTunnels() | Gets SitesWxTunnels |
| SitesZones() | Gets SitesZones |
| SitesStats() | Gets SitesStats |
| SitesStatsApps() | Gets SitesStatsApps |
| SitesStatsAssets() | Gets SitesStatsAssets |
| SitesStatsBeacons() | Gets SitesStatsBeacons |
| SitesStatsBGPPeers() | Gets SitesStatsBGPPeers |
| SitesStatsCalls() | Gets SitesStatsCalls |
| SitesStatsClientsWireless() | Gets SitesStatsClientsWireless |
| SitesStatsClientsSDK() | Gets SitesStatsClientsSDK |
| SitesStatsDevices() | Gets SitesStatsDevices |
| SitesStatsMxEdges() | Gets SitesStatsMxEdges |
| SitesStatsPorts() | Gets SitesStatsPorts |
| SitesStatsWxRules() | Gets SitesStatsWxRules |
| SitesStatsZones() | Gets SitesStatsZones |
| SitesStatsDiscoveredSwitches() | Gets SitesStatsDiscoveredSwitches |
| ConstantsDefinitions() | Gets ConstantsDefinitions |
| ConstantsEvents() | Gets ConstantsEvents |
| ConstantsModels() | Gets ConstantsModels |
| SelfAccount() | Gets SelfAccount |
| SelfAPIToken() | Gets SelfAPIToken |
| SelfOAuth2() | Gets SelfOAuth2 |
| SelfMFA() | Gets SelfMFA |
| SelfAlarms() | Gets SelfAlarms |
| SelfAuditLogs() | Gets SelfAuditLogs |
| SamplesWebhooks() | Gets SamplesWebhooks |
| UtilitiesCommon() | Gets UtilitiesCommon |
| UtilitiesWAN() | Gets UtilitiesWAN |
| UtilitiesLAN() | Gets UtilitiesLAN |
| UtilitiesWiFi() | Gets UtilitiesWiFi |
| UtilitiesPCAPs() | Gets UtilitiesPCAPs |
| UtilitiesLocation() | Gets UtilitiesLocation |
| UtilitiesMxEdge() | Gets UtilitiesMxEdge |
| UtilitiesUpgrade() | Gets UtilitiesUpgrade |

