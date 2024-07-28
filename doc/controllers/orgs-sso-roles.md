# Orgs SSO Roles

```go
orgsSSORoles := client.OrgsSSORoles()
```

## Class Name

`OrgsSSORoles`

## Methods

* [Create Org Sso Role](../../doc/controllers/orgs-sso-roles.md#create-org-sso-role)
* [Delete Org Sso Role](../../doc/controllers/orgs-sso-roles.md#delete-org-sso-role)
* [Get Org Sso Role](../../doc/controllers/orgs-sso-roles.md#get-org-sso-role)
* [List Org Sso Roles](../../doc/controllers/orgs-sso-roles.md#list-org-sso-roles)
* [Update Org Sso Role](../../doc/controllers/orgs-sso-roles.md#update-org-sso-role)


# Create Org Sso Role

Create Org SSO Role

```go
CreateOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SsoRoleOrg) (
    models.ApiResponse[models.SsoRoleOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsoRoleOrg`](../../doc/models/sso-role-org.md) | Body, Optional | Request Body |

## Response Type

[`models.SsoRoleOrg`](../../doc/models/sso-role-org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleOrg{
    Name:         "string",
    Privileges:   []models.PrivilegeOrg{
        models.PrivilegeOrg{
            Role:        models.PrivilegeOrgRoleEnum("admin"),
            Scope:       models.PrivilegeOrgScopeEnum("org"),
            SiteId:      models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
        },
    },
}

apiResponse, err := orgsSSORoles.CreateOrgSsoRole(ctx, orgId, &body)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "privileges": [
    {
      "role": "admin",
      "scope": "sitegroup",
      "sitegroup_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Delete Org Sso Role

Delete Org SSO Role

```go
DeleteOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoroleId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSSORoles.DeleteOrgSsoRole(ctx, orgId, ssoroleId)
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


# Get Org Sso Role

Get Org SSO Role Details

```go
GetOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID) (
    models.ApiResponse[models.SsoRoleOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoroleId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SsoRoleOrg`](../../doc/models/sso-role-org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSORoles.GetOrgSsoRole(ctx, orgId, ssoroleId)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "privileges": [
    {
      "role": "admin",
      "scope": "sitegroup",
      "sitegroup_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# List Org Sso Roles

Get List of Org SSO Roles

```go
ListOrgSsoRoles(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.SsoRoleMsp],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.SsoRoleMsp`](../../doc/models/sso-role-msp.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsSSORoles.ListOrgSsoRoles(ctx, orgId, &page, &limit)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Update Org Sso Role

Update Org SSO Role

```go
UpdateOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID,
    body *models.SsoRoleOrg) (
    models.ApiResponse[models.SsoRoleOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoroleId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsoRoleOrg`](../../doc/models/sso-role-org.md) | Body, Optional | Request Body |

## Response Type

[`models.SsoRoleOrg`](../../doc/models/sso-role-org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleOrg{
    Name:         "string",
    Privileges:   []models.PrivilegeOrg{
        models.PrivilegeOrg{
            Role:        models.PrivilegeOrgRoleEnum("admin"),
            Scope:       models.PrivilegeOrgScopeEnum("org"),
            SiteId:      models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
        },
    },
}

apiResponse, err := orgsSSORoles.UpdateOrgSsoRole(ctx, orgId, ssoroleId, &body)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "privileges": [
    {
      "role": "admin",
      "scope": "sitegroup",
      "sitegroup_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"
    }
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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

