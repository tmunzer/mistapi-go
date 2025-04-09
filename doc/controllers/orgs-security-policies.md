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
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

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

