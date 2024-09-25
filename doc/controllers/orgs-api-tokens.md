# Orgs API Tokens

```go
orgsAPITokens := client.OrgsAPITokens()
```

## Class Name

`OrgsAPITokens`

## Methods

* [Create Org Api Token](../../doc/controllers/orgs-api-tokens.md#create-org-api-token)
* [Delete Org Api Token](../../doc/controllers/orgs-api-tokens.md#delete-org-api-token)
* [Get Org Api Token](../../doc/controllers/orgs-api-tokens.md#get-org-api-token)
* [List Org Api Tokens](../../doc/controllers/orgs-api-tokens.md#list-org-api-tokens)
* [Update Org Api Token](../../doc/controllers/orgs-api-tokens.md#update-org-api-token)


# Create Org Api Token

Create Org API Token
Note that the token key is only available during creation time.

```go
CreateOrgApiToken(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgApitoken) (
    models.ApiResponse[models.OrgApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgApitoken`](../../doc/models/org-apitoken.md) | Body, Optional | - |

## Response Type

[`models.OrgApitoken`](../../doc/models/org-apitoken.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgApitoken{
    CreatedBy:   models.NewOptional(models.ToPointer("user@mycorp.com")),
    CreatedTime: models.ToPointer(float64(1626875902)),
    Id:          models.ToPointer(uuid.MustParse("497f6eca-6276-4993-bfeb-53ecbbba6f08")),
    Key:         models.ToPointer("1qkb...QQCL"),
    LastUsed:    models.NewOptional(models.ToPointer(float64(1690115110))),
    Name:        "org_token_xyz",
    OrgId:       models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    Privileges:  []models.PrivilegeOrg{
        models.PrivilegeOrg{
            OrgId:       models.ToPointer(uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b")),
            Role:        models.PrivilegeOrgRoleEnum("admin"),
            Scope:       models.PrivilegeOrgScopeEnum("org"),
        },
    },
    SrcIps:      []string{
        "63.3.56.0/24",
        "63.3.55.4",
    },
}

apiResponse, err := orgsAPITokens.CreateOrgApiToken(ctx, orgId, &body)
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
  "created_by": "user@mycorp.com",
  "created_time": 1626875902,
  "id": "497f6eca-6276-4993-bfeb-53efbbba6f08",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "privileges": [
    {
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "role": "admin",
      "scope": "org"
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


# Delete Org Api Token

Delete Org API Token

```go
DeleteOrgApiToken(
    ctx context.Context,
    orgId uuid.UUID,
    apitokenId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `apitokenId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAPITokens.DeleteOrgApiToken(ctx, orgId, apitokenId)
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


# Get Org Api Token

Get Org API Token

```go
GetOrgApiToken(
    ctx context.Context,
    orgId uuid.UUID,
    apitokenId uuid.UUID) (
    models.ApiResponse[models.OrgApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `apitokenId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.OrgApitoken`](../../doc/models/org-apitoken.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAPITokens.GetOrgApiToken(ctx, orgId, apitokenId)
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
  "created_by": "user@mycorp.com",
  "created_time": 1626875902,
  "id": "497f6eca-6276-4993-bfeb-53efbbba6f08",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "privileges": [
    {
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "role": "admin",
      "scope": "org"
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


# List Org Api Tokens

Get List of Org API Tokens

```go
ListOrgApiTokens(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.OrgApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.OrgApitoken`](../../doc/models/org-apitoken.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAPITokens.ListOrgApiTokens(ctx, orgId)
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
    "created_by": "user@mycorp.com",
    "created_time": 1626875902,
    "id": "497f6eca-6276-4993-bfeb-53f0bbba6f08",
    "key": "1qkb...QQCL",
    "last_used": 1690115110,
    "name": "org_token_xyz",
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "privileges": [
      {
        "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
        "role": "admin",
        "scope": "org"
      }
    ]
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Api Token

Update Org API Token

```go
UpdateOrgApiToken(
    ctx context.Context,
    orgId uuid.UUID,
    apitokenId uuid.UUID,
    body *models.OrgApitoken) (
    models.ApiResponse[models.OrgApitoken],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `apitokenId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgApitoken`](../../doc/models/org-apitoken.md) | Body, Optional | Request Body |

## Response Type

[`models.OrgApitoken`](../../doc/models/org-apitoken.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apitokenId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgApitoken{
    Name:        "org_token_xyz",
    Privileges:  []models.PrivilegeOrg{
        models.PrivilegeOrg{
            OrgId:       models.ToPointer(uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b")),
            Role:        models.PrivilegeOrgRoleEnum("admin"),
            Scope:       models.PrivilegeOrgScopeEnum("org"),
        },
    },
}

apiResponse, err := orgsAPITokens.UpdateOrgApiToken(ctx, orgId, apitokenId, &body)
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
  "created_by": "user@mycorp.com",
  "created_time": 1626875902,
  "id": "497f6eca-6276-4993-bfeb-53efbbba6f08",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "privileges": [
    {
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "role": "admin",
      "scope": "org"
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

