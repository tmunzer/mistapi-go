# Admins Recover Password

```go
adminsRecoverPassword := client.AdminsRecoverPassword()
```

## Class Name

`AdminsRecoverPassword`

## Methods

* [Recover Password](../../doc/controllers/admins-recover-password.md#recover-password)
* [Verify Recover Passsword](../../doc/controllers/admins-recover-password.md#verify-recover-passsword)


# Recover Password

Recover Password
An email will also be sent to the user with a link to https://manage.mist.com/verify/recover?token=:token

```go
RecoverPassword(
    ctx context.Context,
    body *models.Recover) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.Recover`](../../doc/models/recover.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

body := models.Recover{
    Email:                "test@mistsys.com",
    Recaptcha:            models.ToPointer("string"),
    RecaptchaFlavor:      models.ToPointer(models.RecaptchaFlavorEnum_HCAPTCHA),
}

resp, err := adminsRecoverPassword.RecoverPassword(ctx, &body)
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


# Verify Recover Passsword

Verify Recover Password
With correct verification, the user will be authenticated. UI can then prompt for new password

```go
VerifyRecoverPasssword(
    ctx context.Context,
    token string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `token` | `string` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

token := "token6"

resp, err := adminsRecoverPassword.VerifyRecoverPasssword(ctx, token)
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

