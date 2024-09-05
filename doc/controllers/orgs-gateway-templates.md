# Orgs Gateway Templates

```go
orgsGatewayTemplates := client.OrgsGatewayTemplates()
```

## Class Name

`OrgsGatewayTemplates`

## Methods

* [Create Org Gateway Template](../../doc/controllers/orgs-gateway-templates.md#create-org-gateway-template)
* [Delete Org Gateway Template](../../doc/controllers/orgs-gateway-templates.md#delete-org-gateway-template)
* [Get Org Gateway Template](../../doc/controllers/orgs-gateway-templates.md#get-org-gateway-template)
* [List Org Gateway Templates](../../doc/controllers/orgs-gateway-templates.md#list-org-gateway-templates)
* [Update Org Gateway Template](../../doc/controllers/orgs-gateway-templates.md#update-org-gateway-template)


# Create Org Gateway Template

Create Org Gateway Template

```go
CreateOrgGatewayTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.GatewayTemplate) (
    models.ApiResponse[models.GatewayTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.GatewayTemplate`](../../doc/models/gateway-template.md) | Body, Optional | Gateway Template |

## Response Type

[`models.GatewayTemplate`](../../doc/models/gateway-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.GatewayTemplate{
    DhcpdConfig:           models.ToPointer(models.DhcpdConfig{
    }),
    DnsOverride:           models.ToPointer(true),
    DnsServers:            []string{
        "10.3.20.201",
        "10.3.51.222",
        "1.1.1.1",
    },
    DnsSuffix:             []string{
        "example.com",
    },
    ExtraRoutes:           map[string]models.GatewayExtraRoute{
        "10.101.0.0/16": models.GatewayExtraRoute{
            Via: models.ToPointer("10.3.100.10"),
        },
    },
    IpConfigs:             map[string]models.GatewayIpConfigProperty{
        "Corp-Core": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.100.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-Mgmt": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.172.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-lan": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.171.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
    },
    Name:                  "ITParis",
    NtpOverride:           models.ToPointer(true),
    NtpServers:            []string{
        "10.3.51.222",
    },
    PathPreferences:       map[string]models.GatewayPathPreferences{
        "core": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Core",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "lab": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-lan",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "mgmt": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Mgmt",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "untrust": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Name:           models.ToPointer("wan"),
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("wan")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
    },
    PortConfig:            map[string]models.GatewayPortConfig{
        "ge-0/0/0": models.GatewayPortConfig{
            IpConfig:        models.ToPointer(models.GatewayPortConfigIpConfig{
                Gateway:       models.ToPointer("192.168.1.1"),
                Ip:            models.ToPointer("192.168.1.9"),
                Netmask:       models.ToPointer("/24"),
                Type:          models.ToPointer(models.GatewayWanTypeEnum("static")),
            }),
            Name:            models.ToPointer("wan"),
            Redundant:       models.ToPointer(false),
            TrafficShaping:  models.ToPointer(models.GatewayTrafficShaping{
                Enabled:          models.ToPointer(false),
            }),
            Usage:           models.GatewayPortUsageEnum("wan"),
            WanType:         models.ToPointer(models.GatewayPortWanTypeEnum("broadband")),
        },
        "ge-0/0/6-7": models.GatewayPortConfig{
            Networks:        []string{
                "Corp-lan",
                "Corp-Mgmt",
                "Corp-Core",
            },
            Usage:           models.GatewayPortUsageEnum("lan"),
        },
    },
    ServicePolicies:       []models.ServicePolicy{
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-Internal"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "drive",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-internet"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-core"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "internal_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-mxe-tt-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "mxedge-updates",
                "radsec",
                "icmp",
                "internet_dns",
                "internet_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mxe-data-0-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-mgt"),
            PathPreference:  models.ToPointer("mgmt"),
            Services:        []string{
                "mgmt",
            },
            Tenants:         []string{
                "domain.Corp-Core",
                "lan.Corp-Core",
                "servers.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-edge-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-iot"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "iot",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "servers-hassio.Corp-Core",
                "servers-kubes.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("tanker-to-cctv"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "rtsp",
            },
            Tenants:         []string{
                "servers-tanker.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("core-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "domain.Corp-Core",
                "servers.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("deny")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("iot-upgrade-cctv"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "motioneye",
                "nodejs",
                "raspbian",
            },
            Tenants:         []string{
                "printer",
            },
        },
    },
    Type:                  models.ToPointer(models.GatewayTemplateTypeEnum("standalone")),
}

apiResponse, err := orgsGatewayTemplates.CreateOrgGatewayTemplate(ctx, orgId, &body)
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


# Delete Org Gateway Template

Delete Organization Gateway Template

```go
DeleteOrgGatewayTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    gatewaytemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `gatewaytemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

gatewaytemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsGatewayTemplates.DeleteOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
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


# Get Org Gateway Template

Get Organization Gateway Template details

```go
GetOrgGatewayTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    gatewaytemplateId uuid.UUID) (
    models.ApiResponse[models.GatewayTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `gatewaytemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GatewayTemplate`](../../doc/models/gateway-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

gatewaytemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsGatewayTemplates.GetOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
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


# List Org Gateway Templates

Get List of Org Gateway Templates

```go
ListOrgGatewayTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.GatewayTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.GatewayTemplate`](../../doc/models/gateway-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsGatewayTemplates.ListOrgGatewayTemplates(ctx, orgId, &limit, &page)
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


# Update Org Gateway Template

Update Organization Gateway Template

```go
UpdateOrgGatewayTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    gatewaytemplateId uuid.UUID,
    body *models.GatewayTemplate) (
    models.ApiResponse[models.GatewayTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `gatewaytemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.GatewayTemplate`](../../doc/models/gateway-template.md) | Body, Optional | Gateway Template |

## Response Type

[`models.GatewayTemplate`](../../doc/models/gateway-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

gatewaytemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.GatewayTemplate{
    DhcpdConfig:           models.ToPointer(models.DhcpdConfig{
    }),
    DnsOverride:           models.ToPointer(true),
    DnsServers:            []string{
        "10.3.20.201",
        "10.3.51.222",
        "1.1.1.1",
    },
    DnsSuffix:             []string{
        "example.com",
    },
    ExtraRoutes:           map[string]models.GatewayExtraRoute{
        "10.101.0.0/16": models.GatewayExtraRoute{
            Via: models.ToPointer("10.3.100.10"),
        },
    },
    IpConfigs:             map[string]models.GatewayIpConfigProperty{
        "Corp-Core": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.100.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-Mgmt": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.172.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
        "Corp-lan": models.GatewayIpConfigProperty{
            Ip:           models.ToPointer("10.3.171.9"),
            Netmask:      models.ToPointer("/24"),
            Type:         models.ToPointer(models.IpTypeEnum("static")),
        },
    },
    Name:                  "ITParis",
    NtpOverride:           models.ToPointer(true),
    NtpServers:            []string{
        "10.3.51.222",
    },
    PathPreferences:       map[string]models.GatewayPathPreferences{
        "core": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Core",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "lab": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-lan",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "mgmt": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:       []string{
                        "Corp-Mgmt",
                    },
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("local")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
        "untrust": models.GatewayPathPreferences{
            Paths:    []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Name:           models.ToPointer("wan"),
                    Type:           models.ToPointer(models.GatewayPathTypeEnum("wan")),
                },
            },
            Strategy: models.ToPointer(models.GatewayPathStrategyEnum("ordered")),
        },
    },
    PortConfig:            map[string]models.GatewayPortConfig{
        "ge-0/0/0": models.GatewayPortConfig{
            IpConfig:        models.ToPointer(models.GatewayPortConfigIpConfig{
                Gateway:       models.ToPointer("192.168.1.1"),
                Ip:            models.ToPointer("192.168.1.9"),
                Netmask:       models.ToPointer("/24"),
                Type:          models.ToPointer(models.GatewayWanTypeEnum("static")),
            }),
            Name:            models.ToPointer("wan"),
            Redundant:       models.ToPointer(false),
            TrafficShaping:  models.ToPointer(models.GatewayTrafficShaping{
                Enabled:          models.ToPointer(false),
            }),
            Usage:           models.GatewayPortUsageEnum("wan"),
            WanType:         models.ToPointer(models.GatewayPortWanTypeEnum("broadband")),
        },
        "ge-0/0/6-7": models.GatewayPortConfig{
            Networks:        []string{
                "Corp-lan",
                "Corp-Mgmt",
                "Corp-Core",
            },
            Usage:           models.GatewayPortUsageEnum("lan"),
        },
    },
    ServicePolicies:       []models.ServicePolicy{
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-Internal"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "drive",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("ITParis-internet"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-core"),
            PathPreference:  models.ToPointer("core"),
            Services:        []string{
                "internal_dns",
                "internal_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-mxe-tt-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mgmt-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "mxedge-updates",
                "radsec",
                "icmp",
                "internet_dns",
                "internet_ntp",
            },
            Tenants:         []string{
                "Corp-Mgmt",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("mxe-data-0-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-mgt"),
            PathPreference:  models.ToPointer("mgmt"),
            Services:        []string{
                "mgmt",
            },
            Tenants:         []string{
                "domain.Corp-Core",
                "lan.Corp-Core",
                "servers.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-edge-in"),
            PathPreference:  models.ToPointer("mxe-in"),
            Services:        []string{
                "internal_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("core-to-iot"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "iot",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "servers-hassio.Corp-Core",
                "servers-kubes.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("tanker-to-cctv"),
            PathPreference:  models.ToPointer("iot"),
            Services:        []string{
                "rtsp",
            },
            Tenants:         []string{
                "servers-tanker.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("allow")),
            Idp:             models.ToPointer(models.IdpConfig{
                Enabled:      models.ToPointer(false),
            }),
            Name:            models.ToPointer("core-to-untrust"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "internet_any",
            },
            Tenants:         []string{
                "lan.Corp-Core",
                "domain.Corp-Core",
                "servers.Corp-Core",
            },
        },
        models.ServicePolicy{
            Action:          models.ToPointer(models.AllowDenyEnum("deny")),
            Idp:             models.ToPointer(models.IdpConfig{
                AlertOnly:    models.ToPointer(true),
                Enabled:      models.ToPointer(true),
                Profile:      models.ToPointer("standard"),
            }),
            Name:            models.ToPointer("iot-upgrade-cctv"),
            PathPreference:  models.ToPointer("untrust"),
            Services:        []string{
                "motioneye",
                "nodejs",
                "raspbian",
            },
            Tenants:         []string{
                "printer",
            },
        },
    },
    Type:                  models.ToPointer(models.GatewayTemplateTypeEnum("standalone")),
}

apiResponse, err := orgsGatewayTemplates.UpdateOrgGatewayTemplate(ctx, orgId, gatewaytemplateId, &body)
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

