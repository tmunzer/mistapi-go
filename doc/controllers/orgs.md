# Orgs

An organization usually represents a customer - which has inventories, licenses. An Organization can contain multiple sites. A site usually represents a deployment at the same location (a campus, an office).

```go
orgs := client.Orgs()
```

## Class Name

`Orgs`

## Methods

* [Clone Org](../../doc/controllers/orgs.md#clone-org)
* [Create Org](../../doc/controllers/orgs.md#create-org)
* [Delete Org](../../doc/controllers/orgs.md#delete-org)
* [Get Org](../../doc/controllers/orgs.md#get-org)
* [Search Org Events](../../doc/controllers/orgs.md#search-org-events)
* [Update Org](../../doc/controllers/orgs.md#update-org)


# Clone Org

Create an Org by cloning from another one. Org Settings, Templates, Wxlan Tags, Wxlan Tunnels, Wxlan Rules, Org Wlans will be copied. Sites and Site Groups will not be copied, and therefore, the copied template will not be applied to any site/sitegroups.

```go
CloneOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NameString) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NameString`](../../doc/models/name-string.md) | Body, Optional | Request Body |

## Response Type

[`models.Org`](../../doc/models/org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NameString{
    Name: models.ToPointer("New Org"),
}

apiResponse, err := orgs.CloneOrg(ctx, orgId, &body)
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
  "alarmtemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "allow_mist": true,
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "orggroup_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "session_expiry": 1440
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


# Create Org

Org admin can invite people to manage the org. Furthermore, he can dictate the level of security those accounts are. The check is enforced when the invited admin tries to “accept” the invitation and every time the admin tries to login

```go
CreateOrg(
    ctx context.Context,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.Org`](../../doc/models/org.md) | Body, Optional | - |

## Response Type

[`models.Org`](../../doc/models/org.md)

## Example Usage

```go
ctx := context.Background()

body := models.Org{
    AlarmtemplateId: models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    AllowMist:       models.ToPointer(true),
    Name:            "string",
    SessionExpiry:   models.ToPointer(1440),
}

apiResponse, err := orgs.CreateOrg(ctx, &body)
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
  "alarmtemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "allow_mist": true,
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "orggroup_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "session_expiry": 1440
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


# Delete Org

Delete Org

```go
DeleteOrg(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgs.DeleteOrg(ctx, orgId)
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


# Get Org

Get Organization information

```go
GetOrg(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Org`](../../doc/models/org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgs.GetOrg(ctx, orgId)
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
  "alarmtemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "allow_mist": true,
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "orggroup_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "session_expiry": 1440
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


# Search Org Events

Search Org events

Supported Event Types:

- CRADLEPOINT_SYNC_FAILED
- ORG_CA_CERT_ADDED
- ORG_CA_CERT_REGENERATED

```go
SearchOrgEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsOrgsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | event type |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.ResponseEventsOrgsSearch`](../../doc/models/response-events-orgs-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







duration := "10m"

limit := 100

apiResponse, err := orgs.SearchOrgEvents(ctx, orgId, nil, nil, nil, &duration, &limit)
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


# Update Org

Update Org

```go
UpdateOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Org`](../../doc/models/org.md) | Body, Optional | Request Body |

## Response Type

[`models.Org`](../../doc/models/org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Org{
    AlarmtemplateId: models.NewOptional(models.ToPointer(uuid.MustParse("1984805d-2be2-4aec-a8d4-3ddf67fab0df"))),
    AllowMist:       models.ToPointer(true),
    Name:            "string",
    OrggroupIds:     []uuid.UUID{
    },
    SessionExpiry:   models.ToPointer(1440),
}

apiResponse, err := orgs.UpdateOrg(ctx, orgId, &body)
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
  "alarmtemplate_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "allow_mist": true,
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "orggroup_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "session_expiry": 1440
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

