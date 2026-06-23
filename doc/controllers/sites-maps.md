# Sites Maps

```go
sitesMaps := client.SitesMaps()
```

## Class Name

`SitesMaps`

## Methods

* [Add Site Map Image](../../doc/controllers/sites-maps.md#add-site-map-image)
* [Bulk Assign Site Aps to Map](../../doc/controllers/sites-maps.md#bulk-assign-site-aps-to-map)
* [Create Site Map](../../doc/controllers/sites-maps.md#create-site-map)
* [Delete Site Map](../../doc/controllers/sites-maps.md#delete-site-map)
* [Delete Site Map Image](../../doc/controllers/sites-maps.md#delete-site-map-image)
* [Get Site Map](../../doc/controllers/sites-maps.md#get-site-map)
* [Import Site Maps](../../doc/controllers/sites-maps.md#import-site-maps)
* [Import Site Wayfindings](../../doc/controllers/sites-maps.md#import-site-wayfindings)
* [List Site Maps](../../doc/controllers/sites-maps.md#list-site-maps)
* [Replace Site Map Image](../../doc/controllers/sites-maps.md#replace-site-map-image)
* [Start Site Map Auto Geofence](../../doc/controllers/sites-maps.md#start-site-map-auto-geofence)
* [Start Site Maps Auto Geofence](../../doc/controllers/sites-maps.md#start-site-maps-auto-geofence)
* [Update Site Map](../../doc/controllers/sites-maps.md#update-site-map)


# Add Site Map Image

Add image map is a multipart POST which has an file (Image) and an optional json parameter

```go
AddSiteMapImage(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Image file content uploaded as multipart form data |
| `json` | `*string` | Form, Optional | Optional JSON metadata submitted with the image upload |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

resp, err := sitesMaps.AddSiteMapImage(ctx, siteId, mapId, file, nil)
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


# Bulk Assign Site Aps to Map

This API can be used to assign a list of AP Macs associated with site_id to the specified map_id. Note that map_id must be associated with corresponding site_id. This API obeys the following rules

1. if AP is unassigned to any Map, it gets associated with map_id
2. Any moved APs are returned in the response
3. If the AP is considered a locked AP, no action will be taken

```go
BulkAssignSiteApsToMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseSetDevicesMap],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSetDevicesMap](../../doc/models/response-set-devices-map.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "5c5b35000001",
        "5c5b35584a6f",
    },
}

apiResponse, err := sitesMaps.BulkAssignSiteApsToMap(ctx, siteId, mapId, &body)
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
  "locked": [
    "5c5b35584a6f"
  ],
  "moved": [
    "5c5b35000001"
  ]
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


# Create Site Map

Create Site Map

```go
CreateSiteMap(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Map`](../../doc/models/map.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Map](../../doc/models/map.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Map{
    GroupIdx:             models.ToPointer(0),
    GroupName:            models.ToPointer("string"),
    Height:               models.ToPointer(0),
    HeightM:              models.ToPointer(float64(0)),
    LatlngBr:             models.ToPointer(models.LatlngBr{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    LatlngTl:             models.ToPointer(models.LatlngTl{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    Locked:               models.ToPointer(true),
    Name:                 models.ToPointer("string"),
    OccupancyLimit:       models.ToPointer(0),
    Orientation:          models.ToPointer(0),
    OriginX:              models.ToPointer(0),
    OriginY:              models.ToPointer(0),
    Ppm:                  models.ToPointer(float64(0)),
    SitesurveyPath:       []models.MapSitesurveyPathItems{
        models.MapSitesurveyPathItems{
            Coordinate:           models.ToPointer("string"),
            Name:                 models.ToPointer("string"),
            Nodes:                []models.MapNode{
                models.MapNode{
                    Edges:                map[string]string{
                        "N2": "string",
                    },
                    Name:                 "string",
                    Position:             models.ToPointer(models.MapNodePosition{
                        X:                    float64(0),
                        Y:                    float64(0),
                    }),
                },
            },
        },
    },
    Type:                 models.ToPointer(models.MapTypeEnum_IMAGE),
    View:                 models.NewOptional(models.ToPointer(models.MapViewEnum_ROADMAP)),
    WallPath:             models.ToPointer(models.MapWallPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Wayfinding:           models.ToPointer(models.MapWayfinding{
        Micello:              models.ToPointer(models.MapWayfindingMicello{
            AccountKey:           models.ToPointer("string"),
            DefaultLevelId:       models.ToPointer(0),
        }),
        SnapToPath:           models.ToPointer(true),
    }),
    WayfindingPath:       models.ToPointer(models.MapWayfindingPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Width:                models.ToPointer(0),
    WidthM:               models.ToPointer(float64(0)),
}

apiResponse, err := sitesMaps.CreateSiteMap(ctx, siteId, &body)
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


# Delete Site Map

Delete Site Map

```go
DeleteSiteMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMaps.DeleteSiteMap(ctx, siteId, mapId)
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


# Delete Site Map Image

Delete Site Map Image

```go
DeleteSiteMapImage(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMaps.DeleteSiteMapImage(ctx, siteId, mapId)
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


# Get Site Map

Get Site Map Details

```go
GetSiteMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[models.Map],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Map](../../doc/models/map.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesMaps.GetSiteMap(ctx, siteId, mapId)
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


# Import Site Maps

Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches.

# Note

This endpoint (at the site level), the AP must be already assigned to the site to be placed on the floorplan. If you want to place APs from the Org inventory, it is required to use the endpoint at the Org level [importOrgMaps](#operation/importOrgMaps)

# CSV File Format

```csv
Vendor AP name,Mist AP Mac
US Office AP-2,5c:5b:35:00:00:02
US Office AP-3,5c5b35000002
```

```go
ImportSiteMaps(
    ctx context.Context,
    siteId uuid.UUID,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    models.ApiResponse[models.ResponseMapImport],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `autoDeviceprofileAssignment` | `*bool` | Form, Optional | Whether to auto assign device to deviceprofile by name |
| `csv` | `*models.FileWrapper` | Form, Optional | Optional AP name-mapping CSV file |
| `file` | `*models.FileWrapper` | Form, Optional | Ekahau or iBwave floorplan file to import |
| `json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Form, Optional | Options for importing map data from Ekahau or iBwave JSON |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMapImport](../../doc/models/response-map-import.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

autoDeviceprofileAssignment := true

json := models.MapImportJson{
    ImportAllFloorplans:  models.ToPointer(false),
    ImportHeight:         models.ToPointer(true),
    ImportOrientation:    models.ToPointer(true),
    VendorName:           models.MapImportJsonVendorNameEnum_EKAHAU,
}

apiResponse, err := sitesMaps.ImportSiteMaps(ctx, siteId, &autoDeviceprofileAssignment, nil, nil, &json)
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
  "aps": [
    {
      "action": "assigned-placed",
      "floorplan_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
      "height": 3,
      "mac": "5c5b35000001",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "orientation": 45
    }
  ],
  "floorplans": [
    {
      "action": "imported",
      "id": "cbdb7f0b-3be0-4872-88f9-58790b509c23",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "name": "map1"
    }
  ],
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "summary": {
    "num_ap_assigned": 1,
    "num_inv_assigned": 1,
    "num_map_assigned": 1
  }
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


# Import Site Wayfindings

This imports the vendor map meta data into the Map JSON. This is required by the SDK and App in order to access/render the vendor Map properly.

```go
ImportSiteWayfindings(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.WayfindingImportJson) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WayfindingImportJson`](../../doc/models/containers/wayfinding-import-json.md) | Body, Optional | Vendor wayfinding map metadata imported from Jibestream or Micello |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WayfindingImportJsonContainer.FromMapJibestream(models.MapJibestream{
    ClientId:             uuid.MustParse("199d6770-0f6f-407a-9bd5-fc33c7840194"),
    ClientSecret:         "/9Nog3yDzcYj0bY91XJZQLCt+m9DXaIVhx+Ghk3ddd",
    CustomerId:           123,
    EndpointUrl:          "https://api.jibestream.com",
    MapId:                uuid.MustParse("b069b358-4c97-5319-1f8c-7c5ca64d6ab1"),
    Mmpp:                 223,
    Ppm:                  float64(4),
    VendorName:           "jibestream",
    VenueId:              123,
})

resp, err := sitesMaps.ImportSiteWayfindings(ctx, siteId, mapId, &body)
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


# List Site Maps

Get List of Site Maps

```go
ListSiteMaps(
    ctx context.Context,
    siteId uuid.UUID,
    mapstackId *uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Map],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapstackId` | `*uuid.UUID` | Query, Optional | Filter maps by mapstack UUID; returns only maps belonging to that mapstack |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Map](../../doc/models/map.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesMaps.ListSiteMaps(ctx, siteId, nil, &limit, &page)
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
    "flags": {},
    "group_idx": 1,
    "group_name": "East Wing",
    "height": 0,
    "height_m": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "latlng_br": {
      "lat": "string",
      "lng": "string"
    },
    "latlng_tl": {
      "lat": "string",
      "lng": "string"
    },
    "locked": true,
    "modified_time": 0,
    "name": "string",
    "occupancy_limit": 0,
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "orientation": 0,
    "origin_x": 0,
    "origin_y": 0,
    "ppm": 0,
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "sitesurvey_path": [
      {
        "coordinate": "string",
        "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
        "name": "string",
        "nodes": [
          {
            "edges": {
              "N2": "string"
            },
            "name": "string",
            "position": {
              "x": 0,
              "y": 0
            }
          }
        ]
      }
    ],
    "thumbnail_url": "string",
    "type": "image",
    "url": "string",
    "view": "roadmap",
    "wall_path": {
      "coordinate": "string",
      "nodes": [
        {
          "edges": {
            "N2": "string"
          },
          "name": "string",
          "position": {
            "x": 0,
            "y": 0
          }
        }
      ]
    },
    "wayfinding": {
      "micello": {
        "account_key": "string",
        "default_level_id": 0,
        "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      },
      "snap_to_path": true
    },
    "wayfinding_path": {
      "coordinate": "string",
      "nodes": [
        {
          "edges": {
            "N2": "string"
          },
          "name": "string",
          "position": {
            "x": 0,
            "y": 0
          }
        }
      ]
    },
    "width": 0,
    "width_m": 0
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


# Replace Site Map Image

Replace Map Image

This works like an PUT where the image will be replaced. If transform is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)

```go
ReplaceSiteMapImage(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    file models.FileWrapper,
    json *models.MapSiteReplaceFileJson) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Map image file used to replace the existing site map |
| `json` | [`*models.MapSiteReplaceFileJson`](../../doc/models/map-site-replace-file-json.md) | Form, Optional | Options for replacing a site map image |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

resp, err := sitesMaps.ReplaceSiteMapImage(ctx, siteId, mapId, file, nil)
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


# Start Site Map Auto Geofence

The auto geofence service is a map parsing service that uses map image data to identify the exterior of buildings in the map image also known as "geofences". This API processes a single given MapId. This map must have an image to parse for the auto geofence service. Repeated POST requests to this endpoint while the auto geofence service is processing the map will be rejected.

```go
StartSiteMapAutoGeofence(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMaps.StartSiteMapAutoGeofence(ctx, mapId, siteId)
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


# Start Site Maps Auto Geofence

The auto geofence service is a map parsing service that uses map image data to identify the exterior of buildings in the map image also known as "geofences". This API processes all maps for a given SiteId. The maps must have an image to parse for the auto geofence service. Repeated POST requests to this endpoint while the auto geofence service is processing the map will be rejected.

```go
StartSiteMapsAutoGeofence(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesMaps.StartSiteMapsAutoGeofence(ctx, siteId)
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


# Update Site Map

Update Site Map

```go
UpdateSiteMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.Map) (
    models.ApiResponse[models.Map],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Map`](../../doc/models/map.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Map](../../doc/models/map.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Map{
    Height:               models.ToPointer(0),
    HeightM:              models.ToPointer(float64(0)),
    LatlngBr:             models.ToPointer(models.LatlngBr{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    LatlngTl:             models.ToPointer(models.LatlngTl{
        Lat:                  models.ToPointer("string"),
        Lng:                  models.ToPointer("string"),
    }),
    Locked:               models.ToPointer(true),
    Name:                 models.ToPointer("string"),
    OccupancyLimit:       models.ToPointer(0),
    Orientation:          models.ToPointer(0),
    OriginX:              models.ToPointer(0),
    OriginY:              models.ToPointer(0),
    Ppm:                  models.ToPointer(float64(0)),
    SitesurveyPath:       []models.MapSitesurveyPathItems{
        models.MapSitesurveyPathItems{
            Coordinate:           models.ToPointer("string"),
            Name:                 models.ToPointer("string"),
            Nodes:                []models.MapNode{
                models.MapNode{
                    Edges:                map[string]string{
                        "N2": "string",
                    },
                    Name:                 "string",
                    Position:             models.ToPointer(models.MapNodePosition{
                        X:                    float64(0),
                        Y:                    float64(0),
                    }),
                },
            },
        },
    },
    Type:                 models.ToPointer(models.MapTypeEnum_IMAGE),
    View:                 models.NewOptional(models.ToPointer(models.MapViewEnum_ROADMAP)),
    WallPath:             models.ToPointer(models.MapWallPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Wayfinding:           models.ToPointer(models.MapWayfinding{
        Micello:              models.ToPointer(models.MapWayfindingMicello{
            AccountKey:           models.ToPointer("string"),
            DefaultLevelId:       models.ToPointer(0),
        }),
        SnapToPath:           models.ToPointer(true),
    }),
    WayfindingPath:       models.ToPointer(models.MapWayfindingPath{
        Coordinate:           models.ToPointer("string"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "N2": "string",
                },
                Name:                 "string",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(0),
                    Y:                    float64(0),
                }),
            },
        },
    }),
    Width:                models.ToPointer(0),
    WidthM:               models.ToPointer(float64(0)),
}

apiResponse, err := sitesMaps.UpdateSiteMap(ctx, siteId, mapId, &body)
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

