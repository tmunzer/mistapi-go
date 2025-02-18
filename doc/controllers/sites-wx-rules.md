# Sites Wx Rules

```go
sitesWxRules := client.SitesWxRules()
```

## Class Name

`SitesWxRules`

## Methods

* [Create Site Wx Rule](../../doc/controllers/sites-wx-rules.md#create-site-wx-rule)
* [Delete Site Wx Rule](../../doc/controllers/sites-wx-rules.md#delete-site-wx-rule)
* [Get Site Wx Rule](../../doc/controllers/sites-wx-rules.md#get-site-wx-rule)
* [List Site Wx Rules](../../doc/controllers/sites-wx-rules.md#list-site-wx-rules)
* [List Site Wx Rules Derived](../../doc/controllers/sites-wx-rules.md#list-site-wx-rules-derived)
* [Update Site Wx Rule](../../doc/controllers/sites-wx-rules.md#update-site-wx-rule)


# Create Site Wx Rule

Create Site WxLan Rule

```go
CreateSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.WxlanRule) (
    models.ApiResponse[models.WxlanRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanRule`](../../doc/models/wxlan-rule.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanRule`](../../doc/models/wxlan-rule.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanRule{
    Action:               models.ToPointer(models.WxlanRuleActionEnum_ALLOW),
    ApplyTags:            []string{
        "c049dfcd-0c73-5014-1c64-062e9903f1e5",
    },
    BlockedApps:          []string{
        "mist",
        "all-videos",
    },
    DstAllowWxtags:       []string{
        "fff34466-eec0-3756-6765-381c728a6037",
        "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
    },
    DstDenyWxtags:        []string{
        "aaa34466-eec0-3756-6765-381c728a6037",
        "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
    },
    Enabled:              models.ToPointer(true),
    Order:                1,
    SrcWxtags:            []string{
        "8bfc2490-d726-3587-038d-cb2e71bd2330",
        "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9",
    },
}

apiResponse, err := sitesWxRules.CreateSiteWxRule(ctx, siteId, &body)
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
  "action": "allow",
  "apply_tags": [
    "c049dfcd-0c73-5014-1c64-062e9903f1e5"
  ],
  "blocked_apps": [
    "mist",
    "all-videos"
  ],
  "created_time": 0,
  "dst_allow_wxtags": [
    "fff34466-eec0-3756-6765-381c728a6037",
    "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_deny_wxtags": [
    "aaa34466-eec0-3756-6765-381c728a6037",
    "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "enabled": true,
  "for_site": true,
  "id": "497f6eca-6276-4993-9feb-53cbbbba6f08",
  "modified_time": 0,
  "order": 1,
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "src_wxtags": [
    "8bfc2490-d726-3587-038d-cb2e71bd2330",
    "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
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


# Delete Site Wx Rule

Delete Site WxLan Rule

```go
DeleteSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxruleId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesWxRules.DeleteSiteWxRule(ctx, siteId, wxruleId)
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


# Get Site Wx Rule

Get Site WxLan Rule Details

```go
GetSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID) (
    models.ApiResponse[models.WxlanRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxruleId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.WxlanRule`](../../doc/models/wxlan-rule.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxRules.GetSiteWxRule(ctx, siteId, wxruleId)
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
  "action": "allow",
  "apply_tags": [
    "c049dfcd-0c73-5014-1c64-062e9903f1e5"
  ],
  "blocked_apps": [
    "mist",
    "all-videos"
  ],
  "created_time": 0,
  "dst_allow_wxtags": [
    "fff34466-eec0-3756-6765-381c728a6037",
    "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_deny_wxtags": [
    "aaa34466-eec0-3756-6765-381c728a6037",
    "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "enabled": true,
  "for_site": true,
  "id": "497f6eca-6276-4993-9feb-53cbbbba6f08",
  "modified_time": 0,
  "order": 1,
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "src_wxtags": [
    "8bfc2490-d726-3587-038d-cb2e71bd2330",
    "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
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


# List Site Wx Rules

Get List of Site WxLan Rules

```go
ListSiteWxRules(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.WxlanRule`](../../doc/models/wxlan-rule.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesWxRules.ListSiteWxRules(ctx, siteId, &limit, &page)
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
      "c049dfcd-0c73-5014-1c64-062e9903f1e5"
    ],
    "blocked_apps": [
      "mist",
      "all-videos"
    ],
    "created_time": 0,
    "dst_allow_wxtags": [
      "fff34466-eec0-3756-6765-381c728a6037",
      "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "dst_deny_wxtags": [
      "aaa34466-eec0-3756-6765-381c728a6037",
      "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "enabled": true,
    "for_site": true,
    "id": "497f6eca-6276-4993-bfeb-53ebbbba6f08",
    "modified_time": 0,
    "order": 1,
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "src_wxtags": [
      "8bfc2490-d726-3587-038d-cb2e71bd2330",
      "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
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


# List Site Wx Rules Derived

Get Site WxLan Rule Derived

```go
ListSiteWxRulesDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.WxlanRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.WxlanRule`](../../doc/models/wxlan-rule.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxRules.ListSiteWxRulesDerived(ctx, siteId)
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
      "c049dfcd-0c73-5014-1c64-062e9903f1e5"
    ],
    "blocked_apps": [
      "mist",
      "all-videos"
    ],
    "created_time": 0,
    "dst_allow_wxtags": [
      "fff34466-eec0-3756-6765-381c728a6037",
      "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "dst_deny_wxtags": [
      "aaa34466-eec0-3756-6765-381c728a6037",
      "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "enabled": true,
    "for_site": true,
    "id": "497f6eca-6276-4993-bfeb-53ebbbba6f08",
    "modified_time": 0,
    "order": 1,
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "src_wxtags": [
      "8bfc2490-d726-3587-038d-cb2e71bd2330",
      "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
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


# Update Site Wx Rule

Update Site WxLan Rule

```go
UpdateSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID,
    body *models.WxlanRule) (
    models.ApiResponse[models.WxlanRule],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxruleId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanRule`](../../doc/models/wxlan-rule.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanRule`](../../doc/models/wxlan-rule.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxruleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanRule{
    Action:               models.ToPointer(models.WxlanRuleActionEnum_ALLOW),
    ApplyTags:            []string{
        "c049dfcd-0c73-5014-1c64-062e9903f1e5",
    },
    BlockedApps:          []string{
        "mist",
        "all-videos",
    },
    DstAllowWxtags:       []string{
        "fff34466-eec0-3756-6765-381c728a6037",
        "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
    },
    DstDenyWxtags:        []string{
        "aaa34466-eec0-3756-6765-381c728a6037",
        "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
    },
    Enabled:              models.ToPointer(true),
    Order:                1,
    SrcWxtags:            []string{
        "8bfc2490-d726-3587-038d-cb2e71bd2330",
        "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9",
    },
}

apiResponse, err := sitesWxRules.UpdateSiteWxRule(ctx, siteId, wxruleId, &body)
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
  "action": "allow",
  "apply_tags": [
    "c049dfcd-0c73-5014-1c64-062e9903f1e5"
  ],
  "blocked_apps": [
    "mist",
    "all-videos"
  ],
  "created_time": 0,
  "dst_allow_wxtags": [
    "fff34466-eec0-3756-6765-381c728a6037",
    "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_deny_wxtags": [
    "aaa34466-eec0-3756-6765-381c728a6037",
    "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "enabled": true,
  "for_site": true,
  "id": "497f6eca-6276-4993-9feb-53cbbbba6f08",
  "modified_time": 0,
  "order": 1,
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "src_wxtags": [
    "8bfc2490-d726-3587-038d-cb2e71bd2330",
    "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
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

