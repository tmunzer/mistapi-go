# Sites

A site represents a project, a deployment. For MSP, it can be as small as a coffee shop or a five-star 600-room hotel. A site contains a set of Maps, Wlans, Policies, Zones.

```go
sites := client.Sites()
```

## Class Name

`Sites`

## Methods

* [Delete Site](../../doc/controllers/sites.md#delete-site)
* [Get Site Info](../../doc/controllers/sites.md#get-site-info)
* [Update Site Info](../../doc/controllers/sites.md#update-site-info)


# Delete Site

Delete Site

```go
DeleteSite(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sites.DeleteSite(ctx, siteId)
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


# Get Site Info

Provides information about the site, including its name, address,
timezone, and associated templates. This endpoint is useful for retrieving
the current configuration and details of a specific site.

```go
GetSiteInfo(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.Site],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Site](../../doc/models/site.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sites.GetSiteInfo(ctx, siteId)
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
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "apporttemplate_id": "string",
  "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
  "country_code": "US",
  "created_time": 0,
  "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
  "id": "497f6eca-6276-5005-bfeb-53cbbbba6f17",
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946
  },
  "modified_time": 0,
  "name": "Mist Office",
  "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
  "notes": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
  "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
  "sitegroup_ids": [
    "497f6eca-6276-5006-bfeb-53cbbbba6f18"
  ],
  "timezone": "America/Los_Angeles"
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


# Update Site Info

Updates the configuration and metadata for an existing site.

This endpoint allows modification of site properties including location details (address, coordinates, timezone), template associations (alarm, network, RF, security policy templates), site group memberships, and general information (name, notes).

All fields are optional and only provided fields will be updated.

```go
UpdateSiteInfo(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Site) (
    models.ApiResponse[models.Site],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Site`](../../doc/models/site.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Site](../../doc/models/site.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Site{
    Address:              models.NewOptional(models.ToPointer("string")),
    AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    CountryCode:          models.ToPointer("string"),
    Latlng:               models.ToPointer(models.LatLng{
        Lat:                  float64(0),
        Lng:                  float64(0),
    }),
    Name:                 "string",
    NetworktemplateId:    models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    Notes:                models.NewOptional(models.ToPointer("string")),
    RftemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    SecpolicyId:          models.NewOptional(models.ToPointer(uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"))),
    SitegroupIds:         []uuid.UUID{
        uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
    },
    Timezone:             models.ToPointer("string"),
}

apiResponse, err := sites.UpdateSiteInfo(ctx, siteId, &body)
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
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "apporttemplate_id": "string",
  "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
  "country_code": "US",
  "created_time": 0,
  "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
  "id": "497f6eca-6276-5005-bfeb-53cbbbba6f17",
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946
  },
  "modified_time": 0,
  "name": "Mist Office",
  "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
  "notes": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
  "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
  "sitegroup_ids": [
    "497f6eca-6276-5006-bfeb-53cbbbba6f18"
  ],
  "timezone": "America/Los_Angeles"
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

