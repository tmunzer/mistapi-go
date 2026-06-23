# Orgs Devices-SSR

```go
orgsDevicesSSR := client.OrgsDevicesSSR()
```

## Class Name

`OrgsDevicesSSR`

## Methods

* [Export Org Ssr Id Tokens](../../doc/controllers/orgs-devices-ssr.md#export-org-ssr-id-tokens)
* [Get Org 128 T Registration Commands](../../doc/controllers/orgs-devices-ssr.md#get-org-128-t-registration-commands)
* [Get Org Ssr Registration Commands](../../doc/controllers/orgs-devices-ssr.md#get-org-ssr-registration-commands)


# Export Org Ssr Id Tokens

Export SSR ID tokens for the requested device MAC addresses so they can be imported into Conductor during SSR onboarding.

```go
ExportOrgSsrIdTokens(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseSsrExportIdTokens],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsrExportIdTokens](../../doc/models/response-ssr-export-id-tokens.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "683b679ac024",
    },
}

apiResponse, err := orgsDevicesSSR.ExportOrgSsrIdTokens(ctx, orgId, &body)
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
  "results": [
    {
      "mac": "025b35000020",
      "token": "string"
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


# Get Org 128 T Registration Commands

**This endpoint is deprecated.**

128T devices can be managed/adopted by Mist.

```go
GetOrg128TRegistrationCommands(
    ctx context.Context,
    orgId uuid.UUID,
    ttl *int,
    assetIds []string) (
    models.ApiResponse[models.ResponseRouterSsrRegisterCmd],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ttl` | `*int` | Query, Optional | Token validity duration in seconds. Defaults to 1 year (31536000 seconds) |
| `assetIds` | `[]string` | Query, Optional | When specified restricts registration to listed assets only. Prefer HTTP body over headers for this parameter, especially with long lists to avoid header size limits.<br><br>**Constraints**: *Unique Items Required* |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRouterSsrRegisterCmd](../../doc/models/response-router-ssr-register-cmd.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ttl := 31536000

apiResponse, err := orgsDevicesSSR.GetOrg128TRegistrationCommands(ctx, orgId, &ttl, nil)
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
  "conductor_cmd": "register mist eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "registration_code": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "router_shell_cmd": "128agent register --registration-code eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ"
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


# Get Org Ssr Registration Commands

Return the registration token and conductor or router commands used to register SSR routers with Mist. The optional TTL controls token validity, and asset IDs can restrict registration to specific assets.

```go
GetOrgSsrRegistrationCommands(
    ctx context.Context,
    orgId uuid.UUID,
    ttl *int,
    assetIds []string) (
    models.ApiResponse[models.ResponseRouterSsrRegisterCmd],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ttl` | `*int` | Query, Optional | Token validity duration in seconds. Defaults to 1 year (31536000 seconds) |
| `assetIds` | `[]string` | Query, Optional | When specified restricts registration to listed assets only. Prefer HTTP body over headers for this parameter, especially with long lists to avoid header size limits.<br><br>**Constraints**: *Unique Items Required* |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRouterSsrRegisterCmd](../../doc/models/response-router-ssr-register-cmd.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ttl := 31536000

apiResponse, err := orgsDevicesSSR.GetOrgSsrRegistrationCommands(ctx, orgId, &ttl, nil)
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
  "conductor_cmd": "register mist eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "registration_code": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "router_shell_cmd": "128agent register --registration-code eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ"
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

