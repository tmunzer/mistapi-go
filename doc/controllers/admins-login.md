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

Authenticate an administrator with email and password. A successful login creates the browser session cookies, including the `csrftoken` value used with the `X-CSRFToken` header on later API requests.

When 2FA is enabled, either include the `two_factor` code in this request or submit the first factor here and complete the login with [Two Factor](../../doc/controllers/admins-login.md#two-factor).

:information_source: **Note** This endpoint does not require authentication.

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

**200**: Login Success

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLoginSuccess](../../doc/models/response-login-success.md).

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
    switch typedErr := err.(type) {
        case *errors.ResponseLoginFailure:
            log.Fatalln("ResponseLoginFailureException: ", typedErr)
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

Complete a two-factor login by submitting the 2FA code after the initial email/password step has created a pending login session.

:information_source: **Note** This endpoint does not require authentication.

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

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.TwoFactorString{
    TwoFactor:            "123456",
}

resp, err := adminsLogin.TwoFactor(ctx, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
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
| 401 | two_factor code is incorrect or the user hasn't login yet | `ApiError` |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | The user doesn't have 2FA enabled | `ApiError` |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

