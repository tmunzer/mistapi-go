# Self O Auth 2

```go
selfOAuth2 := client.SelfOAuth2()
```

## Class Name

`SelfOAuth2`

## Methods

* [Get Oauth 2 Url for Linking](../../doc/controllers/self-o-auth-2.md#get-oauth-2-url-for-linking)
* [Link Oauth 2 Mist Account](../../doc/controllers/self-o-auth-2.md#link-oauth-2-mist-account)


# Get Oauth 2 Url for Linking

Obtain Authorization URL for Linking

```go
GetOauth2UrlForLinking(
    ctx context.Context,
    provider string,
    forward *string) (
    models.ApiResponse[models.ResponseSelfOauthUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `provider` | `string` | Template, Required | - |
| `forward` | `*string` | Query, Optional | - |

## Response Type

[`models.ResponseSelfOauthUrl`](../../doc/models/response-self-oauth-url.md)

## Example Usage

```go
ctx := context.Background()

provider := "provider8"

forward := "http://manage.mist.com/oauth/callback.html"

apiResponse, err := selfOAuth2.GetOauth2UrlForLinking(ctx, provider, &forward)
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
  "authorization_url": "https://accounts.google.com/o/oauth2/v2/auth?.....",
  "linked": false
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Link Oauth 2 Mist Account

Link Mist account with an OAuth2 Provider

```go
LinkOauth2MistAccount(
    ctx context.Context,
    provider string,
    body *models.CodeString) (
    models.ApiResponse[models.ResponseSelfOauthLinkSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `provider` | `string` | Template, Required | - |
| `body` | [`*models.CodeString`](../../doc/models/code-string.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseSelfOauthLinkSuccess`](../../doc/models/response-self-oauth-link-success.md)

## Example Usage

```go
ctx := context.Background()

provider := "provider8"

body := models.CodeString{
    Code: "4/S9tegDeLkrYg0L9pWNXV4cgMVbbr3SR9t693A2kSHzw",
}

apiResponse, err := selfOAuth2.LinkOauth2MistAccount(ctx, provider, &body)
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
  "action": "oauth account linked",
  "id": "google_user_id"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Authorization Error | [`ResponseSelfOauthLinkFailureException`](../../doc/models/response-self-oauth-link-failure-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

