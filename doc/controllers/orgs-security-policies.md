# Orgs Security Policies

```go
orgsSecurityPolicies := client.OrgsSecurityPolicies()
```

## Class Name

`OrgsSecurityPolicies`

## Methods

* [Create Org Sec Policy](../../doc/controllers/orgs-security-policies.md#create-org-sec-policy)
* [Delete Org Sec Policy](../../doc/controllers/orgs-security-policies.md#delete-org-sec-policy)
* [Get Org Sec Policy](../../doc/controllers/orgs-security-policies.md#get-org-sec-policy)
* [List Org Sec Policies](../../doc/controllers/orgs-security-policies.md#list-org-sec-policies)
* [Update Org Sec Policy](../../doc/controllers/orgs-security-policies.md#update-org-sec-policy)


# Create Org Sec Policy

Create Org Security Policy

```go
CreateOrgSecPolicy(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Secpolicy) (
    models.ApiResponse[models.Secpolicy],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Secpolicy`](../../doc/models/secpolicy.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Secpolicy](../../doc/models/secpolicy.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Secpolicy{
    Name:                 models.ToPointer("string"),
    Wlans:                []models.Wlan{
        models.Wlan{
            AcctImmediateUpdate:                  models.ToPointer(false),
            AcctInterimInterval:                  models.ToPointer(0),
            AcctServers:                          []models.RadiusAcctServer{
                models.RadiusAcctServer{
                    Host:                 "1.2.3.4",
                    KeywrapEnabled:       models.ToPointer(true),
                    KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                    KeywrapKek:           models.ToPointer("1122334455"),
                    KeywrapMack:          models.ToPointer("1122334455"),
                    Port:                 models.ToPointer(models.RadiusAcctPortContainer.FromNumber(1813)),
                    Secret:               "testing123",
                },
            },
            AllowIpv6Ndp:                         models.ToPointer(true),
            AllowMdns:                            models.ToPointer(false),
            AllowSsdp:                            models.ToPointer(false),
            AppLimit:                             models.ToPointer(models.WlanAppLimit{
                Apps:                 map[string]int{
                    "dropbox": 300,
                    "netflix": 60,
                },
                Enabled:              models.ToPointer(false),
                WxtagIds:             map[string]int{
                    "f99862d9-2726-931f-7559-3dfdf5d070d3": 30,
                },
            }),
            AppQos:                               models.ToPointer(models.WlanAppQos{
                Apps:                 map[string]models.WlanAppQosAppsProperties{
                    "skype-business-video": models.WlanAppQosAppsProperties{
                        Dscp:                 models.ToPointer(models.DscpContainer.FromNumber(32)),
                        DstSubnet:            models.ToPointer("10.2.0.0/16"),
                        SrcSubnet:            models.ToPointer("10.2.0.0/16"),
                    },
                },
                Enabled:              models.ToPointer(true),
                Others:               []models.WlanAppQosOthersItem{
                    models.WlanAppQosOthersItem{
                        Dscp:                 models.ToPointer(models.DscpContainer.FromNumber(32)),
                        DstSubnet:            models.ToPointer("10.2.0.0/16"),
                        PortRanges:           models.ToPointer("80,1024-6553"),
                        Protocol:             models.ToPointer("udp"),
                        SrcSubnet:            models.ToPointer("10.2.0.0/16"),
                    },
                },
            }),
            ArpFilter:                            models.ToPointer(false),
            Auth:                                 models.ToPointer(models.WlanAuth{
                AnticlogThreshold:    models.ToPointer(16),
                EapReauth:            models.ToPointer(false),
                EnableMacAuth:        models.ToPointer(false),
                KeyIdx:               models.ToPointer(1),
                Keys:                 []string{
                    "string",
                },
                MultiPskOnly:         models.ToPointer(false),
                Pairwise:             []models.WlanAuthPairwiseItemEnum{
                    models.WlanAuthPairwiseItemEnum_WPA2CCMP,
                },
                PrivateWlan:          models.ToPointer(true),
                Psk:                  models.NewOptional(models.ToPointer("foryoureyesonly")),
                Type:                 models.WlanAuthTypeEnum_PSK,
                WepAsSecondaryAuth:   models.ToPointer(true),
            }),
            AuthServerSelection:                  models.ToPointer(models.WlanAuthServerSelectionEnum_ORDERED),
            AuthServers:                          []models.RadiusAuthServer{
                models.RadiusAuthServer{
                    Host:                        "1.2.3.4",
                    KeywrapEnabled:              models.ToPointer(true),
                    KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
                    KeywrapKek:                  models.ToPointer("1122334455"),
                    KeywrapMack:                 models.ToPointer("1122334455"),
                    Port:                        models.ToPointer(models.RadiusAuthPortContainer.FromNumber(1812)),
                    Secret:                      "testing123",
                },
            },
            AuthServersNasId:                     models.NewOptional(models.ToPointer("5c5b350e0101-nas")),
            AuthServersNasIp:                     models.NewOptional(models.ToPointer("15.3.1.5")),
            AuthServersRetries:                   models.ToPointer(5),
            AuthServersTimeout:                   models.ToPointer(5),
            Band:                                 models.ToPointer("string"),
            BandSteer:                            models.ToPointer(false),
            BandSteerForceBand5:                  models.ToPointer(false),
            Bands:                                []models.Dot11BandEnum{
                models.Dot11BandEnum_ENUM24,
                models.Dot11BandEnum_ENUM5,
            },
            BlockBlacklistClients:                models.ToPointer(false),
            Bonjour:                              models.ToPointer(models.WlanBonjour{
                AdditionalVlanIds:    models.ToPointer(models.AdditionalVlanIdsContainer.FromString("10,20")),
                Enabled:              models.ToPointer(false),
                Services:             map[string]models.WlanBonjourServiceProperties{
                    "airplay": models.WlanBonjourServiceProperties{
                        RadiusGroups:         []string{
                            "teachers",
                        },
                        Scope:                models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum_SAMEAP),
                    },
                },
            }),
            CiscoCwa:                             models.ToPointer(models.WlanCiscoCwa{
                AllowedHostnames:     []string{
                    "snapchat.com",
                },
                AllowedSubnets:       []string{
                    "63.5.3.0/24",
                },
                BlockedSubnets:       []string{
                    "192.168.0.0/16",
                },
                Enabled:              models.ToPointer(false),
            }),
            ClientLimitDown:                      models.ToPointer(models.WlanLimitContainer.FromNumber(1000)),
            ClientLimitDownEnabled:               models.ToPointer(false),
            ClientLimitUp:                        models.ToPointer(models.WlanLimitContainer.FromNumber(512)),
            ClientLimitUpEnabled:                 models.ToPointer(false),
            CoaServers:                           []models.CoaServer{
                models.CoaServer{
                    DisableEventTimestampCheck: models.ToPointer(false),
                    Enabled:                    models.ToPointer(false),
                    Ip:                         "1.2.3.4",
                    Port:                       models.ToPointer(models.RadiusCoaPortContainer.FromNumber(3799)),
                    Secret:                     "testing456",
                },
            },
            Disable11ax:                          models.ToPointer(false),
            DisableHtVhtRates:                    models.ToPointer(false),
            DisableUapsd:                         models.ToPointer(false),
            DisableV1RoamNotify:                  models.ToPointer(false),
            DisableV2RoamNotify:                  models.ToPointer(false),
            DisableWmm:                           models.ToPointer(false),
            DnsServerRewrite:                     models.NewOptional(models.ToPointer(models.WlanDnsServerRewrite{
                Enabled:              models.ToPointer(false),
                RadiusGroups:         map[string]string{
                    "contractor": "172.1.1.1",
                    "guest": "8.8.8.8",
                },
            })),
            Dtim:                                 models.ToPointer(2),
            DynamicPsk:                           models.NewOptional(models.ToPointer(models.WlanDynamicPsk{
                DefaultPsk:           models.ToPointer("foryoureyesonly"),
                DefaultVlanId:        models.ToPointer(models.VlanIdWithVariableContainer.FromNumber(999)),
                Enabled:              models.ToPointer(false),
                Source:               models.ToPointer(models.DynamicPskSourceEnum_CLOUDPSKS),
            })),
            DynamicVlan:                          models.NewOptional(models.ToPointer(models.WlanDynamicVlan{
                DefaultVlanId:        models.ToPointer(models.WlanDynamicVlanDefaultVlanIdDeprecatedContainer.FromNumber(999)),
                Enabled:              models.ToPointer(false),
                LocalVlanIds:         []models.VlanIdWithVariable{
                    models.VlanIdWithVariableContainer.FromNumber(1),
                },
                Type:                 models.ToPointer(models.WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                Vlans:                map[string]string{
                    "131": "default",
                    "322": "fast,video",
                },
            })),
            EnableLocalKeycaching:                models.ToPointer(false),
            EnableWirelessBridging:               models.ToPointer(false),
            Enabled:                              models.ToPointer(true),
            FastDot1xTimers:                      models.ToPointer(false),
            HideSsid:                             models.ToPointer(false),
            HostnameIe:                           models.ToPointer(false),
            Hotspot20:                            models.ToPointer(models.WlanHotspot20{
                DomainName:           []string{
                    "mist.com",
                },
                Enabled:              models.ToPointer(true),
                NaiRealms:            []string{
                    "string",
                },
                Operators:            []models.WlanHotspot20OperatorsItemEnum{
                    models.WlanHotspot20OperatorsItemEnum_GOOGLE,
                    models.WlanHotspot20OperatorsItemEnum_ATT,
                },
                Rcoi:                 []string{
                    "5A03BA0000",
                },
                VenueName:            models.ToPointer("some_name"),
            }),
            Interface:                            models.ToPointer(models.WlanInterfaceEnum_ALL),
            Isolation:                            models.ToPointer(false),
            L2Isolation:                          models.ToPointer(false),
            LegacyOverds:                         models.ToPointer(false),
            LimitBcast:                           models.ToPointer(false),
            LimitProbeResponse:                   models.ToPointer(true),
            MaxIdletime:                          models.ToPointer(1800),
            MistNac:                              models.ToPointer(models.WlanMistNac{
                Enabled:              models.ToPointer(false),
            }),
            NoStaticDns:                          models.ToPointer(false),
            NoStaticIp:                           models.ToPointer(false),
            Portal:                               models.ToPointer(models.WlanPortal{
                AmazonClientId:              models.NewOptional(models.ToPointer("string")),
                AmazonClientSecret:          models.NewOptional(models.ToPointer("string")),
                AmazonEmailDomains:          []string{
                    "string",
                },
                AmazonEnabled:               models.ToPointer(false),
                Auth:                        models.ToPointer(models.WlanPortalAuthEnum_NONE),
                AzureClientId:               models.NewOptional(models.ToPointer("string")),
                AzureClientSecret:           models.NewOptional(models.ToPointer("string")),
                AzureEnabled:                models.ToPointer(false),
                AzureTenantId:               models.NewOptional(models.ToPointer("string")),
                BroadnetPassword:            models.ToPointer("password"),
                BroadnetSid:                 models.ToPointer("MIST"),
                BroadnetUserId:              models.ToPointer("juniper"),
                BypassWhenCloudDown:         models.ToPointer(false),
                ClickatellApiKey:            models.ToPointer("string"),
                CrossSite:                   models.ToPointer(false),
                EmailEnabled:                models.ToPointer(true),
                Enabled:                     models.ToPointer(false),
                Expire:                      models.ToPointer(1440),
                ExternalPortalUrl:           models.ToPointer("string"),
                FacebookClientId:            models.NewOptional(models.ToPointer("string")),
                FacebookClientSecret:        models.NewOptional(models.ToPointer("string")),
                FacebookEmailDomains:        []string{
                    "string",
                },
                FacebookEnabled:             models.ToPointer(false),
                Forward:                     models.ToPointer(false),
                ForwardUrl:                  models.NewOptional(models.ToPointer("https://abc.com/promotions")),
                GoogleClientId:              models.NewOptional(models.ToPointer("string")),
                GoogleClientSecret:          models.NewOptional(models.ToPointer("string")),
                GoogleEmailDomains:          []string{
                    "mydomain.edu",
                    "mydomain.org",
                },
                GoogleEnabled:               models.ToPointer(false),
                GupshupPassword:             models.ToPointer("string"),
                GupshupUserid:               models.ToPointer("string"),
                MicrosoftClientId:           models.NewOptional(models.ToPointer("string")),
                MicrosoftClientSecret:       models.NewOptional(models.ToPointer("string")),
                MicrosoftEmailDomains:       []string{
                    "string",
                },
                MicrosoftEnabled:            models.ToPointer(false),
                PassphraseEnabled:           models.ToPointer(false),
                Password:                    models.NewOptional(models.ToPointer("let me in")),
                PredefinedSponsorsEnabled:   models.ToPointer(true),
                Privacy:                     models.ToPointer(true),
                PuzzelPassword:              models.ToPointer("string"),
                PuzzelServiceId:             models.ToPointer("string"),
                PuzzelUsername:              models.ToPointer("string"),
                SmsMessageFormat:            models.ToPointer("string"),
                SmsEnabled:                  models.ToPointer(false),
                SmsProvider:                 models.ToPointer(models.WlanPortalSmsProviderEnum_TWILIO),
                SponsorAutoApprove:          models.ToPointer(false),
                SponsorEmailDomains:         []string{
                    "reserved.net",
                    "reserved.org",
                },
                SponsorEnabled:              models.ToPointer(false),
                SponsorLinkValidityDuration: models.ToPointer(models.SponsorLinkValidityDurationContainer.FromString("30")),
                SponsorNotifyAll:            models.ToPointer(false),
                SponsorStatusNotify:         models.ToPointer(false),
                Sponsors:                    models.ToPointer(models.WlanPortalSponsorsContainer.FromMapOfString(map[string]string{
                    "sponsor1@company.com": "FirstName1 LastName1",
                    "sponsor2@company.com": "FirstName2 LastName2",
                })),
                SsoDefaultRole:              models.ToPointer("string"),
                SsoForcedRole:               models.ToPointer("string"),
                SsoIdpCert:                  models.ToPointer("string"),
                SsoIdpSignAlgo:              models.ToPointer(models.WlanPortalIdpSignAlgoEnum_SHA256),
                SsoIdpSsoUrl:                models.ToPointer("string"),
                SsoIssuer:                   models.ToPointer("string"),
                SsoNameidFormat:             models.ToPointer(models.WlanPortalSsoNameidFormatEnum_EMAIL),
                TelstraClientId:             models.ToPointer("string"),
                TelstraClientSecret:         models.ToPointer("string"),
                TwilioAuthToken:             models.NewOptional(models.ToPointer("af9dac44c344a875ab5d31cb7abcdefg")),
                TwilioPhoneNumber:           models.NewOptional(models.ToPointer("+18548888888")),
                TwilioSid:                   models.NewOptional(models.ToPointer("AC72ec6ba0ec5af30e6731c5e47abcdefgh")),
                AdditionalProperties:        map[string]interface{}{
                    "portal_api_secret": interface{}("string"),
                    "portal_image": interface{}("https://url/to/image.png"),
                    "thumbnail": interface{}("string"),
                },
            }),
            PortalAllowedHostnames:               []string{
                "snapchat.com",
                "ibm.com",
            },
            PortalAllowedSubnets:                 []string{
                "63.5.3.0/24",
            },
            PortalApiSecret:                      models.NewOptional(models.ToPointer("EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY")),
            PortalDeniedHostnames:                []string{
                "msg.snapchat.com",
            },
            Qos:                                  models.ToPointer(models.WlanQos{
                Class:                models.ToPointer(models.WlanQosClassEnum_BESTEFFORT),
                Overwrite:            models.ToPointer(false),
            }),
            Radsec:                               models.ToPointer(models.Radsec{
                Enabled:              models.ToPointer(true),
                IdleTimeout:          models.ToPointer(models.RadsecIdleTimeoutContainer.FromNumber(60)),
                MxclusterIds:         []uuid.UUID{
                    uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"),
                },
                ProxyHosts:           []string{
                    "mxedge1.local",
                },
                ServerName:           models.ToPointer("radsec.abc.com"),
                Servers:              []models.RadsecServer{
                    models.RadsecServer{
                        Host:                 models.ToPointer("1.1.1.1"),
                        Port:                 models.ToPointer(1812),
                    },
                },
                UseMxedge:            models.ToPointer(true),
                UseSiteMxedge:        models.ToPointer(false),
            }),
            Rateset:                              map[string]models.WlanDatarates{
                "24": models.WlanDatarates{
                    Ht:                   models.NewOptional(models.ToPointer("00ff00ff00ff")),
                    Legacy:               []models.WlanDataratesLegacyItemEnum{
                        models.WlanDataratesLegacyItemEnum_ENUM6,
                        models.WlanDataratesLegacyItemEnum_ENUM9,
                        models.WlanDataratesLegacyItemEnum_ENUM12,
                        models.WlanDataratesLegacyItemEnum_ENUM18,
                        models.WlanDataratesLegacyItemEnum_ENUM24B,
                        models.WlanDataratesLegacyItemEnum_ENUM36,
                        models.WlanDataratesLegacyItemEnum_ENUM48,
                        models.WlanDataratesLegacyItemEnum_ENUM54,
                    },
                    MinRssi:              models.ToPointer(-70),
                    Template:             models.NewOptional(models.ToPointer(models.WlanDataratesTemplateEnum_CUSTOM)),
                    Vht:                  models.NewOptional(models.ToPointer("03ff03ff03ff01ff")),
                },
                "5": models.WlanDatarates{
                    Ht:                   models.NewOptional(models.ToPointer("00ff00ff00ff")),
                    Legacy:               []models.WlanDataratesLegacyItemEnum{
                        models.WlanDataratesLegacyItemEnum_ENUM6,
                        models.WlanDataratesLegacyItemEnum_ENUM9,
                        models.WlanDataratesLegacyItemEnum_ENUM12,
                        models.WlanDataratesLegacyItemEnum_ENUM18,
                        models.WlanDataratesLegacyItemEnum_ENUM24B,
                        models.WlanDataratesLegacyItemEnum_ENUM36,
                        models.WlanDataratesLegacyItemEnum_ENUM48,
                        models.WlanDataratesLegacyItemEnum_ENUM54,
                    },
                    MinRssi:              models.ToPointer(-70),
                    Template:             models.NewOptional(models.ToPointer(models.WlanDataratesTemplateEnum_CUSTOM)),
                    Vht:                  models.NewOptional(models.ToPointer("03ff03ff03ff01ff")),
                },
            },
            RoamMode:                             models.ToPointer(models.WlanRoamModeEnum_NONE),
            Schedule:                             models.ToPointer(models.WlanSchedule{
                Enabled:              models.ToPointer(false),
                Hours:                models.ToPointer(models.Hours{
                    Fri:                  models.ToPointer("09:00-17:00"),
                    Mon:                  models.ToPointer("09:00-17:00"),
                }),
            }),
            SleExcluded:                          models.ToPointer(false),
            Ssid:                                 "corporate",
            TemplateId:                           models.NewOptional(models.ToPointer(uuid.MustParse("c6d67e98-83ea-49f0-8812-e4abae2b68bc"))),
            UseEapolV1:                           models.ToPointer(false),
            VlanEnabled:                          models.ToPointer(false),
            VlanIds:                              models.ToPointer(models.WlanVlanIdsContainer.FromArrayOfVlanIdWithVariable4([]models.VlanIdWithVariable{
                models.VlanIdWithVariableContainer.FromNumber(3),
                models.VlanIdWithVariableContainer.FromNumber(4),
                models.VlanIdWithVariableContainer.FromNumber(5),
            })),
            VlanPooling:                          models.ToPointer(false),
            WxtagIds:                             models.NewOptional(models.ToPointer([]uuid.UUID{
                uuid.MustParse("497f6eca-6276-4993-bfeb-53e4bbba6f08"),
            })),
            WxtunnelId:                           models.NewOptional(models.ToPointer("string")),
            WxtunnelRemoteId:                     models.NewOptional(models.ToPointer("string")),
        },
    },
}

apiResponse, err := orgsSecurityPolicies.CreateOrgSecPolicy(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "name": "corporate only",
  "wlans": [
    {
      "auth": {
        "pairwise": [
          "wpa1-tkip",
          "wpa2-tkip"
        ],
        "type": "psk"
      },
      "band": "both",
      "ssid": "office"
    },
    {
      "auth": {
        "type": "open"
      },
      "band": "5",
      "ssid": "office-guest"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Sec Policy

Delete Org Security Policy

```go
DeleteOrgSecPolicy(
    ctx context.Context,
    orgId uuid.UUID,
    secpolicyId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secpolicyId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSecurityPolicies.DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Sec Policy

Get Org Security Policy

```go
GetOrgSecPolicy(
    ctx context.Context,
    orgId uuid.UUID,
    secpolicyId uuid.UUID) (
    models.ApiResponse[models.Secpolicy],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secpolicyId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Secpolicy](../../doc/models/secpolicy.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSecurityPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "name": "corporate only",
  "wlans": [
    {
      "auth": {
        "pairwise": [
          "wpa1-tkip",
          "wpa2-tkip"
        ],
        "type": "psk"
      },
      "band": "both",
      "ssid": "office"
    },
    {
      "auth": {
        "type": "open"
      },
      "band": "5",
      "ssid": "office-guest"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Sec Policies

Get List of Org Security Policies

```go
ListOrgSecPolicies(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Secpolicy],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Secpolicy](../../doc/models/secpolicy.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSecurityPolicies.ListOrgSecPolicies(ctx, orgId, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "name": "corporate only",
    "wlans": [
      {
        "auth": {
          "pairwise": [
            "wpa1-tkip",
            "wpa2-tkip"
          ],
          "type": "psk"
        },
        "band": "both",
        "ssid": "office"
      },
      {
        "auth": {
          "type": "open"
        },
        "band": "5",
        "ssid": "office-guest"
      }
    ]
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Sec Policy

Update Org Security Policy

```go
UpdateOrgSecPolicy(
    ctx context.Context,
    orgId uuid.UUID,
    secpolicyId uuid.UUID,
    body *models.Secpolicy) (
    models.ApiResponse[models.Secpolicy],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secpolicyId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Secpolicy`](../../doc/models/secpolicy.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Secpolicy](../../doc/models/secpolicy.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Secpolicy{
}

apiResponse, err := orgsSecurityPolicies.UpdateOrgSecPolicy(ctx, orgId, secpolicyId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "name": "corporate only",
  "wlans": [
    {
      "auth": {
        "pairwise": [
          "wpa1-tkip",
          "wpa2-tkip"
        ],
        "type": "psk"
      },
      "band": "both",
      "ssid": "office"
    },
    {
      "auth": {
        "type": "open"
      },
      "band": "5",
      "ssid": "office-guest"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

