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
        AdditionalProperties: map[string]models.DhcpdConfigProperty{
            "Corp-Mgmt": models.DhcpdConfigProperty{
                DnsServers:           []string{
                    "8.8.8.8",
                },
                DnsSuffix:            []string{
                    "stag.one",
                },
                Gateway:              models.ToPointer("10.3.172.9"),
                IpEnd:                models.ToPointer("10.3.172.99"),
                IpStart:              models.ToPointer("10.3.172.50"),
                Type:                 models.ToPointer(models.DhcpdConfigTypeEnum_LOCAL),
            },
            "Corp-lan": models.DhcpdConfigProperty{
                DnsServers:           []string{
                    "8.8.8.8",
                },
                DnsSuffix:            []string{
                    "stag.one",
                },
                Gateway:              models.ToPointer("10.3.171.9"),
                IpEnd:                models.ToPointer("10.3.171.99"),
                IpStart:              models.ToPointer("10.3.171.50"),
                Type:                 models.ToPointer(models.DhcpdConfigTypeEnum_LOCAL),
            },
        },
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
            Via:                  models.ToPointer("10.3.100.10"),
        },
    },
    IpConfigs:             map[string]models.GatewayIpConfigProperty{
        "Corp-Core": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.100.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
        "Corp-Mgmt": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.172.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
        "Corp-lan": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.171.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
    },
    Name:                  "ITParis",
    NtpOverride:           models.ToPointer(true),
    NtpServers:            []string{
        "10.3.51.222",
    },
    PathPreferences:       map[string]models.GatewayPathPreferences{
        "core": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-Core",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "lab": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-lan",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "mgmt": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-Mgmt",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "untrust": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Name:                 models.ToPointer("wan"),
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_WAN),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
    },
    PortConfig:            map[string]models.GatewayPortConfig{
        "ge-0/0/0": models.GatewayPortConfig{
            Aggregated:           models.ToPointer(false),
            IpConfig:             models.ToPointer(models.GatewayPortConfigIpConfig{
                Gateway:              models.ToPointer("192.168.1.1"),
                Ip:                   models.ToPointer("192.168.1.9"),
                Netmask:              models.ToPointer("/24"),
                Type:                 models.ToPointer(models.GatewayWanTypeEnum_STATIC),
            }),
            Name:                 models.ToPointer("wan"),
            Redundant:            models.ToPointer(false),
            TrafficShaping:       models.ToPointer(models.GatewayTrafficShaping{
                Enabled:              models.ToPointer(false),
            }),
            Usage:                models.GatewayPortUsageEnum_WAN,
            WanType:              models.ToPointer(models.GatewayPortWanTypeEnum_BROADBAND),
        },
        "ge-0/0/6-7": models.GatewayPortConfig{
            AeDisableLacp:        models.ToPointer(false),
            AeIdx:                models.NewOptional(models.ToPointer("0")),
            AeLacpForceUp:        models.ToPointer(true),
            Aggregated:           models.ToPointer(true),
            Networks:             []string{
                "Corp-lan",
                "Corp-Mgmt",
                "Corp-Core",
            },
            Usage:                models.GatewayPortUsageEnum_LAN,
        },
    },
    ServicePolicies:       []models.ServicePolicy{
        models.ServicePolicy{
            Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
            Idp:                  models.ToPointer(models.IdpConfig{
                Enabled:              models.ToPointer(false),
            }),
            Name:                 models.ToPointer("ITParis-Internal"),
            PathPreference:       models.ToPointer("core"),
            Services:             []string{
                "internal_dns",
                "drive",
            },
            Tenants:              []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:               models.ToPointer(models.AllowDenyEnum_DENY),
            Idp:                  models.ToPointer(models.IdpConfig{
                Enabled:              models.ToPointer(false),
            }),
            Name:                 models.ToPointer("ITParis-internet"),
            PathPreference:       models.ToPointer("untrust"),
            Services:             []string{
                "internet_any",
            },
            Tenants:              []string{
                "ITParis",
            },
        },
    },
    Type:                  models.ToPointer(models.GatewayTemplateTypeEnum_STANDALONE),
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
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

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
        AdditionalProperties: map[string]models.DhcpdConfigProperty{
            "Corp-Mgmt": models.DhcpdConfigProperty{
                DnsServers:           []string{
                    "8.8.8.8",
                },
                DnsSuffix:            []string{
                    "stag.one",
                },
                Gateway:              models.ToPointer("10.3.172.9"),
                IpEnd:                models.ToPointer("10.3.172.99"),
                IpStart:              models.ToPointer("10.3.172.50"),
                Type:                 models.ToPointer(models.DhcpdConfigTypeEnum_LOCAL),
            },
            "Corp-lan": models.DhcpdConfigProperty{
                DnsServers:           []string{
                    "8.8.8.8",
                },
                DnsSuffix:            []string{
                    "stag.one",
                },
                Gateway:              models.ToPointer("10.3.171.9"),
                IpEnd:                models.ToPointer("10.3.171.99"),
                IpStart:              models.ToPointer("10.3.171.50"),
                Type:                 models.ToPointer(models.DhcpdConfigTypeEnum_LOCAL),
            },
        },
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
            Via:                  models.ToPointer("10.3.100.10"),
        },
    },
    IpConfigs:             map[string]models.GatewayIpConfigProperty{
        "Corp-Core": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.100.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
        "Corp-Mgmt": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.172.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
        "Corp-lan": models.GatewayIpConfigProperty{
            Ip:                   models.ToPointer("10.3.171.9"),
            Netmask:              models.ToPointer("/24"),
            Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        },
    },
    Name:                  "ITParis",
    NtpOverride:           models.ToPointer(true),
    NtpServers:            []string{
        "10.3.51.222",
    },
    PathPreferences:       map[string]models.GatewayPathPreferences{
        "core": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-Core",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "lab": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-lan",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "mgmt": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Networks:             []string{
                        "Corp-Mgmt",
                    },
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_LOCAL),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
        "untrust": models.GatewayPathPreferences{
            Paths:                []models.GatewayPathPreferencesPath{
                models.GatewayPathPreferencesPath{
                    Name:                 models.ToPointer("wan"),
                    Type:                 models.ToPointer(models.GatewayPathTypeEnum_WAN),
                },
            },
            Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
        },
    },
    PortConfig:            map[string]models.GatewayPortConfig{
        "ge-0/0/0": models.GatewayPortConfig{
            Aggregated:           models.ToPointer(false),
            IpConfig:             models.ToPointer(models.GatewayPortConfigIpConfig{
                Gateway:              models.ToPointer("192.168.1.1"),
                Ip:                   models.ToPointer("192.168.1.9"),
                Netmask:              models.ToPointer("/24"),
                Type:                 models.ToPointer(models.GatewayWanTypeEnum_STATIC),
            }),
            Name:                 models.ToPointer("wan"),
            Redundant:            models.ToPointer(false),
            TrafficShaping:       models.ToPointer(models.GatewayTrafficShaping{
                Enabled:              models.ToPointer(false),
            }),
            Usage:                models.GatewayPortUsageEnum_WAN,
            WanType:              models.ToPointer(models.GatewayPortWanTypeEnum_BROADBAND),
        },
        "ge-0/0/6-7": models.GatewayPortConfig{
            AeDisableLacp:        models.ToPointer(false),
            AeIdx:                models.NewOptional(models.ToPointer("0")),
            AeLacpForceUp:        models.ToPointer(true),
            Aggregated:           models.ToPointer(true),
            Networks:             []string{
                "Corp-lan",
                "Corp-Mgmt",
                "Corp-Core",
            },
            Usage:                models.GatewayPortUsageEnum_LAN,
        },
    },
    ServicePolicies:       []models.ServicePolicy{
        models.ServicePolicy{
            Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
            Idp:                  models.ToPointer(models.IdpConfig{
                Enabled:              models.ToPointer(false),
            }),
            Name:                 models.ToPointer("ITParis-Internal"),
            PathPreference:       models.ToPointer("core"),
            Services:             []string{
                "internal_dns",
                "drive",
            },
            Tenants:              []string{
                "ITParis",
            },
        },
        models.ServicePolicy{
            Action:               models.ToPointer(models.AllowDenyEnum_DENY),
            Idp:                  models.ToPointer(models.IdpConfig{
                Enabled:              models.ToPointer(false),
            }),
            Name:                 models.ToPointer("ITParis-internet"),
            PathPreference:       models.ToPointer("untrust"),
            Services:             []string{
                "internet_any",
            },
            Tenants:              []string{
                "ITParis",
            },
        },
    },
    Type:                  models.ToPointer(models.GatewayTemplateTypeEnum_STANDALONE),
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

