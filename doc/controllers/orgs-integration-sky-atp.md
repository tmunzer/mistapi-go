# Orgs Integration Sky ATP

```go
orgsIntegrationSkyATP := client.OrgsIntegrationSkyATP()
```

## Class Name

`OrgsIntegrationSkyATP`

## Methods

* [Delete Org Sky Atp Integration](../../doc/controllers/orgs-integration-sky-atp.md#delete-org-sky-atp-integration)
* [Get Org Sky Atp Integration](../../doc/controllers/orgs-integration-sky-atp.md#get-org-sky-atp-integration)
* [Setup Org Atp Integration](../../doc/controllers/orgs-integration-sky-atp.md#setup-org-atp-integration)
* [Udpate Org Atp Allowed List](../../doc/controllers/orgs-integration-sky-atp.md#udpate-org-atp-allowed-list)
* [Udpate Org Atp Blocked List](../../doc/controllers/orgs-integration-sky-atp.md#udpate-org-atp-blocked-list)
* [Udpate Org Atp Integration](../../doc/controllers/orgs-integration-sky-atp.md#udpate-org-atp-integration)


# Delete Org Sky Atp Integration

Delete SkyATP Integration

```go
DeleteOrgSkyAtpIntegration(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsIntegrationSkyATP.DeleteOrgSkyAtpIntegration(ctx, orgId)
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


# Get Org Sky Atp Integration

Get Org SkyATP Integration

```go
GetOrgSkyAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountSkyatpData],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountSkyatpData](../../doc/models/account-skyatp-data.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsIntegrationSkyATP.GetOrgSkyAtpIntegration(ctx, orgId)
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
  "Example": {
    "value": {
      "secintel": {
        "third_party_threat_feeds": [
          "block_list"
        ]
      },
      "secintel_allowlist_url": "https://papi.s3.amazonaws.com/secintel_allowlist/xxx...",
      "secintel_blocklist_url": "https://papi.s3.amazonaws.com/secintel_blocklist/xxx..."
    }
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


# Setup Org Atp Integration

1. Login to the Sky ATP realm through the Mist UI by providing the realm, username and password.
2. Sky ATP API is invoked which creates the realm using above details.
3. Sky ATP by default will provide functionality for Security-Intelligence and Advanced Anti Malware.
4. Security Intelligence will provide configuration for CC, DNS Feeds, Infected Host, Blocklists and Allowlists.

```go
SetupOrgAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountSkyatpConfig) (
    models.ApiResponse[models.AccountSkyatpData],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountSkyatpConfig`](../../doc/models/account-skyatp-config.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountSkyatpData](../../doc/models/account-skyatp-data.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountSkyatpConfig{
    Password:             "foryoureyesonly",
    Realm:                "mist-team",
    Username:             "john@abc.com",
}

apiResponse, err := orgsIntegrationSkyATP.SetupOrgAtpIntegration(ctx, orgId, &body)
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
  "Example": {
    "value": {
      "secintel": {
        "third_party_threat_feeds": [
          "block_list"
        ]
      },
      "secintel_allowlist_url": "https://papi.s3.amazonaws.com/secintel_allowlist/xxx...",
      "secintel_blocklist_url": "https://papi.s3.amazonaws.com/secintel_blocklist/xxx..."
    }
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


# Udpate Org Atp Allowed List

Update Sky ATP Allowed List

```go
UdpateOrgAtpAllowedList(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SkyatpList) (
    models.ApiResponse[models.SkyatpList],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SkyatpList`](../../doc/models/skyatp-list.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SkyatpList](../../doc/models/skyatp-list.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SkyatpList{
    Domains:              []models.SkyatpListDomain{
        models.SkyatpListDomain{
            Comment:              models.ToPointer("restricted"),
            Value:                "unsafe.xxx",
        },
    },
    AdditionalProperties: map[string]interface{}{
        "ips": interface{}("System.Collections.Generic.Dictionary`2[System.String,System.Object]"),
    },
}

apiResponse, err := orgsIntegrationSkyATP.UdpateOrgAtpAllowedList(ctx, orgId, &body)
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
  "domains": [
    {
      "comment": "restricted",
      "value": "unsafe.xxx"
    }
  ],
  "ips": [
    {
      "comment": "nas",
      "value": "10.1.3.5"
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


# Udpate Org Atp Blocked List

Update Sky ATP Blocked List

```go
UdpateOrgAtpBlockedList(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SkyatpList) (
    models.ApiResponse[models.SkyatpList],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SkyatpList`](../../doc/models/skyatp-list.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SkyatpList](../../doc/models/skyatp-list.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SkyatpList{
    Domains:              []models.SkyatpListDomain{
        models.SkyatpListDomain{
            Comment:              models.ToPointer("restricted"),
            Value:                "unsafe.xxx",
        },
    },
    AdditionalProperties: map[string]interface{}{
        "ips": interface{}("System.Collections.Generic.Dictionary`2[System.String,System.Object]"),
    },
}

apiResponse, err := orgsIntegrationSkyATP.UdpateOrgAtpBlockedList(ctx, orgId, &body)
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
  "domains": [
    {
      "comment": "restricted",
      "value": "unsafe.xxx"
    }
  ],
  "ips": [
    {
      "comment": "nas",
      "value": "10.1.3.5"
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


# Udpate Org Atp Integration

Update Sky ATP config

```go
UdpateOrgAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountSkyatpData) (
    models.ApiResponse[models.AccountSkyatpInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountSkyatpData`](../../doc/models/account-skyatp-data.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountSkyatpInfo](../../doc/models/account-skyatp-info.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountSkyatpData{
    Secintel:             models.ToPointer(models.AccountSkyatpDataSecintel{
        ThirdPartyThreatFeeds: []string{
            "block_list",
        },
    }),
}

apiResponse, err := orgsIntegrationSkyATP.UdpateOrgAtpIntegration(ctx, orgId, &body)
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

