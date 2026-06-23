# Orgs Guests

```go
orgsGuests := client.OrgsGuests()
```

## Class Name

`OrgsGuests`

## Methods

* [Count Org Guest Authorizations](../../doc/controllers/orgs-guests.md#count-org-guest-authorizations)
* [Delete Org Guest Authorization](../../doc/controllers/orgs-guests.md#delete-org-guest-authorization)
* [Get Org Guest Authorization](../../doc/controllers/orgs-guests.md#get-org-guest-authorization)
* [List Org Guest Authorizations](../../doc/controllers/orgs-guests.md#list-org-guest-authorizations)
* [Search Org Guest Authorization](../../doc/controllers/orgs-guests.md#search-org-guest-authorization)
* [Update Org Guest Authorization](../../doc/controllers/orgs-guests.md#update-org-guest-authorization)


# Count Org Guest Authorizations

Count organization guest authorization records, optionally grouped by `distinct` and filtered by time range.

```go
CountOrgGuestAuthorizations(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgGuestsCountDistinctEnum,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgGuestsCountDistinctEnum`](../../doc/models/org-guests-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `auth_method`, `company`, `ssid`<br><br>**Default**: `"auth_method"` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgGuestsCountDistinctEnum_AUTHMETHOD

duration := "10m"

limit := 100

apiResponse, err := orgsGuests.CountOrgGuestAuthorizations(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Guest Authorization

Delete a guest authorization record by guest MAC address.

```go
DeleteOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

resp, err := orgsGuests.DeleteOrgGuestAuthorization(ctx, orgId, guestMac)
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


# Get Org Guest Authorization

Retrieve the guest authorization record associated with a guest MAC address.

```go
GetOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string) (
    models.ApiResponse[models.Guest],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Guest](../../doc/models/guest.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

apiResponse, err := orgsGuests.GetOrgGuestAuthorization(ctx, orgId, guestMac)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Guest Authorizations

List guest authorization records across the organization, including WLAN, SSID, authentication method, expiration, and guest identity details.

```go
ListOrgGuestAuthorizations(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Guest],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Guest](../../doc/models/guest.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsGuests.ListOrgGuestAuthorizations(ctx, orgId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Guest Authorization

Search organization guest authorization records with filters for WLAN, SSID, authentication method, and time range.

```go
SearchOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    wlanId *string,
    authMethod *string,
    ssid *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string) (
    models.ApiResponse[models.ResponseGuestSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wlanId` | `*string` | Query, Optional | Filter results by WLAN identifier |
| `authMethod` | `*string` | Query, Optional | Filter results by authentication method |
| `ssid` | `*string` | Query, Optional | Filter results by SSID |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseGuestSearch](../../doc/models/response-guest-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wlanId := "00000000-0000-0000-0000-000000000000"

authMethod := "psk"

ssid := "Guest"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsGuests.SearchOrgGuestAuthorization(ctx, orgId, &wlanId, &authMethod, &ssid, &limit, nil, nil, &duration, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Guest Authorization

Update the organization guest authorization record for a guest MAC address, including authorization state, duration, WLAN, and guest identity fields.

```go
UpdateOrgGuestAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    guestMac string,
    body *models.GuestOrg) (
    models.ApiResponse[models.Guest],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `guestMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `body` | [`*models.GuestOrg`](../../doc/models/guest-org.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Guest](../../doc/models/guest.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

guestMac := "0000000000ab"

body := models.GuestOrg{
    Company:                models.ToPointer("string"),
    Email:                  models.ToPointer("user@example.com"),
    Mac:                    models.ToPointer("string"),
    Minutes:                models.ToPointer(1440),
    Name:                   models.ToPointer("John Smith"),
    WlanId:                 uuid.MustParse("6748cfa6-4e12-11e6-9188-0242ac110007"),
}

apiResponse, err := orgsGuests.UpdateOrgGuestAuthorization(ctx, orgId, guestMac, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

