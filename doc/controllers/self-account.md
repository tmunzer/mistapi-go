# Self Account

```go
selfAccount := client.SelfAccount()
```

## Class Name

`SelfAccount`

## Methods

* [Delete Self](../../doc/controllers/self-account.md#delete-self)
* [Get Self](../../doc/controllers/self-account.md#get-self)
* [Get Self Api Usage](../../doc/controllers/self-account.md#get-self-api-usage)
* [Get Self Login Failures](../../doc/controllers/self-account.md#get-self-login-failures)
* [Update Self](../../doc/controllers/self-account.md#update-self)
* [Update Self Email](../../doc/controllers/self-account.md#update-self-email)
* [Verify Self Email](../../doc/controllers/self-account.md#verify-self-email)


# Delete Self

To delete ones account and every associated with it. The effects:

the account would be deleted
any orphaned Org (that only has this account as admin) will be deleted
along with all data with Org (sites, wlans, devices) will be gone.

```go
DeleteSelf(
    ctx context.Context) (
    http.Response,
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

resp, err := selfAccount.DeleteSelf(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request | [`ErrorDeleteFailedException`](../../doc/models/error-delete-failed-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Self

Get ‘whoami’ and privileges (which org and which sites I have access to)

```go
GetSelf(
    ctx context.Context) (
    models.ApiResponse[models.Admin],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := selfAccount.GetSelf(ctx)
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
  "admin_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "email": "user@example.com",
  "first_name": "string",
  "last_name": "string",
  "password_modified_time": 1656353525,
  "privileges": [
    {
      "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "msp_name": "string",
      "name": "string",
      "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "org_name": "string",
      "orggroup_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ],
      "role": "admin",
      "scope": "org",
      "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "sitegroup_ids": [
        "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
      ]
    }
  ]
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


# Get Self Api Usage

Get the status of the API usage for the current user or API Token

```go
GetSelfApiUsage(
    ctx context.Context) (
    models.ApiResponse[models.ApiUsage],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ApiUsage](../../doc/models/api-usage.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := selfAccount.GetSelfApiUsage(ctx)
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
  "request_limit": 5000,
  "requests": 5
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


# Get Self Login Failures

Get a list of failed login attempts across all Orgs for the current admin

```go
GetSelfLoginFailures(
    ctx context.Context) (
    models.ApiResponse[models.LoginFailures],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.LoginFailures](../../doc/models/login-failures.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := selfAccount.GetSelfLoginFailures(ctx)
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
  "email": "admin@test.com",
  "last_failure_at": 1509161968,
  "num_attempts": 40,
  "src_ips": [
    "192.168.1.39",
    "192.168.1.38",
    "192.168.1.37"
  ],
  "user_agents": [
    "Test UA 39",
    "Test UA 38",
    "Test UA 37"
  ]
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


# Update Self

Update Account Information

```go
UpdateSelf(
    ctx context.Context,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

body := models.Admin{
    Email:                models.ToPointer("john.smith@mycorp.net"),
    FirstName:            models.ToPointer("John"),
    LastName:             models.ToPointer("Smith"),
    Phone:                models.ToPointer("14081112222"),
    Phone2:               models.ToPointer("14083334444"),
    AdditionalProperties: map[string]interface{}{
        "persona": interface{}("security"),
    },
}

apiResponse, err := selfAccount.UpdateSelf(ctx, &body)
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
  "admin_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "email": "user@example.com",
  "first_name": "string",
  "last_name": "string",
  "password_modified_time": 1656353525,
  "privileges": [
    {
      "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "msp_name": "string",
      "name": "string",
      "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "org_name": "string",
      "orggroup_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ],
      "role": "admin",
      "scope": "org",
      "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "sitegroup_ids": [
        "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
      ]
    }
  ]
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


# Update Self Email

Change Email
We require the user to verify that they actually own the email address they intend to change it to.

After the API call, the user will receive an email to the new email address with a link like https://manage.mist.com/verify/update?expire=:exp_time&email=:admin_email&token=:token

Upon clicking the link, the user is provided with a login page to authenticate using existing credentials. After successful login, the email address of the user gets updated

**Note**: The request parameter email can be used by UI to validate that the current session (if any) belongs to the admin or provide a login page (by pre-populating the email on login screen). UI can also use the request parameter expire to validate token expiry.

```go
UpdateSelfEmail(
    ctx context.Context,
    body *models.EmailString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.EmailString`](../../doc/models/email-string.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

body := models.EmailString{
    Email:                "new@mistsys.com",
}

resp, err := selfAccount.UpdateSelfEmail(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid email address or new email address already exists | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Verify Self Email

Verify Email change

```go
VerifySelfEmail(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

token := "token6"

resp, err := selfAccount.VerifySelfEmail(ctx, token)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Request | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

