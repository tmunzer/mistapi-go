# Admins Logout

```go
adminsLogout := client.AdminsLogout()
```

## Class Name

`AdminsLogout`


# Logout

Logout

```go
Logout(
    ctx context.Context) (
    models.ApiResponse[models.ResponseLogout],
    error)
```

## Response Type

[`models.ResponseLogout`](../../doc/models/response-logout.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := adminsLogout.Logout(ctx)
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
  "forward_url": "https://my.sso/custom_logout_url"
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

