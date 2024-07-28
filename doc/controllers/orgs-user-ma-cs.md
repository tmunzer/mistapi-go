# Orgs User MA Cs

```go
orgsUserMACs := client.OrgsUserMACs()
```

## Class Name

`OrgsUserMACs`

## Methods

* [Create Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#create-org-user-macs)
* [Delete Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#delete-org-user-mac)
* [Get Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#get-org-user-mac)
* [Import Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#import-org-user-macs)
* [List Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#list-org-user-macs)
* [Search Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#search-org-user-macs)


# Create Org User Macs

Create Org User MACs

### Usermacs import CSV file format

mac,labels,vlan,notes
921b638445cd,”bldg1,flor1”,vlan-100
721b638445ef,”bldg2,flor2”,vlan-101,Canon Printers
721b638445ee,”bldg3,flor3”,vlan-102
921b638445ce,”bldg4,flor4”,vlan-103
921b638445cf,”bldg5,flor5”,vlan-104

```go
CreateOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UserMac) (
    models.ApiResponse[models.UserMacImport],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UserMac`](../../doc/models/user-mac.md) | Body, Optional | - |

## Response Type

[`models.UserMacImport`](../../doc/models/user-mac-import.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UserMac{
    Labels: []string{
        "byod",
        "flr1",
    },
    Mac:    models.ToPointer("5684dae9ac8b"),
    Notes:  models.ToPointer("mac address refers to Canon printers"),
    Vlan:   models.ToPointer(30),
}

apiResponse, err := orgsUserMACs.CreateOrgUserMacs(ctx, orgId, &body)
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
  "added": [
    "921b638445cd"
  ],
  "errors": [
    "921b638445ce - mac invalid",
    "921b638445cf - mac already provided"
  ],
  "updated": [
    "721b638445ef",
    "721b638445ee"
  ]
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


# Delete Org User Mac

Delete Org User MAC

```go
DeleteOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `usermacId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

usermacId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsUserMACs.DeleteOrgUserMac(ctx, orgId, usermacId)
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


# Get Org User Mac

Get Org User MAC

```go
GetOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID) (
    models.ApiResponse[models.UserMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `usermacId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.UserMac`](../../doc/models/user-mac.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

usermacId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsUserMACs.GetOrgUserMac(ctx, orgId, usermacId)
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
  "id": "111cafd2-ba1b-5169-bfcb-9cdf1d473ddb",
  "labels": [
    "flor1",
    "bld4"
  ],
  "mac": "921b638445cd",
  "notes": "mac address refers to Canon printers",
  "vlan": 30
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


# Import Org User Macs

Import Org User MACs

```go
ImportOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    body []models.UserMac) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.UserMac`](../../doc/models/user-mac.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []models.UserMac{
    models.UserMac{
        Labels: []string{
            "label1",
        },
        Mac:    models.ToPointer("921b638445cd"),
        Vlan:   models.ToPointer(30),
    },
    models.UserMac{
        Labels: []string{
            "label2",
            "label3",
        },
        Mac:    models.ToPointer("721b638445ef"),
        Notes:  models.ToPointer("mac address refers to Canon printers"),
    },
    models.UserMac{
        Labels: []string{
            "label4",
        },
        Mac:    models.ToPointer("721b638445ee"),
    },
    models.UserMac{
        Labels: []string{
            "label5",
            "label6",
            "label7",
        },
        Mac:    models.ToPointer("921b638445ce"),
    },
    models.UserMac{
        Mac:    models.ToPointer("921b638445cf"),
        Vlan:   models.ToPointer(100),
    },
}

resp, err := orgsUserMACs.ImportOrgUserMacs(ctx, orgId, body)
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


# List Org User Macs

List Org User MACs

```go
ListOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    blacklisted *bool,
    forGuestWifi *bool,
    crossSite *bool,
    siteId *string,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.UserMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `blacklisted` | `*bool` | Query, Optional | - |
| `forGuestWifi` | `*bool` | Query, Optional | - |
| `crossSite` | `*bool` | Query, Optional | - |
| `siteId` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.UserMac`](../../doc/models/user-mac.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")









page := 1

limit := 100

apiResponse, err := orgsUserMACs.ListOrgUserMacs(ctx, orgId, nil, nil, nil, nil, &page, &limit)
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
    "id": "111cafd2-ba1b-5169-bfcb-9cdf1d473ddb",
    "labels": [
      "flor1",
      "bld4"
    ],
    "mac": "921b638445cd",
    "notes": "mac address refers to Canon printers",
    "vlan": 30
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org User Macs

Search Org User MACs

```go
SearchOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    labels []string,
    page *int,
    limit *int) (
    models.ApiResponse[models.UserMacsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | mac address |
| `labels` | `[]string` | Query, Optional | optional, array of strings of labels |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`models.UserMacsSearch`](../../doc/models/user-macs-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





page := 1

limit := 100

apiResponse, err := orgsUserMACs.SearchOrgUserMacs(ctx, orgId, nil, nil, &page, &limit)
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
  "limit": 100,
  "page": 1,
  "results": [
    {
      "id": "111cafd2-ba1b-5169-bfcb-9cdf1d473ddb",
      "labels": [
        "flor1",
        "bld4"
      ],
      "mac": "921b638445cd",
      "notes": "mac address refers to Canon printers",
      "vlan": 30
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

