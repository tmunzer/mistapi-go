# Admins

Admin API calls can be used to create, manage or authenticate Mist administrators.

To register administrators into an existing MSP account or Organization, please check: * [Invite Msp Admin](../../doc/controllers/ms-ps-admins.md#invite-msp-admin) * [Invite Org Admin](../../doc/controllers/orgs-admins.md#invite-org-admin)

```go
admins := client.Admins()
```

## Class Name

`Admins`

## Methods

* [Get Admin Registration Info](../../doc/controllers/admins.md#get-admin-registration-info)
* [Register New Admin](../../doc/controllers/admins.md#register-new-admin)
* [Verify Admin Invite](../../doc/controllers/admins.md#verify-admin-invite)
* [Verify Registration](../../doc/controllers/admins.md#verify-registration)


# Get Admin Registration Info

Return the public CAPTCHA settings required for administrator registration. This public endpoint does not require authentication. Use the returned `flavor`, `required`, and `sitekey` values to render the correct CAPTCHA challenge before calling [Register New Admin](../../doc/controllers/admins.md#register-new-admin).

:information_source: **Note** This endpoint does not require authentication.

```go
GetAdminRegistrationInfo(
    ctx context.Context,
    recaptchaFlavor *models.RecaptchaFlavorEnum) (
    models.ApiResponse[models.Recaptcha],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Query, Optional | Filter login settings by reCAPTCHA flavor. enum: `google`, `hcaptcha`<br><br>**Default**: `"google"` |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Recaptcha](../../doc/models/recaptcha.md).

## Example Usage

```go
ctx := context.Background()

recaptchaFlavor := models.RecaptchaFlavorEnum_HCAPTCHA

apiResponse, err := admins.GetAdminRegistrationInfo(ctx, &recaptchaFlavor)
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
  "flavor": "google",
  "required": true,
  "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
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


# Register New Admin

Register a new administrator account and initial organization. This public endpoint does not require authentication. Mist sends a verification email containing a link such as `/verify/register?token={token}`; use [Verify Registration](../../doc/controllers/admins.md#verify-registration) to complete registration with that token.

Use [Get Registration Information](../../doc/controllers/admins.md#get-admin-registration-info) before submitting this request to determine whether CAPTCHA is required, which CAPTCHA provider to render, and which public site key to use. If CAPTCHA is required, include the provider response token in `recaptcha` and the provider name in `recaptcha_flavor`.

:information_source: **Note** This endpoint does not require authentication.

```go
RegisterNewAdmin(
    ctx context.Context,
    body *models.AdminInvite) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.AdminInvite`](../../doc/models/admin-invite.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.AdminInvite{
    AccountOnly:          models.ToPointer(false),
    AllowMist:            models.ToPointer(false),
    City:                 models.ToPointer("Cupertino"),
    Country:              models.ToPointer("United States"),
    Email:                "test@mistsys.com",
    FirstName:            "John",
    InviteCode:           models.ToPointer("MISTROCKS"),
    LastName:             "Smith",
    OrgName:              "Smith LLC",
    Password:             "foryoureyesonly",
    Recaptcha:            "string",
    RecaptchaFlavor:      models.ToPointer(models.RecaptchaFlavorEnum_HCAPTCHA),
    RefererInviteToken:   models.ToPointer("Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat"),
    ReturnTo:             models.ToPointer("https://mist.zendesk.com/hc/quickstart.pdf"),
    State:                models.ToPointer("California"),
    StreetAddress:        models.ToPointer("1601 S De Anza Blvd Ste 248"),
    StreetAddress2:       models.ToPointer("1601 S De Anza Blvd Ste 248"),
    Zipcode:              models.ToPointer("95014"),
}

resp, err := admins.RegisterNewAdmin(ctx, &body)
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


# Verify Admin Invite

Accept an administrator invite using the invite verification token from the invite email. This public endpoint does not require authentication. After a successful verification, call [Get Self](../../doc/controllers/self-account.md#get-self) to refresh the authenticated admin profile and retrieve the newly granted privileges.

:information_source: **Note** This endpoint does not require authentication.

```go
VerifyAdminInvite(
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

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

token := "token6"

resp, err := admins.VerifyAdminInvite(ctx, token)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseDetailString:
            log.Fatalln("ResponseDetailStringException: ", typedErr)
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
| 404 | Not Found | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Verify Registration

Verify a new administrator registration using the token from the registration email. This public endpoint does not require authentication. A successful verification creates a login session and may also apply a pending invitation; the response indicates whether an invitation could not be applied automatically.

:information_source: **Note** This endpoint does not require authentication.

```go
VerifyRegistration(
    ctx context.Context,
    token string) (
    models.ApiResponse[models.ResponseVerifyTokenSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `token` | `string` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseVerifyTokenSuccess](../../doc/models/response-verify-token-success.md).

## Example Usage

```go
ctx := context.Background()

token := "token6"

apiResponse, err := admins.VerifyRegistration(ctx, token)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseDetailString:
            log.Fatalln("ResponseDetailStringException: ", typedErr)
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
  "return_to": "https://mist.zendesk.com/hc/quickstart.pdf"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Response if verification expired or already registered | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Response if secret is invalid | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

