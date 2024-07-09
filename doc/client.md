
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.MIST_GLOBAL_01`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `loggerConfiguration` | [`LoggerConfiguration`](logger-configuration.md) | Represents the logger configurations for API calls |
| `apiTokenCredentials` | [`ApiTokenCredentials`](auth/custom-header-signature.md) | The Credentials Setter for Custom Header Signature |
| `basicAuthCredentials` | [`BasicAuthCredentials`](auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

The API client can be initialized as follows:

```go
client := mistapigo.NewClient(
    mistapigo.CreateConfiguration(
        mistapigo.WithHttpConfiguration(
            mistapigo.CreateHttpConfiguration(
                mistapigo.WithTimeout(0),
            ),
        ),
        mistapigo.WithEnvironment(mistapigo.MIST_GLOBAL_01),
        mistapigo.WithApiTokenCredentials(
            mistapigo.NewApiTokenCredentials("Authorization"),
        ),
        mistapigo.WithBasicAuthCredentials(
            mistapigo.NewBasicAuthCredentials(
                "Username",
                "Password",
            ),
        ),
        mistapigo.WithLoggerConfiguration(
            mistapigo.WithLevel("info"),
            mistapigo.WithRequestConfiguration(
                mistapigo.WithRequestBody(true),
            ),
            mistapigo.WithResponseConfiguration(
                mistapigo.WithResponseHeaders(true),
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
| orgs | Gets Orgs |
| orgsAPTemplates | Gets OrgsAPTemplates |
| orgsAPITokens | Gets OrgsAPITokens |
| orgsAdmins | Gets OrgsAdmins |
| orgsAlarmTemplates | Gets OrgsAlarmTemplates |
| orgsAlarms | Gets OrgsAlarms |
| orgsAssetFilters | Gets OrgsAssetFilters |
| orgsAssets | Gets OrgsAssets |
| orgsCRL | Gets OrgsCRL |
| orgsCert | Gets OrgsCert |
| orgsClientsMarvis | Gets OrgsClientsMarvis |
| orgsClientsNAC | Gets OrgsClientsNAC |
| orgsClientsSDK | Gets OrgsClientsSDK |
| orgsClientsWan | Gets OrgsClientsWan |
| orgsClientsWired | Gets OrgsClientsWired |
| orgsClientsWireless | Gets OrgsClientsWireless |
| orgsCradlepoint | Gets OrgsCradlepoint |
| orgsDeviceProfiles | Gets OrgsDeviceProfiles |
| orgsDevices | Gets OrgsDevices |
| orgsDevicesOthers | Gets OrgsDevicesOthers |
| orgsDevicesSSR | Gets OrgsDevicesSSR |
| orgsDevicesWANCluster | Gets OrgsDevicesWANCluster |
| orgsDevicesStats | Gets OrgsDevicesStats |
| orgsDevicesZscaler | Gets OrgsDevicesZscaler |
| orgsEVPNTopologies | Gets OrgsEVPNTopologies |
| orgsGatewayTemplates | Gets OrgsGatewayTemplates |
| orgsGuests | Gets OrgsGuests |
| orgsIDPProfiles | Gets OrgsIDPProfiles |
| orgsInventory | Gets OrgsInventory |
| orgsJSE | Gets OrgsJSE |
| orgsJSI | Gets OrgsJSI |
| orgsLicenses | Gets OrgsLicenses |
| orgsLinkedApplications | Gets OrgsLinkedApplications |
| orgsLogs | Gets OrgsLogs |
| orgsMaps | Gets OrgsMaps |
| orgsMarvis | Gets OrgsMarvis |
| orgsMxClusters | Gets OrgsMxClusters |
| orgsMxEdges | Gets OrgsMxEdges |
| orgsMxTunnels | Gets OrgsMxTunnels |
| orgsNACCRL | Gets OrgsNACCRL |
| orgsNACIDP | Gets OrgsNACIDP |
| orgsNACPortals | Gets OrgsNACPortals |
| orgsNACRules | Gets OrgsNACRules |
| orgsNACTags | Gets OrgsNACTags |
| orgsNetworkTemplates | Gets OrgsNetworkTemplates |
| orgsNetworks | Gets OrgsNetworks |
| orgsPremiumAnalytics | Gets OrgsPremiumAnalytics |
| orgsPskPortals | Gets OrgsPskPortals |
| orgsPsks | Gets OrgsPsks |
| orgsRFTemplates | Gets OrgsRFTemplates |
| orgsSDKInvites | Gets OrgsSDKInvites |
| orgsSDKTemplates | Gets OrgsSDKTemplates |
| orgsSLEs | Gets OrgsSLEs |
| orgsSSO | Gets OrgsSSO |
| orgsSSORoles | Gets OrgsSSORoles |
| orgsSecPolicies | Gets OrgsSecPolicies |
| orgsServicePolicies | Gets OrgsServicePolicies |
| orgsServices | Gets OrgsServices |
| orgsSetting | Gets OrgsSetting |
| orgsSiteTemplates | Gets OrgsSiteTemplates |
| orgsSitegroups | Gets OrgsSitegroups |
| orgsSites | Gets OrgsSites |
| orgsSubscriptions | Gets OrgsSubscriptions |
| orgsTickets | Gets OrgsTickets |
| orgsTunnels | Gets OrgsTunnels |
| orgsUserMACs | Gets OrgsUserMACs |
| orgsVPNs | Gets OrgsVPNs |
| orgsVars | Gets OrgsVars |
| orgsWLANTemplates | Gets OrgsWLANTemplates |
| orgsWebhooks | Gets OrgsWebhooks |
| orgsWlans | Gets OrgsWlans |
| orgsWxRules | Gets OrgsWxRules |
| orgsWxTags | Gets OrgsWxTags |
| orgsWxTunnels | Gets OrgsWxTunnels |
| utilitiesPCAPs | Gets UtilitiesPCAPs |
| utilitiesUpgrade | Gets UtilitiesUpgrade |
| utilitiesWiFi | Gets UtilitiesWiFi |
| installer | Gets Installer |
| admins | Gets Admins |
| adminsLogin | Gets AdminsLogin |
| adminsLoginOAuth2 | Gets AdminsLoginOAuth2 |
| adminsLogout | Gets AdminsLogout |
| adminsLookup | Gets AdminsLookup |
| adminsRecoverPassword | Gets AdminsRecoverPassword |
| constantsDefinitions | Gets ConstantsDefinitions |
| constantsEvents | Gets ConstantsEvents |
| constantsMisc | Gets ConstantsMisc |
| constantsModels | Gets ConstantsModels |
| samplesWebhook | Gets SamplesWebhook |
| selfAPIToken | Gets SelfAPIToken |
| selfAccount | Gets SelfAccount |
| selfAlarms | Gets SelfAlarms |
| selfAuditLogs | Gets SelfAuditLogs |
| selfMFA | Gets SelfMFA |
| selfOAuth2 | Gets SelfOAuth2 |
| mSPs | Gets MSPs |
| mSPsAdmins | Gets MSPsAdmins |
| mSPsInventory | Gets MSPsInventory |
| mSPsLicenses | Gets MSPsLicenses |
| mSPsLogo | Gets MSPsLogo |
| mSPsLogs | Gets MSPsLogs |
| mSPsMarvis | Gets MSPsMarvis |
| mSPsOrgGroups | Gets MSPsOrgGroups |
| mSPsOrgs | Gets MSPsOrgs |
| mSPsSLEs | Gets MSPsSLEs |
| mSPsSSO | Gets MSPsSSO |
| mSPsSSORoles | Gets MSPsSSORoles |
| mSPsTickets | Gets MSPsTickets |
| sites | Gets Sites |
| sitesAPTemplates | Gets SitesAPTemplates |
| sitesAlarms | Gets SitesAlarms |
| sitesAnomaly | Gets SitesAnomaly |
| sitesApplications | Gets SitesApplications |
| sitesAssetFilters | Gets SitesAssetFilters |
| sitesAssets | Gets SitesAssets |
| sitesBeacons | Gets SitesBeacons |
| sitesCalls | Gets SitesCalls |
| sitesClientsNAC | Gets SitesClientsNAC |
| sitesClientsSDK | Gets SitesClientsSDK |
| sitesClientsWan | Gets SitesClientsWan |
| sitesClientsWired | Gets SitesClientsWired |
| sitesClientsWireless | Gets SitesClientsWireless |
| sitesDeviceProfiles | Gets SitesDeviceProfiles |
| sitesDevices | Gets SitesDevices |
| sitesDevicesDiscoveredSwitches | Gets SitesDevicesDiscoveredSwitches |
| sitesDevicesOthers | Gets SitesDevicesOthers |
| sitesDevicesWANCluster | Gets SitesDevicesWANCluster |
| sitesDevicesWired | Gets SitesDevicesWired |
| sitesDevicesWiredVirtualChassis | Gets SitesDevicesWiredVirtualChassis |
| sitesDevicesWireless | Gets SitesDevicesWireless |
| sitesDevicesStats | Gets SitesDevicesStats |
| sitesEVPNTopologies | Gets SitesEVPNTopologies |
| sitesEvents | Gets SitesEvents |
| sitesGatewayTemplates | Gets SitesGatewayTemplates |
| sitesGuests | Gets SitesGuests |
| sitesInsights | Gets SitesInsights |
| sitesJSE | Gets SitesJSE |
| sitesLicenses | Gets SitesLicenses |
| sitesLocation | Gets SitesLocation |
| sitesMaps | Gets SitesMaps |
| sitesMapsAutoOrientation | Gets SitesMapsAutoOrientation |
| sitesMapsAutoPlacement | Gets SitesMapsAutoPlacement |
| sitesMxEdges | Gets SitesMxEdges |
| sitesNetworkTemplates | Gets SitesNetworkTemplates |
| sitesNetworks | Gets SitesNetworks |
| sitesPsks | Gets SitesPsks |
| sitesRFTemplates | Gets SitesRFTemplates |
| sitesRRM | Gets SitesRRM |
| sitesRSSIZones | Gets SitesRSSIZones |
| sitesRfdiags | Gets SitesRfdiags |
| sitesRogues | Gets SitesRogues |
| sitesSLEs | Gets SitesSLEs |
| sitesServicePolicies | Gets SitesServicePolicies |
| sitesServices | Gets SitesServices |
| sitesSetting | Gets SitesSetting |
| sitesSiteTemplates | Gets SitesSiteTemplates |
| sitesSkyatp | Gets SitesSkyatp |
| sitesSyntheticTests | Gets SitesSyntheticTests |
| sitesUISettings | Gets SitesUISettings |
| sitesVPNs | Gets SitesVPNs |
| sitesWANUsages | Gets SitesWANUsages |
| sitesWebhooks | Gets SitesWebhooks |
| sitesWlans | Gets SitesWlans |
| sitesWxRules | Gets SitesWxRules |
| sitesWxTags | Gets SitesWxTags |
| sitesWxTunnels | Gets SitesWxTunnels |
| sitesZones | Gets SitesZones |
| sitesVBeacons | Gets SitesVBeacons |
| utilitiesCommon | Gets UtilitiesCommon |
| utilitiesLAN | Gets UtilitiesLAN |
| utilitiesLocation | Gets UtilitiesLocation |
| utilitiesMxEdge | Gets UtilitiesMxEdge |
| utilitiesWAN | Gets UtilitiesWAN |

