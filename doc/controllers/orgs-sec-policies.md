# Orgs Sec Policies

```go
orgsSecPolicies := client.OrgsSecPolicies()
```

## Class Name

`OrgsSecPolicies`

## Methods

* [Create Org Sec Policies](../../doc/controllers/orgs-sec-policies.md#create-org-sec-policies)
* [Delete Org Sec Policy](../../doc/controllers/orgs-sec-policies.md#delete-org-sec-policy)
* [Get Org Sec Policy](../../doc/controllers/orgs-sec-policies.md#get-org-sec-policy)
* [List Org Sec Policies](../../doc/controllers/orgs-sec-policies.md#list-org-sec-policies)
* [Update Org Sec Policies](../../doc/controllers/orgs-sec-policies.md#update-org-sec-policies)


# Create Org Sec Policies

Create Org Security Policy

```go
CreateOrgSecPolicies(
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

[`models.Secpolicy`](../../doc/models/secpolicy.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Secpolicy{
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := orgsSecPolicies.CreateOrgSecPolicies(ctx, orgId, &body)
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

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSecPolicies.DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
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

[`models.Secpolicy`](../../doc/models/secpolicy.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSecPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
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
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Secpolicy],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Secpolicy`](../../doc/models/secpolicy.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsSecPolicies.ListOrgSecPolicies(ctx, orgId, &page, &limit)
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


# Update Org Sec Policies

Update Org Security Policy

```go
UpdateOrgSecPolicies(
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

[`models.Secpolicy`](../../doc/models/secpolicy.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secpolicyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Secpolicy{
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := orgsSecPolicies.UpdateOrgSecPolicies(ctx, orgId, secpolicyId, &body)
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

