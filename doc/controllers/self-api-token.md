# Self API Token

```go
selfAPIToken := client.SelfAPIToken()
```

## Class Name

`SelfAPIToken`

## Methods

* [Create Api Token](../../doc/controllers/self-api-token.md#create-api-token)
* [Delete Api Token](../../doc/controllers/self-api-token.md#delete-api-token)
* [Get Api Token](../../doc/controllers/self-api-token.md#get-api-token)
* [List Api Tokens](../../doc/controllers/self-api-token.md#list-api-tokens)
* [Update Api Token](../../doc/controllers/self-api-token.md#update-api-token)


# Create Api Token

Create API Token
Note that the key is only available during creation time.

```go
CreateApiToken(
    ctx context.Context,
    body *models.UserApitoken) (
    models.ApiResponse[[]models.UserApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.UserApitoken`](../../doc/models/user-apitoken.md) | Body, Optional | - |

## Response Type

[`[]models.UserApitoken`](../../doc/models/user-apitoken.md)

## Example Usage

```go
ctx := context.Background()

body := models.UserApitoken{
    Name:        models.ToPointer("org_token_xyz"),
}

apiResponse, err := selfAPIToken.CreateApiToken(ctx, &body)
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
[
  {
    "created_time": 1626875902,
    "id": "864f351a-1377-4ad9-83f8-72f3fe6199ba",
    "key": "1qkb...QQCL",
    "last_used": 1690115110,
    "name": "org_token_xyz"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Api Token

Delete an API Token

```go
DeleteApiToken(
    ctx context.Context,
    apitokenId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `apitokenId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := selfAPIToken.DeleteApiToken(ctx, apitokenId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Api Token

Get User API Token

```go
GetApiToken(
    ctx context.Context,
    apitokenId uuid.UUID) (
    models.ApiResponse[models.UserApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `apitokenId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.UserApitoken`](../../doc/models/user-apitoken.md)

## Example Usage

```go
ctx := context.Background()

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := selfAPIToken.GetApiToken(ctx, apitokenId)
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
  "created_time": 1626875902,
  "id": "864f351a-1377-4ad9-83f8-72f3fe6199ba",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz"
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


# List Api Tokens

Get List of Current User API Tokens

```go
ListApiTokens(
    ctx context.Context) (
    models.ApiResponse[[]models.UserApitoken],
    error)
```

## Response Type

[`[]models.UserApitoken`](../../doc/models/user-apitoken.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := selfAPIToken.ListApiTokens(ctx)
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
[
  {
    "created_time": 1626875902,
    "id": "864f351a-1377-4ad9-83f8-72f3fe6199ba",
    "key": "1qkb...QQCL",
    "last_used": 1690115110,
    "name": "org_token_xyz"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Api Token

Update User API Token

```go
UpdateApiToken(
    ctx context.Context,
    apitokenId uuid.UUID,
    body *models.UserApitoken) (
    models.ApiResponse[models.UserApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `apitokenId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UserApitoken`](../../doc/models/user-apitoken.md) | Body, Optional | - |

## Response Type

[`models.UserApitoken`](../../doc/models/user-apitoken.md)

## Example Usage

```go
ctx := context.Background()

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UserApitoken{
    Name:        models.ToPointer("org_token_xyz"),
}

apiResponse, err := selfAPIToken.UpdateApiToken(ctx, apitokenId, &body)
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
  "created_time": 1626875902,
  "id": "864f351a-1377-4ad9-83f8-72f3fe6199ba",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz"
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

