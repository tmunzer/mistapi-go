# Orgs NAC Rules

```go
orgsNACRules := client.OrgsNACRules()
```

## Class Name

`OrgsNACRules`

## Methods

* [Create Org Nac Rule](../../doc/controllers/orgs-nac-rules.md#create-org-nac-rule)
* [Delete Org Nac Rule](../../doc/controllers/orgs-nac-rules.md#delete-org-nac-rule)
* [Get Org Nac Rule](../../doc/controllers/orgs-nac-rules.md#get-org-nac-rule)
* [List Org Nac Rules](../../doc/controllers/orgs-nac-rules.md#list-org-nac-rules)
* [Update Org Nac Rule](../../doc/controllers/orgs-nac-rules.md#update-org-nac-rule)


# Create Org Nac Rule

create Org NAC Rule

```go
CreateOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacRule) (
    models.ApiResponse[models.NacRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacRule`](../../doc/models/nac-rule.md) | Body, Optional | - |

## Response Type

[`models.NacRule`](../../doc/models/nac-rule.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacRule{
    Action:       models.NacRuleActionEnum("allow"),
    ApplyTags:    []string{
        "c049dfcd-0c73-5014-1c64-062e9903f1e5\"",
    },
    Matching:     models.ToPointer(models.NacRuleMatching{
        AuthType:     models.ToPointer(models.NacRuleMatchingAuthTypeEnum("eap-tls")),
        Nactags:      []string{
            "041d5d36-716c-4cfb-4988-3857c6aa14a2",
            "a809a97f-d599-f812-eb8c-c3f84aabf6ba",
        },
        PortTypes:    []models.NacRuleMatchingPortTypeEnum{
            models.NacRuleMatchingPortTypeEnum("wired"),
        },
        SiteIds:      []uuid.UUID{
            uuid.MustParse("bb19fc3e-4124-4b57-80d9-c3f6edce47c4"),
            uuid.MustParse("bb19fc3e-6564-4b57-80d9-c3f6edce47c1"),
        },
        SitegroupIds: []uuid.UUID{
            uuid.MustParse("bb19fc3e-4124-4b57-80d9-c3f6edce47c4"),
            uuid.MustParse("bb19fc3e-6564-4b57-80d9-c3f6edce47c1"),
        },
    }),
    Name:         "name1",
    NotMatching:  models.ToPointer(models.NacRuleMatching{
    }),
    Order:        models.ToPointer(1),
}

apiResponse, err := orgsNACRules.CreateOrgNacRule(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Org Nac Rule

Delete Org NAC Rule

```go
DeleteOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nacruleId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nacruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNACRules.DeleteOrgNacRule(ctx, orgId, nacruleId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Nac Rule

Get Org NAC Rule

```go
GetOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID) (
    models.ApiResponse[models.NacRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nacruleId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.NacRule`](../../doc/models/nac-rule.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nacruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNACRules.GetOrgNacRule(ctx, orgId, nacruleId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Nac Rules

Get List of Org NAC Rules

```go
ListOrgNacRules(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.NacRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.NacRule`](../../doc/models/nac-rule.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsNACRules.ListOrgNacRules(ctx, orgId, &page, &limit)
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
    "action": "allow",
    "apply_tags": [
      "string"
    ],
    "created_time": 0,
    "id": "455f6eca-6276-4993-bfeb-53cbbbba6208",
    "matching": {
      "auth_type": "eap-tls",
      "nactags": [
        "string"
      ],
      "port_types": [
        "wireless"
      ],
      "site_ids": [
        "454f6eca-6276-4993-bfeb-53cbbbba6308"
      ],
      "sitegroup_ids": [
        "453f6eca-6276-4993-bfeb-53cbbbba6408"
      ]
    },
    "modified_time": 0,
    "name": "string",
    "not_matching": {
      "auth_type": "eap-tls",
      "nactags": [
        "string"
      ],
      "port_types": [
        "wireless"
      ],
      "site_ids": [
        "452f6eca-6276-4993-bfeb-53cbbbba6508"
      ],
      "sitegroup_ids": [
        "451f6eca-6276-4993-bfeb-53cbbbba6608"
      ]
    },
    "order": 1,
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Nac Rule

Update Org NAC Rule

```go
UpdateOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID,
    body *models.NacRule) (
    models.ApiResponse[models.NacRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nacruleId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacRule`](../../doc/models/nac-rule.md) | Body, Optional | - |

## Response Type

[`models.NacRule`](../../doc/models/nac-rule.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nacruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacRule{
    Action:       models.NacRuleActionEnum("allow"),
    ApplyTags:    []string{
        "c049dfcd-0c73-5014-1c64-062e9903f1e5",
    },
    Enabled:      models.ToPointer(true),
    Name:         "name6",
    Order:        models.ToPointer(1),
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
}

apiResponse, err := orgsNACRules.UpdateOrgNacRule(ctx, orgId, nacruleId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
