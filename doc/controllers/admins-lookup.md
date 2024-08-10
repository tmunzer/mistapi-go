# Admins Lookup

```go
adminsLookup := client.AdminsLookup()
```

## Class Name

`AdminsLookup`


# Lookup

Login Lookup

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

[`models.ResponseLoginLookup`](../../doc/models/response-login-lookup.md)

## Example Usage

```go
ctx := context.Background()

body := models.EmailString{
    Email: "test@mistsys.com",
}

apiResponse, err := adminsLookup.Lookup(ctx, &body)
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
  "sso_url": "https://my.sso/idp_sso_url"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | user does not exist | `ApiError` |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

