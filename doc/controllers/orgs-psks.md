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

Create Org PSK

When `usage`==`macs`, corresponding "macs" field will hold a list consisting of client mac addresses (["xx:xx:xx:xx:xx",...]) or mac patterns(["xx:xx:*","xx*",...]) or both (["xx:xx:xx:xx:xx:xx", "xx:*", ...]). This list is capped at 5000

```go
CreateOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    upsert *bool,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `upsert` | `*bool` | Query, Optional | If a key exists with the same `name`, replace it with the new one |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

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


# Delete Org Psk

Delete Org PSK

```go
DeleteOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsPsks.DeleteOrgPsk(ctx, orgId, pskId)
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


# Delete Org Psk List

Delete Org PSK List

Delete list of psks on the org. This API accepts single string or list of strings

```go
DeleteOrgPskList(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PskIdList) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PskIdList`](../../doc/models/psk-id-list.md) | Body, Optional | - |

## Response Type

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


# Delete Org Psk Old Passphrase

Delete `old_passphrase` from PSK.
If successful, response is same as GET, returns the PSK with `old_passphrase` removed.

```go
DeleteOrgPskOldPassphrase(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPsks.DeleteOrgPskOldPassphrase(ctx, orgId, pskId)
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


# Get Org Psk

Get Org PSK Details

```go
GetOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPsks.GetOrgPsk(ctx, orgId, pskId)
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


# Import Org Psks

Import PSK from CSV file or JSON

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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := orgsPsks.ImportOrgPsks(ctx, orgId, nil)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Psks

Get List of Org Psks

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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `name` | `*string` | Query, Optional | - |
| `ssid` | `*string` | Query, Optional | - |
| `role` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Psk](../../doc/models/psk.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

name := "psk_name"





limit := 100

page := 1

apiResponse, err := orgsPsks.ListOrgPsks(ctx, orgId, &name, nil, nil, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Multiple Psks

Update Multiple PSKs

```go
UpdateOrgMultiplePsks(
    ctx context.Context,
    orgId uuid.UUID,
    body []models.Psk) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.Psk`](../../doc/models/psk.md) | Body, Optional | **Constraints**: *Unique Items Required* |

## Response Type

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Psk

Update Org PSK

```go
UpdateOrgPsk(
    ctx context.Context,
    orgId uuid.UUID,
    pskId uuid.UUID,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

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

