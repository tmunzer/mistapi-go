# Orgs User MA Cs

```go
orgsUserMACs := client.OrgsUserMACs()
```

## Class Name

`OrgsUserMACs`

## Methods

* [Create Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#create-org-user-mac)
* [Delete Org Multiple User Macs](../../doc/controllers/orgs-user-ma-cs.md#delete-org-multiple-user-macs)
* [Delete Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#delete-org-user-mac)
* [Get Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#get-org-user-mac)
* [Import Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#import-org-user-macs)
* [Search Org User Macs](../../doc/controllers/orgs-user-ma-cs.md#search-org-user-macs)
* [Update Org Multiple User Macs](../../doc/controllers/orgs-user-ma-cs.md#update-org-multiple-user-macs)
* [Update Org User Mac](../../doc/controllers/orgs-user-ma-cs.md#update-org-user-mac)


# Create Org User Mac

Create Org User MACs

### Usermacs import CSV file format

mac,labels,vlan,notes
921b638445cd,"bldg1,flor1",vlan-100
721b638445ef,"bldg2,flor2",vlan-101,Canon Printers
721b638445ee,"bldg3,flor3",vlan-102
921b638445ce,"bldg4,flor4",vlan-103
921b638445cf,"bldg5,flor5",vlan-104

```go
CreateOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UserMac) (
    models.ApiResponse[models.UserMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UserMac`](../../doc/models/user-mac.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UserMac](../../doc/models/user-mac.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UserMac{
    Labels:               []string{
        "byod",
        "flr1",
    },
    Mac:                  "5684dae9ac8b",
    Name:                 models.ToPointer("Printer2"),
    Notes:                models.ToPointer("mac address refers to Canon printers"),
    RadiusGroup:          models.ToPointer("VIP"),
    Vlan:                 models.ToPointer("30"),
}

apiResponse, err := orgsUserMACs.CreateOrgUserMac(ctx, orgId, &body)
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
  "vlan": "30"
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


# Delete Org Multiple User Macs

Delete Multiple Org User MACs

```go
DeleteOrgMultipleUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UsermacsId) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UsermacsId`](../../doc/models/usermacs-id.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



resp, err := orgsUserMACs.DeleteOrgMultipleUserMacs(ctx, orgId, nil)
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UserMac](../../doc/models/user-mac.md).

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
  "vlan": "30"
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


# Import Org User Macs

Import Org User MACs

### CSV Import example

```csv
mac,labels,vlan,notes,name,radius_group
921b638445cd,"bldg1,flor1",vlan-100
721b638445ef,"bldg2,flor2",vlan-101,Canon Printers
721b638445ee,"bldg3,flor3",vlan-102,Printer2,VIP
921b638445ce,"bldg4,flor4",vlan-103
921b638445cf,"bldg5,flor5",vlan-104
```

```go
ImportOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    file models.FileWrapper) (
    models.ApiResponse[models.UserMacImport],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | File to upload |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UserMacImport](../../doc/models/user-mac-import.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

apiResponse, err := orgsUserMACs.ImportOrgUserMacs(ctx, orgId, file)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org User Macs

Search Org User MACs

```go
SearchOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    labels []string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.UserMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Partial/full MAC address |
| `labels` | `[]string` | Query, Optional | Optional, array of strings of labels |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.UserMac](../../doc/models/user-mac.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





limit := 100

page := 1

apiResponse, err := orgsUserMACs.SearchOrgUserMacs(ctx, orgId, nil, nil, &limit, &page)
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
    "vlan": "30"
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


# Update Org Multiple User Macs

Update Multiple Org User MACs

```go
UpdateOrgMultipleUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    body []models.UserMac) (
    models.ApiResponse[models.UserMacsUpdate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.UserMac`](../../doc/models/user-mac.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UserMacsUpdate](../../doc/models/user-macs-update.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []models.UserMac{
    models.UserMac{
        Labels:               []string{
            "byod",
            "flr1",
        },
        Mac:                  "5684dae9ac8b",
        Name:                 models.ToPointer("Printer2"),
        Notes:                models.ToPointer("mac address refers to Canon printers"),
        RadiusGroup:          models.ToPointer("VIP"),
        Vlan:                 models.ToPointer("30"),
    },
}

apiResponse, err := orgsUserMACs.UpdateOrgMultipleUserMacs(ctx, orgId, body)
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
  "errors": [
    "2feacc8e-5893-418a-acaa-4d7c1afd01fe - invalid id"
  ],
  "updated": [
    "1041c16c-ca87-4d3f-bb94-b97c5819fc09",
    "a016cc8e-5893-418a-acaa-4d7c1af6ac0f"
  ]
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


# Update Org User Mac

Update Org User MAC

```go
UpdateOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID,
    body *models.UserMac) (
    models.ApiResponse[models.UserMac],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `usermacId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UserMac`](../../doc/models/user-mac.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UserMac](../../doc/models/user-mac.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

usermacId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UserMac{
    Labels:               []string{
        "byod",
        "flr1",
    },
    Mac:                  "5684dae9ac8b",
    Name:                 models.ToPointer("Printer2"),
    Notes:                models.ToPointer("mac address refers to Canon printers"),
    RadiusGroup:          models.ToPointer("VIP"),
    Vlan:                 models.ToPointer("30"),
}

apiResponse, err := orgsUserMACs.UpdateOrgUserMac(ctx, orgId, usermacId, &body)
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
  "vlan": "30"
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

