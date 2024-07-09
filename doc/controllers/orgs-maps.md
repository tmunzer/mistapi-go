# Orgs Maps

```go
orgsMaps := client.OrgsMaps()
```

## Class Name

`OrgsMaps`

## Methods

* [Import Org Map to Site](../../doc/controllers/orgs-maps.md#import-org-map-to-site)
* [Import Org Maps](../../doc/controllers/orgs-maps.md#import-org-maps)


# Import Org Map to Site

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

#### Request

```
"json": a JSON string describing your upload
"file": a binary file
```

```go
ImportOrgMapToSite(
    ctx context.Context,
    orgId uuid.UUID,
    siteName string,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapImportJson) (
    models.ApiResponse[models.ResponseMapImport],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteName` | `string` | Template, Required | - |
| `autoDeviceprofileAssignment` | `*bool` | Form, Optional | whether to auto assign device to deviceprofile by name |
| `csv` | `*models.FileWrapper` | Form, Optional | csv file for ap name mapping, optional |
| `file` | `*models.FileWrapper` | Form, Optional | ekahau or ibwave file |
| `json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Form, Optional | - |

## Response Type

[`models.ResponseMapImport`](../../doc/models/response-map-import.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteName := "site_name8"

autoDeviceprofileAssignment := true





json := models.MapImportJson{
    ImportAllFloorplans: models.ToPointer(false),
    ImportHeight:        models.ToPointer(true),
    ImportOrientation:   models.ToPointer(true),
    VendorName:          models.MapImportJsonVendorNameEnum("ekahau"),
}

apiResponse, err := orgsMaps.ImportOrgMapToSite(ctx, orgId, siteName, &autoDeviceprofileAssignment, nil, nil, &json)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Import Org Maps

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

### CSV File Format

```csv
Vendor AP name,Mist AP Mac
US Office AP-2 - 5c:5b:35:00:00:02,5c5b35000002
```

```go
ImportOrgMaps(
    ctx context.Context,
    orgId uuid.UUID,
    autoDeviceprofileAssignment *bool,
    csv *models.FileWrapper,
    file *models.FileWrapper,
    json *models.MapOrgImportFileJson) (
    models.ApiResponse[models.ResponseMapImport],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `autoDeviceprofileAssignment` | `*bool` | Form, Optional | whether to auto assign device to deviceprofile by name |
| `csv` | `*models.FileWrapper` | Form, Optional | csv file for ap name mapping, optional |
| `file` | `*models.FileWrapper` | Form, Optional | ekahau or ibwave file |
| `json` | [`*models.MapOrgImportFileJson`](../../doc/models/map-org-import-file-json.md) | Form, Optional | - |

## Response Type

[`models.ResponseMapImport`](../../doc/models/response-map-import.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

autoDeviceprofileAssignment := true





json := models.MapOrgImportFileJson{
    ImportAllFloorplans: models.ToPointer(false),
    ImportHeight:        models.ToPointer(true),
    ImportOrientation:   models.ToPointer(true),
    SiteId:              models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    VendorName:          models.MapOrgImportFileJsonVendorNameEnum("ekahau"),
}

apiResponse, err := orgsMaps.ImportOrgMaps(ctx, orgId, &autoDeviceprofileAssignment, nil, nil, &json)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

