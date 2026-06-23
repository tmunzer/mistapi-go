# Admins Lookup

```go
adminsLookup := client.AdminsLookup()
```

## Class Name

`AdminsLookup`


# Lookup

Check the login method for an administrator email address. This public lookup is mainly used by UI clients to determine whether the user should continue with local login or be redirected to SSO.

:information_source: **Note** This endpoint does not require authentication.

```go
Lookup(
    ctx context.Context,
    body *models.EmailString) (
    models.ApiResponse[models.ResponseLoginLookup],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.EmailString`](../../doc/models/email-string.md) | Body, Optional | Request Body |

## Response Type

**200**: Account exists

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLoginLookup](../../doc/models/response-login-lookup.md).

## Example Usage

```go
ctx := context.Background()

body := models.EmailString{
    Email:                "test@mistsys.com",
}

apiResponse, err := adminsLookup.Lookup(ctx, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
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
  "sso_url": "https://my.sso/idp_sso_url"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | User does not exist | `ApiError` |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

