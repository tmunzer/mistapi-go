# Admins Login

```go
adminsLogin := client.AdminsLogin()
```

## Class Name

`AdminsLogin`

## Methods

* [Login](../../doc/controllers/admins-login.md#login)
* [Two Factor](../../doc/controllers/admins-login.md#two-factor)


# Login

Log in with email/password.
When 2FA is enabled, there are two ways to login:

1. login with two_factor token (with Google Authenticator, etc)
2. login with email/password, generate the token, and use /login/two_factor with the token

```go
Login(
    ctx context.Context,
    body *models.Login) (
    models.ApiResponse[models.ResponseLoginSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.Login`](../../doc/models/login.md) | Body, Optional | - |

## Response Type

[`models.ResponseLoginSuccess`](../../doc/models/response-login-success.md)

## Example Usage

```go
ctx := context.Background()

body := models.Login{
    Email:                "test@mistsys.com",
    Password:             "foryoureyesonly",
    TwoFactor:            models.ToPointer("123456"),
}

apiResponse, err := adminsLogin.Login(ctx, &body)
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
  "email": "test@mistsys.com",
  "two_factor_passed": false,
  "two_factor_required": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Login Failed | [`ResponseLoginFailureException`](../../doc/models/response-login-failure-exception.md) |


# Two Factor

Send 2FA Code

```go
TwoFactor(
    ctx context.Context,
    body *models.TwoFactorString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.TwoFactorString`](../../doc/models/two-factor-string.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

body := models.TwoFactorString{
    TwoFactor:            "123456",
}

resp, err := adminsLogin.TwoFactor(ctx, &body)
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
| 401 | two_factor code is incorrect or the user hasn’t login yet | `ApiError` |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | the user doesn’t have 2FA enabled | `ApiError` |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

