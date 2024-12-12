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

Get Registration Information

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
| `recaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Query, Optional | **Default**: `"google"` |

## Response Type

[`models.Recaptcha`](../../doc/models/recaptcha.md)

## Example Usage

```go
ctx := context.Background()

recaptchaFlavor := models.RecaptchaFlavorEnum("hcaptcha")

apiResponse, err := admins.GetAdminRegistrationInfo(ctx, &recaptchaFlavor)
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
  "flavor": "google",
  "required": true,
  "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
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


# Register New Admin

Register a new admin and his/her org
An email will also be sent to the user with a link to `/verify/register?token={token}`

### reCAPTCHA

Google reCAPTCHA is the choice to prevent bot registration

It needs this

&lt;script src='https://www.google.com/recaptcha/api.js' &gt;&lt;/script&gt;

and this &lt;div&gt; in the desired place

```html
<div class="g-recaptcha" data_sitekey="6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"></div>
```

Use GET /api/v1/register/recaptcha to read the current setting.
Response example:

```json
{    
  "flavor": "google",
  "required": true,    
  "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
}
```

### hCaptcha

Alternative to reCAPTCHA is hCaptcha to prevent bot registration

It needs this script

&lt;script src='https://js.hcaptcha.com/1/api.js' async defer &gt;&lt;/script&gt;

and this &lt;div&gt; in the desired place

```html
<div class="h-recaptcha" data_sitekey="6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"></div>
```

Use GET /api/v1/register/recaptcha?recaptcha_flavor=hcaptcha to read the current setting for hcaptcha with reply.
Response example:

```json
{
  "flavor": "hcaptcha",
  "required": true,
  "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
}"
```

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

``

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
    RecaptchaFlavor:      models.ToPointer(models.RecaptchaFlavorEnum("hcaptcha")),
    RefererInviteToken:   models.ToPointer("Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat"),
    ReturnTo:             models.ToPointer("http://mist.zendesk.com/hc/quickstart.pdf"),
    State:                models.ToPointer("California"),
    StreetAddress:        models.ToPointer("1601 S De Anza Blvd Ste 248"),
    StreetAddress2:       models.ToPointer("1601 S De Anza Blvd Ste 248"),
    Zipcode:              models.ToPointer("95014"),
}

resp, err := admins.RegisterNewAdmin(ctx, &body)
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


# Verify Admin Invite

**Note**: another call to `GET /api/v1/self` is required to see the new set of privileges

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

``

## Example Usage

```go
ctx := context.Background()

token := "token6"

resp, err := admins.VerifyAdminInvite(ctx, token)
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
| 404 | Not Found | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Verify Registration

Verify registration

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

[`models.ResponseVerifyTokenSuccess`](../../doc/models/response-verify-token-success.md)

## Example Usage

```go
ctx := context.Background()

token := "token6"

apiResponse, err := admins.VerifyRegistration(ctx, token)
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
  "return_to": "http://mist.zendesk.com/hc/quickstart.pdf"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Response if verification expired or already registered | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Response if secret is invalid | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

