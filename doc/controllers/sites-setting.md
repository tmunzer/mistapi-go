# Sites Setting

```go
sitesSetting := client.SitesSetting()
```

## Class Name

`SitesSetting`

## Methods

* [Create Site Watched Stations](../../doc/controllers/sites-setting.md#create-site-watched-stations)
* [Create Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-allowlist)
* [Create Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#create-site-wireless-clients-blocklist)
* [Delete Site Watched Stations](../../doc/controllers/sites-setting.md#delete-site-watched-stations)
* [Delete Site Wireless Clients Allowlist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-allowlist)
* [Delete Site Wireless Clients Blocklist](../../doc/controllers/sites-setting.md#delete-site-wireless-clients-blocklist)
* [Get Site Setting](../../doc/controllers/sites-setting.md#get-site-setting)
* [Get Site Setting Derived](../../doc/controllers/sites-setting.md#get-site-setting-derived)
* [Update Site Settings](../../doc/controllers/sites-setting.md#update-site-settings)


# Create Site Watched Stations

This endpoint is to provide list of client macs for annotation as watched station.

Retrieve the current clients list from `watched_station_url` under Site:Setting

```go
CreateSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWatchedStations(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Create Site Wireless Clients Allowlist

This endpoint is to provide list of client macs for annotation as whitelist.

Retrieve the current clients list from `whitelist_url` under Site:Setting

```go
CreateSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "683b679ac024",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsAllowlist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Create Site Wireless Clients Blocklist

This endpoint is to provide list of client macs for annotation blacklist.

Retrieve the current clients list `blacklist_url` under Site:Setting

```go
CreateSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.MacAddresses`](../../doc/models/mac-addresses.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := sitesSetting.CreateSiteWirelessClientsBlocklist(ctx, siteId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
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


# Delete Site Watched Stations

Delete Site Watched Station Clients

```go
DeleteSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWatchedStations(ctx, siteId)
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


# Delete Site Wireless Clients Allowlist

Delete Site Whitelist Station Clients

```go
DeleteSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsAllowlist(ctx, siteId)
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


# Delete Site Wireless Clients Blocklist

Delete Site Blacklist Station Clients

```go
DeleteSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesSetting.DeleteSiteWirelessClientsBlocklist(ctx, siteId)
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


# Get Site Setting

Get the Derived Site Settings, generated by merging the Org level templates (network templates, gateway templates) and the Site level configuration. If the same parameter is defined in both scopes, the Site level one is used.

```go
GetSiteSetting(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSetting.GetSiteSetting(ctx, siteId)
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


# Get Site Setting Derived

Get Site Settings

```go
GetSiteSettingDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSetting.GetSiteSettingDerived(ctx, siteId)
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


# Update Site Settings

Update Site Settings

```go
UpdateSiteSettings(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.SiteSetting) (
    models.ApiResponse[models.SiteSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SiteSetting`](../../doc/models/site-setting.md) | Body, Optional | Request Body |

## Response Type

[`models.SiteSetting`](../../doc/models/site-setting.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteSetting{
    AdditionalConfigCmds:            []string{
        "set snmp community public",
    },
    Analytic:                        models.ToPointer(models.SiteSettingAnalytic{
        Enabled: models.ToPointer(false),
    }),
    ApMatching:                      models.ToPointer(models.SiteSettingApMatching{
        Enabled: models.ToPointer(true),
        Rules:   []models.SiteSettingApMatchingRule{
            models.SiteSettingApMatchingRule{
            },
        },
    }),
    ApPortConfig:                    models.ToPointer(models.SiteSettingApPortConfig{
        ModelSpecific: map[string]models.ApPortConfig{
            "AP32": models.ApPortConfig{
            },
        },
    }),
    AutoUpgrade:                     models.ToPointer(models.SiteSettingAutoUpgrade{
        CustomVersions: map[string]string{
            "AP21": "stable",
            "AP41": "0.1.5135",
            "AP61": "0.1.7215",
        },
        DayOfWeek:      models.ToPointer(models.DayOfWeekEnum("sun")),
        Enabled:        models.ToPointer(false),
        TimeOfDay:      models.ToPointer("12:00"),
        Version:        models.ToPointer(models.SiteAutoUpgradeVersionEnum("beta")),
    }),
    ConfigAutoRevert:                models.ToPointer(false),
    DeviceUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
    DnsServers:                      []string{
        "string",
    },
    DnsSuffix:                       []string{
        "string",
    },
    Engagement:                      models.ToPointer(models.SiteEngagement{
        DwellTagNames: models.ToPointer(models.SiteEngagementDwellTagNames{
            Bounce:    models.ToPointer("Bounce"),
            Engaged:   models.ToPointer("Engaged"),
            Passerby:  models.ToPointer("Passer By"),
            Stationed: models.ToPointer("Stationed"),
        }),
        DwellTags:     models.ToPointer(models.SiteEngagementDwellTags{
            Engaged:   models.NewOptional(models.ToPointer("300-14400")),
            Stationed: models.NewOptional(models.ToPointer("14400-43200")),
        }),
        Hours:         models.ToPointer(models.Hours{
            Fri: models.ToPointer("09:00-17:00"),
            Mon: models.ToPointer("09:00-17:00"),
            Sat: models.ToPointer("09:00-12:00"),
            Sun: models.ToPointer("09:00-12:00"),
            Thu: models.ToPointer("09:00-17:00"),
            Tue: models.ToPointer("09:00-17:00"),
            Wed: models.ToPointer("09:00-17:00"),
        }),
        MaxDwell:      models.ToPointer(43200),
        MinDwell:      models.ToPointer(0),
    }),
    EvpnOptions:                     models.ToPointer(models.EvpnOptions{
        AutoLoopbackSubnet:  models.ToPointer("100.101.0.0/16"),
        AutoRouterIdSubnet:  models.ToPointer("100.100.0.0/24"),
        CoreAsBorder:        models.ToPointer(false),
        Overlay:             models.ToPointer(models.EvpnOptionsOverlay{
            As: models.ToPointer(65000),
        }),
        PerVlanVgaV4Mac:     models.ToPointer(false),
        RoutedAt:            models.ToPointer(models.EvpnOptionsRoutedAtEnum("edge")),
        Underlay:            models.ToPointer(models.EvpnOptionsUnderlay{
            AsBase:         models.ToPointer(65001),
            RoutedIdPrefix: models.ToPointer("/24"),
            Subnet:         models.ToPointer("10.255.240.0/20"),
        }),
    }),
    GatewayAdditionalConfigCmds:     []string{
        "set snmp community public",
    },
    GatewayMgmt:                     models.ToPointer(models.SiteSettingGatewayMgmt{
        AdminSshkeys:               []string{
            "string",
        },
        AppProbing:                 models.ToPointer(models.AppProbing{
            Apps:       []string{
                "string",
            },
            CustomApps: []models.AppProbingCustomApp{
                models.AppProbingCustomApp{
                    AppType:    models.ToPointer("string"),
                    Name:       models.ToPointer("string"),
                    Protocol:   models.ToPointer(models.AppProbingCustomAppProtocolEnum("http")),
                },
            },
            Enabled:    models.ToPointer(true),
        }),
        AppUsage:                   models.ToPointer(true),
        AutoSignatureUpdate:        models.ToPointer(models.SiteSettingGatewayMgmtAutoSignatureUpdate{
            DayOfWeek: models.ToPointer(models.DayOfWeekEnum("any")),
            Enable:    models.ToPointer(true),
            TimeOfDay: models.ToPointer("string"),
        }),
        ConfigRevertTimer:          models.ToPointer(10),
        ProbeHosts:                 []string{
            "string",
        },
        RootPassword:               models.ToPointer("string"),
        SecurityLogSourceAddress:   models.ToPointer("192.168.1.1"),
        SecurityLogSourceInterface: models.ToPointer("string"),
    }),
    Led:                             models.ToPointer(models.ApLed{
        Brightness: models.ToPointer(255),
        Enabled:    models.ToPointer(true),
    }),
    MxedgeMgmt:                      models.ToPointer(models.MxedgeMgmt{
        MistPassword: models.ToPointer("MIST_PASSWORD"),
        RootPassword: models.ToPointer("ROOT_PASSWORD"),
    }),
    Networks:                        map[string]models.SwitchNetwork{
        "property1": models.SwitchNetwork{
            Subnet:          models.ToPointer("string"),
            VlanId:          models.VlanIdWithVariableContainer.FromNumber(10),
        },
        "property2": models.SwitchNetwork{
            Subnet:          models.ToPointer("string"),
            VlanId:          models.VlanIdWithVariableContainer.FromNumber(10),
        },
    },
    NtpServers:                      []string{
        "string",
    },
    Occupancy:                       models.ToPointer(models.SiteOccupancyAnalytics{
        AssetsEnabled:             models.ToPointer(false),
        ClientsEnabled:            models.ToPointer(true),
        MinDuration:               models.ToPointer(3000),
        SdkclientsEnabled:         models.ToPointer(false),
        UnconnectedClientsEnabled: models.ToPointer(false),
    }),
    OspfAreas:                       map[string]models.OspfArea{
        "property1": models.OspfArea{
            IncludeLoopback: models.ToPointer(false),
            Networks:        map[string]models.OspfAreasNetwork{
                "corp": models.OspfAreasNetwork{
                    AuthKeys:               map[string]string{
                        "1": "auth-key-1",
                    },
                    AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum("md5")),
                    BfdMinimumInterval:     models.ToPointer(500),
                    DeadInterval:           models.ToPointer(40),
                    HelloInterval:          models.ToPointer(10),
                    InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum("nbma")),
                    Metric:                 models.NewOptional(models.ToPointer(10000)),
                },
                "guest": models.OspfAreasNetwork{
                    Passive:                models.ToPointer(true),
                },
            },
            Type:            models.ToPointer(models.OspfAreaTypeEnum("default")),
        },
        "property2": models.OspfArea{
            IncludeLoopback: models.ToPointer(false),
            Networks:        map[string]models.OspfAreasNetwork{
                "corp": models.OspfAreasNetwork{
                    AuthKeys:               map[string]string{
                        "1": "auth-key-1",
                    },
                    AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum("md5")),
                    BfdMinimumInterval:     models.ToPointer(500),
                    DeadInterval:           models.ToPointer(40),
                    HelloInterval:          models.ToPointer(10),
                    InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum("nbma")),
                    Metric:                 models.NewOptional(models.ToPointer(10000)),
                },
                "guest": models.OspfAreasNetwork{
                    Passive:                models.ToPointer(true),
                },
            },
            Type:            models.ToPointer(models.OspfAreaTypeEnum("default")),
        },
    },
    PersistConfigOnDevice:           models.ToPointer(false),
    PortMirroring:                   map[string]models.SwitchPortMirroringProperty{
        "property1": models.SwitchPortMirroringProperty{
            InputNetworksIngress: []string{
                "corp",
            },
            InputPortIdsEgress:   []string{
                "ge-0/0/3",
            },
            InputPortIdsIngress:  []string{
                "ge-0/0/3",
            },
            OutputNetwork:        models.ToPointer("analyze"),
            OutputPortId:         models.ToPointer("ge-0/0/5"),
        },
        "property2": models.SwitchPortMirroringProperty{
            InputNetworksIngress: []string{
                "corp",
            },
            InputPortIdsEgress:   []string{
                "ge-0/0/3",
            },
            InputPortIdsIngress:  []string{
                "ge-0/0/3",
            },
            OutputNetwork:        models.ToPointer("analyze"),
            OutputPortId:         models.ToPointer("ge-0/0/5"),
        },
    },
    PortUsages:                      map[string]models.SwitchPortUsage{
        "dynamic": models.SwitchPortUsage{
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("dynamic")),
            ResetDefaultWhen:                         models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum("link_down")),
            Rules:                                    []models.SwitchPortUsageDynamicRule{
                models.SwitchPortUsageDynamicRule{
                    Equals:     models.ToPointer("string"),
                    EqualsAny:  []string{
                        "string",
                    },
                    Expression: models.ToPointer("string"),
                    Src:        models.SwitchPortUsageDynamicRuleSrcEnum("lldp_chassis_id"),
                    Usage:      models.ToPointer("string"),
                },
            },
        },
        "property1": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum("auto")),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(0),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("access")),
            Mtu:                                      models.ToPointer(0),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortAuth:                                 models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum("dot1x"))),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer("string"),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.ToPointer("string"),
        },
        "property2": models.SwitchPortUsage{
            AllNetworks:                              models.ToPointer(false),
            AllowDhcpd:                               models.ToPointer(true),
            BypassAuthWhenServerDown:                 models.ToPointer(true),
            Description:                              models.ToPointer("string"),
            DisableAutoneg:                           models.ToPointer(false),
            Disabled:                                 models.ToPointer(false),
            Duplex:                                   models.ToPointer(models.SwitchPortUsageDuplexEnum("auto")),
            EnableMacAuth:                            models.ToPointer(true),
            EnableQos:                                models.ToPointer(true),
            GuestNetwork:                             models.NewOptional(models.ToPointer("string")),
            MacAuthOnly:                              models.ToPointer(true),
            MacLimit:                                 models.ToPointer(0),
            Mode:                                     models.ToPointer(models.SwitchPortUsageModeEnum("access")),
            Mtu:                                      models.ToPointer(0),
            Networks:                                 []string{
                "string",
            },
            PersistMac:                               models.ToPointer(false),
            PoeDisabled:                              models.ToPointer(false),
            PortNetwork:                              models.ToPointer("string"),
            Speed:                                    models.ToPointer("string"),
            StormControl:                             models.ToPointer(models.SwitchPortUsageStormControl{
                NoBroadcast:           models.ToPointer(false),
                NoMulticast:           models.ToPointer(false),
                NoRegisteredMulticast: models.ToPointer(false),
                NoUnknownUnicast:      models.ToPointer(false),
                Percentage:            models.ToPointer(80),
            }),
            StpEdge:                                  models.ToPointer(true),
            VoipNetwork:                              models.ToPointer("string"),
        },
    },
    Proxy:                           models.ToPointer(models.Proxy{
        Url: models.ToPointer("http://proxy.internal:8080/*"),
    }),
    Rogue:                           models.ToPointer(models.SiteRogue{
        Enabled:           models.ToPointer(false),
        HoneypotEnabled:   models.ToPointer(false),
        MinDuration:       models.ToPointer(10),
        MinRssi:           models.ToPointer(-80),
        WhitelistedBssids: []string{
            "NeighborSSID",
        },
        WhitelistedSsids:  []string{
            "cc:8e:6f:d4:bf:16",
            "cc-8e-6f-d4-bf-16",
            "cc-73-*",
            "cc:82:*",
        },
    }),
    SimpleAlert:                     models.ToPointer(models.SimpleAlert{
        ArpFailure:  models.ToPointer(models.SimpleAlertArpFailure{
            ClientCount:   models.ToPointer(10),
            Duration:      models.ToPointer(20),
            IncidentCount: models.ToPointer(10),
        }),
        DhcpFailure: models.ToPointer(models.SimpleAlertDhcpFailure{
            ClientCount:   models.ToPointer(10),
            Duration:      models.ToPointer(10),
            IncidentCount: models.ToPointer(20),
        }),
        DnsFailure:  models.ToPointer(models.SimpleAlertDnsFailure{
            ClientCount:   models.ToPointer(20),
            Duration:      models.ToPointer(10),
            IncidentCount: models.ToPointer(30),
        }),
    }),
    Skyatp:                          models.ToPointer(models.SiteSettingSkyatp{
        Enabled:          models.ToPointer(true),
        SendIpMacMapping: models.ToPointer(true),
    }),
    SrxApp:                          models.ToPointer(models.SiteSettingSrxApp{
        Enabled: models.ToPointer(false),
    }),
    SshKeys:                         []string{
        "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host",
    },
    Ssr:                             models.ToPointer(models.SiteSettingSsr{
        ConductorHosts: []string{
            "\"1.1.1.1\", \"2.2.2.2\"",
        },
        DisableStats:   models.ToPointer(true),
    }),
    StatusPortal:                    models.ToPointer(models.SiteSettingStatusPortal{
        Enabled:   models.ToPointer(false),
        Hostnames: []string{
            "my.misty.com",
        },
    }),
    Vars:                            map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    Vna:                             models.ToPointer(models.SiteSettingVna{
        Enabled: models.ToPointer(false),
    }),
    WanVna:                          models.ToPointer(models.SiteSettingWanVna{
        Enabled: models.ToPointer(false),
    }),
    Wids:                            models.ToPointer(models.SiteWids{
        RepeatedAuthFailures: models.ToPointer(models.SiteWidsRepeatedAuthFailures{
            Duration:  models.ToPointer(60),
            Threshold: models.ToPointer(0),
        }),
    }),
    Wifi:                            models.ToPointer(models.SiteWifi{
        CiscoEnabled:                      models.ToPointer(true),
        Disable11k:                        models.ToPointer(false),
        DisableRadiosWhenPowerConstrained: models.ToPointer(false),
        EnableArpSpoofCheck:               models.ToPointer(false),
        EnableSharedRadioScanning:         models.ToPointer(true),
        Enabled:                           models.ToPointer(true),
        LocateConnected:                   models.ToPointer(true),
        LocateUnconnected:                 models.ToPointer(false),
        MeshAllowDfs:                      models.ToPointer(false),
        MeshEnableCrm:                     models.ToPointer(false),
        MeshEnabled:                       models.ToPointer(false),
        MeshPsk:                           models.NewOptional(models.ToPointer("string")),
        MeshSsid:                          models.NewOptional(models.ToPointer("string")),
        ProxyArp:                          models.NewOptional(models.ToPointer(models.SiteWifiProxyArpEnum("default"))),
    }),
    WiredVna:                        models.ToPointer(models.SiteSettingWiredVna{
        Enabled: models.ToPointer(false),
    }),
    ZoneOccupancyAlert:              models.ToPointer(models.SiteZoneOccupancyAlert{
        EmailNotifiers: []string{
            "foo@juniper.net",
            "bar@juniper.net",
        },
        Enabled:        models.ToPointer(false),
        Threshold:      models.ToPointer(5),
    }),
}

apiResponse, err := sitesSetting.UpdateSiteSettings(ctx, siteId, &body)
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

