# MS Ps SSO Roles

```go
mSPsSSORoles := client.MSPsSSORoles()
```

## Class Name

`MSPsSSORoles`

## Methods

* [Create Msp Sso Role](../../doc/controllers/ms-ps-sso-roles.md#create-msp-sso-role)
* [Delete Msp Sso Role](../../doc/controllers/ms-ps-sso-roles.md#delete-msp-sso-role)
* [List Msp Sso Roles](../../doc/controllers/ms-ps-sso-roles.md#list-msp-sso-roles)
* [Update Msp Sso Role](../../doc/controllers/ms-ps-sso-roles.md#update-msp-sso-role)


# Create Msp Sso Role

Create MSP Role

```go
CreateMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.SsoRoleMsp) (
    models.ApiResponse[models.SsoRoleMsp],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsoRoleMsp`](../../doc/models/sso-role-msp.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoRoleMsp](../../doc/models/sso-role-msp.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleMsp{
    Name:                 "name6",
    Privileges:           []models.PrivilegeMsp{
        models.PrivilegeMsp{
            Role:                 models.PrivilegeMspRoleEnum_ADMIN,
            Scope:                models.PrivilegeMspScopeEnum_ORG,
        },
    },
}

apiResponse, err := mSPsSSORoles.CreateMspSsoRole(ctx, mspId, &body)
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
      "orggroup_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "role": "read",
      "scope": "orggroup"
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


# Delete Msp Sso Role

Delete MSP SSO Roles

```go
DeleteMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    ssoroleId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoroleId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsSSORoles.DeleteMspSsoRole(ctx, mspId, ssoroleId)
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


# List Msp Sso Roles

Get List of MSP SSO Roles

```go
ListMspSsoRoles(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.SsoRoleMsp],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.SsoRoleMsp](../../doc/models/sso-role-msp.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSORoles.ListMspSsoRoles(ctx, mspId)
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


# Update Msp Sso Role

Update SSO Role

```go
UpdateMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    ssoroleId uuid.UUID,
    body *models.SsoRoleMsp) (
    models.ApiResponse[models.SsoRoleMsp],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoroleId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsoRoleMsp`](../../doc/models/sso-role-msp.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoRoleMsp](../../doc/models/sso-role-msp.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoroleId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoRoleMsp{
    Name:                 "name6",
    Privileges:           []models.PrivilegeMsp{
        models.PrivilegeMsp{
            Role:                 models.PrivilegeMspRoleEnum_ADMIN,
            Scope:                models.PrivilegeMspScopeEnum_ORG,
        },
    },
}

apiResponse, err := mSPsSSORoles.UpdateMspSsoRole(ctx, mspId, ssoroleId, &body)
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
      "orggroup_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "role": "read",
      "scope": "orggroup"
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

