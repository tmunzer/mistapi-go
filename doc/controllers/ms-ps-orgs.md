# MS Ps Orgs

```go
mSPsOrgs := client.MSPsOrgs()
```

## Class Name

`MSPsOrgs`

## Methods

* [Create Msp Org](../../doc/controllers/ms-ps-orgs.md#create-msp-org)
* [Delete Msp Org](../../doc/controllers/ms-ps-orgs.md#delete-msp-org)
* [Get Msp Org](../../doc/controllers/ms-ps-orgs.md#get-msp-org)
* [List Msp Org Stats](../../doc/controllers/ms-ps-orgs.md#list-msp-org-stats)
* [List Msp Orgs](../../doc/controllers/ms-ps-orgs.md#list-msp-orgs)
* [Manage Msp Orgs](../../doc/controllers/ms-ps-orgs.md#manage-msp-orgs)
* [Search Msp Orgs](../../doc/controllers/ms-ps-orgs.md#search-msp-orgs)
* [Update Msp Org](../../doc/controllers/ms-ps-orgs.md#update-msp-org)


# Create Msp Org

Create an Org under MSP

```go
CreateMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Org`](../../doc/models/org.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Org](../../doc/models/org.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Org{
    AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    AllowMist:            models.ToPointer(true),
    Name:                 "string",
    OrggroupIds:          []uuid.UUID{
        uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
    },
    SessionExpiry:        models.ToPointer(10),
}

apiResponse, err := mSPsOrgs.CreateMspOrg(ctx, mspId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Msp Org

Delete MSP Org

```go
DeleteMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsOrgs.DeleteMspOrg(ctx, mspId, orgId)
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


# Get Msp Org

Get MSP Org Details

```go
GetMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Org](../../doc/models/org.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsOrgs.GetMspOrg(ctx, mspId, orgId)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Msp Org Stats

Get List of MSP Orgs Stats

```go
ListMspOrgStats(
    ctx context.Context,
    mspId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsOrg](../../doc/models/stats-org.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := mSPsOrgs.ListMspOrgStats(ctx, mspId, &limit, &page)
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


# List Msp Orgs

Get List of MSP Orgs

```go
ListMspOrgs(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Org](../../doc/models/org.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsOrgs.ListMspOrgs(ctx, mspId)
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


# Manage Msp Orgs

Assign or Unassign Orgs to an MSP account

```go
ManageMspOrgs(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.MspOrgChange) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MspOrgChange`](../../doc/models/msp-org-change.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MspOrgChange{
    Op:                   models.MspOrgChangeOperationEnum_ASSIGN,
    OrgIds:               []string{
        "2b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    },
}

resp, err := mSPsOrgs.ManageMspOrgs(ctx, mspId, &body)
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


# Search Msp Orgs

Search Org in MSP

```go
SearchMspOrgs(
    ctx context.Context,
    mspId uuid.UUID,
    name *string,
    orgId *uuid.UUID,
    subInsufficient *bool,
    trialEnabled *bool,
    usageTypes []string,
    limit *int) (
    models.ApiResponse[models.ResponseOrgSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `name` | `*string` | Query, Optional | - |
| `orgId` | `*uuid.UUID` | Query, Optional | Org id |
| `subInsufficient` | `*bool` | Query, Optional | If this org has sufficient subscription |
| `trialEnabled` | `*bool` | Query, Optional | If this org is under trial period |
| `usageTypes` | `[]string` | Query, Optional | List of types that enabled by usage<br><br>**Constraints**: *Unique Items Required* |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgSearch](../../doc/models/response-org-search.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orgId := uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1")

subInsufficient := false

trialEnabled := false

limit := 100

apiResponse, err := mSPsOrgs.SearchMspOrgs(ctx, mspId, nil, &orgId, &subInsufficient, &trialEnabled, nil, &limit)
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
  "end": 1614383378.4365287,
  "limit": 10,
  "results": [
    {
      "msp_id": "d287e62f-0000-0000-0000-f2b9ba0a531f",
      "name": "Test Org",
      "num_aps": 9,
      "num_sites": 5,
      "num_switches": 1,
      "num_unassigned_aps": 1,
      "org_id": "bb1a8bf6-0000-0000-0000-8053a663cf65",
      "sub_ana_required": 9,
      "sub_ast_entitled": 5,
      "sub_ast_required": 3,
      "sub_eng_required": 3,
      "sub_ex12_required": 1,
      "sub_insufficient": true,
      "sub_man_required": 9,
      "sub_vna_entitled": 1,
      "timestamp": 1614322563.513937,
      "trial_enabled": false,
      "usage_types": [
        "sub_eng"
      ]
    },
    {
      "msp_id": "d287e62f-0000-0000-0000-f2b9ba0a531f",
      "name": "Rogue Test1",
      "num_aps": 1,
      "num_sites": 1,
      "org_id": "0fb81690-0000-0000-0000-9596d1d1534f",
      "sub_ana_entitled": 1,
      "sub_ana_required": 1,
      "sub_insufficient": false,
      "sub_man_entitled": 1,
      "sub_man_required": 1,
      "timestamp": 1614309876.500955
    }
  ],
  "start": 1613778578.4365668,
  "total": 2
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


# Update Msp Org

Update MSP Org

```go
UpdateMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Org`](../../doc/models/org.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Org](../../doc/models/org.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Org{
    AllowMist:            models.ToPointer(true),
    Name:                 "Org",
    SessionExpiry:        models.ToPointer(1440),
}

apiResponse, err := mSPsOrgs.UpdateMspOrg(ctx, mspId, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

