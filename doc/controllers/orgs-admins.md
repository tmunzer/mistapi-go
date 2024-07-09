# Orgs Admins

```go
orgsAdmins := client.OrgsAdmins()
```

## Class Name

`OrgsAdmins`

## Methods

* [Invite Org Admin](../../doc/controllers/orgs-admins.md#invite-org-admin)
* [List Org Admins](../../doc/controllers/orgs-admins.md#list-org-admins)
* [Revoke Org Admin](../../doc/controllers/orgs-admins.md#revoke-org-admin)
* [Uninvite Org Admin](../../doc/controllers/orgs-admins.md#uninvite-org-admin)
* [Update Org Admin](../../doc/controllers/orgs-admins.md#update-org-admin)
* [Update Org Admin Invite](../../doc/controllers/orgs-admins.md#update-org-admin-invite)


# Invite Org Admin

If the request is successful, an email will also be sent to the user with a link to `https://manage.mist.com/verify/invite?token=:token&expire=1459632743&org=OrgName`

```go
InviteOrgAdmin(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Admin) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:             "user@example.com",
    FirstName:         "string",
    LastName:          "string",
    Privileges:        []models.PrivilegeSelf{
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

resp, err := orgsAdmins.InviteOrgAdmin(ctx, orgId, &body)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Admins

Get List of people who can manage the Site/Org under the Org

```go
ListOrgAdmins(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsAdmins.ListOrgAdmins(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Revoke Org Admin

This removes all privileges this admin has against the org

```go
RevokeOrgAdmin(
    ctx context.Context,
    orgId uuid.UUID,
    adminId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAdmins.RevokeOrgAdmin(ctx, orgId, adminId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Uninvite Org Admin

Delete Admin Invite

```go
UninviteOrgAdmin(
    ctx context.Context,
    orgId uuid.UUID,
    inviteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsAdmins.UninviteOrgAdmin(ctx, orgId, inviteId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Admin

Invite Org Admin

```go
UpdateOrgAdmin(
    ctx context.Context,
    orgId uuid.UUID,
    adminId uuid.UUID,
    body *models.Admin) (
    models.ApiResponse[models.Admin],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `adminId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

[`models.Admin`](../../doc/models/admin.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adminId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:             "jsnow@abc.com",
    ExpireTime:        models.ToPointer(0),
    FirstName:         "John",
    Hours:             models.ToPointer(24),
    LastName:          "Sno",
    Phone:             models.ToPointer("string"),
    Phone2:            models.ToPointer("string"),
    Privileges:        []models.PrivilegeSelf{
        models.PrivilegeSelf{
            MspId:        models.ToPointer(uuid.MustParse("c0cf23fc-d82f-4219-988c-82fb61d8c875")),
            Name:         models.ToPointer("string"),
            OrgId:        models.ToPointer(uuid.MustParse("a40f5d1f-d889-42e9-94ea-b9b33585fc6b")),
            OrgName:      models.ToPointer("string"),
            OrggroupIds:  []uuid.UUID{
                uuid.MustParse("497f6eca-6276-4993-bfeb-53d5bbba6f08"),
            },
            Role:         models.PrivilegeSelfRoleEnum("admin"),
            Scope:        models.PrivilegeSelfScopeEnum("org"),
            SiteId:       models.ToPointer(uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")),
            SitegroupIds: []uuid.UUID{
                uuid.MustParse("497f6eca-6276-4993-bfeb-53d6bbba6f08"),
            },
            Views:        models.ToPointer(models.PrivilegeSelfViewsEnum("switch_admin")),
        },
    },
}

apiResponse, err := orgsAdmins.UpdateOrgAdmin(ctx, orgId, adminId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Org Admin Invite

Update Admin Invite

```go
UpdateOrgAdminInvite(
    ctx context.Context,
    orgId uuid.UUID,
    inviteId uuid.UUID,
    body *models.Admin) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `inviteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Admin`](../../doc/models/admin.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

inviteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Admin{
    Email:             "user@example.com",
    FirstName:         "string",
    LastName:          "string",
    Privileges:        []models.PrivilegeSelf{
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

resp, err := orgsAdmins.UpdateOrgAdminInvite(ctx, orgId, inviteId, &body)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
