# Sites Guests

```go
sitesGuests := client.SitesGuests()
```

## Class Name

`SitesGuests`

## Methods

* [Count Site Guest Authorizations](../../doc/controllers/sites-guests.md#count-site-guest-authorizations)
* [Delete Site Guest Authorization](../../doc/controllers/sites-guests.md#delete-site-guest-authorization)
* [Get Site Guest Authorization](../../doc/controllers/sites-guests.md#get-site-guest-authorization)
* [List Site All Guest Authorizations](../../doc/controllers/sites-guests.md#list-site-all-guest-authorizations)
* [List Site All Guest Authorizations Derived](../../doc/controllers/sites-guests.md#list-site-all-guest-authorizations-derived)
* [Search Site Guest Authorization](../../doc/controllers/sites-guests.md#search-site-guest-authorization)
* [Update Site Guest Authorization](../../doc/controllers/sites-guests.md#update-site-guest-authorization)


# Count Site Guest Authorizations

Count Authorized Guest

```go
CountSiteGuestAuthorizations(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteGuestsCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteGuestsCountDistinctEnum`](../../doc/models/site-guests-count-distinct-enum.md) | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteGuestsCountDistinctEnum("auth_method")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesGuests.CountSiteGuestAuthorizations(ctx, siteId, &distinct, &page, &limit, nil, nil, &duration)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
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


# Delete Site Guest Authorization

Delete Guest Authorization

```go
DeleteSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

resp, err := sitesGuests.DeleteSiteGuestAuthorization(ctx, siteId, guestMac)
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


# Get Site Guest Authorization

Get Guest Authorization

```go
GetSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string) (
    models.ApiResponse[models.Guest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | - |

## Response Type

[`models.Guest`](../../doc/models/guest.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

apiResponse, err := sitesGuests.GetSiteGuestAuthorization(ctx, siteId, guestMac)
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
  "authorized": true,
  "authorized_expiring_time": 0,
  "authorized_time": 0,
  "company": "string",
  "email": "user@example.com",
  "field1": "string",
  "field2": "string",
  "field3": "string",
  "field4": "string",
  "mac": "string",
  "minutes": 0,
  "name": "string"
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


# List Site All Guest Authorizations

Get List of Site Guest Authorizations

```go
ListSiteAllGuestAuthorizations(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string) (
    models.ApiResponse[[]models.Guest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `*string` | Query, Optional | UUID of single or multiple (Comma separated) WLAN under Site `site_id` (to filter by WLAN) |

## Response Type

[`[]models.Guest`](../../doc/models/guest.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesGuests.ListSiteAllGuestAuthorizations(ctx, siteId, nil)
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
    "authorized": true,
    "authorized_expiring_time": 0,
    "authorized_time": 0,
    "company": "string",
    "email": "user@example.com",
    "field1": "string",
    "field2": "string",
    "field3": "string",
    "field4": "string",
    "mac": "string",
    "minutes": 0,
    "name": "string"
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


# List Site All Guest Authorizations Derived

Get List of Site Guest Authorizations

```go
ListSiteAllGuestAuthorizationsDerived(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string,
    crossSite *bool) (
    models.ApiResponse[[]models.Guest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `*string` | Query, Optional | UUID of single or multiple (Comma separated) WLAN under Site `site_id` (to filter by WLAN) |
| `crossSite` | `*bool` | Query, Optional | whether to get org level guests, default is false i.e get site level guests |

## Response Type

[`[]models.Guest`](../../doc/models/guest.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



crossSite := false

apiResponse, err := sitesGuests.ListSiteAllGuestAuthorizationsDerived(ctx, siteId, nil, &crossSite)
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
    "authorized": true,
    "authorized_expiring_time": 0,
    "authorized_time": 0,
    "company": "string",
    "email": "user@example.com",
    "field1": "string",
    "field2": "string",
    "field3": "string",
    "field4": "string",
    "mac": "string",
    "minutes": 0,
    "name": "string"
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


# Search Site Guest Authorization

Search Authorized Guest

```go
SearchSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string,
    authMethod *string,
    ssid *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseGuestSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `*string` | Query, Optional | - |
| `authMethod` | `*string` | Query, Optional | - |
| `ssid` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseGuestSearch`](../../doc/models/response-guest-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := "00000000-0000-0000-0000-000000000000"





limit := 100





duration := "10m"

apiResponse, err := sitesGuests.SearchSiteGuestAuthorization(ctx, siteId, &wlanId, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1531862583,
  "limit": 2,
  "next": "/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/guests/search?wlan_id=88ffe630-95b8-11e8-b294-346895ed1b7d&end=1531855849.000&limit=2&start=1531776183.0",
  "results": [
    {
      "ap": "5c5b350e0001",
      "auth_method": "passphrase",
      "authorized_expiring_time": 1531810258.186273,
      "authorized_time": 1531782218,
      "company": "mistsystems",
      "email": "user@mistsys.com",
      "name": "john",
      "ssid": "openNet",
      "timestamp": 1531782218
    },
    {
      "ap": "5c5b350e0001",
      "auth_method": "facebook",
      "authorized_expiring_time": 1531810821.145,
      "authorized_time": 1531782632,
      "company": "xyz inc.",
      "email": "cool_user@yahoo.com",
      "name": "John White",
      "ssid": "openNet",
      "timestamp": 1531782632
    }
  ],
  "start": 1531776183,
  "total": 14
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


# Update Site Guest Authorization

Update Guest Authorization

```go
UpdateSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string,
    body *models.Guest) (
    models.ApiResponse[models.Guest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | - |
| `body` | [`*models.Guest`](../../doc/models/guest.md) | Body, Optional | Request Body |

## Response Type

[`models.Guest`](../../doc/models/guest.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

body := models.Guest{
    AuthorizedExpiringTime: models.ToPointer(float64(1480704955)),
    AuthorizedTime:         models.ToPointer(float64(1480704355)),
    Company:                models.ToPointer("abc"),
    Email:                  models.ToPointer("john@abc.com"),
    Name:                   models.ToPointer("John Smith"),
    Ssid:                   models.ToPointer("Guest-SSID"),
    WlanId:                 models.ToPointer(uuid.MustParse("6748cfa6-4e12-11e6-9188-0242ac110007")),
}

apiResponse, err := sitesGuests.UpdateSiteGuestAuthorization(ctx, siteId, guestMac, &body)
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
  "authorized": true,
  "authorized_expiring_time": 0,
  "authorized_time": 0,
  "company": "string",
  "email": "user@example.com",
  "field1": "string",
  "field2": "string",
  "field3": "string",
  "field4": "string",
  "mac": "string",
  "minutes": 0,
  "name": "string"
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
