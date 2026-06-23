# Orgs Psks

```go
orgsPsks := client.OrgsPsks()
```

## Class Name

`OrgsPsks`

## Methods

* [Create Org Psk](../../doc/controllers/orgs-psks.md#create-org-psk)
* [Delete Org Psk](../../doc/controllers/orgs-psks.md#delete-org-psk)
* [Delete Org Psk List](../../doc/controllers/orgs-psks.md#delete-org-psk-list)
* [Delete Org Psk Old Passphrase](../../doc/controllers/orgs-psks.md#delete-org-psk-old-passphrase)
* [Get Org Psk](../../doc/controllers/orgs-psks.md#get-org-psk)
* [Import Org Psks](../../doc/controllers/orgs-psks.md#import-org-psks)
* [List Org Psks](../../doc/controllers/orgs-psks.md#list-org-psks)
* [Update Org Multiple Psks](../../doc/controllers/orgs-psks.md#update-org-multiple-psks)
* [Update Org Psk](../../doc/controllers/orgs-psks.md#update-org-psk)


# Create Org Psk

Create an organization personal PSK for WLAN access, including SSID, passphrase, usage mode, optional client MAC binding, role, VLAN, and expiration settings.

When `usage`==`macs`, corresponding "macs" field will hold a list consisting of client MAC addresses (["xx:xx:xx:xx:xx",...]) or mac patterns(["xx:xx:*","xx*",...]) or both (["xx:xx:xx:xx:xx:xx", "xx:*", ...]). This list is capped at 5000

```go
CreateOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    upsert *bool,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upsert` | `*bool` | Query, Optional | If a key exists with the same `name`, replace it with the new one |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Psk{
    ExpireTime:             models.NewOptional(models.ToPointer(1614990263)),
    Macs:                   []string{
        "112233abcedf",
        "aabbcc*",
    },
    MaxUsage:               models.ToPointer(0),
    Name:                   "name6",
    NotifyExpiry:           models.ToPointer(false),
    Passphrase:             "passphrase6",
    Ssid:                   "ssid6",
    Usage:                  models.ToPointer(models.PskUsageEnum_MULTI),
}

apiResponse, err := orgsPsks.CreateOrgPsk(ctx, orgId, nil, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Psk

Delete an organization personal PSK by PSK ID so clients can no longer authenticate with that key.

```go
DeleteOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsPsks.DeleteOrgPsk(ctx, orgId, pskId)
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


# Delete Org Psk List

Delete one or more organization PSKs by ID.

The request accepts a single PSK ID string or a list of PSK ID strings.
**Warning**: If no PSK IDs are provided in the request, all organization PSKs will be deleted and clients will no longer be able to authenticate to any SSID with a PSK.

```go
DeleteOrgPskList(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PskIdList) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PskIdList`](../../doc/models/psk-id-list.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PskIdList{
    PskIds:               []uuid.UUID{
        uuid.MustParse("0039c16c-ca87-4d3f-bb94-b97c58199f18"),
        uuid.MustParse("6562cc8e-5893-418a-acaa-4d7c1af8084f"),
    },
}

resp, err := orgsPsks.DeleteOrgPskList(ctx, orgId, &body)
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


# Delete Org Psk Old Passphrase

Remove the stored `old_passphrase` from a PSK after rotation.
If successful, the response returns the PSK with `old_passphrase` removed.

```go
DeleteOrgPskOldPassphrase(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    models.ApiResponse[models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPsks.DeleteOrgPskOldPassphrase(ctx, orgId, pskId)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Psk

Retrieve details for an organization personal PSK, including SSID, usage mode, MAC binding, role, VLAN, expiration, and notification settings.

```go
GetOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    models.ApiResponse[models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPsks.GetOrgPsk(ctx, orgId, pskId)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Import Org Psks

Import organization PSKs from a CSV file or JSON payload.

## CSV File Format

```
PSK Import CSV File Format:
name,ssid,passphrase,usage,vlan_id,mac,max_usage,role,expire_time,notify_expiry,expiry_notification_time,notify_on_create_or_edit,email
Common,warehouse,foryoureyesonly,single,35,a31425f31278,0,student,1618594236
Justin,reception,visible,multi,1002,200,teacher,1618594236
Common2,ssid,1245678-xx,single,35,a31425f31278,0,student,1618594236,true,7,true,admin@test.com
```

```go
ImportOrgPsks(
    ctx context.Context,
    orgId uuid.UUID,
    file *models.FileWrapper) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | Uploaded CSV or binary file containing PSKs to import |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPsks.ImportOrgPsks(ctx, orgId, nil)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "secretpsk",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
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


# List Org Psks

List organization personal PSKs for WLAN access, optionally filtering by name, SSID, or role.

```go
ListOrgPsks(
    ctx context.Context,
    orgId uuid.UUID,
    name *string,
    ssid *string,
    role *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `name` | `*string` | Query, Optional | Filter results by name. Accepts multiple comma-separated values. |
| `ssid` | `*string` | Query, Optional | Filter results by SSID. Accepts multiple comma-separated values. |
| `role` | `*string` | Query, Optional | Filter PSK results by role. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

name := "psk-a,psk-b"

ssid := "Corp,Guest"

role := "employee,contractor"

limit := 100

page := 1

apiResponse, err := orgsPsks.ListOrgPsks(ctx, orgId, &name, &ssid, &role, &limit, &page)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "secretpsk",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
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


# Update Org Multiple Psks

Update multiple organization personal PSKs in one request, including passphrase, usage mode, role, VLAN, expiration, and notification settings.

```go
UpdateOrgMultiplePsks(
    ctx context.Context,
    orgId uuid.UUID,
    body []models.Psk) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.Psk`](../../doc/models/psk.md) | Body, Optional | **Constraints**: *Unique Items Required* |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []models.Psk{
    models.Psk{
        ExpireTime:             models.NewOptional(models.ToPointer(1614990263)),
        Mac:                    models.ToPointer("string"),
        MaxUsage:               models.ToPointer(0),
        Name:                   "string",
        Passphrase:             "secretpsk",
        Ssid:                   "string",
        Usage:                  models.ToPointer(models.PskUsageEnum_MULTI),
        VlanId:                 models.ToPointer(models.PskVlanIdContainer.FromNumber(10)),
    },
}

apiResponse, err := orgsPsks.UpdateOrgMultiplePsks(ctx, orgId, body)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "secretpsk",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
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


# Update Org Psk

Update an organization personal PSK, including passphrase, usage mode, MAC binding, role, VLAN, expiration, and notification settings.

```go
UpdateOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Psk{
    ExpireTime:             models.NewOptional(models.ToPointer(1614990263)),
    Macs:                   []string{
        "112233abcedf",
        "aabbcc*",
    },
    MaxUsage:               models.ToPointer(0),
    Name:                   "name6",
    NotifyExpiry:           models.ToPointer(false),
    Passphrase:             "passphrase6",
    Ssid:                   "ssid6",
    Usage:                  models.ToPointer(models.PskUsageEnum_MULTI),
}

apiResponse, err := orgsPsks.UpdateOrgPsk(ctx, orgId, pskId, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

