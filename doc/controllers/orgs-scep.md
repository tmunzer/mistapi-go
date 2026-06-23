# Orgs SCEP

```go
orgsSCEP := client.OrgsSCEP()
```

## Class Name

`OrgsSCEP`

## Methods

* [Disable Org Mist Scep](../../doc/controllers/orgs-scep.md#disable-org-mist-scep)
* [Get Org Mist Scep](../../doc/controllers/orgs-scep.md#get-org-mist-scep)
* [List Org Issued Client Certificates](../../doc/controllers/orgs-scep.md#list-org-issued-client-certificates)
* [Revoke Org Issued Client Certificates](../../doc/controllers/orgs-scep.md#revoke-org-issued-client-certificates)
* [Update Org Mist Scep](../../doc/controllers/orgs-scep.md#update-org-mist-scep)


# Disable Org Mist Scep

Disable Mist SCEP for the organization and return the updated read-only SCEP settings.

```go
DisableOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSettingScepResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSettingScepResponse](../../doc/models/org-setting-scep-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSCEP.DisableOrgMistScep(ctx, orgId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "cert_providers": [
    "jamf",
    "intune",
    "byod"
  ],
  "enabled": false,
  "intune_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_access_token": "1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=",
  "jamf_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_webhook_url": "https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Mist Scep

Return Mist SCEP settings for the organization, including enabled and suspended status, configured certificate providers, and generated enrollment or webhook URLs.

```go
GetOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSettingScepResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSettingScepResponse](../../doc/models/org-setting-scep-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSCEP.GetOrgMistScep(ctx, orgId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "cert_providers": [
    "jamf",
    "intune",
    "byod"
  ],
  "enabled": false,
  "intune_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_access_token": "1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=",
  "jamf_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_webhook_url": "https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Issued Client Certificates

List Mist SCEP client certificates issued for the organization. Results can be filtered by common name, certificate provider, serial number, device ID, or time range.

```go
ListOrgIssuedClientCertificates(
    ctx context.Context,
    orgId uuid.UUID,
    commonName *string,
    certProvider *string,
    serialNumber *string,
    deviceId *string,
    expireTime *int,
    createdTime *int,
    limit *int,
    page *int) (
    models.ApiResponse[models.IssuedClientCertificatesResults],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `commonName` | `*string` | Query, Optional | Filter by certificate common name (e.g. user UPN or device identifier) |
| `certProvider` | `*string` | Query, Optional | Filter by MDM or certificate provider that issued the certificate. Accepts multiple comma-separated values. |
| `serialNumber` | `*string` | Query, Optional | Filter by certificate serial number. Accepts multiple comma-separated values. |
| `deviceId` | `*string` | Query, Optional | Filter by device identifier associated with the certificate. Accepts multiple comma-separated values. |
| `expireTime` | `*int` | Query, Optional | Filter by certificate expiry time, in epoch seconds |
| `createdTime` | `*int` | Query, Optional | Filter by certificate issuance time, in epoch seconds |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.IssuedClientCertificatesResults](../../doc/models/issued-client-certificates-results.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSCEP.ListOrgIssuedClientCertificates(ctx, orgId, nil, nil, nil, nil, nil, nil, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "limit": 100,
  "page": 1,
  "results": [
    {
      "cert_provider": "jamf",
      "common_name": "john@corp.com",
      "created_time": 1431382121,
      "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "expire_time": 1718921115,
      "serial_number": "13 00 13 03 23 EE D5 84 01"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Revoke Org Issued Client Certificates

Revoke issued Mist SCEP client certificates by certificate serial number.

```go
RevokeOrgIssuedClientCertificates(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ClientCertSerialNumbers) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ClientCertSerialNumbers`](../../doc/models/client-cert-serial-numbers.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.ClientCertSerialNumbers{
    SerialNumbers:        []string{
        "13 00 13 03 23 EE D5 84 01",
    },
}

resp, err := orgsSCEP.RevokeOrgIssuedClientCertificates(ctx, orgId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Mist Scep

Update Mist SCEP settings for the organization, including enabled state, suspension state, and certificate providers.

```go
UpdateOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgSettingScep) (
    models.ApiResponse[models.OrgSettingScepResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgSettingScep`](../../doc/models/org-setting-scep.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSettingScepResponse](../../doc/models/org-setting-scep-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgSettingScep{
    AdditionalProperties: map[string]interface{}{
        "enabled": interface{}("true"),
    },
}

apiResponse, err := orgsSCEP.UpdateOrgMistScep(ctx, orgId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "cert_providers": [
    "jamf",
    "intune",
    "byod"
  ],
  "enabled": false,
  "intune_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_access_token": "1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=",
  "jamf_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_webhook_url": "https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

