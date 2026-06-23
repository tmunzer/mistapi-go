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

Return administrator details and privilege assignments for one administrator under this MSP.

```go
GetMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID) (
    models.ApiResponse[models.Admin],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsAdmins.GetMspAdmin(ctx, mspId, adminId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Invite Msp Admin

Send an MSP administrator invitation with the requested identity fields and privilege assignments.

**Note**: An email will also be sent to the user with a link to https://manage.mist.com/verify/invite?token=:token

```go
InviteMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("user@example.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.AdminPrivilege{
        models.AdminPrivilege{
            MspId:                models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrgId:                models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrggroupIds:          []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
            Role:                 models.AdminPrivilegeRoleEnum_ADMIN,
            Scope:                models.AdminPrivilegeScopeEnum_ORG,
            SiteId:               models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
    },
}

apiResponse, err := mSPsAdmins.InviteMspAdmin(ctx, mspId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Msp Admins

List administrators that have privileges in this MSP hierarchy, including MSP, organization, organization group, site, or site group scopes.

```go
ListMspAdmins(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Admin],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsAdmins.ListMspAdmins(ctx, mspId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Revoke Msp Admin

Remove all privileges this administrator has through the MSP hierarchy, including inherited organization and site access. This does not delete the administrator account.

```go
RevokeMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsAdmins.RevokeMspAdmin(ctx, mspId, adminId)
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


# Uninvite Msp Admin

Cancel a pending MSP administrator invitation before the invitee accepts it.

```go
UninviteMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    inviteId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsAdmins.UninviteMspAdmin(ctx, mspId, inviteId)
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


# Update Msp Admin

Update identity fields and privilege assignments for an existing administrator under this MSP.

```go
UpdateMspAdmin(
    ctx context.Context,
    mspId uuid.UUID,
    adminId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("jsnow@abc.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.AdminPrivilege{
        models.AdminPrivilege{
            OrgId:                models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            Role:                 models.AdminPrivilegeRoleEnum_ADMIN,
            Scope:                models.AdminPrivilegeScopeEnum_ORG,
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
        models.AdminPrivilege{
            Role:                 models.AdminPrivilegeRoleEnum_ADMIN,
            Scope:                models.AdminPrivilegeScopeEnum_SITE,
            SiteId:               models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
        },
    },
}

apiResponse, err := mSPsAdmins.UpdateMspAdmin(ctx, mspId, adminId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Msp Admin Invite

Update the identity fields or privilege assignments on a pending MSP administrator invitation.

```go
UpdateMspAdminInvite(
    ctx context.Context,
    mspId uuid.UUID,
    inviteId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Admin](../../doc/models/admin.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:                models.ToPointer("user@example.com"),
    FirstName:            models.ToPointer("string"),
    LastName:             models.ToPointer("string"),
    Privileges:           []models.AdminPrivilege{
        models.AdminPrivilege{
            MspId:                models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrgId:                models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            OrggroupIds:          []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
            Role:                 models.AdminPrivilegeRoleEnum_ADMIN,
            Scope:                models.AdminPrivilegeScopeEnum_ORG,
            SiteId:               models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")),
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
            },
        },
    },
}

apiResponse, err := mSPsAdmins.UpdateMspAdminInvite(ctx, mspId, inviteId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

