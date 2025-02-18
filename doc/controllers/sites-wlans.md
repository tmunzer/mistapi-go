# Sites Wlans

```go
sitesWlans := client.SitesWlans()
```

## Class Name

`SitesWlans`

## Methods

* [Create Site Wlan](../../doc/controllers/sites-wlans.md#create-site-wlan)
* [Delete Site Wlan](../../doc/controllers/sites-wlans.md#delete-site-wlan)
* [Delete Site Wlan Portal Image](../../doc/controllers/sites-wlans.md#delete-site-wlan-portal-image)
* [Get Site Wlan](../../doc/controllers/sites-wlans.md#get-site-wlan)
* [List Site Wlan Derived](../../doc/controllers/sites-wlans.md#list-site-wlan-derived)
* [List Site Wlans](../../doc/controllers/sites-wlans.md#list-site-wlans)
* [Update Site Wlan](../../doc/controllers/sites-wlans.md#update-site-wlan)
* [Update Site Wlan Portal Template](../../doc/controllers/sites-wlans.md#update-site-wlan-portal-template)
* [Upload Site Wlan Portal Image](../../doc/controllers/sites-wlans.md#upload-site-wlan-portal-image)


# Create Site Wlan

Create Site WLAN

```go
CreateSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Wlan`](../../doc/models/wlan.md) | Body, Optional | Request Body |

## Response Type

[`models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AcctImmediateUpdate:                  models.ToPointer(false),
    AcctInterimInterval:                  models.ToPointer(0),
    AcctServers:                          []models.RadiusAcctServer{
        models.RadiusAcctServer{
            Host:                 "1.2.3.4",
            KeywrapEnabled:       models.ToPointer(true),
            KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
            KeywrapKek:           models.ToPointer("1122334455"),
            KeywrapMack:          models.ToPointer("1122334455"),
            Port:                 models.ToPointer(1813),
            Secret:               "testing123",
        },
    },
    Airwatch:                             models.ToPointer(models.WlanAirwatch{
        ApiKey:               models.ToPointer("aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\""),
        ConsoleUrl:           models.ToPointer("https://hs1.airwatchportals.com"),
        Enabled:              models.ToPointer(true),
        Password:             models.ToPointer("user1"),
        Username:             models.ToPointer("test123"),
    }),
    AllowIpv6Ndp:                         models.ToPointer(true),
    AllowMdns:                            models.ToPointer(false),
    AllowSsdp:                            models.ToPointer(false),
    ApIds:                                models.NewOptional(models.ToPointer([]uuid.UUID{
        uuid.MustParse("497f6eca-6276-4993-bfeb-53cbbbbb6f08"),
    })),
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
                Dscp:                 models.ToPointer(32),
                DstSubnet:            models.ToPointer("10.2.0.0/16"),
                SrcSubnet:            models.ToPointer("10.2.0.0/16"),
            },
        },
        Enabled:              models.ToPointer(true),
        Others:               []models.WlanAppQosOthersItem{
            models.WlanAppQosOthersItem{
                Dscp:                 models.ToPointer(32),
                DstSubnet:            models.ToPointer("10.2.0.0/16"),
                PortRanges:           models.ToPointer("80,1024-6553"),
                Protocol:             models.ToPointer("udp"),
                SrcSubnet:            models.ToPointer("10.2.0.0/16"),
            },
        },
    }),
    ApplyTo:                              models.ToPointer(models.WlanApplyToEnum_SITE),
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
            Port:                        models.ToPointer(1812),
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
        AdditionalVlanIds:    "10,20",
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
    ClientLimitDown:                      models.ToPointer(0),
    ClientLimitDownEnabled:               models.ToPointer(false),
    ClientLimitUp:                        models.ToPointer(0),
    ClientLimitUpEnabled:                 models.ToPointer(false),
    CoaServers:                           []models.CoaServer{
        models.CoaServer{
            DisableEventTimestampCheck: models.ToPointer(false),
            Enabled:                    models.ToPointer(false),
            Ip:                         "1.2.3.4",
            Port:                       models.ToPointer(3799),
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
    Ssid:                                 "demo",
}

apiResponse, err := sitesWlans.CreateSiteWlan(ctx, siteId, &body)
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
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 0,
  "client_limit_down_enabled": false,
  "client_limit_up": 0,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
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


# Delete Site Wlan

Delete Site WLAN

```go
DeleteSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesWlans.DeleteSiteWlan(ctx, siteId, wlanId)
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


# Delete Site Wlan Portal Image

Delete Site WLAN Portal Image

```go
DeleteSiteWlanPortalImage(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesWlans.DeleteSiteWlanPortalImage(ctx, siteId, wlanId)
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


# Get Site Wlan

Get Site WLAN

```go
GetSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWlans.GetSiteWlan(ctx, siteId, wlanId)
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
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 0,
  "client_limit_down_enabled": false,
  "client_limit_up": 0,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
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


# List Site Wlan Derived

Get Wlans Derived

```go
ListSiteWlanDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool,
    wlanId *string) (
    models.ApiResponse[[]models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether to resolve SITE_VARS<br>**Default**: `false` |
| `wlanId` | `*string` | Query, Optional | Filter by WLAN ID |

## Response Type

[`[]models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := true



apiResponse, err := sitesWlans.ListSiteWlanDerived(ctx, siteId, &resolve, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# List Site Wlans

Get List of Site WLANs

```go
ListSiteWlans(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesWlans.ListSiteWlans(ctx, siteId, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Update Site Wlan

Update Site WLAN

```go
UpdateSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Wlan`](../../doc/models/wlan.md) | Body, Optional | Request Body |

## Response Type

[`models.Wlan`](../../doc/models/wlan.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Wlan{
    AcctImmediateUpdate:                  models.ToPointer(false),
    AcctInterimInterval:                  models.ToPointer(0),
    AcctServers:                          []models.RadiusAcctServer{
        models.RadiusAcctServer{
            Host:                 "1.2.3.4",
            KeywrapEnabled:       models.ToPointer(true),
            KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_HEX),
            KeywrapKek:           models.ToPointer("1122334455"),
            KeywrapMack:          models.ToPointer("1122334455"),
            Port:                 models.ToPointer(1813),
            Secret:               "testing123",
        },
    },
    Airwatch:                             models.ToPointer(models.WlanAirwatch{
        ApiKey:               models.ToPointer("aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\""),
        ConsoleUrl:           models.ToPointer("https://hs1.airwatchportals.com"),
        Enabled:              models.ToPointer(true),
        Password:             models.ToPointer("user1"),
        Username:             models.ToPointer("test123"),
    }),
    AllowIpv6Ndp:                         models.ToPointer(true),
    AllowMdns:                            models.ToPointer(false),
    AllowSsdp:                            models.ToPointer(false),
    ApIds:                                models.NewOptional(models.ToPointer([]uuid.UUID{
        uuid.MustParse("497f6eca-6276-4993-bfeb-53cbbbbe6f08"),
    })),
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
                Dscp:                 models.ToPointer(32),
                DstSubnet:            models.ToPointer("10.2.0.0/16"),
                SrcSubnet:            models.ToPointer("10.2.0.0/16"),
            },
        },
        Enabled:              models.ToPointer(true),
        Others:               []models.WlanAppQosOthersItem{
            models.WlanAppQosOthersItem{
                Dscp:                 models.ToPointer(32),
                DstSubnet:            models.ToPointer("10.2.0.0/16"),
                PortRanges:           models.ToPointer("80,1024-6553"),
                Protocol:             models.ToPointer("udp"),
                SrcSubnet:            models.ToPointer("10.2.0.0/16"),
            },
        },
    }),
    ApplyTo:                              models.ToPointer(models.WlanApplyToEnum_SITE),
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
            Port:                        models.ToPointer(1812),
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
        AdditionalVlanIds:    "10,20",
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
    ClientLimitDown:                      models.ToPointer(0),
    ClientLimitDownEnabled:               models.ToPointer(false),
    ClientLimitUp:                        models.ToPointer(0),
    ClientLimitUpEnabled:                 models.ToPointer(false),
    CoaServers:                           []models.CoaServer{
        models.CoaServer{
            DisableEventTimestampCheck: models.ToPointer(false),
            Enabled:                    models.ToPointer(false),
            Ip:                         "1.2.3.4",
            Port:                       models.ToPointer(3799),
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
    Ssid:                                 "demo",
}

apiResponse, err := sitesWlans.UpdateSiteWlan(ctx, siteId, wlanId, &body)
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
  "allow_ipv6_ndp": true,
  "allow_mdns": false,
  "allow_ssdp": false,
  "arp_filter": false,
  "band_steer": false,
  "band_steer_force_band5": false,
  "bands": [
    "24",
    "5"
  ],
  "block_blacklist_clients": false,
  "bonjour": {
    "additional_vlan_ids": "10, 20",
    "enabled": false,
    "services": {
      "airplay": {
        "radius_groups": [
          "teachers"
        ],
        "scope": "same_ap"
      }
    }
  },
  "client_limit_down": 0,
  "client_limit_down_enabled": false,
  "client_limit_up": 0,
  "client_limit_up_enabled": false,
  "disable_11ax": false,
  "disable_ht_vht_rates": false,
  "disable_uapsd": false,
  "disable_v1_roam_notify": false,
  "disable_v2_roam_notify": false,
  "disable_wmm": false,
  "dynamic_vlan": {
    "default_vlan_id": 999,
    "enabled": false,
    "local_vlan_ids": [
      1
    ],
    "type": "airespace-interface-name",
    "vlans": {
      "131": "default",
      "322": "fast,video"
    }
  },
  "enable_local_keycaching": false,
  "enable_wireless_bridging": false,
  "enabled": true,
  "fast_dot1x_timers": false,
  "hide_ssid": false,
  "hostname_ie": false,
  "ssid": "demo"
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


# Update Site Wlan Portal Template

Update a Portal Template

#### Sponsor Email Template

Sponsor Email Template supports following template variables:

| **Name** | **Description** |
| --- | --- |
| approve_url | Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized |
| deny_url | Renders URL to reject the request |
| guest_email | Renders Email ID of the guest |
| guest_name | Renders Name of the guest |
| field1 | Renders value of the Custom Field 1 |
| field2 | Renders value of the Custom Field 2 |
| company | Renders value of the Company field |
| sponsor_link_validity_duration | Renders validity time of the request (i.e. Approve/Deny URL) |
| auth_expire_minutes | Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |

```go
UpdateSiteWlanPortalTemplate(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    body *models.WlanPortalTemplate) (
    models.ApiResponse[models.WlanPortalTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WlanPortalTemplate`](../../doc/models/wlan-portal-template.md) | Body, Optional | Request Body |

## Response Type

[`models.WlanPortalTemplate`](../../doc/models/wlan-portal-template.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WlanPortalTemplate{
    AdditionalProperties: map[string]interface{}{
        "accessCodeAlternateEmail": interface{}("string"),
        "alignment": interface{}("string"),
        "authButtonAmazon": interface{}("string"),
        "authButtonAzure": interface{}("string"),
        "authButtonEmail": interface{}("string"),
        "authButtonFacebook": interface{}("string"),
        "authButtonGoogle": interface{}("string"),
        "authButtonMicrosoft": interface{}("string"),
        "authButtonPassphrase": interface{}("string"),
        "authButtonSms": interface{}("string"),
        "authButtonSponsor": interface{}("string"),
        "authLabel": interface{}("string"),
        "backLink": interface{}("string"),
        "color": interface{}("string"),
        "colorDark": interface{}("string"),
        "colorLight": interface{}("string"),
        "company": interface{}("true"),
        "companyError": interface{}("string"),
        "companyLabel": interface{}("string"),
        "email": interface{}("true"),
        "emailAccessDomainError": interface{}("string"),
        "emailCancel": interface{}("string"),
        "emailCodeError": interface{}("string"),
        "emailError": interface{}("string"),
        "emailFieldLabel": interface{}("string"),
        "emailLabel": interface{}("string"),
        "emailMessage": interface{}("string"),
        "emailSubmit": interface{}("string"),
        "emailTitle": interface{}("string"),
        "field1": interface{}("true"),
        "field1Error": interface{}("string"),
        "field1Label": interface{}("string"),
        "field1Required": interface{}("true"),
        "field2": interface{}("true"),
        "field2Error": interface{}("string"),
        "field2Label": interface{}("string"),
        "field2Required": interface{}("true"),
        "field3": interface{}("true"),
        "field3Error": interface{}("string"),
        "field3Label": interface{}("string"),
        "field3Required": interface{}("true"),
        "field4": interface{}("true"),
        "field4Error": interface{}("string"),
        "field4Label": interface{}("string"),
        "field4Required": interface{}("true"),
        "message": interface{}("string"),
        "name": interface{}("true"),
        "nameError": interface{}("string"),
        "nameLabel": interface{}("string"),
        "optout": interface{}("true"),
        "optoutLabel": interface{}("string"),
        "pageTitle": interface{}("string"),
        "passphraseCancel": interface{}("string"),
        "passphraseError": interface{}("string"),
        "passphraseLabel": interface{}("string"),
        "passphraseMessage": interface{}("string"),
        "passphraseSubmit": interface{}("string"),
        "passphraseTitle": interface{}("string"),
        "poweredBy": interface{}("true"),
        "requiredFieldLabel": interface{}("string"),
        "signInLabel": interface{}("string"),
        "smsCarrierDefault": interface{}("string"),
        "smsCarrierError": interface{}("string"),
        "smsCarrierFieldLabel": interface{}("string"),
        "smsCodeCancel": interface{}("string"),
        "smsCodeError": interface{}("string"),
        "smsCodeFieldLabel": interface{}("string"),
        "smsCodeMessage": interface{}("string"),
        "smsCodeSubmit": interface{}("string"),
        "smsCodeTitle": interface{}("string"),
        "smsCountryFieldLabel": interface{}("string"),
        "smsCountryFormat": interface{}("string"),
        "smsHaveAccessCode": interface{}("string"),
        "smsMessageFormat": interface{}("string"),
        "smsNumberCancel": interface{}("string"),
        "smsNumberError": interface{}("string"),
        "smsNumberFieldLabel": interface{}("string"),
        "smsNumberFormat": interface{}("string"),
        "smsNumberMessage": interface{}("string"),
        "smsNumberSubmit": interface{}("string"),
        "smsNumberTitle": interface{}("string"),
        "smsUsernameFormat": interface{}("string"),
        "smsValidityDuration": interface{}("0"),
        "sponsorBackLink": interface{}("string"),
        "sponsorCancel": interface{}("string"),
        "sponsorEmail": interface{}("string"),
        "sponsorEmailError": interface{}("string"),
        "sponsorEmailTemplate": interface{}("string"),
        "sponsorInfoApproved": interface{}("string"),
        "sponsorInfoDenied": interface{}("string"),
        "sponsorInfoPending": interface{}("string"),
        "sponsorName": interface{}("string"),
        "sponsorNameError": interface{}("string"),
        "sponsorNotePending": interface{}("string"),
        "sponsorStatusApproved": interface{}("string"),
        "sponsorStatusDenied": interface{}("string"),
        "sponsorStatusPending": interface{}("string"),
        "sponsorSubmit": interface{}("string"),
        "tos": interface{}("true"),
        "tosAcceptLabel": interface{}("string"),
        "tosError": interface{}("string"),
        "tosLink": interface{}("string"),
        "tosText": interface{}("string"),
    },
}

apiResponse, err := sitesWlans.UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, &body)
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
  "portal_template": {
    "accessCodeAlternateEmail": "string",
    "alignment": "left",
    "authButtonAmazon": "string",
    "authButtonAzure": "string",
    "authButtonEmail": "string",
    "authButtonFacebook": "string",
    "authButtonGoogle": "string",
    "authButtonMicrosoft": "string",
    "authButtonPassphrase": "string",
    "authButtonSms": "string",
    "authButtonSponsor": "string",
    "authLabel": "string",
    "backLink": "string",
    "color": "string",
    "colorDark": "string",
    "colorLight": "string",
    "company": true,
    "companyError": "string",
    "companyLabel": "string",
    "email": true,
    "emailAccessDomainError": "string",
    "emailCancel": "string",
    "emailCodeError": "string",
    "emailError": "string",
    "emailFieldLabel": "string",
    "emailLabel": "string",
    "emailMessage": "string",
    "emailSubmit": "string",
    "emailTitle": "string",
    "field1": true,
    "field1Error": "string",
    "field1Label": "string",
    "field1Required": true,
    "field2": true,
    "field2Error": "string",
    "field2Label": "string",
    "field2Required": true,
    "field3": true,
    "field3Error": "string",
    "field3Label": "string",
    "field3Required": true,
    "field4": true,
    "field4Error": "string",
    "field4Label": "string",
    "field4Required": true,
    "message": "string",
    "name": true,
    "nameError": "string",
    "nameLabel": "string",
    "optout": true,
    "optoutLabel": "string",
    "pageTitle": "string",
    "passphraseCancel": "string",
    "passphraseError": "string",
    "passphraseLabel": "string",
    "passphraseMessage": "string",
    "passphraseSubmit": "string",
    "passphraseTitle": "string",
    "poweredBy": true,
    "requiredFieldLabel": "string",
    "signInLabel": "string",
    "smsCarrierDefault": "string",
    "smsCarrierError": "string",
    "smsCarrierFieldLabel": "string",
    "smsCodeCancel": "string",
    "smsCodeError": "string",
    "smsCodeFieldLabel": "string",
    "smsCodeMessage": "string",
    "smsCodeSubmit": "string",
    "smsCodeTitle": "string",
    "smsCountryFieldLabel": "string",
    "smsCountryFormat": "string",
    "smsHaveAccessCode": "string",
    "smsMessageFormat": "string",
    "smsNumberCancel": "string",
    "smsNumberError": "string",
    "smsNumberFieldLabel": "string",
    "smsNumberFormat": "string",
    "smsNumberMessage": "string",
    "smsNumberSubmit": "string",
    "smsNumberTitle": "string",
    "smsUsernameFormat": "string",
    "smsValidityDuration": 5,
    "sponsorBackLink": "string",
    "sponsorCancel": "string",
    "sponsorEmail": "string",
    "sponsorEmailError": "string",
    "sponsorEmailTemplate": "string",
    "sponsorInfoApproved": "string",
    "sponsorInfoDenied": "string",
    "sponsorInfoPending": "string",
    "sponsorName": "string",
    "sponsorNameError": "string",
    "sponsorNotePending": "string",
    "sponsorStatusApproved": "string",
    "sponsorStatusDenied": "string",
    "sponsorStatusPending": "string",
    "sponsorSubmit": "string",
    "tos": true,
    "tosAcceptLabel": "string",
    "tosError": "string",
    "tosLink": "string",
    "tosText": "string"
  }
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


# Upload Site Wlan Portal Image

WLAN Portal Image Upload

```go
UploadSiteWlanPortalImage(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Binary file |
| `json` | `*string` | Form, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })



resp, err := sitesWlans.UploadSiteWlanPortalImage(ctx, siteId, wlanId, file, nil)
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

