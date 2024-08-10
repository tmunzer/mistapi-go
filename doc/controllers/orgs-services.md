# Orgs Services

```go
orgsServices := client.OrgsServices()
```

## Class Name

`OrgsServices`

## Methods

* [Create Org Service](../../doc/controllers/orgs-services.md#create-org-service)
* [Delete Org Service](../../doc/controllers/orgs-services.md#delete-org-service)
* [Get Org Service](../../doc/controllers/orgs-services.md#get-org-service)
* [List Org Services](../../doc/controllers/orgs-services.md#list-org-services)
* [Update Org Service](../../doc/controllers/orgs-services.md#update-org-service)


# Create Org Service

Create getOrgServices Service

```go
CreateOrgService(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Service) (
    models.ApiResponse[models.Service],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Service`](../../doc/models/service.md) | Body, Optional | - |

## Response Type

[`models.Service`](../../doc/models/service.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Service{
    Name:                          models.ToPointer("string"),
    Specs:                         []models.ServiceSpec{
        models.ServiceSpec{
            Protocol:  models.ToPointer("any"),
        },
    },
    Type:                          models.ToPointer(models.ServiceTypeEnum("custom")),
}

apiResponse, err := orgsServices.CreateOrgService(ctx, orgId, &body)
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
  "addresses": [
    "string"
  ],
  "app_categories": [
    "string"
  ],
  "apps": [
    "string"
  ],
  "created_time": 0,
  "dscp": 0,
  "failover_policy": "revertable",
  "hostnames": [
    "string"
  ],
  "id": "497f6eca-6276-5004-bfeb-53cbbbba6f16",
  "max_jitter": 0,
  "max_latency": 0,
  "max_loss": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "sle_enabled": false,
  "specs": [
    {
      "port_range": "0",
      "protocol": "any"
    }
  ],
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "vpn_name": "addresses"
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


# Delete Org Service

delete Org Service

```go
DeleteOrgService(
    ctx context.Context,
    orgId uuid.UUID,
    serviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `serviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

serviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsServices.DeleteOrgService(ctx, orgId, serviceId)
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


# Get Org Service

Get Org Service

```go
GetOrgService(
    ctx context.Context,
    orgId uuid.UUID,
    serviceId uuid.UUID) (
    models.ApiResponse[models.Service],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `serviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Service`](../../doc/models/service.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

serviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsServices.GetOrgService(ctx, orgId, serviceId)
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
  "addresses": [
    "string"
  ],
  "app_categories": [
    "string"
  ],
  "apps": [
    "string"
  ],
  "created_time": 0,
  "dscp": 0,
  "failover_policy": "revertable",
  "hostnames": [
    "string"
  ],
  "id": "497f6eca-6276-5004-bfeb-53cbbbba6f16",
  "max_jitter": 0,
  "max_latency": 0,
  "max_loss": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "sle_enabled": false,
  "specs": [
    {
      "port_range": "0",
      "protocol": "any"
    }
  ],
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "vpn_name": "addresses"
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


# List Org Services

Get List of Org Services

```go
ListOrgServices(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Service],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Service`](../../doc/models/service.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsServices.ListOrgServices(ctx, orgId, &page, &limit)
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
    "addresses": [
      "string"
    ],
    "apps": [
      "string"
    ],
    "dscp": 8,
    "hostnames": [
      "string"
    ],
    "max_jitter": 0,
    "max_latency": 0,
    "max_loss": 0,
    "name": "string",
    "specs": [
      {
        "port_range": "0",
        "protocol": "any"
      }
    ],
    "traffic_class": "best_effort",
    "traffic_type": "default",
    "type": "custom"
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


# Update Org Service

update Org Service

```go
UpdateOrgService(
    ctx context.Context,
    orgId uuid.UUID,
    serviceId uuid.UUID,
    body *models.Service) (
    models.ApiResponse[models.Service],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `serviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Service`](../../doc/models/service.md) | Body, Optional | - |

## Response Type

[`models.Service`](../../doc/models/service.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

serviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Service{
    Addresses:                     []string{
        "string",
    },
    AppCategories:                 []string{
        "string",
    },
    Apps:                          []string{
        "string",
    },
    Dscp:                          models.ToPointer(models.ServiceDscpContainer.FromNumber(0)),
    FailoverPolicy:                models.ToPointer(models.ServiceFailoverPolicyEnum("revertable")),
    Hostnames:                     []string{
        "string",
    },
    MaxJitter:                     models.ToPointer(models.ServiceMaxJitterContainer.FromNumber(0)),
    MaxLatency:                    models.ToPointer(models.ServiceMaxLatencyContainer.FromNumber(0)),
    MaxLoss:                       models.ToPointer(models.ServiceMaxLossContainer.FromNumber(0)),
    Name:                          models.ToPointer("string"),
    SleEnabled:                    models.ToPointer(false),
    Specs:                         []models.ServiceSpec{
        models.ServiceSpec{
            PortRange: models.ToPointer("0"),
            Protocol:  models.ToPointer("any"),
        },
    },
    TrafficClass:                  models.ToPointer(models.ServiceTrafficClassEnum("best_effort")),
    TrafficType:                   models.ToPointer("data_best_effort"),
    Type:                          models.ToPointer(models.ServiceTypeEnum("custom")),
}

apiResponse, err := orgsServices.UpdateOrgService(ctx, orgId, serviceId, &body)
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
  "addresses": [
    "string"
  ],
  "app_categories": [
    "string"
  ],
  "apps": [
    "string"
  ],
  "created_time": 0,
  "dscp": 0,
  "failover_policy": "revertable",
  "hostnames": [
    "string"
  ],
  "id": "497f6eca-6276-5004-bfeb-53cbbbba6f16",
  "max_jitter": 0,
  "max_latency": 0,
  "max_loss": 0,
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "sle_enabled": false,
  "specs": [
    {
      "port_range": "0",
      "protocol": "any"
    }
  ],
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "vpn_name": "addresses"
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

