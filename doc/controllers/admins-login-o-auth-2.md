# Admins Login-O Auth 2

```go
adminsLoginOAuth2 := client.AdminsLoginOAuth2()
```

## Class Name

`AdminsLoginOAuth2`

## Methods

* [Get Oauth 2 Authorization Url for Login](../../doc/controllers/admins-login-o-auth-2.md#get-oauth-2-authorization-url-for-login)
* [Login Oauth 2](../../doc/controllers/admins-login-o-auth-2.md#login-oauth-2)
* [Unlink Oauth 2 Provider](../../doc/controllers/admins-login-o-auth-2.md#unlink-oauth-2-provider)


# Get Oauth 2 Authorization Url for Login

Return the provider authorization URL used to start an OAuth2 login or account-linking flow. When `forward` is provided, the provider redirects back to that callback URL after authorization.

:information_source: **Note** This endpoint does not require authentication.

```go
GetOauth2AuthorizationUrlForLogin(
    ctx context.Context,
    provider string,
    forward *string) (
    models.ApiResponse[models.ResponseLoginOauthUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `provider` | `string` | Template, Required | - |
| `forward` | `*string` | Query, Optional | Callback URL used after completing the OAuth login flow |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLoginOauthUrl](../../doc/models/response-login-oauth-url.md).

## Example Usage

```go
ctx := context.Background()

provider := "google"

forward := "https://manage.mist.com/oauth/callback.html"

apiResponse, err := adminsLoginOAuth2.GetOauth2AuthorizationUrlForLogin(ctx, provider, &forward)
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
  "authorization_url": "https://accounts.google.com/o/oauth2/v2/auth?.....",
  "client_id": "173131512-mpbnju32.apps.googleusercontent.com"
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


# Login Oauth 2

Complete an OAuth2 login by exchanging the provider authorization code for a Mist administrator session.

:information_source: **Note** This endpoint does not require authentication.

```go
LoginOauth2(
    ctx context.Context,
    provider string,
    body *models.CodeString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `provider` | `string` | Template, Required | - |
| `body` | [`*models.CodeString`](../../doc/models/code-string.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

provider := "google"

body := models.CodeString{
    Code:                 "4/S9tegDeLkrYg0L9pWNXV4cgMVbbr3SR9t693A2kSHzw",
}

resp, err := adminsLoginOAuth2.LoginOauth2(ctx, provider, &body)
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


# Unlink Oauth 2 Provider

Unlink the specified OAuth2 provider from the authenticated administrator account so it can no longer be used for that account's OAuth login.

```go
UnlinkOauth2Provider(
    ctx context.Context,
    provider string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `provider` | `string` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

provider := "google"

resp, err := adminsLoginOAuth2.UnlinkOauth2Provider(ctx, provider)
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

