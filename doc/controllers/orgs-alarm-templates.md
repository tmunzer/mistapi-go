# Orgs Alarm Templates

```go
orgsAlarmTemplates := client.OrgsAlarmTemplates()
```

## Class Name

`OrgsAlarmTemplates`

## Methods

* [Create Org Alarm Template](../../doc/controllers/orgs-alarm-templates.md#create-org-alarm-template)
* [Delete Org Alarm Template](../../doc/controllers/orgs-alarm-templates.md#delete-org-alarm-template)
* [Get Org Alarm Template](../../doc/controllers/orgs-alarm-templates.md#get-org-alarm-template)
* [List Org Alarm Templates](../../doc/controllers/orgs-alarm-templates.md#list-org-alarm-templates)
* [List Org Suppressed Alarms](../../doc/controllers/orgs-alarm-templates.md#list-org-suppressed-alarms)
* [Suppress Org Alarm](../../doc/controllers/orgs-alarm-templates.md#suppress-org-alarm)
* [Unsuppress Org Suppressed Alarms](../../doc/controllers/orgs-alarm-templates.md#unsuppress-org-suppressed-alarms)
* [Update Org Alarm Template](../../doc/controllers/orgs-alarm-templates.md#update-org-alarm-template)


# Create Org Alarm Template

Available rules can be found in Orgs>Consts>getAlarmDefs

The delivery dict is only required if different from the template delivery settings.

```go
CreateOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AlarmTemplate) (
    models.ApiResponse[models.AlarmTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AlarmTemplate`](../../doc/models/alarm-template.md) | Body, Optional | Request Body |

## Response Type

[`models.AlarmTemplate`](../../doc/models/alarm-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AlarmTemplate{
    CreatedTime:  models.ToPointer(float64(1711001253)),
    Delivery:     models.Delivery{
        Enabled:          true,
        ToOrgAdmins:      models.ToPointer(true),
        ToSiteAdmins:     models.ToPointer(false),
    },
    Id:           models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
    ModifiedTime: models.ToPointer(float64(1711038102)),
    Name:         models.ToPointer("default"),
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Rules:        map[string]models.AlarmTemplateRule{
        "ap_offline": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "bad_cable": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
    },
}

apiResponse, err := orgsAlarmTemplates.CreateOrgAlarmTemplate(ctx, orgId, &body)
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


# Delete Org Alarm Template

Delete Org Alarm Template

```go
DeleteOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `alarmtemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

alarmtemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAlarmTemplates.DeleteOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
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


# Get Org Alarm Template

Get Org Alarm Template Details

```go
GetOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID) (
    models.ApiResponse[models.AlarmTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `alarmtemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.AlarmTemplate`](../../doc/models/alarm-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

alarmtemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAlarmTemplates.GetOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
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


# List Org Alarm Templates

Get List of Org Alarm Templates

```go
ListOrgAlarmTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.AlarmTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.AlarmTemplate`](../../doc/models/alarm-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsAlarmTemplates.ListOrgAlarmTemplates(ctx, orgId, &limit, &page)
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


# List Org Suppressed Alarms

Get List of Org Alarms Currently Suppressed

```go
ListOrgSuppressedAlarms(
    ctx context.Context,
    orgId uuid.UUID,
    scope *models.SuppressedAlarmScopeEnum) (
    models.ApiResponse[models.ResponseOrgSuppressAlarm],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Query, Optional | - |

## Response Type

[`models.ResponseOrgSuppressAlarm`](../../doc/models/response-org-suppress-alarm.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SuppressedAlarmScopeEnum("site")

apiResponse, err := orgsAlarmTemplates.ListOrgSuppressedAlarms(ctx, orgId, &scope)
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
  "results": [
    {
      "duration": 48,
      "site_id": "581328b6-e382-f54e-c9dc-9c998d183a34"
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


# Suppress Org Alarm

In certain situations, for example, scheduled maintenance, you may want to suspend alarms to be triggered against Sites for a period of time.

```go
SuppressOrgAlarm(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SuppressedAlarm) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SuppressedAlarm`](../../doc/models/suppressed-alarm.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SuppressedAlarm{
    Duration:      models.ToPointer(float64(3600)),
    ScheduledTime: models.ToPointer(1678232980),
    Scope:         models.ToPointer(models.SuppressedAlarmScopeEnum("org")),
}

resp, err := orgsAlarmTemplates.SuppressOrgAlarm(ctx, orgId, &body)
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


# Unsuppress Org Suppressed Alarms

Un-Suppress Suppressed Alarms

```go
UnsuppressOrgSuppressedAlarms(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAlarmTemplates.UnsuppressOrgSuppressedAlarms(ctx, orgId)
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


# Update Org Alarm Template

Update Org Alarm Template

```go
UpdateOrgAlarmTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    alarmtemplateId uuid.UUID,
    body *models.AlarmTemplate) (
    models.ApiResponse[models.AlarmTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `alarmtemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AlarmTemplate`](../../doc/models/alarm-template.md) | Body, Optional | Request Body |

## Response Type

[`models.AlarmTemplate`](../../doc/models/alarm-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

alarmtemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AlarmTemplate{
    Delivery:     models.Delivery{
        AdditionalEmails: []string{
            "string",
        },
        Enabled:          true,
        ToOrgAdmins:      models.ToPointer(true),
        ToSiteAdmins:     models.ToPointer(true),
    },
    Name:         models.ToPointer("string"),
    Rules:        map[string]models.AlarmTemplateRule{
        "adhoc_network": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "air_magnet_scan": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "ap_offline": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "bad_cable": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "beacon_flood": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "bssid_spoofing": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "device_down": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "device_restarted": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "dhcp_failure": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "disassociation_flood": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "dot1x_failure": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "eap_dictionary_attack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "eap_failure_injection": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "eap_handshake_flood": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "eap_spoofed_success": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "eapol_logoff_attack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "essid_jack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "excessive_client": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "excessive_eapol_start": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "gateway_down": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "gw_bad_cable": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "gw_negotiation_mismatch": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "honeypot_ssid": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "krack_attack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "missing_vlan": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "monkey_jack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "negotiation_mismatch": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "non_compliant": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "out_of_sequence": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                Enabled:          false,
            }),
            Enabled:  models.ToPointer(true),
        },
        "psk_failure": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "repeated_auth_failures": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "rogue_ap": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "rogue_client": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "secpolicy_violation": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "ssid_injection": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "switch_down": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "tkip_icv_attack": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "vendor_ie_missing": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "watched_station": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
        "zero_ssid_association": models.AlarmTemplateRule{
            Delivery: models.ToPointer(models.Delivery{
                AdditionalEmails: []string{
                    "string",
                },
                Enabled:          true,
                ToOrgAdmins:      models.ToPointer(true),
                ToSiteAdmins:     models.ToPointer(true),
            }),
            Enabled:  models.ToPointer(true),
        },
    },
}

apiResponse, err := orgsAlarmTemplates.UpdateOrgAlarmTemplate(ctx, orgId, alarmtemplateId, &body)
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

