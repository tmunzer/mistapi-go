# MS Ps Admins

```go
mSPsAdmins := client.MSPsAdmins()
```

## Class Name

`MSPsAdmins`

## Methods

* [Get Msp Admin](../../doc/controllers/ms-ps-admins.md#get-msp-admin)
* [Invite Msp Admin](../../doc/controllers/ms-ps-admins.md#invite-msp-admin)
* [List Msp Admins](../../doc/controllers/ms-ps-admins.md#list-msp-admins)
* [Revoke Msp Admin](../../doc/controllers/ms-ps-admins.md#revoke-msp-admin)
* [Uninvite Msp Admin](../../doc/controllers/ms-ps-admins.md#uninvite-msp-admin)
* [Update Msp Admin](../../doc/controllers/ms-ps-admins.md#update-msp-admin)
* [Update Msp Admin Invite](../../doc/controllers/ms-ps-admins.md#update-msp-admin-invite)


# Get Msp Admin

Get MSP Admins

```go
GetMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsAdmins.GetMspAdmin(ctx, mspId, adminId)
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


# Invite Msp Admin

Invite MSP Admin

**Note**: An email will also be sent to the user with a link to https://manage.mist.com/verify/invite?token=:token

```go
InviteMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

[`models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("user@example.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.PrivilegeSelf{
        models.PrivilegeSelf{
            MspId:        models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            MspName:      models.NewOptional(models.ToPointer("string")),
            Name:         models.ToPointer("string"),
            OrgId:        models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrgName:      models.ToPointer("string"),
            OrggroupIds:  []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
            Role:         models.PrivilegeSelfRoleEnum("admin"),
            Scope:        models.PrivilegeSelfScopeEnum("org"),
            SiteId:       models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            SitegroupIds: []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
    },
}

apiResponse, err := mSPsAdmins.InviteMspAdmin(ctx, mspId, &body)
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


# List Msp Admins

Get List of MSP Admins

```go
ListMspAdmins(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsAdmins.ListMspAdmins(ctx, mspId)
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
    "admin_id": "456b7016-a916-a4b1-78dd-72b947c152b7",
    "email": "jsmith@mycorp.org",
    "first_name": "Joe",
    "last_name": "Smith",
    "privileges": [
      {
        "role": "admin",
        "scope": "msp"
      },
      {
        "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
        "role": "admin",
        "scope": "org"
      },
      {
        "orggroup_ids": [
          "507f1bab-13ba-73e2-f291-2bcb8d1362b0"
        ],
        "role": "read",
        "scope": "orggroup"
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


# Revoke Msp Admin

This removes all privileges this admin has against the MSP. This goes deep all the way to the sites

```go
RevokeMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsAdmins.RevokeMspAdmin(ctx, mspId, adminId)
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


# Uninvite Msp Admin

Delete admin invite

```go
UninviteMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    inviteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsAdmins.UninviteMspAdmin(ctx, mspId, inviteId)
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


# Update Msp Admin

Update MSP Admin

```go
UpdateMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

[`models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("jsnow@abc.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.PrivilegeSelf{
        models.PrivilegeSelf{
            OrgId:        models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            Role:         models.PrivilegeSelfRoleEnum("admin"),
            Scope:        models.PrivilegeSelfScopeEnum("org"),
            SitegroupIds: []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
        models.PrivilegeSelf{
            Role:         models.PrivilegeSelfRoleEnum("admin"),
            Scope:        models.PrivilegeSelfScopeEnum("site"),
            SiteId:       models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
        },
    },
}

apiResponse, err := mSPsAdmins.UpdateMspAdmin(ctx, mspId, adminId, &body)
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


# Update Msp Admin Invite

Update MSP admin invite

```go
UpdateMspAdminInvite(
    ctx context.Context,
    mspId uuid.UUID,
    inviteId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

[`models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("user@example.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.PrivilegeSelf{
        models.PrivilegeSelf{
            MspId:        models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            MspName:      models.NewOptional(models.ToPointer("string")),
            Name:         models.ToPointer("string"),
            OrgId:        models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrgName:      models.ToPointer("string"),
            OrggroupIds:  []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
            Role:         models.PrivilegeSelfRoleEnum("admin"),
            Scope:        models.PrivilegeSelfScopeEnum("org"),
            SiteId:       models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            SitegroupIds: []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
    },
}

apiResponse, err := mSPsAdmins.UpdateMspAdminInvite(ctx, mspId, inviteId, &body)
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

