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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoRoleOrg](../../doc/models/sso-role-org.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleOrg{
    Name:                 "string",
    Privileges:           []models.PrivilegeOrg{
        models.PrivilegeOrg{
            Role:                 models.PrivilegeOrgRoleEnum_ADMIN,
            Scope:                models.PrivilegeOrgScopeEnum_ORG,
            SiteId:               models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            AdditionalProperties: map[string]interface{}{
                "msp_id": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
                "msp_name": interface{}("string"),
                "name": interface{}("string"),
                "org_name": interface{}("string"),
                "orggroup_ids": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
                "sitegroup_ids": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoRoleOrg](../../doc/models/sso-role-org.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Sso Roles

Get List of Org SSO Roles

```go
ListOrgSsoRoles(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.SsoRoleOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.SsoRoleOrg](../../doc/models/sso-role-org.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSSORoles.ListOrgSsoRoles(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoRoleOrg](../../doc/models/sso-role-org.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleOrg{
    Name:                 "string",
    Privileges:           []models.PrivilegeOrg{
        models.PrivilegeOrg{
            Role:                 models.PrivilegeOrgRoleEnum_ADMIN,
            Scope:                models.PrivilegeOrgScopeEnum_ORG,
            SiteId:               models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            AdditionalProperties: map[string]interface{}{
                "msp_id": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
                "msp_name": interface{}("string"),
                "name": interface{}("string"),
                "org_name": interface{}("string"),
                "orggroup_ids": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
                "sitegroup_ids": interface{}("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

